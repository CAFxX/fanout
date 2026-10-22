package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 bool) uint8 {
    var v4 uint8 = 127
    var v5 uint = (p0 % ((p0) | 1))
    v6 := []int{9223372036854775807, -448, 0, 671}
    if (32767 <= -99) {
        var v7 uint16 = 597
        v8 := GenS1{F0: 42}
        _ = v7
        _ = v8
    }
    if (uint32(v4) > 42) {
        v9 := GenS1{F0: -32768}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (v4 ^ (^(v4)))
}
func genFn1(p0 int16, p1 uint8) uint64 {
    for fi10 := 0; fi10 < 11; fi10++ {
        if false {
            fmt.Println(p1, p1)
            var v11 float64 = ((0.0 * 1.5) + 3.14159)
            v12 := GenS2{F0: 255, F1: 4294967295}
            _ = v11
            _ = v12
        }
        var v13 uint64 = uint64(p1)
        _ = fi10
        _ = v13
    }
    var v14 uint8 = 206
    var v15 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 bool = (775 < 194)
    if true {
        switch "" {
        case "foo bar":
            if (!true) {
                if v3 {
                    fmt.Println(v3)
                } else {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        _ = len(v2)
                        _ = fi16
                    }
                    if (-180 == 32767) {
                        var v17 float32 = (((1e10 * 3.14159) * 3.14159) - 1e10)
                        _ = len(v1)
                        _ = v17
                    }
                }
            }
        case "hello":
            var v18 uint16 = 32767
            _ = v18
        }
    } else {
        switch ("\n\t" + "foo bar") {
        case "":
            var v19 float64 = 3.14159
            _ = v19
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v3 = v3
        if ((float32(1e10) + ((0.1 / 1e10) / (((float32(0.0) - -2.25) / float32(0.1)) + (1e10 - float32(1e10))))) != 0.0) {
            var v20 int16 = -814
            v21 := GenS3{F0: 26}
            v1 = append(v1, 2147483647)
            _ = v20
            _ = v21
        }
    }
    var v22 bool = (false || (!(42 > 9223372036854775807)))
    var v23 int64 = 0
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int = (fi24 & fi24)
        if v3 {
            v22 = (42 > 1000)
        }
        _ = fi24
        _ = v25
    }
    v22 = (1 != 18446744073709551615)
    var v26 string = "a"
    fmt.Println(v22, v22)
    switch 4294967295 {
    case 1:
        var v27 float64 = float64((1e10 / (0.1 / 1.0)))
        _ = v27
    case -374:
        if v22 {
            v28 := GenS1{F0: 1000}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (v3 && v3) {
                v3 = (32767 >= 82)
            }
            _ = v28
        } else {
            var v29 uint32 = 65535
            switch "\n\t" {
            case "a":
                v23 = v23
            case "":
                fmt.Println(v23)
            default:
                fmt.Println(v23)
            }
            if len(v1) > 0 {
                v1[0] = -858
            }
            _ = v29
        }
    }
    fmt.Println(v23, v26, v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
