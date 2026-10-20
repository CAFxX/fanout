package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0() int16 {
    var v7 uint32 = 1
    switch ("a" + "\n\t") {
    case "":
        fmt.Println(v7)
    }
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 417
    var v2 uint32 = 1
    var v3 float64 = -2.25
    v4 := []string{(("hello" + "hello") + "foo bar"), "a", ("" + ((((("a" + "hello") + ("a" + ("foo bar" + "a"))) + "\n\t") + "") + "hello")), "foo bar"}
    var v5 int = 709
    v6 := make(map[uint32]string)
    var v8 int16 = 1000
    var v9 uint64 = 42
    v10 := GenS1{F0: 127, F1: v8}
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            var v12 float32 = 0.1
            _ = v12
        }
        _ = fi11
    }
    var v13 int32 = v1
    var v14 float64 = float64((1.5 / float32((0.1 / ((((1e10 * v3) - v3) / (v3 + ((0.1 * v3) / 3.14159))) - float64((1e10 / 1.5)))))))
    v4 = append(v4, "")
    if len(v4) > 0 {
        v4[0] = ""
    }
    delete(v6, (v2 * v2))
    v15 := GenS1{F0: v5, F1: v8}
    var v16 float32 = 1.5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2)
    v17 := GenS1{F0: (-(v5)), F1: (v8 & (v8 & genFn0()))}
    if true {
        v4 = append(v4, "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
