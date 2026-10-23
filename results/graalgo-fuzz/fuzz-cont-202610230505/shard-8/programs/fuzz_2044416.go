package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint16; F1 int32; F2 int8 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 int8; F1 uint }
func genFn0() int8 {
    if true {
        var v7 int32 = -1
        fmt.Println(v7)
        _ = v7
    } else {
        var v8 int32 = -2147483648
        _ = v8
    }
    var v9 string = (("foo bar" + ("a" + "\n\t")) + "foo bar")
    fmt.Println(v9)
    fmt.Println(v9)
    var v10 uint32 = 4294967295
    _ = v10
    _ = v9
    return 100
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 42
    var v5 int16 = 42
    v6 := []uint32{1, 0, 262, 65535}
    v5 = (v5 / (((^(v5))) | 1))
    var v11 uint8 = 84
    v3[2147483647] = 42
    v12 := GenS1{F0: 1}
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 bool = (false && (v11 == (^(v11))))
            v15 = true
            fmt.Println(fi14, fi13, v4)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v16 bool = (674 <= 4294967295)
    v17 := GenS1{F0: 0}
    if v16 {
        for fi18 := 0; fi18 < 2; fi18++ {
            if true {
                v19 := GenS2{F0: 2, F1: (fi18 << ((0) & 63)), F2: genFn0()}
                for fi20 := 0; fi20 < 6; fi20++ {
                    v21 := GenS3{F0: false}
                    v5 += v5
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                v3[2] = 2147483647
            }
            _ = fi18
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v3[4294967295] = 42
    }
    var v22 int16 = (v5 >> ((18446744073709551615) & 63))
    v23 := GenS1{F0: v4}
    switch 1 {
    case -253:
        v24 := GenS1{F0: (v4 & genFn0())}
        _ = v24
    case -221:
        v25 := GenS3{F0: (-897 < 2147483647)}
        _ = v25
    case 128:
        var v26 string = "\n\t"
        _ = v26
    }
    if (3.14159 >= float64((3.14159 / (3.14159 / 0.1)))) {
        v27 := GenS3{F0: true}
        if v16 {
            v28 := GenS4{F0: 42}
            _ = v28
        } else {
            _ = len(v6)
        }
        _ = v27
    } else {
        for fi29 := 0; fi29 < 15; fi29++ {
            switch (("foo bar" + ("foo bar" + "foo bar")) + "\n\t") {
            case "hello":
                v11 -= v11
            default:
                if (42 < 0) {
                    if v16 {
                        v11 = v11
                        var v30 uint64 = 994
                        _ = v30
                    } else {
                        fmt.Println(v5, v16)
                        fmt.Println(fi29, v11)
                        for fi31 := 0; fi31 < 11; fi31++ {
                            delete(v2, "hello")
                            v32 := GenS3{F0: v16}
                            _ = fi31
                            _ = v32
                        }
                    }
                    v11 -= (v11 * v11)
                    var v33 int = 42
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v35 := GenS5{F0: ((3.14159 * 3.14159) != (0.0 - (float64((1e10 - 0.0)) * float64(-2.25))))}
                        var v36 uint16 = 1
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
            switch "" {
            case "a":
                delete(v1, "a")
            case "foo bar":
                if v16 {
                    v5 += (v5 % (((v22 << ((836) & 63))) | 1))
                } else {
                    switch "foo bar" {
                    case "":
                        var v37 int16 = 1
                        _ = v37
                    case "hello":
                        var v38 uint8 = v11
                        _ = v38
                    default:
                        switch "" {
                        case "foo bar":
                            v2["foo bar"] = false
                        case "\n\t":
                            var v39 float32 = (-2.25 / 0.1)
                            _ = v39
                        case "hello":
                            switch fi29 {
                            case 2147483647:
                                var v40 uint16 = 65535
                                _ = v40
                            case -1:
                                v11 = v11
                            case -2147483648:
                                v41 := GenS6{F0: (1.5 <= ((-2.25 + (float32((-2.25 - 0.0)) + 0.0)) * 3.14159)), F1: (fi29 << ((2) & 63))}
                                _ = v41
                            }
                        default:
                            if ((false || v16) || v16) {
                                if (301 != 1) {
                                    var v42 uint64 = 0
                                    if v16 {
                                        v42 = 18446744073709551615
                                        v43 := GenS7{F0: 32767, F1: 37, F2: v4}
                                        _ = v43
                                    }
                                    v44 := GenS5{F0: true}
                                    _ = v42
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 10; fi45++ {
                                        if v16 {
                                            var v46 bool = ((0 <= 42) && v16)
                                            fmt.Println(v4, v22)
                                            _ = v46
                                        }
                                        _ = fi45
                                    }
                                }
                                var v47 float64 = 1e10
                                var v48 int32 = 1
                                _ = v47
                                _ = v48
                            } else {
                                if (3.14159 <= 1.5) {
                                    fmt.Println(v4, fi29, v4)
                                } else {
                                    v4 -= genFn0()
                                    fmt.Println(v11)
                                }
                            }
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = 1
                    }
                }
            case "hello":
                for fi49 := 0; fi49 < 19; fi49++ {
                    var v50 uint64 = 1
                    var v51 int = 0
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            default:
                fmt.Println(v11, v11)
            }
            _ = fi29
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            switch "foo bar" {
            case "foo bar":
                v53 := GenS8{F0: ("" + "foo bar"), F1: 42}
                _ = v53
            default:
                if ((v11 * 1) == (v11 * 127)) {
                    for fi54 := 0; fi54 < 5; fi54++ {
                        var v55 int64 = -587
                        if (v11 < (v11 ^ (v11 ^ (v11 / ((v11) | 1))))) {
                            var v56 string = "\n\t"
                            _ = v56
                        }
                        if (v16 && true) {
                            _ = len(v2)
                        } else {
                            var v57 uint8 = (v11 / (((v11 - v11)) | 1))
                            v16 = v16
                            switch fi54 {
                            case 2147483647:
                                v57 -= (v57 | (v57 / (((v11 >> ((2) & 63))) | 1)))
                                if ((!v16) || false) {
                                    fmt.Println(fi52, v55, v16)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v58 := GenS9{F0: (v4 % ((42) | 1))}
                                    _ = v58
                                }
                                var v59 int8 = v4
                                _ = v59
                            default:
                                if false {
                                    var v60 int32 = 789
                                    switch (-(fi54)) {
                                    case 128:
                                        v61 := GenS10{F0: (-(v60)), F1: ("a" + "hello")}
                                        _ = v61
                                    case 4294967295:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case 42:
                                        fmt.Println(v55, fi52, fi54)
                                    default:
                                        switch (^(fi54)) {
                                        case 931:
                                            var v62 string = ""
                                            _ = v62
                                        case -9223372036854775808:
                                            v60 -= -2147483648
                                        case 42:
                                            fmt.Println(v16, v57)
                                        default:
                                            fmt.Println(v11)
                                        }
                                    }
                                    if v16 {
                                        var v63 uint8 = 97
                                        _ = v63
                                    } else {
                                        if false {
                                            v55 += v55
                                        } else {
                                            v57 = (v57 / ((v11) | 1))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        v11 = (v11 & v11)
                                        delete(v2, "")
                                    }
                                    _ = len(v2)
                                    _ = v60
                                }
                            }
                            _ = v57
                        }
                        _ = fi54
                        _ = v55
                    }
                } else {
                    v64 := GenS2{F0: 2, F1: fi52, F2: v4}
                    var v65 int8 = -128
                    var v66 int = (fi52 % (((fi52 | (fi52 + fi52))) | 1))
                    _ = v64
                    _ = v65
                    _ = v66
                }
            }
            if (!v16) {
                v4 -= (v4 >> ((65535) & 63))
                v22 = (v5 - -32768)
            } else {
                fmt.Println(fi52, v4)
                for fi67 := 0; fi67 < 6; fi67++ {
                    v16 = (!true)
                    v68 := v2["foo bar"]
                    _ = fi67
                    _ = v68
                }
            }
            _ = fi52
        }
        v69 := GenS11{F0: 42, F1: 2}
        _ = v69
    }
    v5 = v22
    v22 = (-(v22))
    v70 := GenS6{F0: false, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
