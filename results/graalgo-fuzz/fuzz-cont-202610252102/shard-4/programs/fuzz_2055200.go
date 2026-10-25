package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int }
func genFn0() uint64 {
    var v4 uint64 = 0
    if (((0.1 + 3.14159) * 0.1) > 1.5) {
        var v5 uint64 = (v4 * (v4 * v4))
        v5 = (^(v4))
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint8 = 209
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint64 = v4
            fmt.Println(fi8, v9)
            _ = fi8
            _ = v9
        }
        v4 = 2
        v4 -= 42
    }
    _ = v4
    return (v4 << ((0) & 63))
}

func main() {
    var v1 int64 = 42
    var v2 int = 4294967295
    var v3 int16 = -271
    if ("foo bar" == "\n\t") {
        v10 := GenS1{F0: 834, F1: -1, F2: (v1 - v1)}
        var v11 uint16 = 65535
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v1 = v1
        var v13 string = "a"
        _ = fi12
        _ = v13
    }
    v14 := []int{(v2 / ((v2) | 1)), v2, v2, v2}
    v3 = (^(v3))
    v15 := GenS2{F0: 2}
    _ = len(v14)
    if (255 > 911) {
        var v16 int16 = (v3 % (((^(v3))) | 1))
        _ = v16
    } else {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        fmt.Println(v2, v2)
        var v17 uint8 = 1
        _ = v17
    }
    v18 := GenS3{F0: 0, F1: 2}
    if ((^(v2)) < (-(v2))) {
        v19 := GenS4{F0: 127, F1: 42}
        for fi20 := 0; fi20 < 17; fi20++ {
            if (v1 <= v1) {
                var v21 uint64 = 1
                v22 := GenS5{F0: 123}
                v1 -= (-(v1))
                _ = v21
                _ = v22
            } else {
                switch fi20 {
                case 0:
                    var v23 float32 = 1e10
                    _ = v23
                case 128:
                    v1 = v1
                default:
                    var v24 int = fi20
                    _ = v24
                }
            }
            v25 := GenS3{F0: 100, F1: 42}
            _ = fi20
            _ = v25
        }
        _ = v19
    } else {
        switch 1 {
        case 0:
            v2 = v2
        }
        switch (v2 << ((18446744073709551615) & 63)) {
        case 2147483647:
            var v26 uint32 = 219
            _ = v26
        case 128:
            switch 42 {
            case 32767:
                v27 := GenS6{F0: 65535, F1: v2}
                _ = v27
            case 32768:
                fmt.Println(v3, v3, v1)
            case 869:
                fmt.Println(v1, v2)
            default:
                var v28 uint32 = 287
                _ = v28
            }
        }
    }
    var v29 uint16 = 255
    v14 = append(v14, v2)
    if false {
        if false {
            v14 = append(v14, 2147483647)
        }
        v3 += (-(v3))
        _ = len(v14)
    } else {
        fmt.Println(v2, v29, v1)
        switch (("foo bar" + "hello") + "hello") {
        case "foo bar":
            v30 := GenS6{F0: 274, F1: 9223372036854775807}
            _ = v30
        case "":
            v3 = v3
        }
    }
    fmt.Println(v29)
    if false {
        if (127 == 100) {
            var v31 int64 = (v1 ^ v1)
            v31 -= 9223372036854775807
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v31
        } else {
            switch (v2 % ((0) | 1)) {
            case -388:
                var v32 int32 = 1
                _ = v32
            default:
                v1 += (^(v1))
            }
            v1 += v1
            v29 -= v29
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            v2 += 42
            v34 := GenS3{F0: 42, F1: 42}
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 int8 = 79
                var v37 float64 = -2.25
                v38 := GenS3{F0: 0, F1: 0}
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi33
            _ = v34
        }
        v29 -= v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v29
    _ = v3
}
