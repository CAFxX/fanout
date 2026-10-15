package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int64) int32 {
    if (1000 >= 0) {
        fmt.Println(p0, p0)
        p0 = (p0 ^ (p0 / ((4294967295) | 1)))
        v5 := GenS1{F0: (p1 | -1)}
        _ = v5
    }
    p0 = (p0 & (p0 * (p0 << ((255) & 63))))
    if true {
        p1 += (^(p1))
    }
    v6 := GenS2{F0: 1, F1: int64(p0)}
    if (!(-9223372036854775808 < 694)) {
        switch 2147483647 {
        case 32768:
            var v7 int16 = -32768
            _ = v7
        case -307:
            v8 := GenS2{F0: 2147483647, F1: 410}
            _ = v8
        case 2147483647:
            v9 := GenS3{F0: 724, F1: 2147483647}
            _ = v9
        }
    } else {
        var v10 int = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}
func genFn1(p0 bool, p1 bool) int {
    v11 := []int{42, 127, 1, 1}
    switch -1 {
    case 829:
        if ((0 != 1) || (!(2 > 0))) {
            _ = len(v11)
            if len(v11) > 0 {
                v11[0] = 897
            }
            if len(v11) > 0 {
                v11[0] = 128
            }
        } else {
            switch "hello" {
            case "foo bar":
                if p0 {
                    p0 = false
                    switch "hello" {
                    case "hello":
                        var v12 string = "a"
                        _ = v12
                    case "foo bar":
                        v13 := GenS4{F0: 31}
                        _ = v13
                    case "\n\t":
                        for fi14 := 0; fi14 < 10; fi14++ {
                            p1 = p1
                            v15 := GenS5{F0: ("hello" + ("\n\t" + "foo bar"))}
                            var v16 string = "a"
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    default:
                        switch "hello" {
                        case "a":
                            if false {
                                var v17 uint32 = 463
                                if (63 < v17) {
                                    var v18 float32 = -2.25
                                    fmt.Println(v18)
                                    _ = v18
                                } else {
                                    v19 := GenS6{F0: 0, F1: 28}
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    switch 127 {
                                    case 416:
                                        p0 = false
                                    case -2147483648:
                                        switch "hello" {
                                        case "hello":
                                            for fi20 := 0; fi20 < 4; fi20++ {
                                                v21 := GenS2{F0: fi20, F1: -775}
                                                if p1 {
                                                    v17 -= v17
                                                    var v22 float64 = (-2.25 * (-2.25 / -2.25))
                                                    _ = v22
                                                }
                                                v17 += 4294967295
                                                _ = fi20
                                                _ = v21
                                            }
                                        case "":
                                            var v23 bool = false
                                            _ = v23
                                        case "foo bar":
                                            p1 = (!p1)
                                        }
                                    default:
                                        p1 = false
                                    }
                                    _ = v19
                                }
                                _ = v17
                            } else {
                                for fi24 := 0; fi24 < 7; fi24++ {
                                    p0 = (259 == 9223372036854775807)
                                    _ = fi24
                                }
                                v25 := GenS7{F0: 32767, F1: 100, F2: 588}
                                v26 := GenS8{F0: 1}
                                _ = v25
                                _ = v26
                            }
                        case "":
                            v11 = append(v11, 9223372036854775807)
                        case "foo bar":
                            var v27 uint64 = 861
                            _ = v27
                        }
                    }
                }
            case "hello":
                var v28 bool = p0
                _ = v28
            default:
                for fi29 := 0; fi29 < 18; fi29++ {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        p1 = true
                    case 32767:
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case -958:
                        if len(v11) > 0 {
                            v11[0] = (fi29 / (((fi29 - 4294967295)) | 1))
                        }
                    }
                    if (!p0) {
                        fmt.Println(p0)
                    }
                    _ = fi29
                }
            }
            _ = len(v11)
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 32768
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 uint8 = 0
    var v3 float64 = float64(float32(v1))
    var v4 uint16 = 1
    if (true || false) {
        v3 = v1
        fmt.Println(v4)
    }
    var v30 uint32 = 853
    v31 := GenS9{F0: true, F1: 1}
    var v32 int64 = -9223372036854775808
    if true {
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 uint8 = (^(v2))
            var v35 int64 = -917
            v36 := []int{}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        if false {
            var v37 bool = false
            _ = v37
        } else {
            switch "foo bar" {
            case "":
                v3 = 1.5
            case "\n\t":
                v30 += (^(v30))
            case "a":
                v3 += ((v3 / ((3.14159 + v1) - (v3 * float64(3.14159)))) * ((float64(1e10) + 3.14159) / v1))
            }
            v38 := GenS10{F0: 2}
            v39 := []int{-9223372036854775808, genFn1((867 >= 61), (false || (255 > (^(v2))))), -1, 42}
            _ = v38
            _ = v39
        }
    } else {
        fmt.Println(v2)
        v40 := []int{32768, 479, 0, 32767}
        if true {
            fmt.Println(v3, v1, v1)
            fmt.Println(v4, v2)
        }
        _ = v40
    }
    fmt.Println(v3)
    var v41 int32 = int32(v2)
    fmt.Println(v30)
    v2 = v2
    if (v4 >= 1) {
        v2 = 2
        v42 := GenS6{F0: 100, F1: 1}
        _ = v42
    }
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
}
