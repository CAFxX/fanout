package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int, p1 int64) int64 {
    v6 := []int{p0, (p0 >> ((255) & 63))}
    if len(v6) > 0 {
        v6[0] = 42
    }
    if false {
        v6 = append(v6, (p0 << ((18446744073709551615) & 63)))
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int64 = (p1 * (p1 ^ (^(p1))))
        v8 = p1
        switch "" {
        case "foo bar":
            switch "" {
            case "foo bar":
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: -2147483648, F1: (fi9 >> ((1) & 63))}
                    _ = fi9
                    _ = v10
                }
            default:
                v8 -= -9223372036854775808
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return (p1 | 84)
}

func main() {
    v1 := []uint32{2147483647, 0}
    var v2 int8 = 99
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint8 = 1
    switch "hello" {
    case "hello":
        switch (("\n\t" + ("hello" + "")) + (("hello" + ("" + ("foo bar" + ""))) + ("" + "foo bar"))) {
        case "\n\t":
            var v11 float32 = 1e10
            _ = v11
        case "foo bar":
            var v12 float64 = float64((((-2.25 / 1.5) - 3.14159) * ((((0.0 - -2.25) / (1.5 - (1.5 + (1e10 * (((0.1 * (-2.25 - -2.25)) / 1e10) - float32(3.14159)))))) + (1.5 - (-2.25 + float32((1e10 - 3.14159))))) + 0.0)))
            _ = v12
        case "a":
            if true {
                v5 += (^(v5))
                v2 += v2
            }
        default:
            if (4294967295 <= 127) {
                var v13 uint16 = 42
                var v14 uint64 = 42
                switch ("hello" + "\n\t") {
                case "foo bar":
                    v14 += 4294967295
                default:
                    v15 := GenS2{F0: 360, F1: 0, F2: (v2 ^ v2)}
                    _ = v15
                }
                _ = v13
                _ = v14
            } else {
                fmt.Println(v5, v5)
                fmt.Println(v5, v2)
            }
        }
    case "a":
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v5)
            if true {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v2 -= (v2 / (((-(v2))) | 1))
                    _ = len(v1)
                    _ = fi17
                }
                for fi18 := 0; fi18 < 17; fi18++ {
                    switch fi18 {
                    case 32768:
                        v1 = append(v1, 0)
                    }
                    v1 = append(v1, 644)
                    _ = fi18
                }
            } else {
                v5 = v5
                var v19 bool = (((float64(1.5) / (-2.25 * float64(float32(float64(0.1))))) < float64(1e10)) && ((true && true) && (!(2147483647 != -2147483648))))
                _ = v19
            }
            _ = fi16
        }
    }
    switch 4294967295 {
    case -9223372036854775808:
        v2 = (-(v2))
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v20 uint16 = 2
    if (v5 == (v5 | v5)) {
        v21 := GenS1{F0: -2147483648, F1: 0}
        var v22 uint = 255
        var v23 int8 = (-(v2))
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int = 0
    v5 = 2
    _ = len(v1)
    v25 := GenS2{F0: 2, F1: 2147483647, F2: v2}
    for fi26 := 0; fi26 < 1; fi26++ {
        v5 = 128
        _ = fi26
    }
    v5 = (v5 >> ((141) & 63))
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS3{F0: 209}
        _ = fi27
        _ = v28
    }
    if (18446744073709551615 <= 0) {
        fmt.Println(v20)
    }
    var v29 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
