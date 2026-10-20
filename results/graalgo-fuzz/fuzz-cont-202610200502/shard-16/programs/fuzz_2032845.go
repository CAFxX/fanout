package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{(("" + (((("a" + "") + "hello") + "\n\t") + ("\n\t" + "\n\t"))) + "foo bar")}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 int8 = 127
    var v6 uint = 2
    v5 -= v5
    switch 42 {
    case 127:
        var v7 int8 = v5
        _ = v7
    default:
        v8 := GenS1{F0: -776}
        _ = v8
    }
    var v9 int = 1
    var v10 bool = true
    v11 := GenS2{F0: 218, F1: 443}
    switch int(v6) {
    case 32768:
        var v12 uint32 = 584
        _ = v12
    case -9223372036854775808:
        var v13 int = v9
        _ = v13
    default:
        v10 = v10
    }
    delete(v4, -420)
    v14 := GenS2{F0: 918, F1: 362}
    var v15 int32 = 2147483647
    var v16 uint16 = 42
    for fi17 := 0; fi17 < 15; fi17++ {
        if v10 {
            for fi18 := 0; fi18 < 3; fi18++ {
                _ = len(v3)
                _ = fi18
            }
            var v19 int16 = -610
            _ = v19
        }
        v5 = (v5 | -128)
        _ = fi17
    }
    var v20 uint16 = 202
    fmt.Println(v5, v16, v15)
    var v21 int8 = 104
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
