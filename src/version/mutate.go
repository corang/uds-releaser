package version

import (
	"github.com/corang/uds-releaser/src/types"
	"github.com/corang/uds-releaser/src/utils"
	uds "github.com/defenseunicorns/uds-cli/src/types"
	zarf "github.com/zarf-dev/zarf/src/api/v1alpha1"
)

func MutateYamls(flavor types.Flavor) error {
	packageName, err := mutateZarfYaml(flavor)
	if err != nil {
		return err
	}

	err = mutateBundleYaml(flavor, packageName)
	if err != nil {
		return err
	}
	return nil
}

func mutateZarfYaml(flavor types.Flavor) (packageName string, err error) {
	var zarfPackage zarf.ZarfPackage
	err = utils.LoadYaml("zarf.yaml", &zarfPackage)
	if err != nil {
		return "", err
	}

	zarfPackage.Metadata.Version = flavor.Version

	err = utils.UpdateYaml("zarf.yaml", zarfPackage)
	if err != nil {
		return zarfPackage.Metadata.Name, err
	}

	return zarfPackage.Metadata.Name, nil
}

func mutateBundleYaml(flavor types.Flavor, packageName string) error {
	var bundle uds.UDSBundle
	err := utils.LoadYaml("bundle/uds-bundle.yaml", &bundle)
	if err != nil {
		return err
	}

	bundle.Metadata.Version = flavor.Version

	// Find the package that matches the package name and update its ref
	for i, bundledPackage := range bundle.Packages {
		if bundledPackage.Name == packageName {
			bundle.Packages[i].Ref = flavor.Version
		}
	}

	err = utils.UpdateYaml("bundle/uds-bundle.yaml", bundle)
	if err != nil {
		return err
	}
	return nil
}
