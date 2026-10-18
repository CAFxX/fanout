package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 string) bool {
    v4 := []int{9223372036854775807}
    _ = len(v4)
    v5 := GenS1{F0: 970, F1: 1, F2: 665}
    v6 := GenS1{F0: 0, F1: 32768, F2: -1}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (false && true)
}
func genFn1(p0 string, p1 uint32, p2 uint8) int {
    var v7 float32 = float32(-2.25)
    if true {
        switch 1 {
        case 9223372036854775807:
            v8 := GenS1{F0: 1, F1: -1, F2: -2147483648}
            _ = v8
        case -117:
            if false {
                v9 := GenS1{F0: -860, F1: -1, F2: 1}
                _ = v9
            } else {
                if (65535 != p1) {
                    p0 = "hello"
                    p2 -= p2
                }
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{}
                    v12 := GenS1{F0: 590, F1: fi10, F2: -102}
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 int16 = 0
                        _ = fi13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        case 32768:
            var v15 uint64 = 749
            _ = v15
        default:
            var v16 uint32 = p1
            _ = v16
        }
        if true {
            p1 = uint32(p2)
            v17 := []int{-1}
            _ = v17
        }
        if (-967 > 582) {
            p0 = p0
            switch -113 {
            case -1:
                fmt.Println(p0, p2)
            }
            fmt.Println(v7, p2)
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 932
}

func main() {
    v1 := []int{4294967295, -20, 42}
    var v2 int8 = 42
    v3 := make(map[string]int)
    var v18 uint8 = 2
    v2 += v2
    var v19 int8 = (v2 ^ 0)
    v20 := GenS1{F0: -1, F1: 128, F2: 971}
    var v21 int64 = 1
    v22 := GenS1{F0: -1, F1: 246, F2: 0}
    if false {
        for fi23 := 0; fi23 < 1; fi23++ {
            if (fi23 < genFn1(("hello" + "foo bar"), 0, v18)) {
                v2 = (v19 & (v19 / ((42) | 1)))
                var v24 uint64 = 18446744073709551615
                fmt.Println(fi23, v24, v19)
                _ = v24
            }
            _ = fi23
        }
    } else {
        v19 = v19
        switch "a" {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            fmt.Println(v18)
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v19 -= (v19 & (v19 >> ((2) & 63)))
    var v25 int8 = v19
    var v26 int = -9223372036854775808
    v27 := GenS2{F0: -1, F1: -1, F2: v19}
    if (42 == 773) {
        v25 -= (-(v19))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
