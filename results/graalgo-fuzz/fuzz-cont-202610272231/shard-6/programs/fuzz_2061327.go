package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint8; F1 uint16; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int16; F2 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint32; F1 int64 }
func genFn0() int16 {
    v6 := []int{990, 796, 424, 42}
    var v7 bool = true
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v6)
        fmt.Println(v7, fi8)
        if (!((true && true) && false)) {
            for fi9 := 0; fi9 < 14; fi9++ {
                if false {
                    v7 = v7
                    v10 := GenS1{F0: 2}
                    v7 = true
                    _ = v10
                } else {
                    switch fi9 {
                    case 32768:
                        fmt.Println(fi9, v7, v7)
                    }
                }
                v11 := GenS2{F0: 42, F1: -475, F2: 190}
                v6 = append(v6, fi9)
                _ = fi9
                _ = v11
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                v7 = (0.0 != (1.5 * 0.1))
                v7 = true
                _ = fi12
            }
            var v13 int8 = -1
            _ = v13
        } else {
            if v7 {
                if false {
                    if len(v6) > 0 {
                        v6[0] = fi8
                    }
                    v7 = v7
                    var v14 int32 = 0
                    _ = v14
                } else {
                    fmt.Println(fi8, v7)
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v7 = false
                        if len(v6) > 0 {
                            v6[0] = (fi8 & -1)
                        }
                        var v16 int64 = 1
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS3{F0: 42, F1: (2 >= 0)}
                    _ = v17
                }
                if len(v6) > 0 {
                    v6[0] = (fi8 - fi8)
                }
            } else {
                v18 := GenS4{F0: 224, F1: 371, F2: fi8}
                switch "hello" {
                case "hello":
                    v7 = (42 != -123)
                case "foo bar":
                    if v7 {
                        var v19 uint32 = 0
                        v19 -= v19
                        _ = v19
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v20 int64 = 1
                        _ = v20
                    }
                case "":
                    var v21 bool = (!true)
                    _ = v21
                }
                _ = v18
            }
            v22 := GenS5{F0: fi8}
            v23 := GenS6{F0: fi8, F1: 0, F2: 853}
            _ = v22
            _ = v23
        }
        _ = fi8
    }
    switch "" {
    case "\n\t":
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS4{F0: 2, F1: 42, F2: fi24}
            var v26 uint64 = 18446744073709551615
            _ = fi24
            _ = v25
            _ = v26
        }
    case "a":
        _ = len(v6)
    case "foo bar":
        v27 := GenS7{F0: false, F1: -1}
        _ = v27
    }
    var v28 bool = v7
    _ = v28
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint32) int8 {
    v29 := GenS2{F0: -534, F1: -1, F2: 114}
    fmt.Println(p0)
    v30 := GenS5{F0: -2147483648}
    v31 := []int{}
    _ = p0
    _ = v29
    _ = v30
    _ = v31
    return 127
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{1, 4294967295, 65535, 65535}
    v3 := make(map[int]string)
    var v4 uint16 = 1
    var v5 int8 = 42
    delete(v3, 9223372036854775807)
    fmt.Println(v4, v5)
    switch "foo bar" {
    case "a":
        var v32 uint = 4294967295
        _ = v32
    }
    if (v1 == (((((v1 / (0.1 - v1)) - (0.1 - 0.0)) + v1) * (v1 / v1)) * float64((0.0 / 1.0)))) {
        for fi33 := 0; fi33 < 17; fi33++ {
            switch "hello" {
            case "hello":
                var v34 uint64 = 18446744073709551615
                _ = v34
            default:
                switch fi33 {
                case -1:
                    for fi35 := 0; fi35 < 2; fi35++ {
                        v4 = 273
                        v36 := v3[fi33]
                        var v37 float32 = 3.14159
                        v5 = (v5 % ((v5) | 1))
                        v38 := GenS8{F0: 711, F1: 2147483647}
                        v37 = v37
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                case -2147483648:
                    v3[(fi33 + fi33)] = "\n\t"
                default:
                    var v39 uint = 567
                    _ = v39
                }
            }
            var v40 uint16 = (^(v4))
            _ = fi33
            _ = v40
        }
    }
    v5 += 100
    if (-731 > 127) {
        var v41 string = ("\n\t" + (((("" + ("hello" + "foo bar")) + ("hello" + ("\n\t" + "\n\t"))) + "hello") + ("a" + ("hello" + ("" + "a")))))
        v42 := GenS9{F0: 2, F1: -1}
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        v1 -= 1e10
        var v44 int16 = genFn0()
        var v45 string = ("hello" + "")
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := GenS3{F0: -2147483648, F1: (42 == 0)}
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
