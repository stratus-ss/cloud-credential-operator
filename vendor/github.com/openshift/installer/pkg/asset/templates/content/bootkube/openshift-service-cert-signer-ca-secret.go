package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftServiceCertSignerSecretFileName = "openshift-service-cert-signer-ca-secret.yaml.template"
)

var _ asset.WritableAsset = (*OpenshiftServiceCertSignerSecret)(nil)

// OpenshiftServiceCertSignerSecret is the constant to represent the contents of openshift-service-signer-secret.yaml.template
type OpenshiftServiceCertSignerSecret struct {
	fileName string
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *OpenshiftServiceCertSignerSecret) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *OpenshiftServiceCertSignerSecret) Name() string {
	return "OpenshiftServiceCertSignerSecret"
}

// Generate generates the actual files by this asset
func (t *OpenshiftServiceCertSignerSecret) Generate(parents asset.Parents) error {
	t.fileName = openshiftServiceCertSignerSecretFileName
	data, err := content.GetBootkubeTemplate(t.fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, t.fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *OpenshiftServiceCertSignerSecret) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *OpenshiftServiceCertSignerSecret) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftServiceCertSignerSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
