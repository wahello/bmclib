# EventDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Context** | **string** | A client-supplied string that is stored with the event destination subscription. | 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Destination** | **string** | The URI of the destination Event Service. | [optional] 
**EventFormatType** | [**EventFormatType**](EventFormatType.md) |  | [optional] 
**EventTypes** | [**[]EventType**](EventType.md) | This property contains the types of events that will be sent to the desination. | [optional] 
**HttpHeaders** | [**[]map[string]interface{}**](map[string]interface{}.md) | This is for setting HTTP headers, such as authorization information.  This object will be null on a GET. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**MessageIds** | **[]string** | A list of MessageIds that the service will only send.  If this property is absent or the array is empty, then Events with any MessageId will be sent to the subscriber. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Oem extension object. | [optional] 
**OriginResources** | [**[]IdRef**](idRef.md) | A list of resources for which the service will only send related events.  If this property is absent or the array is empty, then Events originating from any resource will be sent to the subscriber. | [optional] 
**OriginResourcesodataCount** | **int32** | The number of items in a collection. | [optional] 
**Protocol** | [**EventDestinationProtocol**](EventDestinationProtocol.md) |  | [optional] 
**RegistryPrefixes** | **[]string** | A list of the Prefixes for the Message Registries that contain the MessageIds that will be sent to this event destination. | [optional] 
**ResourceTypes** | **[]string** | A list of Resource Type values (Schema names) that correspond to the OriginOfCondition.  The version and full namespace should not be specified. | [optional] 
**SubordinateResources** | **bool** | By setting this to true and specifying OriginResources, this indicates the subscription will be for events from the OriginsResources specified and also all subordinate resources.  Note that resources associated via the Links section are not considered subordinate. | [optional] 
**SubscriptionType** | [**SubscriptionType**](SubscriptionType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


