package session

type GameMode struct{}

func (GameMode) AllowsEditing() bool      { return false }
func (GameMode) AllowsTakingDamage() bool { return false }
func (GameMode) CreativeInventory() bool  { return false }
func (GameMode) HasCollision() bool       { return true }
func (GameMode) AllowsFlying() bool       { return false }
func (GameMode) AllowsInteraction() bool  { return true }
func (GameMode) Visible() bool            { return true }
