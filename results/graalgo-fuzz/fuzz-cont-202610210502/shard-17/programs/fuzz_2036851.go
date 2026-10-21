package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32 }
func genFn0(p0 int8, p1 int16) uint {
    p1 = p1
    var v5 string = ""
    p0 = (p0 % ((p0) | 1))
    fmt.Println(p0, v5, p0)
    if true {
        switch v5 {
        case "hello":
            switch -667 {
            case 4294967295:
                v5 = v5
            case 0:
                for fi6 := 0; fi6 < 13; fi6++ {
                    v7 := GenS1{F0: int64(p0)}
                    for fi8 := 0; fi8 < 11; fi8++ {
                        var v9 uint16 = 0
                        var v10 uint16 = v9
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
            }
        case "foo bar":
            p1 -= (p1 & (p1 >> ((0) & 63)))
        }
        var v11 uint32 = 0
        _ = v11
    } else {
        v5 = "foo bar"
    }
    _ = p0
    _ = p1
    _ = v5
    return 255
}
func genFn1() uint16 {
    var v12 uint64 = 42
    fmt.Println(v12)
    _ = v12
    return uint16(v12)
}

func main() {
    v1 := []string{"foo bar", (("hello" + "a") + "\n\t"), "foo bar"}
    v2 := make(map[int]uint32)
    var v3 uint64 = 627
    var v4 int8 = 42
    v13 := v2[42]
    switch "foo bar" {
    case "a":
        for fi14 := 0; fi14 < 12; fi14++ {
            v13 += 373
            _ = fi14
        }
    }
    switch 42 {
    case 128:
        fmt.Println(v3, v4, v3)
    case 42:
        fmt.Println(v3, v3, v13)
    case 9223372036854775807:
        if false {
            for fi15 := 0; fi15 < 11; fi15++ {
                if (!true) {
                    if false {
                        v3 = v3
                    }
                } else {
                    switch ((("hello" + "hello") + "foo bar") + (("foo bar" + "hello") + ("hello" + ("" + "hello")))) {
                    case "a":
                        v3 -= v3
                    default:
                        switch ("\n\t" + ("" + ("hello" + (("\n\t" + (("" + "") + "foo bar")) + "")))) {
                        case "foo bar":
                            var v16 uint32 = v13
                            _ = v16
                        case "":
                            if (0.0 > 0.0) {
                                switch 32768 {
                                case 42:
                                    if true {
                                        fmt.Println(fi15)
                                    }
                                case -423:
                                    if false {
                                        if (!(-2147483648 <= 0)) {
                                            v17 := GenS1{F0: 1}
                                            v18 := GenS2{F0: (97 >= 127), F1: 4294967295}
                                            v4 -= (v4 >> ((4294967295) & 63))
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        for fi19 := 0; fi19 < 4; fi19++ {
                                            v4 = v4
                                            switch (fi19 - (fi19 % ((32768) | 1))) {
                                            case 42:
                                                _ = len(v2)
                                            default:
                                                v20 := GenS2{F0: false, F1: genFn0((v4 % ((v4) | 1)), -32768)}
                                                _ = v20
                                            }
                                            _ = fi19
                                        }
                                        var v21 int = (fi15 ^ fi15)
                                        _ = v21
                                    }
                                }
                                v3 += (v3 ^ uint64(v13))
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if len(v1) > 0 {
                                        v1[0] = (("" + "hello") + "a")
                                    }
                                    _ = fi22
                                }
                            } else {
                                v4 = -128
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch ("hello" + "a") {
                                        case "hello":
                                            v3 += v3
                                        case "\n\t":
                                            var v23 int = (fi15 * 32767)
                                            _ = v23
                                        }
                                    case "foo bar":
                                        switch "a" {
                                        case "foo bar":
                                            var v24 float64 = (float64(1e10) * (0.0 + ((1.5 * float64(1e10)) - ((1.5 * (-2.25 - (float64(3.14159) / 0.1))) * ((1e10 / (1.5 / 3.14159)) + ((0.1 * 0.0) * (3.14159 * (1e10 - 3.14159))))))))
                                            _ = v24
                                        case "":
                                            fmt.Println(fi15)
                                        case "a":
                                            var v25 int = (fi15 ^ fi15)
                                            _ = v25
                                        }
                                    }
                                    v13 = v13
                                    switch fi15 {
                                    case 127:
                                        v4 += (v4 % ((v4) | 1))
                                    default:
                                        var v26 float32 = 0.0
                                        _ = v26
                                    }
                                }
                                v4 += -92
                            }
                        case "\n\t":
                            _ = len(v1)
                        default:
                            switch (fi15 - -412) {
                            case 128:
                                var v27 int8 = -16
                                _ = v27
                            case 32768:
                                v2[fi15] = (v13 << ((18446744073709551615) & 63))
                            }
                        }
                    }
                }
                v13 -= 1
                fmt.Println(fi15, v4, v3)
                _ = fi15
            }
            _ = len(v2)
            if (128 >= 127) {
                fmt.Println(v4)
            } else {
                v28 := GenS2{F0: (!true), F1: 255}
                v29 := GenS3{F0: 1000, F1: (v3 | 1)}
                _ = v28
                _ = v29
            }
        }
    default:
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 string = "hello"
            fmt.Println(fi30, v13)
            delete(v2, fi30)
            _ = fi30
            _ = v31
        }
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 int8 = v4
            var v34 uint64 = 18446744073709551615
            _ = fi32
            _ = v33
            _ = v34
        }
    case 2147483647:
        v3 -= v3
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if true {
            for fi36 := 0; fi36 < 8; fi36++ {
                switch (fi35 - (fi35 | (fi35 / ((fi35) | 1)))) {
                case 32767:
                    var v37 int64 = -1
                    _ = v37
                default:
                    v38 := GenS4{F0: 0, F1: -32768, F2: 2}
                    _ = v38
                }
                if (true && false) {
                    fmt.Println(fi35, v4)
                } else {
                    switch "foo bar" {
                    case "":
                        var v39 uint32 = v13
                        _ = v39
                    case "a":
                        fmt.Println(fi35, v3, fi36)
                    case "hello":
                        var v40 uint16 = 255
                        _ = v40
                    }
                }
                fmt.Println(v3, fi36)
                _ = fi36
            }
            fmt.Println(fi35)
            var v41 bool = (false && (true && true))
            _ = v41
        } else {
            v4 = (v4 | v4)
            var v42 uint = genFn0(v4, 32767)
            if true {
                v43 := GenS1{F0: -1}
                _ = v43
            }
            _ = v42
        }
        switch "\n\t" {
        case "\n\t":
            v3 -= v3
        case "hello":
            var v44 uint64 = v3
            _ = v44
        default:
            fmt.Println(v3)
        }
        v1 = append(v1, "foo bar")
        _ = fi35
    }
    v45 := v2[1]
    for fi46 := 0; fi46 < 12; fi46++ {
        v45 -= v13
        var v47 string = "\n\t"
        v48 := GenS5{F0: v3, F1: -1}
        _ = fi46
        _ = v47
        _ = v48
    }
    var v49 int = 1
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch ((("" + ("hello" + "")) + ("foo bar" + ((("hello" + ("a" + "a")) + "\n\t") + ("a" + "a")))) + "\n\t") {
    case "hello":
        if (!true) {
            for fi50 := 0; fi50 < 15; fi50++ {
                if true {
                    v13 = uint32(v3)
                } else {
                    switch -9223372036854775808 {
                    case 4294967295:
                        switch "a" {
                        case "":
                            fmt.Println(fi50, v45, v4)
                        }
                    case -9223372036854775808:
                        for fi51 := 0; fi51 < 18; fi51++ {
                            delete(v2, 42)
                            _ = fi51
                        }
                    }
                }
                if false {
                    delete(v2, v49)
                    v3 = v3
                } else {
                    delete(v2, v49)
                    delete(v2, (v49 & (v49 ^ fi50)))
                }
                _ = fi50
            }
            v52 := GenS6{F0: genFn0((v4 % ((v4) | 1)), -1)}
            _ = v52
        } else {
            if false {
                var v53 string = "foo bar"
                _ = v53
            }
        }
    case "a":
        for fi54 := 0; fi54 < 20; fi54++ {
            for fi55 := 0; fi55 < 18; fi55++ {
                switch ("" + ("a" + ("" + (("a" + "foo bar") + "a")))) {
                case "foo bar":
                    var v56 float32 = 3.14159
                    _ = v56
                default:
                    if false {
                        fmt.Println(v13, v4, v4)
                        var v57 uint64 = v3
                        var v58 float64 = 0.1
                        v45 = (v45 * (v13 % ((v45) | 1)))
                        _ = v57
                        _ = v58
                    } else {
                        var v59 uint8 = 120
                        _ = v59
                    }
                    v60 := v2[v49]
                    var v61 float32 = float32(float64(1e10))
                    _ = v60
                    _ = v61
                }
                v62 := v2[int(v13)]
                v63 := GenS7{F0: 0, F1: 0}
                _ = fi55
                _ = v62
                _ = v63
            }
            v64 := GenS2{F0: false, F1: genFn0(v4, -1)}
            _ = fi54
            _ = v64
        }
    default:
        v49 -= -626
    }
    var v65 int8 = (v4 + v4)
    v49 -= 32768
    v45 += (v13 - (v13 >> ((979) & 63)))
    switch (("a" + "foo bar") + "hello") {
    case "hello":
        v3 = (v3 | (v3 % ((1) | 1)))
    case "":
        var v66 uint8 = 255
        _ = v66
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v65
}
