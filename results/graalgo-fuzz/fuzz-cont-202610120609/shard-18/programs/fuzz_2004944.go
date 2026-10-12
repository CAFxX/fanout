package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }
func genFn0(p0 uint32) int32 {
    var v4 int32 = 0
    p0 -= (p0 / ((395) | 1))
    p0 = (p0 << ((138) & 63))
    _ = p0
    _ = v4
    return (v4 / ((40) | 1))
}
func genFn1(p0 uint16, p1 string) string {
    var v5 float64 = 1e10
    p1 = "foo bar"
    var v6 string = (p1 + (("a" + "foo bar") + p1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("" + ("\n\t" + p1))
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int8 = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 = (v3 & (v3 - (v3 % ((v3) | 1))))
        var v8 uint16 = 255
        _ = fi7
        _ = v8
    }
    if true {
        v9 := []int{2147483647, 2147483647, 9223372036854775807, -9223372036854775808}
        _ = v9
    } else {
        fmt.Println(v3)
        v10 := GenS1{F0: (-394 < 1)}
        _ = v10
    }
    var v11 int32 = -200
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch fi13 {
            case 9223372036854775807:
                if (42 == 0) {
                    var v14 int8 = (-(v3))
                    if false {
                        var v15 int32 = v11
                        _ = v15
                    } else {
                        var v16 int16 = 32767
                        delete(v2, fi12)
                        var v17 uint64 = 2
                        _ = v16
                        _ = v17
                    }
                    switch fi13 {
                    case 32767:
                        var v18 int = (^(fi13))
                        _ = v18
                    }
                    _ = v14
                }
            case 4294967295:
                v11 = (v11 >> ((18446744073709551615) & 63))
            case 215:
                var v19 uint8 = 161
                _ = v19
            default:
                var v20 uint8 = 127
                _ = v20
            }
            var v21 float32 = 1e10
            if false {
                var v22 bool = (2 < 258)
                _ = v22
            }
            _ = fi13
            _ = v21
        }
        if true {
            var v23 int32 = 0
            _ = v23
        }
        var v24 float32 = 3.14159
        _ = fi12
        _ = v24
    }
    var v25 int = -727
    switch "\n\t" {
    case "":
        v25 = (v25 - 0)
    case "hello":
        var v26 int64 = 509
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := GenS2{F0: (!false), F1: 120, F2: -987}
            _ = fi27
            _ = v28
        }
    }
    v29 := GenS1{F0: (v11 >= (-(v11)))}
    v30 := []int{249, 0, (v25 >> ((879) & 63))}
    var v31 int64 = 9223372036854775807
    var v32 uint64 = 969
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
