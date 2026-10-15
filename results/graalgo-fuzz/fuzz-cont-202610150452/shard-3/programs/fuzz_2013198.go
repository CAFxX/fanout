package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint; F2 int32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int16 = 1000
    v3 := []int{-2147483648, 32768, -306, 2147483647}
    var v4 uint = 65535
    v2 += v2
    var v5 int8 = -1
    if (!(false || ((!(float32(0.1) <= 0.1)) || (((2 == 2) && (-2.25 == ((3.14159 - float32(1.5)) * (-2.25 - (float32((-2.25 / 1.0)) * 1.5))))) || (((42 >= 42) && (false || (!((v1 < 378) && true)))) || (138 < 707)))))) {
        v6 := GenS1{F0: -558, F1: 938}
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 32768, F1: 1}
            if false {
                v5 += -1
            } else {
                switch ("" + "a") {
                case "\n\t":
                    v9 := GenS2{F0: 0, F1: 2147483647, F2: ((float64(0.0) + -2.25) < float64(0.0))}
                    _ = v9
                case "":
                    _ = len(v3)
                case "foo bar":
                    v3 = append(v3, int(v2))
                default:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v11 := GenS1{F0: (fi10 + (fi10 ^ 42)), F1: 42}
                        v2 = v2
                        if (!true) {
                            v3 = append(v3, (fi7 % ((-9223372036854775808) | 1)))
                            var v12 int64 = 42
                            _ = v12
                        } else {
                            var v13 uint16 = 86
                            _ = v13
                        }
                        if (((0 <= 255) || (true && (390 <= 322))) || (!(v4 < 4294967295))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (!false) {
                                if len(v3) > 0 {
                                    v3[0] = fi7
                                }
                                var v14 bool = false
                                fmt.Println(fi7, v14, v2)
                                _ = v14
                            }
                        }
                        v1 = (v1 & 65535)
                        _ = fi10
                        _ = v11
                    }
                }
            }
            _ = fi7
            _ = v8
        }
        v3 = append(v3, 975)
        _ = v6
    } else {
        v15 := GenS2{F0: 652, F1: 42, F2: (!false)}
        var v16 uint8 = 81
        if true {
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v5 += (v5 | v5)
        }
        _ = v15
        _ = v16
    }
    if (false || true) {
        if (uint64(v1) < 0) {
            v17 := GenS2{F0: 933, F1: -879, F2: ((0 < 493) && ((1 < -885) || false))}
            if true {
                v18 := GenS3{F0: "foo bar", F1: (v4 - v4), F2: -1}
                fmt.Println(v5, v1, v4)
                _ = v18
            } else {
                switch 32767 {
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 3; fi19++ {
                        _ = len(v3)
                        var v20 uint64 = 42
                        var v21 int32 = 42
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS3{F0: "foo bar", F1: (v4 << ((0) & 63)), F2: 997}
                    if true {
                        if (("foo bar" + "a") >= "\n\t") {
                            fmt.Println(v5)
                            switch (fi22 << ((v4) & 63)) {
                            case -2147483648:
                                v3 = append(v3, 128)
                            case 2147483647:
                                fmt.Println(v2, v2)
                            default:
                                var v24 float64 = float64(3.14159)
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 18; fi25++ {
                                if true {
                                    v3 = append(v3, (fi22 - fi22))
                                    var v26 uint8 = 78
                                    _ = v26
                                } else {
                                    v2 = (-(v2))
                                }
                                _ = fi25
                            }
                            for fi27 := 0; fi27 < 7; fi27++ {
                                switch fi27 {
                                case 312:
                                    var v28 uint64 = 624
                                    _ = v28
                                case 127:
                                    fmt.Println(v4)
                                case -2147483648:
                                    if len(v3) > 0 {
                                        v3[0] = (fi22 / (((fi27 + 2147483647)) | 1))
                                    }
                                }
                                _ = fi27
                            }
                        }
                        if false {
                            _ = len(v3)
                            v1 = v1
                        }
                        v4 = v4
                    }
                    if len(v3) > 0 {
                        v3[0] = (fi22 - fi22)
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = v17
        } else {
            for fi29 := 0; fi29 < 14; fi29++ {
                v5 -= int8(v1)
                _ = fi29
            }
            if len(v3) > 0 {
                v3[0] = 841
            }
            v2 = (v2 ^ (-(v2)))
        }
        fmt.Println(v4, v4, v4)
        v4 = (v4 | v4)
    }
    switch 9223372036854775807 {
    case 353:
        v1 = (v1 - v1)
    default:
        var v30 uint64 = 958
        _ = v30
    }
    switch (("hello" + ("a" + "\n\t")) + ("foo bar" + "foo bar")) {
    case "":
        fmt.Println(v2)
    default:
        v3 = append(v3, -346)
    }
    if (-2.25 != 0.0) {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS3{F0: "foo bar", F1: 0, F2: -951}
            _ = fi31
            _ = v32
        }
        var v33 uint = v4
        _ = v33
    } else {
        var v34 string = "a"
        v5 -= int8(v4)
        fmt.Println(v5, v5)
        _ = v34
    }
    v5 = v5
    fmt.Println(v2, v1)
    v3 = append(v3, 610)
    v1 += 239
    fmt.Println(v2)
    v4 = (v4 / (((^(v4))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
