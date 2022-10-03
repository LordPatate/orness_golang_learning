package exercise

type DeathListener interface {
	OnDeathOf(Mortal)
}

type Mortal interface {
	Die()
	AddOnDeathListener(DeathListener)
}

// A simple robot, prone to malfunction
type Robot struct {
	Id               int
	onDeathListeners []DeathListener
}

func (r *Robot) Die() {
	for _, listener := range r.onDeathListeners {
		listener.OnDeathOf(r)
	}
}

func (r *Robot) AddOnDeathListener(listener DeathListener) {
	r.onDeathListeners = append(r.onDeathListeners, listener)
}

// Manages an swarm of helper bots to keep the house clean
// Makes sure to replace them when they break
type Butler struct {
	helperBots map[int]*Robot
}

func NewButler() *Butler {
	newButler := &Butler{}
	newButler.helperBots = make(map[int]*Robot, 3)
	for i := range newButler.helperBots {
		newButler.helperBots[i] = &Robot{Id: i}
		newButler.helperBots[i].AddOnDeathListener(newButler)
	}
	return newButler
}

func (b *Butler) OnDeathOf(mortal Mortal) {
	robot := mortal.(*Robot)
	newRobot := &Robot{Id: robot.Id}
	b.helperBots[robot.Id] = newRobot
	newRobot.AddOnDeathListener(b)
}
