package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 uint8) uint16 {
    p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
    switch "\n\t" {
    case "hello":
        v6 := GenS1{F0: 1, F1: 0, F2: 2}
        _ = v6
    case "":
        v7 := []int{389, 32768, -4}
        _ = v7
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 = p0
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: 0, F1: 295, F2: 0}
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: (fi8 / (((fi8 ^ fi8)) | 1)), F1: 1}
            _ = fi8
            _ = v11
        }
    }
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{}
    var v3 uint16 = 845
    v4 := make(map[string]uint32)
    var v5 string = (("foo bar" + "foo bar") + "foo bar")
    if (true && false) {
        var v12 uint = 255
        switch 0 {
        case 4294967295:
            v13 := GenS1{F0: 1, F1: (v12 / ((100) | 1)), F2: (v12 / (((^(v12))) | 1))}
            _ = v13
        case -494:
            var v14 bool = ((!(!(!false))) && ((v1 ^ -742) <= v1))
            _ = v14
        case 32768:
            v5 = "a"
        }
        _ = v12
    } else {
        v2 = append(v2, 51)
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int16 = (v1 >> ((uint(v3)) & 63))
            v5 = ""
            _ = fi15
            _ = v16
        }
    }
    delete(v4, v5)
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint = 227
        _ = fi17
        _ = v18
    }
    fmt.Println(v1, v3, v5)
    var v19 int32 = 42
    v20 := v4["a"]
    var v21 int = -575
    var v22 float64 = (1.5 * float64(0.0))
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint32 = v20
        v19 += v19
        v25 := GenS2{F0: v21, F1: 2147483647}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int32 = v19
    if (!false) {
        v3 = 1
        v20 = v20
    } else {
        v19 -= v26
        var v27 uint16 = (v3 / ((genFn0(128)) | 1))
        v4[v5] = 108
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 int64 = 9223372036854775807
        fmt.Println(v26)
        if ((v1 & (v1 | v1)) > -32768) {
            v20 += 1
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
