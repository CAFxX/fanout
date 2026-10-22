package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 int; F2 string }
type GenS9 struct { F0 int64; F1 uint; F2 uint16 }
func genFn0(p0 uint32, p1 uint32) int16 {
    v7 := GenS1{F0: 0, F1: 127}
    v8 := GenS1{F0: 42, F1: -94}
    switch "" {
    case "hello":
        v9 := GenS1{F0: 31, F1: 1}
        _ = v9
    default:
        var v10 bool = (1.5 != 1.5)
        _ = v10
    }
    v11 := []int{32768, -882}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64, p1 uint8, p2 uint32) uint16 {
    if false {
        v12 := []int{32767, -246, 473, -9223372036854775808}
        _ = v12
    } else {
        fmt.Println(p0, p2, p1)
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int = fi14
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: "foo bar"}
        _ = fi13
        _ = v16
    }
    p1 = uint8(p2)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    var v5 bool = false
    v6 := make(map[int]string)
    v17 := v4[""]
    switch 42 {
    case 291:
        if true {
            var v18 bool = v17
            _ = v18
        } else {
            delete(v6, 127)
            v19 := GenS3{F0: v2, F1: v2}
            _ = v19
        }
    }
    v2 += -779
    if true {
        switch 28 {
        case -1:
            switch "hello" {
            case "foo bar":
                delete(v6, -357)
            default:
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 uint32 = 2147483647
                    v17 = v5
                    var v22 bool = v5
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        default:
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int = fi23
                _ = fi23
                _ = v24
            }
        }
        v5 = (84 > 2)
    } else {
        var v25 int32 = (^(v2))
        v2 -= v2
        var v26 uint16 = 65535
        _ = v25
        _ = v26
    }
    switch ("hello" + "foo bar") {
    case "":
        var v27 int16 = 42
        _ = v27
    case "a":
        if true {
            v28 := []int{-84, 2147483647, 42, 42}
            _ = v28
        } else {
            switch 0 {
            case 32768:
                for fi29 := 0; fi29 < 20; fi29++ {
                    v30 := []int{(^(fi29)), fi29, (fi29 & fi29)}
                    _ = fi29
                    _ = v30
                }
            case 42:
                var v31 bool = (v2 < (v2 & (^(v2))))
                _ = v31
            case 9223372036854775807:
                fmt.Println(v2)
            default:
                fmt.Println(v2)
            }
            switch "foo bar" {
            case "a":
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 uint32 = 4294967295
                    v17 = v17
                    _ = fi32
                    _ = v33
                }
            case "foo bar":
                fmt.Println(v17, v2, v5)
            default:
                v2 -= v2
            }
        }
    case "\n\t":
        for fi34 := 0; fi34 < 3; fi34++ {
            for fi35 := 0; fi35 < 11; fi35++ {
                if ((false || v5) || (48 < 127)) {
                    fmt.Println(v17, v2)
                    var v36 uint64 = 0
                    v2 -= (v2 + 778)
                    _ = v36
                }
                _ = fi35
            }
            _ = fi34
        }
    default:
        var v37 int = 42
        _ = v37
    }
    if v17 {
        v38 := GenS4{F0: (((float32((((0.0 * 0.1) / ((-2.25 * float64((3.14159 * 1e10))) * (3.14159 - 0.1))) / -2.25)) / 1.0) * 1e10) >= 1e10)}
        v5 = (v5 && v17)
        var v39 int8 = 123
        _ = v38
        _ = v39
    } else {
        fmt.Println(v2)
        for fi40 := 0; fi40 < 10; fi40++ {
            if ((v5 || false) || v17) {
                var v41 int16 = genFn0(0, 2147483648)
                _ = v41
            } else {
                var v42 int64 = 338
                if (v5 && v17) {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        v44 := GenS5{F0: 4294967295, F1: "foo bar"}
                        var v45 int16 = -32768
                        fmt.Println(v5, fi40, v5)
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    var v46 uint = 42
                    _ = v46
                } else {
                    _ = len(v4)
                }
                if (127 == 12) {
                    v47 := GenS2{F0: "a"}
                    v5 = (!(fi40 == (fi40 - fi40)))
                    _ = v47
                }
                _ = v42
            }
            fmt.Println(v17)
            _ = fi40
        }
        for fi48 := 0; fi48 < 13; fi48++ {
            switch "\n\t" {
            case "hello":
                if (v17 && false) {
                    if (546 >= genFn1(1, 2, 557)) {
                        if (((v2 < v2) || (false || ((-952 > 0) && v17))) || true) {
                            if (-2.25 != -2.25) {
                                fmt.Println(v2, v2, v2)
                                for fi49 := 0; fi49 < 17; fi49++ {
                                    v50 := GenS2{F0: "a"}
                                    var v51 uint64 = 0
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    fmt.Println(v17, fi48)
                                    _ = fi52
                                }
                            } else {
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    var v53 uint32 = 2
                                    _ = v53
                                case "hello":
                                    var v54 int16 = 1000
                                    _ = v54
                                case "foo bar":
                                    switch "" {
                                    case "":
                                        var v55 int = fi48
                                        _ = v55
                                    }
                                default:
                                    var v56 int = (fi48 << ((0) & 63))
                                    _ = v56
                                }
                                var v57 uint32 = 42
                                var v58 uint16 = 1
                                _ = v57
                                _ = v58
                            }
                        } else {
                            for fi59 := 0; fi59 < 3; fi59++ {
                                v5 = true
                                _ = fi59
                            }
                            var v60 int64 = 0
                            switch -9223372036854775808 {
                            case 32768:
                                v60 -= (v60 % (((v60 ^ 125)) | 1))
                            default:
                                var v61 uint8 = 42
                                _ = v61
                            }
                            _ = v60
                        }
                        var v62 uint8 = 1
                        switch "\n\t" {
                        case "":
                            var v63 string = "\n\t"
                            _ = v63
                        }
                        _ = v62
                    }
                }
            default:
                v2 = (^(v2))
            }
            v64 := []int{}
            v5 = ((42 <= 127) && true)
            _ = fi48
            _ = v64
        }
    }
    var v65 float64 = (((((-2.25 + (0.1 + 1.5)) * 1e10) + 0.1) + 0.0) + ((1e10 / 3.14159) * float64(3.14159)))
    var v66 uint64 = 1
    for fi67 := 0; fi67 < 9; fi67++ {
        fmt.Println(v17, v66, v5)
        _ = fi67
    }
    if (!((!v5) && (false && true))) {
        if v17 {
            for fi68 := 0; fi68 < 4; fi68++ {
                v5 = true
                _ = fi68
            }
        } else {
            var v69 uint16 = 900
            delete(v4, ("a" + ((("a" + ((("" + ("foo bar" + "foo bar")) + (("" + ("a" + "foo bar")) + "a")) + ("\n\t" + "hello"))) + ("hello" + "a")) + "foo bar")))
            _ = v69
        }
    } else {
        v70 := GenS6{F0: 1}
        _ = v70
    }
    v66 = v66
    for fi71 := 0; fi71 < 4; fi71++ {
        v72 := []int{fi71, (-(fi71))}
        for fi73 := 0; fi73 < 7; fi73++ {
            v74 := GenS7{F0: 1}
            v75 := GenS8{F0: 255, F1: (fi73 + fi73), F2: "hello"}
            var v76 uint = 42
            v77 := GenS1{F0: 42, F1: -128}
            _ = fi73
            _ = v74
            _ = v75
            _ = v76
            _ = v77
        }
        _ = fi71
        _ = v72
    }
    var v78 uint = 18446744073709551615
    if (true || true) {
        var v79 bool = (!v17)
        v65 = 0.1
        _ = v79
    } else {
        var v80 int64 = -9223372036854775808
        _ = v80
    }
    var v81 uint64 = 4294967295
    var v82 bool = v17
    for fi83 := 0; fi83 < 11; fi83++ {
        var v84 uint8 = 128
        v85 := GenS9{F0: 1, F1: uint(v81), F2: 2}
        _ = fi83
        _ = v84
        _ = v85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v78
    _ = v81
    _ = v82
}
