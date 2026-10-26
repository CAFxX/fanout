package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0() string {
    switch "foo bar" {
    case "":
        var v4 uint16 = 1
        _ = v4
    }
    v5 := []int{0, 832, -9223372036854775808, 32768}
    v6 := GenS1{F0: 42}
    var v7 float32 = 0.1
    v7 -= float32(float64(v7))
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint16, p1 uint, p2 int8) int32 {
    var v8 int32 = 1
    v9 := GenS1{F0: 294}
    v10 := GenS2{F0: -9223372036854775808, F1: 0, F2: ("hello" + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 int64 = 678
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    v2 -= v2
    var v11 int16 = 0
    v12 := []int{}
    delete(v3, 759)
    var v13 int8 = -126
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    fmt.Println(v13, v1)
    for fi15 := 0; fi15 < 5; fi15++ {
        v13 = v13
        if (2 < 114) {
            v1 = (v1 >> ((42) & 63))
            for fi16 := 0; fi16 < 3; fi16++ {
                v11 += (^(v11))
                var v17 uint64 = 1
                _ = fi16
                _ = v17
            }
        } else {
            v18 := GenS1{F0: v11}
            _ = v18
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        switch 1 {
        case -2147483648:
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 float64 = -2.25
                _ = fi20
                _ = v21
            }
        case 9223372036854775807:
            var v22 uint32 = 2147483647
            _ = v22
        default:
            var v23 int64 = (v1 + 9223372036854775807)
            _ = v23
        }
        if false {
            v2 = (v2 | (^(v2)))
            v3[42] = fi19
        } else {
            v24 := GenS3{F0: 4294967295, F1: 847}
            _ = v24
        }
        v1 += v1
        _ = fi19
    }
    v25 := GenS3{F0: 18446744073709551615, F1: 4294967295}
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
}
