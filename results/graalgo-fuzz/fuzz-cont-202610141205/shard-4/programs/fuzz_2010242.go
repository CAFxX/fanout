package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 uint16) uint16 {
    var v4 float64 = 3.14159
    switch ("a" + (("foo bar" + "a") + "a")) {
    case "hello":
        fmt.Println(v4, p0)
    case "foo bar":
        if ((2147483647 >= 69) && ("" < ("hello" + "foo bar"))) {
            for fi5 := 0; fi5 < 5; fi5++ {
                switch -9223372036854775808 {
                case 1:
                    for fi6 := 0; fi6 < 4; fi6++ {
                        p0 -= (p0 + 0)
                        switch "a" {
                        case "hello":
                            v7 := []int{fi5}
                            _ = v7
                        case "\n\t":
                            var v8 int = fi5
                            _ = v8
                        }
                        var v9 int16 = -24
                        var v10 int32 = 1
                        _ = fi6
                        _ = v10
                        _ = v9
                    }
                case -887:
                    v11 := GenS1{F0: -292, F1: (!(-2.25 != (0.0 / 1e10))), F2: 32767}
                    _ = v11
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    p0 -= (^(p0))
                    var v13 bool = true
                    _ = fi12
                    _ = v13
                }
                _ = fi5
            }
            v14 := []int{128, 9223372036854775807}
            _ = v14
        } else {
            fmt.Println(p0, v4)
            v15 := GenS1{F0: 14, F1: (128 > 255), F2: 1000}
            _ = v15
        }
    case "\n\t":
        v4 = (((v4 * float64(-2.25)) + ((1e10 / (1e10 / ((v4 - v4) / v4))) / (v4 - v4))) / (1.5 + v4))
    default:
        v16 := GenS2{F0: -1, F1: 1}
        _ = v16
    }
    switch "hello" {
    case "hello":
        v17 := GenS3{F0: 109, F1: 18446744073709551615}
        _ = v17
    case "a":
        v4 = float64(0.0)
    default:
        fmt.Println(v4, v4)
    }
    var v18 string = "foo bar"
    _ = p0
    _ = v18
    _ = v4
    return (p0 / ((2) | 1))
}

func main() {
    v1 := []int{1, -997, 128, 42}
    var v2 int64 = -978
    var v3 int32 = -779
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint64 = 1
            _ = fi19
            _ = v20
        }
    default:
        var v21 uint8 = 0
        _ = v21
    }
    fmt.Println(v3)
    var v22 bool = false
    v23 := GenS3{F0: 0, F1: 8}
    v22 = false
    var v24 float32 = -2.25
    fmt.Println(v22)
    fmt.Println(v24, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
