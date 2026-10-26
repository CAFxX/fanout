package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int }

func main() {
    var v1 uint = 1
    var v2 string = ("\n\t" + "")
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    var v5 int64 = 719
    v6 := GenS1{F0: int16(v5)}
    v7 := GenS2{F0: 46, F1: v5}
    switch 983 {
    case 2147483647:
        for fi8 := 0; fi8 < 16; fi8++ {
            v4[0] = 65535
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: 475}
                var v11 uint = v1
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(v3, fi8)
                fmt.Println(v1, fi12, fi12)
                v3 = "hello"
                _ = fi12
            }
            _ = fi8
        }
    case 127:
        v13 := GenS3{F0: 1, F1: "", F2: 2}
        _ = v13
    default:
        v14 := GenS2{F0: -128, F1: v5}
        _ = v14
    }
    if (float32(1e10) != 0.0) {
        v2 = v3
    }
    var v15 int32 = 731
    v1 = 42
    if true {
        var v16 int8 = 20
        v2 = (v2 + ((v2 + ("hello" + v3)) + "foo bar"))
        v17 := GenS4{F0: 18446744073709551615, F1: 1, F2: -9223372036854775808}
        _ = v16
        _ = v17
    } else {
        v3 = ("\n\t" + "a")
        var v18 float32 = 0.0
        var v19 uint64 = 4294967295
        _ = v18
        _ = v19
    }
    var v20 int64 = v5
    var v21 int64 = v5
    var v22 uint = 287
    fmt.Println(v15, v2)
    fmt.Println(v22, v21)
    v15 -= (v15 ^ int32(v5))
    if ((976 == 42) || true) {
        var v23 uint8 = 0
        _ = v23
    }
    v24 := GenS2{F0: -1, F1: (v21 >> (((v1 % ((v1) | 1))) & 63))}
    if true {
        var v25 int64 = v20
        v15 -= v15
        v26 := []int{-196, 26}
        _ = v25
        _ = v26
    } else {
        v15 = (v15 % (((v15 | (v15 >> ((161) & 63)))) | 1))
    }
    v2 = (v2 + v3)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
