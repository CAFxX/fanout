package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool) bool {
    p0 = false
    if p0 {
        if p0 {
            p0 = (42 > -1)
        } else {
            p0 = (!(1 < 0))
        }
        p0 = (127 < -2147483648)
        v4 := GenS1{F0: 1000}
        _ = v4
    } else {
        v5 := []int{32767, -705, 631}
        v6 := GenS1{F0: 0}
        var v7 int8 = -128
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            p0 = ((p0 && p0) && (((255 != 1) || false) || true))
            _ = fi9
        }
        _ = fi8
    }
    if (!(-9223372036854775808 > -457)) {
        v10 := []int{9223372036854775807, -41}
        _ = v10
    } else {
        switch 463 {
        case 32768:
            fmt.Println(p0)
        default:
            p0 = true
        }
        var v11 uint = 18446744073709551615
        v11 = v11
        _ = v11
    }
    _ = p0
    return true
}
func genFn1(p0 string) uint8 {
    var v12 bool = (1 > 65535)
    var v13 uint32 = 2147483647
    _ = p0
    _ = v12
    _ = v13
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 int8 = 100
    var v14 int16 = -1
    v15 := []int{32767, 2147483647, 0, 509}
    v14 += v14
    switch ("hello" + "foo bar") {
    case "foo bar":
        v16 := GenS2{F0: 385, F1: genFn1(("\n\t" + "a"))}
        _ = v16
    default:
        var v17 bool = false
        _ = v17
    }
    v14 = 375
    v18 := GenS1{F0: -1}
    v3 -= v3
    if len(v15) > 0 {
        v15[0] = 617
    }
    var v19 uint32 = 65535
    if true {
        var v20 uint8 = 127
        var v21 string = ("a" + "")
        v22 := GenS1{F0: (v14 | (v14 & -1))}
        _ = v20
        _ = v21
        _ = v22
    }
    switch "" {
    case "":
        if true {
            var v23 uint32 = v19
            _ = v23
        } else {
            _ = len(v2)
        }
    case "\n\t":
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float64 = (3.14159 * 1e10)
            var v26 string = ("foo bar" + "\n\t")
            var v27 int32 = 0
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case "hello":
        v28 := GenS2{F0: 2, F1: 0}
        _ = v28
    default:
        if (true || (false && (!(2 == 2147483647)))) {
            if true {
                v29 := GenS1{F0: (v14 & v14)}
                var v30 int8 = (-(v3))
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := GenS2{F0: 255, F1: 127}
                    _ = fi31
                    _ = v32
                }
                v15 = append(v15, 32767)
                var v33 uint = 2
                _ = v33
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
