package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 int64, p2 int) uint8 {
    var v4 int16 = 0
    var v5 uint32 = 2147483647
    v5 -= 119
    if (!false) {
        p1 = -1
        p1 += p1
        switch ((("a" + "foo bar") + "\n\t") + "hello") {
        case "\n\t":
            v5 = (v5 / ((v5) | 1))
        case "hello":
            p0 -= (^(p0))
        case "foo bar":
            var v6 float64 = (3.14159 / 3.14159)
            _ = v6
        default:
            var v7 uint64 = 4294967295
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{65535}
    var v3 string = ("" + ("hello" + "a"))
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 876
        var v10 int8 = 100
        v10 += 1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint16 = 65535
    for fi12 := 0; fi12 < 8; fi12++ {
        _ = len(v2)
        v13 := GenS1{F0: false, F1: v3, F2: "a"}
        var v14 uint64 = 429
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v3, v11)
    fmt.Println(v11, v3, v1)
    v15 := GenS1{F0: true, F1: v3, F2: (v3 + ("\n\t" + ((((("foo bar" + "\n\t") + "\n\t") + "hello") + v3) + v3)))}
    if (274 < 42) {
        var v16 string = v3
        fmt.Println(v16)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch fi17 {
            case -301:
                v18 := GenS1{F0: ((v1 ^ (-(v1))) >= v1), F1: "foo bar", F2: v3}
                _ = v18
            case 1:
                var v19 float32 = (((float32(float64((1e10 - 1e10))) * 1.5) - 0.1) / 1.0)
                _ = v19
            default:
                v1 += v1
            }
            var v20 string = (v3 + "foo bar")
            _ = fi17
            _ = v20
        }
    }
    v11 = v11
    v3 = ("a" + v3)
    v21 := GenS2{F0: 1}
    _ = len(v2)
    v11 -= v11
    v11 = (v11 * (v11 | (v11 & v11)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
