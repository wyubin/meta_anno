package main

type Category struct {
	ID uint
	Name string `gorm:"type:varchar(100);unique_index"`
	Articles []Article
}

type Article struct {
	ID uint
	CategoryID uint
	DOI string `gorm:"type:text;unique_index"`
	Title string `gorm:"type:text"`
	Sentences []Sentence
}

type Sentence struct {
	ID uint
	ArticleID uint
	Syndrome string `gorm:"type:text"`
	Part string `gorm:"type:text"`
	Links []Link
}

type Link struct {
	ID uint
	SentenceID uint
	BacteriaID uint
	TrendID uint
}

type Bacteria struct {
	ID uint
	Name string `gorm:"type:varchar(100);unique_index"`
	Links []Link
}

type Trend struct {
	ID uint
	Mark string `gorm:"type:varchar(100);unique_index"`
	Links []Link
}
