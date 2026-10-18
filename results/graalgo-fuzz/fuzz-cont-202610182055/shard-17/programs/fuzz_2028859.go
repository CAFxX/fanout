package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 int32, p1 int64, p2 uint8) int64 {
    var v6 uint8 = (p2 << ((65535) & 63))
    v7 := GenS1{F0: 18446744073709551615, F1: "hello", F2: p1}
    v8 := []int{918, 551, 127}
    fmt.Println(p0, p1, p1)
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool, p1 int, p2 uint16) uint64 {
    if p0 {
        p2 -= 1
    } else {
        var v9 uint32 = 1
        _ = v9
    }
    switch p1 {
    case 42:
        v10 := GenS1{F0: 53, F1: (((("hello" + ("a" + "hello")) + "foo bar") + "\n\t") + "\n\t"), F2: -129}
        _ = v10
    case 862:
        switch (p1 / (((p1 | p1)) | 1)) {
        case 2147483647:
            if (true || (-92 == 127)) {
                var v11 uint32 = 218
                _ = v11
            }
        case -794:
            for fi12 := 0; fi12 < 19; fi12++ {
                if ((("a" + "hello") >= ("hello" + "foo bar")) && (-98 < 282)) {
                    if (!(!(!p0))) {
                        var v13 int64 = genFn0(2147483647, 9223372036854775807, 127)
                        switch "hello" {
                        case "":
                            var v14 uint16 = p2
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 11; fi15++ {
                                v16 := []int{}
                                if p0 {
                                    var v17 int8 = 42
                                    var v18 float32 = (0.1 + 0.1)
                                    switch (p1 ^ 128) {
                                    case -9223372036854775808:
                                        if len(v16) > 0 {
                                            v16[0] = (^(fi12))
                                        }
                                    default:
                                        fmt.Println(p2)
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi15
                                _ = v16
                            }
                        }
                        fmt.Println(p2)
                        _ = v13
                    } else {
                        if (221 >= 128) {
                            var v19 int = fi12
                            for fi20 := 0; fi20 < 7; fi20++ {
                                v21 := []int{}
                                p2 -= (p2 << ((65535) & 63))
                                p0 = (!false)
                                switch -36 {
                                case 562:
                                    if p0 {
                                        v22 := GenS2{F0: p2, F1: 2, F2: 10}
                                        var v23 uint32 = 42
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                switch (p1 * -2147483648) {
                                case 32767:
                                    p0 = true
                                case 32768:
                                    fmt.Println(fi12)
                                case 9223372036854775807:
                                    v24 := GenS3{F0: 42, F1: p0, F2: -32768}
                                    _ = v24
                                default:
                                    var v25 float32 = (1e10 * float32(float64(0.1)))
                                    _ = v25
                                }
                                _ = fi20
                                _ = v21
                            }
                            for fi26 := 0; fi26 < 14; fi26++ {
                                v27 := []int{42, fi12}
                                p1 = (p1 >> ((18446744073709551615) & 63))
                                _ = fi26
                                _ = v27
                            }
                            _ = v19
                        }
                    }
                } else {
                    v28 := []int{}
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := GenS4{F0: (!(!p0)), F1: genFn0(1, 0, 0)}
                        v31 := GenS5{F0: 801, F1: 2147483648}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
                v32 := GenS3{F0: 4294967295, F1: ((p2 >> ((929) & 63)) < p2), F2: 32767}
                _ = fi12
                _ = v32
            }
        }
    case 2147483647:
        fmt.Println(p1)
    }
    fmt.Println(p2)
    if false {
        p0 = p0
        switch "\n\t" {
        case "":
            p1 = 4294967295
        case "foo bar":
            p0 = p0
        }
        var v33 int32 = 0
        _ = v33
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]string)
    var v3 int16 = 1
    var v4 int8 = 1
    var v5 int16 = v3
    for fi34 := 0; fi34 < 20; fi34++ {
        switch "foo bar" {
        case "foo bar":
            v1 = 1
        case "hello":
            var v35 uint16 = 2
            _ = v35
        case "":
            switch (fi34 | fi34) {
            case -745:
                if false {
                    switch fi34 {
                    case -256:
                        for fi36 := 0; fi36 < 8; fi36++ {
                            switch (fi34 ^ (fi36 + (fi36 - 0))) {
                            case 1:
                                if (!(97 < 144)) {
                                    switch 4294967295 {
                                    case -9223372036854775808:
                                        var v37 bool = false
                                        _ = v37
                                    case 128:
                                        v38 := GenS3{F0: (v1 ^ v1), F1: (v4 > (v4 % (((v4 | (v4 % (((v4 % ((100) | 1))) | 1)))) | 1))), F2: v5}
                                        _ = v38
                                    default:
                                        v5 = (-(v3))
                                    }
                                    fmt.Println(fi34)
                                }
                            default:
                                fmt.Println(fi36, fi34, fi34)
                            }
                            delete(v2, ("\n\t" + (("hello" + ("foo bar" + "foo bar")) + "hello")))
                            _ = fi36
                        }
                    }
                }
            case 9223372036854775807:
                for fi39 := 0; fi39 < 1; fi39++ {
                    v40 := GenS2{F0: 201, F1: 4294967295, F2: 1}
                    fmt.Println(v4, fi34, fi34)
                    v1 = v1
                    _ = fi39
                    _ = v40
                }
            }
        default:
            var v41 uint16 = 255
            _ = v41
        }
        var v42 bool = (714 >= 2)
        v4 = v4
        _ = fi34
        _ = v42
    }
    var v43 int8 = (v4 / ((v4) | 1))
    v44 := []int{1}
    for fi45 := 0; fi45 < 11; fi45++ {
        if ((v4 % ((-128) | 1)) <= -1) {
            fmt.Println(v3)
            for fi46 := 0; fi46 < 4; fi46++ {
                v43 = (v4 / ((v4) | 1))
                var v47 uint64 = 911
                switch "a" {
                case "":
                    v48 := GenS6{F0: (-(v43)), F1: 1}
                    _ = v48
                case "a":
                    var v49 uint8 = 1
                    _ = v49
                }
                _ = fi46
                _ = v47
            }
            if (false && ((!false) && (18446744073709551615 > v1))) {
                for fi50 := 0; fi50 < 17; fi50++ {
                    if ((fi45 % ((-51) | 1)) <= fi50) {
                        v51 := GenS4{F0: (false || (!false)), F1: 9223372036854775807}
                        v52 := GenS4{F0: true, F1: -516}
                        _ = len(v2)
                        _ = v51
                        _ = v52
                    }
                    var v53 int8 = (v43 | v43)
                    _ = fi50
                    _ = v53
                }
                v3 += (v5 + v3)
                v44 = append(v44, (fi45 + fi45))
            }
        } else {
            var v54 float64 = (3.14159 / 1.0)
            var v55 float64 = (v54 + 3.14159)
            _ = v54
            _ = v55
        }
        v4 -= (-(v43))
        v5 = (v3 | v5)
        _ = fi45
    }
    fmt.Println(v4)
    v43 += (v43 << ((18446744073709551615) & 63))
    v3 = v5
    fmt.Println(v5, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
