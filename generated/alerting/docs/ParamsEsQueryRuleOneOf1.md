# ParamsEsQueryRuleOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggField** | Pointer to **string** | The name of the numeric field that is used in the aggregation. This property is required when &#x60;aggType&#x60; is &#x60;avg&#x60;, &#x60;max&#x60;, &#x60;min&#x60; or &#x60;sum&#x60;.  | [optional] 
**AggType** | Pointer to [**Aggtype**](Aggtype.md) |  | [optional] [default to COUNT]
**ExcludeHitsFromPreviousRun** | Pointer to **bool** | Indicates whether to exclude matches from previous runs. If &#x60;true&#x60;, you can avoid alert duplication by excluding documents that have already been detected by the previous rule run. This option is not available when a grouping field is specified.  | [optional] 
**GroupBy** | Pointer to [**Groupby**](Groupby.md) |  | [optional] [default to ALL]
**SearchConfiguration** | Pointer to [**ParamsEsQueryRuleOneOf1SearchConfiguration**](ParamsEsQueryRuleOneOf1SearchConfiguration.md) |  | [optional] 
**SearchType** | **string** | The type of query, in this case a text-based query that uses KQL or Lucene. | 
**Size** | **int32** | The number of documents to pass to the configured actions when the threshold condition is met.  | 
**TermField** | Pointer to [**Termfield**](Termfield.md) |  | [optional] 
**TermSize** | Pointer to **int32** | This property is required when &#x60;groupBy&#x60; is &#x60;top&#x60;. It specifies the number of groups to check against the threshold and therefore limits the number of alerts on high cardinality fields.  | [optional] 
**Threshold** | **[]int32** | The threshold value that is used with the &#x60;thresholdComparator&#x60;. If the &#x60;thresholdComparator&#x60; is &#x60;between&#x60; or &#x60;notBetween&#x60;, you must specify the boundary values.  | 
**ThresholdComparator** | [**Thresholdcomparator**](Thresholdcomparator.md) |  | 
**TimeField** | Pointer to **string** | The field that is used to calculate the time window. | [optional] 
**TimeWindowSize** | **int32** | The size of the time window (in &#x60;timeWindowUnit&#x60; units), which determines how far back to search for documents. Generally it should be a value higher than the rule check interval to avoid gaps in detection.  | 
**TimeWindowUnit** | [**Timewindowunit**](Timewindowunit.md) |  | 

## Methods

### NewParamsEsQueryRuleOneOf1

`func NewParamsEsQueryRuleOneOf1(searchType string, size int32, threshold []int32, thresholdComparator Thresholdcomparator, timeWindowSize int32, timeWindowUnit Timewindowunit, ) *ParamsEsQueryRuleOneOf1`

NewParamsEsQueryRuleOneOf1 instantiates a new ParamsEsQueryRuleOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsEsQueryRuleOneOf1WithDefaults

`func NewParamsEsQueryRuleOneOf1WithDefaults() *ParamsEsQueryRuleOneOf1`

NewParamsEsQueryRuleOneOf1WithDefaults instantiates a new ParamsEsQueryRuleOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggField

`func (o *ParamsEsQueryRuleOneOf1) GetAggField() string`

GetAggField returns the AggField field if non-nil, zero value otherwise.

### GetAggFieldOk

`func (o *ParamsEsQueryRuleOneOf1) GetAggFieldOk() (*string, bool)`

GetAggFieldOk returns a tuple with the AggField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggField

`func (o *ParamsEsQueryRuleOneOf1) SetAggField(v string)`

SetAggField sets AggField field to given value.

### HasAggField

`func (o *ParamsEsQueryRuleOneOf1) HasAggField() bool`

HasAggField returns a boolean if a field has been set.

### GetAggType

`func (o *ParamsEsQueryRuleOneOf1) GetAggType() Aggtype`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *ParamsEsQueryRuleOneOf1) GetAggTypeOk() (*Aggtype, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *ParamsEsQueryRuleOneOf1) SetAggType(v Aggtype)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *ParamsEsQueryRuleOneOf1) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf1) GetExcludeHitsFromPreviousRun() bool`

GetExcludeHitsFromPreviousRun returns the ExcludeHitsFromPreviousRun field if non-nil, zero value otherwise.

### GetExcludeHitsFromPreviousRunOk

`func (o *ParamsEsQueryRuleOneOf1) GetExcludeHitsFromPreviousRunOk() (*bool, bool)`

GetExcludeHitsFromPreviousRunOk returns a tuple with the ExcludeHitsFromPreviousRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf1) SetExcludeHitsFromPreviousRun(v bool)`

SetExcludeHitsFromPreviousRun sets ExcludeHitsFromPreviousRun field to given value.

### HasExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf1) HasExcludeHitsFromPreviousRun() bool`

HasExcludeHitsFromPreviousRun returns a boolean if a field has been set.

### GetGroupBy

`func (o *ParamsEsQueryRuleOneOf1) GetGroupBy() Groupby`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ParamsEsQueryRuleOneOf1) GetGroupByOk() (*Groupby, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ParamsEsQueryRuleOneOf1) SetGroupBy(v Groupby)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ParamsEsQueryRuleOneOf1) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetSearchConfiguration

`func (o *ParamsEsQueryRuleOneOf1) GetSearchConfiguration() ParamsEsQueryRuleOneOf1SearchConfiguration`

GetSearchConfiguration returns the SearchConfiguration field if non-nil, zero value otherwise.

### GetSearchConfigurationOk

`func (o *ParamsEsQueryRuleOneOf1) GetSearchConfigurationOk() (*ParamsEsQueryRuleOneOf1SearchConfiguration, bool)`

GetSearchConfigurationOk returns a tuple with the SearchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchConfiguration

`func (o *ParamsEsQueryRuleOneOf1) SetSearchConfiguration(v ParamsEsQueryRuleOneOf1SearchConfiguration)`

SetSearchConfiguration sets SearchConfiguration field to given value.

### HasSearchConfiguration

`func (o *ParamsEsQueryRuleOneOf1) HasSearchConfiguration() bool`

HasSearchConfiguration returns a boolean if a field has been set.

### GetSearchType

`func (o *ParamsEsQueryRuleOneOf1) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *ParamsEsQueryRuleOneOf1) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *ParamsEsQueryRuleOneOf1) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.


### GetSize

`func (o *ParamsEsQueryRuleOneOf1) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ParamsEsQueryRuleOneOf1) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ParamsEsQueryRuleOneOf1) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTermField

`func (o *ParamsEsQueryRuleOneOf1) GetTermField() Termfield`

GetTermField returns the TermField field if non-nil, zero value otherwise.

### GetTermFieldOk

`func (o *ParamsEsQueryRuleOneOf1) GetTermFieldOk() (*Termfield, bool)`

GetTermFieldOk returns a tuple with the TermField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermField

`func (o *ParamsEsQueryRuleOneOf1) SetTermField(v Termfield)`

SetTermField sets TermField field to given value.

### HasTermField

`func (o *ParamsEsQueryRuleOneOf1) HasTermField() bool`

HasTermField returns a boolean if a field has been set.

### GetTermSize

`func (o *ParamsEsQueryRuleOneOf1) GetTermSize() int32`

GetTermSize returns the TermSize field if non-nil, zero value otherwise.

### GetTermSizeOk

`func (o *ParamsEsQueryRuleOneOf1) GetTermSizeOk() (*int32, bool)`

GetTermSizeOk returns a tuple with the TermSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSize

`func (o *ParamsEsQueryRuleOneOf1) SetTermSize(v int32)`

SetTermSize sets TermSize field to given value.

### HasTermSize

`func (o *ParamsEsQueryRuleOneOf1) HasTermSize() bool`

HasTermSize returns a boolean if a field has been set.

### GetThreshold

`func (o *ParamsEsQueryRuleOneOf1) GetThreshold() []int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ParamsEsQueryRuleOneOf1) GetThresholdOk() (*[]int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ParamsEsQueryRuleOneOf1) SetThreshold(v []int32)`

SetThreshold sets Threshold field to given value.


### GetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf1) GetThresholdComparator() Thresholdcomparator`

GetThresholdComparator returns the ThresholdComparator field if non-nil, zero value otherwise.

### GetThresholdComparatorOk

`func (o *ParamsEsQueryRuleOneOf1) GetThresholdComparatorOk() (*Thresholdcomparator, bool)`

GetThresholdComparatorOk returns a tuple with the ThresholdComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf1) SetThresholdComparator(v Thresholdcomparator)`

SetThresholdComparator sets ThresholdComparator field to given value.


### GetTimeField

`func (o *ParamsEsQueryRuleOneOf1) GetTimeField() string`

GetTimeField returns the TimeField field if non-nil, zero value otherwise.

### GetTimeFieldOk

`func (o *ParamsEsQueryRuleOneOf1) GetTimeFieldOk() (*string, bool)`

GetTimeFieldOk returns a tuple with the TimeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeField

`func (o *ParamsEsQueryRuleOneOf1) SetTimeField(v string)`

SetTimeField sets TimeField field to given value.

### HasTimeField

`func (o *ParamsEsQueryRuleOneOf1) HasTimeField() bool`

HasTimeField returns a boolean if a field has been set.

### GetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf1) GetTimeWindowSize() int32`

GetTimeWindowSize returns the TimeWindowSize field if non-nil, zero value otherwise.

### GetTimeWindowSizeOk

`func (o *ParamsEsQueryRuleOneOf1) GetTimeWindowSizeOk() (*int32, bool)`

GetTimeWindowSizeOk returns a tuple with the TimeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf1) SetTimeWindowSize(v int32)`

SetTimeWindowSize sets TimeWindowSize field to given value.


### GetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf1) GetTimeWindowUnit() Timewindowunit`

GetTimeWindowUnit returns the TimeWindowUnit field if non-nil, zero value otherwise.

### GetTimeWindowUnitOk

`func (o *ParamsEsQueryRuleOneOf1) GetTimeWindowUnitOk() (*Timewindowunit, bool)`

GetTimeWindowUnitOk returns a tuple with the TimeWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf1) SetTimeWindowUnit(v Timewindowunit)`

SetTimeWindowUnit sets TimeWindowUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


