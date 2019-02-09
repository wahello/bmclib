# VirtualMedia

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**ConnectedVia** | [**ConnectedVia**](ConnectedVia.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Image** | **string** | A URI providing the location of the selected image. | [optional] 
**ImageName** | **string** | The current image name. | [optional] 
**Inserted** | **bool** | Indicates if virtual media is inserted in the virtual device. | [optional] 
**MediaTypes** | [**[]MediaType2**](MediaType_2.md) | This is the media types supported as virtual media. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Oem extension object. | [optional] 
**WriteProtected** | **bool** | Indicates the media is write protected. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


