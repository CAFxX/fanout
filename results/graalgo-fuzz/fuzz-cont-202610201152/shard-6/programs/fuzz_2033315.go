package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0(p0 int) int64 {
    v7 := GenS1{F0: 1, F1: -1, F2: 1}
    p0 -= p0
    switch "a" {
    case "":
        if false {
            switch "a" {
            case "\n\t":
                switch 4294967295 {
                case 1:
                    var v8 int64 = 473
                    _ = v8
                }
            }
        }
    case "hello":
        var v9 int8 = -1
        _ = v9
    default:
        var v10 float64 = (float64(-2.25) + float64((1e10 / (0.1 / 3.14159))))
        _ = v10
    }
    _ = p0
    _ = v7
    return -133
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    v3 := make(map[int]int)
    var v4 int16 = 662
    var v5 float64 = (((float64(-2.25) - (1e10 - float64((0.0 * 3.14159)))) / (float64(3.14159) / (float64(3.14159) - 1e10))) + (0.0 - 1.5))
    var v6 float64 = float64((1e10 + -2.25))
    switch "" {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        var v11 int64 = 42
        _ = v11
    default:
        v12 := GenS2{F0: true}
        _ = v12
    }
    v4 = (v4 % (((v4 + -1)) | 1))
    v13 := []int{32768, 32768, 42, 1}
    var v14 uint8 = 2
    var v15 int16 = 1
    var v16 int32 = -33
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS1{F0: -849, F1: -54, F2: (^(v14))}
        if true {
            if false {
                switch "a" {
                case "a":
                    if true {
                        var v19 int32 = -2147483648
                        v19 = 0
                        _ = v19
                    } else {
                        if len(v13) > 0 {
                            v13[0] = 128
                        }
                        fmt.Println(v2)
                        var v20 string = "foo bar"
                        _ = v20
                    }
                }
            } else {
                switch (fi17 * fi17) {
                case 1:
                    if true {
                        if false {
                            _ = len(v13)
                        } else {
                            _ = len(v1)
                        }
                    } else {
                        switch (fi17 >> ((18446744073709551615) & 63)) {
                        case -1:
                            v21 := GenS3{F0: 4294967295, F1: 4294967295, F2: v14}
                            _ = v21
                        case 128:
                            v22 := GenS1{F0: -1, F1: 0, F2: (v14 - v14)}
                            _ = v22
                        default:
                            switch fi17 {
                            case -245:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            case -1:
                                v23 := GenS4{F0: 301}
                                _ = v23
                            }
                        }
                        v24 := GenS2{F0: (2 != 1)}
                        switch (fi17 / ((-757) | 1)) {
                        case 194:
                            v25 := GenS5{F0: 4294967295}
                            _ = v25
                        case 128:
                            if false {
                                v3[fi17] = fi17
                            } else {
                                v26 := GenS6{F0: (fi17 + fi17)}
                                var v27 int16 = -211
                                _ = v26
                                _ = v27
                            }
                        case 326:
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v16)
                            }
                        }
                        _ = v24
                    }
                }
            }
        }
        var v28 uint16 = 32767
        _ = fi17
        _ = v18
        _ = v28
    }
    fmt.Println(v4)
    v29 := GenS7{F0: true}
    var v30 int16 = (v4 % ((42) | 1))
    fmt.Println(v6, v16)
    fmt.Println(v6)
    for fi31 := 0; fi31 < 12; fi31++ {
        v30 = 32767
        switch "a" {
        case "foo bar":
            var v32 string = ("" + ("foo bar" + ("\n\t" + ("foo bar" + ((("foo bar" + "foo bar") + "") + "a")))))
            _ = v32
        }
        if (false || (!(-2147483648 < (fi31 ^ fi31)))) {
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 uint = 690
                var v35 uint64 = 588
                _ = fi33
                _ = v34
                _ = v35
            }
            if (0 < 2) {
                v16 -= (-(v16))
            }
        }
        _ = fi31
    }
    v6 = float64((-2.25 + float32((-2.25 / v6))))
    if true {
        fmt.Println(v2, v14)
        var v36 float64 = v6
        if false {
            for fi37 := 0; fi37 < 8; fi37++ {
                v14 = 255
                v38 := GenS8{F0: 2, F1: 107, F2: (v2 / (((v2 - (v2 | 0))) | 1))}
                _ = fi37
                _ = v38
            }
            v16 = 42
        } else {
            fmt.Println(v16, v5)
            v16 = -1
            v6 = (v6 / v6)
        }
        _ = v36
    } else {
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
