package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64 }
func genFn0() int64 {
    var v4 float64 = ((1e10 / 1.0) * float64(-2.25))
    v5 := []int{42}
    fmt.Println(v4)
    if false {
        var v6 uint16 = 697
        _ = v6
    }
    _ = v4
    _ = v5
    return 714
}
func genFn1(p0 int16, p1 uint64) uint64 {
    fmt.Println(p1)
    var v7 uint8 = 255
    var v8 string = ("hello" + ("foo bar" + ("\n\t" + ("a" + "a"))))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []string{"hello", ("\n\t" + (("" + "") + "")), ("" + "")}
    var v2 uint32 = 786
    v3 := []string{"foo bar", ("\n\t" + ("hello" + (("a" + "") + ""))), "foo bar", ("foo bar" + "hello")}
    v2 = (v2 - 941)
    var v9 float32 = ((1.5 + -2.25) - float32(1.5))
    switch 4294967295 {
    case 32767:
        if ((!false) || ((1.5 / 3.14159) != 0.1)) {
            switch "" {
            case "a":
                v10 := GenS1{F0: 42, F1: genFn0(), F2: (!false)}
                _ = v10
            case "\n\t":
                fmt.Println(v2)
            }
            v9 = v9
            var v11 uint8 = 0
            _ = v11
        }
    case 4294967295:
        if (0 == 1) {
            fmt.Println(v2)
            v9 = v9
            if (v2 < (^(v2))) {
                var v12 int32 = -90
                if (1 > 104) {
                    var v13 int16 = -32768
                    _ = v13
                }
                v14 := GenS2{F0: ((^(v2)) >= (v2 - (v2 << ((980) & 63))))}
                _ = v12
                _ = v14
            } else {
                v9 = (float32(((1.5 / 1.0) / 0.1)) / v9)
                fmt.Println(v2, v2)
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint64 = genFn1(int16(v2), 0)
                    fmt.Println(v2, fi15, v16)
                    v9 -= (((float32(1e10) * v9) + v9) * float32((-2.25 * float64(v9))))
                    _ = fi15
                    _ = v16
                }
            }
        } else {
            var v17 uint16 = 1
            _ = v17
        }
    default:
        switch ("hello" + "a") {
        case "":
            switch -98 {
            case 4294967295:
                switch "" {
                case "a":
                    if (!false) {
                        var v18 int64 = -1
                        _ = v18
                    }
                default:
                    if (42 != 100) {
                        if false {
                            var v19 bool = (!((genFn0() <= 0) || (4294967295 != 4294967295)))
                            v20 := GenS2{F0: v19}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v19
                            _ = v20
                        } else {
                            v3 = append(v3, "")
                            switch -9223372036854775808 {
                            case 135:
                                for fi21 := 0; fi21 < 15; fi21++ {
                                    var v22 int8 = -128
                                    _ = fi21
                                    _ = v22
                                }
                            case 166:
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v24 int64 = -402
                                    _ = fi23
                                    _ = v24
                                }
                            case -1:
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            }
                        }
                        if (!true) {
                            if true {
                                v25 := GenS2{F0: ((32767 < 42) && true)}
                                v9 -= (((v9 / v9) - 0.0) * -2.25)
                                _ = v25
                            }
                        }
                    } else {
                        var v26 int = 6
                        v9 += v9
                        _ = v26
                    }
                }
            }
        case "a":
            var v27 float64 = -2.25
            _ = v27
        }
    }
    v28 := GenS1{F0: genFn0(), F1: -1, F2: (2 == genFn1(0, genFn1(1, 2)))}
    if (false && true) {
        switch "a" {
        case "foo bar":
            v2 -= (v2 + v2)
        case "\n\t":
            v29 := GenS1{F0: -384, F1: -9223372036854775808, F2: false}
            _ = v29
        default:
            if (!true) {
                var v30 uint8 = 207
                _ = v30
            } else {
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 uint = 255
                    if false {
                        if (18446744073709551615 >= 18446744073709551615) {
                            var v33 uint64 = 18446744073709551615
                            if (1000 < 32767) {
                                var v34 int8 = 100
                                _ = v34
                            } else {
                                v9 += float32(1.5)
                            }
                            _ = v33
                        } else {
                            var v35 uint32 = 0
                            _ = v35
                        }
                    } else {
                        v9 = v9
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    var v36 float32 = 1.5
                    _ = fi31
                    _ = v32
                    _ = v36
                }
            }
        }
        if (0.1 >= 1e10) {
            if false {
                var v37 bool = (false && (true || false))
                v38 := GenS3{F0: v2, F1: "a", F2: 127}
                _ = v37
                _ = v38
            }
            switch "foo bar" {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                v2 = (v2 & (v2 | (v2 >> ((2) & 63))))
            }
            fmt.Println(v2, v9)
        }
        var v39 float64 = (-2.25 + 1.5)
        _ = v39
    }
    v40 := GenS4{F0: (v9 != float32((0.1 - 0.1))), F1: 1, F2: -139}
    v3 = append(v3, "foo bar")
    var v41 int = -9223372036854775808
    if ((-1 != 0) && (455 != 2147483647)) {
        v9 = (3.14159 / v9)
        fmt.Println(v9, v2, v2)
    }
    v2 += 4294967295
    for fi42 := 0; fi42 < 11; fi42++ {
        switch 32767 {
        case 413:
            v2 = 42
        case 127:
            if (!(true && (-9223372036854775808 >= -1))) {
                var v43 uint = 0
                switch (("hello" + "hello") + "foo bar") {
                case "hello":
                    v44 := GenS1{F0: 356, F1: genFn0(), F2: true}
                    _ = v44
                default:
                    var v45 int = -817
                    _ = v45
                }
                _ = v43
            } else {
                if (true || true) {
                    if (v2 != 65535) {
                        v9 += (v9 - ((float32((0.1 * -2.25)) - v9) - ((((((v9 - 0.0) - 1.5) + float32(0.1)) * v9) * 3.14159) - ((0.0 * (v9 - 1e10)) * v9))))
                        if ((v2 + v2) <= 42) {
                            fmt.Println(fi42, fi42, v41)
                            _ = len(v3)
                            for fi46 := 0; fi46 < 17; fi46++ {
                                if (1 > 962) {
                                    fmt.Println(v2)
                                    if (((v9 * 0.0) * (1.5 * v9)) == v9) {
                                        var v47 int64 = 42
                                        _ = len(v3)
                                        var v48 uint8 = 0
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v2 += v2
                                        fmt.Println(fi46)
                                        var v49 bool = (1 >= 128)
                                        _ = v49
                                    }
                                } else {
                                    v3 = append(v3, ("a" + ("a" + ("a" + "hello"))))
                                    v9 -= 1e10
                                    if false {
                                        v1 = append(v1, ("" + "a"))
                                        v50 := GenS1{F0: 9223372036854775807, F1: -943, F2: ((v2 / (((v2 + v2)) | 1)) <= v2)}
                                        v2 += v2
                                        switch (v41 ^ (fi46 % ((-9223372036854775808) | 1))) {
                                        case 42:
                                            var v51 int8 = -128
                                            _ = v51
                                        case -2147483648:
                                            switch 0 {
                                            case -157:
                                                if len(v1) > 0 {
                                                    v1[0] = ""
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                            case -577:
                                                var v52 bool = true
                                                switch fi42 {
                                                case 127:
                                                    switch -2147483648 {
                                                    case 2147483647:
                                                        var v53 uint16 = 2
                                                        _ = v53
                                                    }
                                                case 4294967295:
                                                    if false {
                                                        if v52 {
                                                            fmt.Println(fi46, v2, fi42)
                                                            var v54 int32 = -893
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS1{F0: 703, F1: 9223372036854775807, F2: (false && v52)}
                                                            _ = v55
                                                        }
                                                        if (false || v52) {
                                                            if (2 < 706) {
                                                                v2 -= (v2 * (v2 >> ((154) & 63)))
                                                                v56 := GenS2{F0: v52}
                                                                _ = v56
                                                            } else {
                                                                v57 := GenS2{F0: (2 != 65535)}
                                                                if true {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v2 = 42
                                                                    default:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    if (4294967295 == fi42) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v2 = 2147483647
                                                                    }
                                                                    v2 = 4294967295
                                                                    var v58 float32 = v9
                                                                    _ = v58
                                                                }
                                                                _ = v57
                                                            }
                                                            if v52 {
                                                                v59 := GenS1{F0: 9223372036854775807, F1: genFn0(), F2: (!(v52 && (-1 != 1000)))}
                                                                _ = v59
                                                            }
                                                            var v60 uint64 = 2
                                                            _ = v60
                                                        }
                                                        v9 += (v9 * v9)
                                                    }
                                                default:
                                                    v61 := GenS3{F0: v2, F1: "foo bar", F2: -84}
                                                    _ = v61
                                                }
                                                v9 += (0.0 * float32(-2.25))
                                                _ = v52
                                            }
                                        case 4294967295:
                                            switch ("" + "") {
                                            case "hello":
                                                switch (-(v41)) {
                                                case -9223372036854775808:
                                                    v2 -= v2
                                                case 32767:
                                                    v2 = (v2 >> ((255) & 63))
                                                case -305:
                                                    fmt.Println(fi46, fi46, v41)
                                                }
                                            }
                                        }
                                        var v62 int64 = 42
                                        _ = v50
                                        _ = v62
                                    }
                                }
                                _ = fi46
                            }
                        }
                    } else {
                        v9 = 1.5
                    }
                }
                var v63 bool = (!true)
                _ = v63
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v64 := GenS5{F0: 18446744073709551615}
        _ = fi42
        _ = v64
    }
    v41 -= v41
    switch "hello" {
    case "\n\t":
        var v65 float64 = (1e10 - 0.1)
        _ = v65
    case "foo bar":
        v2 = (v2 << ((0) & 63))
    case "":
        for fi66 := 0; fi66 < 7; fi66++ {
            if false {
                var v67 int16 = 216
                _ = v67
            } else {
                var v68 int8 = 1
                if (!(!false)) {
                    fmt.Println(v41)
                } else {
                    switch (fi66 % ((2147483647) | 1)) {
                    case -544:
                        v9 -= 0.0
                    case 32767:
                        for fi69 := 0; fi69 < 1; fi69++ {
                            _ = len(v1)
                            v9 += (-2.25 + float32((float64(3.14159) / 0.1)))
                            _ = fi69
                        }
                    case 9223372036854775807:
                        v3 = append(v3, "\n\t")
                    default:
                        _ = len(v1)
                    }
                    v41 = -2147483648
                }
                var v70 int64 = -9223372036854775808
                _ = v68
                _ = v70
            }
            _ = fi66
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v40
    _ = v41
    _ = v9
}
