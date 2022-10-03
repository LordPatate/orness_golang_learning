package exercise

import "testing"

func TestAristocratFamily(t *testing.T) {
	mother := NewAristocrat("Juliet")
	father := NewAristocrat("Romeo")
	father.FallInLove(mother)
	mother.Heirs = []*Aristocrat{
		NewAristocrat("Joe"),
		NewAristocrat("Jack"),
		NewAristocrat("William"),
		NewAristocrat("Averell"),
	}
	father.Heirs = make([]*Aristocrat, 4)
	copy(father.Heirs, mother.Heirs)
	mother.Money = 16
	father.Money = 12
	notary := &Notary{}
	mother.AddOnDeathListener(notary)
	father.AddOnDeathListener(notary)
	mother.Die()
	for _, heir := range mother.Heirs {
		if heir.Money != 7 {
			t.Errorf("Each heir should get 4 money from their mother and 3 from their father, got %v", heir.Money)
		}
	}
}
