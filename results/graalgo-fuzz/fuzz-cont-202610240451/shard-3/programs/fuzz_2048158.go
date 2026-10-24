package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = ("foo bar" == "a")
    v2 := []int{}
    var v3 float64 = (((1e10 / 3.14159) + 0.0) * 1.5)
    var v4 int8 = -128
    switch 32768 {
    case 4294967295:
        v5 := GenS1{F0: false, F1: 198, F2: -32768}
        _ = v5
    case 2147483647:
        var v6 uint = 320
        _ = v6
    case 32767:
        fmt.Println(v4, v1)
    default:
        _ = len(v2)
    }
    var v7 bool = (-2.25 <= -2.25)
    v2 = append(v2, 32767)
    var v8 uint = 1
    switch "" {
    case "a":
        if (0 >= 255) {
            switch ("hello" + "") {
            case "":
                switch ("" + ("\n\t" + "a")) {
                case "a":
                    fmt.Println(v7, v1, v4)
                }
            case "\n\t":
                v1 = (0 == 2147483648)
            case "foo bar":
                if (1 >= 1) {
                    v9 := GenS2{F0: 125, F1: 1, F2: 1}
                    v3 += ((((3.14159 - 1e10) - (v3 + (1.5 - float64(((float32(v3) * 0.0) * (1e10 * 0.1)))))) * ((v3 * 1e10) + v3)) + 3.14159)
                    _ = v9
                }
            default:
                if v7 {
                    v7 = (1 <= -735)
                    v7 = v1
                } else {
                    if false {
                        v8 -= (v8 / ((v8) | 1))
                        if v7 {
                            fmt.Println(v7, v3, v8)
                            v10 := GenS2{F0: 42, F1: 900, F2: 127}
                            _ = len(v2)
                            _ = v10
                        } else {
                            v2 = append(v2, 0)
                            var v11 uint64 = 4294967295
                            _ = v11
                        }
                    }
                    if (!(!(0 <= 448))) {
                        fmt.Println(v7, v7)
                    } else {
                        v8 = (v8 >> (((v8 << (((v8 | (v8 - 42))) & 63))) & 63))
                    }
                }
            }
            var v12 bool = false
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 int16 = 0
                v12 = (780 > 42)
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    case "":
        v2 = append(v2, 0)
    }
    var v15 string = "foo bar"
    v15 = (((v15 + "hello") + "foo bar") + (((v15 + v15) + v15) + "foo bar"))
    v16 := GenS1{F0: (float32((v3 / v3)) < 0.0), F1: -1, F2: 1000}
    v1 = false
    if v7 {
        if v1 {
            v3 -= float64(((3.14159 / 0.1) * 1e10))
            v4 = v4
        }
        v7 = (v4 <= v4)
        for fi17 := 0; fi17 < 19; fi17++ {
            _ = len(v2)
            v18 := GenS3{F0: 0}
            var v19 uint = v8
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v3 -= (float64(1e10) + (0.1 * 3.14159))
        var v20 uint16 = 65535
        _ = v20
    }
    if v7 {
        v7 = (uint8(v8) >= 42)
    }
    var v21 uint = v8
    fmt.Println(v8, v15, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
