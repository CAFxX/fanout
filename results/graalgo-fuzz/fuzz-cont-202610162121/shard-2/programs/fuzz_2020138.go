package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 int, p1 uint16) string {
    p1 = (^(p1))
    var v5 uint8 = 2
    v6 := []int{p0, (p0 % (((p0 << ((1) & 63))) | 1)), p0, 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (("foo bar" + ("" + ("\n\t" + "\n\t"))) + "a")
}
func genFn1(p0 uint64, p1 int16, p2 int32) int16 {
    var v7 uint = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = -128
        v10 := []int{fi8, fi8}
        _ = fi8
        _ = v10
        _ = v9
    }
    p2 -= p2
    v11 := []int{128, 42, -935, -181}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return (p1 >> ((v7) & 63))
}

func main() {
    var v1 int8 = 127
    var v2 bool = (54 == 0)
    v3 := []uint32{255, 42, 150}
    v4 := []string{(("hello" + ("foo bar" + ((("" + "hello") + "") + ""))) + ""), "\n\t", "a"}
    v1 = v1
    v1 = 45
    v4 = append(v4, "")
    var v12 float64 = (((1e10 / float64(-2.25)) + -2.25) / 3.14159)
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: v1}
        _ = fi13
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!v2) {
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                if (!(!(!v2))) {
                    var v17 uint64 = 4294967295
                    v12 = 0.1
                    _ = v17
                } else {
                    var v18 string = ("hello" + ((("" + "hello") + ("hello" + "")) + ""))
                    switch fi15 {
                    case 42:
                        fmt.Println(fi16, fi16)
                    case 0:
                        fmt.Println(v1, fi16)
                    case 4294967295:
                        var v19 uint16 = 65535
                        _ = v19
                    }
                    var v20 int = (^(fi15))
                    _ = v18
                    _ = v20
                }
                v21 := GenS2{F0: 534, F1: -2147483648}
                _ = fi16
                _ = v21
            }
            _ = fi15
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        _ = len(v4)
        var v23 int64 = -1
        switch fi22 {
        case 42:
            v24 := GenS2{F0: 1, F1: 1}
            _ = v24
        case 0:
            switch "" {
            case "a":
                switch (fi22 | fi22) {
                case 2147483647:
                    var v25 string = ("" + "")
                    _ = v25
                case -859:
                    v12 = v12
                default:
                    if ("\n\t" > "") {
                        if (88 != 53) {
                            if (-2147483648 != 2147483647) {
                                if (!true) {
                                    v12 = v12
                                    if false {
                                        for fi26 := 0; fi26 < 10; fi26++ {
                                            if true {
                                                v4 = append(v4, ("a" + ("foo bar" + ("\n\t" + "hello"))))
                                                v1 -= v1
                                                v12 -= v12
                                            } else {
                                                if ("" == ("foo bar" + "hello")) {
                                                    if (v2 && v2) {
                                                        v27 := GenS1{F0: v1}
                                                        _ = v27
                                                    } else {
                                                        fmt.Println(v1, v2, v2)
                                                    }
                                                    var v28 float32 = 1e10
                                                    var v29 uint8 = 163
                                                    if (v2 || v2) {
                                                        if (0 > 18446744073709551615) {
                                                            fmt.Println(v12, fi26)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v28 -= float32(((v12 * v12) * v12))
                                                        }
                                                        if (((v12 / ((float64((v28 / (v28 + v28))) + ((0.0 + (v12 / v12)) * v12)) - ((float64(v28) / (((v12 - (v12 - v12)) - 1e10) * v12)) / v12))) / 3.14159) >= v12) {
                                                            v30 := GenS2{F0: 1, F1: 0}
                                                            _ = v30
                                                        } else {
                                                            var v31 float64 = (float64(0.1) + 3.14159)
                                                            var v32 uint64 = 42
                                                            v2 = v2
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = "\n\t"
                                                        }
                                                        v33 := GenS2{F0: 4294967295, F1: 2147483647}
                                                        _ = v33
                                                    }
                                                    fmt.Println(v1)
                                                    switch fi22 {
                                                    case 2147483647:
                                                        if false {
                                                            v34 := GenS2{F0: 0, F1: -1}
                                                            if true {
                                                                v29 += 252
                                                            } else {
                                                                var v35 string = ("hello" + "foo bar")
                                                                if true {
                                                                    if v2 {
                                                                        if true {
                                                                            var v36 int = 128
                                                                            var v37 bool = true
                                                                            _ = v36
                                                                            _ = v37
                                                                        } else {
                                                                            v38 := GenS1{F0: v1}
                                                                            var v39 bool = (2 < 828)
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        v35 = v35
                                                                        v40 := GenS3{F0: 0, F1: -2147483648, F2: 18446744073709551615}
                                                                        _ = v40
                                                                    } else {
                                                                        var v41 uint32 = uint32(v29)
                                                                        v42 := GenS4{F0: 2, F1: v2}
                                                                        switch (fi22 / (((fi22 ^ 123)) | 1)) {
                                                                        case 1:
                                                                            var v43 string = ((v35 + "\n\t") + v35)
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case "":
                                                                                v4 = append(v4, v35)
                                                                            case "foo bar":
                                                                                var v44 float32 = v28
                                                                                _ = v44
                                                                            }
                                                                            var v45 int8 = 127
                                                                            var v46 int = (fi26 - 32767)
                                                                            v23 += int64(v41)
                                                                            _ = v43
                                                                            _ = v45
                                                                            _ = v46
                                                                        default:
                                                                            switch fi22 {
                                                                            case -9223372036854775808:
                                                                                v47 := GenS5{F0: v41, F1: 211}
                                                                                v48 := GenS2{F0: 1, F1: 652}
                                                                                var v49 float64 = v12
                                                                                _ = v47
                                                                                _ = v48
                                                                                _ = v49
                                                                            default:
                                                                                if (!(!(18446744073709551615 >= 255))) {
                                                                                    fmt.Println(v29)
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v41
                                                                        _ = v42
                                                                    }
                                                                    v12 = -2.25
                                                                } else {
                                                                    _ = len(v3)
                                                                    switch v35 {
                                                                    case "hello":
                                                                        fmt.Println(v1, v1)
                                                                    case "\n\t":
                                                                        fmt.Println(v1, v29, v12)
                                                                    }
                                                                }
                                                                _ = v35
                                                            }
                                                            _ = v34
                                                        } else {
                                                            switch (("foo bar" + ((((("\n\t" + "hello") + "hello") + ("" + ("foo bar" + ""))) + "hello") + "")) + ("a" + (("foo bar" + "a") + ("a" + "")))) {
                                                            case "":
                                                                if (42 > v23) {
                                                                    v50 := GenS2{F0: 255, F1: -1}
                                                                    var v51 int = -203
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            }
                                                            var v52 float64 = 0.1
                                                            if v2 {
                                                                var v53 int32 = -949
                                                                _ = v53
                                                            } else {
                                                                var v54 uint32 = 2
                                                                _ = v54
                                                            }
                                                            _ = v52
                                                        }
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v55 := GenS6{F0: 634, F1: 4294967295}
                                                _ = v55
                                            }
                                            _ = fi26
                                        }
                                        v12 = v12
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        v23 += v23
                                    }
                                }
                                fmt.Println(v23, v23, v2)
                            }
                            v1 = v1
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 = v1
                    } else {
                        v23 -= -9223372036854775808
                        fmt.Println(v2, v23)
                    }
                }
            case "foo bar":
                v4 = append(v4, (("\n\t" + "hello") + (("hello" + "") + ("a" + "\n\t"))))
            case "":
                v23 -= 32
            }
        case 4294967295:
            v56 := GenS3{F0: 195, F1: 42, F2: 0}
            _ = v56
        default:
            var v57 bool = (!(v2 || ((42 >= 31) && (false && (genFn1(42, -32768, -2147483648) >= -32768)))))
            _ = v57
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
