package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
func genFn0() int8 {
    var v7 bool = (!true)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint32 = 925
        var v10 bool = v7
        _ = fi8
        _ = v10
        _ = v9
    }
    if (18446744073709551615 < 255) {
        v11 := []int{}
        if ((!v7) || false) {
            switch 2147483647 {
            case 4294967295:
                var v12 float64 = 0.1
                _ = v12
            }
            fmt.Println(v7)
            _ = len(v11)
        } else {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        _ = v11
    } else {
        var v13 uint = 760
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint64 = 4294967295
                v16 := GenS1{F0: 0, F1: v15}
                _ = fi14
                _ = v15
                _ = v16
            }
        case -624:
            v17 := []int{32767, 1, 9223372036854775807}
            _ = v17
        case -2147483648:
            var v18 float64 = 1.5
            _ = v18
        default:
            if (v7 || v7) {
                var v19 int8 = -54
                var v20 uint16 = 42
                _ = v19
                _ = v20
            } else {
                var v21 int8 = -128
                for fi22 := 0; fi22 < 8; fi22++ {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if ((!(917 == 2)) || false) {
                            var v24 int16 = 954
                            var v25 uint16 = 1
                            _ = v24
                            _ = v25
                        } else {
                            v7 = true
                            fmt.Println(fi23)
                        }
                        _ = fi23
                    }
                    var v26 uint = (v13 << (((v13 - v13)) & 63))
                    _ = fi22
                    _ = v26
                }
                switch ("hello" + "hello") {
                case "\n\t":
                    v27 := GenS1{F0: v21, F1: 18446744073709551615}
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v29 := []int{128, fi28, (fi28 / ((fi28) | 1)), int(v13)}
                        _ = fi28
                        _ = v29
                    }
                default:
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v31 := []int{fi30, 128, 2147483647}
                        _ = fi30
                        _ = v31
                    }
                }
                _ = v21
            }
        }
        _ = v13
    }
    switch ("\n\t" + ("" + ("foo bar" + (("foo bar" + "foo bar") + "")))) {
    case "a":
        v7 = v7
    case "foo bar":
        v7 = false
    }
    _ = v7
    return -128
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{4294967295, 73}
    var v3 int8 = -1
    var v4 int32 = 286
    var v5 float32 = ((-2.25 - 3.14159) - ((float32(1e10) + 1.5) / (1e10 + (0.1 * 1e10))))
    var v6 int64 = -9223372036854775808
    fmt.Println(v1, v4, v4)
    var v32 float64 = 0.1
    for fi33 := 0; fi33 < 7; fi33++ {
        if false {
            v4 -= 42
            v34 := GenS2{F0: (v1 | genFn0()), F1: v6}
            _ = v34
        }
        var v35 int16 = 1000
        _ = fi33
        _ = v35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        if true {
            if len(v2) > 0 {
                v2[0] = 0
            }
            v37 := GenS2{F0: genFn0(), F1: (v6 >> ((42) & 63))}
            _ = v37
        }
        _ = fi36
    }
    var v38 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v32)
    v39 := GenS2{F0: 1, F1: v6}
    v40 := GenS3{F0: v1, F1: 32767, F2: (-(v4))}
    var v41 bool = true
    v5 += 1e10
    switch 1 {
    case 49:
        if len(v2) > 0 {
            v2[0] = 93
        }
    case 4294967295:
        _ = len(v2)
    case 1:
        fmt.Println(v4, v5)
    default:
        v38 = (v38 & 798)
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        fmt.Println(v4)
        _ = fi42
    }
    v43 := GenS3{F0: (v1 + v3), F1: -9223372036854775808, F2: (v4 ^ (v4 * (v4 | (v4 - int32(v38)))))}
    v38 = v38
    var v44 int32 = (v4 + -474)
    var v45 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
