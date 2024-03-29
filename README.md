# convert-to-dpfm-product-master-from-sap-product-master

convert-to-dpfm-product-master-from-sap-product-master は、周辺業務システム　を データ連携基盤 と統合することを目的に、API で品目マスタデータを取得するマイクロサービスです。  
https://xxx.xxx.io/api/API_PRODUCT_MASTER_SRV/creates/

## 動作環境

convert-to-dpfm-product-master-from-sap-product-master の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-dpfm-product-master-from-sap-product-master が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_PRODUCT_MASTER_SRV/creates/

## 本レポジトリ に 含まれる API名
convert-to-dpfm-product-master-from-sap-product-master には、次の API をコールするためのリソースが含まれています。  

* A_General（データ連携基盤 品目マスタ - 基本データ）
* A_ProductDescription（データ連携基盤 品目マスタ - 品目テキストデータ）
* A_GeneralDoc（データ連携基盤 品目マスタ - 基本文書データ）
* A_BusinessPartner（データ連携基盤 品目マスタ - ビジネスパートナデータ）
* A_ProductDescriptionByBusinessPartner（データ連携基盤 品目マスタ - ビジネスパートナ品目テキストデータ）
* A_BPPlant（データ連携基盤 品目マスタ - ビジネスパートナプラントデータ）
* A_BPPlantDoc（データ連携基盤 品目マスタ - ビジネスパートナプラント文書データ）
* A_StorageLocation （データ連携基盤 品目マスタ - 保管場所データ）
* A_MRPArea（データ連携基盤 品目マスタ - MRPエリアデータ）
* A_WorkScheduling（データ連携基盤 品目マスタ - 作業計画データ）
* A_Tax（データ連携基盤 品目マスタ - 税データ）
* A_Accounting （データ連携基盤 品目マスタ - 会計データ）
* A_Allergen（データ連携基盤 品目マスタ - アレルゲンデータ）
* A_NutritionalInfo（データ連携基盤 品目マスタ - 栄養成分データ）
* A_Calories（データ連携基盤 品目マスタ - 熱量データ）


## API への 値入力条件 の 初期値
convert-to-dpfm-product-master-from-sap-product-master において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "DPFMProductMasterCreates",
	"accepter": ["General"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMProductMasterCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	subFuncFin := make(chan error)
	exconfFin := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exconfAllExist, e = c.confirmor.Conf(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.Errorf("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, errs, input)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 品目マスタ の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "PlusMinusFlag" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
