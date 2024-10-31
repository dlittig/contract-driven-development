# openapi_client.PetsApi

All URIs are relative to *http://petstore.swagger.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_pets**](PetsApi.md#create_pets) | **POST** /pets | Create a pet
[**list_pets**](PetsApi.md#list_pets) | **GET** /pets | List all pets
[**show_pet_by_id**](PetsApi.md#show_pet_by_id) | **GET** /pets/{petId} | Info for a specific pet


# **create_pets**
> create_pets(pet)

Create a pet

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.pet import Pet
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://petstore.swagger.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://petstore.swagger.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.PetsApi(api_client)
    pet = openapi_client.Pet() # Pet | 

    try:
        # Create a pet
        api_instance.create_pets(pet)
    except Exception as e:
        print("Exception when calling PetsApi->create_pets: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Null response |  -  |
**0** | unexpected error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_pets**
> List[Pet] list_pets(limit=limit)

List all pets

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.pet import Pet
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://petstore.swagger.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://petstore.swagger.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.PetsApi(api_client)
    limit = 56 # int | How many items to return at one time (max 100) (optional)

    try:
        # List all pets
        api_response = api_instance.list_pets(limit=limit)
        print("The response of PetsApi->list_pets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PetsApi->list_pets: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int**| How many items to return at one time (max 100) | [optional] 

### Return type

[**List[Pet]**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A paged array of pets |  * x-next - A link to the next page of responses <br>  |
**0** | unexpected error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **show_pet_by_id**
> Pet show_pet_by_id(pet_id)

Info for a specific pet

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.pet import Pet
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://petstore.swagger.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://petstore.swagger.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.PetsApi(api_client)
    pet_id = 'pet_id_example' # str | The id of the pet to retrieve

    try:
        # Info for a specific pet
        api_response = api_instance.show_pet_by_id(pet_id)
        print("The response of PetsApi->show_pet_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PetsApi->show_pet_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet_id** | **str**| The id of the pet to retrieve | 

### Return type

[**Pet**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Expected response to a valid request |  -  |
**0** | unexpected error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

