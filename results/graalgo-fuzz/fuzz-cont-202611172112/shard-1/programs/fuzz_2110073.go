package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 bool; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint) string {
    v5 := GenS1{F0: -31, F1: false}
    if true {
        p0 = p0
        for fi6 := 0; fi6 < 1; fi6++ {
            p0 -= (p0 << (((p0 - 209)) & 63))
            _ = fi6
        }
        var v7 int32 = 344
        _ = v7
    } else {
        p0 += (p0 % ((p0) | 1))
        p0 = p0
    }
    p0 += 42
    var v8 uint8 = 2
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 string = "foo bar"
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v8
    return ""
}

func main() {
    var v1 uint16 = 1
    var v2 bool = false
    var v3 bool = v2
    var v4 string = "hello"
    var v11 int8 = -45
    v12 := []int{-898}
    v11 += v11
    fmt.Println(v3, v2, v2)
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 1, F1: v2}
        var v15 uint64 = 1
        v11 += (-(v11))
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = 1
    v17 := GenS1{F0: (v11 << ((903) & 63)), F1: v3}
    var v18 float32 = 1.5
    var v19 string = v4
    switch 32768 {
    case 436:
        v12 = append(v12, 128)
    case 32768:
        for fi20 := 0; fi20 < 14; fi20++ {
            v12 = append(v12, (fi20 - (fi20 * 32767)))
            _ = fi20
        }
    case 127:
        var v21 uint32 = 42
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(v4, v4)
        switch 42 {
        case 42:
            v18 -= (0.1 * (3.14159 + v18))
        default:
            v2 = v2
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            v4 = v19
            if (!false) {
                switch (v19 + "a") {
                case "foo bar":
                    v19 = v19
                case "":
                    if (-330 != -340) {
                        if (v2 && (false && (1 < (fi23 | 1)))) {
                            if (0.1 < -2.25) {
                                v16 -= 18446744073709551615
                                v11 -= -1
                            }
                            v24 := GenS1{F0: (v11 | (v11 << ((4294967295) & 63))), F1: v2}
                            fmt.Println(v16, v19, v3)
                            _ = v24
                        }
                        v3 = (993 >= 65535)
                        fmt.Println(v18, v18)
                    }
                case "\n\t":
                    var v25 float64 = float64(float32(((0.1 + float64(v18)) * (float64((1e10 / ((0.1 / 1e10) + v18))) - 0.0))))
                    _ = v25
                default:
                    v26 := GenS1{F0: 0, F1: v2}
                    _ = v26
                }
                var v27 float32 = v18
                v27 = 0.1
                _ = v27
            }
            _ = fi23
        }
        _ = fi22
    }
    v4 = "a"
    if (773 < 0) {
        v11 += (v11 >> ((18446744073709551615) & 63))
        v12 = append(v12, 9223372036854775807)
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v11)
            v29 := GenS2{F0: 570, F1: (^(v1))}
            if (v19 != v4) {
                if false {
                    v1 = v1
                    if (42 < 615) {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            fmt.Println(v1)
                            _ = fi30
                        }
                        fmt.Println(v11, fi28)
                    }
                } else {
                    switch fi28 {
                    case 109:
                        fmt.Println(v4, v19, v3)
                    case 4294967295:
                        v19 = (v4 + v4)
                    case 128:
                        v31 := GenS3{F0: 1}
                        _ = v31
                    }
                    if true {
                        switch v19 {
                        case "":
                            for fi32 := 0; fi32 < 9; fi32++ {
                                v33 := GenS4{F0: 924, F1: 42}
                                fmt.Println(v2, v18, fi32)
                                if v2 {
                                    if v2 {
                                        if true {
                                            if false {
                                                fmt.Println(v2)
                                                var v34 uint16 = v1
                                                _ = v34
                                            }
                                        } else {
                                            switch (v19 + v19) {
                                            case "hello":
                                                v35 := GenS5{F0: (v11 / (((v11 >> ((1) & 63))) | 1))}
                                                _ = v35
                                            case "\n\t":
                                                v18 = ((0.1 - 0.0) * v18)
                                            }
                                            fmt.Println(v1, v1)
                                        }
                                        v36 := GenS6{F0: 150}
                                        _ = v36
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                                _ = fi32
                                _ = v33
                            }
                        }
                        if v2 {
                            switch (("foo bar" + ((v4 + ((v4 + v4) + ("a" + ("a" + v19)))) + (("" + ("a" + v4)) + (v4 + "a")))) + "\n\t") {
                            case "a":
                                if len(v12) > 0 {
                                    v12[0] = -782
                                }
                            case "":
                                var v37 int = (fi28 - (^(fi28)))
                                _ = v37
                            case "hello":
                                v38 := GenS5{F0: -128}
                                _ = v38
                            default:
                                fmt.Println(v16)
                            }
                            v19 = v19
                            for fi39 := 0; fi39 < 13; fi39++ {
                                var v40 int16 = int16(v16)
                                _ = fi39
                                _ = v40
                            }
                        }
                        if (v2 && ((105 < 128) || true)) {
                            _ = len(v12)
                        }
                    } else {
                        _ = len(v12)
                        if (v3 || true) {
                            var v41 uint16 = uint16(v16)
                            fmt.Println(v41)
                            switch fi28 {
                            case 9223372036854775807:
                                v42 := GenS7{F0: 18446744073709551615, F1: (v3 && (fi28 < fi28)), F2: 149}
                                _ = v42
                            case 32767:
                                if ((0.0 + 3.14159) > (float64(float32((float64(float32(float64(v18))) * 1e10))) / float64((1e10 * 1.5)))) {
                                    if (false || (!(v2 && ((fi28 * (fi28 - 1)) <= fi28)))) {
                                        switch v19 {
                                        case "a":
                                            if v3 {
                                                if v3 {
                                                    var v43 int8 = 1
                                                    _ = v43
                                                } else {
                                                    if v3 {
                                                        for fi44 := 0; fi44 < 10; fi44++ {
                                                            v12 = append(v12, fi44)
                                                            var v45 int = 42
                                                            if (true && true) {
                                                                v19 = v19
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                                var v46 int8 = (^(v11))
                                                                if (v3 || (1 > -32768)) {
                                                                    v47 := GenS7{F0: 146, F1: v2, F2: 42}
                                                                    var v48 uint = 2
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                if true {
                                                                    var v49 uint = uint(v16)
                                                                    _ = v49
                                                                }
                                                                _ = v46
                                                            } else {
                                                                var v50 uint64 = (v16 | 0)
                                                                _ = v50
                                                            }
                                                            _ = fi44
                                                            _ = v45
                                                        }
                                                        for fi51 := 0; fi51 < 10; fi51++ {
                                                            fmt.Println(v16)
                                                            _ = fi51
                                                        }
                                                    }
                                                }
                                                v16 -= v16
                                            } else {
                                                for fi52 := 0; fi52 < 3; fi52++ {
                                                    var v53 int = -2147483648
                                                    v54 := GenS8{F0: 57}
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v55 int8 = (v11 * v11)
                                                        _ = v55
                                                    case "a":
                                                        var v56 int = fi28
                                                        _ = v56
                                                    case "hello":
                                                        if v2 {
                                                            var v57 bool = (false && true)
                                                            _ = v57
                                                        }
                                                    default:
                                                        fmt.Println(v11, v53, v1)
                                                    }
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                        }
                                        v3 = v2
                                    }
                                    v58 := GenS8{F0: -9223372036854775808}
                                    _ = v58
                                }
                            }
                            _ = v41
                        } else {
                            if false {
                                v18 -= 1e10
                            } else {
                                var v59 uint8 = 1
                                _ = v59
                            }
                        }
                        fmt.Println(v19, v18)
                    }
                }
                v60 := GenS2{F0: 25, F1: v1}
                _ = v60
            } else {
                v61 := GenS9{F0: 42, F1: 0}
                _ = v61
            }
            _ = fi28
            _ = v29
        }
    }
    if true {
        fmt.Println(v16)
        var v62 int = 128
        if (!v3) {
            for fi63 := 0; fi63 < 13; fi63++ {
                var v64 uint32 = 65535
                v12 = append(v12, v62)
                _ = fi63
                _ = v64
            }
            fmt.Println(v16, v11, v1)
            fmt.Println(v11, v2, v18)
        }
        _ = v62
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        fmt.Println(v19)
    case "hello":
        var v65 string = ""
        _ = v65
    default:
        if v3 {
            v3 = (false || ((42 >= 1) || v3))
        }
    }
    for fi66 := 0; fi66 < 15; fi66++ {
        var v67 int64 = -1
        var v68 uint64 = v16
        _ = fi66
        _ = v67
        _ = v68
    }
    var v69 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v69
}
