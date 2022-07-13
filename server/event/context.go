package event

// Context represents the context of an event. Handlers of an event may call methods on the context to change
// the result of the event.
type Context struct {
	cancel bool
}

type MessageContext struct {
	cancel  bool
	message string
	format string
}

// C returns a new event context.
func C() *Context {
	return &Context{}
}

// Cancelled returns whether the context has been cancelled.
func (ctx *Context) Cancelled() bool {
	return ctx.cancel
}

// Cancel cancels the context.
func (ctx *Context) Cancel() {
	ctx.cancel = true
}

func MC(m string, f string) *MessageContext {
	return &MessageContext{message: m, format: f}
}

func (ctx *MessageContext) Cancel() {
	ctx.cancel = true
}

func (ctx *MessageContext) Cancelled() bool {
	return ctx.cancel
}

func (ctx *MessageContext) SetMessage(m string) {
	ctx.message = m
}

func (ctx *MessageContext) Message() string {
	return ctx.message
}

func (ctx *MessageContext) SetFormat(f string) {
	ctx.format = f
}

func (ctx *MessageContext) Format() string {
	return ctx.format
}