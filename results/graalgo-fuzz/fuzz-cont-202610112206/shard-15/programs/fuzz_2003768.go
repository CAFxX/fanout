package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint; F2 int32 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
func genFn0(p0 int32, p1 uint16) int16 {
    v6 := GenS1{F0: 1, F1: 42}
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            var v8 uint8 = 248
            _ = v8
        } else {
            if false {
                fmt.Println(fi7, p1)
                p1 -= (p1 & (p1 - (p1 | 42)))
            }
            p1 = 837
        }
        v9 := GenS2{F0: 2, F1: -2147483648, F2: 42}
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 333
}
func genFn1(p0 int64, p1 uint32) uint32 {
    switch 31 {
    case -377:
        var v10 float32 = 0.1
        _ = v10
    case 9223372036854775807:
        var v11 uint16 = 42
        _ = v11
    case -867:
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 bool = true
            _ = fi12
            _ = v13
        }
    }
    if false {
        fmt.Println(p0)
        p1 = 1
        var v14 int8 = 61
        _ = v14
    }
    var v15 float32 = 1.5
    v16 := GenS3{F0: 808, F1: 32767}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = float32((((1.5 - ((-2.25 * (-2.25 - float64(1e10))) / 1.5)) + (float64(1e10) - 1.5)) / 3.14159))
    v3 := make(map[string]uint32)
    var v4 float32 = v2
    var v5 uint8 = 13
    var v17 float64 = (1e10 / 1e10)
    _ = len(v3)
    var v18 int32 = 42
    v19 := GenS4{F0: v18}
    if ((!(true || (915 >= 9223372036854775807))) || (("foo bar" + (("\n\t" + "\n\t") + "foo bar")) > ("\n\t" + "\n\t"))) {
        if (19 != 271) {
            v2 += v4
            fmt.Println(v18, v17)
            v20 := GenS5{F0: ""}
            _ = v20
        }
        v21 := GenS6{F0: 939, F1: v5, F2: 42}
        v22 := []int{9223372036854775807, -251, 127, 0}
        _ = v21
        _ = v22
    } else {
        v23 := GenS6{F0: uint16(v5), F1: v5, F2: -165}
        var v24 uint32 = 65535
        _ = v23
        _ = v24
    }
    v4 += ((3.14159 * float32(0.1)) / (v2 + v4))
    v25 := []int{4294967295, -2147483648, -340}
    for fi26 := 0; fi26 < 5; fi26++ {
        fmt.Println(v17, v18, fi26)
        v5 -= v5
        v27 := GenS3{F0: 2147483647, F1: 0}
        _ = fi26
        _ = v27
    }
    v17 = v17
    var v28 uint16 = 255
    v25 = append(v25, -502)
    switch "\n\t" {
    case "":
        if (!(!true)) {
            fmt.Println(v28, v28, v2)
            if (65535 >= 42) {
                fmt.Println(v2, v18)
                if false {
                    switch 188 {
                    case 127:
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    default:
                        var v29 int16 = 1000
                        _ = v29
                    }
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        var v30 float64 = v17
                        _ = v30
                    case "a":
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    case "\n\t":
                        for fi31 := 0; fi31 < 3; fi31++ {
                            v4 -= v4
                            var v32 bool = true
                            if false {
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    v28 = v28
                                    var v34 int16 = 1
                                    _ = fi33
                                    _ = v34
                                }
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                            }
                            _ = fi31
                            _ = v32
                        }
                    default:
                        for fi35 := 0; fi35 < 19; fi35++ {
                            if true {
                                if (468 <= fi35) {
                                    v36 := v1[fi35]
                                    switch (("foo bar" + "foo bar") + ("" + ("\n\t" + ("a" + "foo bar")))) {
                                    case "":
                                        var v37 bool = (!false)
                                        _ = v37
                                    }
                                    var v38 int8 = 127
                                    _ = v36
                                    _ = v38
                                }
                                if len(v25) > 0 {
                                    v25[0] = -866
                                }
                            }
                            if (((v17 / v17) <= 3.14159) && (48 > 1)) {
                                switch 127 {
                                case 1:
                                    switch (fi35 + fi35) {
                                    case -579:
                                        var v39 uint16 = v28
                                        _ = v39
                                    default:
                                        fmt.Println(fi35)
                                    }
                                case 128:
                                    v40 := GenS7{F0: 4294967295}
                                    _ = v40
                                }
                                if false {
                                    for fi41 := 0; fi41 < 7; fi41++ {
                                        var v42 uint32 = 74
                                        var v43 float32 = v4
                                        if true {
                                            fmt.Println(v18)
                                            fmt.Println(fi41)
                                            _ = len(v1)
                                            if (-9223372036854775808 < 995) {
                                                v4 += v4
                                            } else {
                                                v28 += (v28 / (((v28 / ((uint16(v5)) | 1))) | 1))
                                                var v44 uint32 = 2147483647
                                                _ = v44
                                            }
                                            if (uint32(v28) < genFn1(1, (v42 / ((uint32(v5)) | 1)))) {
                                                var v45 int64 = -1
                                                _ = v45
                                            } else {
                                                switch fi41 {
                                                case 252:
                                                    fmt.Println(fi41, v43)
                                                case 516:
                                                    v17 = 1e10
                                                }
                                                v1[0] = ("" > "foo bar")
                                            }
                                        }
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                    v3[("" + "\n\t")] = 1
                                }
                                v46 := GenS7{F0: 2}
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 7; fi47++ {
                                if true {
                                    var v48 uint16 = (v28 >> ((18446744073709551615) & 63))
                                    var v49 int = (fi47 ^ 1)
                                    _ = v48
                                    _ = v49
                                }
                                switch fi35 {
                                case -187:
                                    v50 := GenS8{F0: ("" + "hello"), F1: 65535, F2: (v18 >> ((425) & 63))}
                                    _ = v50
                                case 608:
                                    v51 := GenS9{F0: false, F1: genFn1(62, 2)}
                                    var v52 uint = 0
                                    v53 := GenS10{F0: false}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                default:
                                    switch (fi35 % (((fi35 + fi35)) | 1)) {
                                    case 9223372036854775807:
                                        fmt.Println(fi47, v4, fi47)
                                        var v54 uint32 = 733
                                        var v55 uint32 = (v54 >> ((0) & 63))
                                        _ = v54
                                        _ = v55
                                    default:
                                        v17 -= 0.1
                                    }
                                }
                                _ = len(v3)
                                _ = fi47
                            }
                            _ = fi35
                        }
                    }
                }
            }
            v18 -= v18
        } else {
            v5 += (v5 ^ v5)
            var v56 int8 = 127
            if len(v25) > 0 {
                v25[0] = 718
            }
            _ = v56
        }
    }
    var v57 float64 = v17
    if ((false && (9223372036854775807 > 688)) && ((4294967295 < 2147483647) || (-278 < -361))) {
        var v58 uint = 253
        delete(v3, "foo bar")
        fmt.Println(v28, v2, v5)
        _ = v58
    } else {
        var v59 int16 = genFn0(v18, v28)
        if (3.14159 <= (v4 + float32(v57))) {
            for fi60 := 0; fi60 < 1; fi60++ {
                var v61 int8 = 42
                if false {
                    for fi62 := 0; fi62 < 19; fi62++ {
                        fmt.Println(fi60)
                        v63 := GenS11{F0: "a"}
                        fmt.Println(v17)
                        _ = fi62
                        _ = v63
                    }
                    var v64 int = (fi60 << ((538) & 63))
                    var v65 int16 = (v59 / (((v59 | (v59 << ((42) & 63)))) | 1))
                    _ = v64
                    _ = v65
                }
                _ = fi60
                _ = v61
            }
        }
        var v66 int64 = 1
        _ = v59
        _ = v66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
