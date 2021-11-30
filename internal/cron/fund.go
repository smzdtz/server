package cron

import "fmt"

// SyncFund 同步基金数据
func SyncFund() {
	fmt.Printf("hello")
	// if !util.IsTradingDay() {
	// 	return
	// }
	// ctx := context.Background()
	// // logging.Infof(ctx, "SyncFund request start...")

	// // 获取全量列表
	// efundlist, err := datacenter.EastMoney.QueryAllFundList(ctx, eastmoney.FundTypeALL)
	// if err != nil {
	// 	// logging.Error(ctx, "SyncFund QueryAllFundList error:"+err.Error())
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// 	return
	// }

	// fundCodes := []string{}
	// for _, efund := range efundlist {
	// 	fundCodes = append(fundCodes, efund.Fcode)
	// }
	// s := core.NewSearcher(ctx)
	// data, err := s.SearchFunds(ctx, fundCodes)
	// fundlist := service.FundList{}
	// typeMap := map[string]struct{}{}
	// for _, fund := range data {
	// 	fundlist = append(fundlist, fund)
	// 	typeMap[fund.Type] = struct{}{}
	// }

	// 更新 services 变量
	// models.FundAllList = fundlist
	// fundtypes := []string{}
	// for k := range typeMap {
	// 	fundtypes = append(fundtypes, k)
	// }
	// models.FundTypeList = fundtypes

	// 更新文件
	// b, err := json.Marshal(efundlist)
	// if err != nil {
	// 	logging.Errorf(ctx, "SyncFund json marshal efundlist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// } else if err := ioutil.WriteFile(models.RawFundAllListFilename, b, 0666); err != nil {
	// 	logging.Errorf(ctx, "SyncFund WriteFile efundlist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// }
	// b, err = json.Marshal(fundlist)
	// if err != nil {
	// 	logging.Errorf(ctx, "SyncFund json marshal fundlist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// } else if err := ioutil.WriteFile(models.FundAllListFilename, b, 0666); err != nil {
	// 	logging.Errorf(ctx, "SyncFund WriteFile fundlist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// }
	// b, err = json.Marshal(models.FundTypeList)
	// if err != nil {
	// 	logging.Errorf(ctx, "SyncFund json marshal fundtypelist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// } else if err := ioutil.WriteFile(models.FundTypeListFilename, b, 0666); err != nil {
	// 	logging.Errorf(ctx, "SyncFund WriteFile fundtypelist error:%v", err)
	// 	promSyncError.WithLabelValues("SyncFund").Inc()
	// }

	// 更新4433列表
	// Update4433()

	// 更新同步时间
	// models.SyncFundTime = time.Now()
}
