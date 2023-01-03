package modelos

// Representa o formato da requição de alteração de senha
type Senha struct {
	Nova  string `json:nova`
	Atual string `json:atual`
}
