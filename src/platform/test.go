package platform

import (
	"pisces_i/src/platform/examination"

	"github.com/kataras/iris"
)

// ExaminationPaper ...
type ExaminationPaper struct {
	Title    string
	SubTitle string
	Content  string
	examination.Examination
}

//VocabularyTest ...
func VocabularyTest(ctx *iris.Context) {
	ctx.Render("chooseVocabularyTest.html", nil)
}

// ChooseVocabularyTest ...
func ChooseVocabularyTest(ctx *iris.Context) {
	testType := ctx.Param("test")

	var e ExaminationPaper

	switch testType {
	case "en":
		e.Title = "英译中单词测试"
		e.SubTitle = "Translate english words into chinese"
		e.Content = "本测试会在所有雅思词汇中随机选择20个单词作为测试题。请为单词选择所对应的中文释义。"
		e.Examination = examination.GenerateEnToCnExamination()
	case "cn":
		e.Title = "中译英单词测试"
		e.SubTitle = "Translate chinese words into english"
		e.Content = "本测试会在所有雅思词汇中随机选择20个单词作为测试题。请为单词选择所对应的英文释义。"
		e.Examination = examination.GenerateCnToEnExamination()
	default:
		// fixme: should redirect to not found page.
	}

	ctx.Render("vocabularyTest.html", e)
}

//EnVocabularyTest ...
func EnVocabularyTest(ctx *iris.Context) {
	var e ExaminationPaper
	e.Title = "英译中单词测试"
	e.SubTitle = "Translate english words into chinese"
	e.Content = "本测试会在所有雅思词汇中随机选择20个单词作为测试题。请为单词选择所对应的中文释义。"
	e.Examination = examination.GenerateEnToCnExamination()

	ctx.Render("vocabularyTest.html", e)
}

//CnVocabularyTest ...
func CnVocabularyTest(ctx *iris.Context) {
	var e ExaminationPaper
	e.Title = "中译英单词测试"
	e.SubTitle = "Translate chinese words into english"
	e.Content = "本测试会在所有雅思词汇中随机选择20个单词作为测试题。请为单词选择所对应的英文释义。"
	e.Examination = examination.GenerateCnToEnExamination()

	ctx.Render("vocabularyTest.html", e)
}
