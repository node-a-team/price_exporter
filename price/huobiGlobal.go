package price

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	"strings"
	"strconv"

	"go.uber.org/zap"

	cfg "github.com/node-a-team/price_exporter/config"
)

func (ps *PriceService) HuobiGlobal(log *zap.Logger, paymentCurrency string, currency string) {


	for {
		func() {

			defer func() {
				if r := recover(); r != nil {
				// log
				}

				time.Sleep(cfg.Config.Options.Interval * time.Second)
			}()


			changeCurrency := currency +paymentCurrency

			resp, err := http.Get(cfg.Config.APIs.HuobiGlobal +changeCurrency)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to fetch from HuobiGlobal_" +fmt.Sprint(err)))
                        }


			defer func() {
				resp.Body.Close()
			}()

			body, err := ioutil.ReadAll(resp.Body)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to read body from Binance_" +fmt.Sprint(err)))
			}

			hm := huobiMarket{}
			json.Unmarshal(body, &hm)


			price := hm.Tick.Data[0].Price
			priceString := strconv.FormatFloat(price, 'f', -1, 64)

			log.Info("Price", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("HuobiGlobal_" +strings.ToUpper(currency) +" to " +strings.ToUpper(paymentCurrency), priceString))

			ps.SetPrice(currency +"/" +paymentCurrency +"/huobiGlobal", price)
		}()
	}
}


