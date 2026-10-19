package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 bool; F1 int8; F2 uint }
func genFn0(p0 int8, p1 uint64) bool {
    var v4 float64 = 1.5
    var v5 uint = 42
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint32 = 769
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := []uint32{873, 2147483647, 2147483648}
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v8 float32 = 1.5
    switch -183 {
    case 0:
        var v9 float64 = 0.1
        _ = v9
    case 297:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 255}
            var v12 uint64 = (^(v3))
            switch "" {
            case "\n\t":
                fmt.Println(v12, v8)
            default:
                var v13 int32 = -359
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case 4294967295:
        var v14 int32 = -47
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        if false {
            fmt.Println(v8, v8, fi15)
            var v16 int64 = -9223372036854775808
            for fi17 := 0; fi17 < 6; fi17++ {
                v16 -= v16
                v16 = v16
                switch fi17 {
                case 32768:
                    var v18 float64 = ((float64(v8) * 0.1) / 1e10)
                    _ = v18
                case 42:
                    v19 := GenS1{F0: 2}
                    _ = v19
                case 370:
                    var v20 int = (fi17 % (((fi15 & fi17)) | 1))
                    v3 += 18446744073709551615
                    _ = v20
                }
                _ = fi17
            }
            _ = v16
        } else {
            v21 := GenS1{F0: 0}
            var v22 int16 = 0
            fmt.Println(v22, fi15)
            _ = v21
            _ = v22
        }
        var v23 int8 = 42
        var v24 uint32 = 2
        _ = fi15
        _ = v23
        _ = v24
    }
    v3 = 18446744073709551615
    var v25 int64 = 1
    v1 = append(v1, 882)
    fmt.Println(v8, v25)
    for fi26 := 0; fi26 < 20; fi26++ {
        switch "hello" {
        case "foo bar":
            var v27 float32 = 1e10
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS2{F0: 0, F1: v3, F2: (v3 >> ((18446744073709551615) & 63))}
            _ = fi28
            _ = v29
        }
        v3 = (v3 * (v3 | 42))
        _ = fi26
    }
    var v30 uint8 = 255
    if false {
        switch "" {
        case "foo bar":
            v31 := GenS2{F0: 421, F1: v3, F2: 4294967295}
            _ = v31
        case "\n\t":
            switch -580 {
            case 865:
                for fi32 := 0; fi32 < 3; fi32++ {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        var v34 float32 = 3.14159
                        if true {
                            switch "hello" {
                            case "":
                                _ = len(v2)
                            case "a":
                                var v35 uint64 = (v3 / ((18446744073709551615) | 1))
                                _ = v35
                            }
                            var v36 int32 = -166
                            _ = v36
                        } else {
                            fmt.Println(v25, fi32, fi33)
                            var v37 uint8 = v30
                            _ = v37
                        }
                        _ = fi33
                        _ = v34
                    }
                    switch fi32 {
                    case 936:
                        var v38 uint32 = 42
                        _ = v38
                    case 128:
                        var v39 uint16 = 2
                        _ = v39
                    default:
                        v40 := GenS3{F0: (42 != 302)}
                        _ = v40
                    }
                    _ = fi32
                }
            case 1:
                var v41 uint32 = 2147483648
                _ = v41
            default:
                if (!true) {
                    switch -189 {
                    case 942:
                        var v42 string = ("\n\t" + "")
                        _ = v42
                    case 389:
                        var v43 int64 = v25
                        _ = v43
                    }
                    var v44 float32 = v8
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 4; fi45++ {
                        v30 = v30
                        v30 = 171
                        _ = fi45
                    }
                    for fi46 := 0; fi46 < 18; fi46++ {
                        v8 -= (v8 / v8)
                        for fi47 := 0; fi47 < 9; fi47++ {
                            v48 := GenS2{F0: (v3 & v3), F1: (v3 + 1), F2: v3}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v49 string = ""
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        v2[(("foo bar" + ("hello" + "hello")) + "")] = ("" + "\n\t")
                        _ = fi46
                    }
                    v50 := GenS3{F0: (-213 >= (v25 - (v25 + v25)))}
                    _ = v50
                }
            }
        default:
            v51 := GenS4{F0: uint16(v30), F1: -1, F2: 9223372036854775807}
            _ = v51
        }
    } else {
        var v52 float32 = v8
        var v53 int8 = 104
        _ = v52
        _ = v53
    }
    if len(v1) > 0 {
        v1[0] = uint32(v30)
    }
    if (!(0 >= v25)) {
        v54 := v2[((("a" + (("foo bar" + "hello") + "\n\t")) + "hello") + "a")]
        for fi55 := 0; fi55 < 17; fi55++ {
            if (!false) {
                var v56 uint16 = 962
                for fi57 := 0; fi57 < 19; fi57++ {
                    v30 -= v30
                    var v58 float64 = (float64(v8) - (3.14159 * ((((0.1 + float64((v8 - v8))) + 1e10) + float64(((1e10 + -2.25) - (1.5 / 1.0)))) + -2.25)))
                    _ = fi57
                    _ = v58
                }
                var v59 uint16 = (v56 | v56)
                _ = v56
                _ = v59
            }
            v54 = "foo bar"
            _ = fi55
        }
        _ = v54
    } else {
        for fi60 := 0; fi60 < 3; fi60++ {
            switch ((("" + ("foo bar" + "hello")) + "\n\t") + "\n\t") {
            case "\n\t":
                var v61 uint = 4294967295
                _ = v61
            default:
                if true {
                    v62 := GenS2{F0: 0, F1: (v3 << ((910) & 63)), F2: (v3 ^ 1)}
                    _ = len(v2)
                    _ = v62
                } else {
                    _ = len(v1)
                    v8 += 1.5
                }
            }
            _ = fi60
        }
        switch 2147483647 {
        case 32768:
            fmt.Println(v3, v25)
        }
        switch -686 {
        case -977:
            var v63 uint8 = v30
            _ = v63
        case -666:
            if (631 != 65535) {
                fmt.Println(v3, v3, v30)
                if true {
                    v8 += 1.5
                }
                _ = len(v1)
            } else {
                var v64 int32 = 0
                for fi65 := 0; fi65 < 20; fi65++ {
                    for fi66 := 0; fi66 < 8; fi66++ {
                        v67 := GenS1{F0: (v30 ^ (v30 - v30))}
                        fmt.Println(v30, v8, v64)
                        switch fi66 {
                        case -1:
                            fmt.Println(fi65)
                        case 0:
                            v8 = (1e10 / (v8 * v8))
                        case -365:
                            v25 += (v25 | v25)
                        }
                        v68 := GenS5{F0: true, F1: 0}
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                    v69 := GenS6{F0: (float32(float64(((3.14159 - v8) - (((v8 + v8) + v8) / v8)))) > 3.14159), F1: 127, F2: 2}
                    _ = fi65
                    _ = v69
                }
                _ = v64
            }
        default:
            v2[""] = "hello"
        }
    }
    var v70 int32 = 1
    v30 -= (v30 % ((v30) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v70
    _ = v8
}
