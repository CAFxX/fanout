package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    var v1 uint8 = 127
    v2 := []int{32767, -798}
    var v3 int16 = -1
    fmt.Println(v1, v3)
    v4 := GenS1{F0: 306, F1: 4294967295}
    v3 = v3
    switch (("foo bar" + "hello") + "\n\t") {
    case "\n\t":
        var v5 uint32 = 1
        _ = v5
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                v3 = v3
                v3 = 0
                fmt.Println(v3, fi6)
                _ = fi7
            }
            _ = fi6
        }
    default:
        var v8 float64 = (0.0 - (1.5 + 3.14159))
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        _ = len(v2)
        _ = fi9
    }
    if (((0.0 / 1.5) == (float64((1.5 - 0.0)) - (-2.25 - float64(1.5)))) && false) {
        switch "" {
        case "hello":
            var v10 bool = (42 < 298)
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v3 = -1
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float64 = (float64(0.1) + (0.0 * 1e10))
            _ = fi12
            _ = v13
        }
        if (false || false) {
            v2 = append(v2, fi11)
            v14 := GenS2{F0: "foo bar", F1: 2147483647, F2: 233}
            _ = v14
        }
        _ = fi11
    }
    var v15 uint8 = v1
    switch -983 {
    case 128:
        if ((0.1 + float64(float32((1.5 / (3.14159 * 1e10))))) >= 1.5) {
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi16, v3)
                for fi17 := 0; fi17 < 17; fi17++ {
                    v15 = v1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi17
                }
                fmt.Println(v15, v15, v3)
                _ = fi16
            }
        } else {
            v15 = (v15 - v15)
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int32 = -1
                v19 += 42
                v3 += (v3 + v3)
                _ = fi18
                _ = v19
            }
            var v20 float32 = 1e10
            _ = v20
        }
    case -442:
        switch 245 {
        case -237:
            v1 += 145
        case 846:
            fmt.Println(v15, v15)
        case -2147483648:
            if (49 > 42) {
                v2 = append(v2, 82)
            } else {
                if false {
                    var v21 int16 = v3
                    _ = v21
                } else {
                    var v22 int64 = -9223372036854775808
                    var v23 string = "a"
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v23 = ("hello" + v23)
                        if ((!true) || ((255 != 254) || true)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v25 int64 = v22
                            _ = v25
                        } else {
                            v2 = append(v2, 1)
                            if ((!((65535 > 2147483648) && (190 <= 179))) && (3.14159 <= (1e10 / (1e10 / float64(1e10))))) {
                                if true {
                                    fmt.Println(v15, v1)
                                }
                                v23 = ("foo bar" + ("foo bar" + v23))
                                v26 := GenS1{F0: v22, F1: 255}
                                _ = v26
                            } else {
                                fmt.Println(fi24, fi24)
                                var v27 bool = ((!false) || (((v22 + 42) < v22) && true))
                                _ = v27
                            }
                            switch (v23 + "foo bar") {
                            case "foo bar":
                                switch v23 {
                                case "hello":
                                    var v28 float32 = float32(-2.25)
                                    _ = v28
                                case "a":
                                    var v29 uint64 = 447
                                    _ = v29
                                default:
                                    if ((832 <= 85) || (((0.1 - 0.0) * -2.25) != (1.5 * ((0.1 / 1.0) + -2.25)))) {
                                        var v30 float64 = 0.1
                                        for fi31 := 0; fi31 < 2; fi31++ {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v32 := GenS3{F0: v23}
                                            fmt.Println(v22, v30, v3)
                                            _ = fi31
                                            _ = v32
                                        }
                                        if (65535 == uint(v1)) {
                                            v1 += (v1 * 2)
                                            for fi33 := 0; fi33 < 1; fi33++ {
                                                _ = len(v2)
                                                v1 = (v1 - v15)
                                                _ = fi33
                                            }
                                        }
                                        _ = v30
                                    } else {
                                        v22 = v22
                                    }
                                }
                            case "hello":
                                v23 = "\n\t"
                            case "\n\t":
                                fmt.Println(v3, v3)
                            }
                        }
                        _ = fi24
                    }
                    _ = v22
                    _ = v23
                }
            }
        }
    }
    fmt.Println(v15, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        v15 -= (v15 >> ((251) & 63))
        if false {
            for fi35 := 0; fi35 < 1; fi35++ {
                v3 -= v3
                switch (((((("" + "foo bar") + ("foo bar" + "hello")) + "foo bar") + "foo bar") + ("hello" + ("" + "\n\t"))) + "a") {
                case "hello":
                    if true {
                        v36 := GenS4{F0: 42, F1: fi35, F2: 43}
                        var v37 uint8 = (v1 - (v1 & (v1 % (((v1 - (v1 ^ 0))) | 1))))
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint16 = 2
                        fmt.Println(fi34, v15)
                        v2 = append(v2, fi35)
                        _ = v38
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = fi34
                    }
                default:
                    v3 = (v3 + 32767)
                }
                var v39 bool = false
                _ = fi35
                _ = v39
            }
        } else {
            var v40 int8 = 0
            _ = v40
        }
        _ = fi34
    }
    if (true || (-1 == 1)) {
        var v41 int = 0
        _ = v41
    }
    var v42 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
