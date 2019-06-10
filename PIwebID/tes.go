func (d *RestDG) getMockOilGasStats() error {

	deliveryRealTime := make([][]influxdb.BaseData, 2)
	nowHourPlus := time.Now().Hour() + 1

	for index, _ := range deliveryRealTime {
		switch index {
		case 0:
			for i := 0; i < nowHourPlus; i++ {
				deliveryRealTime[index] = append(deliveryRealTime[index],
					influxdb.BaseData{
						Name:  utils.Get24HoursFromZero(i).Format("15:04"),
						Value: utils.FormatFloatDigit(0.0554 * (float64(i) + rand.Float64())),
					})
			}
		case 1:
			for i := 0; i < nowHourPlus; i++ {
				deliveryRealTime[index] = append(deliveryRealTime[index],
					influxdb.BaseData{
						Name:  utils.Get24HoursFromZero(i).Format("15:04"),
						Value: utils.FormatFloatDigit(0.0075 * (float64(i) + rand.Float64())),
					})
			}
		}
		// for i := 0; i < nowHourPlus; i++ {
		// 	deliveryRealTime[index] = append(deliveryRealTime[index],
		// 		influxdb.BaseData{
		// 			Name:  utils.Get24HoursFromZero(i).Format("15:04"),
		// 			Value: utils.FormatFloatDigit((2+rand.Float64()/12)*float64(i) + rand.Float64()),
		// 		})
		// }
	}
	zyGasToday := utils.FormatFloatDigit(0.0554 * (float64(nowHourPlus) + rand.Float64()))
	zmGasToday := utils.FormatFloatDigit(0.0075 * (float64(nowHourPlus) + rand.Float64()))
	deliveryRealTimeByte, err := json.Marshal(deliveryRealTime)
	if err != nil {
		return fmt.Errorf("GasStats假数据错误(序列化失败):%s", err)
	}

	//if db := d.cxt.Cxt.GetDB().Table("t_gas_stats").Where("id = ?", 1).Updates(map[string]interface{}{"delivery_real_time": string(deliveryRealTimeByte)}); db.Error != nil {
	//	return fmt.Errorf("GasStats假数据存储失败:%s", db.Error)
	//}
	if db := d.cxt.Cxt.GetDB().Table("t_gas_stats").Updates(map[string]interface{}{"delivery_real_time": string(deliveryRealTimeByte), "delivery_today_zy": zyGasToday, "delivery_today_zm": zmGasToday}); db.Error != nil {
		return fmt.Errorf("GasStats假数据存储失败:%s", db.Error)
	}

	err = d.cxt.GetEmqtt().Send("GasStats")
	if err != nil {
		logrus.Errorf("天然气统计数据发送失败.%+v", err)
	}

	deliveryRealTime = make([][]influxdb.BaseData, 2)

	for index, _ := range deliveryRealTime {
		switch index {
		case 0:
			for i := 0; i < nowHourPlus; i++ {
				deliveryRealTime[index] = append(deliveryRealTime[index],
					influxdb.BaseData{
						Name:  utils.Get24HoursFromZero(i).Format("15:04"),
						Value: utils.FormatFloatDigit(0.1512 * (float64(i) + rand.Float64())),
					})
			}
		case 1:
			for i := 0; i < nowHourPlus; i++ {
				deliveryRealTime[index] = append(deliveryRealTime[index],
					influxdb.BaseData{
						Name:  utils.Get24HoursFromZero(i).Format("15:04"),
						Value: utils.FormatFloatDigit(0.1171 * (float64(i) + rand.Float64())),
					})
			}
		}
		// for i := 0; i < nowHourPlus; i++ {
		// 	deliveryRealTime[index] = append(deliveryRealTime[index],
		// 		influxdb.BaseData{
		// 			Name:  utils.Get24HoursFromZero(i).Format("15:04"),
		// 			Value: utils.FormatFloatDigit((2+rand.Float64()/12)*float64(i) + rand.Float64()),
		// 		})
		// }
	}
	zyOilToday := utils.FormatFloatDigit(0.1512 * (float64(nowHourPlus) + rand.Float64()))
	zmOilToday := utils.FormatFloatDigit(0.1171 * (float64(nowHourPlus) + rand.Float64()))
	deliveryRealTimeByte, err = json.Marshal(deliveryRealTime)
	if err != nil {
		return fmt.Errorf("OilStats假数据错误(序列化失败):%s", err)
	}

	if db := d.cxt.Cxt.GetDB().Table("t_oil_stats").Updates(map[string]interface{}{"delivery_real_time": string(deliveryRealTimeByte), "delivery_today_zy": zyOilToday, "delivery_today_zm": zmOilToday}); db.Error != nil {
		return fmt.Errorf("OilStats假数据存储失败:%s", db.Error)
	}
	// if db := d.cxt.Cxt.GetDB().Table("t_oil_stats").Updates(map[string]interface{}{"delivery_today_zy": zyOilToday}); db.Error != nil {
	// 	return fmt.Errorf("OilStats假数据存储失败:%s", db.Error)
	// }
	// if db := d.cxt.Cxt.GetDB().Table("t_oil_stats").Updates(map[string]interface{}{"delivery_today_zm": zmOilToday}); db.Error != nil {
	// 	return fmt.Errorf("OilStats假数据存储失败:%s", db.Error)
	// }
	err = d.cxt.GetEmqtt().Send("OilStats")
	if err != nil {
		logrus.Errorf("原油统计数据发送失败.%+v", err)
	}

	return nil
}



// 天然气数据
type GasStats struct {
	Id                 int     `json:"id" gorm:"primary_key"`
	DeliveryTodayZY    float64 `json:"deliveryTodayZY"`
	DeliveryTodayZM    float64 `json:"deliveryTodayZM"`
	DeliveryPlanAnnual string  `json:"deliveryPlanAnnual" gorm:"type:mediumtext"`
	DeliveryRealTime   string  `json:"deliveryRealTime" gorm:"type:mediumtext"`
}

func (GasStats) TableName() string {
	return "t_gas_stats"
}

// 原油数据
type OilStats struct {
	Id                 int     `json:"id" gorm:"primary_key"`
	DeliveryTodayZY    float64 `json:"delivery_today_zy"`
	DeliveryTodayZM    float64 `json:"delivery_today_zm"`
	DeliveryPlanAnnual string  `json:"delivery_plan_annual" gorm:"type:mediumtext"`
	DeliveryRealTime   string  `json:"delivery_real_time" gorm:"type:mediumtext"`
	Zy   string  `json:"zy" gorm:"type:mediumtext"`
	Zm   string  `json:"zm" gorm:"type:mediumtext"`
}

func (OilStats) TableName() string {
	return "t_oil_stats"
}