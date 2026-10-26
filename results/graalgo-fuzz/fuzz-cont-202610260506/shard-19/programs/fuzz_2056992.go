package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 65535
    v3 := make(map[int]string)
    var v4 uint64 = 0
    v5 := []int{618, 0, 178}
    if (!(true && (0.1 <= -2.25))) {
        v6 := GenS1{F0: ((v4 << (((v2 * v2)) & 63)) > v4)}
        var v7 int64 = -325
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v7, v7, v2)
        case "hello":
            var v8 float64 = (1e10 * (((3.14159 / 1e10) - -2.25) / 1.0))
            _ = v8
        default:
            var v9 uint8 = 2
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        fmt.Println(v4, v4)
        v2 -= 0
        v10 := v3[704]
        _ = v10
    }
    v4 += v4
    switch "a" {
    case "hello":
        fmt.Println(v4, v2)
    case "":
        fmt.Println(v2, v2)
    default:
        v4 = (v4 | (v4 % ((209) | 1)))
    }
    v2 -= v2
    v2 -= v2
    var v11 uint16 = 65535
    if len(v5) > 0 {
        v5[0] = -200
    }
    if (1 == 42) {
        v4 -= (v4 | 42)
        var v12 uint = (v2 * v2)
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS1{F0: (-51 >= 42)}
                switch fi14 {
                case 2147483647:
                    var v16 bool = ((2147483647 == 126) || true)
                    _ = v16
                }
                var v17 int64 = 9223372036854775807
                _ = fi14
                _ = v15
                _ = v17
            }
            var v18 uint32 = 65535
            _ = fi13
            _ = v18
        }
        _ = v12
    }
    if false {
        v19 := GenS1{F0: (!(false && ((^(v2)) == v2)))}
        var v20 bool = (!(((!true) && false) && (!true)))
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            v11 = v11
            v11 = v11
            if false {
                if false {
                    var v22 uint = (v2 + v2)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v22
                }
                var v23 uint = 255
                _ = v23
            } else {
                switch ("" + ("a" + ("hello" + "hello"))) {
                case "\n\t":
                    var v24 uint16 = 1
                    _ = v24
                case "":
                    if true {
                        v11 -= v11
                    }
                case "foo bar":
                    var v25 uint32 = 2147483647
                    _ = v25
                }
            }
            _ = fi21
        }
    }
    if (false && false) {
        _ = len(v5)
    } else {
        v4 -= (v4 / ((v4) | 1))
        if len(v5) > 0 {
            v5[0] = 128
        }
    }
    v26 := GenS1{F0: false}
    v27 := GenS1{F0: false}
    _ = len(v1)
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
