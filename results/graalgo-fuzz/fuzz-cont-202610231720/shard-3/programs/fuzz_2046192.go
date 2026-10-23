package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 uint16) int32 {
    var v5 float64 = -2.25
    fmt.Println(p0)
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 int) uint64 {
    v6 := []int{p0, (p0 / ((303) | 1)), -1, (p0 + p0)}
    var v7 int64 = 436
    var v8 int = (p0 << ((2) & 63))
    var v9 string = "a"
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{143, 2147483647, 42}
    v3 := []uint32{}
    v4 := []int{-122, 127, 0, 9223372036854775807}
    switch "hello" {
    case "a":
        var v10 uint32 = 65535
        _ = v10
    case "hello":
        var v11 int8 = 51
        _ = v11
    case "foo bar":
        switch 4294967295 {
        case 32767:
            var v12 int16 = 1
            _ = v12
        case -349:
            var v13 float64 = 1.5
            _ = v13
        case 0:
            if (((3.14159 - 1.5) + 0.1) != ((float64(3.14159) - 0.1) - 0.1)) {
                switch 42 {
                case 9223372036854775807:
                    if len(v2) > 0 {
                        v2[0] = 852
                    }
                case 42:
                    if (-9223372036854775808 < 9223372036854775807) {
                        v1 = (v1 % (((v1 / ((42) | 1))) | 1))
                        v1 -= (v1 | v1)
                    } else {
                        var v14 int64 = -639
                        v14 -= (v14 | v14)
                        var v15 int64 = (v14 % ((-1) | 1))
                        _ = v14
                        _ = v15
                    }
                default:
                    v1 += 233
                }
            } else {
                var v16 int16 = -32768
                if ((v16 / ((1000) | 1)) > (v16 & 32767)) {
                    var v17 uint = 2
                    _ = v17
                }
                _ = v16
            }
        default:
            fmt.Println(v1)
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch ("a" + "hello") {
    case "foo bar":
        var v18 string = (("\n\t" + "") + "a")
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: 255}
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if true {
            if false {
                if len(v3) > 0 {
                    v3[0] = 2
                }
            } else {
                v20 := GenS1{F0: 2147483647, F1: 2}
                _ = v20
            }
        }
    }
    fmt.Println(v1)
    var v21 uint16 = 2
    var v22 float32 = float32((1e10 / -2.25))
    switch 641 {
    case 510:
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int16 = 0
            fmt.Println(v24, v24, v24)
            var v25 uint = 2
            _ = fi23
            _ = v24
            _ = v25
        }
    case 32768:
        v1 -= (v1 >> ((65535) & 63))
    }
    fmt.Println(v22)
    switch "foo bar" {
    case "":
        v21 = (v21 << ((4294967295) & 63))
    case "hello":
        if (true || (true || true)) {
            if ((0.1 / 0.1) >= 0.1) {
                fmt.Println(v21, v21, v21)
                v26 := GenS1{F0: 781, F1: 248}
                var v27 uint32 = 1
                _ = v26
                _ = v27
            }
        }
    default:
        v21 = 2
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v22)
        if (!true) {
            var v28 int16 = 32767
            switch "" {
            case "":
                v29 := GenS1{F0: 284, F1: 0}
                _ = v29
            }
            _ = v28
        } else {
            v30 := GenS1{F0: 4294967295, F1: 2}
            _ = v30
        }
        var v31 uint8 = (v1 >> ((uint(v1)) & 63))
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
