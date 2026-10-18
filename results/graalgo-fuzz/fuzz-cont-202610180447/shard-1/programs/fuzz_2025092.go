package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8; F2 int }
type GenS6 struct { F0 uint64 }
func genFn0() uint32 {
    if true {
        switch -2147483648 {
        case 127:
            switch "foo bar" {
            case "":
                var v5 int = 127
                _ = v5
            case "foo bar":
                var v6 int32 = 1
                _ = v6
            }
        case 966:
            v7 := []int{438, -9223372036854775808, 127, 128}
            _ = v7
        case -976:
            var v8 int64 = -9223372036854775808
            _ = v8
        default:
            var v9 int = 1
            _ = v9
        }
        var v10 uint32 = 2
        _ = v10
    } else {
        var v11 int32 = -164
        _ = v11
    }
    if (-9223372036854775808 > 1) {
        var v12 int64 = 480
        var v13 uint16 = 32767
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: -2147483648, F1: ((!(((true && true) || true) || false)) && true)}
    _ = v14
    return 2147483648
}
func genFn1(p0 uint8, p1 uint32, p2 bool) int8 {
    var v15 string = "foo bar"
    var v16 int64 = -1
    fmt.Println(p1, v15)
    var v17 int8 = -128
    var v18 int64 = v16
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return 71
}

func main() {
    var v1 int8 = 100
    var v2 int32 = 1
    var v3 uint64 = 2
    var v4 uint64 = v3
    v4 = (v4 & (v3 << ((255) & 63)))
    v19 := []int{128, 0, 127}
    if false {
        if true {
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS2{F0: 32767}
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(v4, fi22, fi22)
                fmt.Println(v3, v1, v4)
                v3 = v3
                _ = fi22
            }
            v2 = (-(v2))
        } else {
            var v23 uint32 = 773
            if (18446744073709551615 == 64) {
                v2 = v2
            }
            var v24 float32 = float32(1.5)
            _ = v23
            _ = v24
        }
        var v25 string = ("hello" + ("foo bar" + "a"))
        _ = v25
    } else {
        v4 = v3
        switch "foo bar" {
        case "\n\t":
            _ = len(v19)
        case "foo bar":
            fmt.Println(v1, v2, v1)
        case "":
            v26 := GenS3{F0: 1, F1: "\n\t", F2: v2}
            _ = v26
        default:
            var v27 int16 = 1000
            _ = v27
        }
        v2 -= v2
    }
    v19 = append(v19, -253)
    v3 += v3
    if len(v19) > 0 {
        v19[0] = 32767
    }
    v28 := GenS4{F0: 515}
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS5{F0: 2, F1: v1, F2: fi29}
        _ = fi29
        _ = v30
    }
    var v31 int = -1
    for fi32 := 0; fi32 < 1; fi32++ {
        if true {
            v33 := GenS6{F0: (v4 | (v3 / ((v4) | 1)))}
            v3 = v3
            for fi34 := 0; fi34 < 10; fi34++ {
                v2 += v2
                _ = fi34
            }
            _ = v33
        } else {
            if true {
                var v35 bool = (!(0.1 == 0.1))
                for fi36 := 0; fi36 < 17; fi36++ {
                    fmt.Println(v35, fi36, fi32)
                    switch 32767 {
                    case 2147483647:
                        v37 := GenS3{F0: 0, F1: ("a" + ("" + "foo bar")), F2: (v2 | (v2 - (^(v2))))}
                        _ = v37
                    case 32768:
                        switch "a" {
                        case "a":
                            if (2147483648 <= 42) {
                                v4 = 4294967295
                                var v38 float64 = (0.0 / 0.1)
                                fmt.Println(fi36, fi32)
                                _ = v38
                            } else {
                                if v35 {
                                    v4 = 18446744073709551615
                                } else {
                                    var v39 uint = 638
                                    v40 := GenS5{F0: 2, F1: v1, F2: v31}
                                    _ = v39
                                    _ = v40
                                }
                                v41 := GenS3{F0: 0, F1: "a", F2: v2}
                                _ = v41
                            }
                        case "\n\t":
                            v1 = v1
                        }
                    case 9223372036854775807:
                        if false {
                            v35 = (!((v35 && (v35 || (genFn1(88, 0, v35) < (v1 % (((^(v1))) | 1))))) && (v35 && v35)))
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                    default:
                        v31 -= (^(v31))
                    }
                    var v42 uint64 = (v4 & 0)
                    _ = fi36
                    _ = v42
                }
                for fi43 := 0; fi43 < 3; fi43++ {
                    v3 -= v4
                    _ = fi43
                }
                _ = v35
            }
            fmt.Println(v31)
            v2 = v2
        }
        v2 += v2
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
