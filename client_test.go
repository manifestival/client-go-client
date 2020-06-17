package client_test

import (
	"testing"

	mfClient "github.com/manifestival/client-go-client"
	mf "github.com/manifestival/manifestival"

	"k8s.io/apimachinery/pkg/runtime"
	kFake "k8s.io/client-go/dynamic/fake"
)

func TestNewUnsafeDynamicClient(t *testing.T) {
	scheme := runtime.NewScheme()

	kfdc := kFake.NewSimpleDynamicClient(scheme)
	client, err := mfClient.NewUnsafeDynamicClient(kfdc)

	if err != nil {
		t.Fatalf("received client creation error %v", err)
	}

	source := mf.Slice{}
	_, err = mf.ManifestFrom(source, mf.UseClient(client))

	if err != nil {
		t.Fatalf("received creating manifest %v", err)
	}
	// TODO: validate calls with objects in fake dynamic client
}
