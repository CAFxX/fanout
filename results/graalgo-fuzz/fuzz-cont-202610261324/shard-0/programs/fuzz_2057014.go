package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 float32 = (0.0 * (1e10 / 0.1))
    var v4 int64 = -9223372036854775808
    if ((v1 < (v1 + v1)) && false) {
        v3 = 1e10
        var v5 bool = (42 < 933)
        _ = v5
    } else {
        fmt.Println(v1, v4, v1)
        switch v1 {
        case "":
            fmt.Println(v3, v4, v3)
        case "a":
            fmt.Println(v3, v3)
        case "\n\t":
            for fi6 := 0; fi6 < 17; fi6++ {
                delete(v2, (^(fi6)))
                v2[fi6] = (!false)
                _ = fi6
            }
        default:
            for fi7 := 0; fi7 < 5; fi7++ {
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 float32 = v3
                    if true {
                        _ = len(v2)
                    }
                    var v10 int16 = -280
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        }
    }
    var v11 int = 42
    v3 += v3
    switch (v11 + v11) {
    case 879:
        v4 = (v4 - v4)
    }
    var v12 int32 = 1
    if (4294967295 < 42) {
        var v13 string = v1
        fmt.Println(v13, v11)
        v14 := GenS1{F0: 1, F1: (v11 - 32768)}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 bool = (!false)
            fmt.Println(fi15)
            _ = fi16
            _ = v17
        }
        if false {
            v3 = v3
        }
        fmt.Println(v11, v4, v12)
        _ = fi15
    }
    v18 := []int{v11, (v11 / (((v11 >> ((901) & 63))) | 1)), v11, v11}
    switch 128 {
    case -40:
        switch "a" {
        case "foo bar":
            v3 += -2.25
        default:
            v12 += (v12 << ((0) & 63))
        }
    case -577:
        if (((-2.25 < ((1.5 - -2.25) - 0.0)) || false) && (65535 > 42)) {
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 bool = false
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 float32 = (3.14159 + v3)
                v1 = "foo bar"
                v23 := GenS2{F0: 2}
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            switch v1 {
            case "foo bar":
                switch v1 {
                case "a":
                    v2[v11] = true
                default:
                    v24 := GenS1{F0: 0, F1: 1}
                    _ = v24
                }
            case "hello":
                var v25 bool = true
                _ = v25
            default:
                fmt.Println(v12, v12, v4)
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 bool = (v1 <= v1)
                var v28 int = (-(fi26))
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    var v29 string = (v1 + v1)
    v4 = (v4 + 0)
    var v30 int16 = -32768
    v29 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
