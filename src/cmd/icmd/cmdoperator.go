package icmd

// ICMDOperator はcmd.goを操作するインターフェース。
type ICMDOperator interface {
	// Run はコマンドライン引数を受け取ってプログラムを実行する。
	Run(args []string) error
}
