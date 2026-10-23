package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
func genFn0() uint8 {
    var v7 uint = 391
    fmt.Println(v7)
    v8 := GenS1{F0: 128, F1: "hello"}
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int16) int64 {
    p0 = (p0 << ((378) & 63))
    var v9 uint32 = 2
    v10 := []int{-501}
    var v11 float64 = (0.0 / 1.0)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{2147483648, 4294967295, 33}
    var v2 int32 = -1
    var v3 int = -87
    var v4 bool = false
    var v5 string = "\n\t"
    var v6 float32 = (((3.14159 + (1e10 - 0.0)) * float32(((((0.0 * 1.5) - ((0.1 + 0.1) - 1e10)) * 0.1) + 0.0))) * -2.25)
    v12 := GenS2{F0: 1}
    if v4 {
        v3 += (v3 >> ((2) & 63))
        v3 -= 0
    } else {
        var v13 uint8 = 255
        switch v3 {
        case 9223372036854775807:
            v13 = 134
        default:
            switch (v5 + "hello") {
            case "hello":
                if v4 {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 string = (v5 + v5)
                        v2 = (v2 & (-(v2)))
                        v5 = v15
                        _ = fi14
                        _ = v15
                    }
                    if (v4 && v4) {
                        fmt.Println(v6, v3)
                        v16 := GenS3{F0: (!v4), F1: 236, F2: 32767}
                        v13 += (v13 * (v13 | (v13 | v13)))
                        _ = v16
                    } else {
                        var v17 uint8 = v13
                        v6 = v6
                        _ = v17
                    }
                }
            case "\n\t":
                var v18 int16 = 32767
                _ = v18
            default:
                v3 -= (^(v3))
            }
        }
        var v19 int32 = (^(v2))
        _ = v13
        _ = v19
    }
    v4 = true
    v2 = 42
    v20 := GenS1{F0: 128, F1: v5}
    for fi21 := 0; fi21 < 6; fi21++ {
        if v4 {
            v1 = append(v1, 2147483648)
        } else {
            var v22 int32 = v2
            fmt.Println(v4, v2, v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
        }
        if true {
            v3 = (fi21 + -664)
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(v4)
                v24 := GenS3{F0: (!(255 < 42)), F1: 706, F2: 1}
                var v25 uint = 0
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: (-(fi21)), F1: (v5 + v5)}
            _ = v26
        } else {
            var v27 uint = 2
            v4 = v4
            v28 := GenS2{F0: 9223372036854775807}
            _ = v27
            _ = v28
        }
        _ = fi21
    }
    v3 = 1
    var v29 float64 = 1e10
    var v30 uint8 = 36
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
