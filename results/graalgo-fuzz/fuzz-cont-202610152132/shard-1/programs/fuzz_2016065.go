package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 int8, p2 uint8) int64 {
    v7 := []int{}
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "a":
        p2 -= p2
    case "\n\t":
        p2 -= p2
    case "foo bar":
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 689
    var v3 float64 = 1.5
    v4 := []string{"", "a", "a", "a"}
    var v5 int16 = 32767
    var v6 uint16 = 2
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            _ = len(v4)
            v3 += v3
            _ = fi9
        }
        if true {
            switch "hello" {
            case "hello":
                fmt.Println(v2)
            case "":
                v3 = v3
            default:
                v10 := GenS1{F0: v5, F1: 2}
                _ = v10
            }
            v11 := GenS1{F0: 0, F1: 128}
            _ = v11
        } else {
            var v12 int32 = 645
            var v13 uint32 = 964
            var v14 uint = 2
            _ = v12
            _ = v13
            _ = v14
        }
        v6 += v6
        _ = fi8
    }
    if (977 == 1) {
        switch -622 {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v15 int32 = -263
            _ = v15
        }
    }
    if (2 <= uint(v6)) {
        v16 := GenS1{F0: v5, F1: 253}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int64 = 499
            var v19 int = 2147483647
            v3 = 1e10
            _ = fi17
            _ = v18
            _ = v19
        }
        if (false && (!true)) {
            switch "a" {
            case "hello":
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            case "":
                if true {
                    var v20 int64 = 9223372036854775807
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        fmt.Println(v3, v6, v6)
                        _ = fi21
                    }
                }
            case "a":
                var v22 bool = (true && true)
                _ = v22
            }
            switch ("" + "\n\t") {
            case "foo bar":
                v23 := GenS2{F0: 127, F1: 4294967295, F2: 1}
                _ = v23
            }
            v24 := GenS3{F0: false, F1: 0, F2: 232}
            _ = v24
        }
    }
    switch 679 {
    case 32768:
        var v25 int = 4294967295
        _ = v25
    }
    v5 = 1000
    for fi26 := 0; fi26 < 15; fi26++ {
        if false {
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(v6)
                _ = fi27
            }
            v3 += (v3 / 3.14159)
        } else {
            v3 = (v3 / ((float64(0.1) / v3) - v3))
            var v28 float32 = -2.25
            _ = v28
        }
        v3 += v3
        _ = fi26
    }
    v2 += v2
    var v29 uint64 = 1
    fmt.Println(v3, v29, v3)
    v30 := GenS1{F0: 169, F1: 127}
    switch -1 {
    case 128:
        v29 = v29
    case 2147483647:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v31 := GenS1{F0: 32767, F1: 192}
        _ = v31
    }
    var v32 string = "foo bar"
    v33 := GenS2{F0: 42, F1: 18446744073709551615, F2: (v29 >> ((65535) & 63))}
    fmt.Println(v2, v3)
    var v34 float64 = 1e10
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
