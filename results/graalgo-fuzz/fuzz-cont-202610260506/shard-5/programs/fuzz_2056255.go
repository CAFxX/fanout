package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := []string{"", "a", ("" + "\n\t")}
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    var v5 string = "foo bar"
    var v6 int32 = 42
    var v7 int8 = -128
    v3 = v3
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 = (v6 | (v6 | (v6 | 776)))
        _ = fi8
    }
    v1 += (v1 << ((1) & 63))
    switch -9223372036854775808 {
    case 2147483647:
        v3 = (v3 ^ 223)
    case 32768:
        switch -399 {
        case 42:
            v6 = -601
        }
    case -1:
        for fi9 := 0; fi9 < 1; fi9++ {
            if true {
                fmt.Println(v7)
                if false {
                    v5 = v5
                    v10 := GenS1{F0: 2, F1: -699, F2: v3}
                    v11 := GenS2{F0: (^(v3)), F1: 329}
                    _ = v10
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                v7 -= v7
                fmt.Println(v5, v1)
                _ = fi12
            }
            _ = fi9
        }
    default:
        fmt.Println(v6, v6, v5)
    }
    v3 += v3
    v13 := GenS1{F0: 428, F1: int64(v7), F2: v3}
    var v14 int8 = v7
    if false {
        var v15 int32 = (v6 >> ((65535) & 63))
        delete(v4, 1)
        v1 = (v1 & (v1 * (v1 ^ 4294967295)))
        _ = v15
    }
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 17; fi17++ {
        if (v3 < v3) {
            v4[1] = fi17
            v6 -= 2147483647
        } else {
            v18 := v4[(^(v1))]
            var v19 float32 = -2.25
            _ = v18
            _ = v19
        }
        var v20 string = v5
        _ = fi17
        _ = v20
    }
    if true {
        v14 = v14
    }
    v21 := GenS3{F0: 587, F1: 42}
    fmt.Println(v16)
    var v22 int16 = 1
    v23 := GenS3{F0: 2, F1: 4294967295}
    if true {
        v24 := GenS3{F0: 255, F1: 18446744073709551615}
        v1 += (v1 ^ 42)
        fmt.Println(v1, v1)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
