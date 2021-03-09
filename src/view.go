package main

import (
	"strconv"
)

func (vi *View) init(){
	vi.db.AutoMigrate(&Category{},&Article{},&Sentence{},&Link{},&Bacteria{},&Trend{})
}

func (vi *View) add_article(doi,title,cate_name string) *Article {
	art := Article{DOI:doi}
	cate := Category{Name:cate_name}
	if vi.db.Where(&art).First(&art).NewRecord(&art) {
		art.Title = title
		vi.db.Create(&art)
	}
	if vi.db.Where(&cate).First(&cate).NewRecord(&cate) {
		vi.db.Create(&cate)
	}
	cate.Articles = append(cate.Articles,art)
	vi.db.Save(&cate).Preload("Sentences").First(&art)
	return &art
}

func (vi *View) get_articles(art Article) *[]Article {
	arts := []Article{}
	vi.db.Where(&art).Find(&arts)
	return &arts
}

func (vi *View) id2cate() *map[uint]string {
	res := map[uint]string{}
	cates := []Category{}
	vi.db.Where(Category{}).Find(&cates)
	for _, obj := range cates {
		res[obj.ID] = obj.Name
	}
	return &res
}
/*
func (vi *View) get_articles(art Article) *[][4]string {
	res := [][4]string{}
	rows,_ := vi.db.Table("article").Select("article.id, article.doi, article.title, category.name").Joins("join category on article.category_id = category.id").Rows()
	defer rows.Close()
	for rows.Next() {
		var id string
		var doi string
		var title string
		var cate string
		if err := rows.Scan(&id, &doi, &title, &cate); err != nil {
			log.Fatal(err)
		}
		t_res := [4]string{id,doi,title,cate}
		res = append(res,t_res)
	}
	//arts := []Article{}
	//vi.db.Where(&art).Preload("Category").Find(&arts)
	//res := map[string]interface{}{"arts":arts,}
	return &res
}
*/
func (vi *View) add_sentence(art_id,syndrome,part string) *Sentence {
	a_id,_ := strconv.Atoi(art_id)
	sente := Sentence{Syndrome:syndrome,Part:part,ArticleID:uint(a_id)}
	vi.db.Save(&sente).Preload("Links").First(&sente)
	return &sente
}

func (vi *View) add_link(sente_id,bacteria_name,trend_mark string) *Link {
	bacte := Bacteria{Name:bacteria_name}
	trend := Trend{Mark:trend_mark}
	for _, obj := range [2]interface{}{&bacte,&trend} {
		if vi.db.Where(obj).First(obj).NewRecord(obj) {
			vi.db.Create(obj)
		}
	}
	s_id,_ := strconv.Atoi(sente_id)
	link := Link{SentenceID:uint(s_id),BacteriaID:bacte.ID,TrendID:trend.ID}
	vi.db.Save(&link)
	return &link
}
