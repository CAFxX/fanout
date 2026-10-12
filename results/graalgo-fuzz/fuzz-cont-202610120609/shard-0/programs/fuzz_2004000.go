package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int8; F2 bool }
func genFn0() int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: -12, F1: (false || true)}
        v7 := GenS2{F0: -1}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (false && (215 != 2147483647))
    if (42 > 1) {
        var v9 bool = v8
        _ = v9
    }
    v10 := []int{-1, 42, 191}
    _ = v10
    _ = v8
    return 100
}

func main() {
    var v1 int16 = 1
    var v2 int64 = -9223372036854775808
    var v3 float64 = (0.1 * -2.25)
    var v4 uint = 826
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{42, (^(fi11)), fi11}
        _ = fi11
        _ = v12
    }
    switch ("" + (("\n\t" + "foo bar") + "hello")) {
    case "foo bar":
        v13 := GenS1{F0: 127, F1: true}
        _ = v13
    case "\n\t":
        if ((920 == 2147483647) || (true || (!false))) {
            var v14 int64 = v2
            var v15 uint16 = 785
            _ = v14
            _ = v15
        } else {
            v1 += (v1 >> ((1) & 63))
        }
    case "":
        v16 := GenS3{F0: 2147483647, F1: -771}
        _ = v16
    default:
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint8 = 30
            v2 += 9223372036854775807
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int = (^(fi17))
                v20 = v20
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    }
    var v21 int16 = (v1 >> (((v4 / ((1) | 1))) & 63))
    v22 := GenS4{F0: 188, F1: 141, F2: 42}
    switch 42 {
    case -2147483648:
        fmt.Println(v21, v2)
    case 332:
        switch "foo bar" {
        case "a":
            switch "\n\t" {
            case "hello":
                v23 := []int{-834}
                _ = v23
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 float64 = v3
                    _ = fi24
                    _ = v25
                }
            }
        default:
            v4 = v4
        }
    default:
        if true {
            if true {
                v26 := GenS1{F0: 1, F1: false}
                fmt.Println(v3, v1)
                _ = v26
            }
            v27 := GenS5{F0: 0}
            _ = v27
        }
    }
    v28 := GenS6{F0: false, F1: genFn0(), F2: (-1 >= 2147483647)}
    var v29 bool = (945 < 65535)
    var v30 int8 = 79
    var v31 uint = (v4 * 1)
    v31 -= v31
    v32 := []int{127, 2147483647, -391, 42}
    if len(v32) > 0 {
        v32[0] = 4294967295
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v21, v4, v2)
    case "a":
        v30 -= (-(v30))
    case "hello":
        var v33 float64 = ((((v3 + v3) * 0.0) + v3) + ((v3 + (v3 + 1e10)) / v3))
        _ = v33
    default:
        var v34 uint32 = 186
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
