package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 int16, p2 uint) string {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 42
        _ = fi6
        _ = v7
    }
    if false {
        fmt.Println(p0, p1)
        var v8 float64 = (1e10 / 0.1)
        _ = v8
    } else {
        var v9 int16 = -32768
        v10 := GenS1{F0: 0, F1: 0, F2: 4294967295}
        _ = v10
        _ = v9
    }
    if false {
        v11 := GenS2{F0: 42, F1: p2}
        fmt.Println(p2, p1, p0)
        _ = v11
    }
    fmt.Println(p2)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 float32 = -2.25
    var v2 float64 = -2.25
    var v3 int16 = 32767
    var v4 string = "foo bar"
    var v5 uint64 = 214
    var v12 int = 0
    v5 += (v5 << ((2) & 63))
    var v13 float32 = float32(float64((v1 + (1e10 / (v1 * ((float32(v2) * 0.0) + (v1 * (1e10 - v1))))))))
    fmt.Println(v1, v4, v2)
    switch v4 {
    case "":
        v14 := GenS2{F0: -86, F1: 2}
        _ = v14
    case "a":
        fmt.Println(v2, v1)
    case "hello":
        for fi15 := 0; fi15 < 9; fi15++ {
            if false {
                v16 := GenS1{F0: v3, F1: 2, F2: (v5 * (v5 << ((2) & 63)))}
                var v17 int = (fi15 >> ((18446744073709551615) & 63))
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: v5}
            _ = fi15
            _ = v18
        }
    }
    fmt.Println(v3)
    switch "a" {
    case "\n\t":
        v19 := GenS1{F0: v3, F1: 42, F2: (v5 - v5)}
        _ = v19
    case "hello":
        var v20 uint32 = 710
        _ = v20
    case "a":
        v21 := []int{v12, -340, v12, (v12 << ((2) & 63))}
        _ = v21
    }
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
