package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32; F1 int; F2 int }
type GenS6 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0(p0 int, p1 int16, p2 bool) bool {
    v5 := GenS1{F0: (!p2)}
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "\n\t" {
        case "a":
            switch (("\n\t" + "") + (((("" + (("hello" + "") + "foo bar")) + "\n\t") + "") + ("" + ""))) {
            case "":
                p1 += (p1 - 32767)
            case "hello":
                for fi7 := 0; fi7 < 16; fi7++ {
                    var v8 bool = (134 > -405)
                    p1 = p1
                    fmt.Println(v8, fi7, p0)
                    _ = fi7
                    _ = v8
                }
            case "\n\t":
                var v9 uint = 255
                _ = v9
            }
        }
        p1 = 32767
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (-2147483648 < 42)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{0}
    var v3 int32 = -2147483648
    var v4 int16 = 1
    v10 := v1[32767]
    v3 = v3
    var v11 uint8 = 160
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            _ = len(v1)
            v10 = v10
            v11 -= v11
            if (!(42 == 47)) {
                var v14 uint64 = 1
                _ = v14
            } else {
                v15 := GenS2{F0: 688, F1: v11, F2: 1}
                fmt.Println(fi12)
                v16 := GenS3{F0: 0}
                _ = v15
                _ = v16
            }
            var v17 uint8 = 127
            _ = fi13
            _ = v17
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            if (-9223372036854775808 == 0) {
                var v19 int32 = v3
                v4 = v4
                _ = v19
            }
            var v20 bool = (v10 || (v10 && v10))
            switch "foo bar" {
            case "foo bar":
                switch 0 {
                case 127:
                    switch fi12 {
                    case 9223372036854775807:
                        var v21 uint8 = (v11 + (^(v11)))
                        _ = v21
                    case -1:
                        if v20 {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                fmt.Println(v3)
                            }
                            v4 -= (v4 ^ (v4 | (v4 ^ -92)))
                            var v22 float32 = 0.0
                            _ = v22
                        } else {
                            v11 += v11
                        }
                    case 2147483647:
                        var v23 uint32 = 2
                        _ = v23
                    }
                case -9223372036854775808:
                    fmt.Println(v11, fi12)
                case -429:
                    v24 := GenS3{F0: 255}
                    _ = v24
                default:
                    fmt.Println(v10, v11)
                }
            case "":
                v1[int(v4)] = (70 != -18)
            case "\n\t":
                if ((((((("hello" + "a") + ("\n\t" + ("a" + "foo bar"))) + "a") + ("\n\t" + "hello")) + "a") + ("a" + "\n\t")) <= "") {
                    v25 := GenS4{F0: (-49 < 1), F1: 65535, F2: -1}
                    if (uint16(v11) < 910) {
                        if ((true || v20) || v20) {
                            var v26 int8 = 40
                            var v27 int32 = (v3 << ((336) & 63))
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = v25
                } else {
                    if (!((int64(fi12) > int64(v11)) && true)) {
                        fmt.Println(v3, v20, v4)
                        var v28 uint32 = 2
                        switch "" {
                        case "":
                            v29 := v1[fi18]
                            _ = v29
                        }
                        _ = v28
                    } else {
                        v20 = (3.14159 != -2.25)
                        v30 := GenS2{F0: 959, F1: (v11 >> ((650) & 63)), F2: 2}
                        if true {
                            v31 := GenS5{F0: 889, F1: 0, F2: (fi18 >> ((uint(v11)) & 63))}
                            if ((v3 == (-(v3))) || (!(42 != 4294967295))) {
                                switch fi12 {
                                case 57:
                                    v20 = ("a" > "foo bar")
                                }
                                var v32 float32 = 0.0
                                fmt.Println(v3, v32)
                                var v33 int8 = 55
                                _ = v32
                                _ = v33
                            }
                            if (!v20) {
                                fmt.Println(fi18, fi18)
                                var v34 float32 = float32(1.5)
                                switch (fi18 | (fi12 / ((-2147483648) | 1))) {
                                case -9223372036854775808:
                                    if (v20 || v20) {
                                        v35 := GenS4{F0: v20, F1: 42, F2: -106}
                                        var v36 int32 = v3
                                        fmt.Println(v36, v4, v34)
                                        _ = len(v2)
                                        v4 += 155
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                _ = v34
                            } else {
                                v37 := GenS2{F0: 552, F1: v11, F2: 4294967295}
                                var v38 uint32 = 4294967295
                                var v39 uint8 = (v11 << ((0) & 63))
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            _ = v31
                        }
                        fmt.Println(v10, v10, fi18)
                        v3 = (v3 | (^(v3)))
                        switch ("hello" + "hello") {
                        case "\n\t":
                            delete(v1, (fi12 >> ((0) & 63)))
                        case "":
                            switch (("\n\t" + (("\n\t" + ("" + "hello")) + ("hello" + "hello"))) + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                var v40 int = fi18
                                _ = v40
                            case "foo bar":
                                if v10 {
                                    var v41 uint32 = 408
                                    v20 = (v3 != v3)
                                    var v42 int8 = -1
                                    v11 += (^(v11))
                                    if true {
                                        v10 = true
                                        var v43 float32 = 0.0
                                        _ = v43
                                    }
                                    _ = v41
                                    _ = v42
                                }
                            }
                        case "foo bar":
                            var v44 bool = false
                            _ = v44
                        default:
                            v4 -= 899
                        }
                        v45 := GenS3{F0: 0}
                        if (32767 == 756) {
                            v11 += v11
                        }
                        if true {
                            switch "" {
                            case "a":
                                if v10 {
                                    _ = len(v1)
                                    if (v3 == -1) {
                                        if true {
                                            if ((v4 - (v4 + (^(v4)))) >= v4) {
                                                _ = len(v2)
                                            }
                                            v3 = v3
                                            v11 = (v11 | v11)
                                        } else {
                                            if (v4 != (v4 + v4)) {
                                                if (true || ("foo bar" != "")) {
                                                    var v46 int8 = -1
                                                    var v47 int = 4294967295
                                                    v48 := GenS3{F0: 42}
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v20)
                                } else {
                                    v2 = append(v2, 2147483647)
                                    v3 = 1
                                }
                            case "hello":
                                v20 = ((((!v20) || (-58 > 100)) || true) && (!(!(38 < 1))))
                            case "foo bar":
                                var v49 uint32 = 2147483647
                                _ = v49
                            }
                            var v50 int16 = (v4 & (v4 - -32768))
                            var v51 float64 = (-2.25 / 1.0)
                            _ = v50
                            _ = v51
                        }
                        _ = v30
                        _ = v45
                    }
                }
            }
            v4 = -1
            _ = fi18
            _ = v20
        }
        _ = fi12
    }
    switch -1 {
    case 128:
        v52 := GenS6{F0: 255, F1: (v3 << ((211) & 63)), F2: v4}
        _ = v52
    case 4294967295:
        v11 += 228
    }
    if (!v10) {
        if false {
            for fi53 := 0; fi53 < 19; fi53++ {
                v11 = v11
                _ = fi53
            }
            var v54 int = -1
            delete(v1, v54)
            _ = v54
        }
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        var v55 int8 = 127
        _ = v55
    } else {
        fmt.Println(v4, v3, v4)
    }
    for fi56 := 0; fi56 < 2; fi56++ {
        var v57 uint16 = 527
        var v58 int16 = (v4 - (^(v4)))
        _ = fi56
        _ = v57
        _ = v58
    }
    fmt.Println(v3)
    switch "" {
    case "hello":
        if (1.5 > float64(1e10)) {
            v59 := GenS1{F0: (!false)}
            var v60 int = -997
            _ = v59
            _ = v60
        }
    case "\n\t":
        for fi61 := 0; fi61 < 18; fi61++ {
            if (599 <= -1) {
                var v62 uint64 = 0
                v11 += 197
                v63 := GenS2{F0: 42, F1: 128, F2: 255}
                _ = v62
                _ = v63
            }
            v3 -= -1
            fmt.Println(v4, v4, v10)
            _ = fi61
        }
    case "a":
        for fi64 := 0; fi64 < 9; fi64++ {
            fmt.Println(v11)
            switch "a" {
            case "\n\t":
                v1[(fi64 >> ((626) & 63))] = true
            case "hello":
                var v65 float32 = 1e10
                _ = v65
            }
            v66 := v1[fi64]
            _ = fi64
            _ = v66
        }
    default:
        var v67 uint64 = 0
        _ = v67
    }
    var v68 uint8 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v68
}
