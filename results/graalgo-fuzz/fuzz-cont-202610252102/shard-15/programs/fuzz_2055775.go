package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 string, p1 uint, p2 uint8) uint64 {
    fmt.Println(p2, p2)
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return 408
}
func genFn1(p0 int16, p1 int16) int64 {
    if true {
        var v6 int16 = p1
        if true {
            var v7 float64 = float64(1e10)
            var v8 uint = 255
            var v9 uint64 = 564
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: (0.0 <= 1e10), F1: 128}
            p1 += p1
            p0 += -32768
            _ = v10
        }
        v11 := GenS1{F0: true, F1: 1}
        _ = v11
        _ = v6
    } else {
        v12 := GenS1{F0: (0.1 != ((1e10 + 1.5) * 1e10)), F1: 42}
        _ = v12
    }
    if false {
        if false {
            if (-9223372036854775808 != -9223372036854775808) {
                switch "" {
                case "hello":
                    if ((("" + "a") + "foo bar") > "") {
                        p1 += p1
                        if ((false || false) && (false || (861 > 42))) {
                            fmt.Println(p0)
                        } else {
                            var v13 int8 = -1
                            _ = v13
                        }
                        p0 += 1000
                    } else {
                        if (1 <= 2147483647) {
                            for fi14 := 0; fi14 < 19; fi14++ {
                                p0 = p0
                                if true {
                                    fmt.Println(p1, p0)
                                    p0 = (p0 * 155)
                                    fmt.Println(fi14, p1, fi14)
                                }
                                _ = fi14
                            }
                            var v15 int64 = 0
                            var v16 uint64 = 1
                            _ = v15
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 18; fi17++ {
                                fmt.Println(fi17, p1, p1)
                                p1 -= 1000
                                _ = fi17
                            }
                        }
                        v18 := GenS2{F0: 193, F1: (!(true || (0 <= 0)))}
                        var v19 int16 = p0
                        _ = v18
                        _ = v19
                    }
                case "":
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int32 = 0
                        if true {
                            fmt.Println(fi20, p1, v21)
                        }
                        p1 += p1
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    switch "hello" {
                    case "foo bar":
                        var v22 uint16 = 42
                        _ = v22
                    case "\n\t":
                        p1 = (p1 - (p0 % ((p1) | 1)))
                    case "":
                        fmt.Println(p1, p0)
                    }
                default:
                    switch 4294967295 {
                    case -1:
                        v23 := []int{128}
                        _ = v23
                    case 32768:
                        fmt.Println(p1, p1)
                    }
                }
                fmt.Println(p0, p1)
                fmt.Println(p1, p0)
            }
        }
        v24 := []int{4294967295, 1, 32767}
        if (!(true || (995 > 2147483647))) {
            fmt.Println(p0, p0)
            v25 := GenS1{F0: false, F1: 0}
            _ = v25
        }
        _ = v24
    } else {
        var v26 int64 = -840
        var v27 uint16 = 32767
        _ = v26
        _ = v27
    }
    var v28 uint16 = 255
    _ = p0
    _ = p1
    _ = v28
    return -864
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 32767
    var v3 uint8 = 128
    var v4 int16 = (-(v1))
    var v5 uint8 = (v3 & (v3 ^ 153))
    var v29 float64 = ((((0.1 - -2.25) - 1e10) - (float64(-2.25) / float64(float32((-2.25 * 0.1))))) - 1.5)
    v2 += v2
    v30 := GenS1{F0: ((v3 >> ((0) & 63)) >= (v5 / ((5) | 1))), F1: -9223372036854775808}
    var v31 uint = 18446744073709551615
    v29 -= ((v29 - 1.5) - v29)
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 int32 = -2147483648
        _ = fi32
        _ = v33
    }
    switch "\n\t" {
    case "\n\t":
        switch "" {
        case "a":
            v3 = (v3 >> ((v31) & 63))
        }
    case "":
        v34 := []int{42}
        _ = v34
    default:
        var v35 int32 = 589
        _ = v35
    }
    if ("" != ("" + ("hello" + "\n\t"))) {
        switch 4294967295 {
        case 128:
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 string = "hello"
                var v38 string = ""
                for fi39 := 0; fi39 < 9; fi39++ {
                    if ((v29 / 1.5) != (v29 * v29)) {
                        v29 += v29
                        var v40 int32 = 300
                        _ = v40
                    } else {
                        var v41 int32 = -2147483648
                        _ = v41
                    }
                    _ = fi39
                }
                _ = fi36
                _ = v37
                _ = v38
            }
        }
    } else {
        var v42 string = "\n\t"
        v3 += 1
        _ = v42
    }
    v4 = -733
    v5 = v5
    v5 = v5
    v31 = (v31 + (v31 * v31))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
