package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int8, p1 string) uint8 {
    switch p1 {
    case "\n\t":
        var v4 uint8 = 1
        _ = v4
    case "a":
        switch p1 {
        case "foo bar":
            if true {
                v5 := GenS1{F0: 4294967295}
                if (-727 >= 668) {
                    var v6 int64 = 766
                    v7 := GenS1{F0: 32768}
                    _ = v6
                    _ = v7
                }
                _ = v5
            } else {
                p0 = p0
            }
        case "a":
            v8 := []int{42}
            _ = v8
        case "\n\t":
            p0 = 127
        }
    default:
        v9 := GenS2{F0: 18446744073709551615, F1: "a", F2: 18446744073709551615}
        _ = v9
    }
    var v10 int = -2147483648
    v11 := []int{}
    v12 := GenS1{F0: v10}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 28
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 2
    var v3 bool = true
    v13 := []int{519, -2147483648, 127}
    if len(v13) > 0 {
        v13[0] = 1
    }
    var v14 float64 = 0.1
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 float64 = 1e10
        v3 = ((true || (67 >= 0)) || v3)
        fmt.Println(v16, v3, v16)
        _ = fi15
        _ = v16
    }
    v2 = v2
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 uint32 = 42
        v19 := GenS3{F0: 1, F1: (v18 >> ((643) & 63))}
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        switch fi20 {
        case 32767:
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            var v21 uint = 920
            _ = v21
        }
        _ = fi20
    }
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
