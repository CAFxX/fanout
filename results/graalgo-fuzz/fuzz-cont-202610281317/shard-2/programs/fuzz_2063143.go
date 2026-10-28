package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int64; F1 int16; F2 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 uint; F2 bool }
func genFn0(p0 uint) uint8 {
    if true {
        fmt.Println(p0)
        p0 -= (p0 % ((4294967295) | 1))
    } else {
        fmt.Println(p0)
        v5 := GenS1{F0: "foo bar", F1: -1}
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS2{F0: -1, F1: (fi6 | -104), F2: 828}
                fmt.Println(fi7, fi6)
                v9 := GenS1{F0: "a", F1: 32767}
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        _ = v5
    }
    var v10 int = -961
    v11 := GenS2{F0: 1, F1: (v10 - v10), F2: 255}
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v12 uint = uint(v1)
    var v13 bool = true
    var v14 int32 = -2147483648
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 bool = false
        if v13 {
            v14 = 787
            if v16 {
                var v17 int64 = (v2 | 1)
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := []int{869, fi15, (^(fi15)), (fi18 % ((-1) | 1))}
                    var v20 uint64 = 860
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            }
        }
        var v21 float32 = -2.25
        _ = fi15
        _ = v16
        _ = v21
    }
    if ("hello" > "\n\t") {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int16 = -1
            switch "foo bar" {
            case "foo bar":
                var v24 uint32 = 786
                _ = v24
            case "a":
                if v13 {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v1 = 255
                        _ = fi25
                    }
                    v12 = (v12 | v12)
                    if v13 {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            if ((3.14159 * 0.0) != -2.25) {
                                var v27 string = "hello"
                                _ = len(v3)
                                _ = v27
                            } else {
                                if ((((!false) || v13) || (!true)) && (v13 && (v13 || (((^(v14)) == v14) || (!v13))))) {
                                    var v28 uint16 = v1
                                    v29 := GenS3{F0: "hello", F1: (fi22 & 2147483647)}
                                    var v30 int16 = -220
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    if v13 {
                                        var v31 uint = 447
                                        _ = v31
                                    } else {
                                        switch "a" {
                                        case "":
                                            v32 := GenS4{F0: fi26}
                                            _ = v32
                                        case "hello":
                                            var v33 uint32 = 4294967295
                                            _ = v33
                                        default:
                                            if (!true) {
                                                if v13 {
                                                    if (0.1 <= 1.5) {
                                                        var v34 float32 = (0.1 - 0.1)
                                                        var v35 int32 = (^(v14))
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                                if (1.5 != -2.25) {
                                                    fmt.Println(v23)
                                                } else {
                                                    v23 = v23
                                                }
                                            } else {
                                                v36 := GenS4{F0: fi22}
                                                _ = v36
                                            }
                                        }
                                    }
                                }
                                v37 := GenS5{F0: (fi26 + fi26), F1: 2}
                                _ = v37
                            }
                            v23 += v23
                            _ = fi26
                        }
                    } else {
                        fmt.Println(v23)
                        if (v1 >= 255) {
                            v1 = uint16(v12)
                            var v38 uint32 = 65535
                            _ = v38
                        }
                    }
                } else {
                    for fi39 := 0; fi39 < 9; fi39++ {
                        if (42 <= 1) {
                            v40 := []int{-439, int(v2), fi39}
                            _ = v40
                        } else {
                            v41 := GenS2{F0: v2, F1: int(v12), F2: 443}
                            v42 := GenS6{F0: v2, F1: (v23 & v23), F2: 127}
                            delete(v4, 1)
                            _ = v41
                            _ = v42
                        }
                        var v43 int32 = (v14 | v14)
                        switch ("hello" + "a") {
                        case "a":
                            v23 += 42
                        case "hello":
                            switch "\n\t" {
                            case "":
                                v44 := GenS1{F0: "a", F1: 9223372036854775807}
                                fmt.Println(v43, fi22)
                                var v45 int32 = (v43 >> ((v12) & 63))
                                fmt.Println(v14, v12, v12)
                                _ = v44
                                _ = v45
                            case "a":
                                var v46 int8 = 1
                                _ = v46
                            case "foo bar":
                                var v47 int16 = v23
                                _ = v47
                            default:
                                v48 := v4[(fi22 ^ 4294967295)]
                                v4[(v48 - fi39)] = 128
                                v23 -= (^(v23))
                                _ = v48
                            }
                        case "\n\t":
                            _ = len(v3)
                        default:
                            v43 += v43
                        }
                        v49 := []int{}
                        _ = fi39
                        _ = v43
                        _ = v49
                    }
                    v13 = (v2 > (v2 / (((^(v2))) | 1)))
                    var v50 uint = 255
                    _ = v50
                }
            case "":
                v51 := []int{int(v2), -9223372036854775808}
                _ = v51
            }
            _ = fi22
            _ = v23
        }
        v52 := GenS4{F0: 4294967295}
        _ = v52
    } else {
        v13 = false
    }
    if v13 {
        switch 9223372036854775807 {
        case -620:
            v53 := GenS4{F0: -9223372036854775808}
            _ = v53
        case 0:
            if v13 {
                var v54 int8 = int8(v14)
                _ = v54
            }
        }
    }
    var v55 uint8 = 214
    v56 := []int{32767, 42, -393, 127}
    v13 = v13
    v57 := GenS7{F0: 604}
    v13 = ((v55 + v55) != v55)
    v58 := GenS8{F0: 0, F1: 1, F2: false}
    for fi59 := 0; fi59 < 15; fi59++ {
        _ = len(v4)
        _ = fi59
    }
    fmt.Println(v13, v2, v1)
    v13 = (3.14159 <= -2.25)
    v60 := GenS3{F0: ("hello" + ("" + ("\n\t" + ("\n\t" + "foo bar")))), F1: -9223372036854775808}
    v61 := GenS8{F0: 2, F1: v12, F2: (280 <= 454)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v60
    _ = v61
}
