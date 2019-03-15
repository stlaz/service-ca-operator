package operator

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

func (c *serviceCAOperator) setFailingStatus(operatorConfig *operatorv1.ServiceCA, reason, message string) {
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions,
		operatorv1.OperatorCondition{
			Type:    operatorv1.OperatorStatusTypeFailing,
			Status:  operatorv1.ConditionTrue,
			Reason:  reason,
			Message: message,
		})

	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{
		Type:   operatorv1.OperatorStatusTypeProgressing,
		Status: operatorv1.ConditionFalse,
	})

	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions,
		operatorv1.OperatorCondition{
			Type:   operatorv1.OperatorStatusTypeAvailable,
			Status: operatorv1.ConditionFalse,
		})
}

func (c *serviceCAOperator) setAvailableStatus(operatorConfig *operatorv1.ServiceCA) {
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{
		Type:   operatorv1.OperatorStatusTypeAvailable,
		Status: operatorv1.ConditionTrue,
	})

	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{
		Type:   operatorv1.OperatorStatusTypeProgressing,
		Status: operatorv1.ConditionFalse,
	})

	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{
		Type:   operatorv1.OperatorStatusTypeFailing,
		Status: operatorv1.ConditionFalse,
	})
}
