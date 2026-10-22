package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0() int8 {
    v4 := GenS1{F0: -1}
    if (false || true) {
        v5 := []int{2147483647}
        v6 := GenS2{F0: -643, F1: 4294967295}
        v7 := GenS3{F0: 42, F1: -1}
        _ = v5
        _ = v6
        _ = v7
    } else {
        if (221 < 306) {
            if (false && false) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    fmt.Println(fi8)
                    var v9 string = ("hello" + "\n\t")
                    _ = fi8
                    _ = v9
                }
                v10 := GenS3{F0: 42, F1: -9223372036854775808}
                _ = v10
            }
        } else {
            v11 := []int{}
            _ = v11
        }
    }
    v12 := []int{}
    var v13 uint16 = 65535
    if (255 <= 0) {
        fmt.Println(v13)
        for fi14 := 0; fi14 < 16; fi14++ {
            if len(v12) > 0 {
                v12[0] = (fi14 / ((fi14) | 1))
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                v12 = append(v12, (fi14 | (-(fi14))))
                _ = fi15
            }
            _ = fi14
        }
    } else {
        var v16 string = "foo bar"
        if true {
            if false {
                var v17 uint = 255
                _ = v17
            }
        }
        _ = v16
    }
    _ = v12
    _ = v13
    _ = v4
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    v3 -= (0.0 + 1e10)
    fmt.Println(v3)
    fmt.Println(v3)
    var v18 int64 = 1
    v18 += (v18 / (((v18 + (^(v18)))) | 1))
    v19 := GenS1{F0: -675}
    v18 = v18
    var v20 float32 = (v3 / v3)
    if true {
        v21 := []int{128}
        _ = v21
    } else {
        if (!true) {
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int64 = v18
                _ = fi22
                _ = v23
            }
        } else {
            var v24 int = -406
            _ = v24
        }
        switch -2147483648 {
        case -2147483648:
            v25 := GenS2{F0: 42, F1: 4294967295}
            _ = v25
        case 9223372036854775807:
            var v26 uint16 = 65535
            _ = v26
        case -9223372036854775808:
            var v27 int32 = 0
            _ = v27
        }
        var v28 float32 = (v20 * float32((0.1 - (1e10 - 0.1))))
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
