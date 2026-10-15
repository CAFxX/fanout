package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 bool, p1 uint16, p2 uint8) int16 {
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(p1, p1)
        var v8 string = "\n\t"
        _ = fi7
        _ = v8
    }
    v9 := []int{-9223372036854775808, -2147483648, 2147483647}
    for fi10 := 0; fi10 < 12; fi10++ {
        if len(v9) > 0 {
            v9[0] = (-(fi10))
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 598
}

func main() {
    var v1 float32 = 0.1
    var v2 int64 = 1
    var v3 float32 = v1
    var v4 float64 = (1e10 - 0.1)
    var v5 int8 = 127
    v6 := make(map[int]uint32)
    v11 := []int{}
    v3 -= v1
    var v12 uint = 18446744073709551615
    for fi13 := 0; fi13 < 2; fi13++ {
        if (!false) {
            var v14 int = (fi13 | fi13)
            _ = v14
        }
        var v15 int64 = (v2 % ((v2) | 1))
        _ = fi13
        _ = v15
    }
    var v16 uint64 = 129
    for fi17 := 0; fi17 < 2; fi17++ {
        if (!false) {
            fmt.Println(v16, v5)
            fmt.Println(v5)
            v1 = v3
        }
        v5 = (v5 - -1)
        v18 := GenS1{F0: -1, F1: 1}
        _ = fi17
        _ = v18
    }
    switch "\n\t" {
    case "a":
        switch 32767 {
        case 0:
            for fi19 := 0; fi19 < 17; fi19++ {
                if ("a" < "") {
                    fmt.Println(fi19)
                }
                var v20 bool = false
                _ = fi19
                _ = v20
            }
        case 4294967295:
            v21 := GenS2{F0: (v2 - v2), F1: ("" + ("\n\t" + "foo bar"))}
            _ = v21
        default:
            v22 := GenS2{F0: 0, F1: "\n\t"}
            _ = v22
        }
    case "\n\t":
        var v23 int16 = genFn0(false, 5, 44)
        _ = v23
    }
    v3 -= float32(-2.25)
    v12 = 168
    fmt.Println(v12)
    _ = len(v6)
    var v24 uint32 = 647
    v24 = v24
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
