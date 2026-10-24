package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16) int8 {
    var v5 bool = (1 > -118)
    switch "foo bar" {
    case "hello":
        var v6 uint64 = 0
        _ = v6
    }
    if false {
        p0 = 407
        var v7 uint32 = 2147483647
        _ = v7
    }
    p0 += (p0 / ((65535) | 1))
    _ = p0
    _ = v5
    return -75
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 69
    var v3 uint8 = 42
    var v4 float64 = 1.5
    v8 := []int{128, -232, -870}
    if (42 == 65535) {
        var v9 float32 = 0.0
        v10 := GenS1{F0: false}
        _ = v10
        _ = v9
    }
    fmt.Println(v1, v4)
    v4 += (((1e10 * 0.1) / v4) / ((v4 / (v4 + (float64((1.5 + 1e10)) * ((float64((0.0 + 0.0)) - 1.5) + (v4 - (v4 + (v4 - 3.14159))))))) * ((float64(3.14159) - v4) + v4)))
    switch 9223372036854775807 {
    case 935:
        if ((v3 * v3) != (v3 << ((18446744073709551615) & 63))) {
            switch -1 {
            case 68:
                v11 := GenS2{F0: 1, F1: 1}
                _ = v11
            case -881:
                v2 = (v2 >> ((255) & 63))
            case -1:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            var v12 uint = 504
            _ = v12
        } else {
            var v13 uint16 = v2
            for fi14 := 0; fi14 < 19; fi14++ {
                switch (-(fi14)) {
                case 32767:
                    var v15 int64 = -1
                    _ = v15
                case 0:
                    if (((^(v3)) > (v3 >> ((0) & 63))) && (-32768 != 42)) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        switch fi14 {
                        case 32768:
                            v8 = append(v8, (fi14 * (fi14 >> ((620) & 63))))
                        case -2147483648:
                            fmt.Println(v1)
                        case 127:
                            fmt.Println(v13, v4, v4)
                        }
                        fmt.Println(v1, v3, fi14)
                    } else {
                        var v16 string = "hello"
                        switch (v1 + "foo bar") {
                        case "foo bar":
                            for fi17 := 0; fi17 < 3; fi17++ {
                                v13 -= v13
                                var v18 int = fi17
                                _ = fi17
                                _ = v18
                            }
                        default:
                            v4 = (float64((-2.25 - (((-2.25 * 0.0) / 1.0) * 0.1))) + v4)
                        }
                        _ = v16
                    }
                case 9223372036854775807:
                    v19 := GenS2{F0: (v2 & (v2 ^ (^(v2)))), F1: (^(fi14))}
                    _ = v19
                }
                _ = fi14
            }
            var v20 uint16 = v2
            _ = v13
            _ = v20
        }
    case 42:
        v1 = ("foo bar" + v1)
    case -742:
        fmt.Println(v3)
    }
    _ = len(v8)
    if (true || true) {
        var v21 uint64 = 42
        v3 = uint8(v2)
        var v22 uint8 = (^(v3))
        _ = v21
        _ = v22
    } else {
        _ = len(v8)
        var v23 int32 = -392
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v3 += 128
        v1 = "\n\t"
        var v25 int64 = -672
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v8)
        switch "a" {
        case "a":
            var v27 int64 = 1
            _ = v27
        case "foo bar":
            v28 := GenS3{F0: 9223372036854775807}
            _ = v28
        default:
            var v29 uint = 255
            _ = v29
        }
        _ = fi26
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
