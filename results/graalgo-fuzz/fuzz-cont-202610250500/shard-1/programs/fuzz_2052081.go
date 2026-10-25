package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int32; F1 int; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int64 }
func genFn0(p0 uint8, p1 int32) int64 {
    v5 := []int{4294967295, 128}
    p1 = (p1 + (p1 * (p1 % ((2147483647) | 1))))
    v6 := GenS1{F0: 4294967295}
    if (!true) {
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -935
}

func main() {
    var v1 bool = false
    var v2 uint32 = 2147483647
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    fmt.Println(v1)
    if (!(9223372036854775807 != 9223372036854775807)) {
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 string = ("foo bar" + (((("foo bar" + ("\n\t" + ("hello" + "hello"))) + "foo bar") + ("a" + "a")) + ""))
                v9 := []int{fi7, fi7, (fi7 % ((int(v4)) | 1))}
                var v10 uint32 = (^(v2))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
        var v11 float32 = 1.5
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := v3[42]
            switch v13 {
            case 1:
                v13 = 525
            case -1:
                var v14 int32 = -949
                _ = v14
            default:
                if false {
                    var v15 int16 = -32768
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        switch "a" {
                        case "a":
                            var v17 uint8 = 247
                            _ = v17
                        case "hello":
                            v11 -= v11
                            v18 := GenS1{F0: 0}
                            _ = v18
                        default:
                            fmt.Println(fi12, v4, v13)
                        }
                        _ = fi16
                    }
                }
            }
            if (v1 || v1) {
                if (100 < 127) {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 uint64 = 42
                        v20 = 42
                        _ = fi19
                        _ = v20
                    }
                } else {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v13 -= (fi21 << ((v4) & 63))
                        fmt.Println(v11)
                        switch ("" + "hello") {
                        case "hello":
                            switch "" {
                            case "foo bar":
                                v22 := GenS1{F0: 42}
                                _ = v22
                            case "\n\t":
                                v2 += (^(v2))
                            }
                        }
                        var v23 uint8 = 0
                        _ = fi21
                        _ = v23
                    }
                    if false {
                        fmt.Println(v1, v13)
                        if (404 > 32767) {
                            if true {
                                switch v13 {
                                case -1:
                                    var v24 uint64 = 2
                                    _ = v24
                                }
                                fmt.Println(v2)
                            }
                        } else {
                            if false {
                                if false {
                                    v11 -= (1e10 - (v11 * (v11 - float32(3.14159))))
                                } else {
                                    fmt.Println(v11)
                                    switch (fi12 << (((v4 - (v4 ^ 42))) & 63)) {
                                    case 42:
                                        if v1 {
                                            switch (fi12 | (fi12 >> ((v4) & 63))) {
                                            case 42:
                                                if v1 {
                                                    if (!(!(830 > uint64(v2)))) {
                                                        v11 = v11
                                                        v25 := GenS2{F0: 18446744073709551615, F1: -1}
                                                        fmt.Println(v4, v4, fi12)
                                                        _ = v25
                                                    } else {
                                                        var v26 uint64 = 1
                                                        _ = v26
                                                    }
                                                } else {
                                                    var v27 int8 = 0
                                                    var v28 int32 = 0
                                                    v27 += v27
                                                    _ = v27
                                                    _ = v28
                                                }
                                            case 1:
                                                if (255 >= v4) {
                                                    fmt.Println(v13, v2, v11)
                                                    v2 += v2
                                                    v2 -= 42
                                                } else {
                                                    v29 := []int{}
                                                    var v30 uint64 = 4294967295
                                                    _ = v29
                                                    _ = v30
                                                }
                                            default:
                                                if (497 == v2) {
                                                    v4 = (v4 + v4)
                                                } else {
                                                    fmt.Println(v4, v2)
                                                    v31 := GenS3{F0: true, F1: 2}
                                                    v2 += v2
                                                    _ = v31
                                                }
                                            }
                                            v32 := v3[(v13 | 4294967295)]
                                            var v33 string = ""
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 int32 = -1
                                            if (!false) {
                                                fmt.Println(v11)
                                            }
                                            switch ("" + (("" + ("\n\t" + "a")) + ("foo bar" + "hello"))) {
                                            case "a":
                                                switch v13 {
                                                case 42:
                                                    v35 := GenS4{F0: "foo bar", F1: true}
                                                    _ = v35
                                                case -989:
                                                    switch v13 {
                                                    case 205:
                                                        v36 := []int{fi12, v13}
                                                        _ = v36
                                                    default:
                                                        v37 := []int{fi12, (v13 * 32767), fi12, fi12}
                                                        _ = v37
                                                    }
                                                }
                                            case "":
                                                var v38 int8 = -21
                                                _ = v38
                                            }
                                            _ = v34
                                        }
                                    }
                                }
                                var v39 bool = (42 < 4294967295)
                                _ = v39
                            } else {
                                v40 := GenS4{F0: "hello", F1: v1}
                                v41 := []int{(v13 % (((v13 - 969)) | 1)), int(v4), v13}
                                var v42 uint32 = (v2 << ((255) & 63))
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            var v43 int32 = -1
                            v44 := []int{}
                            _ = v43
                            _ = v44
                        }
                        v11 += v11
                    }
                    v13 = v13
                }
                v2 = v2
                v45 := GenS3{F0: (-2147483648 <= -860), F1: 128}
                _ = v45
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v46 uint16 = 272
    if v1 {
        fmt.Println(v4, v4)
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 uint64 = 42
            _ = fi47
            _ = v48
        }
    } else {
        switch 2147483647 {
        case 128:
            v49 := GenS2{F0: 0, F1: -418}
            _ = v49
        case -539:
            v50 := GenS2{F0: 409, F1: 1}
            _ = v50
        case 9223372036854775807:
            if (int32(v46) <= 395) {
                var v51 int64 = 656
                var v52 uint = (v4 % ((1) | 1))
                var v53 uint16 = 635
                _ = v51
                _ = v52
                _ = v53
            }
        }
        for fi54 := 0; fi54 < 19; fi54++ {
            switch fi54 {
            case 0:
                if (320 <= 9223372036854775807) {
                    if (1.5 >= (0.1 * 0.0)) {
                        var v55 float32 = 0.0
                        for fi56 := 0; fi56 < 18; fi56++ {
                            if v1 {
                                v57 := GenS4{F0: "a", F1: (v1 || (uint8(v4) >= 42))}
                                v55 = v55
                                if v1 {
                                    var v58 uint = (v4 << (((^(v4))) & 63))
                                    v1 = (!v1)
                                    var v59 uint16 = (v46 - v46)
                                    v1 = (255 > 2)
                                    v60 := GenS1{F0: 4294967295}
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    switch fi56 {
                                    case -967:
                                        if v1 {
                                            v1 = v1
                                            var v61 bool = ((v4 + (^(v4))) <= 2)
                                            v62 := GenS5{F0: -352, F1: fi56, F2: 127}
                                            _ = v61
                                            _ = v62
                                        }
                                    case 128:
                                        fmt.Println(fi56, v4)
                                    default:
                                        switch 127 {
                                        case -1:
                                            if ((127 > 127) && ((2147483647 <= 1) || ("\n\t" != ""))) {
                                                v4 = (v4 << ((2) & 63))
                                                var v63 float64 = 3.14159
                                                _ = v63
                                            } else {
                                                var v64 uint32 = v2
                                                fmt.Println(v4, v2, fi56)
                                                v1 = (!false)
                                                _ = v64
                                            }
                                        case -2147483648:
                                            fmt.Println(v4, v55, v4)
                                        }
                                    }
                                }
                                _ = v57
                            }
                            _ = fi56
                        }
                        switch ("\n\t" + ("" + ((("" + "a") + "") + "\n\t"))) {
                        case "":
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v1, v2)
                            default:
                                for fi65 := 0; fi65 < 10; fi65++ {
                                    v66 := []int{fi65}
                                    v67 := GenS6{F0: 42, F1: 2, F2: 1}
                                    if (!((42 >= -387) || (-27 < -25))) {
                                        fmt.Println(v46, fi54)
                                        v4 -= v4
                                        fmt.Println(fi65)
                                    }
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                            }
                        case "a":
                            if (!v1) {
                                for fi68 := 0; fi68 < 16; fi68++ {
                                    v1 = false
                                    var v69 uint16 = (v46 * (^(v46)))
                                    _ = fi68
                                    _ = v69
                                }
                                v70 := GenS7{F0: 1}
                                var v71 uint32 = 330
                                _ = v70
                                _ = v71
                            }
                        case "hello":
                            v4 -= v4
                        default:
                            v46 += 2
                        }
                        _ = v55
                    } else {
                        v2 -= 2147483647
                    }
                    fmt.Println(v4, v46, v46)
                    fmt.Println(v2)
                }
            case -2147483648:
                v2 -= v2
            }
            v4 += (v4 << ((v4) & 63))
            v72 := GenS7{F0: 42}
            _ = fi54
            _ = v72
        }
        for fi73 := 0; fi73 < 15; fi73++ {
            v74 := GenS4{F0: "\n\t", F1: v1}
            _ = fi73
            _ = v74
        }
    }
    delete(v3, 902)
    v75 := []int{187, 9223372036854775807}
    _ = len(v3)
    var v76 float64 = 3.14159
    v2 += 4294967295
    fmt.Println(v46, v46, v4)
    v77 := GenS5{F0: 2147483647, F1: 4294967295, F2: 128}
    var v78 uint16 = 2
    if (4294967295 > 127) {
        v46 = uint16(v2)
        if (-135 == 1) {
            v4 += v4
            var v79 uint = v4
            v80 := GenS3{F0: ((v4 | (v79 >> ((v79) & 63))) >= (v4 << ((v4) & 63))), F1: 128}
            _ = v79
            _ = v80
        }
        v46 = (^(v46))
    } else {
        v81 := GenS8{F0: -9223372036854775808, F1: 1}
        var v82 string = "hello"
        v76 = 0.0
        _ = v81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
