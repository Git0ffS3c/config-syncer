package watcher

import (
	"errors"
	"fmt"

	acrt "github.com/appscode/go/runtime"
	"github.com/appscode/log"
	tapi "github.com/k8sdb/apimachinery/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

// Blocks caller. Intended to be called as a Go routine.
func (c *Controller) WatchPostgreses() {
	defer acrt.HandleCrash()

	lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return c.KubeDBClient.Postgreses(apiv1.NamespaceAll).List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return c.KubeDBClient.Postgreses(apiv1.NamespaceAll).Watch(metav1.ListOptions{})
		},
	}
	_, ctrl := cache.NewInformer(lw,
		&tapi.Postgres{},
		c.SyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				if alert, ok := obj.(*tapi.Postgres); ok {
					fmt.Println(alert)
				}
			},
			UpdateFunc: func(old, new interface{}) {
				oldAlert, ok := old.(*tapi.Postgres)
				if !ok {
					log.Errorln(errors.New("Invalid Postgres object"))
					return
				}
				newAlert, ok := new.(*tapi.Postgres)
				if !ok {
					log.Errorln(errors.New("Invalid Postgres object"))
					return
				}
				fmt.Println(oldAlert, newAlert)
			},
			DeleteFunc: func(obj interface{}) {
				if alert, ok := obj.(*tapi.Postgres); ok {
					fmt.Println(alert)
				}
			},
		},
	)
	ctrl.Run(wait.NeverStop)
}