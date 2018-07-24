package template

import (
	"github.com/openshift/api/template/v1"
	"github.com/syndesisio/syndesis-operator/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)


func GetDeclaredResourceTypes() ([]metav1.TypeMeta, error) {
	types := make(map[metav1.TypeMeta]bool)
	res, err := util.LoadKubernetesResourceFromAsset("template.yaml")
	if err != nil {
		return nil, err
	}
	if templ, ok := res.(*v1.Template); ok {
		for _, obj := range templ.Objects {
			u := unstructured.Unstructured{}
			err := u.UnmarshalJSON(obj.Raw)
			if err != nil {
				return nil, err
			}
			meta := metav1.TypeMeta{
				APIVersion: u.GetAPIVersion(),
				Kind: u.GetKind(),
			}
			types[meta] = true
		}

		ret := make([]metav1.TypeMeta, len(types))
		i := 0
		for k := range types {
			ret[i] = k
			i++
		}
		return ret, nil
	}
	return nil, nil
}
