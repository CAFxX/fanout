package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 uint) uint8 {
    v5 := GenS1{F0: 65535}
    if true {
        var v6 float64 = 0.1
        _ = v6
    } else {
        var v7 float32 = (-2.25 * (0.0 + 1e10))
        _ = v7
    }
    var v8 string = "foo bar"
    if (3.14159 <= 0.0) {
        var v9 float32 = float32(1.5)
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v8
    return 50
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{("a" + ""), ("" + "")}
    var v3 uint64 = 4294967295
    var v4 int32 = 1
    if false {
        v4 += (v4 & (-(v4)))
    }
    fmt.Println(v3, v3, v3)
    v10 := GenS1{F0: 2147483647}
    for fi11 := 0; fi11 < 15; fi11++ {
        if len(v2) > 0 {
            v2[0] = (("hello" + (("hello" + "foo bar") + "a")) + "a")
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS2{F0: 0, F1: true, F2: (v1 * (v1 / ((255) | 1)))}
            if (100 >= 115) {
                var v14 int64 = -9223372036854775808
                switch "a" {
                case "a":
                    var v15 uint64 = v3
                    _ = v15
                }
                _ = v14
            }
            v2 = append(v2, "hello")
            _ = fi12
            _ = v13
        }
        v4 -= (-(v4))
        _ = fi11
    }
    var v16 uint8 = 2
    v16 = v16
    for fi17 := 0; fi17 < 1; fi17++ {
        v2 = append(v2, "foo bar")
        fmt.Println(v3)
        v18 := GenS2{F0: 1, F1: (!false), F2: 502}
        _ = fi17
        _ = v18
    }
    var v19 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
