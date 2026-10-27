package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int64; F2 int }
type GenS5 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 bool; F1 uint8 }
type GenS13 struct { F0 uint64 }
func genFn0() int {
    v7 := []int{948}
    var v8 int32 = 42
    fmt.Println(v8)
    var v9 int64 = 9223372036854775807
    var v10 uint16 = 878
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int32 = 42
    var v2 int8 = 0
    v3 := []string{}
    v4 := []int{}
    v5 := []int{1, 1}
    v6 := make(map[int]uint32)
    var v11 string = (((("a" + ("foo bar" + "hello")) + "hello") + ("foo bar" + "hello")) + "")
    fmt.Println(v1, v1)
    v12 := GenS1{F0: (-(v2))}
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v15 int64 = 0
            switch fi13 {
            case -9223372036854775808:
                if len(v4) > 0 {
                    v4[0] = fi14
                }
            case 945:
                fmt.Println(fi14)
            }
            _ = fi14
            _ = v15
        }
        if len(v4) > 0 {
            v4[0] = (^(fi13))
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        _ = len(v3)
        var v17 int16 = -927
        _ = fi16
        _ = v17
    }
    if true {
        switch "hello" {
        case "foo bar":
            if (151 >= 167) {
                switch ((v11 + v11) + ((v11 + v11) + "a")) {
                case "":
                    fmt.Println(v2)
                }
                if (2147483648 > 817) {
                    v18 := GenS2{F0: 245, F1: -1, F2: 32768}
                    fmt.Println(v11, v11)
                    _ = v18
                } else {
                    v19 := GenS2{F0: 108, F1: 9223372036854775807, F2: 127}
                    fmt.Println(v1, v1, v2)
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v21 := GenS2{F0: 28, F1: fi20, F2: -9223372036854775808}
                        switch v11 {
                        case "\n\t":
                            var v22 uint64 = 1
                            _ = v22
                        case "":
                            delete(v6, (fi20 | fi20))
                        default:
                            v23 := GenS3{F0: (fi20 & int(v1))}
                            _ = v23
                        }
                        switch v11 {
                        case "\n\t":
                            for fi24 := 0; fi24 < 19; fi24++ {
                                _ = len(v5)
                                if ((v1 & v1) >= v1) {
                                    var v25 uint16 = 42
                                    v26 := GenS4{F0: 0, F1: 1, F2: 609}
                                    var v27 float64 = 0.0
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    switch "" {
                                    case "hello":
                                        if false {
                                            var v28 uint16 = 2
                                            fmt.Println(v1, v2)
                                            _ = v28
                                        } else {
                                            var v29 uint = 0
                                            switch (v11 + "hello") {
                                            case "a":
                                                fmt.Println(fi20)
                                            case "\n\t":
                                                if true {
                                                    v30 := GenS5{F0: v29, F1: 255, F2: 65535}
                                                    v11 = ("\n\t" + "foo bar")
                                                    _ = v30
                                                }
                                            case "hello":
                                                var v31 int64 = int64(v2)
                                                _ = v31
                                            default:
                                                var v32 uint8 = 127
                                                _ = v32
                                            }
                                            if false {
                                                v29 = v29
                                                v29 = 18446744073709551615
                                            } else {
                                                if true {
                                                    v1 = v1
                                                    var v33 int16 = -294
                                                    v34 := GenS2{F0: 127, F1: fi24, F2: (fi24 | -918)}
                                                    _ = v33
                                                    _ = v34
                                                }
                                                delete(v6, fi20)
                                                v35 := GenS6{F0: v2, F1: 65535}
                                                _ = v35
                                            }
                                            _ = v29
                                        }
                                    default:
                                        if false {
                                            fmt.Println(v2, v2, v1)
                                            v11 = v11
                                            v36 := GenS3{F0: genFn0()}
                                            _ = v36
                                        }
                                    }
                                    v3 = append(v3, (v11 + v11))
                                }
                                v6[fi24] = 2147483648
                                _ = fi24
                            }
                        case "hello":
                            v37 := GenS7{F0: -879}
                            _ = v37
                        default:
                            if true {
                                v6[4294967295] = 2147483648
                                switch "hello" {
                                case "":
                                    for fi38 := 0; fi38 < 4; fi38++ {
                                        v1 += v1
                                        v2 -= v2
                                        _ = fi38
                                    }
                                default:
                                    v39 := GenS8{F0: (v11 < "a")}
                                    _ = v39
                                }
                                v2 -= v2
                            } else {
                                v40 := GenS2{F0: 152, F1: (fi20 * fi20), F2: (^(fi20))}
                                _ = v40
                            }
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
                if true {
                    switch 32767 {
                    case 621:
                        v41 := GenS9{F0: 802, F1: 255}
                        _ = v41
                    default:
                        _ = len(v3)
                    }
                    var v42 float64 = 3.14159
                    v1 -= (-(v1))
                    _ = v42
                } else {
                    fmt.Println(v2, v2, v11)
                    v11 = ((v11 + (v11 + "foo bar")) + "a")
                }
            } else {
                v43 := v6[32767]
                for fi44 := 0; fi44 < 18; fi44++ {
                    v1 -= (v1 + (^(v1)))
                    v45 := GenS4{F0: (v1 | (v1 ^ v1)), F1: -1, F2: (fi44 >> ((42) & 63))}
                    fmt.Println(v2)
                    _ = fi44
                    _ = v45
                }
                if (-2.25 == -2.25) {
                    var v46 uint16 = 32767
                    _ = v46
                } else {
                    v47 := GenS10{F0: 276}
                    v11 = "a"
                    v6[genFn0()] = (v43 << ((0) & 63))
                    _ = v47
                }
                _ = v43
            }
        default:
            if true {
                if ((v2 << ((65535) & 63)) < (v2 << ((1) & 63))) {
                    for fi48 := 0; fi48 < 5; fi48++ {
                        var v49 int = genFn0()
                        _ = fi48
                        _ = v49
                    }
                    var v50 uint = 42
                    v51 := GenS11{F0: 110, F1: 1}
                    _ = v50
                    _ = v51
                }
                var v52 int64 = 9223372036854775807
                _ = v52
            } else {
                switch 96 {
                case 409:
                    fmt.Println(v11, v11, v2)
                case 32767:
                    for fi53 := 0; fi53 < 5; fi53++ {
                        if (!((v11 + "\n\t") > (("\n\t" + v11) + v11))) {
                            var v54 uint = 1
                            _ = len(v4)
                            _ = v54
                        }
                        _ = fi53
                    }
                case -524:
                    v5 = append(v5, 1)
                default:
                    switch -9223372036854775808 {
                    case 42:
                        var v55 int32 = (v1 & v1)
                        _ = v55
                    }
                }
                fmt.Println(v1, v2, v2)
            }
        }
        var v56 int = 1
        v57 := v6[(v56 << ((321) & 63))]
        _ = v56
        _ = v57
    } else {
        var v58 float32 = (3.14159 - -2.25)
        var v59 float32 = v58
        _ = v58
        _ = v59
    }
    v60 := GenS7{F0: -172}
    for fi61 := 0; fi61 < 14; fi61++ {
        for fi62 := 0; fi62 < 8; fi62++ {
            v63 := GenS6{F0: v2, F1: 255}
            if (4294967295 > 291) {
                v2 = (v2 ^ v2)
                switch ("\n\t" + "") {
                case "":
                    switch genFn0() {
                    case -413:
                        if (true && false) {
                            var v64 uint32 = 1
                            _ = v64
                        }
                    case -2147483648:
                        if (84 != -128) {
                            v65 := GenS10{F0: 137}
                            v66 := GenS2{F0: 135, F1: (fi61 - 0), F2: genFn0()}
                            fmt.Println(fi62)
                            _ = v65
                            _ = v66
                        } else {
                            fmt.Println(v2, v11)
                        }
                    case -766:
                        v5 = append(v5, (fi62 * -2147483648))
                    }
                case "\n\t":
                    switch v11 {
                    case "a":
                        v67 := GenS12{F0: false, F1: 26}
                        _ = v67
                    case "\n\t":
                        v4 = append(v4, (^(fi61)))
                    default:
                        if (!(42 != -127)) {
                            if false {
                                fmt.Println(v2, fi61)
                            }
                        }
                    }
                case "hello":
                    var v68 uint = 1
                    _ = v68
                }
                var v69 string = v11
                _ = v69
            }
            if true {
                v1 = (v1 | v1)
                if (v1 == v1) {
                    v70 := GenS13{F0: 4294967295}
                    _ = v70
                }
            }
            _ = fi62
            _ = v63
        }
        v11 = ""
        _ = fi61
    }
    v11 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
