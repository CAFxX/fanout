package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8) string {
    var v5 int8 = 100
    switch 0 {
    case 2147483647:
        var v6 uint64 = 18446744073709551615
        _ = v6
    case -9223372036854775808:
        fmt.Println(v5, p0)
    case 127:
        var v7 uint32 = 625
        _ = v7
    default:
        v8 := GenS1{F0: -1, F1: 2}
        _ = v8
    }
    _ = p0
    _ = v5
    return "hello"
}

func main() {
    var v1 int8 = -3
    var v2 int32 = 2147483647
    var v3 int16 = 1000
    v4 := make(map[int]string)
    if ((v1 ^ v1) < (v1 & (^(v1)))) {
        switch 9223372036854775807 {
        case 32767:
            v9 := []int{4294967295, 225, -2147483648, -17}
            _ = v9
        }
        var v10 bool = ((105 > 1) && true)
        _ = v10
    }
    v11 := []int{}
    v3 = v3
    switch "a" {
    case "foo bar":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float32 = 3.14159
            _ = fi12
            _ = v13
        }
    case "":
        if (false && true) {
            v1 -= -128
            v1 += 100
            fmt.Println(v2)
        }
    }
    switch "" {
    case "":
        var v14 float32 = 1.5
        _ = v14
    case "\n\t":
        v3 = v3
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint8 = 128
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 float32 = 0.1
        v18 -= v18
        fmt.Println(v1)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v2 = 403
        v20 := GenS1{F0: 42, F1: 1}
        v3 -= v3
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v2)
        _ = fi21
    }
    v1 += -1
    v1 = (v1 ^ (v1 - (-(v1))))
    var v22 int32 = v2
    if (42 < 1) {
        for fi23 := 0; fi23 < 3; fi23++ {
            v1 += (-(v1))
            fmt.Println(fi23)
            _ = fi23
        }
    }
    v3 = -32768
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS1{F0: 32767, F1: 0}
            v3 += -128
            var v27 string = ("\n\t" + "")
            fmt.Println(fi25, v3)
            if (55 != 221) {
                switch ("foo bar" + v27) {
                case "\n\t":
                    var v28 float32 = (0.0 * 0.0)
                    _ = v28
                }
                _ = len(v11)
                var v29 int32 = 64
                v2 = (v22 * (v22 & v29))
                v30 := GenS2{F0: (v1 % ((-6) | 1)), F1: 42}
                if (-895 <= -106) {
                    var v31 uint32 = 4294967295
                    _ = v31
                }
                _ = v29
                _ = v30
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        var v32 uint64 = 18446744073709551615
        _ = fi24
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        _ = len(v4)
        switch fi33 {
        case 32768:
            v34 := GenS2{F0: v1, F1: -584}
            _ = v34
        case -9223372036854775808:
            if true {
                v1 = (v1 & (v1 & (v1 / ((int8(v3)) | 1))))
            }
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
