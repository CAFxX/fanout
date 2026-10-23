package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int8 }
func genFn0() int {
    var v4 float32 = -2.25
    fmt.Println(v4)
    switch 127 {
    case 128:
        for fi5 := 0; fi5 < 1; fi5++ {
            for fi6 := 0; fi6 < 14; fi6++ {
                v7 := []int{(fi6 % ((fi6) | 1)), 32768}
                switch 622 {
                case 2147483647:
                    v7 = append(v7, fi6)
                case 32768:
                    fmt.Println(fi5, fi6)
                }
                _ = fi6
                _ = v7
            }
            var v8 int64 = 1
            v9 := []int{1}
            _ = fi5
            _ = v8
            _ = v9
        }
    case 4294967295:
        var v10 int32 = 174
        _ = v10
    case 9223372036854775807:
        v4 = (((v4 + 1e10) / (((0.0 / 3.14159) + ((0.0 - (0.1 * (v4 / v4))) + 1.5)) + v4)) / float32(3.14159))
    }
    _ = v4
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -9223372036854775808
    v3 := []string{"\n\t", ("foo bar" + "a"), "hello"}
    var v11 bool = false
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 1; fi12++ {
        if v11 {
            var v13 float64 = 3.14159
            _ = v13
        } else {
            fmt.Println(fi12, v11, v2)
            v14 := GenS1{F0: 4294967295}
            var v15 uint8 = 119
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v11 = v11
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: 2147483648}
        v18 := GenS1{F0: 2147483648}
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        fmt.Println(v2)
        var v19 bool = true
        v11 = false
        _ = v19
    } else {
        if false {
            switch 4294967295 {
            case 128:
                if v11 {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 int = genFn0()
                        _ = fi20
                        _ = v21
                    }
                    if (!(v2 <= v2)) {
                        var v22 int8 = 0
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 uint8 = 128
                            for fi25 := 0; fi25 < 2; fi25++ {
                                if (v11 || false) {
                                    if (v2 < (v2 % (((v2 / ((v2) | 1))) | 1))) {
                                        if v11 {
                                            var v26 uint16 = 176
                                            v27 := GenS2{F0: v22}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 uint64 = 2
                                            fmt.Println(v28, v11, v24)
                                            if true {
                                                var v29 string = ("" + "\n\t")
                                                v30 := v1["a"]
                                                switch fi25 {
                                                case 42:
                                                    v31 := GenS1{F0: (v30 * (v30 * v30))}
                                                    _ = v31
                                                case -9223372036854775808:
                                                    var v32 int64 = 42
                                                    _ = v32
                                                case 0:
                                                    var v33 int32 = 0
                                                    _ = v33
                                                default:
                                                    v11 = v11
                                                }
                                                v34 := GenS3{F0: 858, F1: 223}
                                                var v35 uint32 = v30
                                                _ = v29
                                                _ = v30
                                                _ = v34
                                                _ = v35
                                            }
                                            _ = v28
                                        }
                                        switch fi25 {
                                        case 394:
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v36 float32 = 0.0
                                                _ = v36
                                            default:
                                                if v11 {
                                                    _ = len(v3)
                                                    v2 = v2
                                                    fmt.Println(fi25, v22)
                                                }
                                            }
                                            if v11 {
                                                v24 += 255
                                            } else {
                                                _ = len(v3)
                                            }
                                            fmt.Println(v11, v22, fi25)
                                        }
                                        var v37 uint16 = 42
                                        _ = v37
                                    }
                                    v38 := GenS2{F0: (v22 >> ((207) & 63))}
                                    _ = v38
                                } else {
                                    var v39 int16 = 0
                                    v39 += int16(v22)
                                    _ = len(v3)
                                    _ = v39
                                }
                                switch "\n\t" {
                                case "hello":
                                    v2 -= v2
                                case "foo bar":
                                    v11 = v11
                                case "a":
                                    if (v2 < -1) {
                                        switch ((("foo bar" + (("" + ((("hello" + ("\n\t" + "hello")) + "a") + "")) + "\n\t")) + "") + "\n\t") {
                                        case "foo bar":
                                            var v40 float64 = 0.1
                                            _ = v40
                                        case "hello":
                                            var v41 uint64 = 1
                                            _ = v41
                                        case "":
                                            fmt.Println(v24, v11)
                                        }
                                    } else {
                                        var v42 uint16 = 255
                                        _ = len(v1)
                                        _ = v42
                                    }
                                default:
                                    v22 = (v22 / ((v22) | 1))
                                }
                                _ = fi25
                            }
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    } else {
                        var v43 uint64 = 865
                        _ = v43
                    }
                } else {
                    switch (("a" + "\n\t") + "") {
                    case "":
                        v44 := GenS1{F0: 740}
                        _ = v44
                    default:
                        v45 := GenS1{F0: 0}
                        _ = v45
                    }
                }
            case 634:
                if false {
                    for fi46 := 0; fi46 < 18; fi46++ {
                        var v47 int8 = -128
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        switch -1 {
                        case 42:
                            var v48 uint16 = 489
                            _ = v48
                        case 1:
                            v47 += v47
                        }
                        _ = fi46
                        _ = v47
                    }
                    switch "a" {
                    case "\n\t":
                        v2 = (v2 & (v2 ^ v2))
                    case "hello":
                        fmt.Println(v2, v11)
                    case "":
                        if false {
                            var v49 int32 = -290
                            _ = v49
                        }
                    default:
                        fmt.Println(v2)
                    }
                } else {
                    _ = len(v1)
                }
            }
            v2 += (v2 + (-(v2)))
        } else {
            v11 = ((float64((0.0 + ((float32((0.1 / 0.1)) + (1e10 - (3.14159 * (3.14159 * -2.25)))) - 0.0))) * 1.5) < (((1e10 * (-2.25 / (-2.25 * 1e10))) + 0.0) - 3.14159))
            for fi50 := 0; fi50 < 20; fi50++ {
                if false {
                    switch "a" {
                    case "foo bar":
                        for fi51 := 0; fi51 < 12; fi51++ {
                            var v52 uint32 = 0
                            if (42 == 42) {
                                if (v52 != (v52 * (v52 * 0))) {
                                    _ = len(v3)
                                    if true {
                                        var v53 uint16 = 1
                                        _ = v53
                                    }
                                } else {
                                    v11 = (-216 > 1000)
                                    v2 = (v2 ^ (v2 - (-(v2))))
                                }
                                v11 = (true || (65535 > 888))
                            }
                            v54 := GenS4{F0: 409, F1: 9223372036854775807, F2: ("\n\t" + "hello")}
                            _ = fi51
                            _ = v52
                            _ = v54
                        }
                    case "":
                        v2 = 550
                    }
                    v55 := v1["\n\t"]
                    for fi56 := 0; fi56 < 17; fi56++ {
                        v57 := GenS5{F0: v2}
                        var v58 int8 = 0
                        v55 += v55
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    _ = v55
                } else {
                    delete(v1, (((("hello" + ("foo bar" + "a")) + ("" + ("\n\t" + "foo bar"))) + "hello") + "hello"))
                }
                v2 = (v2 % (((-(v2))) | 1))
                _ = fi50
            }
            _ = len(v1)
        }
    }
    _ = len(v1)
    v59 := GenS2{F0: -86}
    if (687 >= v2) {
        for fi60 := 0; fi60 < 9; fi60++ {
            var v61 string = (("a" + "\n\t") + "\n\t")
            _ = fi60
            _ = v61
        }
        fmt.Println(v2, v11)
    } else {
        fmt.Println(v2, v11)
    }
    switch -159 {
    case -1:
        fmt.Println(v2, v2)
    }
    switch 9223372036854775807 {
    case 32767:
        fmt.Println(v2, v2)
    case 0:
        v62 := GenS6{F0: 4294967295, F1: 543}
        _ = v62
    }
    if (((!v11) && (true && (1e10 == 0.1))) || (!(816 == 305))) {
        v63 := GenS2{F0: 1}
        _ = v63
    }
    switch ("hello" + "foo bar") {
    case "hello":
        var v64 bool = v11
        _ = v64
    }
    for fi65 := 0; fi65 < 8; fi65++ {
        for fi66 := 0; fi66 < 4; fi66++ {
            v2 -= v2
            _ = fi66
        }
        v67 := GenS7{F0: 2, F1: 42}
        var v68 uint16 = 209
        _ = fi65
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 1; fi69++ {
        var v70 uint32 = 2147483648
        var v71 int8 = 23
        var v72 int = (fi69 - (-(fi69)))
        _ = fi69
        _ = v70
        _ = v71
        _ = v72
    }
    _ = len(v1)
    v11 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v59
}
