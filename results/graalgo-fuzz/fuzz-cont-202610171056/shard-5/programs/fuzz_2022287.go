package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 float32 = 1.5
    var v4 string = ("hello" + ("foo bar" + "\n\t"))
    if true {
        v4 = "foo bar"
    } else {
        var v5 uint8 = 127
        if (2 <= 0) {
            v5 += (v5 ^ (v5 - (v5 + (v5 ^ 138))))
        }
        _ = v5
    }
    v3 = float32((3.14159 + (1e10 + (1.5 - 0.0))))
    v4 = (v4 + "a")
    switch v4 {
    case "a":
        switch "a" {
        case "\n\t":
            var v6 float64 = -2.25
            _ = v6
        case "foo bar":
            var v7 int64 = -561
            _ = v7
        case "":
            delete(v1, 591)
        default:
            delete(v1, 281)
        }
    case "":
        var v8 float32 = (v3 * v3)
        _ = v8
    case "hello":
        if true {
            v3 -= v3
        }
    default:
        v9 := GenS1{F0: 1, F1: -32768, F2: (127 > -1)}
        _ = v9
    }
    v4 = v4
    fmt.Println(v3)
    v3 -= v3
    switch -487 {
    case 128:
        var v10 int8 = 127
        _ = v10
    default:
        v11 := v1[-9223372036854775808]
        _ = v11
    }
    switch -979 {
    case 2147483647:
        fmt.Println(v4, v3)
    case -2147483648:
        var v12 bool = (-778 == 1)
        _ = v12
    default:
        var v13 string = "a"
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: 0, F1: 652, F2: true}
            fmt.Println(fi15)
            v4 = (v4 + "a")
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := []int{127}
    if ((122 < 2) && (116 >= -1)) {
        v3 += -2.25
        var v18 int32 = -2147483648
        _ = v18
    } else {
        v19 := GenS1{F0: 0, F1: 32767, F2: true}
        _ = v19
    }
    var v20 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
