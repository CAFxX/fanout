package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64) string {
    if (0 != -1) {
        var v5 float32 = 0.0
        _ = v5
    } else {
        if (int8(p0) >= -12) {
            v6 := []int{2147483647, 128}
            _ = v6
        } else {
            switch "a" {
            case "\n\t":
                var v7 uint = 638
                _ = v7
            }
            p0 -= 544
            v8 := []int{364, -732, int(p0), 32767}
            _ = v8
        }
    }
    var v9 uint16 = 42
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v9)
        var v11 float64 = -2.25
        v12 := GenS1{F0: -499}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []int{}
    var v2 int64 = -9223372036854775808
    var v3 int64 = v2
    v4 := make(map[int]string)
    fmt.Println(v3)
    fmt.Println(v2, v3)
    switch "hello" {
    case "":
        v13 := v4[-179]
        _ = v13
    }
    fmt.Println(v3, v2)
    v14 := GenS1{F0: -861}
    if true {
        for fi15 := 0; fi15 < 17; fi15++ {
            switch "foo bar" {
            case "foo bar":
                v3 += 42
            case "a":
                var v16 int32 = -548
                _ = v16
            default:
                if false {
                    var v17 uint32 = 4294967295
                    var v18 int64 = v3
                    _ = len(v1)
                    _ = v17
                    _ = v18
                }
            }
            _ = fi15
        }
        var v19 int = 498
        var v20 int64 = (v3 + (v2 | v3))
        _ = v19
        _ = v20
    }
    var v21 float32 = (0.1 * float32((((0.1 / 3.14159) + 1.5) / float64(-2.25))))
    v22 := GenS1{F0: 1}
    if (-696 == v3) {
        v3 = v3
        for fi23 := 0; fi23 < 13; fi23++ {
            v24 := GenS1{F0: -1}
            v3 += 176
            _ = fi23
            _ = v24
        }
        var v25 int16 = -195
        _ = v25
    } else {
        var v26 float32 = (v21 + (float32(1.5) - 0.0))
        var v27 uint32 = 42
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
