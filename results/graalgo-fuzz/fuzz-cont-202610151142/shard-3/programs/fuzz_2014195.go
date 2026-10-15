package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint8; F1 uint8; F2 string }
type GenS9 struct { F0 uint64 }

func main() {
    var v1 bool = ((-1 == 32767) || (((-32768 > 1000) && (-712 < 807)) && (!false)))
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 757
    var v4 int64 = 430
    v4 -= (v4 % ((v4) | 1))
    if (!v1) {
        if ((!(false && (!(!(-291 >= -2147483648))))) || false) {
            v5 := GenS1{F0: "", F1: 98, F2: 42}
            _ = v5
        } else {
            v4 += (-(v4))
            v6 := GenS1{F0: ("" + "a"), F1: -96, F2: 0}
            _ = v6
        }
        switch -705 {
        case 341:
            switch 32768 {
            case -572:
                v1 = v1
            case 4294967295:
                v4 = v4
            case -1:
                for fi7 := 0; fi7 < 3; fi7++ {
                    var v8 uint16 = 2
                    var v9 int32 = 2147483647
                    v10 := GenS2{F0: 0, F1: v9}
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
        }
        switch 9223372036854775807 {
        case 128:
            var v11 int64 = 42
            _ = v11
        case 127:
            var v12 float64 = (0.0 * ((0.1 + 3.14159) - 0.0))
            _ = v12
        default:
            var v13 int8 = 100
            _ = v13
        }
    } else {
        v3 = (^(v3))
        var v14 int = 1
        if ((v14 ^ (v14 | (v14 % ((v14) | 1)))) != (v14 ^ (v14 + (v14 | (v14 % ((int(v3)) | 1)))))) {
            var v15 uint8 = 0
            var v16 bool = (v1 || v1)
            _ = v15
            _ = v16
        } else {
            v17 := GenS3{F0: 42}
            _ = v17
        }
        _ = v14
    }
    delete(v2, (v3 + (^(v3))))
    var v18 float64 = -2.25
    v19 := GenS1{F0: "a", F1: -1, F2: 42}
    var v20 bool = false
    if v1 {
        var v21 int = 366
        _ = v21
    } else {
        var v22 int8 = 42
        var v23 uint8 = 128
        var v24 bool = (-2.25 != (3.14159 * 1.5))
        _ = v22
        _ = v23
        _ = v24
    }
    v4 = v4
    v4 += 0
    v25 := GenS3{F0: 110}
    if true {
        if (!v20) {
            if ((((!v1) && v1) && v20) || (1 >= 100)) {
                var v26 uint32 = v3
                v18 -= v18
                var v27 uint8 = 0
                _ = v26
                _ = v27
            }
            v4 -= v4
            v20 = (32767 > 42)
        } else {
            switch -2147483648 {
            case 0:
                if v1 {
                    v3 += 845
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v3 = v3
                        var v29 bool = (v18 != v18)
                        if ((v20 || v1) && (false || true)) {
                            fmt.Println(v3, v4)
                            v30 := GenS2{F0: 1, F1: 2147483647}
                            v1 = (!v29)
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            case 42:
                v31 := v2[v3]
                _ = v31
            default:
                v32 := GenS3{F0: 128}
                _ = v32
            }
        }
        v3 += (v3 + v3)
    }
    if v20 {
        var v33 uint = 0
        switch (("hello" + "") + "foo bar") {
        case "\n\t":
            if v1 {
                v34 := GenS1{F0: "\n\t", F1: 1, F2: 2}
                _ = v34
            } else {
                for fi35 := 0; fi35 < 3; fi35++ {
                    for fi36 := 0; fi36 < 4; fi36++ {
                        if (!v1) {
                            var v37 int16 = 1000
                            v38 := GenS2{F0: 0, F1: -2147483648}
                            v39 := []int{}
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v33, fi36, fi36)
                            v40 := []int{32767, (fi36 << (((^(v33))) & 63)), (fi36 + (^(fi36)))}
                            _ = v40
                        }
                        v20 = v1
                        _ = fi36
                    }
                    _ = fi35
                }
                var v41 uint = (v33 << ((v33) & 63))
                var v42 int8 = 0
                _ = v41
                _ = v42
            }
        case "a":
            v33 += (v33 >> ((v33) & 63))
        case "foo bar":
            v43 := GenS3{F0: 42}
            _ = v43
        }
        _ = v33
    } else {
        v4 = (v4 % ((v4) | 1))
        for fi44 := 0; fi44 < 12; fi44++ {
            var v45 uint32 = v3
            v45 -= v3
            var v46 uint32 = (v45 % (((v3 - (v3 - v45))) | 1))
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    v47 := v2[(v3 % (((v3 - (v3 * 417))) | 1))]
    switch 42 {
    case -9223372036854775808:
        for fi48 := 0; fi48 < 12; fi48++ {
            for fi49 := 0; fi49 < 11; fi49++ {
                fmt.Println(v47, v18)
                _ = fi49
            }
            var v50 int16 = -32768
            if (!v20) {
                var v51 string = ("foo bar" + "\n\t")
                v50 = -941
                switch "" {
                case "":
                    if (true || (uint16(v47) != 32767)) {
                        fmt.Println(v18, v51, v4)
                        v20 = (false || (2147483647 >= v3))
                        var v52 int8 = -128
                        _ = v52
                    }
                }
                _ = v51
            } else {
                var v53 uint = uint(v3)
                for fi54 := 0; fi54 < 2; fi54++ {
                    fmt.Println(v18, v50)
                    var v55 float64 = float64(float32(v18))
                    switch ((("foo bar" + ("a" + (("a" + "\n\t") + "a"))) + "") + "") {
                    case "":
                        v47 = (v47 & 4294967295)
                        switch -776 {
                        case -1:
                            var v56 int32 = 0
                            _ = v56
                        case 385:
                            switch "" {
                            case "foo bar":
                                v57 := GenS4{F0: 1000, F1: 753}
                                _ = v57
                            case "hello":
                                v2[4294967295] = v3
                            case "\n\t":
                                if false {
                                    switch -1 {
                                    case 680:
                                        delete(v2, 877)
                                    case 32768:
                                        v47 = (v3 << ((65535) & 63))
                                    }
                                    v20 = (false || true)
                                } else {
                                    v58 := []int{}
                                    v3 = (v3 >> ((v53) & 63))
                                    switch fi54 {
                                    case -1:
                                        v59 := GenS4{F0: v50, F1: 2}
                                        _ = v59
                                    case 128:
                                        v18 += (-2.25 / v18)
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            var v60 uint16 = 255
                                            _ = v60
                                        case "":
                                            v61 := GenS5{F0: 437, F1: 1}
                                            _ = v61
                                        case "foo bar":
                                            v50 += (v50 / ((v50) | 1))
                                        default:
                                            v62 := GenS6{F0: false}
                                            _ = v62
                                        }
                                    }
                                    _ = v58
                                }
                            }
                        case 42:
                            v63 := GenS7{F0: 596, F1: (-(v4))}
                            _ = v63
                        }
                        v64 := GenS5{F0: 42, F1: 18446744073709551615}
                        _ = v64
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v53
            }
            _ = fi48
            _ = v50
        }
    case 0:
        v4 += (v4 / ((v4) | 1))
    default:
        fmt.Println(v18, v18)
    }
    if v20 {
        _ = len(v2)
        var v65 int8 = -1
        switch "" {
        case "a":
            switch -834 {
            case -372:
                for fi66 := 0; fi66 < 13; fi66++ {
                    if (-2147483648 >= 42) {
                        v67 := []int{int(v4), (fi66 << ((54) & 63))}
                        v3 = 269
                        fmt.Println(v3, v1, v20)
                        _ = v67
                    } else {
                        v1 = (true || v1)
                    }
                    _ = fi66
                }
            case 2147483647:
                var v68 float32 = 1e10
                _ = v68
            case 32767:
                var v69 uint32 = (v47 / ((614) | 1))
                _ = v69
            default:
                var v70 uint8 = 255
                _ = v70
            }
        case "":
            v71 := []int{0, 4294967295, -9223372036854775808, 32768}
            _ = v71
        case "foo bar":
            v2[733] = (v3 << ((688) & 63))
        default:
            v47 = v3
        }
        _ = v65
    } else {
        var v72 float32 = (3.14159 + 1e10)
        switch 2147483647 {
        case 32768:
            v73 := GenS8{F0: 1, F1: 137, F2: ("foo bar" + ("a" + "\n\t"))}
            _ = v73
        case 183:
            fmt.Println(v20)
        default:
            for fi74 := 0; fi74 < 5; fi74++ {
                if false {
                    if true {
                        v75 := []int{32768, fi74, (fi74 & fi74)}
                        _ = v75
                    } else {
                        for fi76 := 0; fi76 < 4; fi76++ {
                            var v77 int8 = -41
                            v78 := []int{fi74}
                            _ = fi76
                            _ = v77
                            _ = v78
                        }
                    }
                    v18 = float64(v72)
                }
                var v79 uint16 = 733
                if false {
                    v80 := GenS2{F0: 65535, F1: -1}
                    _ = v80
                } else {
                    v18 = v18
                    _ = len(v2)
                    for fi81 := 0; fi81 < 17; fi81++ {
                        var v82 float32 = float32(0.1)
                        v83 := GenS9{F0: 42}
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                }
                _ = fi74
                _ = v79
            }
        }
        _ = v72
    }
    var v84 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v84
}
