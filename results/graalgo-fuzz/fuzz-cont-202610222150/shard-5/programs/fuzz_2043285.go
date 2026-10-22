package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]bool)
    var v3 string = ("a" + "a")
    var v4 string = ""
    v5 := []string{v4}
    var v6 uint8 = 86
    v4 = v4
    v5 = append(v5, ("hello" + "foo bar"))
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int8 = -1
        v9 := GenS1{F0: -32768, F1: "", F2: int64(v6)}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float32 = 0.1
    if false {
        var v11 int32 = 42
        v10 = v10
        _ = v11
    } else {
        var v12 uint8 = (v6 / ((128) | 1))
        v13 := GenS1{F0: -32768, F1: v3, F2: 0}
        _ = v12
        _ = v13
    }
    var v14 int8 = 62
    var v15 string = ("foo bar" + (((v3 + "hello") + v4) + v3))
    if true {
        v16 := GenS1{F0: -32768, F1: v3, F2: -217}
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
