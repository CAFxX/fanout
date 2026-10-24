package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 uint64 = 673
    v5 := []string{}
    v6 := make(map[string]int)
    fmt.Println(v2, v2)
    fmt.Println(v4)
    var v7 uint64 = v4
    v8 := GenS1{F0: 0, F1: -749, F2: 68}
    var v9 uint8 = 128
    v2 -= v2
    _ = len(v5)
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "foo bar" {
        case "foo bar":
            for fi11 := 0; fi11 < 19; fi11++ {
                v4 += (v4 / (((v7 * (v4 | (v4 | (v7 + 303))))) | 1))
                fmt.Println(v7, fi10, fi10)
                _ = fi11
            }
        default:
            if true {
                var v12 uint = 4294967295
                _ = v12
            }
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: 0, F1: int(v9), F2: (fi13 >> ((uint(v9)) & 63))}
        var v15 int64 = -459
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 -= (v4 % (((v4 + (v7 << ((4294967295) & 63)))) | 1))
    v16 := GenS2{F0: 2, F1: int32(v9), F2: false}
    _ = len(v6)
    if (!(!(v7 >= v4))) {
        v2 = ((v2 / 1e10) - -2.25)
        switch "\n\t" {
        case "a":
            var v17 float64 = float64(v2)
            _ = v17
        case "foo bar":
            var v18 string = (("" + ("foo bar" + "\n\t")) + (((("a" + ("hello" + "\n\t")) + "hello") + "") + "hello"))
            _ = v18
        }
        if false {
            v7 -= (v7 + (^(v4)))
            if (!false) {
                v9 = 255
                var v19 float64 = ((float64((v2 - ((v2 / 1.0) + v2))) / (-2.25 * (-2.25 * -2.25))) / 1.0)
                _ = v19
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 float32 = (v2 / 1.0)
                _ = fi20
                _ = v21
            }
        } else {
            v4 = v7
            var v22 bool = false
            _ = v22
        }
    }
    var v23 uint = 18446744073709551615
    var v24 int = 9223372036854775807
    var v25 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
