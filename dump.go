package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func dumpReviews(url string) {
	reviews := loadRecentReviews(url)

	var file *reviewFile
	if fileExists(url) {
		file = load(url)
	} else {
		file = &reviewFile{
			URL:     url,
			Reviews: reviews,
		}
	}

	newReviews := filterNewReviews(reviews, file.Reviews)

	fmt.Printf("\033[1m%d New Reviews\033[0m\n\n", len(newReviews))
	for _, r := range newReviews {
		file.Reviews = append(file.Reviews, r)
		printReview(r)
	}

	save(file)
}

func filterNewReviews(reviews []*review, savedReviews []*review) []*review {
	newReviews := []*review{}
	for _, r := range reviews {
		exists := false
		for _, rr := range savedReviews {
			if r.User == rr.User && r.Text == rr.Text {
				exists = true
			}
		}
		if !exists {
			newReviews = append(newReviews, r)
		}
	}
	return newReviews
}

func loadRecentReviews(url string) []*review {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url+"#see-all/reviews"),
		chromedp.Sleep(1*time.Second),
	)
	if err != nil {
		panic(err)
	}

	scrollPage(ctx, 0)

	var res []byte
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
			var reviews = document.getElementsByClassName('we-customer-review');
			var json = [];
			for (var i = 0; i < reviews.length; i++) {
				json.push({
					title: reviews[i].getElementsByTagName('h3')[0].textContent.trim(),
					text: reviews[i].getElementsByClassName('we-clamp')[0].textContent.trim(),
					date: Date.parse(reviews[i].getElementsByClassName('we-customer-review__date')[0].textContent.trim()),
					dateString: reviews[i].getElementsByClassName('we-customer-review__date')[0].textContent.trim(),
					user: reviews[i].getElementsByClassName('we-customer-review__user')[0].textContent.trim(),
					rating: parseInt(reviews[i].getElementsByClassName('we-star-rating')[0].getAttribute('aria-label').trim().split(' ')[0])
				});
			}
			return JSON.stringify(json);
		})()
		`, &res),
	)
	if err != nil {
		panic(err)
	}

	s, _ := strconv.Unquote(string(res))

	reviews := []*review{}
	json.Unmarshal([]byte(s), &reviews)

	return reviews
}

func scrollPage(ctx context.Context, lastScrollHeight int) {
	var scrollHeight []byte
	err := chromedp.Run(ctx,
		chromedp.Evaluate(`(() => { document.documentElement.scrollTop = document.documentElement.scrollHeight; })()`, nil),
		chromedp.Sleep(1*time.Second),
		chromedp.Evaluate(`document.documentElement.scrollHeight`, &scrollHeight),
	)
	if err != nil {
		panic(err)
	}
	scrollHeightInt, _ := strconv.Atoi(string(scrollHeight))
	if scrollHeightInt != lastScrollHeight {
		scrollPage(ctx, scrollHeightInt)
	}
}
