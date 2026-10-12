package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 int64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 42
    var v3 int32 = 0
    var v4 uint8 = (v1 % ((v1) | 1))
    v5 := []int{2147483647, 42, 2147483647}
    v5 = append(v5, 32768)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        if (9223372036854775807 == 1) {
            if (0.0 > 1e10) {
                var v6 string = "foo bar"
                _ = v6
            } else {
                v3 += (v3 * v3)
            }
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                v2 -= 2
                v8 := GenS1{F0: 1, F1: (v1 / (((v4 >> ((863) & 63))) | 1)), F2: 0}
                _ = fi7
                _ = v8
            }
        }
        _ = len(v5)
    } else {
        var v9 int8 = 103
        v1 = v1
        v1 = v4
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi10, v3, v1)
            case "foo bar":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "a":
                if false {
                    v3 -= 42
                    fmt.Println(v3, v1)
                    switch "foo bar" {
                    case "hello":
                        v3 += v3
                    case "a":
                        v12 := GenS1{F0: 42, F1: (v4 % ((183) | 1)), F2: -1}
                        v13 := GenS1{F0: 4294967295, F1: v1, F2: 1}
                        _ = v12
                        _ = v13
                    case "foo bar":
                        v14 := GenS1{F0: 0, F1: uint8(v2), F2: 908}
                        _ = v14
                    default:
                        fmt.Println(v2, v1)
                    }
                } else {
                    v15 := GenS2{F0: 1}
                    v1 += 19
                    if false {
                        v4 -= v1
                        var v16 string = "a"
                        switch fi10 {
                        case 127:
                            var v17 uint32 = 1
                            _ = v17
                        }
                        fmt.Println(v1, fi10)
                        if ((!false) && (1e10 != 1e10)) {
                            if true {
                                v18 := GenS3{F0: v3, F1: 42, F2: 325}
                                if (false || (float64(1e10) <= 0.1)) {
                                    v19 := GenS4{F0: 282, F1: 32767}
                                    v4 += (v1 | 225)
                                    _ = v19
                                }
                                if (!true) {
                                    if (false && false) {
                                        fmt.Println(v2, v16)
                                        if false {
                                            var v20 int = (fi10 * fi11)
                                            var v21 int64 = 0
                                            if false {
                                                v16 = ""
                                            }
                                            v22 := GenS3{F0: (v3 / ((v3) | 1)), F1: 0, F2: (v21 - 9223372036854775807)}
                                            v23 := GenS3{F0: (-(v3)), F1: 2147483647, F2: 1}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    if true {
                                        if false {
                                            var v24 int8 = -128
                                            var v25 string = v16
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        v16 = (v16 + ((((v16 + (("foo bar" + v16) + v16)) + v16) + v16) + ("foo bar" + ("\n\t" + (v16 + v16)))))
                                        switch "foo bar" {
                                        case "\n\t":
                                            if false {
                                                var v26 float32 = 0.1
                                                v4 = v1
                                                var v27 int16 = -255
                                                _ = v26
                                                _ = v27
                                            }
                                        case "hello":
                                            if true {
                                                if (!true) {
                                                    if (!(fi10 == (fi11 + fi11))) {
                                                        if true {
                                                            fmt.Println(v1)
                                                            v4 = (v1 + 1)
                                                            if (-1 < 127) {
                                                                if (18446744073709551615 == 1) {
                                                                    var v28 string = v16
                                                                    _ = v28
                                                                }
                                                                v29 := GenS1{F0: 226, F1: v4, F2: -877}
                                                                var v30 float32 = (0.1 / ((1.5 * (0.1 - ((float32((float64(0.1) / 0.1)) * 1.5) + (0.1 * float32(float64((-2.25 / 0.1))))))) + 1.5))
                                                                _ = v29
                                                                _ = v30
                                                            }
                                                            v4 = 232
                                                            v1 += (^(v4))
                                                        }
                                                    }
                                                    switch fi11 {
                                                    case -2147483648:
                                                        fmt.Println(v4, fi10, v2)
                                                    case 128:
                                                        v31 := GenS5{F0: ("foo bar" + ""), F1: (v3 ^ (v3 * v3))}
                                                        _ = v31
                                                    default:
                                                        var v32 float32 = 0.0
                                                        _ = v32
                                                    }
                                                    switch (^(fi10)) {
                                                    case -278:
                                                        if false {
                                                            if (!(v2 == v2)) {
                                                                v4 = v4
                                                                fmt.Println(fi10, fi11)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v4 += 123
                                                    v16 = "\n\t"
                                                    v1 += 132
                                                }
                                                switch v16 {
                                                case "a":
                                                    var v33 int = fi11
                                                    _ = len(v5)
                                                    _ = v33
                                                case "hello":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    v34 := GenS1{F0: 653, F1: (v1 ^ v1), F2: 9223372036854775807}
                                                    _ = v34
                                                }
                                                if true {
                                                    v16 = (v16 + "\n\t")
                                                    if (true && (-9223372036854775808 <= -9223372036854775808)) {
                                                        switch fi10 {
                                                        case -71:
                                                            fmt.Println(v4)
                                                        case 42:
                                                            v16 = v16
                                                        default:
                                                            var v35 int32 = v3
                                                            _ = v35
                                                        }
                                                    } else {
                                                        v36 := GenS5{F0: "hello", F1: -987}
                                                        var v37 int64 = 0
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v38 := GenS6{F0: int(v2)}
                                    var v39 float64 = float64(float32(((0.0 - -2.25) * 1e10)))
                                    _ = v38
                                    _ = v39
                                }
                                v40 := GenS5{F0: "", F1: 0}
                                v16 = ((v16 + (((v16 + v16) + (v16 + v16)) + v16)) + "")
                                v16 = v16
                                var v41 int8 = -73
                                switch v16 {
                                case "hello":
                                    switch v16 {
                                    case "foo bar":
                                        if len(v5) > 0 {
                                            v5[0] = fi11
                                        }
                                    case "":
                                        v42 := GenS2{F0: 100}
                                        _ = v42
                                    case "hello":
                                        var v43 uint8 = v4
                                        _ = v43
                                    default:
                                        v5 = append(v5, -1)
                                    }
                                    var v44 int16 = -568
                                    fmt.Println(v3, v44, v3)
                                    _ = v44
                                case "\n\t":
                                    var v45 float64 = 0.1
                                    _ = v45
                                case "a":
                                    fmt.Println(v4, v16)
                                default:
                                    var v46 float64 = 0.1
                                    _ = v46
                                }
                                _ = v18
                                _ = v40
                                _ = v41
                            }
                            switch "" {
                            case "a":
                                var v47 uint = 0
                                _ = v47
                            default:
                                v48 := GenS7{F0: 181}
                                _ = v48
                            }
                            v3 += (^(v3))
                        } else {
                            v16 = v16
                            if len(v5) > 0 {
                                v5[0] = fi10
                            }
                            if true {
                                if len(v5) > 0 {
                                    v5[0] = fi11
                                }
                                var v49 int = fi11
                                fmt.Println(fi11, fi11, v16)
                                _ = v49
                            }
                        }
                        _ = v16
                    } else {
                        var v50 int32 = v3
                        _ = v50
                    }
                    _ = v15
                }
            }
            _ = fi11
        }
        v51 := GenS1{F0: 18446744073709551615, F1: (v1 - v1), F2: 803}
        _ = fi10
        _ = v51
    }
    var v52 string = "foo bar"
    for fi53 := 0; fi53 < 14; fi53++ {
        v3 += 0
        if (2147483647 <= 4294967295) {
            fmt.Println(v52, v1, v4)
        }
        _ = fi53
    }
    v3 += (v3 + (v3 >> ((148) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
