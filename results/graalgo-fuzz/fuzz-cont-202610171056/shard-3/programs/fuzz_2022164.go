package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int8; F2 int }
type GenS6 struct { F0 bool; F1 bool }
func genFn0() int {
    if false {
        var v4 int = 32767
        _ = v4
    }
    v5 := GenS1{F0: 820, F1: 0, F2: 2147483647}
    _ = v5
    return 551
}
func genFn1(p0 int, p1 uint32, p2 uint8) uint16 {
    p2 = (p2 - (p2 ^ p2))
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            p2 = p2
            _ = fi7
        }
        if ((0 > 42) && (!true)) {
            switch 579 {
            case -118:
                var v8 float32 = 0.1
                _ = v8
            case -917:
                for fi9 := 0; fi9 < 17; fi9++ {
                    p0 = (fi6 | p0)
                    p1 = p1
                    v10 := GenS2{F0: 422, F1: (^(fi9)), F2: (p1 << ((884) & 63))}
                    _ = fi9
                    _ = v10
                }
            default:
                var v11 int8 = -94
                _ = v11
            }
            var v12 int64 = 0
            _ = v12
        }
        p0 -= genFn0()
        _ = fi6
    }
    if false {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{genFn0()}
            var v15 uint16 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
        if (!(!(!(65535 != 1)))) {
            for fi16 := 0; fi16 < 1; fi16++ {
                p2 += 127
                switch "hello" {
                case "":
                    var v17 int64 = 402
                    _ = v17
                default:
                    var v18 uint32 = 4294967295
                    _ = v18
                }
                var v19 int8 = 35
                _ = fi16
                _ = v19
            }
            switch p0 {
            case 637:
                v20 := GenS2{F0: 255, F1: (-(p0)), F2: (p1 - (^(p1)))}
                _ = v20
            case -908:
                fmt.Println(p0, p2, p1)
            case 32767:
                fmt.Println(p2, p1, p0)
            default:
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(fi21, p0)
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v24 := GenS3{F0: false, F1: 1, F2: -1}
                        var v25 uint16 = 32767
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS4{F0: genFn0()}
                    _ = fi21
                    _ = v26
                }
            }
            var v27 uint16 = uint16(p2)
            _ = v27
        } else {
            if (0.1 >= 1.5) {
                switch "a" {
                case "":
                    var v28 int64 = 0
                    _ = v28
                case "hello":
                    var v29 int16 = -612
                    _ = v29
                case "foo bar":
                    switch "hello" {
                    case "":
                        var v30 int16 = -266
                        _ = v30
                    default:
                        v31 := GenS3{F0: true, F1: 100, F2: -128}
                        _ = v31
                    }
                }
                var v32 uint64 = 18446744073709551615
                var v33 string = "foo bar"
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 2; fi34++ {
                for fi35 := 0; fi35 < 14; fi35++ {
                    if (p2 <= p2) {
                        var v36 int = (^(fi34))
                        switch (v36 << ((304) & 63)) {
                        case 42:
                            if true {
                                switch (p0 >> ((4294967295) & 63)) {
                                case 9223372036854775807:
                                    v36 = (v36 + 127)
                                    fmt.Println(fi35, p0)
                                default:
                                    v37 := []int{(fi34 ^ (-(fi34)))}
                                    _ = v37
                                }
                                v38 := GenS5{F0: 100, F1: 1, F2: 448}
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        v39 := []int{(v36 - 713)}
                                        fmt.Println(p2)
                                        var v40 string = ("a" + ("hello" + "hello"))
                                        _ = v39
                                        _ = v40
                                    }
                                case "foo bar":
                                    switch ("" + ("a" + "")) {
                                    case "a":
                                        var v41 string = "a"
                                        _ = v41
                                    default:
                                        v42 := []int{fi35}
                                        _ = v42
                                    }
                                }
                                _ = v38
                            } else {
                                v43 := GenS5{F0: 110, F1: 100, F2: (fi35 | (fi34 | fi34))}
                                _ = v43
                            }
                        case 161:
                            if (316 > 883) {
                                switch (fi34 - (^(v36))) {
                                case 2147483647:
                                    var v44 bool = (("hello" + "hello") < "a")
                                    _ = v44
                                case 0:
                                    switch genFn0() {
                                    case -672:
                                        fmt.Println(p1, fi34)
                                    default:
                                        v45 := []int{int(p2), 32767, genFn0()}
                                        _ = v45
                                    }
                                }
                                fmt.Println(fi34, p2, p0)
                            } else {
                                v36 -= int(p1)
                            }
                        default:
                            fmt.Println(fi34, v36)
                        }
                        _ = v36
                    }
                    v46 := []int{(fi35 * -2147483648), fi34}
                    _ = fi35
                    _ = v46
                }
                _ = fi34
            }
            if true {
                for fi47 := 0; fi47 < 5; fi47++ {
                    var v48 int8 = 114
                    switch "\n\t" {
                    case "foo bar":
                        var v49 uint32 = (p1 ^ (p1 << ((65535) & 63)))
                        _ = v49
                    case "a":
                        v50 := []int{(p0 % ((p0) | 1)), fi47}
                        _ = v50
                    case "":
                        p2 += 2
                    }
                    var v51 bool = false
                    _ = fi47
                    _ = v48
                    _ = v51
                }
            }
        }
    } else {
        for fi52 := 0; fi52 < 9; fi52++ {
            v53 := []int{fi52, (^(p0))}
            fmt.Println(fi52, fi52, p2)
            _ = fi52
            _ = v53
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint8 = 127
    var v2 int16 = 1000
    var v3 uint64 = 904
    v1 += 58
    var v54 float64 = -2.25
    v54 += v54
    if ((v3 & v3) <= v3) {
        v54 += 3.14159
    } else {
        v1 -= (v1 ^ v1)
        v55 := []int{32767, 42}
        _ = v55
    }
    v3 = v3
    var v56 float32 = (0.1 * (float32(float64(3.14159)) - 0.1))
    fmt.Println(v1)
    for fi57 := 0; fi57 < 3; fi57++ {
        for fi58 := 0; fi58 < 6; fi58++ {
            if true {
                fmt.Println(v56)
                v59 := []int{}
                v59 = append(v59, (fi57 ^ fi58))
                _ = v59
            }
            v60 := []int{}
            var v61 uint = uint(v3)
            _ = fi58
            _ = v60
            _ = v61
        }
        _ = fi57
    }
    var v62 bool = true
    if (32767 <= 2) {
        v54 += v54
    } else {
        v63 := []int{}
        v2 -= (v2 - (v2 ^ 42))
        _ = v63
    }
    fmt.Println(v56, v2, v3)
    var v64 int8 = -71
    for fi65 := 0; fi65 < 12; fi65++ {
        for fi66 := 0; fi66 < 10; fi66++ {
            switch "" {
            case "foo bar":
                v67 := []int{}
                _ = v67
            case "a":
                if true {
                    v56 -= 0.1
                } else {
                    v68 := []int{32768, fi66, genFn0()}
                    _ = v68
                }
            default:
                v3 -= v3
            }
            _ = fi66
        }
        var v69 float32 = v56
        _ = fi65
        _ = v69
    }
    v62 = (139 > 255)
    for fi70 := 0; fi70 < 4; fi70++ {
        v71 := GenS6{F0: v62, F1: ((!(!(2 == 1))) || v62)}
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v56
    _ = v62
    _ = v64
}
