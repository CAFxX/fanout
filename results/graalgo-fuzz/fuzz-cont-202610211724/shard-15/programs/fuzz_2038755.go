package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 4294967295
    v3 := []uint32{}
    var v4 string = "foo bar"
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 19; fi6++ {
        delete(v1, v4)
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        _ = len(v1)
        _ = fi7
    }
    v2 = v2
    fmt.Println(v4, v2)
    switch (v4 + (((v4 + ("hello" + v4)) + ("a" + (("foo bar" + (v4 + v4)) + ("" + v4)))) + ("\n\t" + ((v4 + v4) + ("a" + v4))))) {
    case "":
        switch 128 {
        case 128:
            if (65535 < 2) {
                var v8 string = "\n\t"
                var v9 uint32 = 1
                _ = v8
                _ = v9
            } else {
                if (-1 <= 0) {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v11 := GenS1{F0: -2147483648, F1: -1, F2: 0}
                        var v12 int = -2147483648
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    if (-2.25 < 0.1) {
                        var v13 int64 = 42
                        _ = v13
                    } else {
                        v2 = (v2 / ((0) | 1))
                        v4 = v4
                        v4 = v4
                    }
                } else {
                    if (!(32768 == 42)) {
                        fmt.Println(v4, v4)
                    }
                    v2 -= (^(v2))
                }
                v14 := GenS2{F0: 497, F1: 240}
                _ = len(v1)
                _ = v14
            }
        }
    case "foo bar":
        var v15 uint = v2
        _ = v15
    default:
        if (255 < 42) {
            var v16 int8 = 1
            if false {
                v17 := GenS2{F0: -1, F1: 181}
                _ = v17
            }
            _ = v16
        } else {
            var v18 int8 = -128
            _ = v18
        }
    }
    var v19 int8 = -128
    v2 = 0
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := v5[2147483648]
        _ = fi20
        _ = v21
    }
    v19 = (^(v19))
    var v22 uint32 = 1
    v23 := GenS2{F0: 127, F1: 68}
    fmt.Println(v4, v22)
    fmt.Println(v22, v2)
    switch (((v4 + (v4 + v4)) + "hello") + v4) {
    case "hello":
        v19 -= v19
    case "\n\t":
        var v24 uint64 = 870
        _ = v24
    case "foo bar":
        var v25 float64 = (1e10 / 3.14159)
        _ = v25
    }
    v3 = append(v3, (v22 / ((v22) | 1)))
    v4 = v4
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
