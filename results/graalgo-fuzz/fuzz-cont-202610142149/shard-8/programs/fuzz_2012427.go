package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 int32) int32 {
    switch "hello" {
    case "hello":
        switch -2147483648 {
        case -435:
            v5 := GenS1{F0: 924}
            _ = v5
        case 0:
            p0 = p0
        default:
            var v6 int16 = 0
            _ = v6
        }
    case "\n\t":
        fmt.Println(p0)
    default:
        var v7 int = 2147483647
        _ = v7
    }
    var v8 uint = 798
    var v9 int32 = p0
    _ = p0
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 float64 = ((0.1 / -2.25) - -2.25)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v1 -= v1
    v10 := GenS2{F0: "", F1: 42}
    v11 := GenS1{F0: 143}
    fmt.Println(v1)
    var v12 string = "\n\t"
    switch -2147483648 {
    case -2147483648:
        var v13 uint32 = 2
        _ = v13
    case 128:
        if (false && (-2147483648 < -2147483648)) {
            v14 := GenS2{F0: ((v12 + ("\n\t" + v12)) + v12), F1: 2}
            v1 = (v1 * ((1.5 + (v1 * ((v1 / v1) * v1))) - (v1 + v1)))
            if ((false && false) || true) {
                var v15 uint32 = 2147483647
                fmt.Println(v15, v12, v15)
                _ = v15
            } else {
                var v16 float64 = v1
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(v16, v16)
                    _ = fi17
                }
                v18 := v4[((v12 + "foo bar") + v12)]
                _ = v16
                _ = v18
            }
            _ = v14
        }
    case 42:
        var v19 uint8 = 1
        _ = v19
    default:
        var v20 uint32 = 0
        _ = v20
    }
    v12 = (("foo bar" + "a") + ("a" + v12))
    v12 = ("a" + v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
