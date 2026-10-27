package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float32 = ((3.14159 * 0.1) * -2.25)
    v2 := []uint32{}
    var v3 int8 = 127
    v4 := []uint32{771, 4294967295, 286}
    v5 := []string{}
    v6 := []uint32{42}
    v1 += 3.14159
    var v7 uint64 = 1
    if (1 == 42) {
        fmt.Println(v7, v1, v3)
        var v8 int64 = 1
        var v9 uint64 = (^(v7))
        _ = v8
        _ = v9
    } else {
        var v10 float64 = (-2.25 * 3.14159)
        if (!(false || true)) {
            switch "" {
            case "\n\t":
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 int8 = v3
                    v12 = v12
                    var v13 float32 = 0.0
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                v14 := GenS1{F0: 987, F1: 0}
                _ = v14
            }
            switch "" {
            case "a":
                v2 = append(v2, 732)
            case "\n\t":
                v6 = append(v6, 107)
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint32 = 40
                v3 += (v3 % (((v3 / ((v3) | 1))) | 1))
                var v17 string = "foo bar"
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            v10 -= (v10 / v10)
        }
        fmt.Println(v1, v3)
        _ = v10
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v7 += v7
        var v19 uint64 = v7
        _ = fi18
        _ = v19
    }
    var v20 int16 = -32768
    var v21 int32 = 2147483647
    v3 += 1
    v21 += (v21 / (((v21 >> ((855) & 63))) | 1))
    var v22 int32 = v21
    for fi23 := 0; fi23 < 1; fi23++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v24 uint64 = v7
        for fi25 := 0; fi25 < 14; fi25++ {
            v2 = append(v2, 2147483647)
            switch "hello" {
            case "a":
                v26 := GenS2{F0: 42}
                _ = v26
            default:
                var v27 uint16 = 1
                var v28 int8 = v3
                fmt.Println(v7)
                _ = v27
                _ = v28
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v1, v3)
    fmt.Println(v3, v22)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        _ = len(v2)
        var v29 uint = 0
        _ = v29
    } else {
        var v30 int64 = 0
        var v31 int16 = v20
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
