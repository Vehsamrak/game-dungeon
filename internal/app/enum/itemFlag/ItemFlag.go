package itemFlag

type Flag string

const (
	CutTreeTool     Flag = "cut_tree"
	FishTool        Flag = "fish_tool"
	MineTool        Flag = "mine_tool"
	ResourceWood    Flag = "resource_wood"
	ResourceFish    Flag = "resource_fish"
	ResourceOre     Flag = "resource_ore"
	Food            Flag = "food"
	IgnoreWaitstate Flag = "ignore_waitstate"
	CanFly          Flag = "can_fly"
	CliffWalk       Flag = "cliff_walk"
)
