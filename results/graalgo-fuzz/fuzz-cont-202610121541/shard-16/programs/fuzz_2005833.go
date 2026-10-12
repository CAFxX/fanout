package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 bool, p2 int) uint {
    fmt.Println(p1, p2)
    var v5 string = ("a" + "a")
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: (!(42 >= 42)), F1: 362, F2: 255}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 791
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 337
    var v3 uint64 = 1
    v4 := make(map[uint32]uint32)
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(fi8)
        if (-608 == -2147483648) {
            v9 := GenS2{F0: 18446744073709551615, F1: true, F2: "hello"}
            var v10 int8 = 0
            switch ("a" + "hello") {
            case "hello":
                v11 := []int{(fi8 + fi8), -1, (fi8 ^ 247), fi8}
                _ = v11
            }
            _ = v10
            _ = v9
        } else {
            if (true || true) {
                v1 = float64(-2.25)
                var v12 bool = (-128 > 99)
                v13 := GenS3{F0: -1, F1: (^(v3))}
                _ = v12
                _ = v13
            } else {
                var v14 int64 = 9223372036854775807
                if (!false) {
                    v3 = 4294967295
                    v15 := []int{1}
                    _ = v15
                }
                _ = v14
            }
        }
        _ = fi8
    }
    switch "a" {
    case "hello":
        var v16 uint8 = 28
        _ = v16
    case "\n\t":
        v17 := v4[4294967295]
        _ = v17
    default:
        v18 := []int{536}
        _ = v18
    }
    if (65 < 127) {
        v19 := GenS1{F0: true, F1: 570, F2: 240}
        v20 := GenS1{F0: false, F1: 97, F2: 2}
        _ = v19
        _ = v20
    } else {
        var v21 int = 151
        v22 := []int{(v21 | 786), (v21 + v21), v21}
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: genFn0(32767, true, 330), F1: true, F2: ""}
    v24 := GenS2{F0: v2, F1: true, F2: ""}
    var v25 string = "a"
    v4[65535] = 2
    v26 := GenS4{F0: 17}
    v3 -= v3
    v27 := []int{}
    var v28 uint = 4294967295
    fmt.Println(v1, v3)
    fmt.Println(v25, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
