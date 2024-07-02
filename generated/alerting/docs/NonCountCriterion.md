# NonCountCriterion

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

## Methods

### NewNonCountCriterion

`func NewNonCountCriterion() *NonCountCriterion`

NewNonCountCriterion instantiates a new NonCountCriterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonCountCriterionWithDefaults

`func NewNonCountCriterionWithDefaults() *NonCountCriterion`

NewNonCountCriterionWithDefaults instantiates a new NonCountCriterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *NonCountCriterion) GetThreshold() []float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *NonCountCriterion) GetThresholdOk() (*[]float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *NonCountCriterion) SetThreshold(v []float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *NonCountCriterion) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetComparator

`func (o *NonCountCriterion) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *NonCountCriterion) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *NonCountCriterion) SetComparator(v string)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *NonCountCriterion) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetTimeUnit

`func (o *NonCountCriterion) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *NonCountCriterion) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *NonCountCriterion) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *NonCountCriterion) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetTimeSize

`func (o *NonCountCriterion) GetTimeSize() float32`

GetTimeSize returns the TimeSize field if non-nil, zero value otherwise.

### GetTimeSizeOk

`func (o *NonCountCriterion) GetTimeSizeOk() (*float32, bool)`

GetTimeSizeOk returns a tuple with the TimeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSize

`func (o *NonCountCriterion) SetTimeSize(v float32)`

SetTimeSize sets TimeSize field to given value.

### HasTimeSize

`func (o *NonCountCriterion) HasTimeSize() bool`

HasTimeSize returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *NonCountCriterion) GetWarningThreshold() []float32`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *NonCountCriterion) GetWarningThresholdOk() (*[]float32, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *NonCountCriterion) SetWarningThreshold(v []float32)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *NonCountCriterion) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetWarningComparator

`func (o *NonCountCriterion) GetWarningComparator() string`

GetWarningComparator returns the WarningComparator field if non-nil, zero value otherwise.

### GetWarningComparatorOk

`func (o *NonCountCriterion) GetWarningComparatorOk() (*string, bool)`

GetWarningComparatorOk returns a tuple with the WarningComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningComparator

`func (o *NonCountCriterion) SetWarningComparator(v string)`

SetWarningComparator sets WarningComparator field to given value.

### HasWarningComparator

`func (o *NonCountCriterion) HasWarningComparator() bool`

HasWarningComparator returns a boolean if a field has been set.

### GetMetric

`func (o *NonCountCriterion) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *NonCountCriterion) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *NonCountCriterion) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *NonCountCriterion) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetAggType

`func (o *NonCountCriterion) GetAggType() string`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *NonCountCriterion) GetAggTypeOk() (*string, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *NonCountCriterion) SetAggType(v string)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *NonCountCriterion) HasAggType() bool`

HasAggType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


