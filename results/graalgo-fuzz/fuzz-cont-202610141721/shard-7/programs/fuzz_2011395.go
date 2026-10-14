package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint }
func genFn0(p0 int, p1 uint64, p2 int32) uint {
    v7 := []int{p0, (p0 & (-(p0))), p0, 9223372036854775807}
    switch 9223372036854775807 {
    case -23:
        switch "hello" {
        case "foo bar":
            p1 = p1
        case "hello":
            var v8 float64 = -2.25
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: ((false && (65535 != 32767)) && true), F1: true}
                fmt.Println(fi9)
                var v11 int64 = 9223372036854775807
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    case -1:
        switch -283 {
        case 1:
            p1 = p1
        case -626:
            var v12 bool = (0 <= 46)
            _ = v12
        case 263:
            v13 := GenS2{F0: p0, F1: 30}
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 803, 297}
    var v2 float64 = 3.14159
    v3 := []int{}
    v4 := []string{"hello", "foo bar"}
    var v5 uint8 = 91
    v6 := make(map[uint32]bool)
    var v14 uint8 = (v5 << ((453) & 63))
    if ((21 >= -116) && false) {
        switch ("" + "foo bar") {
        case "\n\t":
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS1{F0: true, F1: true}
                _ = fi15
                _ = v16
            }
        }
    } else {
        var v17 int32 = 2147483647
        v2 -= (-2.25 + 1.5)
        _ = v17
    }
    var v18 uint64 = 405
    var v19 float32 = ((1e10 + 3.14159) * (1.5 * (0.1 * float32(((v2 * 0.0) * v2)))))
    fmt.Println(v5, v19)
    var v20 int64 = -621
    v21 := GenS1{F0: (float32((v2 + ((float64((v19 - (v19 + (-2.25 - v19)))) / (v2 * float64(0.1))) - v2))) != ((v19 * -2.25) + -2.25)), F1: false}
    v20 = v20
    v20 -= (v20 * (^(v20)))
    fmt.Println(v5, v20, v5)
    for fi22 := 0; fi22 < 10; fi22++ {
        v20 = v20
        _ = fi22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v20 = (v20 % ((v20) | 1))
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: int(v14), F1: 0}
            if false {
                v18 -= v18
                var v26 int64 = (v20 / ((v20) | 1))
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    delete(v6, 190)
    v2 += 0.0
    v27 := GenS3{F0: "foo bar", F1: -141, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
