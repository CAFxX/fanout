package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16) int32 {
    p0 = p0
    p0 = 257
    if (0.0 < float64(3.14159)) {
        p0 -= p0
    } else {
        fmt.Println(p0)
    }
    switch "\n\t" {
    case "hello":
        if (32768 != 1) {
            switch 0 {
            case 2147483647:
                fmt.Println(p0)
            case 32767:
                p0 -= (p0 & (p0 * (p0 << ((1) & 63))))
            default:
                switch 32767 {
                case 42:
                    for fi7 := 0; fi7 < 20; fi7++ {
                        v8 := GenS1{F0: 127, F1: int8(p0), F2: 42}
                        var v9 float64 = float64(0.0)
                        var v10 bool = true
                        _ = fi7
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                }
            }
            p0 = 42
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS2{F0: -1, F1: p0, F2: 3}
                v13 := GenS3{F0: p0}
                switch fi11 {
                case -545:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS1{F0: 74, F1: 1, F2: 870}
                        p0 = p0
                        v16 := []int{9223372036854775807, (fi14 << ((18446744073709551615) & 63))}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case 581:
                    v17 := GenS3{F0: (p0 / ((p0) | 1))}
                    _ = v17
                case -9223372036854775808:
                    v18 := GenS3{F0: -32768}
                    _ = v18
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(p0)
        }
    case "a":
        p0 = p0
    }
    p0 -= p0
    _ = p0
    return -694
}

func main() {
    var v1 int64 = 723
    v2 := []string{"foo bar", "\n\t", (("foo bar" + "hello") + "hello")}
    var v3 float64 = (3.14159 / 3.14159)
    var v4 int = -1
    var v5 int64 = (v1 * 0)
    var v6 int64 = v1
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: 99, F1: 0, F2: 0}
        v21 := GenS1{F0: -85, F1: 1, F2: 18446744073709551615}
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        fmt.Println(v3)
    } else {
        var v22 bool = false
        _ = v22
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        if true {
            v24 := GenS2{F0: -51, F1: 42, F2: 30}
            var v25 uint8 = 1
            switch (-(v4)) {
            case -180:
                var v26 int32 = 305
                _ = v26
            }
            _ = v24
            _ = v25
        }
        var v27 int8 = int8(v4)
        var v28 float32 = 3.14159
        _ = fi23
        _ = v27
        _ = v28
    }
    var v29 uint32 = 27
    v30 := GenS3{F0: 1}
    switch 1 {
    case 32767:
        var v31 float32 = (3.14159 + (((1e10 / float32(v3)) * 0.0) + -2.25))
        _ = v31
    case 779:
        v29 -= v29
    case -212:
        v32 := GenS3{F0: -1}
        _ = v32
    default:
        v33 := GenS2{F0: 1, F1: -32768, F2: 1}
        _ = v33
    }
    if false {
        if true {
            v34 := GenS3{F0: -32768}
            var v35 string = "hello"
            fmt.Println(v4)
            _ = v34
            _ = v35
        }
    }
    v6 -= (v1 >> ((42) & 63))
    v5 = -9223372036854775808
    v3 = ((float64((3.14159 / -2.25)) + v3) / (v3 / 1.5))
    fmt.Println(v6)
    fmt.Println(v6, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
