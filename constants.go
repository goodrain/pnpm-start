package pnpmstart

const (
	Node        = "node"
	NodeModules = "node_modules"
	Pnpm        = "pnpm"
)

const StartupScript = `trap 'kill -TERM $CPID' TERM
trap 'kill -INT $CPID' INT
( %s ) &
CPID="$!"
wait $CPID
trap - TERM INT
wait $CPID
`
