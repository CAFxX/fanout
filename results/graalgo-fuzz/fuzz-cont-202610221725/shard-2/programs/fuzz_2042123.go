package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint16 = 841
    v2 := []uint32{711, 42}
    var v3 int16 = -1
    v4 := []uint32{2, 2147483648, 101, 826}
    var v5 uint16 = 735
    var v6 int8 = 17
    v7 := GenS1{F0: (v1 | v5), F1: (v3 + v3)}
    switch "" {
    case "hello":
        v8 := GenS1{F0: (v1 | (v5 + v5)), F1: 1}
        _ = v8
    case "foo bar":
        v5 = (^(v5))
    case "a":
        var v9 float32 = ((-2.25 * 1e10) - 1e10)
        _ = v9
    default:
        var v10 uint64 = 42
        _ = v10
    }
    if false {
        var v11 int8 = (v6 << ((18446744073709551615) & 63))
        fmt.Println(v3, v6, v11)
        var v12 float32 = 0.1
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v1)
    if (42 != -95) {
        if (!(0 > 1)) {
            var v13 bool = false
            v3 = v3
            var v14 uint8 = 0
            _ = v13
            _ = v14
        }
        if (255 <= 780) {
            v15 := GenS2{F0: -1}
            var v16 bool = true
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint32 = 42
                v19 := GenS1{F0: 2, F1: (v3 - v3)}
                var v20 int16 = 42
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v15
            _ = v16
        } else {
            v21 := GenS2{F0: 42}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                if len(v2) > 0 {
                    v2[0] = 770
                }
                var v22 uint16 = 32767
                if ((745 <= 2) || (!(4294967295 < 65535))) {
                    v2 = append(v2, 2147483647)
                } else {
                    fmt.Println(v6)
                }
                _ = v22
            } else {
                var v23 float64 = 1.5
                _ = v23
            }
            _ = v21
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v3, v6)
            if false {
                v6 = 0
            }
            var v25 uint64 = 2
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int8 = (v6 % (((v6 << ((0) & 63))) | 1))
        switch fi26 {
        case -65:
            if false {
                if false {
                    var v28 int = (fi26 << ((1) & 63))
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        fmt.Println(v27, v5, v27)
                        if len(v2) > 0 {
                            v2[0] = 448
                        }
                        v1 -= v5
                        _ = fi29
                    }
                }
                v3 = (-(v3))
            } else {
                _ = len(v4)
            }
        case 2147483647:
            v6 -= v6
        case 32768:
            v6 += -128
        default:
            v5 += v5
        }
        _ = fi26
        _ = v27
    }
    var v30 uint = 1
    for fi31 := 0; fi31 < 1; fi31++ {
        switch "a" {
        case "":
            var v32 int64 = -1
            _ = v32
        case "a":
            v33 := GenS3{F0: fi31, F1: 111}
            _ = v33
        default:
            var v34 bool = true
            _ = v34
        }
        if true {
            if true {
                var v35 uint = (v30 | v30)
                for fi36 := 0; fi36 < 8; fi36++ {
                    v6 += v6
                    var v37 float32 = -2.25
                    _ = fi36
                    _ = v37
                }
                v5 -= v5
                _ = v35
            } else {
                var v38 string = "a"
                switch v38 {
                case "hello":
                    if (0 > -2147483648) {
                        fmt.Println(fi31)
                        v30 = (v30 & 373)
                    }
                default:
                    v6 = v6
                }
                _ = v38
            }
        }
        _ = fi31
    }
    if (v1 >= 907) {
        fmt.Println(v5)
        var v39 float32 = (0.1 - (3.14159 / 1e10))
        v40 := GenS3{F0: 42, F1: 42}
        _ = v39
        _ = v40
    } else {
        _ = len(v4)
        var v41 float64 = 0.0
        _ = v41
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v42 int = 4294967295
    switch ("hello" + "hello") {
    case "\n\t":
        if ((^(v6)) <= 101) {
            var v43 int8 = v6
            v3 += (v3 - (v3 | v3))
            _ = v43
        }
    default:
        var v44 bool = false
        _ = v44
    }
    v45 := GenS2{F0: 1}
    _ = len(v2)
    var v46 float64 = -2.25
    v47 := GenS4{F0: v6, F1: -27}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
