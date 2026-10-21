package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int32; F1 bool; F2 int }
type GenS7 struct { F0 int32; F1 string; F2 int8 }
type GenS8 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint64 = 906
    v2 := []string{}
    var v3 uint16 = 255
    v4 := make(map[uint32]bool)
    var v5 bool = (9223372036854775807 <= 50)
    var v6 int32 = 0
    if true {
        _ = len(v4)
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := GenS1{F0: 65535, F1: 42, F2: -509}
                v3 += 32767
                v9 := GenS2{F0: 1}
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            fmt.Println(v1, v6)
            fmt.Println(v3, v5)
            v10 := GenS1{F0: 2147483647, F1: 743, F2: -77}
            _ = v10
        }
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "hello":
            var v11 uint = 4294967295
            _ = v11
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                v5 = false
                if (!(111 != 49)) {
                    var v13 int = fi12
                    var v14 int32 = (v6 | -786)
                    _ = v13
                    _ = v14
                }
                _ = len(v2)
                _ = fi12
            }
        }
    }
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 18; fi16++ {
        v1 = v1
        for fi17 := 0; fi17 < 12; fi17++ {
            v6 = v6
            if true {
                v6 = v6
            }
            var v18 uint16 = 2
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS1{F0: 1, F1: 0, F2: 9223372036854775807}
    if v5 {
        if true {
            for fi20 := 0; fi20 < 3; fi20++ {
                if v5 {
                    var v21 string = "\n\t"
                    _ = v21
                } else {
                    var v22 uint8 = 42
                    v1 -= 0
                    _ = v22
                }
                v23 := GenS3{F0: v5}
                v1 = (^(v1))
                _ = fi20
                _ = v23
            }
            if ((!v5) && true) {
                v24 := GenS3{F0: v5}
                _ = v24
            }
        } else {
            if (v5 || (v5 && true)) {
                v25 := GenS4{F0: 4294967295, F1: (v15 | v15), F2: v1}
                var v26 int8 = 42
                _ = v25
                _ = v26
            }
        }
    } else {
        if (v5 || (43 != 18)) {
            if (("foo bar" + (("" + ("a" + ("hello" + "foo bar"))) + "a")) != "a") {
                v1 = (v1 << ((605) & 63))
                for fi27 := 0; fi27 < 12; fi27++ {
                    v28 := GenS1{F0: 4294967295, F1: 670, F2: 1}
                    for fi29 := 0; fi29 < 20; fi29++ {
                        fmt.Println(v6)
                        _ = fi29
                    }
                    _ = fi27
                    _ = v28
                }
                switch -9223372036854775808 {
                case 32767:
                    v1 -= (v1 - 4294967295)
                case 0:
                    if (v3 == v3) {
                        switch "\n\t" {
                        case "\n\t":
                            var v30 uint32 = 1
                            _ = v30
                        case "hello":
                            v2 = append(v2, "foo bar")
                        }
                        fmt.Println(v3)
                    } else {
                        switch "foo bar" {
                        case "":
                            v3 += v3
                        case "a":
                            if (4 == 2) {
                                switch int(v3) {
                                case 128:
                                    var v31 float32 = (float32(-2.25) - (1.5 * 0.0))
                                    _ = v31
                                case 32768:
                                    for fi32 := 0; fi32 < 6; fi32++ {
                                        var v33 uint64 = (v1 - (v1 << ((18446744073709551615) & 63)))
                                        v1 = (v33 | 4294967295)
                                        _ = fi32
                                        _ = v33
                                    }
                                }
                                var v34 int8 = 64
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    v5 = (!(!v5))
                                    v36 := GenS4{F0: 65535, F1: 1, F2: 1}
                                    _ = fi35
                                    _ = v36
                                }
                                _ = v34
                            } else {
                                var v37 int64 = 220
                                v3 -= v3
                                _ = v37
                            }
                        case "foo bar":
                            var v38 int32 = v6
                            _ = v38
                        }
                    }
                default:
                    v39 := GenS5{F0: 49, F1: (v1 << ((255) & 63)), F2: 2}
                    _ = v39
                }
            } else {
                var v40 int = 128
                v41 := GenS1{F0: 333, F1: 4294967295, F2: -344}
                _ = v40
                _ = v41
            }
            var v42 uint16 = (v3 >> ((1) & 63))
            v1 = uint64(v42)
            _ = v42
        } else {
            v15 = (-(v15))
        }
        v43 := v4[42]
        _ = v43
    }
    if (("foo bar" + ("hello" + (("hello" + "\n\t") + ("hello" + "\n\t")))) > "foo bar") {
        var v44 uint16 = (v3 >> ((0) & 63))
        v5 = true
        _ = v44
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch int(v15) {
    case 32767:
        for fi45 := 0; fi45 < 6; fi45++ {
            for fi46 := 0; fi46 < 9; fi46++ {
                if false {
                    var v47 uint64 = (v1 >> ((2) & 63))
                    _ = v47
                }
                if (1 == 2147483647) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v48 uint32 = 2147483647
                    _ = v48
                }
                _ = fi46
            }
            var v49 bool = (837 == 9223372036854775807)
            v6 -= 1
            _ = fi45
            _ = v49
        }
    }
    v3 -= (v3 >> ((42) & 63))
    var v50 int = -2147483648
    v3 = (v3 % (((v3 + 0)) | 1))
    v51 := GenS2{F0: (v3 / (((^(v3))) | 1))}
    for fi52 := 0; fi52 < 10; fi52++ {
        if (32768 < fi52) {
            if (!v5) {
                if v5 {
                    var v53 int32 = (^(v6))
                    for fi54 := 0; fi54 < 15; fi54++ {
                        v53 = v6
                        _ = fi54
                    }
                    _ = v53
                }
                var v55 int64 = 0
                if (("" >= "\n\t") || (1e10 != 0.1)) {
                    v56 := GenS3{F0: ((v6 ^ v6) > 416)}
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + "")
                    }
                    _ = v56
                } else {
                    var v57 float64 = -2.25
                    _ = len(v4)
                    v58 := GenS6{F0: (v6 ^ v6), F1: false, F2: v50}
                    _ = v57
                    _ = v58
                }
                _ = v55
            } else {
                fmt.Println(fi52)
                if v5 {
                    var v59 int64 = -663
                    v6 = 950
                    _ = v59
                } else {
                    var v60 int16 = -981
                    _ = v60
                }
                v6 = v6
            }
            var v61 float32 = (1e10 * 1.5)
            _ = v61
        } else {
            if (int16(fi52) == v15) {
                for fi62 := 0; fi62 < 20; fi62++ {
                    var v63 int32 = -1
                    if ((false && v5) && ((v15 & 32767) > v15)) {
                        v5 = (!v5)
                        var v64 uint = 65535
                        _ = v64
                    }
                    v2 = append(v2, "hello")
                    v63 = v63
                    _ = fi62
                    _ = v63
                }
                for fi65 := 0; fi65 < 20; fi65++ {
                    _ = len(v4)
                    switch "" {
                    case "hello":
                        fmt.Println(fi65)
                    }
                    _ = fi65
                }
                switch fi52 {
                case 2147483647:
                    if v5 {
                        v66 := GenS6{F0: int32(v3), F1: v5, F2: int(v1)}
                        _ = v66
                    } else {
                        v1 -= (v1 | 1)
                    }
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                var v67 bool = (!v5)
                for fi68 := 0; fi68 < 17; fi68++ {
                    fmt.Println(v67, v3)
                    v2 = append(v2, "")
                    var v69 int32 = (v6 >> ((65535) & 63))
                    _ = fi68
                    _ = v69
                }
                _ = v67
            }
            var v70 uint = 148
            _ = v70
        }
        for fi71 := 0; fi71 < 11; fi71++ {
            var v72 uint64 = (v1 & v1)
            if v5 {
                delete(v4, 860)
                var v73 int64 = -336
                _ = v73
            }
            _ = fi71
            _ = v72
        }
        if false {
            v5 = v5
            v1 = (^(v1))
            var v74 uint = 18446744073709551615
            _ = v74
        } else {
            if (0 != 0) {
                if false {
                    v5 = (1 < 0)
                    if v5 {
                        var v75 uint8 = 1
                        v15 = v15
                        _ = v75
                    } else {
                        var v76 uint32 = 42
                        switch (("\n\t" + (("foo bar" + "\n\t") + "hello")) + "foo bar") {
                        case "foo bar":
                            var v77 string = ""
                            _ = v77
                        case "":
                            v5 = (true || ((v6 % (((v6 / (((v6 | v6)) | 1))) | 1)) > v6))
                        }
                        v1 = (v1 | 42)
                        _ = v76
                    }
                } else {
                    v5 = true
                    if (!v5) {
                        v6 = (v6 ^ v6)
                    }
                    v5 = v5
                }
                var v78 int16 = v15
                _ = v78
            } else {
                for fi79 := 0; fi79 < 15; fi79++ {
                    if true {
                        var v80 bool = v5
                        v50 = (fi79 & 0)
                        var v81 uint32 = 1
                        _ = v80
                        _ = v81
                    } else {
                        v5 = v5
                    }
                    _ = fi79
                }
                var v82 int64 = int64(v6)
                switch int(v82) {
                case 127:
                    if (!(65535 != 1)) {
                        if v5 {
                            var v83 float64 = 3.14159
                            var v84 uint32 = uint32(v1)
                            v50 += (v50 * (fi52 / ((-2147483648) | 1)))
                            _ = v83
                            _ = v84
                        }
                    }
                }
                _ = v82
            }
        }
        _ = fi52
    }
    var v85 int16 = (^(v15))
    fmt.Println(v15)
    var v86 int = (v50 % ((277) | 1))
    for fi87 := 0; fi87 < 12; fi87++ {
        var v88 int8 = -83
        v89 := GenS7{F0: 42, F1: "\n\t", F2: (v88 | -1)}
        _ = fi87
        _ = v88
        _ = v89
    }
    v90 := GenS8{F0: -471, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v85
    _ = v86
    _ = v90
}
