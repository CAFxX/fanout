package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 int32) string {
    if true {
        v5 := GenS1{F0: 9223372036854775807, F1: 42, F2: -32768}
        _ = v5
    }
    v6 := GenS2{F0: -32768, F1: 1, F2: "a"}
    var v7 uint64 = 2
    p0 = p0
    v7 -= 2
    _ = p0
    _ = v6
    _ = v7
    return ((("" + "a") + ((("\n\t" + "") + "a") + "hello")) + "hello")
}

func main() {
    v1 := []uint32{}
    var v2 bool = false
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    switch -9223372036854775808 {
    case 1:
        v2 = v2
    case 128:
        if (!true) {
            fmt.Println(v2)
            fmt.Println(v2)
            v2 = true
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(fi8, fi8)
                v2 = (!(2 < 1))
                _ = fi8
            }
        }
    default:
        var v9 uint16 = 466
        _ = v9
    }
    delete(v3, "hello")
    var v10 int8 = -128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v10 = v10
        var v12 int16 = -844
        _ = fi11
        _ = v12
    }
    var v13 string = "hello"
    switch 0 {
    case 82:
        switch -2147483648 {
        case 32768:
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 uint16 = 523
                v16 := GenS3{F0: 124, F1: 96}
                _ = fi14
                _ = v15
                _ = v16
            }
        case -845:
            v2 = v2
        default:
            v3[v13] = "hello"
        }
    case 1:
        fmt.Println(v10, v2, v10)
    case 203:
        v4[4294967295] = v13
    default:
        var v17 string = ((v13 + v13) + "a")
        _ = v17
    }
    switch v13 {
    case "":
        if (4294967295 != 1) {
            if (v13 == "\n\t") {
                v2 = v2
                v13 = "hello"
                var v18 float64 = ((1.5 * 1e10) - 0.1)
                _ = v18
            }
            var v19 uint = 1
            var v20 int16 = -1
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 float32 = (((1e10 + 1.5) - 0.1) / -2.25)
                _ = fi21
                _ = v22
            }
        }
    case "foo bar":
        if v2 {
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: 626, F1: v10, F2: v13}
                _ = fi23
                _ = v24
            }
            var v25 uint32 = 228
            var v26 float32 = 0.0
            _ = v25
            _ = v26
        }
    case "a":
        var v27 int = 42
        _ = v27
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (fi28 < fi28) {
                v2 = v2
                v10 = (-(v10))
                if v2 {
                    var v29 uint32 = 695
                    v10 -= v10
                    v10 -= v10
                    _ = v29
                } else {
                    v13 = "hello"
                }
            }
            _ = fi28
        }
    }
    v30 := GenS1{F0: 138, F1: v10, F2: 1000}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
