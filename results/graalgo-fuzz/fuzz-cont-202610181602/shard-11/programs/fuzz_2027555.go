package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int, p1 int16, p2 uint8) int8 {
    v7 := GenS1{F0: -1, F1: ("\n\t" + ("a" + "foo bar")), F2: "\n\t"}
    var v8 int16 = (p1 ^ (p1 & (^(p1))))
    fmt.Println(v8, v8)
    p2 = (p2 / (((p2 * (^(p2)))) | 1))
    var v9 int = (p0 | (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 58
}

func main() {
    var v1 int32 = 1
    var v2 int32 = 42
    var v3 string = ""
    var v4 uint16 = 42
    var v5 int = 32767
    var v6 int16 = 32767
    v10 := []int{296}
    var v11 uint32 = 869
    switch "hello" {
    case "\n\t":
        v12 := GenS1{F0: -374, F1: ("\n\t" + v3), F2: v3}
        _ = v12
    case "hello":
        var v13 int = v5
        _ = v13
    case "":
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v6 += 32767
        _ = fi15
    }
    v3 = v3
    v4 = uint16(v11)
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(v6, v3)
            _ = fi17
        }
        if false {
            var v18 bool = false
            if v18 {
                v19 := GenS2{F0: (-(v6)), F1: -581}
                _ = v19
            }
            switch ((v3 + v3) + v3) {
            case "a":
                v20 := GenS1{F0: -9223372036854775808, F1: v3, F2: "foo bar"}
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 6; fi21++ {
                    v22 := GenS3{F0: 233}
                    var v23 string = ""
                    if v18 {
                        if (3.14159 > 3.14159) {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            switch int(v11) {
                            case -9223372036854775808:
                                if (402 >= 467) {
                                    v6 += (v6 % ((v6) | 1))
                                    v18 = v18
                                } else {
                                    _ = len(v10)
                                    var v24 string = v3
                                    v25 := GenS3{F0: 42}
                                    v23 = v23
                                    _ = v24
                                    _ = v25
                                }
                            case 303:
                                fmt.Println(v6)
                            case 2147483647:
                                var v26 float64 = (0.0 + -2.25)
                                _ = v26
                            default:
                                switch ("hello" + (v23 + "hello")) {
                                case "a":
                                    v10 = append(v10, (fi21 | (fi21 & (v5 >> ((4294967295) & 63)))))
                                case "":
                                    var v27 int16 = (v6 ^ (^(v6)))
                                    _ = v27
                                default:
                                    v28 := GenS1{F0: -134, F1: ("foo bar" + "a"), F2: v3}
                                    _ = v28
                                }
                                v29 := GenS2{F0: v6, F1: 1}
                                _ = v29
                            }
                            if v18 {
                                if v18 {
                                    switch v3 {
                                    case "foo bar":
                                        fmt.Println(v11)
                                    default:
                                        var v30 float64 = (-2.25 / (1e10 / 1.0))
                                        _ = v30
                                    }
                                } else {
                                    var v31 uint32 = 562
                                    var v32 uint64 = 584
                                    switch fi21 {
                                    case 128:
                                        v33 := GenS2{F0: (v6 % ((v6) | 1)), F1: -741}
                                        _ = v33
                                    }
                                    _ = v31
                                    _ = v32
                                }
                            }
                            v10 = append(v10, fi21)
                            v34 := GenS1{F0: 317, F1: v23, F2: "a"}
                            _ = v34
                        } else {
                            var v35 uint32 = v11
                            v3 = "\n\t"
                            v36 := GenS2{F0: v6, F1: -484}
                            _ = v35
                            _ = v36
                        }
                        var v37 uint8 = 128
                        var v38 int32 = v2
                        _ = v37
                        _ = v38
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                v39 := GenS4{F0: 65535}
                _ = v39
            }
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v5, v11)
    v40 := GenS1{F0: 794, F1: v3, F2: (v3 + v3)}
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 9; fi42++ {
            if true {
                var v43 uint16 = v4
                v44 := GenS2{F0: (v6 ^ 1000), F1: 9223372036854775807}
                _ = v43
                _ = v44
            }
            _ = fi42
        }
        v4 = 2
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
