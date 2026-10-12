package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]uint32)
    var v3 int8 = 42
    var v4 int16 = 0
    var v5 uint16 = 42
    v6 := make(map[string]int)
    v1 += 31
    v7 := GenS1{F0: 65535}
    fmt.Println(v3, v4)
    v4 = v4
    v5 = (v5 << ((4294967295) & 63))
    if true {
        var v8 float32 = 1e10
        var v9 int16 = v4
        if false {
            switch 32767 {
            case 127:
                v10 := []int{9223372036854775807}
                _ = v10
            case 32768:
                fmt.Println(v5, v5, v5)
            case 32767:
                fmt.Println(v8, v8)
            }
        }
        _ = v8
        _ = v9
    } else {
        v2[4294967295] = 980
    }
    if (!((-13 <= 337) || true)) {
        v4 = (v4 / (((v4 | (v4 * (^(v4))))) | 1))
        v11 := GenS1{F0: uint(v1)}
        if ((!false) && (-2.25 != -2.25)) {
            v12 := []int{}
            var v13 uint64 = 42
            v14 := GenS1{F0: 167}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float64 = float64(float32(0.0))
            v5 = v5
            v4 -= 1
            _ = fi15
            _ = v16
        }
    }
    switch "\n\t" {
    case "\n\t":
        v5 = uint16(v1)
    case "hello":
        v1 = (v1 >> ((101) & 63))
    case "a":
        var v17 bool = (v3 != (v3 ^ (v3 << ((2) & 63))))
        _ = v17
    default:
        var v18 int16 = (v4 ^ v4)
        _ = v18
    }
    v1 = v1
    switch 32767 {
    case 1:
        v5 = v5
    case 2147483647:
        v19 := GenS2{F0: 4294967295, F1: (v3 >> ((204) & 63))}
        _ = v19
    case 32767:
        if ((v4 == v4) && true) {
            var v20 uint64 = 42
            if ((true || (false && false)) || (741 < 1)) {
                switch -538 {
                case 1:
                    for fi21 := 0; fi21 < 1; fi21++ {
                        fmt.Println(v3)
                        _ = fi21
                    }
                case -449:
                    if ((("hello" + "hello") + ("foo bar" + "hello")) > "foo bar") {
                        v4 += 876
                        var v22 bool = true
                        _ = v22
                    } else {
                        var v23 uint = uint(v1)
                        v20 = v20
                        _ = v23
                    }
                default:
                    fmt.Println(v5)
                }
                v24 := v2[2147483648]
                _ = v24
            }
            _ = v20
        } else {
            var v25 int16 = 32767
            v26 := GenS3{F0: 849}
            v4 = 1
            _ = v25
            _ = v26
        }
    default:
        v3 += (^(v3))
    }
    v27 := []int{4294967295, 2147483647, 9223372036854775807}
    delete(v6, "\n\t")
    for fi28 := 0; fi28 < 15; fi28++ {
        fmt.Println(fi28)
        switch (^(fi28)) {
        case 906:
            v29 := GenS4{F0: 32767}
            _ = v29
        default:
            fmt.Println(v3, v1, v4)
        }
        if (true && true) {
            var v30 float32 = (3.14159 + float32(0.1))
            v31 := GenS5{F0: 100}
            var v32 int32 = -2147483648
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
