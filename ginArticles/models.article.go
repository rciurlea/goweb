package main

type article struct {
	ID      int
	Title   string
	Content string
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	return articleList
}
