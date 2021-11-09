// #!/usr/bin/env python
// # -*- coding:utf-8 -*-
// """
// Date: 2021/1/15 16:59
// Desc: 个股新闻数据
// http://so.eastmoney.com/news/s?keyword=%E4%B8%AD%E5%9B%BD%E4%BA%BA%E5%AF%BF&pageindex=1&searchrange=8192&sortfiled=4
// """
// import pandas as pd
// import requests

// def stock_news_em(stock: str = "601628") -> pd.DataFrame:
//     """
//     东方财富-个股新闻-最近 20 条新闻
//     http://so.eastmoney.com/news/s?keyword=%E4%B8%AD%E5%9B%BD%E4%BA%BA%E5%AF%BF&pageindex=1&searchrange=8192&sortfiled=4
//     :param stock: 股票代码
//     :type stock: str
//     :return: 个股新闻
//     :rtype: pandas.DataFrame
//     """
//     url = "http://searchapi.eastmoney.com//bussiness/Web/GetCMSSearchList"
//     params = {
//         "type": "8196",
//         "pageindex": "1",
//         "pagesize": "20",
//         "keyword": f"({stock})()",
//         "name": "zixun",
//         "_": "1608800267874",
//     }
//     headers = {
// "Accept": "*/*",
// "Accept-Encoding": "gzip, deflate",
// "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
// "Cache-Control": "no-cache",
// "Connection": "keep-alive",
// "Host": "searchapi.eastmoney.com",
// "Pragma": "no-cache",
// "Referer": "http://so.eastmoney.com/",
// "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
//     }

//     r = requests.get(url, params=params, headers=headers)
//     data_json = r.json()
//     temp_df = pd.DataFrame(data_json["Data"])
//     temp_df.columns = [
//         "url",
//         "title",
//         "_",
//         "public_time",
//         "content",
//     ]
//     temp_df['code'] = stock
//     temp_df = temp_df[
//         [
//             "code",
//             "title",
//             "content",
//             "public_time",
//             "url",
//         ]
//     ]
//     temp_df["title"] = (
//         temp_df["title"].str.replace(r"\(<em>", "", regex=True).str.replace(r"</em>\)", "", regex=True)
//     )
//     temp_df["content"] = (
//         temp_df["content"].str.replace(r"\(<em>", "", regex=True).str.replace(r"</em>\)", "", regex=True)
//     )
//     temp_df["content"] = (
//         temp_df["content"].str.replace(r"<em>", "", regex=True).str.replace(r"</em>", "", regex=True)
//     )
//     temp_df["content"] = temp_df["content"].str.replace(r"\u3000", "", regex=True)
//     temp_df["content"] = temp_df["content"].str.replace(r"\r\n", " ", regex=True)
//     return temp_df

// if __name__ == "__main__":
//     stock_news_em_df = stock_news_em(stock="601318")
//     print(stock_news_em_df.info())

// {
//     "IsSuccess": true,
//     "Code": 0,
//     "Message": "成功",
//     "TotalPage": 183,
//     "TotalCount": 3659,
//     "Keyword": "1.601628",
//     "Data": [
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111092174321413.html",
//             "Art_Title": "中国人寿：融资净偿还152.4万元，融资余额12.02亿元（11-08）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111092174321413.html",
//             "Art_CreateTime": "2021-11-09 07:58:53",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月8日融资净偿还152.4万元；融资余额12.02亿元，较前一日下降0.13%融资方面，当日融资买入1855.85万元，融资偿还2008.25万元，融资净偿还152.4万元。融券方面，融券卖出6.27万股，融券偿还13.17万股，融券余量1..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111092174151796.html",
//             "Art_Title": "中国人寿11月08日被沪股通减持33.47万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111092174151796.html",
//             "Art_CreateTime": "2021-11-09 07:39:18",
//             "Art_Content": "11月08日，中国人寿被沪股通减持33.47万股，已连续4日被沪股通减持，共计179.82万股，最新持股量为5357.62万股，占公司A股总股本的0.26%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111082173320660.html",
//             "Art_Title": "国寿等2家公司上午10:16异动",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111082173320660.html",
//             "Art_CreateTime": "2021-11-08 10:54:47",
//             "Art_Content": "　　2家上市公司于10:16am录得异动纪录：上市公司现价变幅股份状态 　　------------------------------------------　　国寿 13.820元 +3.13% 主动买沽比率75:25 　　(02628) 　　PACIFIC LEGEND 　..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111062172298673.html",
//             "Art_Title": "中国人寿：融资净偿还972.94万元，融资余额12.03亿元（11-05）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111062172298673.html",
//             "Art_CreateTime": "2021-11-06 07:39:57",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月5日融资净偿还972.94万元；融资余额12.03亿元，较前一日下降0.8%融资方面，当日融资买入1837.55万元，融资偿还2810.49万元，融资净偿还972.94万元。融券方面，融券卖出14.35万股，融券偿还16.3万股，融券余量..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111062172289271.html",
//             "Art_Title": "中国人寿11月05日被沪股通减持47.28万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111062172289271.html",
//             "Art_CreateTime": "2021-11-06 07:39:10",
//             "Art_Content": "11月05日，中国人寿被沪股通减持47.28万股，已连续3日被沪股通减持，共计146.35万股，最新持股量为5391.08万股，占公司A股总股本的0.26%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111052170818090.html",
//             "Art_Title": "中国人寿：融资净买入221.51万元，融资余额12.13亿元（11-04）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111052170818090.html",
//             "Art_CreateTime": "2021-11-05 07:56:24",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月4日融资净买入221.51万元；融资余额12.13亿元，较前一日增加0.18%融资方面，当日融资买入3503.39万元，融资偿还3281.88万元，融资净买入221.51万元。融券方面，融券卖出23.82万股，融券偿还16.99万股，融券..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111052170636647.html",
//             "Art_Title": "中国人寿11月04日被沪股通减持75.43万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111052170636647.html",
//             "Art_CreateTime": "2021-11-05 07:35:42",
//             "Art_Content": "11月04日，中国人寿被沪股通减持75.43万股，最新持股量为5438.36万股，占公司A股总股本的0.26%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数据生产，仅供参考，不构成任何投资建议，据此操作..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111042170289238.html",
//             "Art_Title": "存量投资规模已达1.2万亿！中国人寿将强化综合金融优势服务区域高质量发展",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111042170289238.html",
//             "Art_CreateTime": "2021-11-04 16:41:55",
//             "Art_Content": "　　近日，中国人寿发布《关于服务国家区域发展战略的指导意见》(以下简称《意见》)，就落实落细新时期区域发展的一系列重大战略部署，服务优化区域经济布局，促进区域协调发展作出一系列工作部署。《意见》立足新发展阶段，贯彻新发展理念，发挥综合金融优势，针对不同区域发展重点提出针对性服务举..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111042170045885.html",
//             "Art_Title": "中国人寿寿险公司荣获第三届新财富最佳上市公司奖",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111042170045885.html",
//             "Art_CreateTime": "2021-11-04 14:11:10",
//             "Art_Content": "　　近日，由新财富牵头举办的第三届新财富最佳上市公司评选结果揭晓，中国人寿保险股份有限公司(以下简称“中国人寿寿险公司”)荣获“新财富最佳上市公司”奖。本次获奖是业界对中国人寿寿险公司上市十余年来持之以恒的市值经营工作的认可与鼓励。　　新财富最佳上市公司评选由新财富联合光华-罗特..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111042169232620.html",
//             "Art_Title": "中国人寿：融资净偿还1066.28万元，融资余额12.11亿元（11-03）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111042169232620.html",
//             "Art_CreateTime": "2021-11-04 07:57:02",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月3日融资净偿还1066.28万元；融资余额12.11亿元，较前一日下降0.87%融资方面，当日融资买入1405.77万元，融资偿还2472.05万元，融资净偿还1066.28万元。融券方面，融券卖出12.52万股，融券偿还23.7万股，融..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111042169078242.html",
//             "Art_Title": "中国人寿11月03日被沪股通减持23.64万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111042169078242.html",
//             "Art_CreateTime": "2021-11-04 07:40:01",
//             "Art_Content": "11月03日，中国人寿被沪股通减持23.64万股，最新持股量为5513.8万股，占公司A股总股本的0.26%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数据生产，仅供参考，不构成任何投资建议，据此操作风..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111032168985690.html",
//             "Art_Title": "存量投资规模达1.2万亿元 中国人寿发布服务国家区域发展战略指导意见",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111032168985690.html",
//             "Art_CreateTime": "2021-11-04 07:05:55",
//             "Art_Content": "　　记者获悉，中国人寿近日发布《关于服务国家区域发展战略的指导意见》(下称《意见》)，就落实落细新时期区域发展、优化区域经济布局、促进区域协调发展作出一系列工作部署。\r\n 　　据了解，《意见》立足新发展阶段，针对不同区域发展重点提出针对性服务举措，助力构建要素有序自由流动、主体功..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111032168906416.html",
//             "Art_Title": "中国人寿支持国家重点区域发展存量投资规模达1.2万亿元",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111032168906416.html",
//             "Art_CreateTime": "2021-11-03 19:48:44",
//             "Art_Content": "　　11月3日，记者从中国人寿获悉，截至2021年上半年，中国人寿支持国家重点区域发展存量投资规模已达1.2万亿元，投向地区各类债券及地方交通、棚改、地铁等重大民生工程建设，深度融入区域协调发展。此外，中国人寿承保的社保补充医疗保险覆盖19个省份约4710万人，长期护理保险累计覆..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111032168834161.html",
//             "Art_Title": "国寿资产完成发行“中国人寿-翡翠1 号资产支持计划”",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111032168834161.html",
//             "Art_CreateTime": "2021-11-03 18:23:12",
//             "Art_Content": "　　证券时报网讯，11月3日，国寿资产发布消息称，由其发起设立的“中国人寿-翡翠1号资产支持计划”产品近日顺利完成发行。该资产支持计划以原始权益人租赁债权为基础资产，募集中国人寿系统内资金和系统外三方非金融机构资金。(刘敬元)（文章来源：证券时报）"
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111032168551334.html",
//             "Art_Title": "中国人寿支持国家重点区域发展存量投资规模达1.2万亿",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111032168551334.html",
//             "Art_CreateTime": "2021-11-03 15:26:43",
//             "Art_Content": "　　近日，中国人寿发布《关于服务国家区域发展战略的指导意见》(以下简称《意见》)，就落实落细新时期区域发展的一系列重大战略部署，服务优化区域经济布局，促进区域协调发展作出一系列工作部署。\r\n 　　《意见》明确，到2025年，中国人寿将初步建成高质量服务区域发展的金融保险服务体系，..."
//         },
//         {
//             "Art_UniqueUrl": "http://finance.eastmoney.com/a/202111032168019117.html",
//             "Art_Title": "存量投资规模已达1.2万亿！中国人寿服务国家区域发展战略驰而不息",
//             "Art_Url": "http://finance.eastmoney.com/news/1354,202111032168019117.html",
//             "Art_CreateTime": "2021-11-03 10:23:44",
//             "Art_Content": "　　近日，中国人寿发布《关于服务国家区域发展战略的指导意见》(以下简称《意见》)，就落实落细新时期区域发展的一系列重大战略部署，服务优化区域经济布局，促进区域协调发展作出一系列工作部署。《意见》立足新发展阶段，贯彻新发展理念，发挥综合金融优势，针对不同区域发展重点提出针对性服务举..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111032167528394.html",
//             "Art_Title": "中国人寿：连续3日融资净买入累计4349.96万元（11-02）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111032167528394.html",
//             "Art_CreateTime": "2021-11-03 07:53:49",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月2日融资净买入736.8万元；融资余额12.22亿元，较前一日增加0.61%融资方面，当日融资买入4339.72万元，融资偿还3602.92万元，融资净买入736.8万元，连续3日净买入累计4349.96万元。融券方面，融券卖出27.26..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111032167356968.html",
//             "Art_Title": "中国人寿11月02日获沪股通增持36.75万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111032167356968.html",
//             "Art_CreateTime": "2021-11-03 07:34:11",
//             "Art_Content": "11月02日，中国人寿获沪股通增持36.75万股，最新持股量为5537.43万股，占公司A股总股本的0.27%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数据生产，仅供参考，不构成任何投资建议，据此操作..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111022165792056.html",
//             "Art_Title": "中国人寿：融资净买入59.61万元，融资余额12.14亿元（11-01）",
//             "Art_Url": "http://stock.eastmoney.com/news/1697,202111022165792056.html",
//             "Art_CreateTime": "2021-11-02 07:57:03",
//             "Art_Content": "中国人寿融资融券信息显示，2021年11月1日融资净买入59.61万元；融资余额12.14亿元，较前一日增加0.05%融资方面，当日融资买入3019.15万元，融资偿还2959.54万元，融资净买入59.61万元。融券方面，融券卖出16.66万股，融券偿还23.66万股，融券余量..."
//         },
//         {
//             "Art_UniqueUrl": "http://stock.eastmoney.com/a/202111022165598758.html",
//             "Art_Title": "中国人寿11月01日被沪股通减持10.4万股",
//             "Art_Url": "http://stock.eastmoney.com/news/11072,202111022165598758.html",
//             "Art_CreateTime": "2021-11-02 07:35:14",
//             "Art_Content": "11月01日，中国人寿被沪股通减持10.4万股，最新持股量为5500.68万股，占公司A股总股本的0.26%。近五日持股量数据持股量及股价变动图注：文中持股数量经过前复权处理，可能存在与港交所披露不一致的情形。免责声明：本文基于大数据生产，仅供参考，不构成任何投资建议，据此操作风..."
//         }
//     ],
//     "RelatedWord": "",
//     "StillSearch": [
//         "中国人寿",
//         "601628"
//     ],
//     "StockModel": {
//         "Name": "中国人寿",
//         "Code": "601628"
//     }
// }
