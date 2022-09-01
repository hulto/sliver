package accessibility

// Code generated by cdproto-gen. DO NOT EDIT.

// EventLoadComplete the loadComplete event mirrors the load complete event
// sent by the browser to assistive technology when the web page has finished
// loading.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#event-loadComplete
type EventLoadComplete struct {
	Root *Node `json:"root"` // New document root node.
}

// EventNodesUpdated the nodesUpdated event is sent every time a previously
// requested node has changed the in tree.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#event-nodesUpdated
type EventNodesUpdated struct {
	Nodes []*Node `json:"nodes"` // Updated node data.
}