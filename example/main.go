package main

import (
	"context"
	"fmt"
	"github.com/sarjon/toggler"
	"github.com/sarjon/toggler/operator"
	"github.com/sarjon/toggler/storage"
)

func main() {
	// Feature X is configured with "unanimous" strategy, meaning all conditions must be met.
	featureX := toggler.NewToggle("feature_x", toggler.NewOperatorCondition("ip", operator.NewEqual("127.0.0.1")))
	featureX.SetStrategy(toggler.StrategyUnanimous)

	// Feature Y is configured using multiple operators and boolean logic.
	featureY := toggler.NewToggle("feature_y",
		toggler.NewOperatorCondition("username", operator.NewOr(
			operator.NewEqual("super_admin"),
			operator.NewOr(
				operator.NewEqual("alpha_tester"),
				operator.NewEqual("beta_tester"),
			),
		)),
	)

	// Feature Z is configured to never be active despite its conditions.
	featureZ := toggler.NewToggle("feature_z", toggler.NewOperatorCondition("email", operator.NewEqual("dev@example.com")))
	featureZ.SetActivation(toggler.ActiveNever)

	// Store toggles in memory.
	togglesStorage, _ := storage.NewInMemory([]*toggler.Toggle{
		featureX,
		featureY,
		featureZ,
	})

	// Configure toggle manager with toggle storage and default strategies.
	toggleManager := toggler.NewManager(&toggler.ManagerArgs{
		Storage: togglesStorage,
	})

	// Build context with arbitrary values.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ip", "127.0.0.1")
	ctx = context.WithValue(ctx, "username", "beta_tester")
	ctx = context.WithValue(ctx, "email", "dev@example.com")

	// Use toggle manager to test if features are active or not.
	activeFeatureX, _ := toggleManager.Active("feature_x", ctx)
	activeFeatureY, _ := toggleManager.Active("feature_y", ctx)
	activeFeatureZ, _ := toggleManager.Active("feature_z", ctx)

	printFeature("X", activeFeatureX)
	printFeature("Y", activeFeatureY)
	printFeature("Z", activeFeatureZ)

	// Printed result will be:
	//
	// Feature X is active: true
	// Feature Y is active: true
	// Feature Z is active: false
}

func printFeature(name string, active bool) {
	fmt.Println(fmt.Sprintf("Feature %s is active: %t", name, active))
}
