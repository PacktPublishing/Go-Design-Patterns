package creational

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The memory positions of the shirts are different %p != %p \n\n", &shirt1, &shirt2)

}
