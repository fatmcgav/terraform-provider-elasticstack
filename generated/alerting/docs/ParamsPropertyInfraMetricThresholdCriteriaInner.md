# ParamsPropertyInfraMetricThresholdCriteriaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **[]float32** |  | [optional] 
**Comparator** | Pointer to **string** |  | [optional] 
**TimeUnit** | Pointer to **string** |  | [optional] 
**TimeSize** | Pointer to **float32** |  | [optional] 
**WarningThreshold** | Pointer to **[]float32** |  | [optional] 
**WarningComparator** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**AggType** | Pointer to **string** |  | [optional] 
**CustomMetric** | Pointer to [**[]CustomCriterionCustomMetricInner**](CustomCriterionCustomMetricInner.md) |  | [optional] 
**Equation** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewParamsPropertyInfraMetricThresholdCriteriaInner

`func NewParamsPropertyInfraMetricThresholdCriteriaInner() *ParamsPropertyInfraMetricThresholdCriteriaInner`

NewParamsPropertyInfraMetricThresholdCriteriaInner instantiates a new ParamsPropertyInfraMetricThresholdCriteriaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsPropertyInfraMetricThresholdCriteriaInnerWithDefaults

`func NewParamsPropertyInfraMetricThresholdCriteriaInnerWithDefaults() *ParamsPropertyInfraMetricThresholdCriteriaInner`

NewParamsPropertyInfraMetricThresholdCriteriaInnerWithDefaults instantiates a new ParamsPropertyInfraMetricThresholdCriteriaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetThreshold() []float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetThresholdOk() (*[]float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetThreshold(v []float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetComparator(v string)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetTimeUnit

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetTimeSize

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetTimeSize() float32`

GetTimeSize returns the TimeSize field if non-nil, zero value otherwise.

### GetTimeSizeOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetTimeSizeOk() (*float32, bool)`

GetTimeSizeOk returns a tuple with the TimeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSize

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetTimeSize(v float32)`

SetTimeSize sets TimeSize field to given value.

### HasTimeSize

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasTimeSize() bool`

HasTimeSize returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetWarningThreshold() []float32`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetWarningThresholdOk() (*[]float32, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetWarningThreshold(v []float32)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetWarningComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetWarningComparator() string`

GetWarningComparator returns the WarningComparator field if non-nil, zero value otherwise.

### GetWarningComparatorOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetWarningComparatorOk() (*string, bool)`

GetWarningComparatorOk returns a tuple with the WarningComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetWarningComparator(v string)`

SetWarningComparator sets WarningComparator field to given value.

### HasWarningComparator

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasWarningComparator() bool`

HasWarningComparator returns a boolean if a field has been set.

### GetMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetAggType

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetAggType() string`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetAggTypeOk() (*string, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetAggType(v string)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetCustomMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetCustomMetric() []CustomCriterionCustomMetricInner`

GetCustomMetric returns the CustomMetric field if non-nil, zero value otherwise.

### GetCustomMetricOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetCustomMetricOk() (*[]CustomCriterionCustomMetricInner, bool)`

GetCustomMetricOk returns a tuple with the CustomMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetCustomMetric(v []CustomCriterionCustomMetricInner)`

SetCustomMetric sets CustomMetric field to given value.

### HasCustomMetric

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasCustomMetric() bool`

HasCustomMetric returns a boolean if a field has been set.

### GetEquation

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetEquation() string`

GetEquation returns the Equation field if non-nil, zero value otherwise.

### GetEquationOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetEquationOk() (*string, bool)`

GetEquationOk returns a tuple with the Equation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquation

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetEquation(v string)`

SetEquation sets Equation field to given value.

### HasEquation

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasEquation() bool`

HasEquation returns a boolean if a field has been set.

### GetLabel

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ParamsPropertyInfraMetricThresholdCriteriaInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


