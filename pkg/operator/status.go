package operator

import (
	"time"

	scsv1 "github.com/openshift/service-ca-operator/pkg/apis/serviceca/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *serviceCAOperator) setFailingStatus(operatorConfig *scsv1.ServiceCA, reason, message string) {
	setOperatorCondition(&operatorConfig.Status.Conditions,
		scsv1.OperatorCondition{
			Type:    scsv1.OperatorStatusTypeFailing,
			Status:  scsv1.ConditionTrue,
			Reason:  reason,
			Message: message,
		})

	setOperatorCondition(&operatorConfig.Status.Conditions, scsv1.OperatorCondition{
		Type:   scsv1.OperatorStatusTypeProgressing,
		Status: scsv1.ConditionFalse,
	})

	setOperatorCondition(&operatorConfig.Status.Conditions,
		scsv1.OperatorCondition{
			Type:   scsv1.OperatorStatusTypeAvailable,
			Status: scsv1.ConditionFalse,
		})
}

func (c *serviceCAOperator) setAvailableStatus(operatorConfig *scsv1.ServiceCA) {
	setOperatorCondition(&operatorConfig.Status.Conditions, scsv1.OperatorCondition{
		Type:   scsv1.OperatorStatusTypeAvailable,
		Status: scsv1.ConditionTrue,
	})

	setOperatorCondition(&operatorConfig.Status.Conditions, scsv1.OperatorCondition{
		Type:   scsv1.OperatorStatusTypeProgressing,
		Status: scsv1.ConditionFalse,
	})

	setOperatorCondition(&operatorConfig.Status.Conditions, scsv1.OperatorCondition{
		Type:   scsv1.OperatorStatusTypeFailing,
		Status: scsv1.ConditionFalse,
	})
}

// The below is a copy of functions from library-go
// FIMXE: the configuration from top-level objects needs to use openshift/api
//        but currently we don't have the controller types in there
func setOperatorCondition(conditions *[]scsv1.OperatorCondition, newCondition scsv1.OperatorCondition) {
	if conditions == nil {
		conditions = &[]scsv1.OperatorCondition{}
	}
	existingCondition := findOperatorCondition(*conditions, newCondition.Type)
	if existingCondition == nil {
		newCondition.LastTransitionTime = metav1.NewTime(time.Now())
		*conditions = append(*conditions, newCondition)
		return
	}

	if existingCondition.Status != newCondition.Status {
		existingCondition.Status = newCondition.Status
		existingCondition.LastTransitionTime = metav1.NewTime(time.Now())
	}

	existingCondition.Reason = newCondition.Reason
	existingCondition.Message = newCondition.Message
}

func findOperatorCondition(conditions []scsv1.OperatorCondition, conditionType string) *scsv1.OperatorCondition {
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return &conditions[i]
		}
	}

	return nil
}
