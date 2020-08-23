package kubernetes

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func dataSourceKubernetesPersistentVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesPersistentVolumeRead,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("persistent volume", true),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the desired characteristics of a volume requested by a pod author. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"azure_file": {
							Type:        schema.TypeList,
							Description: "The list of ports that are exposed by this service. More info: http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies",
							MinItems:    1,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"secret_name": {
										Type:        schema.TypeString,
										Description: "The secret name",
										Computed:    true,
									},
									"secret_namespace": {
										Type:        schema.TypeString,
										Description: "The namespace of secret",
										Computed:    true,
									},
									"share_name": {
										Type:        schema.TypeString,
										Description: "The share name on Azure Files service.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceKubernetesPersistentVolumeRead(d *schema.ResourceData, meta interface{}) error {
	metadata := expandMetadata(d.Get("metadata").([]interface{}))

	om := meta_v1.ObjectMeta{
		Namespace: metadata.Namespace,
		Name:      metadata.Name,
	}
	d.SetId(buildId(om))

	return resourceKubernetesPersistentVolumeRead(d, meta)
}
