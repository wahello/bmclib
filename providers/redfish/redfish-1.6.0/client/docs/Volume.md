# Volume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | [optional] 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataType** | **string** | The type of a resource. | [optional] 
**Actions** | [**Volume2Actions**](Volume_2_Actions.md) |  | [optional] 
**BlockSizeBytes** | **float32** | The size of the smallest addressible unit (Block) of this volume in bytes. | [optional] 
**CapacityBytes** | **float32** | The size in bytes of this Volume. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Encrypted** | **bool** | Is this Volume encrypted. | [optional] 
**EncryptionTypes** | [**[]EncryptionTypes**](EncryptionTypes.md) | The types of encryption used by this Volume. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Identifiers** | [**[]Identifier2**](Identifier_2.md) | The Durable names for the volume. | [optional] 
**Links** | [**Volume2Links**](Volume_2_Links.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Oem extension object. | [optional] 
**Operations** | [**[]Operations2**](Operations_2.md) | The operations currently running on the Volume. | [optional] 
**OptimumIOSizeBytes** | **float32** | The size in bytes of this Volume&#39;s optimum IO size. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**VolumeType** | [**VolumeType**](VolumeType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


