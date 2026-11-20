package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 uint16, p1 string) uint32 {
    var v5 string = p1
    for fi6 := 0; fi6 < 2; fi6++ {
        switch "\n\t" {
        case "a":
            var v7 int64 = 0
            _ = v7
        case "hello":
            if (-644 < 27) {
                fmt.Println(fi6, p1, p0)
                v8 := GenS1{F0: 253}
                _ = v8
            } else {
                switch (fi6 | 0) {
                case -2147483648:
                    var v9 float64 = -2.25
                    _ = v9
                case 32768:
                    var v10 uint64 = 42
                    _ = v10
                case 127:
                    v11 := GenS2{F0: 1}
                    _ = v11
                }
            }
        default:
            p1 = (p1 + p1)
        }
        v12 := GenS2{F0: 4294967295}
        fmt.Println(v5, p1)
        _ = fi6
        _ = v12
    }
    v13 := GenS3{F0: 42, F1: (!false), F2: 43}
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "\n\t":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int32 = 1
                _ = fi14
                _ = v15
            }
        case "foo bar":
            var v16 float32 = 3.14159
            _ = v16
        }
    }
    if (1000 > -866) {
        v17 := []int{128}
        var v18 uint16 = (p0 % ((p0) | 1))
        if (3.14159 > -2.25) {
            switch v5 {
            case "\n\t":
                for fi19 := 0; fi19 < 4; fi19++ {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v21 := GenS1{F0: 101}
                        v22 := GenS2{F0: 18446744073709551615}
                        if (1 < -26) {
                            v18 -= p0
                            switch (^(fi19)) {
                            case 32767:
                                p0 -= (v18 & p0)
                            case 4294967295:
                                v5 = "hello"
                            case 1:
                                var v23 int64 = 0
                                _ = v23
                            }
                        } else {
                            v24 := GenS4{F0: -761}
                            _ = v24
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    var v25 int16 = 42
                    var v26 uint8 = 127
                    _ = fi19
                    _ = v25
                    _ = v26
                }
            default:
                var v27 float32 = float32(float64(3.14159))
                _ = v27
            }
            var v28 string = v5
            _ = v28
        } else {
            p1 = ("" + "")
            v5 = ("a" + "")
        }
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v5
    return 2147483648
}
func genFn1() int16 {
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int = 128
        var v31 int8 = 42
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS5{F0: 448}
    if ((!false) && (false || (!false))) {
        var v33 float64 = 1e10
        v33 -= ((0.0 * (v33 + v33)) - v33)
        _ = v33
    }
    var v34 int8 = 77
    _ = v32
    _ = v34
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    v3 := make(map[string]int)
    v4 := []string{("hello" + "a"), "hello", ""}
    fmt.Println(v2)
    if ("a" >= "\n\t") {
        if (!false) {
            fmt.Println(v2)
            if (!true) {
                var v35 bool = (9223372036854775807 == -174)
                var v36 bool = v35
                _ = v35
                _ = v36
            }
            v2 += (v2 * 1)
        }
    }
    var v37 int64 = 0
    for fi38 := 0; fi38 < 1; fi38++ {
        v2 -= v2
        v37 = -909
        if true {
            var v39 uint8 = 1
            switch (^(fi38)) {
            case 804:
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 float32 = 0.1
                    _ = fi40
                    _ = v41
                }
            case -9223372036854775808:
                v42 := GenS6{F0: 1, F1: 42, F2: v37}
                _ = v42
            }
            _ = v39
        }
        _ = fi38
    }
    v2 = (v2 ^ (^(v2)))
    v43 := GenS5{F0: v2}
    _ = len(v4)
    v37 = v37
    _ = len(v3)
    v44 := v1[9223372036854775807]
    for fi45 := 0; fi45 < 14; fi45++ {
        for fi46 := 0; fi46 < 7; fi46++ {
            v37 = (v37 - -989)
            if true {
                var v47 uint = 255
                v2 += (v2 & (v2 + (v2 / (((v2 ^ (v2 << ((255) & 63)))) | 1))))
                v48 := GenS3{F0: 42, F1: ((1000 < 508) && false), F2: -55}
                _ = v47
                _ = v48
            } else {
                var v49 bool = (2 > genFn0(872, ("a" + "")))
                v49 = ((911 < 42) && true)
                v49 = v49
                _ = v49
            }
            v37 += v37
            _ = fi46
        }
        switch "a" {
        case "a":
            v2 = v2
        case "hello":
            v44 += (^(v44))
        }
        v44 = (v44 % ((v44) | 1))
        _ = fi45
    }
    var v50 int32 = -1
    var v51 uint8 = 255
    _ = len(v1)
    if (((("\n\t" + ("foo bar" + "")) + "a") + ("\n\t" + (("foo bar" + (("hello" + (("\n\t" + ("hello" + ("" + "\n\t"))) + "a")) + "foo bar")) + ("a" + "hello")))) <= (("\n\t" + "") + ("hello" + "hello"))) {
        fmt.Println(v44)
        v4 = append(v4, (("hello" + "") + "a"))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v50
    _ = v51
}
