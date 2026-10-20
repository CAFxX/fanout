package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    v1 := []uint32{732}
    var v2 float64 = ((3.14159 - 3.14159) / 3.14159)
    v3 := []int{-1, 4294967295}
    v2 += (-2.25 / v2)
    var v4 uint64 = 204
    var v5 float32 = -2.25
    if true {
        if (!true) {
            switch "a" {
            case "":
                for fi6 := 0; fi6 < 20; fi6++ {
                    v4 = 703
                    _ = fi6
                }
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            case "a":
                v7 := GenS1{F0: -128}
                _ = v7
            default:
                v2 = ((v2 / (-2.25 - 0.1)) / v2)
            }
        }
        if true {
            var v8 uint = 42
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 -= v2
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 float64 = (v2 - (0.0 + (1e10 + float64(v5))))
            fmt.Println(v10, fi9)
            _ = fi9
            _ = v10
        }
    } else {
        if (1000 < 1) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v4, v5)
        } else {
            v11 := GenS1{F0: 42}
            var v12 uint32 = 2147483648
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v11
            _ = v12
        }
        if (-209 <= 1) {
            fmt.Println(v5)
            for fi13 := 0; fi13 < 3; fi13++ {
                for fi14 := 0; fi14 < 20; fi14++ {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "a" {
                        case "foo bar":
                            if ((-1 > 0) && true) {
                                v15 := GenS2{F0: uint32(v4), F1: true, F2: 1}
                                v5 = float32(1e10)
                                v1 = append(v1, 124)
                                var v16 int64 = 42
                                v4 += 2
                                _ = v15
                                _ = v16
                            }
                        default:
                            v5 -= float32(float64(1.5))
                        }
                        v5 = 1.5
                    }
                    _ = fi14
                }
                if true {
                    v17 := GenS3{F0: 1, F1: 4294967295}
                    v18 := GenS2{F0: 1, F1: (127 <= uint8(v4)), F2: 1}
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v4, fi13)
                    v5 += (float32((0.0 * (v2 + (v2 * v2)))) - ((float32(v2) + ((v5 + 1e10) + v5)) * 1e10))
                    var v19 string = ""
                    _ = v19
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 int = fi20
                    _ = fi20
                    _ = v21
                }
                _ = fi13
            }
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 0
    if len(v1) > 0 {
        v1[0] = 2
    }
    switch "" {
    case "foo bar":
        var v25 int64 = -189
        _ = v25
    case "a":
        switch "a" {
        case "a":
            if true {
                var v26 bool = (1 > 81)
                _ = v26
            }
        }
    default:
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: 122}
            switch (fi27 * (fi27 | (fi27 | (^(fi27))))) {
            case 878:
                fmt.Println(fi27, v2, v5)
            case 4294967295:
                v4 += (v4 & (v4 | (v4 / ((1) | 1))))
            default:
                v29 := GenS1{F0: 100}
                _ = v29
            }
            _ = fi27
            _ = v28
        }
    }
    var v30 float32 = 1.5
    if false {
        if ((127 <= 128) && (v30 < v30)) {
            v4 -= v4
            fmt.Println(v24, v4, v5)
            var v31 int8 = 0
            _ = v31
        }
    }
    switch 128 {
    case -910:
        v24 += 65535
    case 32768:
        var v32 bool = true
        _ = v32
    case 127:
        v2 = v2
    }
    var v33 int32 = 1
    if len(v1) > 0 {
        v1[0] = 458
    }
    _ = len(v3)
    var v34 int16 = -774
    v35 := GenS2{F0: 2147483647, F1: (65535 == 2147483647), F2: v34}
    var v36 uint = 146
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
