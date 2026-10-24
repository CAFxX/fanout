package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string; F1 int; F2 bool }
func genFn0() string {
    v5 := []int{9223372036854775807, 776, -672, -9223372036854775808}
    var v6 int32 = 655
    _ = v5
    _ = v6
    return ("" + "hello")
}

func main() {
    var v1 int8 = 118
    v2 := make(map[int]bool)
    v3 := []uint32{1, 2147483647}
    var v4 float32 = 1.5
    v1 = (v1 >> ((65535) & 63))
    v7 := GenS1{F0: 0, F1: 906, F2: 42}
    var v8 int16 = -1
    switch -9223372036854775808 {
    case 576:
        v8 -= v8
    }
    v9 := GenS2{F0: 1, F1: 1, F2: v8}
    v10 := GenS1{F0: 0, F1: 1, F2: 848}
    v4 += 0.1
    switch 27 {
    case -839:
        v11 := GenS2{F0: 2, F1: 760, F2: v8}
        _ = v11
    case -735:
        var v12 uint = 2
        _ = v12
    default:
        v13 := GenS3{F0: 32768, F1: (v1 & v1)}
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    fmt.Println(v1, v4, v4)
    fmt.Println(v8, v8, v1)
    v2[42] = (65535 == 65535)
    var v14 uint = 0
    if true {
        v15 := GenS1{F0: 42, F1: 32767, F2: 178}
        v14 = v14
        v16 := v2[-9223372036854775808]
        _ = v15
        _ = v16
    } else {
        if (-378 != 2147483647) {
            v4 += v4
            switch 1 {
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 1:
                var v17 float64 = 3.14159
                _ = v17
            case 42:
                var v18 float32 = 1.5
                _ = v18
            default:
                var v19 bool = true
                _ = v19
            }
        } else {
            var v20 int16 = 1000
            v14 -= v14
            _ = v20
        }
    }
    v21 := GenS4{F0: ("foo bar" + ("" + "")), F1: -9223372036854775808, F2: false}
    for fi22 := 0; fi22 < 4; fi22++ {
        v1 = (-(v1))
        _ = fi22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        switch 128 {
        case 0:
            fmt.Println(v1)
        case 2147483647:
            var v24 int8 = v1
            _ = v24
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
