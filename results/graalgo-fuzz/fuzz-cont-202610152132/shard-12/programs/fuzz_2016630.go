package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 int64, p1 string, p2 uint64) uint8 {
    var v6 uint8 = 255
    fmt.Println(p0, p0)
    v7 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v6 - v6)
}

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{2147483647}
    v3 := make(map[string]int)
    var v4 float32 = float32((float64(3.14159) * v1))
    v5 := []string{"\n\t"}
    _ = len(v3)
    for fi8 := 0; fi8 < 12; fi8++ {
        switch (-(fi8)) {
        case -409:
            v4 -= ((float32(v1) / v4) + 0.1)
        case 42:
            var v9 int32 = 2147483647
            _ = v9
        }
        var v10 uint = 519
        v1 += v1
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint = 42
        var v13 uint16 = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 7; fi14++ {
            switch (fi14 * fi14) {
            case 0:
                fmt.Println(v4)
            default:
                v15 := GenS1{F0: "", F1: (fi14 | 32767)}
                _ = v15
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint8 = 2
                _ = fi16
                _ = v17
            }
            _ = fi14
        }
        if (123 > 1) {
            v1 = float64(v4)
        }
    } else {
        var v18 float32 = ((v4 / 1e10) * 3.14159)
        v2 = append(v2, 2)
        _ = v18
    }
    var v19 uint64 = 0
    v20 := GenS1{F0: (("foo bar" + "a") + "foo bar"), F1: -589}
    var v21 uint8 = 186
    var v22 int16 = -32768
    switch "\n\t" {
    case "":
        switch -2147483648 {
        case -2147483648:
            var v23 int = 0
            _ = v23
        case -375:
            delete(v3, "\n\t")
        case 128:
            v21 += (v21 >> ((uint(v19)) & 63))
        default:
            if (!true) {
                v21 = 127
            } else {
                v19 = (^(v19))
                var v24 uint32 = 2147483648
                _ = v24
            }
        }
    case "a":
        v25 := GenS1{F0: ("foo bar" + "a"), F1: 857}
        _ = v25
    case "hello":
        v22 -= (v22 - v22)
    }
    v22 = (v22 - (v22 | v22))
    var v26 string = "foo bar"
    v19 = (^(v19))
    fmt.Println(v4, v4)
    v27 := GenS1{F0: v26, F1: 753}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
