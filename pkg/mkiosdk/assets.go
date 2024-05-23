package mkiosdk

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Asset - An Asset.
type Asset struct {
	// The resource properties.
	Properties *AssetProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *armmediaservices.SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AssetProperties - The Asset properties.
type AssetProperties struct {
	// The alternate ID of the Asset.
	AlternateID *string `json:"alternateId,omitempty"`

	// The name of the asset blob container.
	Container *string `json:"container,omitempty"`

	ContainerDeletionPolicy *ContainerDeletePolicy `json:"containerDeletionPolicy,omitempty"`

	// The Asset description.
	Description *string `json:"description,omitempty"`

	// The name of the storage account.
	StorageAccountName *string `json:"storageAccountName,omitempty"`

	// READ-ONLY; The Asset ID.
	AssetID *string `json:"assetId,omitempty" azure:"ro"`

	// READ-ONLY; The creation date of the Asset.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; The last modified date of the Asset.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat *armmediaservices.AssetStorageEncryptionFormat `json:"storageEncryptionFormat,omitempty" azure:"ro"`
}

type ContainerDeletePolicy string

const (
	Retain ContainerDeletePolicy = "Retain"
	Delete ContainerDeletePolicy = "Delete"
)
