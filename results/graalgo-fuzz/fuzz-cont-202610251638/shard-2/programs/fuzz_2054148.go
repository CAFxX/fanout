package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0(p0 int8) uint64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            fmt.Println(fi6, fi5, p0)
            v7 := GenS1{F0: 1, F1: (fi5 * fi5), F2: 2147483647}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if false {
        p0 = (p0 & (p0 << ((2) & 63)))
    } else {
        var v8 uint = 2
        v8 -= (v8 << (((v8 - 18446744073709551615)) & 63))
        for fi9 := 0; fi9 < 4; fi9++ {
            v8 = (v8 % ((466) | 1))
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint16 = 2
                v12 := []int{fi10, fi10, 32768, (fi9 * (fi9 * 2147483647))}
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := []int{(fi9 % (((fi9 / ((fi9) | 1))) | 1)), fi9, -33, (fi9 / ((32768) | 1))}
            _ = fi9
            _ = v13
        }
        _ = v8
    }
    _ = p0
    return 843
}
func genFn1(p0 int32) int64 {
    v14 := []int{844}
    switch "foo bar" {
    case "hello":
        p0 += p0
    case "\n\t":
        v14 = append(v14, -1)
    }
    var v15 uint8 = 2
    if len(v14) > 0 {
        v14[0] = -895
    }
    _ = p0
    _ = v14
    _ = v15
    return -858
}

func main() {
    var v1 uint = 2
    var v2 int16 = int16(v1)
    var v3 int8 = 1
    v4 := make(map[string]int)
    for fi16 := 0; fi16 < 10; fi16++ {
        v3 = (v3 << ((v1) & 63))
        v17 := GenS2{F0: genFn1(-1), F1: 0, F2: 0}
        _ = fi16
        _ = v17
    }
    var v18 int = -135
    if (!false) {
        v19 := GenS1{F0: -1, F1: 525, F2: 712}
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v4["foo bar"] = 0
        fmt.Println(v1, v1, v1)
        v1 = (^(v1))
        _ = fi20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 string = "a"
        _ = fi21
        _ = v22
    }
    var v23 int32 = -29
    var v24 int = v18
    for fi25 := 0; fi25 < 1; fi25++ {
        if true {
            v3 -= 42
        } else {
            if true {
                v2 = (v2 >> (((v1 << ((v1) & 63))) & 63))
            }
            v26 := GenS2{F0: 0, F1: 1, F2: v2}
            _ = v26
        }
        var v27 int = -1
        var v28 bool = (!(v1 <= 4294967295))
        _ = fi25
        _ = v27
        _ = v28
    }
    v29 := v4[""]
    var v30 uint8 = 1
    v31 := GenS1{F0: (v23 % ((-782) | 1)), F1: (v24 % (((^(v18))) | 1)), F2: 636}
    for fi32 := 0; fi32 < 20; fi32++ {
        v4[("foo bar" + (("" + ("" + "hello")) + "hello"))] = v24
        if true {
            fmt.Println(v30)
        } else {
            v33 := []int{127, -2147483648, 128, 127}
            _ = v33
        }
        v34 := []int{}
        _ = fi32
        _ = v34
    }
    v35 := GenS2{F0: -1, F1: 65535, F2: v2}
    var v36 uint8 = (v30 & 31)
    v29 -= (v24 / ((174) | 1))
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 uint32 = 532
        _ = fi37
        _ = v38
    }
    v39 := []int{v24, int(v3), 4294967295, v29}
    switch v24 {
    case 417:
        var v40 float32 = (-2.25 - float32(0.1))
        _ = v40
    case 0:
        if false {
            _ = len(v4)
        }
    case 867:
        if ((v23 % ((42) | 1)) == v23) {
            v2 = v2
            var v41 uint8 = (v30 | (v30 ^ v36))
            _ = v41
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
                if true {
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                } else {
                    var v43 float64 = (0.0 - ((3.14159 - 0.1) * 1.5))
                    _ = v43
                }
                var v44 string = "hello"
                _ = fi42
                _ = v44
            }
        }
    default:
        var v45 int64 = 1
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
