package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint32 = 560
    var v2 float64 = (-2.25 / 3.14159)
    var v3 uint64 = 42
    var v4 int16 = -32768
    var v5 uint = 18446744073709551615
    v5 -= v5
    v2 = 0.0
    v2 -= (3.14159 - float64(3.14159))
    var v6 float64 = (v2 * v2)
    var v7 uint32 = (^(v1))
    if (0 < 1) {
        var v8 int = 32767
        _ = v8
    }
    v9 := GenS1{F0: v1}
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v4, v6, v2)
        v7 = 4294967295
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    v5 += v5
    if (true || (false || true)) {
        var v12 bool = ("\n\t" == ("foo bar" + "foo bar"))
        fmt.Println(v12, v7, v4)
        _ = v12
    } else {
        v13 := []int{243}
        _ = v13
    }
    v3 = (^(v3))
    if (1e10 != ((-2.25 - (v2 - (v6 * (v6 / (1.5 / (0.1 * 1e10)))))) - v2)) {
        var v14 string = ("a" + "")
        _ = v14
    }
    v15 := GenS2{F0: 255, F1: false}
    var v16 int16 = (v4 | -1)
    v2 = 1.5
    if ("" != ("a" + "")) {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int = (fi17 % ((fi17) | 1))
            _ = fi17
            _ = v18
        }
    } else {
        var v19 string = "a"
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := []int{-9223372036854775808}
            var v23 bool = true
            var v24 int16 = v4
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        switch "a" {
        case "foo bar":
            v25 := []int{}
            _ = v25
        default:
            var v26 string = "a"
            _ = v26
        }
        _ = fi20
    }
    v27 := []int{32768, 32767, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
