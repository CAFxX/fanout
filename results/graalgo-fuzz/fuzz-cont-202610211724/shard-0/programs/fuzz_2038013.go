package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 string) uint8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = 2147483647
        var v6 float32 = float32(0.1)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -32768, F1: 208}
    v8 := []int{127, 0, -1, 332}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 125
}
func genFn1() int16 {
    if (950 == -32768) {
        var v9 uint = 42
        _ = v9
    }
    var v10 int64 = 0
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 bool = (fi11 <= fi11)
        _ = fi11
        _ = v12
    }
    _ = v10
    return -32768
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    v3 := make(map[int]uint32)
    var v4 uint = 255
    v4 = (v4 - (v4 - v4))
    var v13 float64 = 1.5
    var v14 uint = 42
    if (false || ((!true) || false)) {
        v15 := GenS2{F0: 1, F1: 2, F2: ((false && false) || (-128 < -128))}
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v4 -= v4
        _ = fi16
    }
    var v17 uint8 = 0
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := []int{531, (fi18 - (fi18 << ((v14) & 63)))}
        v14 = (v4 >> ((v4) & 63))
        _ = fi18
        _ = v19
    }
    var v20 uint64 = (v2 ^ v2)
    if (353 > 1) {
        if false {
            var v21 int16 = 42
            switch "a" {
            case "":
                v1["hello"] = 65535
            case "foo bar":
                var v22 float32 = 1.5
                _ = v22
            }
            var v23 uint64 = (v20 % (((v20 - (v2 << (((v14 - v14)) & 63)))) | 1))
            _ = v21
            _ = v23
        }
        v24 := []int{484, 0}
        v4 -= v4
        _ = v24
    } else {
        v25 := []int{}
        for fi26 := 0; fi26 < 19; fi26++ {
            switch "hello" {
            case "foo bar":
                v27 := GenS2{F0: -9223372036854775808, F1: (v17 % (((v17 & (v17 * uint8(v2)))) | 1)), F2: (true && false)}
                _ = v27
            case "hello":
                v17 += 0
            default:
                fmt.Println(fi26, fi26, fi26)
            }
            _ = len(v3)
            _ = fi26
        }
        var v28 int16 = 1
        _ = v25
        _ = v28
    }
    switch -975 {
    case -2147483648:
        if (!false) {
            for fi29 := 0; fi29 < 10; fi29++ {
                v30 := GenS3{F0: v14}
                switch "a" {
                case "hello":
                    var v31 int64 = 734
                    _ = v31
                default:
                    v32 := GenS3{F0: 18446744073709551615}
                    _ = v32
                }
                v20 = v2
                _ = fi29
                _ = v30
            }
            fmt.Println(v20)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
