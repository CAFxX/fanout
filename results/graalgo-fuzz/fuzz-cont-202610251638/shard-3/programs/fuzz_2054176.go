package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int8) int64 {
    var v7 string = "foo bar"
    p0 = p0
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = v7
    return -1
}
func genFn1() int32 {
    if (1 != 100) {
        var v8 float64 = float64(-2.25)
        _ = v8
    }
    var v9 int8 = 0
    var v10 uint8 = 207
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 2147483647
    var v3 uint64 = 253
    v4 := make(map[string]int)
    var v5 int8 = 42
    var v6 float32 = (0.1 - (-2.25 * (((0.1 / (1.5 * (float32(-2.25) / 1.0))) + 1e10) * float32(0.1))))
    var v11 uint32 = uint32(v1)
    v12 := GenS1{F0: "", F1: v11, F2: 644}
    switch 0 {
    case 128:
        if false {
            fmt.Println(v6, v11)
            v1 -= (v1 >> ((2) & 63))
            v3 = v3
        } else {
            var v13 float32 = (v6 / (v6 / 1.0))
            _ = v13
        }
    default:
        v14 := []int{127, 128}
        _ = v14
    }
    switch "foo bar" {
    case "\n\t":
        for fi15 := 0; fi15 < 17; fi15++ {
            v3 += v3
            _ = fi15
        }
    case "foo bar":
        for fi16 := 0; fi16 < 20; fi16++ {
            if true {
                v17 := GenS2{F0: "hello", F1: 255}
                if ((42 >= 128) || (0 != 2)) {
                    if true {
                        v18 := v4[(("a" + "foo bar") + "a")]
                        var v19 int64 = -9223372036854775808
                        _ = v18
                        _ = v19
                    } else {
                        if (!((true || (107 >= (v3 | v3))) && false)) {
                            fmt.Println(v1, v1)
                            var v20 uint8 = 127
                            var v21 uint32 = v11
                            _ = v20
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 3; fi22++ {
                                switch fi16 {
                                case -265:
                                    switch 498 {
                                    case -187:
                                        v6 = v6
                                    case -868:
                                        switch 127 {
                                        case -142:
                                            if true {
                                                v5 += v5
                                                v6 += 0.0
                                            }
                                            var v23 int32 = v2
                                            v24 := GenS1{F0: (("\n\t" + ("foo bar" + "")) + ""), F1: (v11 * v11), F2: 9223372036854775807}
                                            if true {
                                                var v25 bool = true
                                                if (v1 < v1) {
                                                    v5 = v5
                                                    v26 := GenS3{F0: (^(v1)), F1: (fi16 * (fi16 >> ((2) & 63))), F2: 1}
                                                    var v27 uint64 = v3
                                                    v28 := GenS4{F0: (v11 * (v11 - 283))}
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    v11 += (v11 + v11)
                                                    var v29 uint = uint(v3)
                                                    _ = v29
                                                }
                                                v1 += (v1 | v1)
                                                _ = v25
                                            }
                                            _ = v23
                                            _ = v24
                                        case 2147483647:
                                            v5 = v5
                                        default:
                                            var v30 int16 = 1
                                            _ = v30
                                        }
                                    }
                                case 9223372036854775807:
                                    var v31 uint64 = v3
                                    _ = v31
                                }
                                v32 := GenS5{F0: 4294967295, F1: (v11 / ((v11) | 1)), F2: -1}
                                _ = fi22
                                _ = v32
                            }
                        }
                    }
                    v2 = v2
                } else {
                    fmt.Println(v5, v11)
                }
                _ = v17
            }
            v2 = v2
            var v33 float64 = 0.0
            _ = fi16
            _ = v33
        }
    }
    v6 -= ((v6 + 1.5) - v6)
    var v34 int16 = 32767
    var v35 string = ("hello" + "")
    fmt.Println(v2, v5)
    for fi36 := 0; fi36 < 12; fi36++ {
        fmt.Println(v34, v35)
        _ = fi36
    }
    if (((v6 - -2.25) / v6) <= 0.0) {
        for fi37 := 0; fi37 < 13; fi37++ {
            fmt.Println(v11, v2, v34)
            for fi38 := 0; fi38 < 9; fi38++ {
                if (!true) {
                    v39 := GenS3{F0: uint16(v3), F1: -557, F2: 976}
                    v34 -= 42
                    _ = v39
                }
                fmt.Println(v11, fi38)
                _ = fi38
            }
            v6 = (float32(-2.25) / v6)
            _ = fi37
        }
        if ((v3 << ((uint(v11)) & 63)) != (v3 / ((v3) | 1))) {
            var v40 int32 = v2
            _ = v40
        } else {
            v41 := GenS6{F0: 42, F1: 585}
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        v6 = v6
        fmt.Println(v2, v35)
        var v43 int64 = 506
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
