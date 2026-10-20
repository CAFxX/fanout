package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int; F2 uint16 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]int)
    var v3 uint8 = 1
    var v4 string = "foo bar"
    var v5 uint32 = 2147483648
    var v6 int32 = 42
    var v7 uint16 = 0
    var v8 uint32 = 42
    var v9 int16 = -447
    v5 = (v8 % ((v5) | 1))
    v10 := GenS1{F0: v9, F1: (v3 << ((65535) & 63))}
    if false {
        v11 := []int{(v1 << ((1) & 63)), v1}
        v4 = v4
        switch v1 {
        case 0:
            v6 += v6
        }
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            _ = len(v2)
            if true {
                if (!(255 <= 65535)) {
                    var v13 string = v4
                    switch "" {
                    case "":
                        switch (fi12 % ((int(v6)) | 1)) {
                        case -508:
                            v14 := GenS2{F0: (v1 & (fi12 >> ((uint(v7)) & 63))), F1: v9, F2: v4}
                            _ = v14
                        }
                    case "foo bar":
                        v9 += 0
                    case "hello":
                        var v15 int16 = v9
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 13; fi16++ {
                            switch v13 {
                            case "\n\t":
                                fmt.Println(v7)
                            case "foo bar":
                                v8 = (^(v8))
                            case "":
                                v5 = v5
                            default:
                                if true {
                                    v17 := v2[fi16]
                                    if true {
                                        var v18 string = ""
                                        _ = v18
                                    } else {
                                        var v19 int64 = 0
                                        switch v17 {
                                        case -9223372036854775808:
                                            if (18446744073709551615 <= 1) {
                                                v4 = "hello"
                                            } else {
                                                v20 := []int{-1, (fi12 | v1), int(v9)}
                                                _ = v20
                                            }
                                        case -426:
                                            switch -548 {
                                            case 42:
                                                v13 = (v13 + "\n\t")
                                            case -136:
                                                v21 := GenS1{F0: (v9 * (v9 - (v9 & (v9 % (((v9 - v9)) | 1))))), F1: (^(v3))}
                                                _ = v21
                                            }
                                        case 127:
                                            v22 := GenS3{F0: v7, F1: (fi12 >> ((18446744073709551615) & 63))}
                                            _ = v22
                                        }
                                        var v23 string = v4
                                        _ = v19
                                        _ = v23
                                    }
                                    v5 = (^(v5))
                                    v24 := GenS2{F0: v1, F1: v9, F2: ""}
                                    switch (v1 * 32768) {
                                    case 262:
                                        v25 := GenS3{F0: 906, F1: (fi16 << ((1) & 63))}
                                        _ = v25
                                    }
                                    _ = v17
                                    _ = v24
                                }
                            }
                            var v26 uint8 = (v3 << ((2) & 63))
                            _ = fi16
                            _ = v26
                        }
                    }
                    var v27 int = (fi12 << ((530) & 63))
                    _ = v13
                    _ = v27
                } else {
                    var v28 uint32 = v8
                    if true {
                        var v29 uint16 = (v7 ^ 2)
                        fmt.Println(v7, v5)
                        var v30 float64 = 3.14159
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS4{F0: 4294967295, F1: v1}
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v33 := []int{(v1 + int(v9))}
                            v4 = ("" + "a")
                            var v34 float32 = float32((0.0 / 1e10))
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        switch "" {
                        case "a":
                            for fi35 := 0; fi35 < 4; fi35++ {
                                delete(v2, (fi12 % ((v1) | 1)))
                                v7 = (^(v7))
                                if (false || true) {
                                    v9 = -567
                                    v5 = v8
                                    switch (fi12 - (v1 / (((^(fi35))) | 1))) {
                                    case 127:
                                        var v36 int32 = v6
                                        _ = v36
                                    default:
                                        v37 := []int{(^(v1))}
                                        _ = v37
                                    }
                                    v5 += (^(v28))
                                    fmt.Println(v1, fi12, v4)
                                } else {
                                    v38 := GenS3{F0: v7, F1: (v1 << ((42) & 63))}
                                    switch int(v7) {
                                    case -425:
                                        v5 = 2147483648
                                    case -169:
                                        var v39 uint32 = 2147483647
                                        _ = v39
                                    case 1:
                                        _ = len(v2)
                                    }
                                    _ = v38
                                }
                                v7 += v7
                                _ = fi35
                            }
                        case "\n\t":
                            v40 := v2[int(v7)]
                            _ = v40
                        default:
                            if true {
                                v7 = (v7 << ((255) & 63))
                                for fi41 := 0; fi41 < 15; fi41++ {
                                    if true {
                                        v42 := v2[(v1 | v1)]
                                        v2[v1] = (fi12 / (((v42 - v42)) | 1))
                                        v43 := []int{(fi41 * (^(v1))), fi12, -9223372036854775808}
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = fi41
                                }
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    fmt.Println(v5, fi44)
                                    var v45 int = (^(v1))
                                    var v46 int64 = 42
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                            }
                        }
                        _ = v31
                    }
                    _ = v28
                }
                for fi47 := 0; fi47 < 16; fi47++ {
                    var v48 bool = false
                    v49 := GenS5{F0: uint8(v5), F1: (v8 - 0)}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                if (false && true) {
                    v5 -= v5
                    v50 := GenS1{F0: (v9 << ((65535) & 63)), F1: (v3 & (v3 & (v3 ^ 2)))}
                    _ = v50
                }
            }
            _ = fi12
        }
    }
    if false {
        v51 := GenS6{F0: 218, F1: 2, F2: v9}
        switch 289 {
        case 569:
            for fi52 := 0; fi52 < 13; fi52++ {
                for fi53 := 0; fi53 < 11; fi53++ {
                    v54 := GenS6{F0: 186, F1: 0, F2: 42}
                    if (((0.0 / -2.25) <= 1.5) && true) {
                        var v55 bool = (419 < 4294967295)
                        var v56 uint16 = v7
                        _ = v55
                        _ = v56
                    }
                    _ = fi53
                    _ = v54
                }
                var v57 int16 = -32768
                fmt.Println(v8, v9, v57)
                _ = fi52
                _ = v57
            }
        }
        if false {
            switch 1 {
            case 127:
                switch (v1 + v1) {
                case -9223372036854775808:
                    switch (v1 + v1) {
                    case 0:
                        v58 := []int{}
                        _ = v58
                    case -791:
                        if (642 >= 18446744073709551615) {
                            switch v4 {
                            case "":
                                if true {
                                    var v59 int = 42
                                    if false {
                                        for fi60 := 0; fi60 < 16; fi60++ {
                                            switch (fi60 * v59) {
                                            case 543:
                                                v1 = 2147483647
                                            case -2147483648:
                                                switch (v59 & fi60) {
                                                case 127:
                                                    for fi61 := 0; fi61 < 19; fi61++ {
                                                        v1 = fi60
                                                        var v62 int8 = 39
                                                        v63 := GenS1{F0: int16(v8), F1: v3}
                                                        _ = len(v2)
                                                        v3 += (v3 ^ v3)
                                                        v64 := []int{v59, (fi61 % ((2147483647) | 1)), v1, (v59 * (^(fi60)))}
                                                        var v65 string = v4
                                                        if true {
                                                            v66 := v2[v1]
                                                            var v67 uint16 = (v7 - uint16(v3))
                                                            switch fi60 {
                                                            case 128:
                                                                var v68 uint32 = v5
                                                                _ = v68
                                                            case 638:
                                                                if ((v62 * -94) < (^(v62))) {
                                                                    fmt.Println(v9, v9, fi61)
                                                                    var v69 uint64 = 669
                                                                    fmt.Println(v59, v1, v6)
                                                                    _ = v69
                                                                } else {
                                                                    var v70 int64 = 0
                                                                    v71 := GenS3{F0: uint16(v8), F1: (v59 | -9223372036854775808)}
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            case -49:
                                                                var v72 bool = (3.14159 == float32(-2.25))
                                                                _ = v72
                                                            }
                                                            _ = v66
                                                            _ = v67
                                                        } else {
                                                            v8 -= v8
                                                            delete(v2, int(v62))
                                                            v65 = "hello"
                                                        }
                                                        _ = fi61
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                case -240:
                                                    v73 := v2[v1]
                                                    _ = v73
                                                case 9223372036854775807:
                                                    for fi74 := 0; fi74 < 15; fi74++ {
                                                        v4 = v4
                                                        _ = fi74
                                                    }
                                                default:
                                                    v75 := []int{-1, -419}
                                                    _ = v75
                                                }
                                            default:
                                                fmt.Println(v6, v1)
                                            }
                                            v4 = (v4 + v4)
                                            _ = fi60
                                        }
                                        v9 = 1000
                                        for fi76 := 0; fi76 < 19; fi76++ {
                                            switch fi76 {
                                            case 453:
                                                v77 := GenS6{F0: 42, F1: 2, F2: v9}
                                                _ = v77
                                            case 9223372036854775807:
                                                for fi78 := 0; fi78 < 18; fi78++ {
                                                    var v79 uint64 = 317
                                                    v8 += v8
                                                    _ = fi78
                                                    _ = v79
                                                }
                                            default:
                                                v80 := GenS7{F0: ((603 <= -9223372036854775808) && (1e10 == 3.14159))}
                                                _ = v80
                                            }
                                            v1 = (v1 / ((383) | 1))
                                            _ = fi76
                                        }
                                    } else {
                                        var v81 uint8 = 241
                                        _ = v81
                                    }
                                    _ = v59
                                } else {
                                    v82 := GenS8{F0: 128, F1: v1, F2: 1}
                                    _ = v82
                                }
                            default:
                                var v83 int32 = v6
                                _ = v83
                            }
                            var v84 float32 = -2.25
                            _ = v84
                        }
                    default:
                        v85 := GenS9{F0: -447, F1: (v3 ^ (^(v3)))}
                        _ = v85
                    }
                default:
                    var v86 int16 = 755
                    _ = v86
                }
            }
        }
        _ = v51
    }
    var v87 int = v1
    v88 := GenS5{F0: (v3 & (v3 >> ((741) & 63))), F1: v5}
    if true {
        var v89 string = (("foo bar" + ("foo bar" + v4)) + v4)
        for fi90 := 0; fi90 < 13; fi90++ {
            for fi91 := 0; fi91 < 10; fi91++ {
                v92 := []int{769, (fi91 - -1)}
                var v93 uint = 299
                _ = fi91
                _ = v92
                _ = v93
            }
            v94 := GenS10{F0: v3}
            _ = fi90
            _ = v94
        }
        _ = v89
    } else {
        var v95 int64 = 635
        _ = v95
    }
    v87 += (-(v1))
    delete(v2, v87)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v87:", v87)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v87
    _ = v88
    _ = v9
}
