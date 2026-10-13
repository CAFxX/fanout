package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 string; F1 int32; F2 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 uint64; F2 int }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 int; F1 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint32 }
func genFn0(p0 uint32) uint64 {
    if false {
        p0 += p0
        p0 += p0
    } else {
        v7 := GenS1{F0: 18446744073709551615}
        p0 = 65535
        _ = v7
    }
    p0 = 470
    var v8 float32 = 0.1
    var v9 uint16 = 295
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := []int{-2147483648, -1}
    var v4 uint64 = 4294967295
    var v5 uint16 = 0
    v6 := make(map[uint32]bool)
    switch int(v4) {
    case 480:
        v4 += (v4 * 42)
    }
    var v10 int8 = -30
    v11 := GenS2{F0: 100, F1: false, F2: 2147483647}
    _ = len(v2)
    var v12 bool = (false || (true && ((true && true) && (!(0.1 > 0.0)))))
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            switch "" {
            case "foo bar":
                var v15 uint8 = 98
                _ = v15
            case "":
                v5 -= v5
            case "\n\t":
                var v16 string = (("\n\t" + "\n\t") + "foo bar")
                if false {
                    _ = len(v1)
                    v17 := GenS3{F0: -1, F1: (fi14 ^ 9), F2: 361}
                    fmt.Println(v4)
                    _ = v17
                } else {
                    var v18 string = ("\n\t" + "\n\t")
                    if v12 {
                        if v12 {
                            var v19 int = fi13
                            fmt.Println(v12)
                            _ = v19
                        } else {
                            v20 := GenS4{F0: v10}
                            switch fi13 {
                            case -2147483648:
                                var v21 uint8 = 2
                                _ = v21
                            case -162:
                                if (!true) {
                                    var v22 uint8 = 42
                                    _ = len(v3)
                                    _ = v22
                                } else {
                                    v23 := v2[-357]
                                    v23 += (fi14 + 4294967295)
                                    var v24 int8 = v10
                                    var v25 uint16 = v5
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                var v26 uint32 = 42
                                _ = v26
                            case 32768:
                                var v27 uint8 = 2
                                _ = v27
                            default:
                                delete(v1, v16)
                            }
                            var v28 string = v18
                            _ = v20
                            _ = v28
                        }
                        switch fi14 {
                        case 2147483647:
                            if len(v3) > 0 {
                                v3[0] = (fi14 + fi13)
                            }
                        default:
                            switch (^(fi14)) {
                            case 9223372036854775807:
                                var v29 string = v18
                                _ = v29
                            }
                        }
                    }
                    var v30 uint = 1
                    _ = v18
                    _ = v30
                }
                var v31 string = v16
                _ = v16
                _ = v31
            default:
                _ = len(v2)
            }
            _ = fi14
        }
        if (2 == 255) {
            var v32 int64 = int64(v4)
            for fi33 := 0; fi33 < 3; fi33++ {
                switch "hello" {
                case "":
                    v2[32768] = (fi13 | (-(fi33)))
                case "a":
                    fmt.Println(fi13, fi33)
                case "foo bar":
                    if (2 < 2) {
                        v4 = v4
                        v34 := GenS5{F0: (v10 > -128)}
                        _ = v34
                    } else {
                        v5 = 42
                    }
                }
                _ = fi33
            }
            if true {
                v5 = 42
                fmt.Println(fi13)
                if (!(18446744073709551615 == 196)) {
                    if v12 {
                        v35 := GenS6{F0: "a", F1: "foo bar"}
                        var v36 uint16 = (v5 % ((v5) | 1))
                        v37 := GenS5{F0: v12}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                } else {
                    v38 := GenS3{F0: (v10 / (((v10 & v10)) | 1)), F1: (fi13 + (fi13 & (fi13 | (^(fi13))))), F2: 2147483648}
                    _ = v38
                }
            } else {
                var v39 int = 42
                v4 += (v4 / (((v4 - 4294967295)) | 1))
                _ = v39
            }
            _ = v32
        } else {
            v40 := GenS7{F0: 2, F1: 65535}
            switch int(v10) {
            case 954:
                if ((v12 && true) || false) {
                    switch fi13 {
                    case 609:
                        if v12 {
                            var v41 int16 = 138
                            var v42 uint8 = 17
                            if ((((v12 && v12) && v12) || (!((v12 && (18446744073709551615 == 4294967295)) || (78 == v42)))) || true) {
                                v5 = v5
                                for fi43 := 0; fi43 < 19; fi43++ {
                                    v10 -= v10
                                    _ = fi43
                                }
                            } else {
                                var v44 int16 = v41
                                if len(v3) > 0 {
                                    v3[0] = fi13
                                }
                                _ = v44
                            }
                            _ = v41
                            _ = v42
                        }
                    default:
                        _ = len(v2)
                    }
                    v12 = (false || (42 >= v4))
                } else {
                    v10 = int8(v4)
                    var v45 string = ("" + (("\n\t" + ("" + "a")) + "a"))
                    v5 = v5
                    _ = v45
                }
            case 23:
                delete(v1, ("\n\t" + "foo bar"))
            default:
                _ = len(v3)
            }
            _ = v40
        }
        _ = fi13
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v12 {
        switch -566 {
        case 1:
            var v46 uint = 246
            _ = v46
        }
        v47 := GenS6{F0: "", F1: "hello"}
        _ = v47
    } else {
        v4 = (v4 | (^(v4)))
        var v48 uint32 = uint32(v5)
        _ = v48
    }
    v49 := GenS8{F0: 152, F1: 1}
    v50 := GenS9{F0: ("" + "foo bar"), F1: 1, F2: 42}
    if v12 {
        for fi51 := 0; fi51 < 13; fi51++ {
            for fi52 := 0; fi52 < 2; fi52++ {
                v53 := GenS10{F0: true}
                var v54 uint32 = uint32(v5)
                var v55 int16 = int16(v5)
                _ = fi52
                _ = v53
                _ = v54
                _ = v55
            }
            _ = fi51
        }
        switch 1 {
        case 32767:
            var v56 uint = 1
            _ = v56
        case 613:
            if false {
                v57 := GenS11{F0: 32768, F1: 4294967295, F2: 717}
                fmt.Println(v5, v10)
                if ("\n\t" <= "a") {
                    if v12 {
                        v5 -= v5
                    }
                    switch "foo bar" {
                    case "":
                        v58 := GenS10{F0: (-1 != v10)}
                        _ = v58
                    case "foo bar":
                        v59 := GenS12{F0: 4294967295, F1: v5}
                        _ = v59
                    case "a":
                        if v12 {
                            var v60 uint32 = 42
                            switch "foo bar" {
                            case "":
                                v1[(("foo bar" + ("\n\t" + ("hello" + ""))) + (("" + ("hello" + "\n\t")) + "a"))] = 2147483647
                            case "foo bar":
                                v12 = v12
                            case "\n\t":
                                v61 := GenS13{F0: 1, F1: genFn0(v60)}
                                _ = v61
                            }
                            _ = v60
                        } else {
                            var v62 float32 = float32(1.5)
                            for fi63 := 0; fi63 < 5; fi63++ {
                                for fi64 := 0; fi64 < 4; fi64++ {
                                    fmt.Println(fi64)
                                    fmt.Println(fi63)
                                    _ = fi64
                                }
                                _ = len(v3)
                                v65 := GenS9{F0: "\n\t", F1: 42, F2: -250}
                                _ = fi63
                                _ = v65
                            }
                            var v66 uint16 = v5
                            _ = v62
                            _ = v66
                        }
                    }
                } else {
                    v12 = (42 == v10)
                    var v67 float32 = (((float32(1e10) / 0.1) / 3.14159) - (1.5 + 0.1))
                    var v68 uint64 = v4
                    _ = v67
                    _ = v68
                }
                _ = v57
            } else {
                v12 = v12
                if (2147483647 == 4294967295) {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            v4 = 816
                        }
                        v5 = (v5 + v5)
                        var v69 string = "\n\t"
                        _ = v69
                    }
                    var v70 int = 674
                    _ = v70
                } else {
                    for fi71 := 0; fi71 < 14; fi71++ {
                        var v72 int64 = -1
                        if (255 != 42) {
                            var v73 int8 = (v10 * v10)
                            v5 += v5
                            _ = v73
                        } else {
                            if false {
                                v74 := GenS9{F0: ("\n\t" + "\n\t"), F1: -1, F2: v72}
                                for fi75 := 0; fi75 < 3; fi75++ {
                                    var v76 uint = 461
                                    switch fi71 {
                                    case 0:
                                        fmt.Println(v4)
                                    case 32767:
                                        var v77 uint = 1
                                        _ = v77
                                    case -707:
                                        var v78 float64 = ((-2.25 * ((0.1 * 0.0) * float64(0.1))) / (1e10 / (-2.25 - (1.5 + 3.14159))))
                                        _ = v78
                                    }
                                    v10 -= v10
                                    _ = fi75
                                    _ = v76
                                }
                                for fi79 := 0; fi79 < 1; fi79++ {
                                    var v80 int16 = 1
                                    if (!v12) {
                                        v81 := v1["foo bar"]
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v81
                                    } else {
                                        v82 := GenS12{F0: (v4 | (v4 * v4)), F1: 65535}
                                        v12 = (42 > 42)
                                        _ = v82
                                    }
                                    if v12 {
                                        v83 := GenS14{F0: 42}
                                        v84 := GenS11{F0: -527, F1: (^(v4)), F2: fi79}
                                        v72 += -9223372036854775808
                                        _ = v83
                                        _ = v84
                                    } else {
                                        _ = len(v1)
                                        var v85 int64 = -9223372036854775808
                                        v5 = v5
                                        _ = v85
                                    }
                                    _ = fi79
                                    _ = v80
                                }
                                _ = v74
                            }
                            v86 := GenS13{F0: 42, F1: 1}
                            v72 -= 1
                            _ = v86
                        }
                        _ = fi71
                        _ = v72
                    }
                    v4 -= 18446744073709551615
                }
            }
        case 128:
            v87 := GenS15{F0: v12}
            _ = v87
        default:
            v12 = v12
        }
        var v88 int16 = -32768
        _ = v88
    }
    v89 := GenS16{F0: 796}
    var v90 float32 = 0.1
    var v91 uint16 = v5
    for fi92 := 0; fi92 < 18; fi92++ {
        fmt.Println(fi92, v4, v4)
        _ = fi92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
