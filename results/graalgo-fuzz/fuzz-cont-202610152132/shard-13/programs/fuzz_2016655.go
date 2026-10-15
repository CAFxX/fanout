package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 4294967295
    var v3 int16 = 1000
    v3 += (^(v3))
    if (!false) {
        v4 := GenS1{F0: (1 != 449)}
        if true {
            if (true && (false || (0.1 == (((float64((0.0 - -2.25)) * (((3.14159 * ((-2.25 * (0.1 - (3.14159 + (0.1 + ((-2.25 * (1e10 - (((float64(1.5) + (1e10 * 1.5)) / float64((-2.25 + (float32(((1e10 - 0.0) - 1e10)) * 1.5)))) / (0.0 - (float64(0.1) / 1e10))))) / 1.5))))) / (1e10 / -2.25))) / (3.14159 - 1.5)) * 1e10)) / (0.1 + 3.14159)) - 0.1)))) {
                v2 -= v2
                v2 -= (v2 & v2)
            }
            var v5 int64 = -453
            _ = v5
        } else {
            v6 := GenS1{F0: (false || (0 <= 0))}
            _ = v6
        }
        _ = v4
    }
    var v7 float64 = (0.0 * (float64(3.14159) / 1.5))
    var v8 int64 = -645
    v2 = 65535
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int64 = (v8 - (v8 ^ (^(v8))))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int8 = 1
        var v13 uint16 = 32767
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 = 1
    v2 = 592
    switch 4294967295 {
    case -862:
        fmt.Println(v7, v3)
    case 200:
        var v14 bool = (1 <= 0)
        _ = v14
    default:
        if true {
            v7 = float64((0.0 * (((1e10 + ((1e10 * 3.14159) / 1e10)) - (-2.25 / 1.0)) - (float32(v7) + 3.14159))))
        }
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int = -2147483648
        v17 := GenS1{F0: (!true)}
        _ = fi15
        _ = v16
        _ = v17
    }
    if (2 > 279) {
        if (-2.25 < v7) {
            if (42 >= 1000) {
                v3 = (v3 << ((18446744073709551615) & 63))
            }
            fmt.Println(v8)
        }
        v3 = 1
        var v18 int = -288
        _ = v18
    } else {
        if (!((true && (!(224 <= v2))) && false)) {
            if (0.0 <= v7) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v1["foo bar"] = v2
                    _ = fi19
                }
                fmt.Println(v8, v7)
            } else {
                fmt.Println(v2)
            }
        }
        v20 := GenS1{F0: true}
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := GenS1{F0: false}
            _ = fi22
            _ = v23
        }
        switch "hello" {
        case "foo bar":
            v8 += 461
        case "hello":
            v3 += v3
        case "a":
            var v24 uint32 = v2
            _ = v24
        default:
            for fi25 := 0; fi25 < 8; fi25++ {
                v2 = v2
                fmt.Println(v8)
                _ = fi25
            }
        }
        switch "a" {
        case "foo bar":
            v26 := GenS1{F0: (fi21 < fi21)}
            _ = v26
        case "":
            var v27 int8 = -1
            _ = v27
        case "a":
            var v28 float32 = (-2.25 / -2.25)
            _ = v28
        default:
            if (!true) {
                v29 := []int{(fi21 / ((fi21) | 1)), fi21, fi21, (fi21 & fi21)}
                _ = v29
            }
        }
        _ = fi21
    }
    var v30 uint32 = v2
    v31 := []int{128, -9223372036854775808}
    fmt.Println(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v7
    _ = v8
}
