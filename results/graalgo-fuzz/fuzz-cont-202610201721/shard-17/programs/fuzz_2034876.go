package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 uint16, p1 uint8, p2 int) uint32 {
    switch p2 {
    case -9223372036854775808:
        var v7 float64 = (3.14159 / 1.0)
        _ = v7
    case 32767:
        p1 = (^(p1))
    case 42:
        var v8 uint8 = p1
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{fi9, fi9}
        switch fi9 {
        case 128:
            p2 -= (fi9 / (((p2 | fi9)) | 1))
        }
        _ = fi9
        _ = v10
    }
    switch p2 {
    case 191:
        p2 += 32767
    case 1:
        v11 := GenS1{F0: -174, F1: 1}
        _ = v11
    case 42:
        p0 -= (p0 + 826)
    default:
        p0 = p0
    }
    var v12 float32 = (0.1 + (float32((0.0 - float64((3.14159 - (3.14159 * 1e10))))) + 0.1))
    var v13 int8 = 35
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2147483647
}
func genFn1(p0 uint, p1 int32, p2 uint32) uint {
    for fi14 := 0; fi14 < 11; fi14++ {
        switch "a" {
        case "":
            var v15 bool = true
            _ = v15
        case "a":
            if (-2.25 <= -2.25) {
                v16 := GenS2{F0: ((p2 < (p2 / ((genFn0(2, 2, fi14)) | 1))) || (!true)), F1: (p0 >> (((p0 % ((p0) | 1))) & 63))}
                v17 := []int{(fi14 - fi14), fi14, fi14, (-(fi14))}
                _ = v16
                _ = v17
            } else {
                switch "foo bar" {
                case "hello":
                    var v18 uint16 = 65535
                    _ = v18
                case "foo bar":
                    var v19 float64 = 3.14159
                    _ = v19
                case "a":
                    if true {
                        var v20 bool = (-879 != 9223372036854775807)
                        var v21 int64 = -1
                        _ = v20
                        _ = v21
                    } else {
                        if false {
                            v22 := GenS1{F0: 0, F1: 82}
                            v23 := []int{}
                            _ = v22
                            _ = v23
                        } else {
                            p0 -= p0
                            fmt.Println(p1)
                            switch ("a" + "hello") {
                            case "\n\t":
                                p1 += p1
                            case "hello":
                                var v24 bool = true
                                _ = v24
                            default:
                                v25 := []int{(^(fi14)), 2147483647}
                                _ = v25
                            }
                        }
                        v26 := []int{fi14, (fi14 & (-(fi14))), 32768, 127}
                        fmt.Println(fi14, fi14, p2)
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 2; fi27++ {
                        switch fi14 {
                        case 32767:
                            var v28 uint16 = 255
                            _ = v28
                        }
                        _ = fi27
                    }
                }
                p1 += (p1 & (p1 / ((-1) | 1)))
                var v29 int8 = 1
                _ = v29
            }
        }
        _ = fi14
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        switch "hello" {
        case "":
            var v31 int16 = 1000
            _ = v31
        case "foo bar":
            switch "" {
            case "hello":
                p0 += (p0 >> ((255) & 63))
            default:
                var v32 uint32 = genFn0(2, 52, 127)
                _ = v32
            }
        }
        fmt.Println(fi30)
        _ = fi30
    }
    v33 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return 18446744073709551615
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint32 = 862
    var v3 string = "hello"
    var v4 float32 = 1e10
    v5 := make(map[int]int)
    var v6 uint32 = 4294967295
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := []int{v1, (-(fi34)), (fi34 >> ((0) & 63))}
        var v36 uint64 = 4294967295
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println(v3)
    v1 += (^(v1))
    var v37 uint8 = 128
    var v38 int32 = 0
    if true {
        var v39 int16 = -1
        _ = v39
    } else {
        for fi40 := 0; fi40 < 12; fi40++ {
            for fi41 := 0; fi41 < 15; fi41++ {
                var v42 float32 = -2.25
                _ = fi41
                _ = v42
            }
            var v43 string = "a"
            v1 = fi40
            _ = fi40
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        v45 := GenS2{F0: (1.5 != (3.14159 - ((v4 + 3.14159) + ((v4 / v4) / 3.14159)))), F1: 4294967295}
        var v46 uint8 = v37
        _ = len(v5)
        _ = fi44
        _ = v45
        _ = v46
    }
    v47 := []int{v1, v1, (-(v1)), v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
