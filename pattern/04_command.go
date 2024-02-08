package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// цель паттерна инкапсулировать запрос
// не нужно создавать различные обработчики для каждого отправителя
// отделяет пользовательский интерфейс от связанной с ним бизнес-логики

type Command interface {
	Execute()
}

type PlayerInterface interface {
	Move()
	Attack()
}

type button struct {
	command Command
}

func NewButton(command Command) *button {
	return &button{command: command}
}

func (b *button) Press() {
	b.command.Execute()
}

type moveCommand struct {
	player PlayerInterface
}

func NewMoveCommand(player PlayerInterface) *moveCommand {
	return &moveCommand{
		player: player,
	}
}

func (c *moveCommand) Execute() {
	c.player.Move()
}

type attackCommand struct {
	player PlayerInterface
}

func NewattackCommand(player PlayerInterface) *attackCommand {
	return &attackCommand{
		player: player,
	}
}

func (c *attackCommand) Execute() {
	c.player.Attack()
}

type Player struct{}

func (p *Player) Move() {
	println("Move")
}

func (p *Player) Attack() {
	println("Attack")
}

func CmdPattern() {
	p := &Player{}
	mc := NewMoveCommand(p)
	ac := NewattackCommand(p)
	attackButton := NewButton(ac)
	moveButton := NewButton(mc)
	attackButton.Press()
	moveButton.Press()
}
