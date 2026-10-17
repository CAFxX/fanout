package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint16) int {
    switch "a" {
    case "":
        fmt.Println(p0)
    case "hello":
        var v5 int8 = 1
        _ = v5
    default:
        v6 := []int{}
        _ = v6
    }
    var v7 bool = (!(!true))
    v8 := []int{}
    v8 = append(v8, 32767)
    v9 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 bool) uint64 {
    if p0 {
        var v10 uint = 42
        _ = v10
    } else {
        if (69 < 19) {
            switch genFn0(354) {
            case -1:
                p0 = p0
            case 0:
                for fi11 := 0; fi11 < 13; fi11++ {
                    v12 := []int{(fi11 / ((fi11) | 1)), (fi11 & fi11)}
                    p0 = (true || (910 > 1))
                    p0 = ((p0 || (978 < 4294967295)) && (-1 > 1))
                    _ = fi11
                    _ = v12
                }
            default:
                v13 := []int{42, 127}
                _ = v13
            }
            p0 = (508 <= 755)
            fmt.Println(p0)
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int = fi14
                v16 := []int{fi14, (v15 - (v15 - v15))}
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS2{F0: 896, F1: -1}
            _ = v17
        }
        switch 2147483647 {
        case -1:
            var v18 int32 = 959
            _ = v18
        case -488:
            p0 = (2 != 42)
        }
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := []int{9223372036854775807, 772, 42, 32767}
    var v2 int32 = -398
    v3 := []int{42, 0}
    v4 := make(map[string]string)
    if false {
        var v19 int16 = 1
        _ = v19
    }
    v2 -= (v2 << ((42) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    for fi20 := 0; fi20 < 15; fi20++ {
        delete(v4, "hello")
        var v21 uint8 = 11
        _ = fi20
        _ = v21
    }
    var v22 int64 = 240
    switch "hello" {
    case "hello":
        if (2147483647 == 1) {
            fmt.Println(v22, v22)
        }
    case "":
        switch "\n\t" {
        case "hello":
            switch -2147483648 {
            case 1:
                var v23 string = "foo bar"
                _ = v23
            default:
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v22)
                    _ = fi24
                }
            }
        case "foo bar":
            v25 := GenS3{F0: ((0.0 / 0.1) == -2.25)}
            _ = v25
        }
    case "\n\t":
        v22 = (v22 * 0)
    }
    switch 42 {
    case 32768:
        switch "foo bar" {
        case "a":
            v2 = v2
        case "":
            for fi26 := 0; fi26 < 2; fi26++ {
                switch "" {
                case "a":
                    v22 = (v22 & (v22 + v22))
                }
                v27 := v4["\n\t"]
                var v28 uint = 464
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    case -416:
        var v29 float64 = (((3.14159 / (-2.25 + 1e10)) + ((0.0 + ((((float64(((0.0 + float32(0.0)) + (1e10 / 1.0))) + 0.0) - (-2.25 / 1.0)) - 1.5) / float64((float32(3.14159) / -2.25)))) * ((1.5 / 1.5) * 0.0))) * -2.25)
        _ = v29
    }
    v2 = v2
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS3{F0: (!(genFn0(65535) < fi30))}
        _ = fi30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
