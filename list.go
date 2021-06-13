package main

func listReviews(url string) {
	file := load(url)
	for _, r := range file.Reviews {
		printReview(r)
	}
}
