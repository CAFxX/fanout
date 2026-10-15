package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint64) uint64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float64 = (-2.25 + 3.14159)
        _ = fi5
        _ = v6
    }
    p0 = p0
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float64 = (3.14159 * 0.1)
        v8 -= ((v8 / 1.5) + -2.25)
        _ = fi7
        _ = v8
    }
    _ = p0
    return (p0 - 4294967295)
}
func genFn1(p0 int32) uint16 {
    var v9 float64 = (0.0 * (1.5 * 3.14159))
    if true {
        var v10 uint32 = 1
        v11 := []int{128, 154, 128, 586}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("foo bar" + "foo bar")
    v3 := make(map[int]bool)
    v4 := []string{}
    v2 = (("\n\t" + ("\n\t" + "")) + (v2 + v2))
    v12 := v1[159]
    var v13 int = -1
    var v14 float64 = 3.14159
    v12 = ((v2 + "foo bar") + "foo bar")
    v1[v13] = (v2 + ("hello" + "foo bar"))
    v15 := GenS1{F0: 100, F1: 42}
    fmt.Println(v13, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
