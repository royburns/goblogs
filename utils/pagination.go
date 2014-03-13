package utils

import (
	"errors"
	"math"
	"strconv"
)

type Config struct {
	PageUrl       string
	PageSize      int
	PageNum       int
	RowsNum       int
	LinksNum      int
	AnchorClass   string
	CurrentClass  string
	FullTagOpen   string
	FullTagClose  string
	InfoTagOpen   string
	InfoTagClose  string
	FirstTagOpen  string
	FirstTagClose string
	LastTagOpen   string
	LastTagClose  string
	CurTagOpen    string
	CurTagClose   string
	NextTagOpen   string
	NextTagClose  string
	PrevTagOpen   string
	PrevTagClose  string
	NumTagOpen    string
	NumTagClose   string
}

// create page links
func CreateLinks(conf Config) (string, error) {
	page_url := conf.PageUrl
	rows_num := conf.RowsNum
	page_size := conf.PageSize
	page_num := conf.PageNum
	anchor_class := conf.AnchorClass
	current_class := conf.CurrentClass
	links_num := conf.LinksNum
	full_tag_open := conf.FullTagOpen
	full_tag_close := conf.FullTagClose
	first_tag_open := conf.FirstTagOpen
	first_tag_close := conf.FirstTagClose
	prev_tag_open := conf.PrevTagOpen
	prev_tag_close := conf.PrevTagClose
	cur_tag_open := conf.CurTagOpen
	cur_tag_close := conf.CurTagClose
	num_tag_open := conf.NumTagOpen
	num_tag_close := conf.NumTagClose
	next_tag_open := conf.NextTagOpen
	next_tag_close := conf.NextTagClose
	last_tag_open := conf.LastTagOpen
	last_tag_close := conf.LastTagClose

	var pageStr string
	if rows_num == 0 || page_size == 0 {
		pageStr = ""
		return pageStr, errors.New("PageSize or RowsNum is not zero")
	}

	pages := int(math.Ceil(float64(rows_num) / float64(page_size)))

	if page_num < 1 {
		page_num = 1
	}

	if page_num > pages {
		pageStr = "404"
		return pageStr, errors.New("PageNum Must Less than or equal to Pages")
	}

	if anchor_class != "" {
		anchor_class = "class=\"" + anchor_class + "\" "
	}

	if current_class != "" {
		current_class = "class=\"" + current_class + "\" "
	}

	if pages == 1 {
		pageStr = ""
		return pageStr, nil
	}

	if links_num < 1 {
		pageStr = ""
		return pageStr, errors.New("LinksNum Must be more than equal to one")
	}

	pageStr = full_tag_open
	//first page
	if page_num > 1 {
		pageStr += first_tag_open + "<a " + anchor_class + " href=\"" + page_url + "?page=1\" >首页</a>" + first_tag_close + "\n"
	}
	//prev page
	if page_num > 1 {
		n := page_num - 1
		pageStr += prev_tag_open + "<a " + anchor_class + " href=\"" + page_url + "?page=" + strconv.Itoa(n) + "\" >上一页</a>" + prev_tag_close + "\n"
	}
	//pages
	for i := 1; i <= pages; i++ {
		var pl int
		var pr int
		if page_num-links_num < 0 {
			pl = 0
		} else {
			pl = page_num - links_num
		}

		if page_num+links_num > pages {
			pr = pages
		} else {
			pr = page_num + links_num
		}

		if pr < 2*links_num+1 {
			pr = 2*links_num + 1
		}
		if pl > pages-2*links_num {
			pl = pages - 2*links_num
		}
		//Current page
		if i == page_num {
			pageStr += cur_tag_open + "<span" + current_class + ">" + strconv.Itoa(i) + "</span>" + cur_tag_close + "\n"
		} else if i >= pl && i <= pr {
			pageStr += num_tag_open + "<a " + anchor_class + " href=\"" + page_url + "?page=" + strconv.Itoa(i) + "\" >" + strconv.Itoa(i) + "</a>" + num_tag_close + "\n"
		}
	}

	//next page
	if page_num < pages {
		n := page_num + 1
		pageStr += next_tag_open + "<a " + anchor_class + " href=\"" + page_url + "?page=" + strconv.Itoa(n) + "\" >下一页</a>" + next_tag_close + "\n"
	}
	//end page
	if page_num < pages {
		pageStr += last_tag_open + "<a " + anchor_class + " href=\"" + page_url + "?page=" + strconv.Itoa(pages) + "\" >末页</a>" + last_tag_close
	}

	pageStr += full_tag_close
	return pageStr, nil
}
