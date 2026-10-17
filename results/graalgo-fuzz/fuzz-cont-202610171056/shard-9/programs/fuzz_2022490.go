package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool; F1 int16 }
func genFn0(p0 int, p1 int16) bool {
    fmt.Println(p1, p1)
    v4 := []int{(-(p0))}
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return true
}

func main() {
    var v1 int = 1
    var v2 int8 = 1
    v3 := make(map[int]string)
    v5 := []int{v1}
    var v6 uint = 1
    var v7 int = (-(v1))
    for fi8 := 0; fi8 < 10; fi8++ {
        v7 -= 300
        var v9 uint16 = 32767
        _ = fi8
        _ = v9
    }
    fmt.Println(v7, v7, v2)
    v6 -= 0
    var v10 int = v7
    if false {
        if true {
            for fi11 := 0; fi11 < 4; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v10 = 42
                    v13 := GenS1{F0: 1000, F1: v1}
                    _ = fi12
                    _ = v13
                }
                v10 -= fi11
                _ = fi11
            }
            var v14 float64 = (0.1 * (0.0 * 1.5))
            if false {
                fmt.Println(v7, v2)
            }
            _ = v14
        } else {
            if len(v5) > 0 {
                v5[0] = v1
            }
            v1 = -304
        }
        _ = len(v5)
        var v15 int64 = -9223372036854775808
        _ = v15
    }
    var v16 int8 = (v2 | (v2 % (((v2 % ((v2) | 1))) | 1)))
    switch "" {
    case "hello":
        var v17 float32 = (((3.14159 / -2.25) / (0.0 - -2.25)) * -2.25)
        _ = v17
    case "":
        switch "hello" {
        case "a":
            v2 -= v2
        case "hello":
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS2{F0: (^(v2)), F1: 1000}
                var v20 uint32 = 2147483648
                _ = len(v5)
                _ = fi18
                _ = v19
                _ = v20
            }
        case "\n\t":
            for fi21 := 0; fi21 < 18; fi21++ {
                v7 -= fi21
                _ = fi21
            }
        default:
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int = v1
                v24 := GenS1{F0: 1, F1: fi22}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    case "a":
        v25 := GenS3{F0: 826, F1: 32767, F2: 2}
        _ = v25
    }
    if (18446744073709551615 != 675) {
        v2 = (v2 & -105)
        for fi26 := 0; fi26 < 17; fi26++ {
            if false {
                v1 = fi26
                v27 := GenS3{F0: 1000, F1: -999, F2: 0}
                var v28 float64 = 0.0
                _ = v27
                _ = v28
            } else {
                v6 += (^(v6))
            }
            _ = fi26
        }
        v29 := GenS4{F0: v16, F1: -119}
        _ = v29
    } else {
        var v30 uint64 = 0
        switch (v1 % (((^(v7))) | 1)) {
        case -724:
            v31 := GenS2{F0: v16, F1: 1000}
            _ = v31
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := GenS5{F0: (1 > 32767), F1: -404}
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint = (v6 / ((4294967295) | 1))
                _ = fi34
                _ = v35
            }
            var v36 bool = false
            _ = fi32
            _ = v33
            _ = v36
        }
        _ = v30
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v37 string = (("\n\t" + ((((((("hello" + ("" + ("" + ("" + ("\n\t" + "foo bar"))))) + ("foo bar" + "\n\t")) + "hello") + ("foo bar" + "\n\t")) + "\n\t") + ("" + "foo bar")) + "hello")) + "a")
        if len(v5) > 0 {
            v5[0] = 0
        }
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
