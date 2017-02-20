package visitor

import "testing"

type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T){
		msg := MessageA{
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if testHelper.Received !=  expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	t.Run("MessageB test", func(t *testing.T){
		msg := MessageB {
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"
		if testHelper.Received !=  expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})
}
