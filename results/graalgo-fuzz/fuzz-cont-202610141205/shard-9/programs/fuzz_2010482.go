package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int64) string {
    switch "foo bar" {
    case "a":
        p0 = p0
    case "foo bar":
        fmt.Println(p0)
    case "":
        var v4 int = 42
        _ = v4
    default:
        var v5 uint32 = 65535
        _ = v5
    }
    switch "\n\t" {
    case "":
        v6 := []int{}
        _ = v6
    default:
        p0 = (p0 | p0)
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    if (871 > 255) {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi9, fi9)
            var v10 int16 = -276
            p0 = p0
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    return (("a" + (("hello" + "hello") + (("hello" + ("\n\t" + "\n\t")) + "\n\t"))) + "hello")
}
func genFn1(p0 string, p1 int8) int32 {
    switch "a" {
    case "":
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 bool = (false && (-1 == 1))
            _ = fi11
            _ = v12
        }
    case "\n\t":
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int16 = 95
            v15 := GenS1{F0: 846, F1: 42, F2: (p1 % (((p1 - p1)) | 1))}
            if true {
                v14 += (v14 | v14)
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 int64 = -9223372036854775808
                    var v18 int16 = 1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    case "a":
        var v19 string = (((p0 + "\n\t") + p0) + (p0 + "\n\t"))
        _ = v19
    default:
        p1 = 100
    }
    p0 = "foo bar"
    if ((p1 % (((p1 * (p1 * (p1 | p1)))) | 1)) != -128) {
        for fi20 := 0; fi20 < 15; fi20++ {
            fmt.Println(p1)
            p1 = (p1 % (((p1 - (-(p1)))) | 1))
            _ = fi20
        }
    } else {
        var v21 uint8 = 128
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint8 = 13
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    fmt.Println(p0, p1)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 float64 = 0.1
        fmt.Println(v25, p1, p0)
        p1 = p1
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint16 = 255
    var v2 uint = 2
    v3 := []string{"hello", ("" + (("foo bar" + ("" + ("" + "foo bar"))) + "hello"))}
    var v26 uint64 = 2
    var v27 int8 = 25
    for fi28 := 0; fi28 < 15; fi28++ {
        v26 += 453
        var v29 string = "\n\t"
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v27, v26)
    fmt.Println(v26)
    for fi30 := 0; fi30 < 1; fi30++ {
        if (genFn1(("foo bar" + "hello"), (v27 << ((v2) & 63))) != 0) {
            switch "" {
            case "foo bar":
                fmt.Println(v27, fi30)
            case "a":
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 uint = (^(v2))
                    _ = fi31
                    _ = v32
                }
            default:
                for fi33 := 0; fi33 < 13; fi33++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v27 += v27
                    _ = fi33
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v34 := GenS1{F0: v2, F1: 2, F2: -103}
            _ = v34
        } else {
            fmt.Println(fi30, fi30, v26)
        }
        _ = fi30
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 int = (^(fi35))
        _ = fi35
        _ = v36
    }
    if (((((("foo bar" + (("" + "a") + ("foo bar" + "hello"))) + "foo bar") + ("a" + "foo bar")) >= "") && (2 >= 0)) || (false || (v27 < 109))) {
        fmt.Println(v2, v2, v2)
        fmt.Println(v27, v1)
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        v38 := GenS2{F0: 42, F1: v2}
        fmt.Println(v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi37
        _ = v38
    }
    if false {
        var v39 int16 = 73
        v2 = (^(v2))
        _ = v39
    } else {
        v27 += (v27 ^ v27)
        v1 = v1
    }
    v27 = (v27 + -128)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v40 := GenS3{F0: (true && (42 >= -446))}
    switch -1 {
    case 885:
        var v41 bool = false
        _ = v41
    case 0:
        if true {
            var v42 uint8 = 0
            var v43 int32 = 1
            v2 -= v2
            _ = v42
            _ = v43
        }
    case -645:
        if ((v2 * 4294967295) < (v2 & (v2 % (((v2 % ((v2) | 1))) | 1)))) {
            for fi44 := 0; fi44 < 8; fi44++ {
                fmt.Println(v27, v26)
                var v45 int8 = v27
                for fi46 := 0; fi46 < 5; fi46++ {
                    switch "\n\t" {
                    case "hello":
                        var v47 string = "\n\t"
                        _ = v47
                    }
                    var v48 int = fi44
                    var v49 int16 = -567
                    _ = fi46
                    _ = v48
                    _ = v49
                }
                _ = fi44
                _ = v45
            }
            var v50 uint64 = 18446744073709551615
            fmt.Println(v50)
            _ = v50
        } else {
            var v51 uint = (v2 / ((42) | 1))
            _ = v51
        }
    default:
        for fi52 := 0; fi52 < 18; fi52++ {
            v2 = (v2 >> ((452) & 63))
            var v53 int64 = 1
            if ((("foo bar" == "\n\t") || (!(!(v1 > 0)))) && (0 == 1)) {
                if (false || false) {
                    var v54 float64 = float64(0.0)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v54
                } else {
                    var v55 int8 = (v27 ^ (v27 << (((^(v2))) & 63)))
                    v56 := GenS4{F0: v26}
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = v55
                    _ = v56
                }
            }
            _ = fi52
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v40
}
