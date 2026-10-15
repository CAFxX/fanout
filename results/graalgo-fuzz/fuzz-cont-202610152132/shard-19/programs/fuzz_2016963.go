package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0() uint {
    switch 32767 {
    case 2147483647:
        var v5 uint16 = 0
        _ = v5
    }
    v6 := GenS1{F0: 2, F1: 821, F2: 65535}
    if false {
        v7 := GenS1{F0: 18446744073709551615, F1: 32767, F2: 725}
        _ = v7
    } else {
        v8 := []int{127, 1, -601}
        if len(v8) > 0 {
            v8[0] = 0
        }
        var v9 float32 = float32((1.5 + 0.0))
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: "a"}
    var v11 float64 = (1.5 * 1.5)
    _ = v10
    _ = v11
    _ = v6
    return 4294967295
}

func main() {
    var v1 bool = false
    var v2 int16 = 0
    v3 := make(map[int]int)
    var v4 float64 = 1.5
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v2)
        if (v1 && v1) {
            fmt.Println(v2, fi12)
            var v13 string = "hello"
            for fi14 := 0; fi14 < 2; fi14++ {
                v1 = true
                var v15 uint8 = 127
                _ = len(v3)
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            v1 = (!v1)
            switch 127 {
            case -9223372036854775808:
                v1 = v1
            case -74:
                var v16 int32 = -528
                _ = v16
            case 32768:
                var v17 string = (("a" + "foo bar") + "a")
                _ = v17
            }
        }
        var v18 float32 = 0.1
        _ = fi12
        _ = v18
    }
    var v19 uint64 = 4294967295
    v1 = v1
    if (-1 < 42) {
        if false {
            var v20 int8 = 127
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int = fi21
                _ = fi21
                _ = v22
            }
        }
        var v23 float64 = (((v4 - (float64(3.14159) * (v4 / (1.5 / 0.1)))) / v4) * v4)
        fmt.Println(v19, v19)
        _ = v23
    } else {
        v24 := GenS3{F0: 42, F1: (^(v2))}
        _ = v24
    }
    var v25 uint = 878
    for fi26 := 0; fi26 < 2; fi26++ {
        v3[-909] = (fi26 << ((genFn0()) & 63))
        switch "" {
        case "\n\t":
            for fi27 := 0; fi27 < 6; fi27++ {
                fmt.Println(v4)
                var v28 int32 = -2147483648
                v4 = 1e10
                var v29 int16 = -1
                _ = fi27
                _ = v28
                _ = v29
            }
        case "":
            if v1 {
                v19 += v19
                if (((0.0 + (((-2.25 - (float32(float64(1.5)) / (1e10 - -2.25))) + 1.5) + 0.0)) * (0.1 + (1.5 * ((0.1 / 1e10) - 1e10)))) < (-2.25 / ((1.5 / (-2.25 / (0.0 + 1e10))) + -2.25))) {
                    fmt.Println(fi26, v4)
                } else {
                    v2 = v2
                }
            }
        case "foo bar":
            v3[(fi26 & fi26)] = fi26
        default:
            v1 = v1
        }
        v25 = (v25 % ((18446744073709551615) | 1))
        _ = fi26
    }
    v30 := GenS1{F0: 1, F1: 2, F2: 474}
    v2 -= (v2 + -713)
    switch 32768 {
    case -264:
        fmt.Println(v1, v25, v2)
    case 1:
        fmt.Println(v1)
    default:
        var v31 uint = (^(v25))
        _ = v31
    }
    var v32 int8 = 2
    fmt.Println(v4, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
