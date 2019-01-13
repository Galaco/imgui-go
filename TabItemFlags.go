package imgui

const (
	TabItemFlagsDefault = 0
	// TabItemFlagsUnsavedDocument Appends '*' to title without affecting the ID, as a convenience to avoid using
	// the ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo
	// it without flicker.
	TabItemFlagsUnsavedDocument               = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programatically make the tab selected when calling BeginTabItem()
	TabItemFlagsSetSelected                   = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton Disables behavior of closing tabs (that are submitted with p_open != NULL) with middle
	// mouse button. You can still repro this behavior on user's side with if (IsItemHovered() && IsMouseClicked(2))
	// *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton  = 1 << 2
	// TabItemFlagsNoPushId Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
	TabItemFlagsNoPushId                      = 1 << 3
)
