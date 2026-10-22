package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{26, 9223372036854775807, 0}
    v2 := []int{-300, 790, 581, 9223372036854775807}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 42
    v5 := []uint32{2, 837}
    v6 := make(map[uint32]bool)
    v4 = v4
    for fi7 := 0; fi7 < 12; fi7++ {
        if ((!(false || (18446744073709551615 > 4294967295))) || (((0 != 388) && false) || false)) {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int32 = 0
                v10 := GenS1{F0: fi7, F1: 42}
                switch "foo bar" {
                case "a":
                    var v11 int8 = 25
                    _ = v11
                default:
                    var v12 uint64 = 402
                    _ = v12
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v14 uint = 777
                var v15 string = "\n\t"
                _ = fi13
                _ = v14
                _ = v15
            }
            fmt.Println(v4, v4)
        } else {
            switch "hello" {
            case "a":
                fmt.Println(fi7)
            case "hello":
                for fi16 := 0; fi16 < 19; fi16++ {
                    v4 -= v4
                    var v17 int8 = 117
                    _ = fi16
                    _ = v17
                }
            case "\n\t":
                v4 = 168
            default:
                v18 := GenS2{F0: ((1.5 * 0.0) == 0.0)}
                _ = v18
            }
            var v19 int32 = 255
            _ = v19
        }
        v4 = (^(v4))
        _ = fi7
    }
    v20 := GenS2{F0: false}
    var v21 int32 = 0
    v22 := GenS2{F0: false}
    for fi23 := 0; fi23 < 6; fi23++ {
        v4 += 2
        var v24 uint16 = 255
        _ = fi23
        _ = v24
    }
    var v25 int32 = 1
    switch ("" + "foo bar") {
    case "hello":
        var v26 float32 = 0.1
        _ = v26
    case "\n\t":
        v4 = (v4 >> ((466) & 63))
    case "foo bar":
        if false {
            fmt.Println(v4)
            v4 += (v4 - (v4 & 1))
            switch "hello" {
            case "foo bar":
                for fi27 := 0; fi27 < 6; fi27++ {
                    v21 = (v21 ^ (v21 >> ((834) & 63)))
                    _ = fi27
                }
            }
        } else {
            v28 := GenS1{F0: -411, F1: -1}
            v29 := GenS1{F0: 32768, F1: 9223372036854775807}
            _ = v28
            _ = v29
        }
    default:
        fmt.Println(v25, v25, v25)
    }
    fmt.Println(v4, v21)
    var v30 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
