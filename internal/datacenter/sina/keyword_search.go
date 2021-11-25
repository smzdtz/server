// 关键词搜索

package sina

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"smzdtz-server/pkg/http"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// SearchResult 搜索结果
type SearchResult struct {
	// 数字代码
	SecurityCode string
	// 带后缀的代码
	Secucode string
	// 股票名称
	Name string
	// 股市类型: 11=A股 31=港股 41=美股 103=英股
	Market int
}

// KeywordSearch 关键词搜索， 股票、代码、拼音
func (s Sina) KeywordSearch(ctx context.Context, kw string) (results []SearchResult, err error) {
	apiurl := fmt.Sprintf("https://suggest3.sinajs.cn/suggest/key=%s", kw)
	resp, err := http.HTTPGETRaw(ctx, s.HTTPClient, apiurl, nil)

	if err != nil {
		return nil, err
	}

	trans := transform.NewReader(bytes.NewReader(resp), simplifiedchinese.GBK.NewDecoder())
	utf8resp, err := ioutil.ReadAll(trans)

	ds := strings.Split(string(utf8resp), "=")
	if len(ds) != 2 {
		return nil, errors.New("search resp invalid:" + string(utf8resp))
	}
	data := strings.Trim(ds[1], `"`)
	for _, line := range strings.Split(data, ";") {
		lineitems := strings.Split(line, ",")
		if len(lineitems) < 9 {
			continue
		}
		market, _ := strconv.Atoi(lineitems[1])
		secucode := lineitems[3][2:] + "." + lineitems[3][:2]
		result := SearchResult{
			SecurityCode: lineitems[2],
			Secucode:     secucode,
			Name:         lineitems[6],
			Market:       market,
		}
		if market == 11 {
			results = append(results, result)
		}
	}
	// 按股市编号排序确保A股在前面
	sort.Slice(results, func(i, j int) bool {
		return results[i].Market < results[j].Market
	})

	return
}
