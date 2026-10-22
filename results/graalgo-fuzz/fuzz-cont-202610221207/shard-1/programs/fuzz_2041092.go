package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    var v3 bool = true
    var v4 int8 = -114
    var v5 uint32 = 1
    fmt.Println(v3, v3)
    var v6 uint32 = v5
    if v3 {
        for fi7 := 0; fi7 < 17; fi7++ {
            v1 = append(v1, v5)
            var v8 string = ""
            _ = fi7
            _ = v8
        }
        switch (("hello" + (("\n\t" + (("" + "foo bar") + ("" + ("a" + (("a" + "a") + "\n\t"))))) + "\n\t")) + "\n\t") {
        case "hello":
            v9 := GenS1{F0: -243, F1: 0}
            _ = v9
        case "":
            v10 := GenS1{F0: 822, F1: 1}
            _ = v10
        default:
            var v11 float64 = 1e10
            _ = v11
        }
        switch "" {
        case "":
            v6 -= (v6 << ((18446744073709551615) & 63))
        case "a":
            if (848 != 127) {
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint = 211
                    var v14 int64 = -9223372036854775808
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case "hello":
            v5 += (^(v6))
        default:
            fmt.Println(v6, v3, v3)
        }
    }
    var v15 uint = 42
    if v3 {
        var v16 int8 = (v4 & 42)
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                v15 = v15
                _ = fi18
            }
            switch "" {
            case "a":
                switch 207 {
                case 1:
                    var v19 string = (("a" + ("a" + ("foo bar" + (("hello" + "foo bar") + ("\n\t" + "a"))))) + "foo bar")
                    _ = v19
                case 2147483647:
                    var v20 uint = (v15 - uint(v6))
                    _ = v20
                case -1:
                    delete(v2, "a")
                default:
                    if v3 {
                        var v21 int32 = 1
                        var v22 int8 = v4
                        _ = v21
                        _ = v22
                    } else {
                        v3 = false
                        v15 += v15
                        var v23 int = (fi17 % ((fi17) | 1))
                        _ = v23
                    }
                }
            default:
                var v24 int8 = v4
                _ = v24
            }
            var v25 int8 = -128
            _ = fi17
            _ = v25
        }
        v6 = uint32(v15)
        _ = v16
    }
    var v26 uint = (^(v15))
    var v27 float64 = (((1e10 + float64(-2.25)) / ((((-2.25 / 1.0) * 0.0) * (-2.25 / 1.5)) + (1.5 - ((0.0 * 1e10) + (0.1 * ((0.0 * float64((1e10 * 1e10))) + 1e10)))))) / ((0.1 * -2.25) * 0.1))
    var v28 uint64 = 335
    v15 -= uint(v5)
    var v29 int16 = 985
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
