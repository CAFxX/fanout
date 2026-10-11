package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
func genFn0(p0 int, p1 uint64, p2 uint16) string {
    if true {
        v7 := []int{}
        var v8 string = "foo bar"
        v9 := GenS1{F0: true, F1: ((!false) && ((9223372036854775807 >= 42) || (1 <= 32767))), F2: (p2 / ((p2) | 1))}
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p1, p1, p1)
    p0 = (p0 * p0)
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1() bool {
    var v10 int16 = 32767
    if (!false) {
        fmt.Println(v10)
        switch "foo bar" {
        case "":
            var v11 uint = 187
            _ = v11
        }
    } else {
        var v12 int64 = -800
        _ = v12
    }
    _ = v10
    return (v10 >= v10)
}

func main() {
    v1 := []int{790, 128, 0}
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 string = ("\n\t" + "\n\t")
    var v5 uint64 = 18446744073709551615
    var v6 string = v4
    var v13 float64 = 1e10
    switch v6 {
    case "":
        v13 += float64(-2.25)
    default:
        var v14 float64 = v13
        _ = v14
    }
    fmt.Println(v13)
    fmt.Println(v5)
    if true {
        var v15 int = -9223372036854775808
        v15 = v15
        v15 -= 42
        _ = v15
    } else {
        var v16 bool = false
        _ = v16
    }
    v6 = v6
    for fi17 := 0; fi17 < 19; fi17++ {
        switch fi17 {
        case -585:
            if (!false) {
                v18 := GenS2{F0: 4294967295}
                _ = v18
            }
        case -1:
            var v19 uint16 = 65535
            _ = v19
        case 4294967295:
            var v20 uint64 = (v5 * v5)
            _ = v20
        default:
            var v21 uint = 186
            _ = v21
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v4 = v4
        v23 := GenS3{F0: (!true), F1: 900, F2: ((^(v5)) == 0)}
        v24 := v3[4294967295]
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
