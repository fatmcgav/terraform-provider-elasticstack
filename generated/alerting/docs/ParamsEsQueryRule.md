# ParamsEsQueryRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggField** | Pointer to **string** | The name of the numeric field that is used in the aggregation. This property is required when &#x60;aggType&#x60; is &#x60;avg&#x60;, &#x60;max&#x60;, &#x60;min&#x60; or &#x60;sum&#x60;.  | [optional] 
**AggType** | Pointer to [**Aggtype**](Aggtype.md) |  | [optional] [default to COUNT]
**EsqlQuery** | [**ParamsEsQueryRuleOneOfEsqlQuery**](ParamsEsQueryRuleOneOfEsqlQuery.md) |  | 
**ExcludeHitsFromPreviousRun** | Pointer to **bool** | Indicates whether to exclude matches from previous runs. If &#x60;true&#x60;, you can avoid alert duplication by excluding documents that have already been detected by the previous rule run. This option is not available when a grouping field is specified.  | [optional] 
**GroupBy** | Pointer to [**Groupby**](Groupby.md) |  | [optional] [default to ALL]
**SearchType** | **string** | The type of query, in this case a query that uses Elasticsearch Query DSL. | [default to "esQuery"]
**Size** | **int32** | The number of documents to pass to the configured actions when the threshold condition is met.  | 
**TermSize** | Pointer to **int32** | This property is required when &#x60;groupBy&#x60; is &#x60;top&#x60;. It specifies the number of groups to check against the threshold and therefore limits the number of alerts on high cardinality fields.  | [optional] 
**Threshold** | **[]int32** | The threshold value that is used with the &#x60;thresholdComparator&#x60;. If the &#x60;thresholdComparator&#x60; is &#x60;between&#x60; or &#x60;notBetween&#x60;, you must specify the boundary values.  | 
**ThresholdComparator** | [**Thresholdcomparator**](Thresholdcomparator.md) |  | 
**TimeField** | **string** | The field that is used to calculate the time window. | 
**TimeWindowSize** | **int32** | The size of the time window (in &#x60;timeWindowUnit&#x60; units), which determines how far back to search for documents. Generally it should be a value higher than the rule check interval to avoid gaps in detection.  | 
**TimeWindowUnit** | [**Timewindowunit**](Timewindowunit.md) |  | 
**SearchConfiguration** | Pointer to [**ParamsEsQueryRuleOneOf1SearchConfiguration**](ParamsEsQueryRuleOneOf1SearchConfiguration.md) |  | [optional] 
**TermField** | Pointer to [**Termfield**](Termfield.md) |  | [optional] 
**EsQuery** | **string** | The query definition, which uses Elasticsearch Query DSL. | 
**Index** | [**ParamsEsQueryRuleOneOf2Index**](ParamsEsQueryRuleOneOf2Index.md) |  | 

## Methods

### NewParamsEsQueryRule

`func NewParamsEsQueryRule(esqlQuery ParamsEsQueryRuleOneOfEsqlQuery, searchType string, size int32, threshold []int32, thresholdComparator Thresholdcomparator, timeField string, timeWindowSize int32, timeWindowUnit Timewindowunit, esQuery string, index ParamsEsQueryRuleOneOf2Index, ) *ParamsEsQueryRule`

NewParamsEsQueryRule instantiates a new ParamsEsQueryRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsEsQueryRuleWithDefaults

`func NewParamsEsQueryRuleWithDefaults() *ParamsEsQueryRule`

NewParamsEsQueryRuleWithDefaults instantiates a new ParamsEsQueryRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggField

`func (o *ParamsEsQueryRule) GetAggField() string`

GetAggField returns the AggField field if non-nil, zero value otherwise.

### GetAggFieldOk

`func (o *ParamsEsQueryRule) GetAggFieldOk() (*string, bool)`

GetAggFieldOk returns a tuple with the AggField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggField

`func (o *ParamsEsQueryRule) SetAggField(v string)`

SetAggField sets AggField field to given value.

### HasAggField

`func (o *ParamsEsQueryRule) HasAggField() bool`

HasAggField returns a boolean if a field has been set.

### GetAggType

`func (o *ParamsEsQueryRule) GetAggType() Aggtype`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *ParamsEsQueryRule) GetAggTypeOk() (*Aggtype, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *ParamsEsQueryRule) SetAggType(v Aggtype)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *ParamsEsQueryRule) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetEsqlQuery

`func (o *ParamsEsQueryRule) GetEsqlQuery() ParamsEsQueryRuleOneOfEsqlQuery`

GetEsqlQuery returns the EsqlQuery field if non-nil, zero value otherwise.

### GetEsqlQueryOk

`func (o *ParamsEsQueryRule) GetEsqlQueryOk() (*ParamsEsQueryRuleOneOfEsqlQuery, bool)`

GetEsqlQueryOk returns a tuple with the EsqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsqlQuery

`func (o *ParamsEsQueryRule) SetEsqlQuery(v ParamsEsQueryRuleOneOfEsqlQuery)`

SetEsqlQuery sets EsqlQuery field to given value.


### GetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRule) GetExcludeHitsFromPreviousRun() bool`

GetExcludeHitsFromPreviousRun returns the ExcludeHitsFromPreviousRun field if non-nil, zero value otherwise.

### GetExcludeHitsFromPreviousRunOk

`func (o *ParamsEsQueryRule) GetExcludeHitsFromPreviousRunOk() (*bool, bool)`

GetExcludeHitsFromPreviousRunOk returns a tuple with the ExcludeHitsFromPreviousRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRule) SetExcludeHitsFromPreviousRun(v bool)`

SetExcludeHitsFromPreviousRun sets ExcludeHitsFromPreviousRun field to given value.

### HasExcludeHitsFromPreviousRun

`func (o *ParamsEsQueryRule) HasExcludeHitsFromPreviousRun() bool`

HasExcludeHitsFromPreviousRun returns a boolean if a field has been set.

### GetGroupBy

`func (o *ParamsEsQueryRule) GetGroupBy() Groupby`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ParamsEsQueryRule) GetGroupByOk() (*Groupby, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ParamsEsQueryRule) SetGroupBy(v Groupby)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ParamsEsQueryRule) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetSearchType

`func (o *ParamsEsQueryRule) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *ParamsEsQueryRule) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *ParamsEsQueryRule) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.


### GetSize

`func (o *ParamsEsQueryRule) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ParamsEsQueryRule) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ParamsEsQueryRule) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTermSize

`func (o *ParamsEsQueryRule) GetTermSize() int32`

GetTermSize returns the TermSize field if non-nil, zero value otherwise.

### GetTermSizeOk

`func (o *ParamsEsQueryRule) GetTermSizeOk() (*int32, bool)`

GetTermSizeOk returns a tuple with the TermSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSize

`func (o *ParamsEsQueryRule) SetTermSize(v int32)`

SetTermSize sets TermSize field to given value.

### HasTermSize

`func (o *ParamsEsQueryRule) HasTermSize() bool`

HasTermSize returns a boolean if a field has been set.

### GetThreshold

`func (o *ParamsEsQueryRule) GetThreshold() []int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ParamsEsQueryRule) GetThresholdOk() (*[]int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ParamsEsQueryRule) SetThreshold(v []int32)`

SetThreshold sets Threshold field to given value.


### GetThresholdComparator

`func (o *ParamsEsQueryRule) GetThresholdComparator() Thresholdcomparator`

GetThresholdComparator returns the ThresholdComparator field if non-nil, zero value otherwise.

### GetThresholdComparatorOk

`func (o *ParamsEsQueryRule) GetThresholdComparatorOk() (*Thresholdcomparator, bool)`

GetThresholdComparatorOk returns a tuple with the ThresholdComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdComparator

`func (o *ParamsEsQueryRule) SetThresholdComparator(v Thresholdcomparator)`

SetThresholdComparator sets ThresholdComparator field to given value.


### GetTimeField

`func (o *ParamsEsQueryRule) GetTimeField() string`

GetTimeField returns the TimeField field if non-nil, zero value otherwise.

### GetTimeFieldOk

`func (o *ParamsEsQueryRule) GetTimeFieldOk() (*string, bool)`

GetTimeFieldOk returns a tuple with the TimeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeField

`func (o *ParamsEsQueryRule) SetTimeField(v string)`

SetTimeField sets TimeField field to given value.


### GetTimeWindowSize

`func (o *ParamsEsQueryRule) GetTimeWindowSize() int32`

GetTimeWindowSize returns the TimeWindowSize field if non-nil, zero value otherwise.

### GetTimeWindowSizeOk

`func (o *ParamsEsQueryRule) GetTimeWindowSizeOk() (*int32, bool)`

GetTimeWindowSizeOk returns a tuple with the TimeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowSize

`func (o *ParamsEsQueryRule) SetTimeWindowSize(v int32)`

SetTimeWindowSize sets TimeWindowSize field to given value.


### GetTimeWindowUnit

`func (o *ParamsEsQueryRule) GetTimeWindowUnit() Timewindowunit`

GetTimeWindowUnit returns the TimeWindowUnit field if non-nil, zero value otherwise.

### GetTimeWindowUnitOk

`func (o *ParamsEsQueryRule) GetTimeWindowUnitOk() (*Timewindowunit, bool)`

GetTimeWindowUnitOk returns a tuple with the TimeWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowUnit

`func (o *ParamsEsQueryRule) SetTimeWindowUnit(v Timewindowunit)`

SetTimeWindowUnit sets TimeWindowUnit field to given value.


### GetSearchConfiguration

`func (o *ParamsEsQueryRule) GetSearchConfiguration() ParamsEsQueryRuleOneOf1SearchConfiguration`

GetSearchConfiguration returns the SearchConfiguration field if non-nil, zero value otherwise.

### GetSearchConfigurationOk

`func (o *ParamsEsQueryRule) GetSearchConfigurationOk() (*ParamsEsQueryRuleOneOf1SearchConfiguration, bool)`

GetSearchConfigurationOk returns a tuple with the SearchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchConfiguration

`func (o *ParamsEsQueryRule) SetSearchConfiguration(v ParamsEsQueryRuleOneOf1SearchConfiguration)`

SetSearchConfiguration sets SearchConfiguration field to given value.

### HasSearchConfiguration

`func (o *ParamsEsQueryRule) HasSearchConfiguration() bool`

HasSearchConfiguration returns a boolean if a field has been set.

### GetTermField

`func (o *ParamsEsQueryRule) GetTermField() Termfield`

GetTermField returns the TermField field if non-nil, zero value otherwise.

### GetTermFieldOk

`func (o *ParamsEsQueryRule) GetTermFieldOk() (*Termfield, bool)`

GetTermFieldOk returns a tuple with the TermField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermField

`func (o *ParamsEsQueryRule) SetTermField(v Termfield)`

SetTermField sets TermField field to given value.

### HasTermField

`func (o *ParamsEsQueryRule) HasTermField() bool`

HasTermField returns a boolean if a field has been set.

### GetEsQuery

`func (o *ParamsEsQueryRule) GetEsQuery() string`

GetEsQuery returns the EsQuery field if non-nil, zero value otherwise.

### GetEsQueryOk

`func (o *ParamsEsQueryRule) GetEsQueryOk() (*string, bool)`

GetEsQueryOk returns a tuple with the EsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQuery

`func (o *ParamsEsQueryRule) SetEsQuery(v string)`

SetEsQuery sets EsQuery field to given value.


### GetIndex

`func (o *ParamsEsQueryRule) GetIndex() ParamsEsQueryRuleOneOf2Index`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ParamsEsQueryRule) GetIndexOk() (*ParamsEsQueryRuleOneOf2Index, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ParamsEsQueryRule) SetIndex(v ParamsEsQueryRuleOneOf2Index)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


