package code

import (
	"fmt"
	"math"
	"strings"

	"github.com/huichen/sego"
)

type Seg struct {
}

func (s *Seg) Main() {
	fmt.Println(math.Abs(-1))
	segmenter := new(sego.Segmenter)
	segmenter.LoadDictionary("E:/GOLIB/src/github.com/huichen/sego/data/dictionary.txt")

	//	text := `
	//		Before I went to middle school,
	//		 my parents changed their job,
	//		they found a good job in the big city,
	//		so I needed to leave with them. Moving to another city means I
	//		 have to say goodbye to my friends and this beautiful place. It is hard
	//		 for me to say goodbye, my memories are around here, the water, the mountain and the
	//		 people are so familiar to me. I have to face the new people and new environment, there are
	//		so many challenges waiting for me. But it is time to say goodbye, I need to grow up and be a
	//		strong person. I believe that I can conquer the difficulties.
	//	`
	text := `
	this is the first line it stand for hah.
	this is the second line it stand for heh heh heh heh heh heh.
	zhe shi the third line it stand for 中国是个什么样子的国家.
	this is the fourth line it stand for lala lala heh lala heh lala.
	this is the fifth line it stand for 111 111 111 111 111 111.
`
	//预处理，去掉回车与制表符,并且转为小写。
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	text = strings.ToLower(text)
	//获取总的文档与文档数目，（以及每一个文档的名字，或者作者。这里用index作为作者。）
	docs := strings.Split(text, ".")
	docCount := len(docs)
	fmt.Println("docCount:", docCount)
	//获取每个文档中的所有词,以及出现该词的文档数
	var m = make(map[int]map[string]map[string]float64) //map[index]map[词]文档数
	for i, line := range docs {                         //line 在这里是每一句话
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, ".", "", -1)
		data := []byte(line)
		//对每句话进行分词
		segs := segmenter.Segment(data)
		var mm = make(map[string]float64) //出现文档中每个词的文档数，key是文档中的每个词，value是出现这个词的文档数
		var lineWordCount = 0
		var tfM = make(map[string]float64)
		tmp := make(map[string]string)
		for _, v := range segs {
			if word := strings.Trim(v.Token().Text(), " "); word != "" {
				lineWordCount++
				tfM[word]++
				tmp[word] = word
			}
		}
		for _, v := range tmp {
			word := v
			for _, l := range docs {
				if strings.Contains(l, word) {
					if mm[word] == 0 {
						mm[word] = 1
					}
					mm[word]++
				}
			}
		}
		//计算tf
		for k, _ := range tfM {
			tfM[k] = tfM[k] / ((float64)(lineWordCount))
		}
		tfidfM := make(map[string]map[string]float64)
		tfidfM["idf"] = mm
		tfidfM["tf"] = tfM
		m[i] = tfidfM
	}
	//	计算idf
	for ik, _ := range m {
		for k, _ := range m[ik]["idf"] {
			m[ik]["idf"][k] = math.Log10(((float64)(docCount)) / m[ik]["idf"][k])
		}
	}
	var TF_IDF = make(map[int]map[string]float64)
	for ik, _ := range m {
		tfM := m[ik]["tf"]
		idfM := m[ik]["idf"]
		tfidfM := make(map[string]float64)
		for k, _ := range tfM {
			//			fmt.Println(k, "tf:", tfM[k])
			//			fmt.Println(k, "idf:", idfM[k])
			tfidf := tfM[k] * idfM[k]
			tfidfM[k] = tfidf
		}
		TF_IDF[ik] = tfidfM

		fmt.Println(ik, tfidfM)
	}

}
