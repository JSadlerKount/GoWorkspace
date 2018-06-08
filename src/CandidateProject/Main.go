/*

Created By: Jeff Sadler

Note: Govatar library must be installed on local PC for this to work!
	  See https://github.com/o1egl/govatar
	  go get -u github.com/o1egl/govatar/...

This file creates three wrapper functions for the govatar.Generate function found
in the  Govatar library. These wrapper functions makes it easier to create a male
and female Govatar.

*/

package main

import (
	"fmt"
	"image"

	"github.com/o1egl/govatar"
)

// GovatarGenerateMale generates a male govatar
func GovatarGenerateMale() (image.Image, error) {
	img, err := govatar.Generate(govatar.MALE)
	return img, err
}

// GovatarGenerateFemale generates a female govatar
func GovatarGenerateFemale() (image.Image, error) {
	img, err := govatar.Generate(govatar.FEMALE)
	return img, err
}

// GovatarGenerate generates a govatar with gender decided by passed in parameter
func GovatarGenerate(gender govatar.Gender) (image.Image, error) {
	img, err := govatar.Generate(gender)
	return img, err
}

// Main execution part of app
func main() {
	// Let's generate a female govatar.
	if _, err := GovatarGenerateFemale(); err != nil {
		fmt.Println(fmt.Errorf("Error(s) occurred during Govatar generate. Error: %s", err))
	} else {
		fmt.Println("Govatar generated succesfully.")
	}
}
