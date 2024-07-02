# ParamsEsQueryRuleOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggField** | Pointer to **string** | The name of the numeric field that is used in the aggregation. This property is required when &#x60;aggType&#x60; is &#x60;avg&#x60;, &#x60;max&#x60;, &#x60;min&#x60; or &#x60;sum&#x60;.  | [optional] 
**AggType** | Pointer to [**Aggtype**](Aggtype.md) |  | [optional] [default to COUNT]
**EsQuery** | **string** | The query definition, which uses Elasticsearch Query DSL. | 
**ExcludeHitsFromPreviousRun** | Pointer to **bool** | Indicates whether to exclude matches from previous runs. If &#x60;true&#x60;, you can avoid alert duplication by excluding documents that have already been detected by the previous rule run. This option is not available when a grouping field is specified.  | [optional] 
**GroupBy** | Pointer to [**Groupby**](Groupby.md) |  | [optional] [default to ALL]
**Index** | [**ParamsEsQueryRuleOneOf2Index**](ParamsEsQueryRuleOneOf2Index.md) |  | 
**SearchType** | Pointer to **string** | The type of query, in this case a query that uses Elasticsearch Query DSL. | [optional] [default to "esQuery"]
**Size** | Pointer to **int32** | The number of documents to pass to the configured actions when the threshold condition is met.  | [optional] 
**TermField** | Pointer to [**Termfield**](Termfield.md) |  | [optional] 
**TermSize** | Pointer to **int32** | This property is required when &#x60;groupBy&#x60; is &#x60;top&#x60;. It specifies the number of groups to check against the threshold and therefore limits the number of alerts on high cardinality fields.  | [optional] 
**Threshold** | **[]int32** | The threshold value that is used with the &#x60;thresholdComparator&#x60;. If the &#x60;thresholdComparator&#x60; is &#x60;between&#x60; or &#x60;notBetween&#x60;, you must specify the boundary values.  | 
**ThresholdComparator** | [**Thresholdcomparator**](Thresholdcomparator.md) |  | 
**TimeField** | **string** | The field that is used to calculate the time window. | 
**TimeWindowSize** | **int32** | The size of the time window (in &#x60;timeWindowUnit&#x60; units), which determines how far back to search for documents. Generally it should be a value higher than the rule check interval to avoid gaps in detection.  | 
**TimeWindowUnit** | [**Timewindowunit**](Timewindowunit.md) |  | 

## Methods

### NewParamsEsQueryRuleOneOf2

`func NewParamsEsQueryRuleOneOf2(esQuery string, index ParamsEsQueryRuleOneOf2Index, threshold []int32, thresholdComparator Thresholdcomparator, timeField string, timeWindowSize int32, timeWindowUnit Timewindowunit, ) *ParamsEsQueryRuleOneOf2`

NewParamsEsQueryRuleOneOf2 instantiates a new ParamsEsQueryRuleOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsEsQueryRuleOneOf2WithDefaults

`func NewParamsEsQueryRuleOneOf2WithDefaults() *ParamsEsQueryRuleOneOf2`

NewParamsEsQueryRuleOneOf2WithDefaults instantiates a new ParamsEsQueryRuleOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggField

`func (o *ParamsEsQueryRuleOneOf2) GetAggField() string`

GetAggField returns the AggField field if non-nil, zero value otherwise.

### GetAggFieldOk

`func (o *ParamsEsQueryRuleOneOf2) GetAggFieldOk() (*string, bool)`

GetAggFieldOk returns a tuple with the AggField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggField

`func (o *ParamsEsQueryRuleOneOf2) SetAggField(v string)`

SetAggField sets AggField field to given value.

### HasAggField

`func (o *ParamsEsQueryRuleOneOf2) HasAggField() bool`

HasAggField returns a boolean if a field has been set.

### GetAggType

`func (o *ParamsEsQueryRuleOneOf2) GetAggType() Aggtype`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *ParamsEsQueryRuleOneOf2) GetAggTypeOk() (*Aggtype, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *ParamsEsQueryRuleOneOf2) SetAggType(v Aggtype)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *ParamsEsQueryRuleOneOf2) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetEsQuery

`func (o *ParamsEsQueryRuleOneOf2) GetEsQuery() string`

GetEsQuery returns the EsQuery field if non-nil, zero value otherwise.

### GetEsQueryOk

`func (o *ParamsEsQueryRuleOneOf2) GetEsQueryOk() (*string, bool)`

GetEsQueryOk returns a tuple with the EsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQuery

`func (o *ParamsEsQueryRuleOneOf2) SetEsQuery(v string)`

SetEsQuery sets EsQuery field to given value.


### GetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf2) GetExcludeHitsFromPreviousRun() bool`

GetExcludeHitsFromPreviousRun returns the ExcludeHitsFromPreviousRun field if non-nil, zero value otherwise.

### GetExcludeHitsFromPreviousRunOk

`func (o *ParamsEsQueryRuleOneOf2) GetExcludeHitsFromPreviousRunOk() (*bool, bool)`

GetExcludeHitsFromPreviousRunOk returns a tuple with the ExcludeHitsFromPreviousRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf2) SetExcludeHitsFromPreviousRun(v bool)`

SetExcludeHitsFromPreviousRun sets ExcludeHitsFromPreviousRun field to given value.

### HasExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf2) HasExcludeHitsFromPreviousRun() bool`

HasExcludeHitsFromPreviousRun returns a boolean if a field has been set.

### GetGroupBy

`func (o *ParamsEsQueryRuleOneOf2) GetGroupBy() Groupby`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ParamsEsQueryRuleOneOf2) GetGroupByOk() (*Groupby, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ParamsEsQueryRuleOneOf2) SetGroupBy(v Groupby)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ParamsEsQueryRuleOneOf2) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetIndex

`func (o *ParamsEsQueryRuleOneOf2) GetIndex() ParamsEsQueryRuleOneOf2Index`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ParamsEsQueryRuleOneOf2) GetIndexOk() (*ParamsEsQueryRuleOneOf2Index, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ParamsEsQueryRuleOneOf2) SetIndex(v ParamsEsQueryRuleOneOf2Index)`

SetIndex sets Index field to given value.


### GetSearchType

`func (o *ParamsEsQueryRuleOneOf2) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *ParamsEsQueryRuleOneOf2) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *ParamsEsQueryRuleOneOf2) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *ParamsEsQueryRuleOneOf2) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetSize

`func (o *ParamsEsQueryRuleOneOf2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ParamsEsQueryRuleOneOf2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ParamsEsQueryRuleOneOf2) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ParamsEsQueryRuleOneOf2) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTermField

`func (o *ParamsEsQueryRuleOneOf2) GetTermField() Termfield`

GetTermField returns the TermField field if non-nil, zero value otherwise.

### GetTermFieldOk

`func (o *ParamsEsQueryRuleOneOf2) GetTermFieldOk() (*Termfield, bool)`

GetTermFieldOk returns a tuple with the TermField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermField

`func (o *ParamsEsQueryRuleOneOf2) SetTermField(v Termfield)`

SetTermField sets TermField field to given value.

### HasTermField

`func (o *ParamsEsQueryRuleOneOf2) HasTermField() bool`

HasTermField returns a boolean if a field has been set.

### GetTermSize

`func (o *ParamsEsQueryRuleOneOf2) GetTermSize() int32`

GetTermSize returns the TermSize field if non-nil, zero value otherwise.

### GetTermSizeOk

`func (o *ParamsEsQueryRuleOneOf2) GetTermSizeOk() (*int32, bool)`

GetTermSizeOk returns a tuple with the TermSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSize

`func (o *ParamsEsQueryRuleOneOf2) SetTermSize(v int32)`

SetTermSize sets TermSize field to given value.

### HasTermSize

`func (o *ParamsEsQueryRuleOneOf2) HasTermSize() bool`

HasTermSize returns a boolean if a field has been set.

### GetThreshold

`func (o *ParamsEsQueryRuleOneOf2) GetThreshold() []int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ParamsEsQueryRuleOneOf2) GetThresholdOk() (*[]int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ParamsEsQueryRuleOneOf2) SetThreshold(v []int32)`

SetThreshold sets Threshold field to given value.


### GetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf2) GetThresholdComparator() Thresholdcomparator`

GetThresholdComparator returns the ThresholdComparator field if non-nil, zero value otherwise.

### GetThresholdComparatorOk

`func (o *ParamsEsQueryRuleOneOf2) GetThresholdComparatorOk() (*Thresholdcomparator, bool)`

GetThresholdComparatorOk returns a tuple with the ThresholdComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf2) SetThresholdComparator(v Thresholdcomparator)`

SetThresholdComparator sets ThresholdComparator field to given value.


### GetTimeField

`func (o *ParamsEsQueryRuleOneOf2) GetTimeField() string`

GetTimeField returns the TimeField field if non-nil, zero value otherwise.

### GetTimeFieldOk

`func (o *ParamsEsQueryRuleOneOf2) GetTimeFieldOk() (*string, bool)`

GetTimeFieldOk returns a tuple with the TimeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeField

`func (o *ParamsEsQueryRuleOneOf2) SetTimeField(v string)`

SetTimeField sets TimeField field to given value.


### GetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf2) GetTimeWindowSize() int32`

GetTimeWindowSize returns the TimeWindowSize field if non-nil, zero value otherwise.

### GetTimeWindowSizeOk

`func (o *ParamsEsQueryRuleOneOf2) GetTimeWindowSizeOk() (*int32, bool)`

GetTimeWindowSizeOk returns a tuple with the TimeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf2) SetTimeWindowSize(v int32)`

SetTimeWindowSize sets TimeWindowSize field to given value.


### GetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf2) GetTimeWindowUnit() Timewindowunit`

GetTimeWindowUnit returns the TimeWindowUnit field if non-nil, zero value otherwise.

### GetTimeWindowUnitOk

`func (o *ParamsEsQueryRuleOneOf2) GetTimeWindowUnitOk() (*Timewindowunit, bool)`

GetTimeWindowUnitOk returns a tuple with the TimeWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf2) SetTimeWindowUnit(v Timewindowunit)`

SetTimeWindowUnit sets TimeWindowUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


