package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []int{-2147483648}
    var v2 int32 = 0
    v3 := make(map[string]uint32)
    var v4 int32 = (-(v2))
    var v5 int64 = 42
    v6 := []uint32{130, 32}
    if (true || (12 > 255)) {
        if ((!((false || ((false || (-97 > -9223372036854775808)) || (("hello" + "foo bar") >= "hello"))) && ((1e10 * (-2.25 - (0.1 / float32(-2.25)))) != 3.14159))) || ((^(v5)) >= (v5 - 886))) {
            fmt.Println(v2, v4)
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 string = ""
                _ = fi7
                _ = v8
            }
            if false {
                if true {
                    fmt.Println(v4)
                } else {
                    if true {
                        v6 = append(v6, 1)
                    }
                    v9 := GenS1{F0: 0, F1: 18446744073709551615, F2: 255}
                    switch (("" + "\n\t") + "a") {
                    case "foo bar":
                        v2 -= (v2 ^ (^(v4)))
                    default:
                        var v10 uint16 = 943
                        _ = v10
                    }
                    _ = v9
                }
                if false {
                    var v11 int32 = (v2 / (((v4 | (v4 | v2))) | 1))
                    _ = v11
                }
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 uint8 = 42
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 int64 = (v5 / ((v5) | 1))
                _ = v14
            }
        }
    } else {
        if (false && true) {
            for fi15 := 0; fi15 < 4; fi15++ {
                v5 = v5
                switch ("hello" + ("\n\t" + "foo bar")) {
                case "foo bar":
                    v5 -= v5
                }
                _ = fi15
            }
        }
    }
    switch 90 {
    case 918:
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    default:
        var v16 int64 = v5
        _ = v16
    }
    if false {
        v4 -= (v2 | 1)
    }
    var v17 float64 = (3.14159 + 0.1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        switch "hello" {
        case "":
            if false {
                fmt.Println(v17, v17, v17)
                var v18 float32 = (((0.1 / ((0.1 / 3.14159) + float32((v17 + v17)))) + 1e10) - ((1e10 * 1.5) * 0.0))
                v17 = (v17 - v17)
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v2, v4)
            _ = len(v3)
            var v20 uint64 = 158
            _ = fi19
            _ = v20
        }
        v5 += v5
    }
    switch "hello" {
    case "a":
        var v21 int16 = -1
        _ = v21
    case "":
        if (1000 > 42) {
            var v22 int64 = v5
            v2 = (v4 * (-(v4)))
            switch "\n\t" {
            case "a":
                fmt.Println(v2, v17, v5)
            case "\n\t":
                if (("" + ("a" + ("" + ((("a" + "hello") + "\n\t") + ("foo bar" + "hello"))))) >= "hello") {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 int32 = (^(v4))
                        _ = fi23
                        _ = v24
                    }
                }
            case "hello":
                var v25 string = ("\n\t" + "hello")
                _ = v25
            }
            _ = v22
        } else {
            v3[(("hello" + "\n\t") + (("\n\t" + "") + ("foo bar" + "\n\t")))] = 0
            fmt.Println(v2, v4)
        }
    default:
        v1 = append(v1, 32767)
    }
    var v26 float64 = ((((v17 - ((-2.25 - (float64(((-2.25 * float32(v17)) * -2.25)) * float64(0.0))) / -2.25)) / ((v17 + ((((float64(0.1) + (v17 + 0.0)) + 0.0) - ((1e10 - v17) - (v17 + 0.1))) * v17)) + float64(1.5))) - (v17 + v17)) - (-2.25 / 3.14159))
    v5 = v5
    switch "hello" {
    case "":
        fmt.Println(v5)
    case "foo bar":
        v3["foo bar"] = 2147483647
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 float64 = v26
        var v29 uint32 = 2
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int16 = 32767
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v5, v26)
        _ = fi31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint64 = 0
        switch "" {
        case "":
            for fi34 := 0; fi34 < 16; fi34++ {
                if (false && false) {
                    v35 := GenS2{F0: fi32, F1: 24, F2: 351}
                    _ = v35
                } else {
                    v36 := GenS3{F0: 128, F1: 42, F2: true}
                    v17 = (v17 - 0.1)
                    _ = v36
                }
                var v37 int8 = -1
                v2 -= 42
                if (((^(v33)) == (v33 >> ((817) & 63))) || (34 != 128)) {
                    fmt.Println(v4, fi34, fi34)
                    v4 -= v4
                } else {
                    if false {
                        v2 = (v4 * (v4 >> ((102) & 63)))
                    }
                    var v38 int8 = (v37 ^ (^(v37)))
                    _ = len(v3)
                    if ((v26 + (v17 * 1e10)) == v17) {
                        v39 := GenS3{F0: -1, F1: 266, F2: (!((-786 != (fi32 + 42)) && ((0.0 <= ((3.14159 + 0.1) + float32(v17))) || (2147483648 != 2147483648))))}
                        _ = v39
                    }
                    _ = v38
                }
                switch ("\n\t" + "\n\t") {
                case "":
                    if (85 <= 65535) {
                        v37 -= (^(v37))
                        v3[("hello" + "hello")] = 2
                    }
                    var v40 uint = 4294967295
                    _ = v40
                }
                switch "foo bar" {
                case "a":
                    v4 += v4
                    var v41 uint8 = 255
                    var v42 uint16 = 0
                    v37 = 0
                    switch "hello" {
                    case "":
                        v41 = (v41 / (((^(v41))) | 1))
                    case "foo bar":
                        v43 := GenS4{F0: 18446744073709551615}
                        _ = v43
                    }
                    _ = v41
                    _ = v42
                case "hello":
                    v44 := GenS3{F0: fi34, F1: 434, F2: false}
                    _ = v44
                case "":
                    v45 := GenS3{F0: (fi32 & -503), F1: 255, F2: false}
                    _ = v45
                default:
                    v46 := GenS5{F0: (v33 / ((v33) | 1))}
                    _ = v46
                }
                _ = fi34
                _ = v37
            }
        default:
            v1 = append(v1, -2147483648)
        }
        fmt.Println(v2, fi32)
        _ = fi32
        _ = v33
    }
    v30 += (v30 / (((v30 + v30)) | 1))
    fmt.Println(v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
