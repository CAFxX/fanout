package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() string {
    v6 := []int{}
    switch "\n\t" {
    case "a":
        v7 := GenS1{F0: 84}
        _ = v7
    case "hello":
        v6 = append(v6, 127)
    default:
        if true {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int64 = 42
                var v10 uint16 = 263
                var v11 float32 = (float32(((-2.25 / -2.25) + 3.14159)) - (0.0 * 0.1))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 uint64 = 4294967295
            _ = v12
        } else {
            switch "\n\t" {
            case "hello":
                var v13 uint64 = 1
                _ = v13
            case "foo bar":
                var v14 uint8 = 127
                _ = v14
            default:
                if (-2.25 >= ((1e10 - float32(0.0)) + float32(0.0))) {
                    var v15 bool = false
                    v15 = v15
                    fmt.Println(v15)
                    _ = v15
                }
            }
        }
    }
    v16 := GenS1{F0: 0}
    _ = v16
    _ = v6
    return "foo bar"
}
func genFn1(p0 int, p1 uint) bool {
    p1 += p1
    if (0.1 < 1e10) {
        p0 = p0
        fmt.Println(p0, p1)
    } else {
        var v17 int32 = -586
        _ = v17
    }
    var v18 int = p0
    _ = p0
    _ = p1
    _ = v18
    return (450 <= 42)
}

func main() {
    var v1 int64 = -636
    v2 := []uint32{0, 4294967295, 42, 208}
    var v3 bool = (!((1e10 - 0.0) != 3.14159))
    var v4 string = ""
    var v5 int16 = 42
    v1 = (^(v1))
    v1 = (^(v1))
    v19 := GenS1{F0: 2}
    _ = len(v2)
    if true {
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 string = v4
                v1 = 1
                var v23 string = v4
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: 2}
            _ = fi20
            _ = v24
        }
        var v25 int16 = (v5 & (v5 | (v5 * 223)))
        _ = v25
    }
    v3 = (v3 && (((!false) && (v4 == v4)) || (true || (v5 < v5))))
    var v26 int16 = (v5 << ((18446744073709551615) & 63))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
