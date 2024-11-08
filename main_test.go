package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPanicAndRecover(t *testing.T) {
	// 標準出力をキャプチャするための設定
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// テスト対象の関数を実行
	main()

	// 標準出力のキャプチャを終了
	w.Close()
	os.Stdout = old

	// キャプチャした出力を読み込む
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 期待される出力の検証
	expected := []string{
		"プログラム開始",
		"deepestFunctionでリカバリー: deepestFunctionでパニックが発生",
		"deepestFunctionのdeferが実行されました",
		"deeperFunctionの残りのコード",
		"deeperFunctionのdeferが実行されました",
		"innerFunctionでリカバリー: deeperFunctionでパニックが発生",
		"innerFunctionのdeferが実行されました",
		"outerFunctionでリカバリー: deeperFunctionでパニックが発生",
		"プログラム終了",
	}

	// 出力の順序も検証
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != len(expected) {
		t.Errorf("期待される行数: %d, 実際の行数: %d", len(expected), len(lines))
		t.Errorf("実際の出力:\n%s", output)
		return
	}

	for i, exp := range expected {
		if lines[i] != exp {
			t.Errorf("行 %d:\n期待: %s\n実際: %s", i+1, exp, lines[i])
		}
	}
}
