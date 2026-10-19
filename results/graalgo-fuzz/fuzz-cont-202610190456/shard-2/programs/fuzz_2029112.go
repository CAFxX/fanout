package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 int32; F2 uint64 }
func genFn0(p0 uint32, p1 int32) bool {
    p0 -= (p0 ^ (^(p0)))
    v5 := GenS1{F0: -111, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v5
    return (false && false)
}
func genFn1(p0 uint64, p1 int32) int32 {
    p0 = 42
    v6 := GenS1{F0: 1, F1: 0}
    if (1 <= 727) {
        v7 := GenS1{F0: -128, F1: 0}
        _ = v7
    }
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return (-(p1))
}

func main() {
    v1 := []uint32{1, 2147483647}
    v2 := []int{}
    var v3 int16 = -607
    var v4 bool = false
    v3 += v3
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v4 = v4
    if (!v4) {
        v3 = v3
    } else {
        v9 := GenS1{F0: 42, F1: -9223372036854775808}
        _ = v9
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v4, v3)
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int16 = (^(v3))
        v12 := GenS1{F0: 127, F1: 0}
        _ = len(v1)
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 = ((2 < 0) && (!(v3 >= v3)))
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 = v3
        v14 := GenS2{F0: "a", F1: genFn1(42, 758), F2: 1}
        v4 = ((((0.0 * float64(-2.25)) + (1.5 - ((1.5 - (-2.25 + 3.14159)) - 1.5))) / -2.25) != 0.0)
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 42
    if (-571 < -922) {
        v16 := GenS1{F0: 42, F1: 427}
        v4 = (!v4)
        for fi17 := 0; fi17 < 17; fi17++ {
            fmt.Println(v4)
            var v18 float32 = 0.0
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        if false {
            var v19 float32 = 3.14159
            v20 := GenS2{F0: "hello", F1: -2147483648, F2: 4294967295}
            _ = v19
            _ = v20
        } else {
            _ = len(v1)
            fmt.Println(v3)
        }
        fmt.Println(v3, v3)
    }
    v15 += v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
