# ParamsEsQueryRuleOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggField** | Pointer to **string** | The name of the numeric field that is used in the aggregation. This property is required when &#x60;aggType&#x60; is &#x60;avg&#x60;, &#x60;max&#x60;, &#x60;min&#x60; or &#x60;sum&#x60;.  | [optional] 
**AggType** | Pointer to [**Aggtype**](Aggtype.md) |  | [optional] [default to COUNT]
**EsqlQuery** | [**ParamsEsQueryRuleOneOfEsqlQuery**](ParamsEsQueryRuleOneOfEsqlQuery.md) |  | 
**ExcludeHitsFromPreviousRun** | Pointer to **bool** | Indicates whether to exclude matches from previous runs. If &#x60;true&#x60;, you can avoid alert duplication by excluding documents that have already been detected by the previous rule run. This option is not available when a grouping field is specified.  | [optional] 
**GroupBy** | Pointer to [**Groupby**](Groupby.md) |  | [optional] [default to ALL]
**SearchType** | **string** | The type of query, in this case a query that uses Elasticsearch Query Language (ES|QL). | 
**Size** | **int32** | When &#x60;searchType&#x60; is &#x60;esqlQuery&#x60;, this property is required but it does not affect the rule behavior.  | 
**TermSize** | Pointer to **int32** | This property is required when &#x60;groupBy&#x60; is &#x60;top&#x60;. It specifies the number of groups to check against the threshold and therefore limits the number of alerts on high cardinality fields.  | [optional] 
**Threshold** | **[]int32** | The threshold value that is used with the &#x60;thresholdComparator&#x60;. When &#x60;searchType&#x60; is &#x60;esqlQuery&#x60;, this property is required and must be set to zero.  | 
**ThresholdComparator** | **string** | The comparison function for the threshold. When &#x60;searchType&#x60; is &#x60;esqlQuery&#x60;, this property is required and must be set to \&quot;&gt;\&quot;. Since the &#x60;threshold&#x60; value must be &#x60;0&#x60;, the result is that an alert occurs whenever the query returns results.  | 
**TimeField** | Pointer to **string** | The field that is used to calculate the time window. | [optional] 
**TimeWindowSize** | **int32** | The size of the time window (in &#x60;timeWindowUnit&#x60; units), which determines how far back to search for documents. Generally it should be a value higher than the rule check interval to avoid gaps in detection.  | 
**TimeWindowUnit** | [**Timewindowunit**](Timewindowunit.md) |  | 

## Methods

### NewParamsEsQueryRuleOneOf

`func NewParamsEsQueryRuleOneOf(esqlQuery ParamsEsQueryRuleOneOfEsqlQuery, searchType string, size int32, threshold []int32, thresholdComparator string, timeWindowSize int32, timeWindowUnit Timewindowunit, ) *ParamsEsQueryRuleOneOf`

NewParamsEsQueryRuleOneOf instantiates a new ParamsEsQueryRuleOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsEsQueryRuleOneOfWithDefaults

`func NewParamsEsQueryRuleOneOfWithDefaults() *ParamsEsQueryRuleOneOf`

NewParamsEsQueryRuleOneOfWithDefaults instantiates a new ParamsEsQueryRuleOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggField

`func (o *ParamsEsQueryRuleOneOf) GetAggField() string`

GetAggField returns the AggField field if non-nil, zero value otherwise.

### GetAggFieldOk

`func (o *ParamsEsQueryRuleOneOf) GetAggFieldOk() (*string, bool)`

GetAggFieldOk returns a tuple with the AggField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggField

`func (o *ParamsEsQueryRuleOneOf) SetAggField(v string)`

SetAggField sets AggField field to given value.

### HasAggField

`func (o *ParamsEsQueryRuleOneOf) HasAggField() bool`

HasAggField returns a boolean if a field has been set.

### GetAggType

`func (o *ParamsEsQueryRuleOneOf) GetAggType() Aggtype`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *ParamsEsQueryRuleOneOf) GetAggTypeOk() (*Aggtype, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *ParamsEsQueryRuleOneOf) SetAggType(v Aggtype)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *ParamsEsQueryRuleOneOf) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetEsqlQuery

`func (o *ParamsEsQueryRuleOneOf) GetEsqlQuery() ParamsEsQueryRuleOneOfEsqlQuery`

GetEsqlQuery returns the EsqlQuery field if non-nil, zero value otherwise.

### GetEsqlQueryOk

`func (o *ParamsEsQueryRuleOneOf) GetEsqlQueryOk() (*ParamsEsQueryRuleOneOfEsqlQuery, bool)`

GetEsqlQueryOk returns a tuple with the EsqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsqlQuery

`func (o *ParamsEsQueryRuleOneOf) SetEsqlQuery(v ParamsEsQueryRuleOneOfEsqlQuery)`

SetEsqlQuery sets EsqlQuery field to given value.


### GetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf) GetExcludeHitsFromPreviousRun() bool`

GetExcludeHitsFromPreviousRun returns the ExcludeHitsFromPreviousRun field if non-nil, zero value otherwise.

### GetExcludeHitsFromPreviousRunOk

`func (o *ParamsEsQueryRuleOneOf) GetExcludeHitsFromPreviousRunOk() (*bool, bool)`

GetExcludeHitsFromPreviousRunOk returns a tuple with the ExcludeHitsFromPreviousRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf) SetExcludeHitsFromPreviousRun(v bool)`

SetExcludeHitsFromPreviousRun sets ExcludeHitsFromPreviousRun field to given value.

### HasExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRuleOneOf) HasExcludeHitsFromPreviousRun() bool`

HasExcludeHitsFromPreviousRun returns a boolean if a field has been set.

### GetGroupBy

`func (o *ParamsEsQueryRuleOneOf) GetGroupBy() Groupby`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ParamsEsQueryRuleOneOf) GetGroupByOk() (*Groupby, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ParamsEsQueryRuleOneOf) SetGroupBy(v Groupby)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ParamsEsQueryRuleOneOf) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetSearchType

`func (o *ParamsEsQueryRuleOneOf) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *ParamsEsQueryRuleOneOf) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *ParamsEsQueryRuleOneOf) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.


### GetSize

`func (o *ParamsEsQueryRuleOneOf) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ParamsEsQueryRuleOneOf) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ParamsEsQueryRuleOneOf) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTermSize

`func (o *ParamsEsQueryRuleOneOf) GetTermSize() int32`

GetTermSize returns the TermSize field if non-nil, zero value otherwise.

### GetTermSizeOk

`func (o *ParamsEsQueryRuleOneOf) GetTermSizeOk() (*int32, bool)`

GetTermSizeOk returns a tuple with the TermSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSize

`func (o *ParamsEsQueryRuleOneOf) SetTermSize(v int32)`

SetTermSize sets TermSize field to given value.

### HasTermSize

`func (o *ParamsEsQueryRuleOneOf) HasTermSize() bool`

HasTermSize returns a boolean if a field has been set.

### GetThreshold

`func (o *ParamsEsQueryRuleOneOf) GetThreshold() []int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ParamsEsQueryRuleOneOf) GetThresholdOk() (*[]int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ParamsEsQueryRuleOneOf) SetThreshold(v []int32)`

SetThreshold sets Threshold field to given value.


### GetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf) GetThresholdComparator() string`

GetThresholdComparator returns the ThresholdComparator field if non-nil, zero value otherwise.

### GetThresholdComparatorOk

`func (o *ParamsEsQueryRuleOneOf) GetThresholdComparatorOk() (*string, bool)`

GetThresholdComparatorOk returns a tuple with the ThresholdComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdComparator

`func (o *ParamsEsQueryRuleOneOf) SetThresholdComparator(v string)`

SetThresholdComparator sets ThresholdComparator field to given value.


### GetTimeField

`func (o *ParamsEsQueryRuleOneOf) GetTimeField() string`

GetTimeField returns the TimeField field if non-nil, zero value otherwise.

### GetTimeFieldOk

`func (o *ParamsEsQueryRuleOneOf) GetTimeFieldOk() (*string, bool)`

GetTimeFieldOk returns a tuple with the TimeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeField

`func (o *ParamsEsQueryRuleOneOf) SetTimeField(v string)`

SetTimeField sets TimeField field to given value.

### HasTimeField

`func (o *ParamsEsQueryRuleOneOf) HasTimeField() bool`

HasTimeField returns a boolean if a field has been set.

### GetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf) GetTimeWindowSize() int32`

GetTimeWindowSize returns the TimeWindowSize field if non-nil, zero value otherwise.

### GetTimeWindowSizeOk

`func (o *ParamsEsQueryRuleOneOf) GetTimeWindowSizeOk() (*int32, bool)`

GetTimeWindowSizeOk returns a tuple with the TimeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowSize

`func (o *ParamsEsQueryRuleOneOf) SetTimeWindowSize(v int32)`

SetTimeWindowSize sets TimeWindowSize field to given value.


### GetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf) GetTimeWindowUnit() Timewindowunit`

GetTimeWindowUnit returns the TimeWindowUnit field if non-nil, zero value otherwise.

### GetTimeWindowUnitOk

`func (o *ParamsEsQueryRuleOneOf) GetTimeWindowUnitOk() (*Timewindowunit, bool)`

GetTimeWindowUnitOk returns a tuple with the TimeWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowUnit

`func (o *ParamsEsQueryRuleOneOf) SetTimeWindowUnit(v Timewindowunit)`

SetTimeWindowUnit sets TimeWindowUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


