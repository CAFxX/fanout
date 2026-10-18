package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint }

func main() {
    v1 := []int{438, 127, -1}
    var v2 int32 = 711
    var v3 float64 = ((0.0 - 3.14159) / (3.14159 - 0.1))
    v4 := GenS1{F0: 25, F1: 255}
    var v5 int16 = 42
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            v5 = (v5 | (v5 | v5))
        } else {
            switch fi6 {
            case 495:
                var v7 float32 = 1.5
                _ = v7
            }
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(fi6, v2)
                _ = fi8
            }
            if len(v1) > 0 {
                v1[0] = (fi6 << ((0) & 63))
            }
        }
        _ = fi6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            var v10 int = (fi9 / (((fi9 | fi9)) | 1))
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 float32 = 1e10
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            if len(v1) > 0 {
                v1[0] = 9223372036854775807
            }
            fmt.Println(v5, v3)
            if (!(!((1 >= 14) || true))) {
                var v13 float64 = (0.1 * (0.0 / (v3 - 0.1)))
                if (42 <= 42) {
                    if true {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            _ = len(v1)
                            var v15 float64 = v3
                            _ = fi14
                            _ = v15
                        }
                        if len(v1) > 0 {
                            v1[0] = fi9
                        }
                    }
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v3 += 0.1
                        var v17 uint32 = 760
                        v18 := GenS1{F0: 128, F1: 42}
                        var v19 float64 = v13
                        var v20 int16 = (v5 - v5)
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                }
                _ = v13
            } else {
                var v21 int16 = -564
                if (-58 >= 42) {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint16 = 2
                        _ = len(v1)
                        _ = fi22
                        _ = v23
                    }
                }
                v1 = append(v1, 42)
                _ = v21
            }
        }
        _ = fi9
    }
    v24 := GenS2{F0: 275}
    fmt.Println(v3, v3, v3)
    switch 2147483647 {
    case -728:
        v25 := GenS1{F0: 42, F1: 189}
        _ = v25
    case -2147483648:
        var v26 uint8 = 42
        _ = v26
    case 127:
        v5 += 182
    }
    v1 = append(v1, -2147483648)
    var v27 uint32 = 65
    v27 = (v27 - v27)
    var v28 uint = 58
    v27 += v27
    v29 := GenS1{F0: 44, F1: 1}
    v2 = v2
    if ((1 != 1) || true) {
        if (42 == 127) {
            if ((3.14159 < -2.25) && false) {
                fmt.Println(v5)
                v3 = 0.1
            }
            if true {
                switch "\n\t" {
                case "hello":
                    var v30 uint8 = 128
                    _ = v30
                case "":
                    v2 -= (v2 - (v2 >> ((255) & 63)))
                case "a":
                    v5 = 1
                default:
                    for fi31 := 0; fi31 < 9; fi31++ {
                        var v32 uint = v28
                        v2 += 0
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 uint = (^(v28))
                v2 = int32(v5)
                _ = v33
            } else {
                v34 := GenS1{F0: 84, F1: 165}
                _ = v34
            }
        } else {
            if true {
                var v35 float32 = ((1.5 * 1.5) - (0.0 / -2.25))
                _ = v35
            }
        }
        switch 9223372036854775807 {
        case -867:
            fmt.Println(v28, v27, v5)
        }
    }
    fmt.Println(v27, v3)
    v36 := GenS3{F0: 1, F1: 817}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
