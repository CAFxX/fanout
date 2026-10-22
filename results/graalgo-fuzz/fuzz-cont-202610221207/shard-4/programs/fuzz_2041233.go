package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32, p1 uint64) uint8 {
    p0 -= p0
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: (p1 / ((p1) | 1)), F1: (^(p1)), F2: 2147483647}
        p1 = 42
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: p1, F1: p1, F2: 42}
            if (!false) {
                var v11 int64 = 1
                _ = v11
            } else {
                v12 := GenS2{F0: true}
                var v13 int32 = p0
                _ = v12
                _ = v13
            }
            var v14 bool = ((!true) || false)
            var v15 uint = 986
            v6 = (float32(((0.1 + 3.14159) / -2.25)) * 0.0)
            v16 := GenS3{F0: "hello", F1: 1}
            v17 := GenS4{F0: 528}
            switch int(p1) {
            case -91:
                var v18 float64 = 0.1
                _ = v18
            case 127:
                v19 := []int{-2147483648}
                _ = v19
            default:
                var v20 string = "a"
                _ = v20
            }
            v21 := GenS4{F0: 0}
            _ = fi9
            _ = v10
            _ = v14
            _ = v15
            _ = v16
            _ = v17
            _ = v21
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return uint8(p1)
}
func genFn1(p0 uint32, p1 bool, p2 uint8) uint64 {
    v22 := []int{2147483647, int(p2)}
    v22 = append(v22, 977)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if p1 {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS5{F0: 1, F1: 1}
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint = 385
                if p1 {
                    if p1 {
                        if true {
                            var v27 float32 = 0.1
                            _ = v27
                        }
                        p1 = (!(568 >= 255))
                    } else {
                        if p1 {
                            var v28 int32 = -609
                            if len(v22) > 0 {
                                v22[0] = fi25
                            }
                            fmt.Println(v28, p2, p0)
                            var v29 int16 = 0
                            _ = v28
                            _ = v29
                        } else {
                            var v30 int16 = -45
                            _ = v30
                        }
                        var v31 uint64 = 0
                        if (int8(fi23) >= 127) {
                            var v32 uint32 = (^(p0))
                            var v33 uint = v26
                            if len(v22) > 0 {
                                v22[0] = 128
                            }
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS5{F0: 1, F1: -996}
                        p1 = p1
                        _ = v31
                        _ = v34
                    }
                    switch "a" {
                    case "":
                        if (42 == 891) {
                            var v35 uint32 = 4294967295
                            _ = v35
                        } else {
                            p1 = (p0 < (p0 * p0))
                            var v36 int8 = 100
                            if (0 != 110) {
                                v37 := GenS3{F0: "foo bar", F1: (v36 / ((-112) | 1))}
                                var v38 int = -1
                                var v39 bool = true
                                var v40 int16 = 1000
                                var v41 uint64 = 1
                                fmt.Println(fi23, v26, v38)
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            _ = v36
                        }
                    }
                    if ((v26 % ((v26) | 1)) > (v26 & v26)) {
                        var v42 int64 = -723
                        switch 127 {
                        case 32768:
                            p2 = p2
                        case 464:
                            var v43 int16 = -656
                            _ = v43
                        }
                        var v44 string = "hello"
                        _ = v42
                        _ = v44
                    }
                }
                if (v26 < (v26 ^ 738)) {
                    var v45 float32 = float32(-2.25)
                    var v46 uint16 = 1
                    fmt.Println(p2, p0, fi25)
                    _ = v45
                    _ = v46
                } else {
                    p1 = (!("hello" < "a"))
                    switch 32767 {
                    case 127:
                        if true {
                            var v47 float32 = (1e10 / (1.5 + float32(0.1)))
                            if len(v22) > 0 {
                                v22[0] = fi25
                            }
                            _ = v47
                        } else {
                            if len(v22) > 0 {
                                v22[0] = 32767
                            }
                            p2 = p2
                        }
                    case 2147483647:
                        v22 = append(v22, fi23)
                    case -9223372036854775808:
                        var v48 uint32 = (p0 / ((p0) | 1))
                        _ = v48
                    default:
                        var v49 int64 = 42
                        _ = v49
                    }
                }
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        switch 700 {
        case -917:
            var v50 uint8 = (p2 | 1)
            _ = v50
        case 4294967295:
            p2 = 0
        case 32768:
            _ = len(v22)
        default:
            p0 += p0
        }
    } else {
        p1 = (p1 && (((p1 && (false || p1)) || p1) && p1))
    }
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 2
    var v3 uint = (v2 << ((v2) & 63))
    var v4 int64 = -1
    v5 := make(map[int]string)
    v4 -= v4
    if ((!(!(89 > -1))) || (!(136 <= 65535))) {
        switch 32767 {
        case 898:
            v51 := GenS4{F0: -1}
            _ = v51
        case 211:
            if false {
                v4 -= (v4 >> ((42) & 63))
                var v52 uint16 = 32767
                _ = v52
            }
        }
    } else {
        var v53 uint32 = 4294967295
        _ = v53
    }
    var v54 int32 = -2147483648
    v3 = v2
    v55 := GenS3{F0: "hello", F1: 42}
    v5[-832] = "hello"
    delete(v5, -577)
    var v56 int16 = -1
    for fi57 := 0; fi57 < 10; fi57++ {
        for fi58 := 0; fi58 < 4; fi58++ {
            if (("foo bar" >= "\n\t") || false) {
                var v59 int8 = 127
                if true {
                    fmt.Println(fi57, v2)
                    var v60 uint = (^(v2))
                    _ = v60
                }
                v54 -= (^(v54))
                _ = v59
            }
            if true {
                v61 := GenS1{F0: 267, F1: genFn1(463, true, uint8(v3)), F2: 2}
                v62 := GenS6{F0: 9}
                _ = v61
                _ = v62
            }
            _ = fi58
        }
        var v63 int8 = 48
        v64 := GenS7{F0: 1, F1: (v4 ^ 9223372036854775807)}
        _ = fi57
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 8; fi65++ {
        var v66 uint8 = 212
        _ = fi65
        _ = v66
    }
    v2 += v3
    v5[-162] = ""
    if (!true) {
        var v67 string = "a"
        _ = v67
    } else {
        fmt.Println(v2, v54)
        v2 = v2
        fmt.Println(v56, v54, v2)
    }
    delete(v1, 2147483647)
    if false {
        var v68 bool = true
        var v69 float32 = ((float32(3.14159) + (0.0 - 1.5)) / 0.1)
        _ = v68
        _ = v69
    }
    var v70 float32 = 0.0
    v71 := GenS2{F0: ((42 >= 255) && (!(!false)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v70
    _ = v71
}
