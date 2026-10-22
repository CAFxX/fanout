package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
func genFn0(p0 uint8, p1 string) uint8 {
    var v5 int16 = int16(p0)
    v5 += 1000
    _ = p0
    _ = p1
    _ = v5
    return (p0 << ((1) & 63))
}
func genFn1(p0 uint8, p1 string, p2 int64) uint {
    var v6 float32 = 3.14159
    var v7 int32 = 2147483647
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 float32 = (1.5 - (((0.0 + 0.1) - -2.25) + 0.1))
    var v2 uint = 1
    v3 := make(map[string]string)
    var v4 float32 = (v1 * ((v1 * 3.14159) + v1))
    switch ("foo bar" + ("" + "")) {
    case "":
        v1 = v1
    }
    if (false || ((!true) && true)) {
        var v8 float64 = 0.1
        if true {
            switch -9223372036854775808 {
            case 193:
                switch "\n\t" {
                case "\n\t":
                    v4 = v1
                }
            case -995:
                if (true || (-1 != 9223372036854775807)) {
                    v2 = (v2 | v2)
                    if (!false) {
                        v4 = 1.5
                        v8 -= v8
                        v3["a"] = "a"
                    } else {
                        if true {
                            if (("" < (("a" + "foo bar") + "\n\t")) || (875 >= 751)) {
                                fmt.Println(v2, v8)
                                v9 := GenS1{F0: 42, F1: true, F2: 255}
                                var v10 int8 = -128
                                _ = v10
                                _ = v9
                            } else {
                                v1 -= (((-2.25 + v4) / v4) - v4)
                                switch ((("\n\t" + "a") + "foo bar") + ("a" + ("hello" + ""))) {
                                case "hello":
                                    switch (("" + (("" + "") + "")) + "a") {
                                    case "hello":
                                        switch 128 {
                                        case 559:
                                            v11 := GenS1{F0: 9223372036854775807, F1: true, F2: 49}
                                            _ = v11
                                        case 32767:
                                            v4 = v4
                                        case -1:
                                            v3[("\n\t" + ("a" + ("a" + "\n\t")))] = "foo bar"
                                        default:
                                            v12 := GenS2{F0: 0, F1: -32768, F2: 255}
                                            _ = v12
                                        }
                                    }
                                case "foo bar":
                                    var v13 uint32 = 276
                                    _ = v13
                                case "a":
                                    v14 := GenS2{F0: 1, F1: -32768, F2: 0}
                                    _ = v14
                                }
                            }
                            if true {
                                v4 = (0.0 + v1)
                                var v15 int32 = -2147483648
                                _ = v15
                            } else {
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    v2 = 4294967295
                                    _ = fi16
                                }
                                if (-2147483648 != -2147483648) {
                                    var v17 bool = false
                                    var v18 string = "a"
                                    v4 = v4
                                    _ = v17
                                    _ = v18
                                } else {
                                    v19 := GenS2{F0: -1, F1: -659, F2: 376}
                                    v4 = v4
                                    _ = v19
                                }
                            }
                            for fi20 := 0; fi20 < 15; fi20++ {
                                switch "a" {
                                case "hello":
                                    v2 = 2
                                case "a":
                                    v1 += (((v1 * (1.5 * v1)) * (v4 + (v1 - ((((0.0 * ((-2.25 - v4) / v4)) / v1) / float32(float64(float32(float64(v1))))) - ((v4 / ((0.1 * ((v1 + v4) + v1)) - v4)) + (v4 * -2.25)))))) / v1)
                                }
                                var v21 uint16 = 255
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                }
            case 32768:
                v22 := GenS3{F0: -964, F1: 127}
                _ = v22
            default:
                var v23 float32 = v4
                _ = v23
            }
            v24 := GenS4{F0: true}
            _ = v24
        } else {
            var v25 string = (((("" + "hello") + "\n\t") + ("hello" + "\n\t")) + "")
            v25 = (v25 + v25)
            _ = v25
        }
        _ = v8
    } else {
        for fi26 := 0; fi26 < 1; fi26++ {
            switch fi26 {
            case 761:
                var v27 string = "\n\t"
                _ = v27
            default:
                if false {
                    var v28 float64 = ((-2.25 + (0.0 * (1.5 * (1e10 - (1e10 + -2.25))))) * (3.14159 / (1e10 - 0.1)))
                    var v29 uint8 = 117
                    fmt.Println(v28, fi26, v1)
                    _ = v28
                    _ = v29
                }
            }
            _ = fi26
        }
    }
    var v30 int = 1
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            if ((fi31 << (((^(v2))) & 63)) > (v30 >> ((genFn1(127, "foo bar", 149)) & 63))) {
                if false {
                    fmt.Println(v2, v2, fi31)
                }
                fmt.Println(fi31)
                var v33 int = v30
                v4 -= v1
                v34 := GenS2{F0: -492, F1: 0, F2: 32767}
                fmt.Println(v30)
                _ = v33
                _ = v34
            }
            _ = fi32
        }
        _ = fi31
    }
    var v35 int64 = -9223372036854775808
    if (0 <= 175) {
        switch (^(v30)) {
        case 439:
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := GenS2{F0: (v35 - (v35 % ((878) | 1))), F1: 32767, F2: uint16(v2)}
                _ = len(v3)
                v35 = (v35 << (((v2 / ((v2) | 1))) & 63))
                _ = fi36
                _ = v37
            }
        case 32767:
            if (false && (494 > 42)) {
                v4 -= (v1 + (v4 - 1e10))
            }
        default:
            for fi38 := 0; fi38 < 3; fi38++ {
                if true {
                    v4 += v1
                    if true {
                        v39 := GenS1{F0: -9223372036854775808, F1: true, F2: 128}
                        _ = v39
                    }
                }
                var v40 int16 = 0
                v2 = genFn1(133, "a", (v35 / ((408) | 1)))
                _ = fi38
                _ = v40
            }
        }
        var v41 uint8 = 2
        v42 := GenS5{F0: 4294967295}
        _ = v41
        _ = v42
    } else {
        v4 = v1
        switch (v30 % ((v30) | 1)) {
        case 972:
            if ("foo bar" > "\n\t") {
                for fi43 := 0; fi43 < 11; fi43++ {
                    if (-2.25 <= 0.0) {
                        var v44 float32 = ((0.1 - ((v4 + ((float32(((float64(v4) / float64(v4)) + ((1.5 - (3.14159 / 1.5)) / (1.5 * 0.1)))) * 1.5) * (1.5 * ((v4 / v1) * v4)))) - (-2.25 + (v1 - (-2.25 / -2.25))))) + v1)
                        v4 -= (1e10 / ((((((1e10 / v1) + 0.0) / 1.5) * v1) + v44) + 0.0))
                        var v45 int8 = 1
                        _ = v44
                        _ = v45
                    }
                    _ = fi43
                }
            } else {
                v4 += (v1 / v4)
                v46 := GenS5{F0: v30}
                _ = v46
            }
        case 4294967295:
            var v47 int8 = 127
            _ = v47
        default:
            var v48 float64 = (0.1 - float64(-2.25))
            _ = v48
        }
    }
    fmt.Println(v4, v2)
    switch v30 {
    case 295:
        v35 -= (v35 * int64(v30))
    case -1:
        var v49 bool = false
        _ = v49
    default:
        v2 = (v2 | 255)
    }
    for fi50 := 0; fi50 < 16; fi50++ {
        v4 = float32(float64(v1))
        v51 := []int{}
        _ = fi50
        _ = v51
    }
    v35 = v35
    v52 := GenS2{F0: (v35 * 0), F1: -260, F2: 1}
    switch (^(v30)) {
    case 127:
        v2 = 60
    case 32768:
        fmt.Println(v35, v4, v1)
    case 866:
        fmt.Println(v30, v35)
    }
    v53 := GenS3{F0: -32768, F1: 100}
    for fi54 := 0; fi54 < 3; fi54++ {
        var v55 uint16 = 42
        v56 := GenS1{F0: (v35 >> ((v2) & 63)), F1: true, F2: 160}
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 20; fi57++ {
        fmt.Println(v30, v35, v30)
        _ = fi57
    }
    switch (v30 - -326) {
    case 32767:
        if (255 == genFn1(127, "a", (v35 - int64(v2)))) {
            v58 := []int{-2147483648}
            _ = v58
        }
    case 42:
        fmt.Println(v30, v1)
    }
    v30 += (v30 & v30)
    var v59 uint32 = 855
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v52
    _ = v53
    _ = v59
}
