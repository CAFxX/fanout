package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int8) int {
    var v5 bool = (215 >= 337)
    if (((((1e10 - (((0.1 + 1.5) * -2.25) - (float32((((1e10 * 1e10) + 1.5) / 3.14159)) - 1e10))) * 0.1) * (0.1 + 0.0)) * 1e10) > -2.25) {
        v5 = (127 == 127)
        v5 = v5
        var v6 string = ""
        _ = v6
    }
    var v7 uint = 65535
    var v8 int = -755
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    _ = v8
    return (v8 & (v8 >> ((v7) & 63)))
}
func genFn1(p0 uint8) int {
    var v9 uint16 = 2
    p0 -= p0
    v10 := GenS1{F0: 501}
    v11 := GenS1{F0: 65535}
    var v12 int8 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 662
    v3 := []int{32768, -418}
    v4 := []uint32{375, 3, 4294967295, 0}
    v2 = (^(v2))
    fmt.Println(v2, v2)
    switch "a" {
    case "foo bar":
        if true {
            var v13 uint16 = 300
            _ = v13
        }
    default:
        fmt.Println(v1, v1)
    }
    v2 -= v2
    v1 -= v1
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 = 0
        v2 = -1
        _ = fi14
    }
    if false {
        v3 = append(v3, -814)
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint8 = 127
            _ = fi15
            _ = v16
        }
        var v17 bool = true
        _ = v17
    } else {
        var v18 uint16 = 2
        v4 = append(v4, 42)
        var v19 int = 42
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        if (v2 != v2) {
            var v21 uint32 = 42
            fmt.Println(v21, v2)
            _ = v21
        } else {
            v2 = v2
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v1, v2)
            var v23 uint8 = 1
            _ = fi22
            _ = v23
        }
        _ = fi20
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
