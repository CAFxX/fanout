package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8; F2 int64 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32, p1 int64, p2 int32) int32 {
    var v7 uint16 = 42
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return int32(v7)
}
func genFn1(p0 int16, p1 uint8) int16 {
    p1 += (^(p1))
    var v8 int16 = 0
    v9 := GenS1{F0: -1, F1: 18446744073709551615, F2: p0}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 63, 785, 42}
    var v4 string = "foo bar"
    v5 := []string{v4, v4, v4, (((v4 + v4) + (v4 + v4)) + v4)}
    var v6 int8 = -113
    if true {
        v10 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: -793}
        for fi11 := 0; fi11 < 3; fi11++ {
            v3 = append(v3, 429)
            var v12 uint16 = 0
            _ = fi11
            _ = v12
        }
        var v13 int = 916
        _ = v10
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v4 = ""
        fmt.Println(v4, v1)
        _ = fi14
    }
    var v15 int16 = genFn1(-32768, 0)
    switch -9223372036854775808 {
    case 32767:
        v1 += int64(v15)
    case -210:
        if true {
            var v16 uint16 = 65535
            v17 := GenS1{F0: 944, F1: 18446744073709551615, F2: v15}
            if false {
                var v18 string = ((v4 + "\n\t") + v4)
                v2[-285] = "foo bar"
                var v19 int = 32768
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 3; fi20++ {
                    if false {
                        switch (fi20 - int(v1)) {
                        case 420:
                            v21 := v2[(fi20 % ((fi20) | 1))]
                            _ = v21
                        case -253:
                            var v22 uint8 = 0
                            _ = v22
                        }
                    } else {
                        var v23 float64 = 3.14159
                        v16 = 32767
                        switch v4 {
                        case "":
                            var v24 int8 = v6
                            _ = v24
                        case "a":
                            for fi25 := 0; fi25 < 10; fi25++ {
                                fmt.Println(v6, fi20)
                                v4 = ("" + "\n\t")
                                _ = fi25
                            }
                        }
                        _ = v23
                    }
                    v26 := GenS2{F0: v4, F1: 255}
                    var v27 int32 = int32(v1)
                    _ = fi20
                    _ = v26
                    _ = v27
                }
                var v28 float64 = (3.14159 * -2.25)
                _ = v28
            }
            _ = v16
            _ = v17
        }
    default:
        v15 -= (v15 % ((genFn1(v15, 127)) | 1))
    }
    var v29 int16 = (v15 + (v15 ^ (v15 ^ v15)))
    v1 = (v1 + v1)
    var v30 int8 = 100
    switch (((v4 + ("foo bar" + v4)) + ("foo bar" + v4)) + "hello") {
    case "a":
        for fi31 := 0; fi31 < 20; fi31++ {
            switch fi31 {
            case 647:
                for fi32 := 0; fi32 < 19; fi32++ {
                    fmt.Println(v30)
                    v5 = append(v5, v4)
                    v15 = genFn1((^(v29)), 128)
                    _ = fi32
                }
            case 42:
                var v33 int = fi31
                _ = v33
            case -245:
                var v34 uint16 = 0
                _ = v34
            default:
                for fi35 := 0; fi35 < 3; fi35++ {
                    if true {
                        var v36 float64 = 1e10
                        delete(v2, (fi31 + (fi35 * (fi35 ^ (fi35 * fi31)))))
                        switch "a" {
                        case "\n\t":
                            var v37 uint = 4294967295
                            _ = v37
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                        case "":
                            if false {
                                var v38 uint8 = 128
                                v39 := GenS1{F0: 514, F1: 0, F2: -710}
                                var v40 uint32 = 642
                                v38 = (v38 ^ 55)
                                var v41 uint32 = (^(v40))
                                var v42 int32 = 949
                                v43 := GenS3{F0: v6, F1: (v38 << ((269) & 63)), F2: (v1 | (-(v1)))}
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            if ((((float32(0.0) + (-2.25 - 1.5)) + ((1e10 + -2.25) * 0.1)) * (3.14159 / 0.1)) != -2.25) {
                                v2[(fi35 + fi35)] = "foo bar"
                                if false {
                                    var v44 float64 = v36
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    if ((-(v15)) < v29) {
                                        v45 := GenS1{F0: 1, F1: 0, F2: (v15 - 32767)}
                                        if false {
                                            if (32767 != 1) {
                                                if (322 <= 2) {
                                                    v46 := GenS2{F0: v4, F1: 127}
                                                    _ = v46
                                                } else {
                                                    switch v4 {
                                                    case "foo bar":
                                                        switch 32767 {
                                                        case 156:
                                                            switch 571 {
                                                            case 0:
                                                                v1 += 0
                                                            case 42:
                                                                switch (fi31 ^ 365) {
                                                                case 32768:
                                                                    v30 += (^(v6))
                                                                case 4294967295:
                                                                    var v47 uint16 = 2
                                                                    _ = v47
                                                                }
                                                            case 4294967295:
                                                                v30 = -119
                                                            default:
                                                                switch (-(fi31)) {
                                                                case 128:
                                                                    v48 := GenS1{F0: 0, F1: 42, F2: genFn1((v29 + (v29 - (v29 % ((v29) | 1)))), 118)}
                                                                    _ = v48
                                                                default:
                                                                    switch v4 {
                                                                    case "hello":
                                                                        if (2 > 0) {
                                                                            if false {
                                                                                if (true || false) {
                                                                                    v1 += v1
                                                                                    var v49 uint16 = 2
                                                                                    _ = v49
                                                                                }
                                                                            } else {
                                                                                delete(v2, 0)
                                                                                v6 -= v30
                                                                                switch (fi35 & 128) {
                                                                                case 9223372036854775807:
                                                                                    if (v15 > (v29 << ((18446744073709551615) & 63))) {
                                                                                        fmt.Println(v4)
                                                                                        fmt.Println(v30, v6, v15)
                                                                                    }
                                                                                }
                                                                            }
                                                                            v50 := GenS3{F0: v30, F1: 9, F2: v1}
                                                                            var v51 int8 = -81
                                                                            _ = v50
                                                                            _ = v51
                                                                        } else {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        fmt.Println(v29, v15, v4)
                                                    default:
                                                        v1 = 9223372036854775807
                                                    }
                                                    var v52 uint64 = 1
                                                    _ = v52
                                                }
                                            } else {
                                                v53 := GenS3{F0: -1, F1: 91, F2: (v1 ^ 9223372036854775807)}
                                                v54 := GenS2{F0: "hello", F1: 0}
                                                var v55 int = (fi35 >> ((549) & 63))
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            fmt.Println(v6)
                                            fmt.Println(fi31, v15)
                                        } else {
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v56 uint64 = 1
                                                _ = v56
                                            } else {
                                                var v57 int32 = genFn0(1, v1, -150)
                                                _ = v57
                                            }
                                            var v58 uint8 = 2
                                            _ = v58
                                        }
                                        v1 = (^(v1))
                                        _ = v45
                                    } else {
                                        v36 = (v36 - ((v36 / float64(((((((float32(float64(-2.25)) + (1e10 * (((1.5 / (float32((float64((3.14159 + ((1e10 * 0.1) - (((float32(v36) - 1e10) * -2.25) - -2.25)))) * v44)) * 1.5)) + (1e10 - 1.5)) * (0.0 / 0.1)))) / float32(-2.25)) / 3.14159) - -2.25) * -2.25) - 0.1))) - (v36 + 0.1)))
                                        var v59 int = fi35
                                        if false {
                                            fmt.Println(v6, v6)
                                        }
                                        _ = v59
                                    }
                                    _ = v44
                                }
                                var v60 uint = 42
                                _ = v60
                            }
                        }
                        _ = v36
                    }
                    var v61 int8 = v30
                    _ = fi35
                    _ = v61
                }
            }
            v62 := GenS3{F0: 127, F1: 2, F2: 42}
            var v63 string = v4
            _ = fi31
            _ = v62
            _ = v63
        }
    }
    v30 += 57
    v30 -= v30
    if true {
        v15 += v29
        switch v4 {
        case "a":
            fmt.Println(v1, v4, v6)
        case "\n\t":
            for fi64 := 0; fi64 < 9; fi64++ {
                var v65 string = (("\n\t" + ((v4 + v4) + "foo bar")) + (("hello" + ("" + "foo bar")) + ("\n\t" + v4)))
                fmt.Println(v15)
                var v66 int32 = -1
                _ = fi64
                _ = v65
                _ = v66
            }
        case "":
            v6 = v30
        default:
            var v67 int32 = 42
            _ = v67
        }
    } else {
        switch 1 {
        case -1:
            v68 := GenS3{F0: (v30 << ((1) & 63)), F1: 128, F2: v1}
            _ = v68
        default:
            var v69 int16 = -1
            _ = v69
        }
        var v70 float64 = (float64((0.0 / 1e10)) - (3.14159 / (1e10 - 0.1)))
        fmt.Println(v15, v1)
        _ = v70
    }
    v4 = (v4 + "hello")
    v71 := GenS4{F0: false, F1: 222}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
