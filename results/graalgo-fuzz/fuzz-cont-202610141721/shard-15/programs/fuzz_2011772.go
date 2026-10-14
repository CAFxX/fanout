package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0() uint64 {
    var v7 uint = 157
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: "\n\t", F1: (((0.0 * 1e10) + 1.5) == (0.0 - (((0.0 - float32(0.1)) / -2.25) - 0.0))), F2: 1}
        var v10 uint32 = 820
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (("hello" + "a") + "a") {
    case "\n\t":
        for fi11 := 0; fi11 < 8; fi11++ {
            if false {
                v7 += (v7 / ((v7) | 1))
            }
            _ = fi11
        }
    default:
        if (-9223372036854775808 > 1) {
            fmt.Println(v7)
            var v12 float32 = (3.14159 * (1.5 * float32((0.0 * (0.1 + 0.1)))))
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: 42}
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 739
    _ = v15
    _ = v7
    return (v15 << (((v7 ^ (v7 - v7))) & 63))
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 127
    var v3 int64 = -1
    v4 := []int{9223372036854775807, -200, -2147483648, -258}
    var v5 int32 = 2147483647
    v6 := []uint32{2147483647, 497}
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int64 = v3
        v3 -= v3
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 1
    v18 = (^(v18))
    var v19 int8 = 111
    v2 += 63
    var v20 uint64 = 2
    switch "a" {
    case "":
        if len(v4) > 0 {
            v4[0] = -549
        }
    }
    switch "a" {
    case "\n\t":
        switch 1 {
        case 144:
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 int64 = (v3 ^ v3)
                for fi23 := 0; fi23 < 5; fi23++ {
                    fmt.Println(v22, v19)
                    v5 = -587
                    var v24 int8 = 1
                    var v25 int32 = -1
                    var v26 int16 = v1
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 float32 = 3.14159
                _ = fi21
                _ = v22
                _ = v27
            }
        case 42:
            if (42 != 65535) {
                switch 9223372036854775807 {
                case -147:
                    v3 -= (v3 / (((^(v3))) | 1))
                default:
                    v1 -= v1
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch ("" + "") {
                    case "a":
                        fmt.Println(v3, v5)
                    }
                    var v29 int8 = 42
                    _ = fi28
                    _ = v29
                }
            }
        }
    case "":
        v18 = (v18 >> ((1) & 63))
    case "foo bar":
        v5 = v5
    default:
        v19 -= 100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
