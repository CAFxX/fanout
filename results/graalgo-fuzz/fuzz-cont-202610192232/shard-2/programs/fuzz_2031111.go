package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8, p1 uint32, p2 int32) int8 {
    if (!(743 != 992)) {
        v6 := GenS1{F0: 1}
        _ = v6
    } else {
        if (-386 <= -1) {
            var v7 uint16 = 1
            _ = v7
        }
        switch 127 {
        case -868:
            var v8 int8 = -128
            _ = v8
        case 42:
            if (1 < 255) {
                var v9 int = 441
                p1 = (p1 & p1)
                _ = v9
            }
        case 9223372036854775807:
            v10 := []int{-635, 42, -2147483648}
            _ = v10
        default:
            if false {
                p1 -= 1
                p2 -= 42
            }
        }
        var v11 uint16 = 32767
        _ = v11
    }
    p2 = 42
    fmt.Println(p1, p2)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int, p1 uint) int {
    v12 := []int{(p0 >> ((65535) & 63)), p0, (p0 >> ((p1) & 63)), p0}
    p1 += (p1 + p1)
    v13 := GenS2{F0: 90}
    p0 = int(p1)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 int8 = -49
    var v4 int = 32767
    var v5 uint = 580
    var v14 uint = (v5 & (v5 ^ (v5 / ((0) | 1))))
    switch v4 {
    case 955:
        var v15 uint = (v5 / ((509) | 1))
        _ = v15
    default:
        var v16 float64 = (1.5 - -2.25)
        _ = v16
    }
    var v17 int64 = 149
    v3 = (v3 << (((^(v14))) & 63))
    v17 += 42
    var v18 string = ("a" + "\n\t")
    v19 := GenS3{F0: 824, F1: 2147483647}
    var v20 uint8 = 255
    fmt.Println(v17, v20, v17)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int16 = 0
        var v23 bool = (32767 != 42)
        v20 = 0
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 0}
    v25 := GenS3{F0: int64(v3), F1: 941}
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := v1[v18]
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 float32 = ((3.14159 - -2.25) * (1e10 - (0.0 + (-2.25 - ((-2.25 * -2.25) / float32(-2.25))))))
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    var v30 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
