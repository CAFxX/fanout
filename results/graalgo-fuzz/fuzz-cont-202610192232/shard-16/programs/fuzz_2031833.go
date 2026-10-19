package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 bool) int {
    var v5 int = 42
    var v6 int8 = 63
    if p2 {
        var v7 int64 = 1
        switch v5 {
        case 32767:
            var v8 float64 = (1.5 - ((1.5 - (1e10 / 1.0)) / ((1e10 / (((3.14159 / (float64(1.5) * 3.14159)) - 0.1) / (1e10 + 0.1))) * 1e10)))
            _ = v8
        }
        var v9 int32 = 2147483647
        _ = v7
        _ = v9
    }
    if (((0.0 * -2.25) - float32(3.14159)) >= ((0.1 / 1.5) + 1.5)) {
        fmt.Println(p0, v5, v6)
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v5)
            p2 = true
            v11 := []int{(^(fi10))}
            _ = fi10
            _ = v11
        }
        v12 := []int{v5, (^(v5)), (-(v5)), (v5 ^ v5)}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(v5))
}

func main() {
    var v1 int32 = 757
    v2 := []string{("foo bar" + (("" + ("a" + "a")) + (("" + "hello") + ("hello" + "\n\t"))))}
    var v3 string = ("" + "")
    v4 := make(map[int]uint32)
    v13 := GenS1{F0: 1}
    for fi14 := 0; fi14 < 12; fi14++ {
        v3 = ("hello" + v3)
        if true {
            v4[(fi14 >> ((18446744073709551615) & 63))] = 806
            fmt.Println(fi14, v1)
        } else {
            fmt.Println(v1, v3)
            v1 -= (v1 >> ((429) & 63))
        }
        _ = fi14
    }
    delete(v4, 32768)
    switch -1 {
    case 0:
        for fi15 := 0; fi15 < 9; fi15++ {
            v3 = "foo bar"
            _ = fi15
        }
    case 2147483647:
        v3 = ("hello" + v3)
    }
    v16 := GenS2{F0: 959, F1: 2, F2: 905}
    switch v3 {
    case "a":
        v17 := GenS2{F0: 116, F1: 255, F2: 2}
        _ = v17
    case "\n\t":
        var v18 int8 = -2
        _ = v18
    case "foo bar":
        fmt.Println(v3, v1)
    }
    v3 = "foo bar"
    v1 -= v1
    var v19 int64 = -533
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
