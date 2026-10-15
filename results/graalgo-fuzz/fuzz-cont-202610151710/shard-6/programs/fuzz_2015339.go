package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32) uint32 {
    if (!(!(!(18446744073709551615 >= 148)))) {
        var v5 int64 = -179
        var v6 uint8 = 127
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 0, F1: 42}
                v10 := GenS2{F0: 781, F1: -2, F2: "hello"}
                var v11 uint8 = 2
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            p0 = -1
            var v12 bool = true
            _ = fi7
            _ = v12
        }
    }
    var v13 string = (("" + "foo bar") + ("" + "foo bar"))
    _ = p0
    _ = v13
    return 495
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483647
    var v4 uint64 = (^(v1))
    switch -516 {
    case -907:
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 42, F1: fi14}
            switch 32767 {
            case 42:
                v16 := []int{(-(fi14)), fi14, -2147483648}
                _ = v16
            case 32767:
                var v17 float64 = 0.1
                _ = v17
            case 127:
                var v18 float32 = ((-2.25 + 1e10) / (((-2.25 * 1e10) * 1e10) * (float32(float64(-2.25)) - (((3.14159 + 0.1) - (1.5 + (-2.25 - (3.14159 + (1e10 / (0.1 * 1e10)))))) * float32(-2.25)))))
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(fi14, v1, v1)
                v20 := GenS3{F0: 204}
                var v21 int8 = -1
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi14
            _ = v15
        }
    case -2147483648:
        var v22 int64 = 9
        _ = v22
    default:
        fmt.Println(v4, v1, v1)
    }
    v1 -= (v1 & 2)
    var v23 uint16 = 32767
    v4 = (v1 | (v1 << ((18446744073709551615) & 63)))
    v24 := GenS3{F0: 128}
    v1 = v4
    v1 += 4294967295
    v1 = (^(v1))
    if ("a" != "foo bar") {
        delete(v2, 4294967295)
        v23 += v23
        var v25 uint = 607
        _ = v25
    } else {
        if true {
            var v26 uint8 = 2
            for fi27 := 0; fi27 < 3; fi27++ {
                fmt.Println(fi27, fi27)
                _ = fi27
            }
            _ = v26
        }
        fmt.Println(v3)
        v28 := []int{0, 32767, 1, -2147483648}
        _ = v28
    }
    v29 := GenS1{F0: 42, F1: -1}
    v30 := GenS1{F0: -906, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
