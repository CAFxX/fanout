package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 int32, p1 uint8) uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = 42
        _ = fi6
        _ = v7
    }
    if ((!(float64(1.5) != 1e10)) && true) {
        switch 1 {
        case 32767:
            var v8 uint32 = 645
            _ = v8
        default:
            p0 = 1
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 float32 = 1e10
            p0 = p0
            for fi11 := 0; fi11 < 13; fi11++ {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        var v12 uint64 = uint64(p1)
                        _ = v12
                    case "":
                        v13 := []int{fi9}
                        _ = v13
                    case "a":
                        switch "foo bar" {
                        case "\n\t":
                            p0 -= -166
                        case "hello":
                            var v14 uint8 = (p1 * (p1 % ((p1) | 1)))
                            _ = v14
                        case "":
                            v10 += 1.5
                        }
                    }
                } else {
                    v10 = (((1e10 + v10) - (v10 - 0.0)) / (0.1 * v10))
                }
                switch (fi11 ^ (fi11 % ((fi11) | 1))) {
                case 1:
                    v10 += (v10 - float32(0.1))
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    } else {
        switch 973 {
        case -1:
            v15 := GenS1{F0: 0}
            _ = v15
        case 0:
            var v16 int8 = 42
            _ = v16
        }
        v17 := GenS1{F0: 100}
        var v18 int = -1
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int16, p1 int, p2 uint32) uint8 {
    p2 = genFn0(-2147483648, 2)
    if false {
        fmt.Println(p0)
    }
    p0 = p0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 64
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 0
    v3 := make(map[int]bool)
    var v4 float32 = (-2.25 + (1e10 * 0.1))
    var v5 int32 = 0
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 string = "\n\t"
        v21 := []int{fi19, (fi19 ^ fi19), (^(fi19)), (fi19 >> ((776) & 63))}
        for fi22 := 0; fi22 < 12; fi22++ {
            v2 += (v2 * 0)
            switch fi19 {
            case -2147483648:
                if ((1e10 != 0.0) && true) {
                    v1 = v1
                    v23 := GenS1{F0: -128}
                    fmt.Println(fi22, v5, v2)
                    _ = v23
                }
            case 127:
                var v24 float64 = 3.14159
                _ = v24
            default:
                v25 := GenS1{F0: -128}
                _ = v25
            }
            if false {
                fmt.Println(fi22, v2)
                v4 = v4
            }
            var v26 int16 = 0
            v27 := GenS2{F0: false, F1: 2}
            _ = fi22
            _ = v26
            _ = v27
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v5)
    fmt.Println(v2, v2)
    switch -855 {
    case 54:
        var v28 int8 = 44
        _ = v28
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            if (false && (100 > int8(v2))) {
                v4 = ((v4 / ((((v4 * v4) + 0.1) / v4) + 1e10)) + (float32(3.14159) / v4))
                v29 := []int{0, 9223372036854775807}
                var v30 uint32 = 65535
                _ = v29
                _ = v30
            } else {
                if false {
                    v31 := v3[32767]
                    if v31 {
                        v2 = 42
                        fmt.Println(v4, v5, v4)
                        var v32 int32 = (v5 + (^(v5)))
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 13; fi33++ {
                        var v34 float32 = ((v4 + v4) / ((v4 * (v4 / v4)) * v4))
                        v34 += ((v34 + (-2.25 - (float32(float64(float32((0.0 + 1e10)))) - v4))) / float32(float64(v34)))
                        v35 := GenS3{F0: "foo bar", F1: 18446744073709551615, F2: v31}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = v31
                } else {
                    if ((v2 | 121) >= v2) {
                        if (!false) {
                            v1 = 29
                        } else {
                            fmt.Println(v5, v5)
                        }
                    }
                }
                var v36 uint16 = 65535
                v37 := GenS4{F0: 0}
                _ = v36
                _ = v37
            }
        case "":
            v1 = 9223372036854775807
        }
    }
    v38 := v3[1]
    fmt.Println(v38, v5)
    var v39 int16 = -1
    v40 := []int{9223372036854775807, 128}
    if (false && v38) {
        fmt.Println(v2)
    }
    v38 = ((0.0 - 0.0) >= 1e10)
    v41 := GenS5{F0: -2147483648, F1: 2}
    v38 = ((!false) || v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
