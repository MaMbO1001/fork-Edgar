package app

type App struct {
	jopa Jopa
}

func Mew(j Jopa) *App {
	return &App{
		jopa: j,
	}
}
