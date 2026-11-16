package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 int, p1 int32, p2 int) uint16 {
    var v6 int = 42
    p0 = p2
    v7 := []int{(v6 / ((-2147483648) | 1)), (v6 % ((p0) | 1)), p0, (v6 + 32767)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 float64 = (float64(0.1) * 0.1)
    v2 := []uint32{368}
    var v3 bool = (222 <= 156)
    var v4 int = -2147483648
    v5 := []string{"\n\t", "", ("" + "a")}
    if v3 {
        fmt.Println(v4, v1)
    } else {
        v2 = append(v2, 65)
        v8 := GenS1{F0: 42, F1: 65535, F2: -1}
        v4 = (v4 % (((v4 & (-(v4)))) | 1))
        _ = v8
    }
    switch ("foo bar" + ("hello" + ("hello" + "hello"))) {
    case "a":
        if v3 {
            fmt.Println(v4, v1)
            var v9 uint32 = 1
            _ = v9
        }
    case "foo bar":
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v1)
            v3 = ((1e10 - 1.5) >= (-2.25 - ((-2.25 - 0.1) / 3.14159)))
            _ = fi10
        }
    case "hello":
        var v11 int8 = 1
        _ = v11
    default:
        if (!((v1 * v1) >= 0.1)) {
            var v12 uint16 = 42
            switch ("a" + ("" + ("hello" + ("hello" + "a")))) {
            case "":
                var v13 uint = 818
                _ = v13
            }
            v12 = v12
            _ = v12
        } else {
            v4 = (^(v4))
            for fi14 := 0; fi14 < 15; fi14++ {
                v4 = v4
                fmt.Println(v3, v3, v3)
                var v15 string = (("hello" + "") + ("hello" + "\n\t"))
                _ = fi14
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = -1
        var v18 float64 = 0.0
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint = 18446744073709551615
    v5 = append(v5, ("foo bar" + (("" + "") + "")))
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    if true {
        v1 -= float64(3.14159)
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v1 -= v1
        v3 = v3
        _ = fi20
    }
    if (v3 && (114 == 694)) {
        v3 = (!false)
        v19 += (v19 + (v19 << (((v19 + (v19 % (((^(v19))) | 1)))) & 63)))
    } else {
        fmt.Println(v3)
        v21 := GenS1{F0: -1, F1: 65535, F2: -1}
        v2 = append(v2, 65535)
        _ = v21
    }
    v22 := GenS1{F0: 1000, F1: 1, F2: 9223372036854775807}
    var v23 float64 = (-2.25 * (v1 - v1))
    fmt.Println(v3)
    var v24 int64 = 1
    v25 := GenS1{F0: 42, F1: 455, F2: (v24 | v24)}
    if (!true) {
        if (((false && true) || (v3 || (true && v3))) || (42 <= 1)) {
            v2 = append(v2, 0)
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 string = ("a" + ("hello" + "foo bar"))
                _ = fi26
                _ = v27
            }
            v24 -= 739
        } else {
            var v28 uint16 = 718
            for fi29 := 0; fi29 < 10; fi29++ {
                v30 := GenS1{F0: -416, F1: 2147483648, F2: (v24 ^ 0)}
                if v3 {
                    v4 -= fi29
                    var v31 int32 = -41
                    v23 += (-2.25 * (0.0 - v23))
                    _ = v31
                } else {
                    var v32 int8 = -85
                    if v3 {
                        v33 := GenS2{F0: 160, F1: (v19 << ((4294967295) & 63)), F2: 42}
                        for fi34 := 0; fi34 < 2; fi34++ {
                            switch (fi29 << ((v19) & 63)) {
                            case 127:
                                v2 = append(v2, 1)
                            case -9223372036854775808:
                                v35 := GenS1{F0: -449, F1: 464, F2: v24}
                                _ = v35
                            default:
                                v24 = 9223372036854775807
                            }
                            _ = fi34
                        }
                        for fi36 := 0; fi36 < 10; fi36++ {
                            v3 = true
                            if ((!(0 != 143)) && v3) {
                                v32 = (v32 + v32)
                            }
                            _ = fi36
                        }
                        _ = v33
                    } else {
                        var v37 uint64 = 18446744073709551615
                        _ = v37
                    }
                    var v38 uint16 = 292
                    _ = v32
                    _ = v38
                }
                if ((v24 - (v24 << (((^(v19))) & 63))) == v24) {
                    var v39 uint = v19
                    for fi40 := 0; fi40 < 1; fi40++ {
                        v23 = (v23 / 1.5)
                        var v41 uint16 = v28
                        v42 := GenS1{F0: 0, F1: 0, F2: v24}
                        v3 = (0 != -1)
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v39
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        if true {
            v43 := GenS3{F0: 559}
            var v44 int = (v4 | v4)
            _ = v43
            _ = v44
        }
        var v45 uint = v19
        _ = v45
    }
    if false {
        v46 := GenS4{F0: 336, F1: "\n\t", F2: ((!(("foo bar" + "") <= ("a" + "hello"))) && false)}
        _ = v46
    }
    var v47 int8 = 100
    switch "" {
    case "\n\t":
        for fi48 := 0; fi48 < 12; fi48++ {
            v5 = append(v5, "hello")
            _ = fi48
        }
    default:
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
