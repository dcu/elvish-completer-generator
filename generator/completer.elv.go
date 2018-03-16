package generator

const (
	_templateContent = `
fn spaces [n]{
    repeat $n ' ' | joins ''
}

fn cand [text desc]{
    edit:complex-candidate $text &display-suffix=' '(spaces (- 14 (wcswidth $text)))$desc
}

subcmds~ = (constantly (
{{ range .SubCommands }}
    cand {{ .Name }} {{ .Description | quote }}
{{ end }}
))

build-flags~ = (constantly (
{{ range .Flags }}
    cand {{ .Name }} {{ .Description | quote }}
{{ end }}
))

fn -is-flag [f]{
    has-prefix $f -
}

fn list-files [curr]{
	put $curr*[match-hidden][nomatch-ok]
}

fn compl [cmd @words]{
	curr = $words[-1]
	if (-is-flag $curr) {
		build-flags
	}
	list-files $curr
	subcmds
}

edit:completion:arg-completer[{{.CommandName}}] = $compl~
`
)
