//go:build !darwin

package main

// setActivationPolicy hides the app from the dock, which is macOS only.
func setActivationPolicy() {
}
