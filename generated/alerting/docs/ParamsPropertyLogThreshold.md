# ParamsPropertyLogThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**[][]CountCriteriaInner**]([]CountCriteriaInner.md) |  | [optional] 
**Count** | [**CountCount**](CountCount.md) |  | 
**TimeSize** | **float32** |  | 
**TimeUnit** | **string** |  | 
**LogView** | [**CountLogView**](CountLogView.md) |  | 
**GroupBy** | Pointer to **[]string** |  | [optional] 

## Methods

### NewParamsPropertyLogThreshold

`func NewParamsPropertyLogThreshold(count CountCount, timeSize float32, timeUnit string, logView CountLogView, ) *ParamsPropertyLogThreshold`

NewParamsPropertyLogThreshold instantiates a new ParamsPropertyLogThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsPropertyLogThresholdWithDefaults

`func NewParamsPropertyLogThresholdWithDefaults() *ParamsPropertyLogThreshold`

NewParamsPropertyLogThresholdWithDefaults instantiates a new ParamsPropertyLogThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *ParamsPropertyLogThreshold) GetCriteria() [][]CountCriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ParamsPropertyLogThreshold) GetCriteriaOk() (*[][]CountCriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ParamsPropertyLogThreshold) SetCriteria(v [][]CountCriteriaInner)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ParamsPropertyLogThreshold) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetCount

`func (o *ParamsPropertyLogThreshold) GetCount() CountCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ParamsPropertyLogThreshold) GetCountOk() (*CountCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ParamsPropertyLogThreshold) SetCount(v CountCount)`

SetCount sets Count field to given value.


### GetTimeSize

`func (o *ParamsPropertyLogThreshold) GetTimeSize() float32`

GetTimeSize returns the TimeSize field if non-nil, zero value otherwise.

### GetTimeSizeOk

`func (o *ParamsPropertyLogThreshold) GetTimeSizeOk() (*float32, bool)`

GetTimeSizeOk returns a tuple with the TimeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSize

`func (o *ParamsPropertyLogThreshold) SetTimeSize(v float32)`

SetTimeSize sets TimeSize field to given value.


### GetTimeUnit

`func (o *ParamsPropertyLogThreshold) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ParamsPropertyLogThreshold) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ParamsPropertyLogThreshold) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetLogView

`func (o *ParamsPropertyLogThreshold) GetLogView() CountLogView`

GetLogView returns the LogView field if non-nil, zero value otherwise.

### GetLogViewOk

`func (o *ParamsPropertyLogThreshold) GetLogViewOk() (*CountLogView, bool)`

GetLogViewOk returns a tuple with the LogView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogView

`func (o *ParamsPropertyLogThreshold) SetLogView(v CountLogView)`

SetLogView sets LogView field to given value.


### GetGroupBy

`func (o *ParamsPropertyLogThreshold) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ParamsPropertyLogThreshold) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ParamsPropertyLogThreshold) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ParamsPropertyLogThreshold) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


