package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int32; F1 uint16 }
func genFn0() uint8 {
    v6 := []int{9223372036854775807, -2147483648}
    for fi7 := 0; fi7 < 10; fi7++ {
        _ = len(v6)
        var v8 int8 = 127
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: -2147483648, F1: 2}
        if (!true) {
            fmt.Println(fi9)
            var v11 int16 = 32767
            for fi12 := 0; fi12 < 19; fi12++ {
                v6 = append(v6, fi9)
                var v13 float64 = 1e10
                v14 := GenS2{F0: (v11 << ((671) & 63)), F1: 1}
                v6 = append(v6, (-(fi12)))
                v6 = append(v6, (-(fi12)))
                var v15 string = "\n\t"
                v11 -= -32768
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v16 float32 = 3.14159
    _ = len(v6)
    _ = v16
    _ = v6
    return 128
}

func main() {
    var v1 int32 = -774
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 89
    var v5 bool = (42 < (v1 << ((2) & 63)))
    var v17 float32 = float32((float64(-2.25) / 3.14159))
    var v18 int16 = -287
    switch -2147483648 {
    case 4294967295:
        var v19 float32 = 1.5
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float64 = 1.5
            _ = fi20
            _ = v21
        }
    case 1:
        if true {
            v5 = v5
            var v22 uint = 42
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 uint32 = 1
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        switch "a" {
        case "hello":
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS2{F0: 42, F1: 166}
                if v5 {
                    if ((^(v4)) == (^(v4))) {
                        v28 := GenS3{F0: 1, F1: 42, F2: v4}
                        var v29 uint = 18446744073709551615
                        _ = v28
                        _ = v29
                    }
                    v18 = v18
                    v18 += v18
                    var v30 uint16 = 42
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
        case "":
            v31 := v2["foo bar"]
            _ = v31
        }
        if v5 {
            fmt.Println(v18)
            v4 = (v4 >> ((404) & 63))
        }
        _ = fi25
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS4{F0: 0, F1: (v1 << ((42) & 63)), F2: 927}
        if (!true) {
            var v34 int64 = -677
            _ = v34
        } else {
            v4 = (v4 >> ((42) & 63))
        }
        _ = fi32
        _ = v33
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 int = (fi35 % (((-(fi35))) | 1))
        _ = fi35
        _ = v36
    }
    _ = len(v3)
    for fi37 := 0; fi37 < 10; fi37++ {
        delete(v3, 2147483647)
        _ = fi37
    }
    if ((v1 << ((1) & 63)) < v1) {
        switch -349 {
        case 0:
            fmt.Println(v4, v17)
        case 253:
            v5 = v5
        case -541:
            var v38 uint16 = 255
            _ = v38
        default:
            switch "" {
            case "a":
                v39 := GenS4{F0: 955, F1: (-(v1)), F2: 949}
                _ = v39
            case "":
                switch -9223372036854775808 {
                case 32768:
                    fmt.Println(v18, v4)
                case 127:
                    v40 := []int{-9223372036854775808, -518, 128, 127}
                    _ = v40
                case -2147483648:
                    v17 += float32(0.0)
                default:
                    var v41 int8 = 100
                    _ = v41
                }
            }
        }
        v1 = int32(v18)
        v42 := []int{int(v4), 258, 127, -672}
        _ = v42
    } else {
        v43 := GenS2{F0: (v18 & v18), F1: 65535}
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        switch ("\n\t" + "a") {
        case "\n\t":
            if (-2.25 > 1.5) {
                v18 += (v18 << ((24) & 63))
                if false {
                    switch "a" {
                    case "":
                        v45 := []int{}
                        _ = v45
                    case "hello":
                        v46 := GenS1{F0: v1, F1: 42}
                        _ = v46
                    case "a":
                        v17 = -2.25
                    }
                    v1 += (v1 & -624)
                    v17 -= v17
                } else {
                    v47 := GenS5{F0: -128, F1: 1}
                    _ = v47
                }
                v17 = (float32(1e10) - (v17 / float32(-2.25)))
            }
        case "foo bar":
            v5 = true
        default:
            fmt.Println(v18)
        }
        var v48 string = "\n\t"
        v49 := GenS3{F0: 32767, F1: 1, F2: v4}
        _ = fi44
        _ = v48
        _ = v49
    }
    if false {
        v50 := []int{127, -742, 127}
        if len(v50) > 0 {
            v50[0] = 32768
        }
        _ = v50
    } else {
        v51 := []int{-136, 42, 859, 733}
        _ = v51
    }
    var v52 uint16 = 42
    fmt.Println(v5, v17)
    var v53 uint16 = v52
    var v54 uint16 = v53
    v18 -= -1
    if true {
        if (610 <= 2147483648) {
            v1 = v1
            if true {
                switch 1 {
                case -9223372036854775808:
                    var v55 int32 = v1
                    _ = v55
                case 9223372036854775807:
                    var v56 int64 = 320
                    _ = v56
                case 4294967295:
                    v57 := GenS1{F0: v1, F1: 18446744073709551615}
                    _ = v57
                }
            }
        } else {
            v52 = 2
        }
    }
    if ((float32(1e10) * v17) < v17) {
        if v5 {
            fmt.Println(v53)
            for fi58 := 0; fi58 < 20; fi58++ {
                if v5 {
                    v54 = v53
                } else {
                    switch "a" {
                    case "a":
                        v53 = v52
                    case "\n\t":
                        if false {
                            v59 := GenS5{F0: 127, F1: -1}
                            if (!v5) {
                                var v60 string = ""
                                _ = v60
                            } else {
                                v61 := v2[""]
                                v62 := GenS3{F0: v53, F1: uint(v53), F2: (v4 & v4)}
                                _ = v61
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 7; fi63++ {
                                var v64 uint8 = genFn0()
                                var v65 int = fi63
                                var v66 string = "hello"
                                _ = fi63
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                            _ = v59
                        }
                    }
                    v18 += v18
                    v3[2] = fi58
                }
                if v5 {
                    var v67 int32 = (v1 - (v1 >> ((1) & 63)))
                    if v5 {
                        delete(v2, (("a" + "foo bar") + "\n\t"))
                        _ = len(v3)
                        for fi68 := 0; fi68 < 10; fi68++ {
                            v18 += (v18 | v18)
                            var v69 string = "hello"
                            v54 = v54
                            v52 = (v52 * (v53 & 902))
                            _ = fi68
                            _ = v69
                        }
                    } else {
                        var v70 float64 = 0.1
                        v71 := GenS2{F0: 384, F1: 878}
                        v72 := GenS6{F0: (v67 % (((v1 + 42)) | 1)), F1: v52}
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    switch "foo bar" {
                    case "a":
                        var v73 int16 = v18
                        _ = v73
                    case "":
                        switch (("a" + "hello") + "hello") {
                        case "a":
                            var v74 int8 = 127
                            _ = v74
                        default:
                            switch -1 {
                            case -269:
                                v5 = v5
                            default:
                                var v75 bool = true
                                _ = v75
                            }
                        }
                    default:
                        v17 = (float32(float64((1.5 / 1.5))) / v17)
                    }
                    _ = v67
                }
                var v76 uint32 = 42
                _ = fi58
                _ = v76
            }
        } else {
            var v77 string = "foo bar"
            _ = v77
        }
        v4 -= (v4 + (v4 + 0))
    } else {
        switch "foo bar" {
        case "hello":
            var v78 uint16 = v52
            _ = v78
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
