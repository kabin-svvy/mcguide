package main

import (
	"testing"
)

func TestUpdateMerchantGuideFailed(t *testing.T) {
	t.Run("glob is O", func(t *testing.T) {
		msg, _ := mcguide("glob is O")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("%s", msg)
		}
	})

	t.Run("glob is OO", func(t *testing.T) {
		msg, _ := mcguide("glob is OO")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("%s", msg)
		}
	})

	t.Run("glob is CC", func(t *testing.T) {
		msg, _ := mcguide("glob is CC")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("%s", msg)
		}
	})
}

func TestUpdateMaterialGuideFailed(t *testing.T) {

	t.Run("glob glob Silver is 34 Credits", func(t *testing.T) {
		msg, _ := mcguide("glob glob  Silver is 34 Credits")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("Expected Result I have no idea what you are talking about not : %s", msg)
		}
	})

	t.Run("glob prok Gold is 57800 Credits", func(t *testing.T) {
		msg, _ := mcguide("glob prok Gold is 57800 Credits")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("Expected Result I have no idea what you are talking about not : %s", msg)
		}
	})

	t.Run("pish pish Iron is 3910 Credits", func(t *testing.T) {
		msg, _ := mcguide("pish pish Iron is 3910 Credits")
		if msg != "I have no idea what you are talking about" {
			t.Errorf("Expected Result I have no idea what you are talking about not : %s", msg)
		}
	})
}

func TestUpdateMerchantGuide(t *testing.T) {
	t.Run("glob is I", func(t *testing.T) {
		_, err := mcguide("glob is I")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("prok is V", func(t *testing.T) {
		_, err := mcguide("prok is V")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish is X", func(t *testing.T) {
		_, err := mcguide("pish is X")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("tegj is L", func(t *testing.T) {
		_, err := mcguide("tegj is L")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})
}

func TestUpdateMaterialGuide(t *testing.T) {
	t.Run("glob is I", func(t *testing.T) {
		_, err := mcguide("glob is I")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("prok is V", func(t *testing.T) {
		_, err := mcguide("prok is V")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish is X", func(t *testing.T) {
		_, err := mcguide("pish is X")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("tegj is L", func(t *testing.T) {
		_, err := mcguide("tegj is L")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("glob glob Silver is 34 Credits", func(t *testing.T) {
		_, err := mcguide("glob glob Silver is 34 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("glob prok Gold is 57800 Credits", func(t *testing.T) {
		_, err := mcguide("glob prok Gold is 57800 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish pish Iron is 3910 Credits", func(t *testing.T) {
		_, err := mcguide("pish pish Iron is 3910 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})
}

func TestHowMuchGuide(t *testing.T) {
	t.Run("glob is I", func(t *testing.T) {
		_, err := mcguide("glob is I")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("prok is V", func(t *testing.T) {
		_, err := mcguide("prok is V")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish is X", func(t *testing.T) {
		_, err := mcguide("pish is X")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("tegj is L", func(t *testing.T) {
		_, err := mcguide("tegj is L")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("how much is pish tegj glob glob ?", func(t *testing.T) {
		msg, err := mcguide("how much is pish tegj glob glob ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "pish tegj glob glob is 42" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "pish tegj glob glob is 42", msg)
		}
	})

	t.Run("how much is glob pish ?", func(t *testing.T) {
		msg, err := mcguide("how much is glob pish ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "glob pish is 9" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "glob pish is 9", msg)
		}
	})

	t.Run("how much is pish glob glob glob ?", func(t *testing.T) {
		msg, err := mcguide("how much is pish glob glob glob ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "pish glob glob glob is 13" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "pish glob glob glob is 13", msg)
		}
	})

	t.Run("how much is pish tegj glob ?", func(t *testing.T) {
		msg, err := mcguide("how much is pish tegj glob ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "pish tegj glob is 41" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "pish tegj glob is 41", msg)
		}
	})

	t.Run("how much is tegj pish glob ?", func(t *testing.T) {
		msg, err := mcguide("how much is tegj pish glob ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "tegj pish glob is 61" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "tegj pish glob is 61", msg)
		}
	})

	t.Run("how much is tegj pish prok glob glob ?", func(t *testing.T) {
		msg, err := mcguide("how much is tegj pish prok glob glob ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "tegj pish prok glob glob is 67" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "tegj pish prok glob glob is 67", msg)
		}
	})
}

func TestHowManyGuide(t *testing.T) {
	t.Run("glob is I", func(t *testing.T) {
		_, err := mcguide("glob is I")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("prok is V", func(t *testing.T) {
		_, err := mcguide("prok is V")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish is X", func(t *testing.T) {
		_, err := mcguide("pish is X")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("tegj is L", func(t *testing.T) {
		_, err := mcguide("tegj is L")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("glob glob Silver is 34 Credits", func(t *testing.T) {
		_, err := mcguide("glob glob Silver is 34 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("glob prok Gold is 57800 Credits", func(t *testing.T) {
		_, err := mcguide("glob prok Gold is 57800 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("pish pish Iron is 3910 Credits", func(t *testing.T) {
		_, err := mcguide("pish pish Iron is 3910 Credits")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("how many Credits is glob prok Silver ?", func(t *testing.T) {
		msg, err := mcguide("how many Credits is glob prok Silver ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "glob prok Silver is 68 Credits" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "glob prok Silver is 68 Credits", msg)
		}
	})

	t.Run("how many Credits is glob prok Gold ?", func(t *testing.T) {
		msg, err := mcguide("how many Credits is glob prok Gold ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "glob prok Gold is 57800 Credits" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "glob prok Gold is 57800 Credits", msg)
		}
	})

	t.Run("how many Credits is glob prok Iron ?", func(t *testing.T) {
		msg, err := mcguide("how many Credits is glob prok Iron ?")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if msg != "glob prok Iron is 782 Credits" {
			t.Errorf("Expected shout \"%s\" not \"%s\"", "glob prok Iron is 782 Credits", msg)
		}
	})
}
