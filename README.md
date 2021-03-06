# kubebuilder-demo
Just messing around with kubebuilder as a method of generating custom resource definitions for use as controllers or operators!

## Why this tool is useful

### Setting things up for our controller
![alt text](https://github.com/apaz037/kubebuilder-demo/raw/master/hack/images/setting-things-up.png "applying our generated install.yaml")

### Creating an Object from our Custom Resource Definition
![alt text](https://github.com/apaz037/kubebuilder-demo/raw/master/hack/images/creating-a-custom-resource.png "creating a CR from our CRD")

### Modifying Resources with our Controller/Operator
![alt text](https://github.com/apaz037/kubebuilder-demo/raw/master/hack/images/modifying-a-resource-with-our-controller.png "When our controller's reconcile function runs and finds any sloths, we mark them with a message")

## Links
- [Docker Hub Image](https://hub.docker.com/r/aaronpaz/kubebuilder-demo/)
- [Under The Hood of Kubebuilder](https://itnext.io/under-the-hood-of-kubebuilder-framework-ff6b38c10796)

## Notes
"When simply deploying is not enough is when I start to scratch my head wondering whether its time to start writing a controller or an operator." - Kris Nova

"A controller is a very specific declarative piece of software that reconciles state and an operator is just a domain specific controller." - Joe Beda
