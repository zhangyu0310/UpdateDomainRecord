package sdk

import (
	"UpdateDomainRecord/config"
	dns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zhangyu0310/zlogger"
)

// CreateClient 使用AK&SK初始化账号Client
//
//	@param accessKeyId
//	@param accessKeySecret
//	@return Client
//	@throws Exception
func CreateClient(accessKeyId *string, accessKeySecret *string) (client *dns.Client, err error) {
	cfg := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	cfg.Endpoint = tea.String(config.GetGlobalConfig().EndPoint)
	client = &dns.Client{}
	client, err = dns.NewClient(cfg)
	return client, err
}

func handleSdkErr(err error) error {
	var sdkErr = &tea.SDKError{}
	if err != nil {
		if tmp, ok := err.(*tea.SDKError); ok {
			sdkErr = tmp
		} else {
			sdkErr.Message = tea.String(err.Error())
		}
		return sdkErr
	}
	return nil
}

func RunOnce(ip string) (err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。
	cfg := config.GetGlobalConfig()
	client, err := CreateClient(
		tea.String(cfg.AccessKeyID),
		tea.String(cfg.AccessKeySecret))
	if err != nil {
		return err
	}
	zlogger.Info("Create API client success!")

	request := &dns.DescribeDomainRecordsRequest{
		DomainName: tea.String(cfg.DomainName),
		Type:       tea.String("A"),
	}
	response, err := client.DescribeDomainRecords(request)
	if err = handleSdkErr(err); err != nil {
		return err
	}
	zlogger.Info("Describe domain records success!")

	for _, record := range response.Body.DomainRecords.Record {
		updateRequest := &dns.UpdateDomainRecordRequest{
			Lang:     tea.String("en"),
			RR:       record.RR,
			RecordId: record.RecordId,
			TTL:      tea.Int64(600),
			Type:     record.Type,
			Value:    tea.String(ip),
		}
		_, err := client.UpdateDomainRecord(updateRequest)
		if err = handleSdkErr(err); err != nil {
			return err
		}
		zlogger.InfoF("Update %s domain record success!", *record.RR)
	}

	return err
}
