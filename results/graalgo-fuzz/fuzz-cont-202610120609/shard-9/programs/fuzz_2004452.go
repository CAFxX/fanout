package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int8; F2 int16 }
type GenS5 struct { F0 int; F1 int8 }
func genFn0() int64 {
    var v4 int16 = -528
    if ((false && false) && (v4 >= (v4 ^ v4))) {
        v4 += v4
        v5 := GenS1{F0: 630}
        _ = v5
    }
    v6 := []int{-363}
    v4 = (v4 | -940)
    _ = v4
    _ = v6
    return 712
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    v3 := make(map[int]int)
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS2{F0: 42}
        _ = fi7
        _ = v8
    }
    var v9 float64 = -2.25
    v10 := v1[-9223372036854775808]
    if false {
        v11 := []int{619}
        var v12 uint = v2
        _ = v11
        _ = v12
    }
    var v13 int16 = -1
    switch (v10 + v10) {
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int64 = -9223372036854775808
            for fi16 := 0; fi16 < 5; fi16++ {
                switch 1 {
                case 127:
                    v9 -= 0.0
                }
                var v17 int8 = -128
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if true {
            v19 := []int{-9223372036854775808, fi18, fi18, 32768}
            _ = v19
        } else {
            var v20 int32 = 947
            if (1.5 >= -2.25) {
                delete(v1, 42)
                v21 := GenS1{F0: 4294967295}
                _ = v21
            }
            v2 = (v2 / (((v2 << (((v2 % ((v2) | 1))) & 63))) | 1))
            _ = v20
        }
        _ = fi18
    }
    if (42 < 246) {
        switch 42 {
        case -839:
            for fi22 := 0; fi22 < 18; fi22++ {
                v2 = v2
                v23 := GenS1{F0: 18446744073709551615}
                if false {
                    v24 := []int{-586, -2147483648}
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
        case 8:
            v10 = ""
        }
        if true {
            v25 := GenS3{F0: (v13 & (v13 << ((961) & 63)))}
            v2 += (^(v2))
            _ = v25
        }
        if (1.5 >= (float32(v9) + -2.25)) {
            v9 -= float64((float32(((float64(0.1) / (0.0 / v9)) - 0.0)) / (((float32(float64(0.1)) / 1.5) - -2.25) + ((-2.25 * (1e10 - -2.25)) + 0.0))))
        }
    } else {
        fmt.Println(v9)
    }
    switch 924 {
    case 32768:
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 16; fi27++ {
                delete(v3, fi27)
                v2 = (^(v2))
                if (!(1e10 < v9)) {
                    switch v10 {
                    case "hello":
                        var v28 uint = (v2 % ((1) | 1))
                        _ = v28
                    case "\n\t":
                        var v29 int16 = (^(v13))
                        _ = v29
                    case "a":
                        switch v10 {
                        case "":
                            var v30 int64 = 42
                            _ = v30
                        case "a":
                            v13 = (v13 / ((32767) | 1))
                        }
                        var v31 string = "hello"
                        v32 := v1[(fi27 & 127)]
                        _ = v31
                        _ = v32
                    default:
                        var v33 int = -244
                        _ = v33
                    }
                } else {
                    var v34 int16 = v13
                    switch v10 {
                    case "\n\t":
                        var v35 uint = (v2 ^ v2)
                        _ = v35
                    case "a":
                        fmt.Println(fi26, v13)
                    }
                    fmt.Println(v9, v13, fi27)
                    _ = v34
                }
                _ = fi27
            }
            _ = fi26
        }
    case 127:
        v36 := []int{9223372036854775807, 705, 0}
        _ = v36
    default:
        v37 := GenS4{F0: 0, F1: 0, F2: (^(v13))}
        _ = v37
    }
    fmt.Println(v13)
    var v38 uint16 = 0
    var v39 bool = true
    v40 := GenS5{F0: -9223372036854775808, F1: -128}
    if (127 >= -1) {
        v41 := GenS4{F0: 2147483648, F1: 127, F2: v13}
        v13 = (v13 << (((^(v2))) & 63))
        _ = v41
    }
    var v42 float64 = v9
    var v43 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v9
}
