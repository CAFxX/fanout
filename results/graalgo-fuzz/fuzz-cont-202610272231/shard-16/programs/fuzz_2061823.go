package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    var v3 uint8 = 127
    var v4 float64 = float64(-2.25)
    v5 := make(map[string]int)
    v6 := []string{("" + "foo bar"), "hello", "", ("a" + ((("\n\t" + "") + "foo bar") + (("foo bar" + "foo bar") + (("foo bar" + "\n\t") + (("a" + "hello") + "hello")))))}
    var v7 uint16 = 42
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    v8 := v5["foo bar"]
    var v9 bool = true
    v8 = v8
    fmt.Println(v9, v8)
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    if (v2 >= v2) {
        v10 := GenS1{F0: 127, F1: (^(v7)), F2: (v7 >> ((0) & 63))}
        _ = v10
    }
    var v11 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
