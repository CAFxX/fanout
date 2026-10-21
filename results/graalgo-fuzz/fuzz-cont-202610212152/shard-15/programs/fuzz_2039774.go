package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
func genFn0(p0 int64) int16 {
    if (true && true) {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            var v5 int = -1
            _ = v5
        default:
            v6 := []int{32768, 2147483647, -9223372036854775808}
            _ = v6
        }
        if (("\n\t" == "\n\t") || false) {
            fmt.Println(p0)
            var v7 string = (((("a" + "foo bar") + "hello") + "hello") + "foo bar")
            v8 := []int{}
            _ = v7
            _ = v8
        } else {
            fmt.Println(p0)
        }
    }
    p0 = p0
    v9 := []int{473, -993, 2147483647, 128}
    _ = p0
    _ = v9
    return 1
}
func genFn1() int16 {
    var v10 int16 = -634
    v11 := GenS1{F0: 12, F1: (1.5 == (3.14159 - (0.0 * 1.5)))}
    if false {
        for fi12 := 0; fi12 < 10; fi12++ {
            if false {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v10 += (v10 | genFn0(65))
                    v14 := []int{fi12}
                    v10 = v10
                    _ = fi13
                    _ = v14
                }
            } else {
                switch "" {
                case "\n\t":
                    var v15 int = 1
                    _ = v15
                case "foo bar":
                    var v16 float64 = (-2.25 * 3.14159)
                    _ = v16
                }
                v17 := []int{-1, fi12}
                _ = v17
            }
            if (!true) {
                var v18 int32 = 1
                v19 := GenS2{F0: 1, F1: 0, F2: 1}
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := []int{(-(fi12)), (-(fi12))}
                    _ = fi20
                    _ = v21
                }
                var v22 uint32 = 2147483648
                v23 := GenS2{F0: 1, F1: -1, F2: 255}
                _ = v22
                _ = v23
            }
            _ = fi12
        }
    }
    var v24 int32 = 2147483647
    v25 := []int{-106, 42, -1}
    _ = v10
    _ = v11
    _ = v24
    _ = v25
    return (v10 << ((4294967295) & 63))
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 1
    v3 := []uint32{0, 571, 661, 2147483648}
    var v4 string = "\n\t"
    var v26 int8 = 30
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int32 = 436
        _ = fi27
        _ = v28
    }
    fmt.Println(v26, v26, v1)
    fmt.Println(v26, v26, v1)
    switch 9223372036854775807 {
    case -1:
        v26 = (-(v26))
    case -9223372036854775808:
        fmt.Println(v4, v4, v1)
    case -490:
        v29 := GenS1{F0: 385, F1: true}
        _ = v29
    default:
        switch v4 {
        case "foo bar":
            v3 = append(v3, 65535)
        case "":
            for fi30 := 0; fi30 < 10; fi30++ {
                if (1.5 >= 0.0) {
                    v2 -= (^(v2))
                    v3 = append(v3, 2147483647)
                    fmt.Println(v1, v4, v4)
                }
                var v31 bool = (1 != uint64(v2))
                _ = fi30
                _ = v31
            }
        case "\n\t":
            v32 := GenS3{F0: 1, F1: 2147483647, F2: false}
            _ = v32
        default:
            _ = len(v3)
        }
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(fi33)
            if true {
                v1 += v1
            }
            var v35 int64 = int64(v26)
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    _ = len(v3)
    var v36 int8 = -128
    v2 += 255
    v2 += 517
    fmt.Println(v1, v4, v1)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v4
}
