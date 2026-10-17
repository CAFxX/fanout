package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint16; F2 bool }
type GenS5 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 bool) int {
    v4 := GenS1{F0: 2, F1: 100, F2: 65535}
    p0 = (!(-2147483648 == -2147483648))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int16 = 1000
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    var v8 float64 = (1.5 - float64((3.14159 * (3.14159 * (((((1.5 + 3.14159) - 1e10) / -2.25) - (3.14159 + (0.1 * (1.5 * float32((0.1 - (0.0 / 3.14159))))))) + (-2.25 + 3.14159))))))
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 394
}
func genFn1(p0 uint, p1 uint) int64 {
    p0 = (^(p1))
    v9 := GenS2{F0: (2 <= 304), F1: (!false), F2: 907}
    _ = p0
    _ = p1
    _ = v9
    return -339
}

func main() {
    v1 := []uint32{4294967295, 42, 2147483648, 4294967295}
    var v2 uint = 1
    v3 := make(map[uint32]bool)
    switch -354 {
    case 42:
        var v10 string = "hello"
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        switch "hello" {
        case "foo bar":
            v12 := GenS3{F0: v2}
            _ = v12
        case "hello":
            v2 = 895
        }
        _ = fi11
    }
    v13 := GenS4{F0: 873, F1: 2, F2: false}
    v2 -= v2
    v2 = v2
    var v14 int16 = 32767
    var v15 int32 = 1
    if (!(165 >= 0)) {
        var v16 uint16 = 2
        var v17 uint8 = 0
        for fi18 := 0; fi18 < 7; fi18++ {
            switch fi18 {
            case 4294967295:
                for fi19 := 0; fi19 < 20; fi19++ {
                    if ((^(v15)) < v15) {
                        v20 := GenS5{F0: v14, F1: (v17 & (v17 + v17)), F2: 0}
                        switch int(v14) {
                        case -9223372036854775808:
                            var v21 int32 = (^(v15))
                            _ = v21
                        }
                        v16 -= 65535
                        _ = v20
                    }
                    var v22 string = "\n\t"
                    if ((v15 + (-(v15))) < v15) {
                        fmt.Println(v17)
                        switch -1 {
                        case -803:
                            if (42 >= -684) {
                                if (v14 != v14) {
                                    if false {
                                        var v23 uint = v2
                                        v24 := GenS1{F0: 0, F1: -115, F2: 4294967295}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v2 += v2
                                        v25 := v3[148]
                                        v26 := GenS5{F0: v14, F1: v17, F2: 2147483647}
                                        _ = v25
                                        _ = v26
                                    }
                                    delete(v3, 764)
                                }
                                switch ("" + ((v22 + v22) + (v22 + (("hello" + (v22 + "foo bar")) + v22)))) {
                                case "a":
                                    var v27 float64 = 0.1
                                    _ = v27
                                }
                            }
                            switch (^(fi18)) {
                            case -138:
                                var v28 int64 = -350
                                _ = v28
                            default:
                                var v29 uint16 = 0
                                _ = v29
                            }
                        case 0:
                            fmt.Println(v17, v22, v15)
                        case 2147483647:
                            v15 -= (v15 % (((^(v15))) | 1))
                        }
                        v15 = v15
                    }
                    _ = fi19
                    _ = v22
                }
            case 9223372036854775807:
                v30 := GenS3{F0: v2}
                _ = v30
            case 32768:
                v31 := GenS5{F0: 42, F1: v17, F2: 65535}
                _ = v31
            default:
                v32 := GenS6{F0: (v17 + 42), F1: v14}
                _ = v32
            }
            v33 := GenS6{F0: (v17 << (((v2 & (v2 / ((v2) | 1)))) & 63)), F1: (-(v14))}
            if ((fi18 | (fi18 << (((v2 ^ (v2 >> (((v2 + 255)) & 63)))) & 63))) < (^(fi18))) {
                for fi34 := 0; fi34 < 9; fi34++ {
                    switch ("hello" + "a") {
                    case "hello":
                        var v35 uint64 = 185
                        _ = v35
                    }
                    _ = fi34
                }
                var v36 int = 2147483647
                v37 := GenS7{F0: (fi18 & 2147483647), F1: 1, F2: 4294967295}
                _ = v36
                _ = v37
            }
            _ = fi18
            _ = v33
        }
        _ = v16
        _ = v17
    }
    var v38 uint16 = 255
    var v39 float32 = (0.0 / 1e10)
    for fi40 := 0; fi40 < 2; fi40++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v15, fi40)
        v39 += v39
        _ = fi40
    }
    var v41 int = 128
    v42 := v3[42]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v41
    _ = v42
}
