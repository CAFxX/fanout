package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
func genFn0(p0 bool) uint8 {
    p0 = (1000 > 786)
    var v6 uint16 = 42
    if p0 {
        p0 = (v6 != (v6 & 821))
        v7 := GenS1{F0: 1}
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = 0
            for fi10 := 0; fi10 < 1; fi10++ {
                p0 = p0
                v11 := GenS1{F0: -306}
                _ = fi10
                _ = v11
            }
            var v12 float64 = (1.5 - (1e10 / -2.25))
            _ = fi8
            _ = v12
            _ = v9
        }
        _ = v7
    } else {
        if (p0 || (-32768 == 0)) {
            p0 = (-406 <= 195)
        } else {
            v13 := []int{413, 623, -1, -2147483648}
            for fi14 := 0; fi14 < 18; fi14++ {
                p0 = (("\n\t" + ("a" + "hello")) > "")
                _ = fi14
            }
            v15 := GenS1{F0: 528}
            _ = v13
            _ = v15
        }
        p0 = ((-153 > 378) && p0)
    }
    _ = p0
    _ = v6
    return 197
}
func genFn1(p0 uint64, p1 int16) uint16 {
    fmt.Println(p1)
    v16 := GenS2{F0: (p1 >> ((1) & 63)), F1: 311, F2: 2}
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(fi17, p0, fi17)
        _ = fi17
    }
    v18 := []int{128, -324, 42, 412}
    for fi19 := 0; fi19 < 9; fi19++ {
        p1 -= (-(p1))
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v18
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    var v5 int = -9223372036854775808
    switch v5 {
    case 2147483647:
        fmt.Println(v5)
    case 127:
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint32 = 0
            _ = fi20
            _ = v21
        }
    case 42:
        v5 += 9223372036854775807
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5)
        case "a":
            var v22 float64 = 1.5
            _ = v22
        case "hello":
            if ((v5 - v5) == (v5 % (((v5 + v5)) | 1))) {
                var v23 float64 = ((float64(-2.25) / 3.14159) * -2.25)
                _ = v23
            }
        }
    }
    v24 := v1["a"]
    fmt.Println(v5)
    v24 = "foo bar"
    if (true && true) {
        v24 = "foo bar"
        var v25 int64 = -318
        _ = v25
    } else {
        v26 := []int{(v5 << ((0) & 63)), v5, 2147483647}
        _ = len(v3)
        switch v5 {
        case -1:
            for fi27 := 0; fi27 < 11; fi27++ {
                delete(v2, fi27)
                _ = fi27
            }
        case -828:
            var v28 int8 = -11
            _ = v28
        case 128:
            v29 := v1[""]
            _ = v29
        default:
            var v30 uint16 = 255
            _ = v30
        }
        _ = v26
    }
    v31 := []int{v5, v5, v5, (v5 | v5)}
    var v32 uint = 651
    var v33 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
