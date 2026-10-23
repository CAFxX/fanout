package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    var v3 uint64 = 42
    var v4 int64 = -9223372036854775808
    v5 := []int{-38}
    v6 := []int{}
    v7 := GenS1{F0: false, F1: (v4 >> ((42) & 63))}
    if (71 <= 127) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 = v2
            _ = fi8
        }
    }
    var v9 int16 = 42
    if false {
        var v10 uint8 = 42
        _ = v10
    } else {
        var v11 int64 = v4
        fmt.Println(v9)
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                if (0 == 42) {
                    fmt.Println(fi13, v9, fi12)
                    var v14 float32 = 0.0
                    _ = v14
                } else {
                    switch v2 {
                    case "":
                        var v15 string = v2
                        _ = v15
                    case "a":
                        var v16 bool = (true && (true && true))
                        _ = v16
                    }
                }
                switch v2 {
                case "foo bar":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    if (!(!("hello" <= "foo bar"))) {
                        if (!true) {
                            if true {
                                if (-744 < v9) {
                                    switch v2 {
                                    case "a":
                                        var v17 uint32 = 1
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v17
                                    case "hello":
                                        v1[429] = v2
                                    }
                                    if (!true) {
                                        var v18 float32 = 0.0
                                        if ((0.0 + (float64(v18) - float64(v18))) == 3.14159) {
                                            var v19 int8 = 42
                                            _ = v19
                                        }
                                        v18 -= (v18 + v18)
                                        _ = v18
                                    }
                                }
                            } else {
                                v20 := GenS2{F0: ("" + v2), F1: (fi12 & (fi12 + fi13))}
                                var v21 int32 = -437
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS2{F0: (((("hello" + v2) + v2) + (v2 + v2)) + ""), F1: fi13}
                            fmt.Println(v4, fi13, v4)
                            _ = v22
                        }
                        if false {
                            switch v2 {
                            case "\n\t":
                                v23 := GenS1{F0: false, F1: (v11 - v11)}
                                _ = v23
                            case "foo bar":
                                if (!false) {
                                    delete(v1, 2147483647)
                                } else {
                                    fmt.Println(fi12)
                                    var v24 uint64 = v3
                                    _ = v24
                                }
                            case "hello":
                                v3 = v3
                            default:
                                var v25 uint16 = 32767
                                _ = v25
                            }
                        } else {
                            v26 := GenS2{F0: "a", F1: 42}
                            var v27 uint = 1
                            if (int64(v9) > (v4 % ((v11) | 1))) {
                                switch (v2 + v2) {
                                case "foo bar":
                                    _ = len(v5)
                                case "":
                                    v3 = v3
                                }
                                fmt.Println(v27, v2)
                                delete(v1, 184)
                            }
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS1{F0: true, F1: v4}
                        _ = v28
                    } else {
                        fmt.Println(v3)
                    }
                }
                if (true || true) {
                    v29 := GenS1{F0: true, F1: v4}
                    _ = len(v1)
                    _ = v29
                }
                _ = fi13
            }
            _ = fi12
        }
        _ = v11
    }
    if false {
        for fi30 := 0; fi30 < 9; fi30++ {
            v31 := GenS2{F0: ("" + ("hello" + (v2 + v2))), F1: fi30}
            if false {
                for fi32 := 0; fi32 < 6; fi32++ {
                    if (false || (false || false)) {
                        switch v2 {
                        case "\n\t":
                            switch fi32 {
                            case 2147483647:
                                _ = len(v1)
                            }
                        case "hello":
                            var v33 int16 = (v9 / ((v9) | 1))
                            _ = v33
                        case "foo bar":
                            if false {
                                if true {
                                    switch 32768 {
                                    case -2147483648:
                                        fmt.Println(v3, fi30, v4)
                                    default:
                                        if false {
                                            v9 = v9
                                            v3 = v3
                                        }
                                    }
                                    var v34 int16 = (v9 % ((v9) | 1))
                                    var v35 int32 = 111
                                    _ = v34
                                    _ = v35
                                }
                                v2 = v2
                                v4 = (-(v4))
                            } else {
                                v9 = (v9 >> ((551) & 63))
                            }
                        default:
                            var v36 int64 = (v4 >> ((2) & 63))
                            _ = v36
                        }
                        _ = len(v6)
                    } else {
                        fmt.Println(v4, fi32, v2)
                    }
                    _ = fi32
                }
                v37 := GenS3{F0: 32767}
                switch (v2 + "\n\t") {
                case "foo bar":
                    var v38 float32 = ((-2.25 + 1e10) * 0.1)
                    _ = v38
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "hello":
                    var v39 int8 = 1
                    _ = v39
                }
                _ = v37
            } else {
                v40 := v1[2147483648]
                var v41 int16 = (v9 & (-(v9)))
                _ = v40
                _ = v41
            }
            _ = fi30
            _ = v31
        }
        switch 67 {
        case 783:
            switch (v2 + v2) {
            case "hello":
                var v42 bool = false
                _ = v42
            case "a":
                v3 = 0
            default:
                switch (v2 + "") {
                case "\n\t":
                    var v43 uint32 = 2
                    _ = v43
                case "hello":
                    switch "foo bar" {
                    case "\n\t":
                        for fi44 := 0; fi44 < 2; fi44++ {
                            var v45 float32 = 0.0
                            fmt.Println(v2)
                            switch v2 {
                            case "":
                                v46 := GenS1{F0: ((^(v9)) < (v9 / ((v9) | 1))), F1: (v4 + v4)}
                                _ = v46
                            case "foo bar":
                                _ = len(v1)
                            case "hello":
                                var v47 int8 = 104
                                _ = v47
                            }
                            _ = fi44
                            _ = v45
                        }
                    case "":
                        delete(v1, 2147483648)
                    default:
                        for fi48 := 0; fi48 < 11; fi48++ {
                            for fi49 := 0; fi49 < 10; fi49++ {
                                var v50 uint64 = v3
                                if false {
                                    var v51 int = (fi48 - fi49)
                                    _ = len(v5)
                                    v52 := GenS2{F0: v2, F1: (^(fi49))}
                                    _ = v51
                                    _ = v52
                                } else {
                                    switch (fi48 >> ((74) & 63)) {
                                    case 127:
                                        if (!(fi48 != (fi49 | 2147483647))) {
                                            if (false || false) {
                                                v50 = 2
                                            }
                                        } else {
                                            v53 := GenS1{F0: false, F1: (v4 ^ v4)}
                                            v54 := GenS3{F0: (v9 | v9)}
                                            v3 = (v50 << ((415) & 63))
                                            _ = v53
                                            _ = v54
                                        }
                                    case 0:
                                        v55 := GenS1{F0: false, F1: v4}
                                        _ = v55
                                    }
                                    var v56 string = (v2 + v2)
                                    _ = v56
                                }
                                if (37 == 52) {
                                    v57 := GenS3{F0: int16(fi49)}
                                    fmt.Println(fi48, fi49)
                                    var v58 uint8 = 255
                                    _ = v57
                                    _ = v58
                                }
                                var v59 int64 = (v4 | v4)
                                _ = len(v1)
                                if true {
                                    v9 = -431
                                    if false {
                                        v3 -= (v50 >> ((721) & 63))
                                        if false {
                                            v50 = (v50 / (((v50 - v3)) | 1))
                                            var v60 bool = true
                                            _ = v60
                                        }
                                    }
                                    var v61 int64 = (-(v4))
                                    if (!false) {
                                        fmt.Println(fi49, v3, fi49)
                                        v59 = (-(v4))
                                        v62 := GenS1{F0: true, F1: -1}
                                        _ = v62
                                    } else {
                                        v3 = (v3 | v3)
                                        if (761 < 18446744073709551615) {
                                            v63 := GenS1{F0: (true && true), F1: v4}
                                            var v64 bool = false
                                            v65 := GenS2{F0: v2, F1: 128}
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            fmt.Println(v4)
                                            var v66 uint32 = 2147483647
                                            _ = v66
                                        }
                                    }
                                    _ = v61
                                } else {
                                    v5 = append(v5, (fi49 + (-(fi48))))
                                    var v67 int8 = 127
                                    var v68 int8 = (v67 << ((65535) & 63))
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi49
                                _ = v50
                                _ = v59
                            }
                            _ = fi48
                        }
                    }
                case "":
                    var v69 int32 = 0
                    _ = v69
                }
            }
        }
        switch 1 {
        case 0:
            v2 = v2
        default:
            if true {
                for fi70 := 0; fi70 < 8; fi70++ {
                    if true {
                        v9 -= (v9 & 1000)
                        if true {
                            _ = len(v5)
                        }
                    }
                    var v71 int = -9223372036854775808
                    v72 := GenS4{F0: true, F1: (v3 * 42)}
                    _ = fi70
                    _ = v71
                    _ = v72
                }
            } else {
                var v73 float32 = -2.25
                for fi74 := 0; fi74 < 2; fi74++ {
                    switch (v2 + "") {
                    case "\n\t":
                        var v75 float64 = (0.1 - 0.1)
                        _ = v75
                    case "hello":
                        if true {
                            v76 := GenS2{F0: v2, F1: -550}
                            _ = v76
                        } else {
                            delete(v1, 42)
                            switch (fi74 << ((255) & 63)) {
                            case 32768:
                                var v77 uint32 = 909
                                _ = v77
                            case -486:
                                if true {
                                    var v78 uint64 = v3
                                    v3 -= v78
                                    v2 = v2
                                    _ = v78
                                }
                            }
                        }
                    }
                    _ = fi74
                }
                _ = v73
            }
        }
    }
    var v79 int64 = (v4 << ((42) & 63))
    v80 := v1[478]
    v81 := GenS2{F0: "", F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v79
    _ = v80
    _ = v81
    _ = v9
}
