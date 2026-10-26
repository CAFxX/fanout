package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS6 struct { F0 bool; F1 int; F2 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8 }

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    var v3 uint32 = 2
    v4 := make(map[string]string)
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: (^(v3)), F1: fi6}
        _ = fi6
        _ = v7
    }
    v3 -= 2147483647
    if false {
        switch v1 {
        case "a":
            switch ("" + "hello") {
            case "a":
                if true {
                    if (-647 >= -808) {
                        var v8 int64 = -1
                        _ = v8
                    }
                } else {
                    v3 -= (v3 / ((v3) | 1))
                    var v9 float64 = float64((3.14159 / (3.14159 * (0.1 * (3.14159 * ((float32(-2.25) - float32(float64((-2.25 / (float32(float64(3.14159)) - 1.5))))) - 0.0))))))
                    if false {
                        for fi10 := 0; fi10 < 8; fi10++ {
                            fmt.Println(v1, v3)
                            v11 := GenS1{F0: v3, F1: fi10}
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        v3 = (v3 + (^(v3)))
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v3 -= (v3 & v3)
                            v13 := GenS2{F0: 0, F1: 0}
                            _ = fi12
                            _ = v13
                        }
                        v14 := GenS2{F0: -9, F1: 18446744073709551615}
                        _ = v14
                    }
                    _ = v9
                }
            case "hello":
                v5 += v5
            case "":
                v15 := GenS3{F0: 255, F1: 255}
                _ = v15
            }
        case "foo bar":
            if (true && ((!true) && false)) {
                v16 := []int{}
                v3 = v3
                _ = v16
            }
        }
        switch "hello" {
        case "\n\t":
            var v17 uint = 986
            _ = v17
        case "":
            var v18 uint32 = 4294967295
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 int = fi19
                v21 := GenS1{F0: v3, F1: (-(v20))}
                if true {
                    v20 = 32767
                    for fi22 := 0; fi22 < 7; fi22++ {
                        if (v3 >= 42) {
                            var v23 string = (((v1 + ("hello" + v1)) + v1) + (v1 + ("foo bar" + "a")))
                            _ = v23
                        }
                        fmt.Println(v5)
                        fmt.Println(v5)
                        _ = fi22
                    }
                    var v24 int16 = 1
                    _ = v24
                } else {
                    fmt.Println(fi19, v1, fi19)
                    switch -2147483648 {
                    case -985:
                        v3 -= (^(v3))
                    case -324:
                        var v25 int = v20
                        _ = v25
                    case -905:
                        v26 := GenS4{F0: 2, F1: v1, F2: 42}
                        _ = v26
                    }
                    var v27 float32 = 1.5
                    _ = v27
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            switch ("foo bar" + v1) {
            case "":
                if (!(false || (float32(1e10) != float32(0.1)))) {
                    v28 := GenS5{F0: v5, F1: 1, F2: 24}
                    _ = v28
                }
            case "a":
                for fi29 := 0; fi29 < 11; fi29++ {
                    if true {
                        v3 += 523
                    }
                    _ = fi29
                }
            default:
                v30 := GenS4{F0: 18446744073709551615, F1: (v1 + ((("" + (v1 + v1)) + "a") + v1)), F2: -1}
                _ = v30
            }
        }
    } else {
        var v31 string = ((v1 + "hello") + v1)
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(v5, v5)
            v3 += v3
            if true {
                v33 := GenS2{F0: -569, F1: 2}
                v31 = v31
                _ = v33
            } else {
                delete(v4, (v1 + v31))
                switch v1 {
                case "a":
                    for fi34 := 0; fi34 < 2; fi34++ {
                        var v35 int8 = 78
                        if true {
                            if (!true) {
                                fmt.Println(fi32, v5, fi32)
                                v36 := GenS5{F0: 656, F1: 128, F2: (v35 % (((v35 * 100)) | 1))}
                                _ = v36
                            } else {
                                v37 := GenS3{F0: v5, F1: 147}
                                _ = v37
                            }
                            if (v35 <= v35) {
                                v38 := []int{}
                                v4[v31] = "\n\t"
                                if (1 <= 466) {
                                    v31 = ((("a" + v1) + ("hello" + ((v1 + v31) + (("hello" + v31) + "foo bar")))) + ((("\n\t" + v31) + "hello") + ("hello" + ("foo bar" + v1))))
                                } else {
                                    v1 = "foo bar"
                                }
                                v1 = "hello"
                                if (!(42 >= 0)) {
                                    switch ((v31 + v1) + v31) {
                                    case "a":
                                        _ = len(v38)
                                    default:
                                        _ = len(v38)
                                    }
                                } else {
                                    if len(v38) > 0 {
                                        v38[0] = 1
                                    }
                                }
                                _ = v38
                            } else {
                                if (0 >= 9223372036854775807) {
                                    var v39 int = fi34
                                    var v40 uint16 = (^(v5))
                                    v35 = v35
                                    v41 := GenS6{F0: (false || (!(!false))), F1: (fi34 | int(v3)), F2: v31}
                                    v42 := []int{(v39 | (fi34 / ((fi34) | 1)))}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                var v43 uint8 = 0
                                if (1 < uint(v5)) {
                                    v44 := GenS5{F0: v5, F1: v43, F2: (v35 * v35)}
                                    v3 -= 4294967295
                                    _ = v44
                                }
                                _ = v43
                            }
                        }
                        fmt.Println(v35, v31)
                        _ = fi34
                        _ = v35
                    }
                case "":
                    fmt.Println(v31, v31)
                default:
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 int64 = -9223372036854775808
                        if false {
                            v47 := GenS7{F0: (v1 + ((v1 + v31) + (v31 + "\n\t")))}
                            _ = v47
                        } else {
                            v4[v31] = v1
                        }
                        var v48 uint64 = 858
                        v46 = (v46 | (v46 / (((v46 ^ 42)) | 1)))
                        v49 := GenS1{F0: v3, F1: fi32}
                        switch (v31 + ("foo bar" + "foo bar")) {
                        case "foo bar":
                            delete(v4, v31)
                        case "hello":
                            fmt.Println(v46, fi32, fi45)
                        case "":
                            if false {
                                v50 := []int{}
                                _ = v50
                            }
                        default:
                            var v51 uint16 = (v5 % ((v5) | 1))
                            _ = v51
                        }
                        switch fi45 {
                        case -9223372036854775808:
                            v52 := GenS8{F0: 2, F1: 255}
                            _ = v52
                        case -2147483648:
                            if true {
                                if true {
                                    v53 := GenS1{F0: v3, F1: fi32}
                                    if true {
                                        var v54 uint64 = v48
                                        v3 = v3
                                        var v55 uint32 = (v3 - (v3 % (((v3 / ((v3) | 1))) | 1)))
                                        _ = v54
                                        _ = v55
                                    }
                                    v3 = 967
                                    switch (v31 + "a") {
                                    case "foo bar":
                                        var v56 float64 = 1e10
                                        _ = v56
                                    }
                                    switch 0 {
                                    case 128:
                                        v4[(((("\n\t" + v31) + v31) + ("a" + v1)) + v1)] = v31
                                    case -9223372036854775808:
                                        switch v31 {
                                        case "":
                                            var v57 float64 = -2.25
                                            _ = v57
                                        default:
                                            var v58 int32 = -2147483648
                                            _ = v58
                                        }
                                        delete(v4, ((v1 + (((("a" + (((("a" + v1) + v31) + "foo bar") + v31)) + v31) + v31) + v31)) + v1))
                                        v48 += 1
                                        switch (v31 + (v1 + ((("\n\t" + (v31 + "")) + "foo bar") + "\n\t"))) {
                                        case "foo bar":
                                            fmt.Println(fi45, v1)
                                        case "hello":
                                            v59 := GenS9{F0: 32767}
                                            _ = v59
                                        }
                                        if false {
                                            var v60 float64 = ((0.0 - ((0.0 * 0.1) + 0.1)) - 3.14159)
                                            v61 := v4[v31]
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    _ = v53
                                }
                                v62 := GenS10{F0: int8(v48), F1: 42}
                                var v63 uint8 = 121
                                _ = v62
                                _ = v63
                            } else {
                                var v64 uint64 = v48
                                v65 := GenS11{F0: int16(v64)}
                                v66 := GenS2{F0: 632, F1: (v48 | uint64(v5))}
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                        case 0:
                            _ = len(v2)
                        }
                        v67 := []int{(fi32 + 9223372036854775807)}
                        _ = fi45
                        _ = v46
                        _ = v48
                        _ = v49
                        _ = v67
                    }
                }
            }
            _ = fi32
        }
        v3 -= v3
        _ = v31
    }
    var v68 string = ((v1 + v1) + "a")
    var v69 int64 = 1
    v5 += (v5 * (v5 << ((65535) & 63)))
    v70 := []int{1, -964}
    _ = len(v70)
    fmt.Println(v5, v1)
    var v71 uint64 = 4294967295
    v72 := GenS12{F0: 0}
    var v73 uint8 = 128
    if true {
        var v74 bool = ((v3 & (v3 << ((117) & 63))) != 889)
        v1 = (v68 + v1)
        _ = v74
    }
    v5 -= v5
    fmt.Println(v3, v68, v73)
    var v75 uint = 570
    if true {
        if (false && (true && (197 >= v69))) {
            if true {
                v69 = (v69 >> ((uint(v73)) & 63))
                if (0.0 <= float32(0.1)) {
                    switch 1 {
                    case 42:
                        v1 = v1
                    default:
                        if true {
                            for fi76 := 0; fi76 < 3; fi76++ {
                                v75 -= (v75 << ((220) & 63))
                                _ = len(v2)
                                v5 = (v5 >> (((v75 ^ v75)) & 63))
                                _ = fi76
                            }
                            if ((v73 >> ((1) & 63)) == v73) {
                                v77 := GenS10{F0: 0, F1: 105}
                                fmt.Println(v1)
                                _ = v77
                            } else {
                                var v78 int = 14
                                _ = v78
                            }
                        }
                    }
                }
            } else {
                if len(v70) > 0 {
                    v70[0] = 32768
                }
                v68 = v1
                var v79 string = ((("hello" + v68) + v68) + ("\n\t" + v1))
                _ = v79
            }
            _ = len(v70)
            if true {
                v68 = (v1 + v68)
                v5 = v5
                fmt.Println(v3, v5)
            }
        } else {
            for fi80 := 0; fi80 < 19; fi80++ {
                if true {
                    v73 -= (v73 - (v73 * uint8(v5)))
                }
                _ = fi80
            }
            for fi81 := 0; fi81 < 15; fi81++ {
                for fi82 := 0; fi82 < 11; fi82++ {
                    var v83 uint64 = (v71 << (((v75 / (((v75 / (((v75 + v75)) | 1))) | 1))) & 63))
                    v84 := GenS3{F0: 2, F1: uint8(v3)}
                    v75 = uint(v73)
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                _ = fi81
            }
        }
        var v85 int64 = v69
        if (v3 > uint32(v71)) {
            if len(v70) >= 2 {
                v70 = v70[1:len(v70)]
            }
            if len(v70) > 0 {
                v70[0] = 4294967295
            }
            v5 += 42
        }
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v75
}
