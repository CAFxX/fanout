package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 int }
func genFn0(p0 int, p1 int32) uint8 {
    var v7 uint8 = 180
    switch "foo bar" {
    case "":
        var v8 uint16 = 32767
        _ = v8
    case "foo bar":
        var v9 float32 = 1e10
        _ = v9
    case "a":
        p1 = p1
    default:
        switch "hello" {
        case "\n\t":
            var v10 int32 = int32(p0)
            _ = v10
        case "hello":
            p1 = p1
        case "foo bar":
            var v11 uint32 = 2
            _ = v11
        }
    }
    v12 := []int{p0, p0, int(p1), (p0 & -276)}
    v13 := GenS1{F0: 18446744073709551615, F1: 811, F2: -2147483648}
    v14 := GenS1{F0: 4294967295, F1: 0, F2: (-(p0))}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v7
    return 202
}
func genFn1(p0 int16, p1 uint64, p2 int64) int16 {
    var v15 uint64 = (^(p1))
    v15 -= v15
    var v16 uint16 = 642
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return p0
}

func main() {
    v1 := []int{}
    var v2 int16 = 1
    var v3 uint = 1
    var v4 int32 = 0
    var v5 float64 = 3.14159
    var v6 int64 = 42
    v17 := GenS2{F0: -825}
    if true {
        if ((v4 * 2147483647) == v4) {
            for fi18 := 0; fi18 < 8; fi18++ {
                if true {
                    v2 += (v2 / ((v2) | 1))
                    v19 := GenS2{F0: genFn1((v2 - genFn1(1, 1, v6)), 1, -9223372036854775808)}
                    _ = v19
                } else {
                    var v20 uint = (v3 - 18446744073709551615)
                    fmt.Println(v3, v5, fi18)
                    if false {
                        fmt.Println(v6, v2)
                        var v21 float64 = float64(3.14159)
                        if (426 < 48) {
                            var v22 bool = ((false && true) && (!true))
                            for fi23 := 0; fi23 < 7; fi23++ {
                                _ = len(v1)
                                var v24 int8 = 100
                                var v25 int = (fi18 ^ -9223372036854775808)
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS3{F0: (("" + ("a" + "hello")) + (("foo bar" + ("a" + ("hello" + "\n\t"))) + (("" + "a") + ("foo bar" + "")))), F1: (v6 / ((v6) | 1))}
                            _ = v22
                            _ = v26
                        } else {
                            if len(v1) > 0 {
                                v1[0] = (fi18 | (fi18 - (fi18 / ((fi18) | 1))))
                            }
                            var v27 uint32 = 0
                            _ = v27
                        }
                        _ = v21
                    } else {
                        fmt.Println(v2, v20)
                    }
                    _ = v20
                }
                v28 := GenS4{F0: v6}
                _ = fi18
                _ = v28
            }
            var v29 uint64 = 321
            _ = v29
        }
        if len(v1) > 0 {
            v1[0] = 880
        }
    } else {
        fmt.Println(v5, v3, v6)
        _ = len(v1)
        switch "\n\t" {
        case "":
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 float32 = (((3.14159 * 0.0) / 1.0) - float32((v5 - ((v5 / 1.5) / (v5 / v5)))))
                v32 := GenS5{F0: 898, F1: 402}
                _ = fi30
                _ = v31
                _ = v32
            }
        case "foo bar":
            v1 = append(v1, 860)
        default:
            var v33 uint16 = 54
            _ = v33
        }
    }
    var v34 int32 = -75
    v35 := GenS6{F0: 32768, F1: uint64(v3), F2: 127}
    fmt.Println(v34, v6, v6)
    var v36 int = 32767
    var v37 float64 = (v5 / 1.5)
    v38 := GenS4{F0: -1}
    for fi39 := 0; fi39 < 17; fi39++ {
        v1 = append(v1, 1)
        _ = fi39
    }
    if true {
        var v40 float32 = (0.1 / (float32(0.1) - 1e10))
        var v41 uint64 = 42
        _ = v40
        _ = v41
    }
    var v42 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        v36 += v36
        _ = fi43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 uint32 = 4294967295
        var v46 uint64 = 0
        for fi47 := 0; fi47 < 8; fi47++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    v42 = (((float64(-2.25) + v42) / (v37 * ((0.1 + float64(0.1)) * ((v37 * v37) / v37)))) - v37)
    var v48 float64 = -2.25
    switch "a" {
    case "hello":
        var v49 uint16 = 1
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v6
}
