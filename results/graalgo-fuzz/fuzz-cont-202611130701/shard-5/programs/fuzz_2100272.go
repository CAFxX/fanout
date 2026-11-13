package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 int; F2 uint32 }
type GenS9 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 uint, p1 int32) string {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            fmt.Println(p1, p0, p1)
            if true {
                fmt.Println(fi5)
            }
            switch (("hello" + ("" + ((("foo bar" + "") + "") + "foo bar"))) + ("\n\t" + "foo bar")) {
            case "a":
                if (1.5 < float64(float32(-2.25))) {
                    fmt.Println(fi4, p0, p1)
                    var v6 int8 = 101
                    _ = v6
                }
            }
            var v7 int16 = 1000
            _ = fi5
            _ = v7
        }
        var v8 bool = (true || (!true))
        if (-9223372036854775808 != -1) {
            switch fi4 {
            case 128:
                p1 = (-(p1))
            case 205:
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 float64 = 1.5
                    p0 = p0
                    var v11 int32 = p1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case -2147483648:
                switch ((("hello" + "\n\t") + "foo bar") + ("" + "")) {
                case "foo bar":
                    p0 = 0
                case "":
                    fmt.Println(v8, v8, p1)
                case "hello":
                    for fi12 := 0; fi12 < 14; fi12++ {
                        v13 := []int{}
                        var v14 int64 = -9223372036854775808
                        switch ("" + ("hello" + (("\n\t" + "a") + ("" + "")))) {
                        case "a":
                            var v15 int16 = -1
                            _ = v15
                        case "\n\t":
                            v8 = (false || false)
                        case "foo bar":
                            v13 = append(v13, int(v14))
                        }
                        if (false && true) {
                            v16 := GenS1{F0: p0, F1: 1}
                            _ = len(v13)
                            _ = v16
                        } else {
                            v17 := GenS2{F0: 2, F1: 192}
                            _ = v17
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                default:
                    if (1 > 65535) {
                        var v18 int64 = -329
                        p0 += 255
                        var v19 int16 = 0
                        _ = v18
                        _ = v19
                    } else {
                        p1 += p1
                    }
                }
            }
            v8 = true
        } else {
            var v20 string = ""
            fmt.Println(v8, v8)
            if v8 {
                fmt.Println(p0)
                var v21 float32 = 1e10
                _ = v21
            }
            _ = v20
        }
        _ = fi4
        _ = v8
    }
    switch 32767 {
    case 0:
        v22 := GenS3{F0: 771, F1: 42}
        _ = v22
    case -1:
        p1 = (-(p1))
    case 4294967295:
        v23 := GenS1{F0: 444, F1: 9223372036854775807}
        _ = v23
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    var v1 uint64 = 545
    v2 := make(map[uint32]int)
    v3 := []uint32{42, 221}
    v3 = append(v3, 685)
    var v24 bool = true
    v25 := GenS1{F0: 42, F1: 0}
    v24 = ((!((!(-422 != 0)) && v24)) || (v24 && v24))
    switch -880 {
    case -925:
        if v24 {
            fmt.Println(v24, v1)
            delete(v2, 0)
            v2[65535] = 2147483647
        } else {
            _ = len(v2)
            var v26 uint64 = v1
            _ = v26
        }
    case 638:
        v27 := GenS4{F0: 1, F1: (^(v1))}
        _ = v27
    case 4294967295:
        if true {
            v24 = (!(1e10 >= 1.5))
            var v28 int = 127
            for fi29 := 0; fi29 < 7; fi29++ {
                v28 += v28
                delete(v2, 1)
                for fi30 := 0; fi30 < 5; fi30++ {
                    v3 = append(v3, 42)
                    v1 += v1
                    _ = fi30
                }
                _ = fi29
            }
            _ = v28
        } else {
            var v31 uint8 = 128
            _ = v31
        }
    default:
        v1 = v1
    }
    var v32 int16 = -1
    fmt.Println(v1, v32)
    v1 += 4294967295
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(fi33, v24, v24)
            v35 := GenS5{F0: int32(v32), F1: fi33}
            v36 := v2[65535]
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    var v37 uint64 = (v1 | 42)
    for fi38 := 0; fi38 < 12; fi38++ {
        fmt.Println(fi38, v1)
        var v39 string = "a"
        _ = fi38
        _ = v39
    }
    var v40 uint64 = v37
    if (true && (4294967295 == 2)) {
        if true {
            v32 += (v32 / (((v32 / (((v32 << ((1) & 63))) | 1))) | 1))
        }
    } else {
        delete(v2, 4294967295)
        v24 = true
        v32 += (-(v32))
    }
    if false {
        switch "" {
        case "hello":
            v32 = v32
        case "foo bar":
            var v41 string = "foo bar"
            _ = v41
        }
    } else {
        var v42 int = 42
        switch "foo bar" {
        case "foo bar":
            switch "hello" {
            case "hello":
                if (!v24) {
                    v37 -= (^(v40))
                    switch v42 {
                    case -9223372036854775808:
                        _ = len(v3)
                    }
                    v24 = (0.0 > 1.5)
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        for fi43 := 0; fi43 < 11; fi43++ {
                            var v44 int16 = (v32 ^ (v32 - v32))
                            if false {
                                fmt.Println(v32, v1)
                                switch 32768 {
                                case 32768:
                                    var v45 uint32 = 65535
                                    _ = v45
                                }
                                switch "a" {
                                case "a":
                                    v46 := GenS2{F0: 0, F1: (v44 % ((-865) | 1))}
                                    _ = v46
                                case "":
                                    for fi47 := 0; fi47 < 20; fi47++ {
                                        switch 2147483647 {
                                        case 474:
                                            v48 := GenS6{F0: (v40 ^ (v1 ^ 2)), F1: 4294967295}
                                            _ = v48
                                        case 1:
                                            delete(v2, 4294967295)
                                        case -582:
                                            var v49 int = (v42 + fi43)
                                            _ = v49
                                        }
                                        _ = fi47
                                    }
                                case "hello":
                                    v1 -= 2
                                }
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        for fi50 := 0; fi50 < 12; fi50++ {
                            delete(v2, 42)
                            if v24 {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    v1 = (v1 % ((v1) | 1))
                                    v52 := GenS4{F0: v32, F1: (v1 % ((v1) | 1))}
                                    _ = fi51
                                    _ = v52
                                }
                            }
                            _ = fi50
                        }
                    }
                    v42 += v42
                    switch "\n\t" {
                    case "a":
                        switch "\n\t" {
                        case "\n\t":
                            v2[2] = v42
                        case "hello":
                            v53 := GenS7{F0: 4294967295}
                            _ = v53
                        case "a":
                            for fi54 := 0; fi54 < 9; fi54++ {
                                var v55 uint16 = 42
                                if v24 {
                                    fmt.Println(v42)
                                    fmt.Println(v24)
                                    for fi56 := 0; fi56 < 17; fi56++ {
                                        v1 += v40
                                        v2[2147483648] = 128
                                        if v24 {
                                            v57 := GenS4{F0: (^(v32)), F1: v37}
                                            _ = v57
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                v42 -= fi54
                                            case "\n\t":
                                                var v58 uint = 4294967295
                                                _ = v58
                                            }
                                        }
                                        var v59 uint8 = 127
                                        _ = fi56
                                        _ = v59
                                    }
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        _ = len(v2)
                                    default:
                                        v60 := GenS4{F0: v32, F1: (v37 | (v1 / ((2) | 1)))}
                                        _ = v60
                                    }
                                    var v61 uint = 255
                                    _ = v61
                                }
                                _ = fi54
                                _ = v55
                            }
                        }
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            case "":
                if (true && false) {
                    if (32767 != 42) {
                        switch "a" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        fmt.Println(v42, v37, v42)
                        if false {
                            var v62 int32 = 2147483647
                            v40 = (v1 ^ v1)
                            var v63 float64 = 0.1
                            _ = v62
                            _ = v63
                        }
                    } else {
                        fmt.Println(v32, v42, v32)
                    }
                    if (!(!("a" != "a"))) {
                        var v64 int32 = -510
                        var v65 uint8 = 2
                        var v66 uint = 998
                        _ = v64
                        _ = v65
                        _ = v66
                    } else {
                        var v67 int8 = -128
                        delete(v2, 65535)
                        if (3.14159 > (float64(3.14159) + 1.5)) {
                            v68 := GenS8{F0: 1, F1: int(v40), F2: 0}
                            v69 := GenS9{F0: 0, F1: int64(v32), F2: 32767}
                            if ("\n\t" > "\n\t") {
                                v70 := GenS5{F0: 0, F1: v42}
                                delete(v2, 42)
                                var v71 string = "foo bar"
                                _ = v70
                                _ = v71
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = 719
                                }
                            }
                            _ = v68
                            _ = v69
                        }
                        _ = v67
                    }
                }
            default:
                var v72 uint16 = 1
                _ = v72
            }
        case "a":
            for fi73 := 0; fi73 < 8; fi73++ {
                var v74 int = fi73
                if len(v3) > 0 {
                    v3[0] = 2
                }
                if v24 {
                    var v75 bool = v24
                    if (192 > 2) {
                        v76 := GenS7{F0: 0}
                        fmt.Println(v1, v37, v75)
                        v37 -= (^(v40))
                        _ = v76
                    } else {
                        fmt.Println(v75, v37)
                    }
                    v77 := GenS8{F0: 1, F1: (fi73 - (fi73 * (fi73 >> ((4294967295) & 63)))), F2: 796}
                    _ = v75
                    _ = v77
                }
                _ = fi73
                _ = v74
            }
        default:
            switch 360 {
            case 127:
                v32 = (v32 & (v32 ^ (v32 - -32768)))
            }
        }
        fmt.Println(v32)
        _ = v42
    }
    if v24 {
        var v78 uint64 = 779
        v32 -= v32
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v37
    _ = v40
}
