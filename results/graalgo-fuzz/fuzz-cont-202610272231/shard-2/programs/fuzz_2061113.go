package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 int16; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 int64 }
func genFn0(p0 string) uint8 {
    var v5 int = 445
    v6 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: v5, F1: 42, F2: 940}
        v9 := GenS1{F0: 127, F1: -1, F2: 273}
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float32 = 0.1
            var v12 uint64 = 18446744073709551615
            var v13 uint16 = 42
            v6 = append(v6, (v5 >> ((699) & 63)))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []string{"hello", ("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))}
    var v2 int8 = 42
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    fmt.Println(v2)
    fmt.Println(v2)
    v14 := GenS1{F0: -9223372036854775808, F1: (-(v2)), F2: 78}
    var v15 bool = (909 < 1)
    if v15 {
        var v16 int = 9223372036854775807
        _ = v16
    } else {
        v17 := GenS1{F0: 2147483647, F1: v2, F2: 42}
        var v18 float64 = ((1.5 + 0.0) + 0.0)
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "a":
        if ((float64(0.0) / 3.14159) < (3.14159 * float64((1.5 + -2.25)))) {
            var v19 int64 = 75
            _ = v19
        } else {
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int8 = v2
                v2 = (^(v21))
                _ = fi20
                _ = v21
            }
            v2 += 127
        }
    case "hello":
        fmt.Println(v15, v2)
    default:
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2, v15)
        case "hello":
            var v22 uint = 2
            _ = v22
        case "a":
            var v23 int16 = 518
            _ = v23
        default:
            switch -892 {
            case 127:
                v15 = v15
            default:
                fmt.Println(v2)
            }
        }
    }
    v3["\n\t"] = 539
    v24 := GenS2{F0: -440, F1: -32768, F2: 170}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    v15 = (("a" + ((("foo bar" + "hello") + (("hello" + ("hello" + "\n\t")) + ((("hello" + "foo bar") + "foo bar") + "\n\t"))) + "\n\t")) != "\n\t")
    v25 := GenS3{F0: -1}
    switch 42 {
    case -449:
        var v26 uint64 = 42
        _ = v26
    }
    v2 -= (v2 ^ v2)
    delete(v3, (("hello" + "a") + ("hello" + "\n\t")))
    if false {
        _ = len(v1)
        v15 = false
    } else {
        switch 378 {
        case 42:
            v15 = (255 <= 685)
        case 462:
            v27 := GenS3{F0: -387}
            _ = v27
        default:
            v28 := GenS4{F0: 42, F1: ((1.5 * (float64((0.0 + (0.0 * 0.1))) + (1e10 * -2.25))) >= 1e10), F2: 42}
            _ = v28
        }
        fmt.Println(v15, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
