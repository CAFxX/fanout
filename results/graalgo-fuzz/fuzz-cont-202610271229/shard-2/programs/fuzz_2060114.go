package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int }
func genFn0() uint64 {
    v6 := GenS1{F0: 119, F1: -75, F2: 128}
    v7 := []int{4294967295, 1, -2147483648}
    switch "\n\t" {
    case "hello":
        v8 := GenS2{F0: 1, F1: -834}
        _ = v8
    case "foo bar":
        if (false && true) {
            var v9 uint16 = 65535
            _ = v9
        } else {
            var v10 uint = 93
            _ = v10
        }
    case "\n\t":
        var v11 uint = 1
        _ = v11
    default:
        var v12 uint64 = 596
        _ = v12
    }
    _ = len(v7)
    _ = len(v7)
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint32 {
    v13 := GenS3{F0: 1}
    if (!false) {
        var v14 string = "hello"
        _ = v14
    }
    var v15 int = 128
    v15 += (v15 & (v15 & 42))
    _ = v13
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "foo bar"
    v3 := make(map[int]bool)
    var v4 float32 = -2.25
    var v5 uint64 = 0
    var v16 bool = false
    if v16 {
        v16 = false
        var v17 int32 = -2147483648
        _ = v17
    }
    var v18 uint = 255
    fmt.Println(v5)
    v19 := v1[v2]
    v5 -= (v5 ^ v5)
    var v20 uint64 = 0
    fmt.Println(v16, v2)
    if true {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := []int{fi21, fi21, fi21, 42}
            if (187 <= 2) {
                for fi23 := 0; fi23 < 17; fi23++ {
                    if v16 {
                        v20 += 873
                    } else {
                        if ("a" <= (v2 + (("\n\t" + "") + v2))) {
                            v5 = (v20 ^ v20)
                            var v24 int64 = 1
                            v25 := GenS4{F0: 1}
                            _ = v24
                            _ = v25
                        }
                        _ = len(v3)
                    }
                    v2 = v2
                    _ = fi23
                }
            } else {
                v16 = (42 < (v19 * (v19 / (((v19 % ((v19) | 1))) | 1))))
                v26 := v3[fi21]
                _ = v26
            }
            _ = fi21
            _ = v22
        }
        v27 := GenS2{F0: (v5 << (((v18 + v18)) & 63)), F1: 256}
        _ = v27
    }
    v28 := GenS5{F0: -82, F1: 2147483647}
    switch 32767 {
    case -2147483648:
        v29 := []int{}
        _ = v29
    default:
        var v30 int64 = 423
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
