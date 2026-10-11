package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 string }
func genFn0(p0 bool) uint32 {
    var v5 int = 127
    switch v5 {
    case 0:
        switch (("foo bar" + ("hello" + (((((("foo bar" + "hello") + "") + "hello") + "hello") + "") + ((("hello" + "\n\t") + "foo bar") + ("hello" + "foo bar"))))) + "\n\t") {
        case "a":
            v5 -= v5
        }
    case -2147483648:
        var v6 float64 = (0.1 - 1e10)
        _ = v6
    }
    v5 -= v5
    _ = p0
    _ = v5
    return 1
}
func genFn1() int64 {
    var v7 float32 = 0.0
    v7 += 1.5
    _ = v7
    return -519
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 18446744073709551615
    var v3 bool = (0 == 127)
    v4 := make(map[uint32]uint32)
    v8 := GenS1{F0: 26, F1: -1, F2: -793}
    v9 := []int{42, 9223372036854775807}
    v4[65535] = 2
    v2 = 174
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 string = "a"
        _ = fi10
        _ = v11
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v12 uint8 = (v1 * 221)
    for fi13 := 0; fi13 < 5; fi13++ {
        if false {
            for fi14 := 0; fi14 < 12; fi14++ {
                v3 = (int16(v2) == 51)
                if (1 == v2) {
                    var v15 bool = true
                    _ = v15
                } else {
                    switch (((("hello" + (("a" + "foo bar") + "a")) + ("a" + "\n\t")) + "hello") + (("" + ("a" + "")) + "a")) {
                    case "\n\t":
                        if ((v2 <= v2) || v3) {
                            fmt.Println(v12)
                            var v16 uint32 = 1
                            var v17 string = (("foo bar" + "hello") + ("\n\t" + "\n\t"))
                            v18 := GenS1{F0: v12, F1: -2147483648, F2: -646}
                            v12 = (^(v1))
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int32 = 2147483647
                            _ = v19
                        }
                    case "a":
                        var v20 uint = 828
                        _ = v20
                    }
                    v2 = uint64(v12)
                }
                _ = fi14
            }
        } else {
            v1 += v12
            for fi21 := 0; fi21 < 9; fi21++ {
                switch fi13 {
                case 355:
                    if v3 {
                        if (((-55 <= int(v2)) || v3) && v3) {
                            var v22 int8 = 0
                            v12 += (^(v1))
                            _ = v22
                        } else {
                            v12 = v1
                        }
                        fmt.Println(fi21, fi13, v3)
                    }
                default:
                    switch "foo bar" {
                    case "\n\t":
                        v2 += 1
                    case "hello":
                        v2 -= v2
                    case "":
                        if (((fi13 << ((255) & 63)) != fi13) && (!true)) {
                            _ = len(v9)
                            v23 := v4[1]
                            var v24 string = "a"
                            _ = v23
                            _ = v24
                        }
                    default:
                        switch (fi13 >> ((18446744073709551615) & 63)) {
                        case 230:
                            v25 := GenS1{F0: uint8(v2), F1: -2147483648, F2: 305}
                            _ = v25
                        case 0:
                            if (!v3) {
                                var v26 float32 = (1.5 * 0.1)
                                if (!("hello" > ("" + ("" + ("hello" + (("" + "a") + "\n\t")))))) {
                                    v2 = 2
                                    var v27 int64 = 0
                                    _ = v27
                                }
                                if v3 {
                                    v28 := GenS1{F0: v1, F1: 14, F2: -1}
                                    _ = len(v4)
                                    var v29 uint = 1
                                    _ = v28
                                    _ = v29
                                }
                                _ = v26
                            } else {
                                var v30 uint = 18446744073709551615
                                v30 += (v30 >> ((v30) & 63))
                                _ = v30
                            }
                        }
                        if (322 < 640) {
                            var v31 int32 = -1
                            if v3 {
                                v2 -= v2
                                if len(v9) > 0 {
                                    v9[0] = fi13
                                }
                                delete(v4, genFn0(((v3 || (false && (!("a" >= ((("foo bar" + (("foo bar" + "a") + ("\n\t" + "a"))) + "") + "\n\t"))))) || v3)))
                            }
                            delete(v4, 2147483647)
                            v31 -= (v31 | (v31 | (v31 << ((990) & 63))))
                            _ = v31
                        } else {
                            var v32 uint64 = (v2 - v2)
                            v32 += 18446744073709551615
                            fmt.Println(v3)
                            if v3 {
                                if (v3 || v3) {
                                    if true {
                                        var v33 uint32 = 0
                                        fmt.Println(v32, v32)
                                        _ = v33
                                    } else {
                                        if (127 >= 36) {
                                            var v34 uint8 = (v1 * v12)
                                            v35 := GenS1{F0: v34, F1: 774, F2: -1}
                                            var v36 int16 = 32767
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    if v3 {
                                        var v37 int64 = genFn1()
                                        var v38 uint8 = v12
                                        v39 := GenS1{F0: (^(v38)), F1: -2147483648, F2: -1}
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    } else {
                                        v9 = append(v9, 4294967295)
                                    }
                                    var v40 int8 = 0
                                    _ = v40
                                }
                                fmt.Println(fi13, v3, v2)
                            }
                            v3 = ((!v3) && (!true))
                            var v41 uint = 65535
                            var v42 uint32 = 1
                            if v3 {
                                fmt.Println(fi21, v42, v32)
                            } else {
                                _ = len(v9)
                            }
                            _ = v32
                            _ = v41
                            _ = v42
                        }
                    }
                }
                var v43 uint16 = 0
                _ = fi21
                _ = v43
            }
            v44 := GenS2{F0: fi13}
            _ = v44
        }
        switch ("hello" + ("\n\t" + "a")) {
        case "hello":
            if true {
                v1 = v1
                var v45 int8 = -128
                if v3 {
                    if true {
                        if len(v9) > 0 {
                            v9[0] = fi13
                        }
                        switch fi13 {
                        case 9223372036854775807:
                            if len(v9) > 0 {
                                v9[0] = fi13
                            }
                        case 128:
                            fmt.Println(v2, v12)
                        case 4294967295:
                            if false {
                                switch fi13 {
                                case 123:
                                    v12 -= 42
                                }
                            }
                        }
                        var v46 uint16 = 65535
                        _ = v46
                    }
                    fmt.Println(v3, v3, v12)
                    var v47 uint16 = 2
                    _ = v47
                } else {
                    var v48 string = "hello"
                    var v49 int8 = (v45 << ((254) & 63))
                    v48 = ""
                    _ = v48
                    _ = v49
                }
                _ = v45
            } else {
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "":
                        v12 -= 2
                    case "foo bar":
                        for fi50 := 0; fi50 < 15; fi50++ {
                            var v51 int32 = 42
                            _ = fi50
                            _ = v51
                        }
                    default:
                        switch "a" {
                        case "\n\t":
                            var v52 uint8 = v12
                            _ = v52
                        case "":
                            v53 := GenS2{F0: (^(fi13))}
                            _ = v53
                        case "a":
                            if (1e10 != (float32(float64(0.1)) + (0.1 - 1e10))) {
                                for fi54 := 0; fi54 < 9; fi54++ {
                                    v9 = append(v9, 127)
                                    _ = fi54
                                }
                            } else {
                                v2 = (^(v2))
                                v55 := GenS2{F0: (fi13 >> ((597) & 63))}
                                if ((v12 + 2) != (v12 % ((v1) | 1))) {
                                    var v56 uint8 = v1
                                    if false {
                                        for fi57 := 0; fi57 < 11; fi57++ {
                                            if (42 <= v2) {
                                                v56 += (^(v56))
                                                var v58 int = (fi13 ^ (-(fi57)))
                                                v59 := GenS2{F0: fi57}
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v60 := GenS3{F0: 821, F1: 127}
                                                _ = v60
                                            }
                                            _ = fi57
                                        }
                                    } else {
                                        v4[uint32(v1)] = 65535
                                        v61 := GenS1{F0: (^(v1)), F1: -307, F2: 42}
                                        var v62 uint8 = (^(v56))
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = v56
                                }
                                _ = v55
                            }
                        }
                    }
                case "":
                    _ = len(v9)
                case "foo bar":
                    fmt.Println(v1, v1, fi13)
                default:
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            for fi63 := 0; fi63 < 8; fi63++ {
                                if false {
                                    if v3 {
                                        if (-2147483648 <= 1) {
                                            delete(v4, 42)
                                        } else {
                                            var v64 int8 = 1
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            _ = v64
                                        }
                                    }
                                } else {
                                    fmt.Println(fi13)
                                }
                                switch (fi13 << ((8) & 63)) {
                                case 1:
                                    if v3 {
                                        fmt.Println(v2)
                                    }
                                case -818:
                                    _ = len(v4)
                                case -9223372036854775808:
                                    var v65 uint32 = 4294967295
                                    _ = v65
                                default:
                                    v66 := GenS1{F0: (v1 / (((^(v12))) | 1)), F1: 2147483647, F2: -2147483648}
                                    _ = v66
                                }
                                _ = fi63
                            }
                        default:
                            if true {
                                v67 := GenS3{F0: fi13, F1: v12}
                                switch "" {
                                case "hello":
                                    fmt.Println(v12, v2)
                                case "foo bar":
                                    v68 := GenS4{F0: 973}
                                    _ = v68
                                default:
                                    var v69 uint32 = 161
                                    _ = v69
                                }
                                _ = v67
                            }
                        }
                        v70 := GenS5{F0: -975, F1: 255, F2: 0}
                        _ = v70
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        for fi71 := 0; fi71 < 16; fi71++ {
                            v12 += v12
                            var v72 int64 = 0
                            var v73 uint32 = 1
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                    }
                }
                v2 -= (v2 % ((v2) | 1))
                var v74 uint8 = v1
                _ = v74
            }
        case "foo bar":
            var v75 bool = true
            _ = v75
        default:
            switch (fi13 >> ((1) & 63)) {
            case 461:
                if v3 {
                    if false {
                        if (125 >= 0) {
                            var v76 float32 = 1e10
                            _ = v76
                        }
                    }
                    v77 := GenS5{F0: -387, F1: 747, F2: 2}
                    _ = v77
                } else {
                    var v78 float32 = 3.14159
                    _ = v78
                }
            case -761:
                _ = len(v4)
            }
        }
        v4[42] = 2147483647
        _ = fi13
    }
    var v79 int8 = 47
    v80 := GenS6{F0: "a"}
    switch "" {
    case "hello":
        var v81 float32 = 0.0
        _ = v81
    case "\n\t":
        v79 += (v79 ^ 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v79:", v79)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v79
    _ = v8
    _ = v80
    _ = v9
}
