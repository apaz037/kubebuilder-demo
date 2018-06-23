package sloth

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	mammalsv1beta1 "kubebuilder-demo/pkg/apis/mammals/v1beta1"
	mammalsv1beta1client "kubebuilder-demo/pkg/client/clientset/versioned/typed/mammals/v1beta1"
	mammalsv1beta1informer "kubebuilder-demo/pkg/client/informers/externalversions/mammals/v1beta1"
	mammalsv1beta1lister "kubebuilder-demo/pkg/client/listers/mammals/v1beta1"

	"kubebuilder-demo/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Sloth resources goes here.

func (bc *SlothController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE	
	if !bc.logged {
		log.Printf("Implement the Reconcile function on sloth.SlothController to reconcile %s\n", k.Name)
		bc.logged = true
	}
	return nil
}

// +kubebuilder:controller:group=mammals,version=v1beta1,kind=Sloth,resource=sloths
type SlothController struct {
	// INSERT ADDITIONAL FIELDS HERE
	slothLister mammalsv1beta1lister.SlothLister
	slothclient mammalsv1beta1client.MammalsV1beta1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	slothrecorder record.EventRecorder

	// After we log, we flip this bit
	logged bool
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &SlothController{
		slothLister: arguments.ControllerManager.GetInformerProvider(&mammalsv1beta1.Sloth{}).(mammalsv1beta1informer.SlothInformer).Lister(),

		slothclient:   arguments.Clientset.MammalsV1beta1(),
		slothrecorder: arguments.CreateRecorder("SlothController"),
	}

	// Create a new controller that will call SlothController.Reconcile on changes to Sloths
	gc := &controller.GenericController{
		Name:             "SlothController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&mammalsv1beta1.Sloth{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Sloth Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the SlothController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
