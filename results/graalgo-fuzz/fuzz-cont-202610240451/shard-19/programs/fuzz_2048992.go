package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 uint, p1 uint16) string {
    var v4 int64 = 0
    var v5 int64 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint64 = 912
    var v2 int = -1
    v3 := make(map[int]bool)
    switch ("\n\t" + ("" + ("" + "foo bar"))) {
    case "\n\t":
        var v6 bool = (true || (127 != 255))
        _ = v6
    }
    v7 := GenS1{F0: "\n\t", F1: v1, F2: 100}
    switch (v2 - v2) {
    case 4294967295:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float64 = (((float64(0.0) * ((0.1 + (((((3.14159 * (1e10 / float64(1e10))) + 0.1) + 1.5) - 0.0) / -2.25)) - 0.0)) + (3.14159 - (1.5 / 1e10))) + (((1e10 / 3.14159) / 1.0) - (3.14159 / (-2.25 / 1.0))))
            _ = fi8
            _ = v9
        }
    case 32768:
        if ((false && true) && (false || true)) {
            v2 = v2
        }
    default:
        v2 = 1
    }
    v10 := GenS2{F0: 255}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 bool = (9223372036854775807 == 42)
        switch -396 {
        case 32768:
            var v13 uint32 = 387
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    if (0 <= 42) {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS3{F0: 255, F1: 1, F2: 0}
            switch fi14 {
            case 465:
                fmt.Println(fi14)
            }
            v2 = fi14
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS4{F0: -661, F1: v2}
    v1 = v1
    switch v2 {
    case 1:
        switch (v2 * 128) {
        case -1:
            var v17 uint32 = 384
            _ = v17
        case 32768:
            var v18 int64 = 42
            _ = v18
        }
    case 128:
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS4{F0: 42, F1: (v2 / ((-512) | 1))}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v1)
    v2 -= v2
    if true {
        switch ((("foo bar" + "hello") + "hello") + "") {
        case "\n\t":
            v21 := GenS3{F0: 42, F1: 65535, F2: 65535}
            _ = v21
        }
        v22 := GenS3{F0: 128, F1: 2147483647, F2: 2}
        _ = v22
    } else {
        var v23 int64 = -1
        _ = v23
    }
    var v24 uint32 = 500
    v2 = 2147483647
    fmt.Println(v24, v2, v24)
    var v25 uint64 = v1
    if (-1 == -987) {
        var v26 uint16 = 0
        v24 += (v24 * 42)
        for fi27 := 0; fi27 < 8; fi27++ {
            switch "a" {
            case "foo bar":
                v28 := GenS3{F0: 28, F1: 39, F2: 0}
                _ = v28
            case "\n\t":
                v29 := []int{(v2 % (((-(v2))) | 1)), fi27}
                _ = v29
            default:
                if (42 > 616) {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v26 = (v26 & v26)
                        v25 -= v25
                        switch 699 {
                        case 435:
                            v31 := GenS5{F0: "", F1: 32768}
                            _ = v31
                        default:
                            v25 = v25
                        }
                        _ = fi30
                    }
                } else {
                    var v32 uint32 = (v24 + v24)
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 float32 = 0.0
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 11; fi35++ {
                        v24 -= (v24 + v24)
                        _ = fi35
                    }
                    _ = v32
                }
            }
            _ = fi27
        }
        _ = v26
    }
    if true {
        switch (v2 | v2) {
        case -57:
            v25 = (^(v1))
        case 0:
            v36 := []int{(v2 - (v2 >> ((2) & 63)))}
            _ = v36
        default:
            v2 = v2
        }
        delete(v3, (-(v2)))
        fmt.Println(v25, v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
