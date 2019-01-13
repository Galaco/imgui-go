package imgui

const (
	TabBarFlagsDefault = 0
	// Allow manually dragging tabs to re-order them + New tabs are appended at the end of list
	TabBarFlagsReorderable = 1 << 0
	// Automatically select new tabs when they appear
	TabBarFlagsAutoSelectNewTabs = 1 << 1
	// Disable behavior of closing tabs (that are submitted with p_open != NULL) with middle mouse button. You can
	// still repro this behavior on user's side with if (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton   = 1 << 2
	TabBarFlagsNoTabListPopupButton           = 1 << 3
	TabBarFlagsNoTabListScrollingButtons      = 1 << 4
	// Resize tabs when they don't fit
	TabBarFlagsFittingPolicyResizeDown        = 1 << 5
	// Add scroll buttons when tabs don't fit
	TabBarFlagsFittingPolicyScroll            = 1 << 6
	TabBarFlagsFittingPolicyMask_ 			  = TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	TabBarFlagsFittingPolicyDefault_ 		  = TabBarFlagsFittingPolicyResizeDown
)