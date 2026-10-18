package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int64; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 int8, p2 int16) uint8 {
    var v6 uint16 = 32767
    p1 += 127
    var v7 uint16 = (v6 << ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 174
}
func genFn1() uint32 {
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int64 = 9223372036854775807
            fmt.Println(v10, v10, fi9)
            if ((0.0 != -2.25) && true) {
                v11 := []int{-2147483648, (fi9 >> ((452) & 63))}
                _ = v11
            } else {
                v12 := GenS1{F0: 23, F1: genFn0(0, 70, 0)}
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v13 uint = 18446744073709551615
        var v14 int8 = -128
        _ = fi8
        _ = v13
        _ = v14
    }
    switch ((("\n\t" + "") + ("a" + ("hello" + ((("hello" + "foo bar") + "") + "hello")))) + "foo bar") {
    case "\n\t":
        var v15 uint8 = 128
        _ = v15
    case "hello":
        v16 := GenS1{F0: 128, F1: 2}
        _ = v16
    case "foo bar":
        if (!true) {
            v17 := GenS1{F0: 1, F1: 42}
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(fi18)
                if (-9223372036854775808 != -9223372036854775808) {
                    v19 := GenS2{F0: 0, F1: 0, F2: 365}
                    for fi20 := 0; fi20 < 14; fi20++ {
                        if (!false) {
                            v21 := []int{(fi20 & fi18), (fi18 ^ (fi18 / ((128) | 1))), fi18, fi20}
                            fmt.Println(fi18, fi20)
                            _ = v21
                        } else {
                            var v22 uint32 = 65535
                            if false {
                                fmt.Println(fi20, fi18)
                                v23 := GenS3{F0: 255}
                                v22 = v22
                                _ = v23
                            } else {
                                v24 := GenS3{F0: 2}
                                switch fi20 {
                                case 23:
                                    v22 = v22
                                }
                                if (v22 != (v22 ^ v22)) {
                                    switch ((((((("\n\t" + ("" + (("a" + "a") + "a"))) + "a") + "a") + ("hello" + "\n\t")) + "a") + "foo bar") + (((("a" + "\n\t") + "\n\t") + ("" + ("\n\t" + "foo bar"))) + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t"))) {
                                    case "\n\t":
                                        if (4294967295 >= 1) {
                                            var v25 int16 = 42
                                            _ = v25
                                        }
                                    }
                                    v22 += 35
                                }
                                _ = v24
                            }
                            var v26 uint16 = 1000
                            _ = v22
                            _ = v26
                        }
                        fmt.Println(fi18, fi20)
                        _ = fi20
                    }
                    _ = v19
                } else {
                    v27 := []int{}
                    var v28 int8 = 42
                    _ = v27
                    _ = v28
                }
                _ = fi18
            }
            _ = v17
        } else {
            if false {
                v29 := GenS2{F0: -1, F1: 1, F2: 42}
                _ = v29
            }
            var v30 string = ("" + "hello")
            v30 = v30
            _ = v30
        }
    }
    switch -9223372036854775808 {
    case 4294967295:
        var v31 float32 = 0.1
        _ = v31
    case 127:
        v32 := []int{2147483647, 503}
        _ = v32
    }
    v33 := []int{32768}
    _ = v33
    return 4
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int32 = 42
    var v4 float32 = 0.1
    var v5 float64 = (float64((v4 - v4)) + ((((3.14159 * 0.1) * float64((((1e10 - (1.5 * v4)) * 0.0) + (float32(3.14159) / 1e10)))) + 1e10) - 0.1))
    v34 := GenS3{F0: 251}
    v4 -= 0.0
    var v35 uint32 = 0
    switch ("a" + ("foo bar" + "hello")) {
    case "\n\t":
        var v36 uint32 = 0
        _ = v36
    case "a":
        for fi37 := 0; fi37 < 3; fi37++ {
            delete(v2, "a")
            fmt.Println(v35, fi37, v35)
            v3 += (v3 & v3)
            _ = fi37
        }
    default:
        var v38 int64 = 0
        _ = v38
    }
    v2["foo bar"] = ((("\n\t" + "\n\t") + "") + "hello")
    if (0 <= 127) {
        v4 -= float32(0.0)
        v39 := GenS1{F0: uint8(v35), F1: 39}
        _ = v39
    }
    v4 = v4
    if ((0.1 * v4) < (3.14159 * 1e10)) {
        fmt.Println(v35, v35)
    }
    fmt.Println(v35, v5, v3)
    if (!(92 <= 158)) {
        switch ("foo bar" + "hello") {
        case "a":
            var v40 int8 = 1
            _ = v40
        case "":
            var v41 int32 = 1
            _ = v41
        case "hello":
            v3 += 714
        default:
            if (!(134 > 16)) {
                if false {
                    if (!true) {
                        var v42 uint8 = 128
                        v3 = (v3 - v3)
                        _ = v42
                    }
                }
                switch "hello" {
                case "a":
                    v43 := []int{327}
                    _ = v43
                case "":
                    for fi44 := 0; fi44 < 4; fi44++ {
                        v45 := GenS3{F0: 0}
                        _ = fi44
                        _ = v45
                    }
                case "foo bar":
                    fmt.Println(v5, v4, v5)
                default:
                    v3 = v3
                }
            } else {
                v4 = v4
                v46 := GenS4{F0: -753, F1: -9223372036854775808, F2: v35}
                _ = v46
            }
        }
    } else {
        v5 += v5
    }
    var v47 int32 = v3
    v48 := GenS3{F0: 1}
    v3 = (v47 >> ((2) & 63))
    if (false || (true || (!((v5 <= -2.25) || true)))) {
        v49 := GenS2{F0: -19, F1: 165, F2: 4294967295}
        switch (("a" + (("hello" + "foo bar") + "hello")) + "") {
        case "hello":
            v50 := []int{-906, 32767}
            _ = v50
        case "foo bar":
            for fi51 := 0; fi51 < 5; fi51++ {
                fmt.Println(v47)
                _ = fi51
            }
        case "":
            if true {
                switch 0 {
                case 32768:
                    v3 = (v3 << ((1) & 63))
                case -365:
                    v52 := GenS5{F0: -649}
                    _ = v52
                case 9223372036854775807:
                    if (!false) {
                        var v53 bool = (1 == 42)
                        v5 -= v5
                        _ = v53
                    } else {
                        v47 = v3
                        var v54 int16 = 175
                        _ = v54
                    }
                default:
                    switch "a" {
                    case "":
                        v35 = (v35 * (v35 >> ((2) & 63)))
                    case "foo bar":
                        v47 = (v47 & 520)
                    default:
                        var v55 int16 = 42
                        _ = v55
                    }
                }
            }
        default:
            var v56 int16 = -32768
            _ = v56
        }
        _ = v49
    }
    _ = len(v1)
    if (4294967295 == 1) {
        var v57 uint = 65535
        var v58 int = -9223372036854775808
        v59 := GenS6{F0: 1, F1: 42, F2: int32(v57)}
        _ = v57
        _ = v58
        _ = v59
    } else {
        if true {
            if (18446744073709551615 >= 0) {
                fmt.Println(v5)
            } else {
                delete(v2, ("" + "\n\t"))
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
