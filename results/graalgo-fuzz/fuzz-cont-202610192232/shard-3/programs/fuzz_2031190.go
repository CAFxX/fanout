package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
type GenS4 struct { F0 int16; F1 uint }

func main() {
    var v1 uint = 545
    v2 := make(map[uint32]bool)
    var v3 int64 = -735
    switch -155 {
    case 32767:
        var v4 uint32 = 436
        _ = v4
    case 2147483647:
        var v5 int16 = -660
        _ = v5
    default:
        v1 = 580
    }
    _ = len(v2)
    if (false && false) {
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 12; fi7++ {
                v3 = (v3 % (((v3 - 1)) | 1))
                v1 = v1
                if false {
                    var v8 uint8 = 128
                    v9 := GenS1{F0: fi6}
                    var v10 uint16 = 443
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    v3 += v3
                    v1 = 808
                    if ((true && ("hello" >= "\n\t")) && true) {
                        var v11 int8 = 0
                        var v12 uint16 = 0
                        if false {
                            v13 := GenS1{F0: fi6}
                            var v14 string = ("hello" + "hello")
                            _ = v13
                            _ = v14
                        } else {
                            var v15 float32 = 1e10
                            if true {
                                var v16 float32 = v15
                                v17 := GenS2{F0: (v1 | v1)}
                                delete(v2, 739)
                                _ = v16
                                _ = v17
                            }
                            v1 -= (^(v1))
                            _ = v15
                        }
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            var v18 bool = true
                            var v19 float32 = float32((float64(float32(3.14159)) - float64((0.1 * 0.1))))
                            _ = v18
                            _ = v19
                        } else {
                            v20 := GenS2{F0: v1}
                            var v21 int64 = (v3 + 42)
                            _ = v20
                            _ = v21
                        }
                        switch (fi6 ^ (fi7 * fi6)) {
                        case 148:
                            delete(v2, 4294967295)
                        case 894:
                            var v22 float32 = ((0.0 / (-2.25 * (-2.25 * -2.25))) / -2.25)
                            _ = v22
                        case 4294967295:
                            var v23 float32 = 1.5
                            v24 := []int{int(v3), 0}
                            _ = len(v2)
                            _ = v23
                            _ = v24
                        default:
                            var v25 bool = ((-295 < 0) || false)
                            _ = v25
                        }
                        if (true && true) {
                            if (-2.25 != (-2.25 - (1.5 + 1e10))) {
                                var v26 uint8 = 128
                                v1 -= (v1 & 65535)
                                _ = v26
                            } else {
                                v27 := GenS1{F0: (fi7 + 384)}
                                v28 := []int{-1, 4294967295}
                                _ = len(v2)
                                fmt.Println(v3, fi6, v3)
                                _ = v27
                                _ = v28
                            }
                        } else {
                            var v29 float32 = 0.0
                            v30 := GenS3{F0: 595, F1: false, F2: v1}
                            _ = v29
                            _ = v30
                        }
                    }
                }
                _ = fi7
            }
            _ = fi6
        }
        var v31 int8 = 0
        v32 := []int{127, 231, -9223372036854775808, 128}
        _ = v31
        _ = v32
    } else {
        v1 -= 31
        v33 := GenS3{F0: 42, F1: true, F2: (v1 - v1)}
        _ = v33
    }
    if (true && true) {
        var v34 int32 = 2147483647
        v34 = (v34 - v34)
        _ = v34
    } else {
        v35 := GenS4{F0: -980, F1: 831}
        var v36 int8 = -128
        _ = v35
        _ = v36
    }
    var v37 int16 = 32767
    switch "\n\t" {
    case "foo bar":
        for fi38 := 0; fi38 < 1; fi38++ {
            v1 -= (v1 + (v1 | v1))
            _ = fi38
        }
    case "a":
        v3 -= v3
    default:
        fmt.Println(v37, v1, v3)
    }
    v1 += (v1 | (^(v1)))
    v3 -= v3
    v39 := GenS3{F0: 0, F1: (489 < 1), F2: 0}
    var v40 int32 = 42
    var v41 uint = 65535
    v3 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v41
}
