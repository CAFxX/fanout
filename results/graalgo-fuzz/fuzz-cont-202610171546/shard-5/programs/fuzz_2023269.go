package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 uint, p1 uint32, p2 int8) int {
    switch "" {
    case "foo bar":
        if true {
            if (!(!true)) {
                v4 := GenS1{F0: -163}
                if (!(!false)) {
                    for fi5 := 0; fi5 < 4; fi5++ {
                        var v6 uint64 = 630
                        _ = fi5
                        _ = v6
                    }
                    v7 := []int{-2147483648, 42, 850, 1}
                    var v8 uint16 = 347
                    _ = v7
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v10 := []int{(-(fi9)), (fi9 / ((fi9) | 1))}
                        p2 = p2
                        _ = fi9
                        _ = v10
                    }
                    v11 := []int{463, -2147483648, 286}
                    _ = v11
                }
                var v12 uint = 65535
                _ = v12
                _ = v4
            }
            if true {
                var v13 uint16 = 65535
                _ = v13
            }
            v14 := []int{9223372036854775807, -431, -1, int(p0)}
            _ = v14
        }
    default:
        fmt.Println(p2, p2, p1)
    }
    var v15 uint = (^(p0))
    p1 = p1
    var v16 float32 = (-2.25 * ((1e10 * 1.5) / 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 128
}
func genFn1(p0 uint64) string {
    if ((!(70 < 255)) && (42 != 1)) {
        p0 -= (p0 + p0)
        var v17 uint8 = 29
        var v18 bool = false
        _ = v17
        _ = v18
    }
    fmt.Println(p0)
    var v19 int8 = 0
    _ = p0
    _ = v19
    return "a"
}

func main() {
    v1 := []int{}
    var v2 float32 = 0.0
    var v3 int = 9223372036854775807
    var v20 int16 = -32768
    v2 -= float32(3.14159)
    _ = len(v1)
    switch (v3 >> ((1) & 63)) {
    case 128:
        var v21 float64 = -2.25
        _ = v21
    default:
        for fi22 := 0; fi22 < 17; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                if len(v1) > 0 {
                    v1[0] = (fi23 ^ (fi22 << ((974) & 63)))
                }
                v2 = 1e10
                switch "a" {
                case "":
                    switch (v3 / (((v3 + (^(fi22)))) | 1)) {
                    case 125:
                        fmt.Println(v20, v3)
                    case 459:
                        v24 := GenS1{F0: -579}
                        _ = v24
                    }
                case "\n\t":
                    var v25 bool = true
                    _ = v25
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    var v26 uint32 = 1
    for fi27 := 0; fi27 < 4; fi27++ {
        v28 := GenS2{F0: 531, F1: (false && true), F2: 127}
        for fi29 := 0; fi29 < 13; fi29++ {
            if len(v1) > 0 {
                v1[0] = (fi27 * 128)
            }
            var v30 float64 = 3.14159
            v2 -= (0.1 - v2)
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v31 int32 = 1
    _ = len(v1)
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 uint8 = 255
        _ = fi32
        _ = v33
    }
    v26 -= (v26 % (((v26 >> ((4294967295) & 63))) | 1))
    switch "\n\t" {
    case "foo bar":
        v34 := GenS3{F0: "", F1: 255}
        _ = v34
    case "\n\t":
        v2 += -2.25
    default:
        if true {
            v26 = (v26 % ((385) | 1))
        }
    }
    v35 := GenS2{F0: genFn0(18446744073709551615, v26, -127), F1: false, F2: 0}
    var v36 int16 = v20
    if true {
        v37 := GenS4{F0: ("\n\t" + ("\n\t" + "a")), F1: 128}
        v36 += (v20 * v36)
        _ = v37
    } else {
        v1 = append(v1, v3)
    }
    v2 = v2
    var v38 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v38
}
