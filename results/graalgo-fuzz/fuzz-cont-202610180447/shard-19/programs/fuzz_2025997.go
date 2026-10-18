package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 float64 = 0.1
    if (true || false) {
        if v2 {
            switch "foo bar" {
            case "\n\t":
                switch "hello" {
                case "hello":
                    v4 := v1[65535]
                    _ = v4
                }
            default:
                if (2 <= 2) {
                    v3 = 0.1
                    v3 -= (((1.5 + ((v3 / (v3 - 1e10)) + float64(-2.25))) - v3) * (0.0 * ((float64(float32(3.14159)) + (float64(((3.14159 - 0.0) + (1e10 - 1.5))) + 1.5)) / 1.5)))
                } else {
                    switch 375 {
                    case 32767:
                        v5 := []int{4294967295}
                        _ = v5
                    case 128:
                        for fi6 := 0; fi6 < 1; fi6++ {
                            switch fi6 {
                            case -959:
                                var v7 float32 = float32(float64(((1e10 + (1e10 * 1.5)) / 1.0)))
                                _ = v7
                            case 127:
                                if (4294967295 > 943) {
                                    fmt.Println(v3)
                                }
                            }
                            for fi8 := 0; fi8 < 13; fi8++ {
                                var v9 int16 = -1
                                v3 -= 0.0
                                switch "" {
                                case "foo bar":
                                    var v10 int64 = -830
                                    _ = v10
                                case "hello":
                                    v11 := GenS1{F0: 289, F1: v9, F2: 1}
                                    _ = v11
                                case "":
                                    var v12 bool = v2
                                    _ = v12
                                }
                                fmt.Println(fi8, v2)
                                v13 := GenS1{F0: 0, F1: (v9 >> ((1) & 63)), F2: 606}
                                v14 := GenS2{F0: (v9 / ((1) | 1)), F1: 653}
                                v15 := GenS1{F0: 65535, F1: -1, F2: 255}
                                _ = fi8
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v9
                            }
                            _ = fi6
                        }
                    }
                    var v16 float32 = (0.0 - (1e10 / (0.0 + float32(float64(1.5)))))
                    v16 = (v16 - v16)
                    _ = v16
                }
            }
        }
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            delete(v1, 142)
            v2 = v2
            _ = fi17
        }
    }
    if false {
        var v18 string = "foo bar"
        v19 := GenS3{F0: false, F1: -1, F2: 2}
        v20 := GenS4{F0: 65535, F1: 2}
        _ = v18
        _ = v19
        _ = v20
    }
    v3 = (0.1 / v3)
    switch "a" {
    case "hello":
        fmt.Println(v3)
    case "":
        v21 := GenS2{F0: 42, F1: 453}
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 float64 = (-2.25 - (v3 + v3))
            var v24 int16 = 1000
            v2 = v2
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v25 := GenS5{F0: 18446744073709551615, F1: 42}
    v2 = (127 <= -39)
    var v26 uint = 4294967295
    switch 2147483647 {
    case -9223372036854775808:
        delete(v1, 4294967295)
    case 934:
        var v27 float32 = float32(v3)
        _ = v27
    }
    v28 := GenS6{F0: 0, F1: 4294967295}
    if (v26 == v26) {
        if v2 {
            v29 := v1[961]
            if (v3 >= (0.0 + 3.14159)) {
                v26 = 167
                delete(v1, 42)
            } else {
                for fi30 := 0; fi30 < 7; fi30++ {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 int16 = 42
                        fmt.Println(fi31, v32)
                        v33 := GenS7{F0: ("hello" + ("a" + "a"))}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = fi30
                }
                for fi34 := 0; fi34 < 11; fi34++ {
                    var v35 uint16 = 0
                    var v36 bool = v29
                    var v37 int32 = -1
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            v3 -= 1e10
            _ = v29
        }
    }
    v38 := GenS8{F0: v26, F1: -9223372036854775808, F2: 495}
    v39 := GenS6{F0: uint64(v26), F1: 200}
    fmt.Println(v2, v26)
    var v40 uint = (v26 % (((v26 | 523)) | 1))
    v41 := []int{0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
