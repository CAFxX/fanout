package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 uint32) int64 {
    var v5 int32 = -2147483648
    v5 += 0
    var v6 int32 = (v5 >> ((1) & 63))
    var v7 string = ("" + "")
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    var v1 bool = (1 > -1)
    var v2 int16 = 42
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v1 = (false && (65535 > 725))
    fmt.Println(v2)
    if v1 {
        if (!v1) {
            v8 := GenS1{F0: 42, F1: 127}
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            v4[418] = "foo bar"
            _ = fi9
        }
        var v10 int = 128
        _ = v10
    }
    v1 = false
    fmt.Println(v1, v2)
    v11 := []int{}
    fmt.Println(v1)
    fmt.Println(v2, v1)
    var v12 bool = true
    if (((((1.5 / 1e10) - (1.5 - 3.14159)) * 0.0) + 0.1) != -2.25) {
        v12 = (!((42 <= 1) || ((!((100 <= -128) && (432 < 42))) && v1)))
        if len(v11) > 0 {
            v11[0] = -942
        }
    } else {
        fmt.Println(v2, v2)
        var v13 int8 = 0
        v2 += (v2 - (v2 * 32767))
        _ = v13
    }
    fmt.Println(v2, v12)
    if len(v11) > 0 {
        v11[0] = 1
    }
    var v14 uint8 = 42
    var v15 uint = 65535
    v16 := GenS2{F0: (^(v2))}
    var v17 int32 = 1
    switch 0 {
    case -584:
        var v18 uint16 = 32767
        _ = v18
    case -1:
        var v19 int = 128
        _ = v19
    default:
        switch ("" + "\n\t") {
        case "":
            var v20 int = 175
            _ = v20
        case "a":
            switch "" {
            case "\n\t":
                v21 := GenS1{F0: 1, F1: 61}
                _ = v21
            case "foo bar":
                v14 = 6
            }
        }
    }
    v3["\n\t"] = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
