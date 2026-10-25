package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 42
    v2 := make(map[int]string)
    v3 := []string{}
    v4 := []uint32{69, 754}
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: v1}
    var v7 float32 = (0.0 * 3.14159)
    if false {
        switch "a" {
        case "":
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint = (v1 & (v1 & v1))
                var v10 uint = (v9 | 459)
                _ = fi8
                _ = v10
                _ = v9
            }
        case "a":
            if len(v4) > 0 {
                v4[0] = 1
            }
        default:
            var v11 int8 = 1
            _ = v11
        }
        v1 += (v1 ^ (v1 / ((v1) | 1)))
        v1 -= 0
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v7 += float32(float64(1.5))
        var v13 int16 = 707
        v7 = 1.5
        _ = fi12
        _ = v13
    }
    fmt.Println(v7)
    for fi14 := 0; fi14 < 13; fi14++ {
        if false {
            v7 += (v7 - v7)
            v7 = v7
        } else {
            fmt.Println(v7)
            if (1 != 42) {
                for fi15 := 0; fi15 < 17; fi15++ {
                    switch "foo bar" {
                    case "":
                        var v16 uint32 = 65535
                        v1 += (^(v1))
                        _ = v16
                    }
                    if (v1 > (v1 | (v1 * (v1 >> ((v1) & 63))))) {
                        fmt.Println(v1, v1, v7)
                    } else {
                        var v17 float64 = 0.1
                        _ = v17
                    }
                    var v18 uint16 = 1
                    _ = fi15
                    _ = v18
                }
                var v19 bool = true
                v20 := v2[fi14]
                _ = v19
                _ = v20
            }
        }
        _ = fi14
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float64 = (1e10 * float64((v7 - -2.25)))
        v22 = v22
        _ = fi21
        _ = v22
    }
    switch "a" {
    case "hello":
        var v23 uint = v1
        _ = v23
    case "":
        switch "\n\t" {
        case "\n\t":
            var v24 int = 886
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint8 = 1
                v1 -= 0
                v4 = append(v4, 2147483647)
                _ = fi25
                _ = v26
            }
        case "hello":
            fmt.Println(v7)
        }
    case "foo bar":
        var v27 bool = true
        _ = v27
    }
    switch "hello" {
    case "hello":
        v7 = v7
    case "foo bar":
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int8 = 42
                fmt.Println(v7, v7, fi29)
                v31 := GenS1{F0: 0}
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
    case "":
        if false {
            var v32 uint8 = 42
            var v33 float64 = 0.1
            _ = v32
            _ = v33
        }
    default:
        var v34 uint = 1
        _ = v34
    }
    v4 = append(v4, 65535)
    v35 := GenS1{F0: v1}
    v7 -= v7
    fmt.Println(v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
