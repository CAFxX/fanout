package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32 }
func genFn0() int8 {
    var v6 int32 = -143
    v6 = (v6 >> ((18446744073709551615) & 63))
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "\n\t" {
        case "hello":
            v8 := []int{}
            _ = v8
        case "foo bar":
            var v9 int32 = (-(v6))
            _ = v9
        }
        switch -869 {
        case -492:
            if (true || true) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(v6, fi7)
                    fmt.Println(fi7, fi7, fi10)
                    _ = fi10
                }
                var v11 string = "hello"
                _ = v11
            }
        default:
            var v12 int32 = v6
            _ = v12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            if false {
                switch fi7 {
                case 127:
                    v6 = v6
                case 810:
                    fmt.Println(fi13)
                case 489:
                    var v14 int32 = v6
                    _ = v14
                default:
                    v15 := []int{(fi13 - 9223372036854775807), (fi13 * 42), fi13}
                    _ = v15
                }
            } else {
                var v16 float64 = (1e10 * float64(-2.25))
                _ = v16
            }
            _ = fi13
        }
        _ = fi7
    }
    var v17 int32 = v6
    _ = v17
    _ = v6
    return -128
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -32768
    v3 := make(map[int]bool)
    v4 := []int{-810, -1, 32767, 32767}
    var v5 int16 = v2
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int16 = v5
        var v20 bool = true
        _ = fi18
        _ = v19
        _ = v20
    }
    v5 += v2
    v21 := GenS1{F0: 1, F1: (false || true)}
    v2 = (v2 >> ((18446744073709551615) & 63))
    var v22 int32 = -2147483648
    _ = len(v4)
    var v23 string = "foo bar"
    var v24 bool = false
    for fi25 := 0; fi25 < 13; fi25++ {
        v22 = (v22 ^ v22)
        if v24 {
            v5 = (v2 / (((-(v2))) | 1))
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 uint16 = 0
                _ = fi26
                _ = v27
            }
        }
        if (!v24) {
            if false {
                var v28 int64 = 0
                _ = v28
            }
            v24 = v24
            v29 := GenS2{F0: 679}
            _ = v29
        }
        _ = fi25
    }
    v30 := GenS2{F0: 724}
    var v31 uint = 2
    var v32 string = (v23 + ((("hello" + ("" + v23)) + v23) + v23))
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
