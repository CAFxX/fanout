package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int8; F2 int32 }
type GenS9 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 int16, p1 string, p2 int8) int64 {
    p0 = (p0 / ((58) | 1))
    if false {
        switch 4294967295 {
        case -775:
            v7 := []int{0, 32767, -1, 629}
            _ = v7
        case 2147483647:
            p2 += p2
        default:
            v8 := []int{762}
            _ = v8
        }
        v9 := []int{9223372036854775807}
        _ = v9
    }
    if (1 > 252) {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 float32 = 3.14159
            v12 := []int{(fi10 >> ((310) & 63))}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int8 = -15
        if true {
            p0 -= 136
            if false {
                var v15 uint64 = 2
                if (42 > 819) {
                    if ("\n\t" == (p1 + "hello")) {
                        if true {
                            v15 = (v15 + v15)
                        }
                    }
                } else {
                    v16 := GenS1{F0: false, F1: "foo bar", F2: 65535}
                    for fi17 := 0; fi17 < 17; fi17++ {
                        p2 -= (v14 << ((1) & 63))
                        _ = fi17
                    }
                    _ = v16
                }
                var v18 uint32 = 0
                _ = v15
                _ = v18
            }
        }
        if (42 <= 2147483647) {
            for fi19 := 0; fi19 < 4; fi19++ {
                if ((fi13 << ((21) & 63)) < fi13) {
                    var v20 float32 = 1e10
                    v21 := []int{(fi13 >> ((42) & 63))}
                    v22 := GenS2{F0: 194, F1: 42}
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi19
            }
            v23 := []int{4294967295, (fi13 & 32767), (-(fi13)), -136}
            for fi24 := 0; fi24 < 2; fi24++ {
                switch "" {
                case "":
                    var v25 uint = 0
                    _ = v25
                }
                _ = fi24
            }
            _ = v23
        }
        _ = fi13
        _ = v14
    }
    var v26 float32 = float32(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 0
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    v3 := []int{32768, 32767, 1}
    var v4 string = ""
    var v5 int16 = -234
    v6 := make(map[string]bool)
    if (v1 != (v1 | v1)) {
        var v27 int32 = 42
        v28 := GenS3{F0: v27}
        _ = v27
        _ = v28
    } else {
        var v29 uint8 = 1
        _ = v29
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 int16 = (v5 << ((0) & 63))
        v3 = append(v3, 42)
        _ = fi30
        _ = v31
    }
    var v32 float32 = (1e10 / -2.25)
    var v33 float64 = 3.14159
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v3)
        v35 := GenS4{F0: -1}
        _ = fi34
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 = (v5 & v5)
        } else {
            if true {
                if (!(4294967295 >= 42)) {
                    var v37 int = 0
                    _ = v37
                } else {
                    v5 += (v5 + (v5 | -1))
                }
                switch (fi36 / (((fi36 - -357)) | 1)) {
                case 1:
                    v2[4294967295] = 42
                default:
                    if (-2.25 > 0.1) {
                        var v38 float32 = (1.5 * v32)
                        fmt.Println(v33, v32)
                        v33 = (v33 + v33)
                        _ = v38
                    }
                }
            } else {
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 uint16 = 0
                    switch 1 {
                    case -2147483648:
                        v41 := GenS5{F0: v1, F1: v5}
                        _ = v41
                    default:
                        v42 := GenS2{F0: uint8(v40), F1: 30}
                        switch ((v4 + ("hello" + v4)) + (v4 + "hello")) {
                        case "hello":
                            if (v40 > v40) {
                                switch 4294967295 {
                                case -707:
                                    delete(v2, 1)
                                default:
                                    var v43 uint16 = (v40 + 473)
                                    _ = v43
                                }
                                v44 := GenS5{F0: 9223372036854775807, F1: 627}
                                v45 := GenS6{F0: 762}
                                _ = v44
                                _ = v45
                            }
                        }
                        var v46 int8 = 1
                        switch fi39 {
                        case 127:
                            fmt.Println(v1)
                        }
                        v3 = append(v3, (^(fi39)))
                        _ = v42
                        _ = v46
                    }
                    _ = fi39
                    _ = v40
                }
            }
        }
        v5 -= v5
        _ = fi36
    }
    if false {
        v47 := GenS1{F0: (-83 <= -128), F1: "\n\t", F2: 65535}
        for fi48 := 0; fi48 < 16; fi48++ {
            v2[462] = 16
            var v49 int16 = (v5 >> ((982) & 63))
            _ = fi48
            _ = v49
        }
        _ = v47
    } else {
        for fi50 := 0; fi50 < 3; fi50++ {
            v51 := GenS7{F0: 1}
            delete(v2, 818)
            v52 := GenS8{F0: v4, F1: 127, F2: 0}
            _ = fi50
            _ = v51
            _ = v52
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (false || (true && false)) {
        if (!(1.5 >= v32)) {
            fmt.Println(v4, v32, v33)
        } else {
            fmt.Println(v1, v1)
            var v53 int64 = v1
            _ = v53
        }
    } else {
        if (genFn0(1000, ("foo bar" + v4), -102) > v1) {
            var v54 int32 = 2147483647
            _ = v54
        } else {
            var v55 int32 = 42
            _ = v55
        }
        _ = len(v3)
    }
    for fi56 := 0; fi56 < 2; fi56++ {
        switch (-(fi56)) {
        case 32768:
            v1 = int64(fi56)
        case 32767:
            switch fi56 {
            case -1:
                for fi57 := 0; fi57 < 9; fi57++ {
                    var v58 int32 = 1
                    _ = fi57
                    _ = v58
                }
            case -203:
                fmt.Println(v4)
            default:
                if (int64(v5) < 678) {
                    v1 = v1
                    v59 := GenS1{F0: (true || (false || false)), F1: "", F2: 4294967295}
                    for fi60 := 0; fi60 < 17; fi60++ {
                        fmt.Println(v4)
                        v1 -= v1
                        _ = fi60
                    }
                    _ = v59
                }
            }
        case -9223372036854775808:
            v61 := v6[v4]
            _ = v61
        default:
            v1 = (v1 >> ((226) & 63))
        }
        _ = fi56
    }
    var v62 int = -824
    v32 -= 0.0
    _ = len(v3)
    if (false && (!((!true) && (v1 <= v1)))) {
        var v63 uint32 = 4294967295
        var v64 bool = true
        v65 := GenS9{F0: genFn0(-420, ((v4 + v4) + "hello"), -1), F1: 100, F2: 18446744073709551615}
        _ = v63
        _ = v64
        _ = v65
    } else {
        v33 = v33
        switch v62 {
        case 300:
            switch ("" + ("foo bar" + (v4 + (v4 + (((v4 + ("\n\t" + "\n\t")) + v4) + "a"))))) {
            case "\n\t":
                v3 = append(v3, (v62 ^ v62))
            default:
                switch v4 {
                case "":
                    for fi66 := 0; fi66 < 7; fi66++ {
                        if true {
                            _ = len(v3)
                            v67 := GenS3{F0: -872}
                            _ = v67
                        }
                        var v68 int64 = 792
                        _ = fi66
                        _ = v68
                    }
                }
            }
        case 994:
            v69 := GenS9{F0: 102, F1: -99, F2: 114}
            _ = v69
        }
        switch v62 {
        case 127:
            if true {
                for fi70 := 0; fi70 < 2; fi70++ {
                    v3 = append(v3, (fi70 + fi70))
                    _ = fi70
                }
                var v71 int16 = v5
                _ = v71
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
