package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 int; F2 int64 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 int64) int {
    for fi6 := 0; fi6 < 17; fi6++ {
        p1 -= (p1 & p1)
        v7 := GenS1{F0: 42}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    v8 := []int{1, 128, -349, 257}
    _ = p0
    _ = p1
    _ = v8
    return 32767
}
func genFn1(p0 int16, p1 int32) int16 {
    v9 := []int{}
    fmt.Println(p0, p1)
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 bool = (685 != 65535)
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "foo bar":
        p1 = (p1 << ((698) & 63))
    case "hello":
        p1 -= (-(p1))
    case "\n\t":
        p1 -= 42
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    var v1 int16 = 1000
    v2 := []string{"\n\t", ("a" + "\n\t"), "foo bar"}
    var v3 uint16 = 382
    v4 := []string{("foo bar" + ("" + "\n\t"))}
    var v5 int = 0
    var v12 uint64 = 578
    v3 = uint16(v12)
    for fi13 := 0; fi13 < 12; fi13++ {
        switch "hello" {
        case "foo bar":
            var v14 int8 = 1
            _ = v14
        case "hello":
            v5 = 1
        case "":
            for fi15 := 0; fi15 < 1; fi15++ {
                if false {
                    var v16 uint64 = 52
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + "")
                    }
                    _ = v16
                } else {
                    v12 += (v12 >> ((65535) & 63))
                    v12 += (v12 >> ((uint(v12)) & 63))
                    if true {
                        v17 := GenS1{F0: genFn1(genFn1(v1, 0), -1)}
                        if false {
                            v5 = 128
                        }
                        _ = v17
                    } else {
                        var v18 int = fi15
                        var v19 int16 = v1
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi15
            }
        default:
            switch fi13 {
            case 128:
                v5 = (fi13 & (fi13 >> ((0) & 63)))
            case -1:
                switch fi13 {
                case -9223372036854775808:
                    var v20 float32 = (0.1 / -2.25)
                    _ = v20
                case 42:
                    _ = len(v2)
                case 0:
                    if (!(uint8(v12) <= 0)) {
                        switch ((("hello" + "hello") + "foo bar") + "foo bar") {
                        case "foo bar":
                            if (((0.1 + 0.1) + ((1e10 + (float64(3.14159) / 1.0)) - (-2.25 - 3.14159))) != 0.1) {
                                v5 = fi13
                                v21 := GenS1{F0: (v1 + -865)}
                                _ = v21
                            } else {
                                v4 = append(v4, ((("\n\t" + "") + "foo bar") + "a"))
                            }
                        case "hello":
                            if (true && (1 > 0)) {
                                var v22 uint8 = 0
                                switch "" {
                                case "a":
                                    if (false && true) {
                                        v23 := GenS2{F0: (907 != v12), F1: 1}
                                        var v24 float64 = 0.1
                                        v22 = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        if true {
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = "a"
                                                }
                                            }
                                        }
                                        var v25 string = "hello"
                                        _ = v25
                                    }
                                }
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    if (!false) {
                                        var v27 uint32 = 700
                                        _ = v27
                                    }
                                    v5 = fi13
                                    _ = fi26
                                }
                                _ = v22
                            }
                        default:
                            if (-1 > 0) {
                                var v28 bool = (42 != uint32(v3))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v29 := GenS1{F0: (v1 * (v1 - (v1 / ((genFn1(32767, 42)) | 1))))}
                                _ = v28
                                _ = v29
                            } else {
                                fmt.Println(v3)
                            }
                        }
                        for fi30 := 0; fi30 < 17; fi30++ {
                            fmt.Println(v3, v12, fi30)
                            var v31 uint64 = v12
                            if true {
                                v5 += int(v1)
                                switch "foo bar" {
                                case "hello":
                                    v3 += 1
                                    v3 = v3
                                case "a":
                                    fmt.Println(v1, v5)
                                }
                                var v32 uint64 = v12
                                if false {
                                    if ((float64(float32(1e10)) - (0.0 * (float64(1.5) - float64((float32(1e10) + 1e10))))) == 0.1) {
                                        var v33 uint8 = 2
                                        v34 := GenS3{F0: 18446744073709551615, F1: (v12 % ((uint64(v3)) | 1))}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v12 = v12
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v3 -= (v3 % ((1) | 1))
                                    }
                                    v35 := GenS3{F0: 65535, F1: (v12 << ((18446744073709551615) & 63))}
                                    switch ("\n\t" + "hello") {
                                    case "a":
                                        v36 := GenS4{F0: -2147483648}
                                        _ = v36
                                    case "\n\t":
                                        if true {
                                            var v37 int32 = -1
                                            switch fi13 {
                                            case 0:
                                                if true {
                                                    v1 -= (v1 * (v1 * 898))
                                                    v3 -= (v3 * v3)
                                                    var v38 int16 = 974
                                                    if len(v2) > 0 {
                                                        v2[0] = (("" + "\n\t") + ((("hello" + ("\n\t" + "foo bar")) + "hello") + "foo bar"))
                                                    }
                                                    if true {
                                                        v12 = (v31 & (v12 >> ((4294967295) & 63)))
                                                        var v39 uint8 = 2
                                                        _ = v39
                                                    }
                                                    _ = v38
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v40 float64 = 3.14159
                                                    _ = v40
                                                }
                                            case -2147483648:
                                                _ = len(v2)
                                            case 127:
                                                v41 := GenS1{F0: (v1 + (v1 * 1000))}
                                                _ = v41
                                            }
                                            var v42 int = fi13
                                            _ = v37
                                            _ = v42
                                        } else {
                                            switch ("a" + ("" + "\n\t")) {
                                            case "hello":
                                                var v43 float32 = ((float32(0.1) * ((((0.1 + 1e10) + ((1e10 + ((1.5 - 1e10) + float32(1.5))) - (0.0 * 0.1))) * 1.5) - 1e10)) * (-2.25 / (-2.25 * (1.5 / 1.0))))
                                                _ = v43
                                            }
                                            v44 := GenS4{F0: 583}
                                            v31 = v31
                                            _ = v44
                                        }
                                    }
                                    _ = v35
                                }
                                var v45 int8 = 94
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v3 -= v3
                                v2 = append(v2, "hello")
                                _ = v32
                                _ = v45
                            }
                            var v46 int16 = genFn1((v1 + (v1 << ((4294967295) & 63))), -934)
                            _ = fi30
                            _ = v31
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 18; fi47++ {
                            if (false || (!false)) {
                                v48 := GenS5{F0: false, F1: -1, F2: "\n\t"}
                                if ((v1 * -1) >= (v1 * v1)) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v1 = 0
                                    default:
                                        fmt.Println(v5)
                                    }
                                } else {
                                    if (true || (3.14159 == 1.5)) {
                                        var v49 uint8 = 113
                                        v12 -= v12
                                        switch "a" {
                                        case "foo bar":
                                            var v50 uint32 = 4294967295
                                            _ = v50
                                        default:
                                            v2 = append(v2, "\n\t")
                                        }
                                        v49 = 2
                                        _ = v49
                                    } else {
                                        switch (fi13 >> ((18446744073709551615) & 63)) {
                                        case -9223372036854775808:
                                            if false {
                                                fmt.Println(fi13, v1)
                                                fmt.Println(fi13, v1, v5)
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    v51 := GenS2{F0: true, F1: -2147483648}
                                                    _ = v51
                                                case "a":
                                                    v1 = (v1 - int16(v12))
                                                }
                                                fmt.Println(fi13, v1)
                                            }
                                        default:
                                            if (!false) {
                                                _ = len(v2)
                                            } else {
                                                if true {
                                                    v52 := GenS6{F0: (v12 | v12)}
                                                    v1 += (v1 / ((v1) | 1))
                                                    _ = v52
                                                }
                                                v53 := GenS7{F0: (("" + (("a" + "\n\t") + ("" + "\n\t"))) <= "hello"), F1: (v5 >> ((42) & 63)), F2: -1}
                                                var v54 float32 = -2.25
                                                switch "foo bar" {
                                                case "hello":
                                                    v5 = (v5 >> ((4294967295) & 63))
                                                    var v55 float32 = 3.14159
                                                    v55 = 1e10
                                                    _ = v55
                                                case "":
                                                    v3 = v3
                                                default:
                                                    v3 += (v3 * v3)
                                                }
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        fmt.Println(v12)
                                        v2 = append(v2, "")
                                    }
                                }
                                if (!(false || true)) {
                                    v56 := GenS8{F0: v3, F1: (3.14159 <= ((3.14159 * 0.0) - (-2.25 / (-2.25 * 3.14159))))}
                                    _ = v56
                                } else {
                                    v57 := GenS9{F0: (42 >= v1)}
                                    v58 := GenS10{F0: 0, F1: 42}
                                    _ = v57
                                    _ = v58
                                }
                                _ = v48
                            }
                            v59 := GenS8{F0: 32767, F1: (357 <= 0)}
                            switch "a" {
                            case "foo bar":
                                v1 = 32767
                            case "":
                                var v60 float64 = -2.25
                                _ = v60
                            }
                            _ = fi47
                            _ = v59
                        }
                    }
                default:
                    v3 -= 680
                }
            default:
                switch int(v12) {
                case 127:
                    if len(v4) > 0 {
                        v4[0] = (("hello" + "hello") + "")
                    }
                default:
                    v61 := GenS8{F0: 754, F1: (!(18446744073709551615 >= 42))}
                    _ = v61
                }
            }
        }
        _ = fi13
    }
    var v62 uint8 = uint8(v3)
    v1 = (v1 * (v1 & 0))
    v3 -= v3
    v63 := GenS6{F0: v12}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= v1
    var v64 int16 = (v1 * 1)
    v65 := GenS4{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
