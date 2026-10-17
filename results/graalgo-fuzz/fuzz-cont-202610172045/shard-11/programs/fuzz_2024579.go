package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 int16 = -570
    var v2 int16 = (v1 >> ((42) & 63))
    v3 := []string{"foo bar", "hello", ("\n\t" + "hello"), "foo bar"}
    v4 := []int{0, 133, 9223372036854775807}
    fmt.Println(v1)
    v2 = 42
    if true {
        v1 -= (v2 >> ((134) & 63))
    }
    switch 1 {
    case 128:
        switch 32767 {
        case 2147483647:
            for fi5 := 0; fi5 < 12; fi5++ {
                for fi6 := 0; fi6 < 14; fi6++ {
                    switch "hello" {
                    case "\n\t":
                        v2 += (-(v2))
                    case "foo bar":
                        v1 = -1
                    case "hello":
                        var v7 int32 = 1
                        var v8 uint = 1
                        _ = v7
                        _ = v8
                    default:
                        fmt.Println(fi5, v2)
                    }
                    fmt.Println(v2, fi5, v1)
                    v9 := GenS1{F0: ("" + "\n\t"), F1: 42, F2: 566}
                    switch (("" + "hello") + "foo bar") {
                    case "a":
                        switch 240 {
                        case 32768:
                            v10 := GenS1{F0: ("\n\t" + "hello"), F1: (fi5 & (fi5 / ((fi6) | 1))), F2: 255}
                            _ = v10
                        case -138:
                            if true {
                                if false {
                                    var v11 float64 = (float64(0.0) + 0.1)
                                    _ = v11
                                }
                                if len(v4) > 0 {
                                    v4[0] = (-(fi5))
                                }
                            } else {
                                v12 := GenS2{F0: "foo bar"}
                                var v13 string = ""
                                _ = v12
                                _ = v13
                            }
                        }
                    case "":
                        if len(v4) > 0 {
                            v4[0] = fi6
                        }
                        var v14 int8 = int8(fi5)
                        if false {
                            v15 := GenS1{F0: "\n\t", F1: 0, F2: 18446744073709551615}
                            if ((float64((float32(3.14159) * (((((3.14159 / 1e10) - (float32(0.0) + 0.0)) + (0.1 - 0.1)) + -2.25) + (1e10 / 0.1)))) * 1e10) >= float64(-2.25)) {
                                v1 -= (v2 * v2)
                            } else {
                                fmt.Println(v2)
                            }
                            _ = v15
                        } else {
                            var v16 float64 = (float64((float32(1e10) + 1.5)) - 3.14159)
                            _ = v16
                        }
                        _ = v14
                    case "hello":
                        v17 := GenS2{F0: "a"}
                        _ = v17
                    default:
                        v2 -= v1
                    }
                    v1 = -789
                    _ = fi6
                    _ = v9
                }
                _ = fi5
            }
        case 32768:
            if (327 <= 42) {
                v1 = 32767
                switch 128 {
                case 4294967295:
                    v18 := GenS1{F0: "hello", F1: -240, F2: 537}
                    _ = v18
                case -2147483648:
                    v2 -= 1000
                }
            } else {
                var v19 int8 = 0
                fmt.Println(v2)
                switch 127 {
                case -1:
                    v20 := GenS3{F0: (v2 & int16(v19))}
                    _ = v20
                case 2147483647:
                    v1 += v1
                default:
                    var v21 float64 = 3.14159
                    _ = v21
                }
                _ = v19
            }
        default:
            switch 127 {
            case 951:
                switch 4294967295 {
                case -347:
                    var v22 uint64 = 390
                    _ = v22
                case 4294967295:
                    v23 := GenS4{F0: -7, F1: 128, F2: 100}
                    _ = v23
                case 127:
                    if false {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            var v25 float64 = 1.5
                            var v26 string = ("a" + "\n\t")
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v27 float32 = (3.14159 - 3.14159)
                        _ = v27
                    } else {
                        v28 := GenS1{F0: "", F1: 0, F2: 2}
                        v2 += 846
                        v1 -= (v2 >> ((648) & 63))
                        _ = v28
                    }
                default:
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v2 += 187
                        _ = fi29
                    }
                }
            case -1:
                v3 = append(v3, "\n\t")
            case 32767:
                var v30 int8 = 61
                _ = v30
            }
        }
    case -289:
        var v31 int = -1
        _ = v31
    case -666:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v2 = (v1 * -32768)
    fmt.Println(v2, v2)
    fmt.Println(v1)
    v32 := GenS5{F0: 4294967295, F1: "hello"}
    var v33 float64 = 1e10
    v4 = append(v4, 4294967295)
    var v34 int64 = 0
    if ((false && true) || (true && true)) {
        var v35 string = "foo bar"
        if false {
            if false {
                fmt.Println(v35, v1)
            }
            v36 := GenS4{F0: 42, F1: 255, F2: -1}
            if len(v4) > 0 {
                v4[0] = -180
            }
            _ = v36
        } else {
            switch v35 {
            case "foo bar":
                var v37 uint8 = 200
                _ = v37
            default:
                if true {
                    fmt.Println(v35)
                }
            }
            v38 := GenS6{F0: 4294967295}
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = v38
        }
        if false {
            if len(v4) > 0 {
                v4[0] = 32768
            }
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
