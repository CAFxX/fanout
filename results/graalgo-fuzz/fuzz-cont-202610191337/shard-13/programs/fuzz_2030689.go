package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint32; F2 bool }
func genFn0(p0 int32) uint16 {
    switch 9223372036854775807 {
    case 32768:
        var v7 int16 = -419
        _ = v7
    case 1:
        p0 += p0
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            if ((fi8 >> ((2) & 63)) >= fi8) {
                v9 := GenS1{F0: 876, F1: 0, F2: 975}
                if false {
                    var v10 uint64 = 899
                    p0 += -2147483648
                    var v11 uint32 = 120
                    _ = v10
                    _ = v11
                } else {
                    p0 = (p0 % ((p0) | 1))
                    if true {
                        v12 := []int{fi8, (fi8 + 32767)}
                        for fi13 := 0; fi13 < 7; fi13++ {
                            if false {
                                if false {
                                    var v14 int8 = -74
                                    _ = v14
                                } else {
                                    var v15 uint64 = 4294967295
                                    _ = v15
                                }
                            }
                            if len(v12) > 0 {
                                v12[0] = (fi8 & -9223372036854775808)
                            }
                            _ = fi13
                        }
                        _ = v12
                    } else {
                        p0 -= p0
                        var v16 int = (^(fi8))
                        for fi17 := 0; fi17 < 18; fi17++ {
                            var v18 string = "foo bar"
                            v19 := GenS1{F0: 18446744073709551615, F1: -128, F2: 1}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = v16
                    }
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v21 := GenS1{F0: 1, F1: 84, F2: 106}
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v9
            } else {
                v22 := []int{(fi8 << ((18446744073709551615) & 63)), (fi8 ^ (fi8 * fi8))}
                switch ("hello" + "a") {
                case "foo bar":
                    v23 := GenS2{F0: -1, F1: fi8, F2: 42}
                    _ = v23
                case "\n\t":
                    p0 = (p0 + p0)
                case "hello":
                    p0 += (p0 - p0)
                }
                _ = v22
            }
            p0 += p0
            for fi24 := 0; fi24 < 16; fi24++ {
                p0 = 116
                _ = fi24
            }
            _ = fi8
        }
    }
    var v25 bool = true
    var v26 uint16 = 65535
    v25 = false
    _ = p0
    _ = v25
    _ = v26
    return v26
}

func main() {
    var v1 int64 = 42
    var v2 string = "hello"
    var v3 int8 = 33
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    v6 := []int{}
    fmt.Println(v2, v1)
    if (-654 > 1) {
        var v27 string = v2
        _ = v27
    } else {
        v28 := GenS3{F0: 255, F1: 712}
        fmt.Println(v3, v1, v3)
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
        _ = v28
    }
    if ((42 == 2) && (((((-2.25 != (0.1 + (3.14159 + (3.14159 * 3.14159)))) && true) || false) || (4294967295 != 787)) || (65535 >= 2147483647))) {
        v29 := GenS4{F0: (-(v1))}
        var v30 int64 = v1
        v31 := GenS1{F0: 675, F1: v3, F2: 2}
        _ = v29
        _ = v30
        _ = v31
    }
    switch v2 {
    case "hello":
        if (!true) {
            var v32 float32 = (float32(1e10) - 1.5)
            var v33 int = 42
            _ = v32
            _ = v33
        } else {
            fmt.Println(v3, v2)
            var v34 int8 = (v3 << ((1) & 63))
            var v35 uint8 = 1
            _ = v34
            _ = v35
        }
    case "":
        var v36 uint8 = 2
        _ = v36
    case "a":
        if false {
            var v37 uint32 = 0
            if false {
                v6 = append(v6, -1)
                var v38 uint8 = 42
                _ = v38
            } else {
                var v39 int = 9223372036854775807
                for fi40 := 0; fi40 < 12; fi40++ {
                    v2 = (((v2 + "hello") + v2) + ((v2 + ("foo bar" + (v2 + "hello"))) + ("foo bar" + (v2 + v2))))
                    _ = fi40
                }
                _ = v39
            }
            _ = v37
        }
    default:
        v41 := GenS5{F0: 2, F1: 1, F2: (!(false || (!true)))}
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v3 = v3
        if (!(true && true)) {
            v2 = ("hello" + (v2 + "a"))
            v1 += (^(v1))
        }
        if len(v6) > 0 {
            v6[0] = 1
        }
        _ = fi42
    }
    fmt.Println(v3, v3, v3)
    for fi43 := 0; fi43 < 13; fi43++ {
        v2 = (v2 + ((v2 + v2) + v2))
        _ = fi43
    }
    switch v2 {
    case "hello":
        var v44 uint = 42
        _ = v44
    case "\n\t":
        var v45 bool = (409 >= 1)
        _ = v45
    case "":
        var v46 uint = 4294967295
        _ = v46
    }
    var v47 string = v2
    v5[("foo bar" + v47)] = false
    if false {
        v6 = append(v6, 1)
    } else {
        if (!(255 == 195)) {
            var v48 uint8 = 40
            _ = v48
        }
        v1 = -1
        v49 := GenS2{F0: (v1 - v1), F1: 9223372036854775807, F2: 0}
        _ = v49
    }
    var v50 int8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
