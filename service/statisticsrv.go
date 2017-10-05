package service

import (
	"strconv"
	"sync"

	"github.com/b3log/solo.go/model"
)

var Statistic = &statisticService{
	mutex: &sync.Mutex{},
}

type statisticService struct {
	mutex *sync.Mutex
}

func (srv *statisticService) GetStatistic(statisticName string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if nil != db.Where("name = ? AND category = ? AND blog_id = ?", statisticName, model.SettingCategoryStatistic, blogID).Find(ret).Error {
		return nil
	}

	return ret
}

func (srv *statisticService) GetStatistics(blogID uint, statisticNames ...string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	settings := []*model.Setting{}
	if nil != db.Where("name IN (?) AND category = ? AND blog_id = ?", statisticNames, model.SettingCategoryStatistic, blogID).Find(&settings).Error {
		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}

func (srv *statisticService) IncArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.IncArticleCountWithoutTx(blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncArticleCountWithoutTx(blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := db.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := db.Model(&model.Setting{}).Update(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.DecArticleCountWithoutTx(blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecArticleCountWithoutTx(blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := db.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count - 1)
	if err := db.Model(&model.Setting{}).Update(setting).Error; nil != err {
		return err
	}

	return nil
}