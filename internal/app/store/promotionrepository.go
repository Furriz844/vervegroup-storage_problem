package store

import (
	"storage-api/internal/app/model"
)

var promotions = []*model.Promotion{
	{ID: "172FFC14-D229-4C93-B06B-F48B8C095512", Price: 9.68, ExpirationDate: "2022-06-04 06:01:20"},
	{ID: "2aaf8f7b-872e-4a14-afbe-1a9b4799dee3", Price: 51.787896, ExpirationDate: "2018-08-07 17:41:27 +0200 CEST"},
	{ID: "769b000e-1d48-4716-92a3-4285dd6cc1e8", Price: 21.858542, ExpirationDate: "2018-08-03 15:40:36 +0200 CEST"},
	{ID: "d9433531-5b0a-431d-82d4-b413dc34253f", Price: 32.180885, ExpirationDate: "2018-08-10 12:47:53 +0200 CEST"},
	{ID: "e0cb39dc-fad6-42c8-b05a-8d792fe7ec2c", Price: 28.697746, ExpirationDate: "2018-10-19 06:01:31 +0200 CEST"},
	{ID: "14603ea9-6d69-47cb-9a72-5d827cd5bdfc", Price: 68.084971, ExpirationDate: "2018-06-08 10:48:51 +0200 CEST"},
	{ID: "1a176256-a40a-419e-b0f4-772bcfb51d2e", Price: 20.756935, ExpirationDate: "2018-06-16 18:45:56 +0200 CEST"},
	{ID: "6cff915c-9e8e-4481-a470-18a3dbfb8c04", Price: 57.303457, ExpirationDate: "2018-08-25 07:00:37 +0200 CEST"},
	{ID: "10121f45-45bf-4350-9547-d2961eda4329", Price: 29.700212, ExpirationDate: "2018-08-04 08:25:09 +0200 CEST"},
	{ID: "8aaa4862-c976-42a0-839e-43533837eb2c", Price: 75.393388, ExpirationDate: "2018-09-17 16:34:10 +0200 CEST"},
}

func GetPromotions() []*model.Promotion {
	return promotions
}

func GetPromotionById(id string) *model.Promotion {
	for _, a := range promotions {
		if a.ID == id {
			return a
		}
	}
	return nil
}
