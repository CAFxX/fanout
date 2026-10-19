package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }

func main() {
    var v1 float32 = (-2.25 * 1.5)
    var v2 int16 = 32767
    var v3 uint64 = 4294967295
    var v4 uint16 = 2
    var v5 int32 = -2147483648
    v6 := GenS1{F0: -128, F1: v3}
    v7 := []int{0, 42, 2147483647}
    var v8 int = 599
    v9 := GenS1{F0: 100, F1: v3}
    if (!true) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v10 int8 = 0
        v5 = v5
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint32 = 2
        v5 = (v5 * (v5 << ((65535) & 63)))
        _ = fi11
        _ = v12
    }
    v7 = append(v7, v8)
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float32 = v1
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 0, F1: v5, F2: ((("\n\t" + "a") + "") + ("" + ""))}
    fmt.Println(v5, v5, v5)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 = (^(v8))
    for fi17 := 0; fi17 < 3; fi17++ {
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "foo bar":
            var v18 uint8 = 255
            _ = v18
        default:
            v13 += v13
        }
        _ = len(v7)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
