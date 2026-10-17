package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int64, p1 bool) int8 {
    p1 = (!false)
    switch 0 {
    case -2147483648:
        v5 := GenS1{F0: -128}
        _ = v5
    case 2147483647:
        var v6 uint16 = 767
        _ = v6
    case 128:
        if true {
            p1 = (!(128 != 0))
            if (1 >= 32767) {
                if (320 > 1) {
                    p0 += 138
                    for fi7 := 0; fi7 < 4; fi7++ {
                        var v8 bool = (p0 != 42)
                        _ = fi7
                        _ = v8
                    }
                }
                v9 := []int{1, 9223372036854775807, -2147483648, 128}
                _ = v9
            } else {
                p1 = (-128 <= 100)
                v10 := []int{}
                if (true || ((-47 >= -747) || ((2 > 569) && (p1 && (1 < 0))))) {
                    switch "hello" {
                    case "a":
                        var v11 int = 2147483647
                        _ = v11
                    case "":
                        var v12 string = "hello"
                        _ = v12
                    }
                } else {
                    var v13 float64 = 0.1
                    _ = v13
                }
                _ = v10
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                if p1 {
                    var v15 int64 = (-(p0))
                    var v16 int32 = -904
                    v15 -= (p0 + p0)
                    _ = v15
                    _ = v16
                } else {
                    p0 += 42
                }
                _ = fi14
            }
        }
    default:
        var v17 float32 = 1.5
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            switch 2147483647 {
            case -1:
                if (fi18 != (fi18 << ((42) & 63))) {
                    if p1 {
                        v19 := GenS1{F0: 0}
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 uint8 = 37
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            p1 = (2147483647 < 323)
                        default:
                            v22 := []int{fi20, -9223372036854775808, (fi20 >> ((18446744073709551615) & 63)), fi18}
                            _ = v22
                        }
                        var v23 float64 = 3.14159
                        v24 := []int{fi20}
                        _ = fi20
                        _ = v21
                        _ = v23
                        _ = v24
                    }
                    var v25 float64 = 0.0
                    _ = v25
                }
            }
            p1 = p1
            for fi26 := 0; fi26 < 12; fi26++ {
                if (p1 || ((238 <= 425) && p1)) {
                    v27 := []int{(-(fi18))}
                    _ = v27
                }
                _ = fi26
            }
            _ = fi18
        }
        p1 = p1
        p1 = p1
    }
    if p1 {
        p0 += -263
    }
    p1 = p1
    _ = p0
    _ = p1
    return 29
}
func genFn1(p0 int8, p1 uint32, p2 int8) string {
    var v28 int16 = 1
    v29 := []int{332, 9223372036854775807}
    v30 := GenS2{F0: -688, F1: -960}
    v28 = (v28 >> ((4294967295) & 63))
    var v31 int32 = int32(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    _ = v30
    _ = v31
    return "foo bar"
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = uint16(v1)
    v3 := make(map[uint32]int)
    var v4 int16 = 781
    for fi32 := 0; fi32 < 8; fi32++ {
        _ = len(v3)
        v4 += (v4 + v4)
        v2 -= v2
        _ = fi32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        v2 += v2
        v2 -= v2
        v4 -= (-(v4))
        _ = fi33
    }
    v4 += v4
    var v34 int32 = 42
    switch -757 {
    case 32767:
        if false {
            fmt.Println(v1)
            _ = len(v3)
            for fi35 := 0; fi35 < 14; fi35++ {
                v36 := []int{(fi35 ^ (^(fi35)))}
                fmt.Println(fi35, v1, v4)
                switch (fi35 ^ fi35) {
                case -1:
                    if ((-36 >= 42) && false) {
                        var v37 bool = (!false)
                        _ = v37
                    }
                case 32768:
                    var v38 string = "\n\t"
                    _ = v38
                case -2147483648:
                    var v39 bool = false
                    _ = v39
                default:
                    fmt.Println(fi35, v2)
                }
                _ = fi35
                _ = v36
            }
        }
    case 9223372036854775807:
        v40 := GenS3{F0: 42, F1: 466}
        _ = v40
    case -9223372036854775808:
        v41 := GenS3{F0: 4294967295, F1: v2}
        _ = v41
    default:
        v4 += v4
    }
    var v42 int = -873
    delete(v3, v1)
    fmt.Println(v4, v34)
    switch "\n\t" {
    case "":
        var v43 bool = (-1 <= 0)
        _ = v43
    case "\n\t":
        v34 -= v34
    }
    if ((("hello" + "hello") + ("\n\t" + "hello")) == ("\n\t" + (("a" + "") + "hello"))) {
        v42 -= (v42 / ((v42) | 1))
        switch (("hello" + "\n\t") + "\n\t") {
        case "a":
            v1 = 2
        case "":
            fmt.Println(v4)
        default:
            if (false || (v1 < 42)) {
                v42 = (v42 * 0)
                v42 -= 128
                var v44 uint = 1
                _ = v44
            }
        }
    }
    if true {
        v42 -= v42
    } else {
        for fi45 := 0; fi45 < 5; fi45++ {
            if true {
                v4 += (^(v4))
                v34 = 0
                for fi46 := 0; fi46 < 4; fi46++ {
                    var v47 int64 = 42
                    _ = fi46
                    _ = v47
                }
            }
            v48 := GenS1{F0: -8}
            _ = fi45
            _ = v48
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            fmt.Println(v2, v34)
            _ = fi49
        }
    }
    v50 := GenS1{F0: 100}
    v34 -= v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v50
}
