package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    v1 := []uint32{2147483648, 472, 12, 2147483648}
    var v2 int64 = 745
    var v3 int8 = 127
    v4 := make(map[int]bool)
    if true {
        if false {
            switch (("" + ("" + "\n\t")) + ((("a" + "\n\t") + "hello") + "\n\t")) {
            case "a":
                if (-2.25 >= -2.25) {
                    if (!(false || (!(1 < 255)))) {
                        fmt.Println(v3)
                        delete(v4, 844)
                        if (1.5 > 1.5) {
                            v2 -= v2
                        }
                    }
                    delete(v4, -412)
                }
            default:
                v2 = v2
            }
            var v5 int = 128
            v6 := GenS1{F0: "foo bar", F1: true}
            _ = v5
            _ = v6
        }
        var v7 bool = (!(42 == 932))
        v2 = (v2 + 1)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 155
            var v10 string = "foo bar"
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = len(v4)
    }
    v2 += v2
    for fi11 := 0; fi11 < 19; fi11++ {
        v3 = 1
        _ = fi11
    }
    switch ("" + "hello") {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        if false {
            _ = len(v4)
            v12 := GenS2{F0: 0, F1: 2}
            v3 -= (v3 % ((v3) | 1))
            _ = v12
        }
    }
    if len(v1) > 0 {
        v1[0] = 492
    }
    var v13 float64 = float64(1.5)
    v13 += v13
    var v14 int16 = -1
    var v15 int32 = 0
    v13 += (((1e10 + v13) * (3.14159 / v13)) - ((0.1 / (v13 - (v13 + v13))) * v13))
    var v16 float32 = 0.0
    switch "\n\t" {
    case "foo bar":
        v17 := v4[2147483647]
        _ = v17
    case "a":
        if false {
            v18 := GenS1{F0: ("a" + ("" + "\n\t")), F1: ((false || true) || false)}
            fmt.Println(v13, v14, v15)
            _ = v18
        } else {
            fmt.Println(v3, v14, v15)
            switch -2147483648 {
            case -2147483648:
                v15 = (v15 | 0)
            case 695:
                var v19 int64 = (v2 % ((v2) | 1))
                _ = v19
            }
        }
    default:
        v2 = (v2 << ((4294967295) & 63))
    }
    v13 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
