package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint }
func genFn0() uint8 {
    v7 := GenS1{F0: 1, F1: (("" + "foo bar") + ("foo bar" + "hello"))}
    switch ("\n\t" + "\n\t") {
    case "":
        v8 := GenS2{F0: 1, F1: 0}
        _ = v8
    case "\n\t":
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    var v10 int32 = -2147483648
    _ = v10
    _ = v7
    return 127
}
func genFn1(p0 int8, p1 int, p2 int32) uint {
    var v11 int8 = p0
    switch p1 {
    case 9223372036854775807:
        if ((!false) && true) {
            v12 := []int{-2147483648, 32767, (-(p1))}
            fmt.Println(v11, p1)
            v11 = 100
            _ = v12
        }
    case 469:
        fmt.Println(p0, p1, p2)
    case 32767:
        p2 -= 0
    default:
        if ((true || true) && ((1 > 0) || true)) {
            var v13 uint64 = 0
            for fi14 := 0; fi14 < 10; fi14++ {
                p1 = (p1 >> ((2) & 63))
                v15 := GenS3{F0: 1, F1: 0}
                v16 := GenS4{F0: 0, F1: 0, F2: 48}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        } else {
            if (-111 == p0) {
                p2 -= 0
            }
        }
    }
    var v17 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v17
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 1
    v3 := make(map[int]string)
    v4 := []int{-1, 940}
    var v5 float64 = (1.5 + -2.25)
    v6 := []uint32{42, 713}
    fmt.Println(v5)
    if true {
        for fi18 := 0; fi18 < 6; fi18++ {
            v2 -= genFn1(23, (^(fi18)), 0)
            v19 := GenS2{F0: 0, F1: 807}
            _ = fi18
            _ = v19
        }
        v2 -= 2
        _ = len(v3)
    }
    v5 += (float64(0.1) * (float64(3.14159) / 1.0))
    if (4294967295 != 128) {
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := v3[-1]
                _ = fi21
                _ = v22
            }
            v5 -= -2.25
            v5 = 0.0
            _ = fi20
        }
    } else {
        var v23 uint64 = 904
        if (float64((0.0 + ((-2.25 / 0.1) * 1.5))) == (v5 - v5)) {
            var v24 string = ("a" + "foo bar")
            _ = v24
        } else {
            v23 = v23
            if false {
                var v25 uint = 1
                v23 += 703
                _ = v25
            } else {
                var v26 int64 = -9223372036854775808
                _ = v26
            }
        }
        _ = v23
    }
    var v27 int = 128
    if (-766 >= -523) {
        v5 += v5
        v2 -= v2
        fmt.Println(v27, v2)
    } else {
        var v28 int32 = -1
        v29 := GenS2{F0: -32768, F1: 9223372036854775807}
        _ = v28
        _ = v29
    }
    if true {
        fmt.Println(v5)
        for fi30 := 0; fi30 < 19; fi30++ {
            v27 -= 32768
            _ = fi30
        }
    } else {
        if true {
            fmt.Println(v2)
            v2 = 804
        }
    }
    v27 = v27
    var v31 int16 = 754
    v32 := GenS4{F0: 740, F1: 1, F2: genFn1(100, (v27 >> ((v2) & 63)), -2147483648)}
    delete(v1, "\n\t")
    var v33 uint32 = 2147483648
    fmt.Println(v27, v27)
    if false {
        v34 := GenS4{F0: 0, F1: 2, F2: 608}
        v31 -= v31
        _ = v34
    } else {
        var v35 float32 = 1.5
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 float64 = (float64(3.14159) + float64(1e10))
            v2 -= (v2 | (^(v2)))
            v33 = (^(v33))
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        if (true && false) {
            v31 += (v31 >> ((genFn1(1, v27, 42)) & 63))
            v27 = -528
            v27 = -78
        }
        _ = fi38
    }
    switch "hello" {
    case "a":
        for fi39 := 0; fi39 < 3; fi39++ {
            v2 -= (v2 * (^(v2)))
            v40 := GenS2{F0: (v31 >> ((genFn1(42, v27, 42)) & 63)), F1: -836}
            fmt.Println(v27, v2, fi39)
            _ = fi39
            _ = v40
        }
    case "":
        v33 = (v33 << ((genFn1(-128, (^(v27)), 2147483647)) & 63))
    case "\n\t":
        v41 := GenS2{F0: 1000, F1: -9223372036854775808}
        _ = v41
    default:
        v27 = v27
    }
    if (-6 <= 0) {
        v6 = append(v6, 966)
    }
    v42 := GenS1{F0: 128, F1: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
