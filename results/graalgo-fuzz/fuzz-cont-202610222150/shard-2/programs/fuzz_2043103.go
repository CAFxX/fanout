package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64) int64 {
    var v4 bool = true
    if v4 {
        var v5 uint8 = 0
        switch 127 {
        case 9223372036854775807:
            v6 := GenS1{F0: -32768}
            _ = v6
        default:
            switch ("" + ((("a" + ("\n\t" + (("hello" + (("hello" + "foo bar") + "a")) + "a"))) + "foo bar") + "hello")) {
            case "a":
                var v7 uint16 = 2
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 uint64 = 18446744073709551615
                    v5 = v5
                    _ = fi8
                    _ = v9
                }
            case "":
                var v10 int = -9223372036854775808
                _ = v10
            }
        }
        _ = v5
    } else {
        var v11 uint64 = 54
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint = 843
            _ = fi12
            _ = v13
        }
        switch "foo bar" {
        case "":
            v11 = (v11 + v11)
        default:
            v11 += v11
        }
        _ = v11
    }
    p0 = p0
    if v4 {
        var v14 uint = 255
        p0 -= (p0 & p0)
        v15 := GenS1{F0: -184}
        _ = v14
        _ = v15
    }
    v4 = true
    _ = p0
    _ = v4
    return (p0 + p0)
}
func genFn1(p0 int32) int8 {
    p0 += -249
    switch "\n\t" {
    case "foo bar":
        p0 = p0
    case "hello":
        var v16 int64 = 248
        _ = v16
    default:
        for fi17 := 0; fi17 < 2; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                if false {
                    fmt.Println(p0, fi18, p0)
                    v19 := []int{(fi18 | fi17), fi17, fi18}
                    _ = v19
                }
                var v20 string = ((("\n\t" + "foo bar") + "foo bar") + "\n\t")
                _ = fi18
                _ = v20
            }
            var v21 uint = 65535
            _ = fi17
            _ = v21
        }
    }
    p0 += (p0 >> ((0) & 63))
    _ = p0
    return -1
}

func main() {
    var v1 bool = true
    var v2 float64 = 0.0
    var v3 bool = ((0.0 > (1.5 * (0.0 * ((1.5 + (1e10 / float32((0.0 + 1.5)))) + 1e10)))) || (0 != 1))
    v22 := []int{-2147483648, 9223372036854775807}
    if len(v22) > 0 {
        v22[0] = 9223372036854775807
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(v3, v3, fi23)
        _ = fi23
    }
    v1 = ((v2 >= (v2 * (v2 * (0.1 + (0.1 / v2))))) && true)
    v24 := GenS2{F0: 100, F1: 42, F2: 1}
    var v25 int8 = 0
    v1 = v1
    var v26 int32 = -147
    switch 2147483647 {
    case 1:
        if (-985 > 127) {
            fmt.Println(v26, v3, v25)
            fmt.Println(v1, v2)
        }
    }
    var v27 uint8 = 1
    if v3 {
        v22 = append(v22, -308)
        v28 := GenS2{F0: (^(v25)), F1: 214, F2: -128}
        v29 := GenS3{F0: 127, F1: (v26 * -577)}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
