package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int32) int32 {
    var v7 bool = false
    fmt.Println(v7)
    v8 := GenS1{F0: 100, F1: 255}
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1() uint64 {
    v9 := []int{32768}
    var v10 uint16 = 2
    var v11 uint16 = (v10 & (v10 / ((v10) | 1)))
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    var v1 bool = true
    var v2 int32 = 405
    var v3 int8 = -78
    v4 := []int{4294967295}
    var v5 uint64 = 18446744073709551615
    var v6 float64 = (1.5 + (0.0 + 0.0))
    _ = len(v4)
    var v12 int64 = 689
    var v13 int16 = 260
    fmt.Println(v13, v13)
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            v13 -= (-(v13))
            if false {
                fmt.Println(v6, v12, v5)
                var v16 int64 = v12
                _ = v16
            }
            _ = fi15
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int = fi17
            if true {
                var v19 int8 = (-(v3))
                v18 = v18
                _ = len(v4)
                v3 -= (v3 - v19)
                if (!(-701 < -1)) {
                    v20 := GenS2{F0: ("a" + "")}
                    v1 = false
                    _ = v20
                }
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        _ = fi14
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v1, v1, v2)
    default:
        v13 = v13
    }
    var v21 int64 = 0
    v22 := GenS2{F0: "foo bar"}
    v3 -= v3
    var v23 float64 = (1e10 + (0.0 - v6))
    v24 := GenS3{F0: 42, F1: 32768}
    v23 += v6
    v25 := GenS3{F0: 558, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
