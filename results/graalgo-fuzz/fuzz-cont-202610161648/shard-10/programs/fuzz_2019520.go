package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int16 }

func main() {
    var v1 float64 = float64(0.0)
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v1 = v1
    v4 := []int{}
    var v5 uint16 = 2
    switch ("" + ("hello" + "")) {
    case "a":
        var v6 float32 = (0.0 / 0.1)
        _ = v6
    }
    v1 += (v1 - (((((v1 + v1) * 1.5) + v1) - ((v1 / ((v1 * v1) + v1)) / v1)) - float64((-2.25 + float32(1.5)))))
    var v7 float32 = ((((1.5 / 1.0) + (-2.25 * (3.14159 + -2.25))) + (3.14159 * 0.0)) / -2.25)
    v5 = 777
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 -= float32(v1)
        var v9 int16 = -1
        for fi10 := 0; fi10 < 3; fi10++ {
            if true {
                switch (fi8 + fi8) {
                case -1:
                    v5 = (v5 - v5)
                    var v11 bool = ((0.0 - -2.25) > -2.25)
                    var v12 int32 = 301
                    if (100 < 104) {
                        switch "\n\t" {
                        case "hello":
                            v13 := GenS1{F0: 1, F1: 2147483647, F2: -1}
                            _ = v13
                        case "\n\t":
                            v5 = 0
                        case "":
                            var v14 uint16 = v5
                            _ = v14
                        default:
                            var v15 uint = 4294967295
                            switch ("" + ("foo bar" + ("foo bar" + ("" + ("" + ("foo bar" + ("a" + "foo bar"))))))) {
                            case "a":
                                _ = len(v2)
                            case "hello":
                                v5 += 541
                            case "":
                                var v16 bool = false
                                _ = v16
                            default:
                                var v17 int16 = v9
                                _ = v17
                            }
                            _ = v15
                        }
                        var v18 int64 = -332
                        v18 -= v18
                        _ = v18
                    } else {
                        v19 := GenS1{F0: v12, F1: 296, F2: 9223372036854775807}
                        if (1 > 0) {
                            if v11 {
                                fmt.Println(v1, fi8)
                                delete(v3, fi10)
                                v20 := GenS1{F0: v12, F1: 189, F2: -1}
                                _ = v20
                            }
                        }
                        var v21 float32 = v7
                        _ = v19
                        _ = v21
                    }
                    v9 = (-(v9))
                    _ = v11
                    _ = v12
                }
                if false {
                    if (false && false) {
                        fmt.Println(fi10, v5, fi10)
                    } else {
                        switch (fi10 << ((2) & 63)) {
                        case 32767:
                            var v22 string = "a"
                            _ = v22
                        case 1:
                            v1 -= v1
                        case 128:
                            v23 := GenS2{F0: 354}
                            _ = v23
                        default:
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                fmt.Println(v5, v5, v7)
                            default:
                                if (true || ((v5 << ((961) & 63)) >= 65535)) {
                                    var v24 uint64 = 18446744073709551615
                                    v25 := GenS1{F0: 0, F1: 269, F2: 9223372036854775807}
                                    v9 += 1
                                    _ = v24
                                    _ = v25
                                } else {
                                    switch fi10 {
                                    case 282:
                                        v26 := GenS2{F0: 323}
                                        _ = v26
                                    }
                                    v7 = 3.14159
                                }
                            }
                        }
                        v5 -= 0
                    }
                    _ = len(v4)
                    var v27 uint16 = 42
                    switch fi8 {
                    case 986:
                        v28 := GenS2{F0: 1}
                        _ = v28
                    case 0:
                        v1 = 3.14159
                    }
                    _ = v27
                } else {
                    var v29 uint16 = (v5 << ((875) & 63))
                    v29 = 0
                    _ = v29
                }
                if true {
                    if (!(!true)) {
                        v30 := GenS3{F0: 2}
                        v1 = v1
                        _ = v30
                    } else {
                        var v31 uint16 = v5
                        v9 = (v9 >> ((431) & 63))
                        switch "\n\t" {
                        case "a":
                            if (!true) {
                                if true {
                                    v32 := GenS4{F0: (0 > 42), F1: true, F2: 708}
                                    if (!true) {
                                        v7 -= v7
                                        _ = len(v4)
                                        fmt.Println(v31, v5, v5)
                                    }
                                    var v33 float32 = ((3.14159 * float32(float64(0.0))) / (0.0 * (v7 / v7)))
                                    _ = v32
                                    _ = v33
                                }
                                if (v9 == (v9 / ((0) | 1))) {
                                    var v34 int16 = (v9 * 786)
                                    _ = v34
                                }
                            }
                        case "hello":
                            v7 += ((v7 + v7) - ((v7 / v7) * (v7 + (v7 / (float32(1.5) / (v7 - ((1.5 - 3.14159) - v7)))))))
                        case "foo bar":
                            v35 := GenS5{F0: (128 > 128), F1: ((v1 >= float64((0.0 - (v7 - v7)))) && true)}
                            _ = v35
                        default:
                            if (true && true) {
                                var v36 uint = 42
                                var v37 int32 = 895
                                v7 += ((3.14159 - (v7 / v7)) - (1.5 / v7))
                                _ = v36
                                _ = v37
                            }
                        }
                        fmt.Println(fi8, fi8, v5)
                        fmt.Println(fi10)
                        var v38 string = "\n\t"
                        _ = v31
                        _ = v38
                    }
                    var v39 uint32 = 2147483647
                    _ = v39
                }
            }
            if false {
                fmt.Println(fi8)
                _ = len(v4)
                var v40 uint8 = 127
                _ = len(v4)
                var v41 int8 = 100
                fmt.Println(v41, v5)
                var v42 int = -779
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    if (true || ((!(false && false)) && (uint(v5) > 283))) {
        var v43 int8 = 100
        _ = v43
    }
    v7 = (v7 - v7)
    switch "" {
    case "":
        fmt.Println(v1)
    case "hello":
        for fi44 := 0; fi44 < 14; fi44++ {
            v1 += 3.14159
            _ = fi44
        }
    case "a":
        var v45 uint = 0
        _ = v45
    default:
        if true {
            v7 = v7
            for fi46 := 0; fi46 < 4; fi46++ {
                v7 += 1e10
                for fi47 := 0; fi47 < 4; fi47++ {
                    var v48 int = 153
                    var v49 string = ("hello" + "hello")
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                if false {
                    var v50 float64 = float64(v7)
                    if len(v4) > 0 {
                        v4[0] = fi46
                    }
                    v51 := GenS6{F0: 42, F1: (v5 % ((v5) | 1))}
                    _ = v50
                    _ = v51
                } else {
                    v52 := GenS7{F0: 0, F1: 1000}
                    for fi53 := 0; fi53 < 5; fi53++ {
                        v1 += (v1 + 0.0)
                        var v54 uint = 713
                        var v55 int = 1
                        var v56 string = "foo bar"
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    v1 = float64(-2.25)
                    _ = v52
                }
                _ = fi46
            }
        } else {
            v5 -= v5
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
