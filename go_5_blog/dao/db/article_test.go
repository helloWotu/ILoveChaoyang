package db

import (
	"go_5_blog/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	//root为数据库用户名，后面为密码，tcp代表tcp协议，blogger处填写自己的数据库名称
	dns := "root:tutu55201@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 测插入文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 4
	article.ArticleInfo.CommentCount = 0
	article.Content = `犁头入土真锋利，
先到南面去耕地。
百谷种子播田头，
粒粒孕育富生机。
有人送饭来看你，
挑着方筐和圆篓，
里面装的是黍米。
头戴手编草斗笠，
手持锄头来翻土，
除草田畦得清理。
野草腐烂作肥料，
庄稼生长真茂密。
挥镰收割响声齐，
打下谷子高堆起。
看那高处似城墙，
看那两旁似梳齿，
粮仓成百开不闭。
各个粮仓都装满，
妇女儿童心神怡。
杀头黑唇大黄牛，
弯弯双角真美丽。
不断祭祀后续前，
继承古人的礼仪。`
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "良耜"
	article.ArticleInfo.Username = "sun"
	article.ArticleInfo.Summary = "犁头入土真锋利"
	article.ArticleInfo.ViewCount = 1001
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId : %d\n", articleId)
}

func TestGetAricleList(t *testing.T) {
	articleList, err := GetAricleList(1, 15)
	if err != nil {
		return
	}
	t.Logf("article : %d\n", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	ar,err := GetArticleDetail(1)
	if err != nil {
		return
	}
	t.Logf("article %#v\n", ar)
}
