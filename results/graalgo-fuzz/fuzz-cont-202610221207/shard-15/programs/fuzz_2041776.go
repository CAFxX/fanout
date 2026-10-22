package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint32; F1 string; F2 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 uint16 }
func genFn0() bool {
    var v6 int16 = -855
    v7 := GenS1{F0: 4294967295}
    var v8 float32 = (3.14159 * 0.1)
    var v9 int32 = -2147483648
    v9 = -2147483648
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (true && ((v9 ^ (-(v9))) != (v9 | 0)))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 uint16 = 0
    var v5 int64 = 9223372036854775807
    v5 -= v5
    switch -171 {
    case 0:
        v3 = append(v3, 2147483647)
    default:
        switch "hello" {
        case "":
            v4 = v4
        case "foo bar":
            v5 = 9223372036854775807
        case "hello":
            var v10 int32 = -2147483648
            _ = v10
        default:
            var v11 uint8 = 9
            _ = v11
        }
    }
    var v12 uint64 = 940
    v5 += v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint = 968
    var v14 int16 = 32767
    v15 := GenS1{F0: 974}
    switch "" {
    case "a":
        v16 := GenS1{F0: 950}
        _ = v16
    default:
        delete(v2, "")
    }
    if true {
        delete(v2, "\n\t")
        if true {
            v3 = append(v3, 128)
            v4 -= (v4 ^ v4)
            _ = len(v3)
        } else {
            if len(v3) > 0 {
                v3[0] = -1
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            if (fi17 >= (fi17 * fi17)) {
                var v18 string = ""
                _ = v18
            } else {
                v19 := GenS1{F0: (fi17 % ((fi17) | 1))}
                switch fi17 {
                case -573:
                    fmt.Println(v4, v4)
                case 649:
                    v12 += (v12 & v12)
                default:
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v21 := GenS2{F0: 483}
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v19
            }
            if true {
                if ((((-809 < 9223372036854775807) || (false && false)) || (v4 == v4)) || (v4 == v4)) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if true {
                            if (!false) {
                                if ((v14 * (v14 + 89)) >= 0) {
                                    switch ((("foo bar" + ("\n\t" + ("hello" + "hello"))) + ("hello" + ("foo bar" + "\n\t"))) + ("foo bar" + ("hello" + "foo bar"))) {
                                    case "":
                                        fmt.Println(v4)
                                    case "\n\t":
                                        if (!(true || false)) {
                                            v23 := GenS2{F0: uint16(v13)}
                                            var v24 string = "\n\t"
                                            var v25 int32 = -612
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            switch fi17 {
                                            case 2147483647:
                                                v26 := GenS2{F0: 255}
                                                _ = v26
                                            case 841:
                                                var v27 uint = (v13 << ((v13) & 63))
                                                _ = v27
                                            }
                                            if false {
                                                if true {
                                                    v28 := GenS2{F0: 513}
                                                    if ((("\n\t" + ("" + "hello")) > ("foo bar" + "\n\t")) || true) {
                                                        v14 -= 1000
                                                    } else {
                                                        if true {
                                                            fmt.Println(v13, fi17, v4)
                                                            var v29 uint32 = 422
                                                            v14 -= v14
                                                            _ = v29
                                                        }
                                                    }
                                                    _ = v28
                                                } else {
                                                    v30 := GenS3{F0: v14, F1: 42, F2: 18}
                                                    var v31 uint64 = v12
                                                    _ = v30
                                                    _ = v31
                                                }
                                                v13 = 885
                                                if (-1 < 2147483647) {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi22
                                                    }
                                                    if (v5 == v5) {
                                                        if len(v3) > 0 {
                                                            v3[0] = -2147483648
                                                        }
                                                        v32 := GenS2{F0: (v4 - (v4 ^ 42))}
                                                        _ = v32
                                                    } else {
                                                        var v33 bool = ((!false) || (88 > 42))
                                                        v34 := GenS4{F0: 2147483648}
                                                        v12 = (v12 >> ((v13) & 63))
                                                        v14 -= v14
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    var v35 float32 = 1e10
                                                    var v36 int64 = (v5 | 947)
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    if ("\n\t" == ("hello" + "\n\t")) {
                                                        if (("\n\t" + ("" + "hello")) == "foo bar") {
                                                            v37 := v1[""]
                                                            if true {
                                                                fmt.Println(v12, fi22, v14)
                                                                v38 := GenS3{F0: -1, F1: 2, F2: 52}
                                                                _ = v38
                                                            } else {
                                                                v4 = 644
                                                            }
                                                            v4 = 255
                                                            var v39 uint32 = 433
                                                            v40 := GenS3{F0: v14, F1: (v39 + v39), F2: 171}
                                                            _ = v37
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    } else {
                                                        if true {
                                                            v3 = append(v3, 9223372036854775807)
                                                        } else {
                                                            v5 += int64(v4)
                                                            v41 := GenS2{F0: 255}
                                                            v12 -= v12
                                                            v42 := GenS4{F0: 2}
                                                            v12 -= (v12 ^ (^(v12)))
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        fmt.Println(v14)
                                                    }
                                                    v43 := GenS5{F0: v12}
                                                    _ = v43
                                                }
                                            } else {
                                                fmt.Println(v12, fi17, v14)
                                                if (0 >= v12) {
                                                    v13 = (v13 - 18446744073709551615)
                                                    switch 644 {
                                                    case -2147483648:
                                                        fmt.Println(fi22)
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v44 := GenS6{F0: 2, F1: 127}
                                                    _ = v44
                                                }
                                                fmt.Println(v5, fi22)
                                            }
                                            v45 := GenS7{F0: 2147483648, F1: "\n\t", F2: ""}
                                            _ = v45
                                        }
                                    case "a":
                                        var v46 uint16 = (^(v4))
                                        _ = v46
                                    default:
                                        v5 = (v5 - -479)
                                    }
                                    _ = len(v1)
                                    if true {
                                        if true {
                                            v47 := GenS2{F0: 32767}
                                            _ = v47
                                        }
                                        v12 = v12
                                    }
                                } else {
                                    v13 -= (v13 + (^(v13)))
                                }
                            } else {
                                var v48 uint32 = 2147483647
                                if true {
                                    _ = len(v1)
                                }
                                var v49 uint8 = uint8(v13)
                                _ = v48
                                _ = v49
                            }
                        }
                        v13 = (v13 & 255)
                        var v50 int16 = v14
                        _ = fi22
                        _ = v50
                    }
                    v4 += (v4 >> (((v13 | (v13 ^ v13))) & 63))
                    v13 = (v13 | (v13 >> ((v13) & 63)))
                } else {
                    delete(v2, ("\n\t" + "foo bar"))
                    if len(v3) > 0 {
                        v3[0] = (fi17 * (fi17 >> ((v13) & 63)))
                    }
                }
            } else {
                v4 = 611
            }
            _ = fi17
        }
    } else {
        switch -27 {
        case 127:
            v12 = (v12 % ((v12) | 1))
        case 2147483647:
            v12 -= v12
        case 32768:
            var v51 int = 9223372036854775807
            _ = v51
        default:
            var v52 float32 = (-2.25 / (0.1 / (float32(float64(1.5)) - (0.1 + 1e10))))
            _ = v52
        }
        _ = len(v2)
        switch -1 {
        case -1:
            v12 -= (v12 >> (((v13 % (((v13 | v13)) | 1))) & 63))
        case 2147483647:
            var v53 float64 = 3.14159
            _ = v53
        default:
            if (v12 < 308) {
                v54 := v2["a"]
                v55 := GenS4{F0: v54}
                _ = v54
                _ = v55
            } else {
                if true {
                    v14 -= v14
                } else {
                    fmt.Println(v5)
                }
            }
        }
    }
    for fi56 := 0; fi56 < 17; fi56++ {
        if false {
            switch ((("foo bar" + "a") + "") + "foo bar") {
            case "\n\t":
                for fi57 := 0; fi57 < 1; fi57++ {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v13)
                    case "a":
                        if false {
                            v58 := v2[(("foo bar" + "\n\t") + (("\n\t" + ("\n\t" + "\n\t")) + "hello"))]
                            v59 := GenS6{F0: 113, F1: -51}
                            _ = v58
                            _ = v59
                        } else {
                            v13 += 0
                        }
                    case "\n\t":
                        var v60 int = 127
                        _ = v60
                    }
                    _ = fi57
                }
            case "":
                if (true || false) {
                    v61 := GenS8{F0: -9223372036854775808}
                    v62 := GenS9{F0: (fi56 - int(v13)), F1: v4}
                    v1[(("\n\t" + (("foo bar" + ("" + "foo bar")) + "")) + ("hello" + "hello"))] = ("foo bar" + "\n\t")
                    _ = v61
                    _ = v62
                } else {
                    fmt.Println(v5, fi56)
                }
            case "foo bar":
                if ((true || (630 > fi56)) && false) {
                    _ = len(v3)
                    v13 -= v13
                    v2["\n\t"] = 42
                }
            }
            var v63 int8 = 42
            fmt.Println(v12, v13, v14)
            _ = v63
        } else {
            fmt.Println(v5, v4, fi56)
            for fi64 := 0; fi64 < 10; fi64++ {
                v14 -= (v14 - v14)
                _ = fi64
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        for fi65 := 0; fi65 < 9; fi65++ {
            if (3.14159 <= (0.0 * float32(float64(((3.14159 / 3.14159) - float32(1e10)))))) {
                fmt.Println(v13, v4, fi65)
                fmt.Println(v4)
            } else {
                fmt.Println(v13)
                v5 = v5
                fmt.Println(v12, v5, fi56)
            }
            _ = fi65
        }
        _ = fi56
    }
    if false {
        if (true && true) {
            var v66 bool = true
            v12 += 2
            _ = v66
        }
    } else {
        v67 := GenS5{F0: (v12 << (((v13 * v13)) & 63))}
        var v68 uint16 = 0
        _ = v67
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
