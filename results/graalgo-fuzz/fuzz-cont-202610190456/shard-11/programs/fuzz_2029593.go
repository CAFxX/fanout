package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int8; F2 int }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 uint8, p1 uint32) int8 {
    fmt.Println(p1)
    for fi4 := 0; fi4 < 11; fi4++ {
        if (1 == 0) {
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 int8 = int8(p0)
                fmt.Println(v6, p0, p0)
                p1 = 151
                _ = fi5
                _ = v6
            }
        }
        p1 = 2
        var v7 string = ("" + ((("" + "foo bar") + "") + ""))
        _ = fi4
        _ = v7
    }
    switch -376 {
    case 515:
        for fi8 := 0; fi8 < 16; fi8++ {
            p1 -= (p1 & (p1 - (^(p1))))
            _ = fi8
        }
    case -1:
        v9 := GenS1{F0: 1, F1: "\n\t", F2: 1}
        _ = v9
    }
    var v10 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    return -58
}
func genFn1(p0 int8, p1 int8) string {
    p0 = -128
    p1 += p1
    var v11 int32 = 0
    _ = p0
    _ = p1
    _ = v11
    return ("\n\t" + "hello")
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 0
    v3 := []uint32{}
    v2 = 2147483648
    v2 = v2
    var v12 uint8 = 236
    v13 := GenS2{F0: (((v12 == 255) && false) && (0 > -57)), F1: v12, F2: 1}
    var v14 string = "foo bar"
    _ = len(v3)
    var v15 int64 = -9223372036854775808
    if true {
        var v16 string = v14
        switch v16 {
        case "hello":
            var v17 int16 = -1
            _ = v17
        case "a":
            v18 := v1[v16]
            _ = v18
        case "":
            v19 := v1[("" + "a")]
            _ = v19
        }
        _ = v16
    } else {
        var v20 uint = 0
        _ = v20
    }
    v15 += 42
    var v21 string = v14
    for fi22 := 0; fi22 < 17; fi22++ {
        v2 += v2
        v23 := GenS3{F0: genFn0(v12, (v2 | v2)), F1: fi22}
        v2 = (^(v2))
        _ = fi22
        _ = v23
    }
    var v24 uint16 = uint16(v2)
    for fi25 := 0; fi25 < 14; fi25++ {
        if true {
            fmt.Println(v21, v24)
            switch (("\n\t" + v14) + (v14 + ("\n\t" + v21))) {
            case "\n\t":
                v26 := GenS4{F0: -396}
                _ = v26
            case "":
                v2 = 2147483648
            }
            switch "" {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = (v2 / ((2147483647) | 1))
                }
            case "\n\t":
                for fi27 := 0; fi27 < 14; fi27++ {
                    if (!false) {
                        var v28 float32 = (0.0 - (3.14159 + 3.14159))
                        v14 = "\n\t"
                        v21 = "\n\t"
                        _ = v28
                    } else {
                        if false {
                            if false {
                                v29 := GenS3{F0: 1, F1: fi25}
                                if (false || (-47 >= int8(v12))) {
                                    if true {
                                        v2 = (v2 | 4294967295)
                                    }
                                    if true {
                                        v30 := GenS5{F0: -89, F1: 1, F2: v12}
                                        if (!true) {
                                            if len(v3) > 0 {
                                                v3[0] = uint32(v24)
                                            }
                                        } else {
                                            var v31 int16 = 665
                                            v24 = (v24 % ((v24) | 1))
                                            _ = v31
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v30
                                    }
                                    _ = len(v1)
                                }
                                fmt.Println(fi27, v21)
                                _ = v29
                            }
                            v14 = (v21 + "\n\t")
                            v1[""] = true
                        } else {
                            if false {
                                if false {
                                    if true {
                                        var v32 int8 = -1
                                        _ = v32
                                    } else {
                                        fmt.Println(v2, v14)
                                        if (v21 <= v14) {
                                            fmt.Println(fi25, fi27, fi25)
                                            v33 := GenS1{F0: 42, F1: "a", F2: 1}
                                            v34 := GenS4{F0: 0}
                                            _ = v33
                                            _ = v34
                                        } else {
                                            if false {
                                                var v35 float32 = (((1.5 / 0.1) + (0.1 + (0.1 - 0.0))) - ((float32(((3.14159 / ((1e10 * ((0.0 / 1e10) + 1e10)) / 1e10)) * 1e10)) * 1e10) + -2.25))
                                                var v36 bool = true
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                        var v37 int = 4294967295
                                        fmt.Println(v24, fi27, v2)
                                        var v38 int64 = -9223372036854775808
                                        _ = v37
                                        _ = v38
                                    }
                                    fmt.Println(v14, v12)
                                    if true {
                                        if (v2 >= 275) {
                                            switch v21 {
                                            case "hello":
                                                var v39 int64 = (v15 % ((v15) | 1))
                                                _ = v39
                                            case "\n\t":
                                                fmt.Println(v2, v14)
                                                fmt.Println(fi25, v2)
                                            case "a":
                                                var v40 int32 = -1
                                                _ = v40
                                            }
                                            var v41 int = 127
                                            if ((928 < 2147483647) || ((v2 > 4294967295) || true)) {
                                                var v42 float32 = (0.0 - float32((-2.25 - 0.0)))
                                                _ = v42
                                            }
                                            _ = v41
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v12)
                                switch (((v21 + (("foo bar" + v21) + (v14 + ("hello" + (v21 + v21))))) + v14) + "") {
                                case "a":
                                    fmt.Println(v15)
                                }
                                if (true && (!(1.5 == (1.5 / (3.14159 / 3.14159))))) {
                                    var v43 int16 = 160
                                    _ = v43
                                } else {
                                    v44 := GenS6{F0: -811, F1: 8, F2: (fi25 + fi25)}
                                    v24 = 0
                                    _ = v44
                                }
                            }
                            v15 += (v15 & (v15 * v15))
                        }
                    }
                    fmt.Println(fi27, v21)
                    if (6 < genFn0(127, 537)) {
                        switch v14 {
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v14, v24, v12)
                                var v45 uint16 = (v24 | v24)
                                v24 -= 42
                                _ = v45
                            }
                        }
                    } else {
                        if (!false) {
                            switch (fi27 % ((127) | 1)) {
                            case -977:
                                if false {
                                    fmt.Println(v21, v24)
                                    v14 = "\n\t"
                                    switch fi27 {
                                    case 976:
                                        v3 = append(v3, (^(v2)))
                                    }
                                } else {
                                    var v46 uint8 = (v12 - (v12 % ((v12) | 1)))
                                    _ = v46
                                }
                                var v47 float32 = 1e10
                                var v48 string = (v21 + ("foo bar" + "foo bar"))
                                _ = v47
                                _ = v48
                            case 9223372036854775807:
                                v49 := GenS7{F0: v2, F1: 255}
                                _ = v49
                            }
                            var v50 uint16 = 42
                            switch (("foo bar" + "a") + v21) {
                            case "\n\t":
                                var v51 int16 = 0
                                _ = v51
                            case "":
                                if false {
                                    v50 = 2
                                    if false {
                                        v21 = v21
                                    }
                                    var v52 int8 = -106
                                    var v53 int8 = v52
                                    v14 = v14
                                    _ = v52
                                    _ = v53
                                } else {
                                    if true {
                                        v21 = (("a" + ("\n\t" + ((v14 + v21) + "hello"))) + ("a" + v14))
                                        if len(v3) > 0 {
                                            v3[0] = 65535
                                        }
                                    } else {
                                        var v54 string = ((v14 + v21) + ((v21 + "") + v21))
                                        _ = v54
                                    }
                                }
                                var v55 int64 = 1
                                _ = v55
                            case "hello":
                                var v56 string = v21
                                _ = v56
                            }
                            _ = v50
                        }
                        var v57 bool = (((float64(0.1) + 0.1) + (1.5 + -2.25)) > (1e10 * 0.0))
                        if ((205 > 32767) || (0 == 2147483647)) {
                            v57 = (2 == 7)
                        }
                        if v57 {
                            v58 := GenS8{F0: -1}
                            _ = v58
                        }
                        if v57 {
                            var v59 float64 = (((0.0 / (((1.5 + 0.1) * 1e10) / 1e10)) + -2.25) + float64((0.0 * 0.0)))
                            var v60 string = v21
                            _ = v59
                            _ = v60
                        } else {
                            var v61 int64 = (-(v15))
                            _ = v61
                        }
                        _ = v57
                    }
                    _ = fi27
                }
            default:
                var v62 bool = (true || (fi25 > fi25))
                _ = v62
            }
        }
        _ = fi25
    }
    var v63 uint = 18446744073709551615
    var v64 uint32 = (^(v2))
    var v65 uint32 = (v2 | (^(v2)))
    for fi66 := 0; fi66 < 13; fi66++ {
        var v67 int64 = (v15 << ((v63) & 63))
        fmt.Println(v14)
        v24 -= 361
        _ = fi66
        _ = v67
    }
    v68 := GenS9{F0: 42, F1: 2147483647, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v68
}
