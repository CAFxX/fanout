package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }

func main() {
    var v1 uint16 = 0
    var v2 uint = 456
    var v3 int = -141
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float64 = (1e10 - ((1e10 + 1.5) * 1.5))
        v7 := []int{(v3 | (fi5 >> ((v2) & 63))), v3, 182}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (0.1 != 1e10) {
        fmt.Println(v1)
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{42, -9223372036854775808, (fi9 + 2147483647), -1}
                var v11 int = -996
                v2 -= v2
                _ = fi9
                _ = v10
                _ = v11
            }
            v4[(fi8 ^ fi8)] = "foo bar"
            var v12 string = ("hello" + ("\n\t" + "foo bar"))
            _ = fi8
            _ = v12
        }
    }
    v13 := GenS1{F0: 42, F1: 9223372036854775807}
    v14 := []int{(-(v3)), v3}
    var v15 uint32 = 4294967295
    _ = len(v4)
    v16 := GenS1{F0: 32767, F1: -9223372036854775808}
    if false {
        var v17 string = ("hello" + "hello")
        v18 := GenS2{F0: -2147483648, F1: -1, F2: -18}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (false && false) {
            if (!true) {
                v14 = append(v14, 386)
            }
            fmt.Println(v3, v2, fi19)
        }
        _ = fi19
    }
    _ = len(v4)
    switch v3 {
    case 4294967295:
        if ((v15 * (v15 * (v15 & (v15 ^ uint32(v2))))) < 77) {
            switch "" {
            case "a":
                var v20 uint64 = 2
                _ = v20
            default:
                if (v15 > (^(v15))) {
                    var v21 float32 = ((1e10 / -2.25) * 0.0)
                    v3 = v3
                    _ = v21
                } else {
                    switch "a" {
                    case "foo bar":
                        _ = len(v4)
                    case "hello":
                        v2 += v2
                    }
                }
            }
            fmt.Println(v2, v3)
        } else {
            v1 -= v1
        }
    }
    if false {
        v2 = uint(v15)
        v22 := GenS1{F0: -1, F1: 42}
        if len(v14) > 0 {
            v14[0] = 9223372036854775807
        }
        _ = v22
    }
    v15 = 2147483648
    var v23 float32 = (3.14159 * 0.1)
    v24 := GenS3{F0: 2, F1: v1, F2: 2}
    v3 -= v3
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v23, v2)
    case -9223372036854775808:
        var v25 int = -411
        _ = v25
    case 127:
        delete(v4, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
