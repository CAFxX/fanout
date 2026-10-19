package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0(p0 string) int8 {
    v6 := []int{128, -439}
    if (1 < 9223372036854775807) {
        p0 = "foo bar"
    } else {
        var v7 uint16 = 32767
        fmt.Println(p0, p0)
        v8 := GenS1{F0: 374, F1: 225}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return 111
}
func genFn1(p0 string) int32 {
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(p0, fi9)
        _ = fi9
    }
    var v10 float64 = 1.5
    fmt.Println(p0, p0)
    fmt.Println(v10)
    v11 := GenS1{F0: 4294967295, F1: -2147483648}
    _ = p0
    _ = v10
    _ = v11
    return -670
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int = 407
    var v5 int16 = 0
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: 255, F1: -2147483648}
        var v14 int16 = (v5 - (-(v5)))
        v15 := []int{fi12, (v4 % (((-(v4))) | 1)), 9223372036854775807, 0}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    switch (^(v4)) {
    case -2147483648:
        v4 += v4
    }
    fmt.Println(v1)
    v1 = v1
    if false {
        v16 := GenS1{F0: 917, F1: 0}
        _ = v16
    }
    v1 = (v1 % ((v1) | 1))
    switch v4 {
    case 42:
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(fi17, v1)
            _ = fi17
        }
    case 1:
        fmt.Println(v5)
    default:
        v18 := GenS2{F0: -1, F1: 71, F2: 128}
        _ = v18
    }
    if (uint64(v1) != 765) {
        var v19 int64 = 181
        v4 = v4
        var v20 uint = 2
        _ = v19
        _ = v20
    } else {
        fmt.Println(v4)
        _ = len(v2)
        var v21 uint = 42
        _ = v21
    }
    switch 1 {
    case 243:
        if (!(127 >= 1)) {
            var v22 uint64 = 2
            _ = v22
        }
    case 4294967295:
        var v23 float32 = ((((0.1 / (1.5 * 1.5)) * 0.0) * ((0.1 + (3.14159 * float32((float64((-2.25 + 0.1)) + float64(0.0))))) - -2.25)) * float32(1e10))
        _ = v23
    case -1:
        var v24 uint16 = (v1 ^ (v1 >> ((255) & 63)))
        _ = v24
    }
    v5 += v5
    v25 := GenS2{F0: genFn1("a"), F1: 1, F2: 128}
    switch ("\n\t" + ("hello" + "hello")) {
    case "hello":
        v26 := []int{(v4 >> ((1) & 63)), v4}
        _ = v26
    case "a":
        v1 -= v1
    default:
        var v27 bool = true
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
