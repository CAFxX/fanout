package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 int; F1 int8; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint32) int16 {
    v4 := []int{-83}
    var v5 float32 = (0.0 * ((float32(0.0) - (float32(float64(0.0)) / 0.1)) / 1.0))
    var v6 uint16 = 255
    if (286 <= 4294967295) {
        fmt.Println(v5, v6, v6)
    } else {
        switch -1 {
        case 32767:
            v7 := GenS1{F0: 156, F1: p0, F2: 1}
            _ = v7
        case 127:
            v8 := GenS1{F0: 128, F1: p0, F2: -9223372036854775808}
            _ = v8
        }
        v6 = v6
    }
    var v9 float64 = 0.0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    return 801
}

func main() {
    var v1 bool = false
    var v2 bool = (507 != 841)
    v3 := make(map[int]string)
    var v10 float32 = (0.0 - 0.0)
    v11 := GenS1{F0: 64, F1: 2, F2: 145}
    v10 = (v10 + (((v10 / v10) / (v10 - v10)) / ((0.1 - (v10 - float32(float64(-2.25)))) * 0.0)))
    var v12 float32 = -2.25
    var v13 float64 = 1.5
    for fi14 := 0; fi14 < 18; fi14++ {
        v13 = 1.5
        _ = fi14
    }
    if v2 {
        for fi15 := 0; fi15 < 8; fi15++ {
            switch (fi15 % ((32767) | 1)) {
            case 0:
                var v16 string = "foo bar"
                _ = v16
            }
            var v17 int16 = -32768
            _ = len(v3)
            _ = fi15
            _ = v17
        }
    } else {
        var v18 uint16 = 1
        _ = v18
    }
    v19 := GenS2{F0: v2, F1: 771, F2: v2}
    v20 := GenS3{F0: 1, F1: 0, F2: 127}
    var v21 int32 = 42
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v10, v10, v1)
        _ = fi22
    }
    switch -2147483648 {
    case -153:
        v23 := GenS4{F0: -32768}
        _ = v23
    case 861:
        v13 -= 1.5
    }
    if true {
        if (18446744073709551615 >= 255) {
            v24 := GenS5{F0: 2147483648, F1: 354, F2: 0}
            _ = v24
        } else {
            if v2 {
                v25 := GenS6{F0: 18446744073709551615, F1: 424}
                switch (("foo bar" + (("hello" + (("a" + "hello") + "")) + "a")) + "\n\t") {
                case "hello":
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v2 = ((0 >= 0) && v2)
                        _ = fi26
                    }
                case "a":
                    v12 -= v12
                }
                v27 := []int{42}
                _ = v25
                _ = v27
            }
            v28 := GenS2{F0: false, F1: 575, F2: (855 < 961)}
            v13 = v13
            _ = v28
        }
        switch 2147483647 {
        case -210:
            for fi29 := 0; fi29 < 20; fi29++ {
                for fi30 := 0; fi30 < 4; fi30++ {
                    if ((fi30 / ((fi30) | 1)) != 1) {
                        v10 = (v12 - 1e10)
                    }
                    switch "" {
                    case "":
                        switch 2147483647 {
                        case 32767:
                            v31 := []int{(fi30 + 42), (fi29 >> ((398) & 63)), -1, int(v21)}
                            _ = v31
                        case -917:
                            v10 = v12
                        default:
                            v32 := []int{}
                            _ = v32
                        }
                    }
                    _ = fi30
                }
                fmt.Println(v10, v13, fi29)
                _ = fi29
            }
        default:
            v33 := []int{4294967295, 0, -1}
            _ = v33
        }
    }
    var v34 int32 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
}
