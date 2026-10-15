package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 string, p1 int, p2 int16) uint16 {
    p0 = (("a" + p0) + ((p0 + (p0 + p0)) + ("foo bar" + p0)))
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]uint32)
    var v3 int8 = 118
    v4 := make(map[string]uint32)
    v3 -= v3
    v5 := GenS1{F0: -128}
    v3 -= (v3 | (^(v3)))
    switch "a" {
    case "\n\t":
        switch "\n\t" {
        case "a":
            var v6 int64 = -1
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 string = ("a" + "")
                _ = fi7
                _ = v8
            }
        case "":
            var v9 uint32 = 0
            _ = v9
        }
    case "":
        if true {
            v10 := []int{317, -318}
            _ = v10
        } else {
            var v11 uint32 = 42
            _ = v11
        }
    case "foo bar":
        v3 = (^(v3))
    }
    v3 = 0
    v12 := v4["a"]
    for fi13 := 0; fi13 < 14; fi13++ {
        switch 42 {
        case -581:
            v14 := GenS1{F0: -128}
            _ = v14
        case -1:
            v1 = 0.0
        default:
            var v15 bool = false
            _ = v15
        }
        _ = fi13
    }
    if true {
        v4[""] = (v12 & (v12 >> ((18446744073709551615) & 63)))
        var v16 bool = (1 >= 4294967295)
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v1, v3)
        v3 = v3
        fmt.Println(v12)
        _ = fi17
    }
    var v18 uint32 = v12
    v18 -= 379
    v12 += 811
    if (!(!true)) {
        var v19 int = 42
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(v3, v18, v12)
            _ = fi20
        }
        _ = v19
    } else {
        var v21 int64 = 452
        if (true && false) {
            _ = len(v4)
        } else {
            var v22 int32 = -1
            var v23 int16 = 1
            for fi24 := 0; fi24 < 10; fi24++ {
                v18 += 752
                fmt.Println(fi24, v3)
                v25 := []int{}
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        }
        v26 := GenS2{F0: 0, F1: true}
        _ = v21
        _ = v26
    }
    v1 = 0.1
    fmt.Println(v12, v12, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
