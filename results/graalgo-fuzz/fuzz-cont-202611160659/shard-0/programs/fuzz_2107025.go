package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 uint, p1 int32) int16 {
    fmt.Println(p0)
    var v7 int32 = (p1 / (((p1 - (p1 | (^(p1))))) | 1))
    fmt.Println(p1)
    p1 = v7
    switch -1 {
    case -168:
        var v8 int16 = 0
        _ = v8
    case 516:
        var v9 int16 = -32768
        _ = v9
    case 823:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint16 = 2
            var v12 uint16 = v11
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 64
}

func main() {
    var v1 float32 = ((1.5 - float32(float64(0.1))) + ((0.1 - 0.0) * -2.25))
    v2 := []string{("a" + ("" + "foo bar")), ("" + "\n\t")}
    var v3 bool = false
    v4 := []string{"\n\t"}
    var v5 string = "\n\t"
    v6 := make(map[string]string)
    var v13 string = v5
    v1 = v1
    var v14 uint = 255
    v14 = (^(v14))
    var v15 int8 = 100
    var v16 float64 = 1e10
    v14 = (v14 % ((760) | 1))
    v17 := GenS1{F0: 100, F1: v13}
    v15 -= v15
    _ = len(v6)
    if (42 > 1) {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    var v18 int32 = -2147483648
    var v19 int8 = v15
    var v20 uint64 = 42
    var v21 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
