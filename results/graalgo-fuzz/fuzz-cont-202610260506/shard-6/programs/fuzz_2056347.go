package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []uint32{4294967295}
    var v2 uint64 = 42
    var v3 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi4 := 0; fi4 < 12; fi4++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = fi4
    }
    v5 := GenS1{F0: 9223372036854775807, F1: 65535, F2: (!(false || ((241 >= (v2 % (((v2 ^ (v2 ^ (v2 & (v2 * v2))))) | 1))) && false)))}
    v3 -= (v3 * (((v3 * v3) * (float64(-2.25) + (v3 * v3))) / v3))
    v3 += float64(-2.25)
    fmt.Println(v3, v3)
    v2 = 4294967295
    v6 := GenS1{F0: -9223372036854775808, F1: 2, F2: (1 >= 1)}
    var v7 float64 = v3
    _ = len(v1)
    switch 32768 {
    case 128:
        v8 := GenS2{F0: 581, F1: 445}
        _ = v8
    case 1:
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi9
        }
    case 127:
        if false {
            var v10 int32 = 42
            switch (("\n\t" + "") + "\n\t") {
            case "\n\t":
                v3 = float64(((((0.1 * (3.14159 * float32((v3 + 0.0)))) * ((float32(v3) - (0.1 * (((1e10 - (1e10 * (0.1 - -2.25))) - ((3.14159 - (1e10 * -2.25)) + 1e10)) / 0.1))) + -2.25)) / float32(v7)) - ((1e10 - 1e10) + (float32((0.0 + float64(0.1))) + -2.25))))
            case "foo bar":
                if false {
                    var v11 string = "hello"
                    _ = v11
                }
            default:
                if true {
                    v12 := GenS2{F0: 615, F1: 128}
                    _ = v12
                }
            }
            _ = v10
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                v3 -= v7
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi13
            }
            v7 += ((float64(1e10) - float64(-2.25)) - -2.25)
            var v14 uint = 355
            _ = v14
        }
    }
    var v15 uint16 = 969
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
