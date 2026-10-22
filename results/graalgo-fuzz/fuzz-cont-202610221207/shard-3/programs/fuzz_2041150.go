package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 bool; F2 int64 }
type GenS4 struct { F0 int32; F1 int64; F2 uint }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 float64 = (0.0 + -2.25)
    v4 := []string{"hello", "hello", "a", ("hello" + "a")}
    var v5 float32 = (-2.25 + ((1.5 + (float32(1.5) + float32(float64(float32(v3))))) + 1.5))
    var v6 int = 0
    if (true || (true && (false || false))) {
        v6 = (v6 + (v6 ^ -9223372036854775808))
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                v6 -= (v6 >> ((0) & 63))
                _ = fi8
            }
            switch "hello" {
            case "foo bar":
                v3 += (v3 / (float64(((3.14159 + v5) + ((v5 - ((v5 + (0.0 - float32(float64(v5)))) - 0.0)) - float32(float64(0.1))))) * 0.0))
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi7
        }
    }
    v9 := GenS1{F0: 2147483647, F1: v6, F2: 792}
    switch ("foo bar" + "a") {
    case "foo bar":
        v6 = v6
    case "hello":
        _ = len(v4)
    default:
        switch (v6 & v6) {
        case 42:
            fmt.Println(v6)
        case 595:
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 string = "foo bar"
                fmt.Println(v5)
                for fi12 := 0; fi12 < 13; fi12++ {
                    if (false || ((v3 - float64(1.5)) == (0.1 - (float64(v5) / (v3 / float64(v5)))))) {
                        var v13 uint16 = 65535
                        _ = v13
                    }
                    v6 = (v6 ^ (fi10 - (fi12 / ((fi10) | 1))))
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        }
    }
    if (!true) {
        v5 = (float32(float64(1e10)) / 1.0)
        if false {
            var v14 float32 = float32(float64(v5))
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 uint = 569
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                if ("hello" <= "\n\t") {
                    fmt.Println(v3, v5)
                } else {
                    switch ("a" + ("foo bar" + "a")) {
                    case "":
                        switch "foo bar" {
                        case "hello":
                            var v18 int = (fi17 ^ (fi17 + (^(fi17))))
                            _ = v18
                        case "foo bar":
                            if (false && true) {
                                if ((true || (!(false && true))) && (255 < 2)) {
                                    var v19 uint32 = 2147483648
                                    _ = v19
                                } else {
                                    var v20 int = v6
                                    _ = v20
                                }
                            } else {
                                v21 := GenS2{F0: (("a" + "a") + "foo bar"), F1: 806}
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    v23 := GenS2{F0: "foo bar", F1: 4294967295}
                                    if false {
                                        v6 = (fi17 >> ((78) & 63))
                                        var v24 bool = true
                                        v1[(fi17 % ((fi17) | 1))] = ((v5 * float32((v3 / v3))) == 3.14159)
                                        _ = v24
                                    }
                                    var v25 float64 = v3
                                    v26 := GenS3{F0: 311, F1: false, F2: -9223372036854775808}
                                    v5 += 1.5
                                    _ = fi22
                                    _ = v23
                                    _ = v25
                                    _ = v26
                                }
                                var v27 float64 = v3
                                _ = v21
                                _ = v27
                            }
                        default:
                            switch "\n\t" {
                            case "a":
                                v5 -= float32(-2.25)
                            case "hello":
                                v3 += 0.1
                            case "foo bar":
                                v3 = 1e10
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    case "a":
                        v4 = append(v4, "")
                    default:
                        v6 = (fi17 % ((571) | 1))
                    }
                }
                _ = fi17
            }
            var v28 float64 = float64(1e10)
            var v29 int64 = -617
            _ = v28
            _ = v29
        }
    } else {
        switch 9223372036854775807 {
        case -2147483648:
            var v30 float64 = v3
            _ = v30
        case -34:
            if (1 != 2147483647) {
                var v31 int8 = -1
                var v32 uint32 = 751
                _ = v31
                _ = v32
            } else {
                v5 -= (1.5 - v5)
                if true {
                    v3 -= float64(v5)
                } else {
                    var v33 uint8 = 128
                    if false {
                        _ = len(v4)
                        v34 := GenS4{F0: -368, F1: int64(v6), F2: 397}
                        var v35 uint32 = 42
                        _ = v34
                        _ = v35
                    }
                    v3 = float64(1.5)
                    _ = v33
                }
            }
        case 128:
            v4 = append(v4, ("\n\t" + ""))
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            fmt.Println(fi36)
            v3 += v3
            _ = fi36
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            _ = len(v4)
            _ = fi37
        }
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 int8 = 66
        v2[(v6 / ((-538) | 1))] = ((("\n\t" + "foo bar") + "a") + "foo bar")
        v3 = float64((0.0 - float32(((v3 + 0.1) - (((1e10 / ((0.1 / (float64(v5) * 3.14159)) + (v3 - float64(v5)))) - (0.1 + (float64(-2.25) / v3))) * -2.25)))))
        _ = fi38
        _ = v39
    }
    v6 = v6
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 int64 = 0
        v5 = ((3.14159 - -2.25) + v5)
        for fi42 := 0; fi42 < 4; fi42++ {
            v41 -= -9223372036854775808
            if (((2 < 2) && ((255 >= 42) || ((292 >= 216) || true))) || (255 == 656)) {
                v43 := GenS1{F0: -1, F1: 2147483647, F2: (fi40 * fi42)}
                v44 := GenS5{F0: 42, F1: false}
                v6 += (^(v6))
                _ = v43
                _ = v44
            } else {
                var v45 string = "\n\t"
                _ = v45
            }
            var v46 int16 = 1000
            v47 := GenS3{F0: (v46 / (((v46 << ((65535) & 63))) | 1)), F1: (122 < (-(fi40))), F2: 42}
            if (fi40 == fi40) {
                fmt.Println(fi42, v46, v3)
                switch "" {
                case "a":
                    v3 -= v3
                }
                v6 = v6
            } else {
                switch "a" {
                case "":
                    fmt.Println(v41, v3)
                case "a":
                    switch (fi42 | 637) {
                    case -691:
                        var v48 int64 = v41
                        _ = v48
                    }
                default:
                    if (54 >= 18446744073709551615) {
                        v5 += ((v5 / v5) / 3.14159)
                        v46 = 32767
                        v3 += float64((1e10 * ((v5 / 3.14159) + 3.14159)))
                    }
                }
            }
            _ = fi42
            _ = v46
            _ = v47
        }
        _ = fi40
        _ = v41
    }
    var v49 bool = true
    if len(v4) > 0 {
        v4[0] = "a"
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    fmt.Println(v5, v6, v6)
    if ("hello" < "hello") {
        v6 = (v6 >> ((0) & 63))
        var v50 float64 = v3
        _ = v50
    } else {
        var v51 string = "foo bar"
        v49 = (-37 <= 100)
        for fi52 := 0; fi52 < 10; fi52++ {
            if (!(65535 == 255)) {
                for fi53 := 0; fi53 < 18; fi53++ {
                    v51 = "foo bar"
                    if v49 {
                        var v54 int8 = 1
                        fmt.Println(v49, v3, v51)
                        var v55 uint = 42
                        _ = v54
                        _ = v55
                    }
                    _ = fi53
                }
            }
            fmt.Println(v6, v49, v5)
            v5 -= v5
            _ = fi52
        }
        _ = v51
    }
    if v49 {
        v56 := GenS6{F0: 65535, F1: 1}
        _ = v56
    }
    if v49 {
        v49 = false
        fmt.Println(v3, v5, v3)
    }
    var v57 uint16 = 165
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v57
    _ = v6
    _ = v9
}
