package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint) uint16 {
    var v6 uint8 = 42
    v7 := []int{0, 1}
    var v8 int16 = 32767
    if true {
        var v9 float32 = 1e10
        _ = v9
    } else {
        v10 := GenS1{F0: -9223372036854775808, F1: 565}
        fmt.Println(v8)
        if false {
            var v11 int = -666
            v8 += (-(v8))
            _ = v11
        } else {
            v12 := GenS2{F0: 299, F1: 185, F2: p1}
            v13 := GenS3{F0: 32768}
            _ = v12
            _ = v13
        }
        _ = v10
    }
    v14 := GenS3{F0: -48}
    _ = p0
    _ = p1
    _ = v14
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []string{""}
    v2 := []string{"", "hello"}
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    var v5 uint16 = 2
    var v15 uint = 0
    if true {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch "a" {
            case "hello":
                var v17 uint64 = 938
                _ = v17
            case "a":
                switch fi16 {
                case -1:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        var v19 int = fi18
                        v15 -= (v15 + (v15 % ((v15) | 1)))
                        _ = fi18
                        _ = v19
                    }
                case 573:
                    var v20 bool = ((!true) && (true && ("foo bar" > ("hello" + "hello"))))
                    _ = v20
                }
            }
            _ = fi16
        }
        fmt.Println(v3, v3)
        for fi21 := 0; fi21 < 1; fi21++ {
            _ = len(v4)
            var v22 uint16 = (v5 * 0)
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int16 = 32767
        for fi25 := 0; fi25 < 5; fi25++ {
            switch "\n\t" {
            case "a":
                v24 = (v24 ^ v24)
            case "foo bar":
                if false {
                    v15 -= v15
                } else {
                    var v26 uint8 = 31
                    var v27 int64 = -1
                    v26 = v26
                    switch fi25 {
                    case 127:
                        var v28 int8 = 100
                        _ = v28
                    case 297:
                        fmt.Println(v3, v27)
                    default:
                        v2 = append(v2, "a")
                    }
                    fmt.Println(v3, v5)
                    _ = v26
                    _ = v27
                }
            }
            if (false || (!((^(v5)) >= (v5 - v5)))) {
                var v29 int64 = 42
                switch "" {
                case "hello":
                    v4[fi25] = 370
                case "\n\t":
                    v15 -= 65535
                case "a":
                    if ("\n\t" == "a") {
                        var v30 uint8 = 10
                        v29 = v29
                        fmt.Println(v24)
                        _ = v30
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if (!((128 != 255) && true)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                default:
                    var v31 float64 = 0.1
                    _ = v31
                }
                v32 := GenS4{F0: (v24 << ((v15) & 63)), F1: 1}
                if true {
                    var v33 uint32 = 65535
                    var v34 string = "a"
                    v15 += v15
                    _ = v33
                    _ = v34
                }
                _ = v29
                _ = v32
            } else {
                if ((!false) || true) {
                    var v35 float32 = -2.25
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = v35
                } else {
                    if false {
                        fmt.Println(fi25, v15, v24)
                        v3 -= int32(v5)
                        var v36 int64 = int64(v3)
                        var v37 uint = (v15 - v15)
                        switch "a" {
                        case "":
                            if (v37 <= (v15 ^ (v15 >> (((v37 + v37)) & 63)))) {
                                var v38 uint8 = 42
                                if (true && (v24 <= -1)) {
                                    if ((v37 / ((v37) | 1)) == (^(v15))) {
                                        var v39 float64 = (1.5 / 1.5)
                                        fmt.Println(fi25, v39, v37)
                                        _ = v39
                                    } else {
                                        v40 := GenS3{F0: (fi25 | int(v3))}
                                        switch ("hello" + "a") {
                                        case "foo bar":
                                            var v41 uint16 = v5
                                            _ = v41
                                        case "\n\t":
                                            v4[fi25] = 2
                                        }
                                        _ = v40
                                    }
                                    var v42 int32 = 2147483647
                                    _ = v42
                                } else {
                                    v37 -= (v15 + (^(v15)))
                                }
                                _ = v38
                            } else {
                                var v43 uint16 = (v5 + (v5 + v5))
                                var v44 int32 = -1
                                var v45 uint = v37
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        case "a":
                            if (!(-442 > 9223372036854775807)) {
                                if false {
                                    if false {
                                        switch fi25 {
                                        case 300:
                                            v46 := GenS4{F0: 1000, F1: 22}
                                            _ = v46
                                        case -905:
                                            var v47 int8 = -1
                                            _ = v47
                                        case 1:
                                            if (true && (false || true)) {
                                                if (genFn0(4294967295, (v37 ^ (v15 * v37))) == (v5 - (v5 | v5))) {
                                                    v48 := GenS5{F0: v24}
                                                    v49 := GenS1{F0: fi25, F1: int32(v24)}
                                                    _ = v48
                                                    _ = v49
                                                }
                                                var v50 uint64 = 632
                                                _ = v50
                                            } else {
                                                v36 -= -9223372036854775808
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        default:
                                            var v51 bool = false
                                            _ = v51
                                        }
                                    }
                                    v52 := GenS1{F0: 761, F1: v3}
                                    _ = v52
                                } else {
                                    v3 = (-(v3))
                                    var v53 string = (("a" + "\n\t") + "foo bar")
                                    switch "\n\t" {
                                    case "\n\t":
                                        v54 := GenS4{F0: (v24 | v24), F1: 906}
                                        _ = v54
                                    case "":
                                        fmt.Println(v3, v15)
                                    default:
                                        switch v53 {
                                        case "a":
                                            var v55 uint32 = 1
                                            _ = v55
                                        case "hello":
                                            if (v5 <= 176) {
                                                if (genFn0(42, v37) < (v5 >> ((v15) & 63))) {
                                                    switch 2147483647 {
                                                    case 591:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case 1:
                                                        fmt.Println(v5)
                                                    case 4294967295:
                                                        v53 = v53
                                                    default:
                                                        if false {
                                                            v1 = append(v1, "")
                                                            var v56 uint16 = (^(v5))
                                                            var v57 uint16 = (^(v5))
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            v37 -= v37
                                                            fmt.Println(v5, v15, v36)
                                                            delete(v4, 1)
                                                        }
                                                        var v58 int = fi25
                                                        _ = v58
                                                    }
                                                } else {
                                                    v53 = ""
                                                    fmt.Println(v53)
                                                    switch int(v5) {
                                                    case 127:
                                                        v53 = v53
                                                    case 1:
                                                        v53 = v53
                                                    case 522:
                                                        if true {
                                                            var v59 uint64 = 1
                                                            _ = v59
                                                        }
                                                    default:
                                                        v3 = -952
                                                    }
                                                }
                                                switch (v53 + "a") {
                                                case "\n\t":
                                                    if false {
                                                        var v60 uint = (v15 >> ((uint(v5)) & 63))
                                                        var v61 float32 = 1.5
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v2 = append(v2, "a")
                                                        v53 = v53
                                                    }
                                                case "a":
                                                    switch v53 {
                                                    case "":
                                                        v3 += v3
                                                    case "foo bar":
                                                        var v62 uint32 = 133
                                                        _ = v62
                                                    default:
                                                        var v63 int64 = v36
                                                        _ = v63
                                                    }
                                                case "foo bar":
                                                    var v64 bool = (!false)
                                                    _ = v64
                                                }
                                                var v65 float32 = 3.14159
                                                _ = v65
                                            }
                                        case "\n\t":
                                            v66 := GenS3{F0: (-(fi23))}
                                            _ = v66
                                        default:
                                            if true {
                                                if (false && false) {
                                                    var v67 uint8 = 1
                                                    _ = v67
                                                }
                                                var v68 uint64 = 42
                                                _ = v68
                                            }
                                        }
                                    }
                                    _ = v53
                                }
                            } else {
                                fmt.Println(fi23, v3)
                            }
                        default:
                            v4[fi25] = 2147483648
                        }
                        _ = v36
                        _ = v37
                    } else {
                        switch "hello" {
                        case "":
                            v69 := GenS4{F0: v24, F1: 985}
                            _ = v69
                        case "hello":
                            v24 -= (-(v24))
                        }
                    }
                }
            }
            v15 += (v15 & (v15 + (^(v15))))
            var v70 float32 = float32(0.1)
            fmt.Println(v3, v15)
            _ = len(v2)
            _ = fi25
            _ = v70
        }
        _ = fi23
        _ = v24
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        var v72 uint16 = genFn0(1, (v15 + (v15 & 625)))
        _ = fi71
        _ = v72
    }
    var v73 int64 = 42
    var v74 int = 4294967295
    var v75 uint = v15
    for fi76 := 0; fi76 < 20; fi76++ {
        var v77 float64 = float64(((3.14159 / 1.0) - 3.14159))
        v3 = v3
        _ = fi76
        _ = v77
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v78 bool = false
    if (genFn0(4294967295, v75) <= v5) {
        if v78 {
            v4[682] = 546
        }
        for fi79 := 0; fi79 < 19; fi79++ {
            v5 += (^(v5))
            _ = fi79
        }
    }
    var v80 uint = (v75 % (((v75 & v15)) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v81 := GenS3{F0: (v74 << ((1) & 63))}
    var v82 uint16 = v5
    v3 -= (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
    _ = v75
    _ = v78
    _ = v80
    _ = v81
    _ = v82
}
