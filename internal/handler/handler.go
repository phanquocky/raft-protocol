package handler

// Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:

// the method's type is exported.
// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.

type handler struct {
}