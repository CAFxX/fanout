package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
func genFn0(p0 int32) uint {
    v5 := []int{128}
    switch -568 {
    case 128:
        p0 -= p0
    case -9223372036854775808:
        var v6 uint = 511
        _ = v6
    }
    _ = p0
    _ = v5
    return 418
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := make(map[uint32]uint32)
    switch 2147483647 {
    case -1:
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 float64 = 3.14159
            var v9 int64 = -438
            _ = fi7
            _ = v8
            _ = v9
        }
    case 647:
        v4[1] = 4294967295
    case 994:
        v10 := []int{32767}
        _ = v10
    default:
        if (!(!(false || (622 >= 1)))) {
            var v11 uint8 = 37
            v12 := []int{-2147483648}
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: true, F1: 16, F2: true}
            v3 += v3
            var v14 float64 = -2.25
            _ = v13
            _ = v14
        }
    }
    v3 -= 0
    if ((!(!("\n\t" == ""))) && false) {
        switch "hello" {
        case "":
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 bool = true
                    var v18 float32 = (1e10 / 1.0)
                    v17 = ((v18 + float32(float64(((((v18 - 0.0) + 1.5) * 0.1) * (1e10 / v18))))) > v18)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        }
    } else {
        var v19 int8 = 100
        v19 = 112
        _ = v19
    }
    fmt.Println(v3)
    var v20 bool = false
    v20 = (1 != 1)
    v3 = v3
    v21 := GenS2{F0: 1, F1: 32768}
    var v22 float32 = 3.14159
    if (32767 <= 358) {
        v23 := GenS3{F0: 227, F1: "", F2: v3}
        var v24 uint = 23
        var v25 int8 = 42
        _ = v23
        _ = v24
        _ = v25
    }
    switch -941 {
    case 127:
        var v26 bool = v20
        _ = v26
    case -682:
        v27 := []int{-2147483648, 32767, -619, -587}
        _ = v27
    }
    v28 := GenS3{F0: 175, F1: "foo bar", F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
}
