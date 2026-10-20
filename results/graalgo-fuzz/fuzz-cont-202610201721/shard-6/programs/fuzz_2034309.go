package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 string; F1 int; F2 bool }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint64, p1 bool) uint64 {
    fmt.Println(p1, p1)
    switch "foo bar" {
    case "foo bar":
        if p1 {
            if p1 {
                v6 := GenS1{F0: (1.5 == 1.5)}
                _ = v6
            } else {
                v7 := GenS1{F0: p1}
                switch "" {
                case "hello":
                    v8 := GenS2{F0: 0}
                    _ = v8
                case "foo bar":
                    fmt.Println(p0, p0)
                default:
                    switch 128 {
                    case 128:
                        for fi9 := 0; fi9 < 19; fi9++ {
                            switch "hello" {
                            case "a":
                                for fi10 := 0; fi10 < 15; fi10++ {
                                    v11 := GenS3{F0: 398}
                                    var v12 float64 = 3.14159
                                    _ = fi10
                                    _ = v11
                                    _ = v12
                                }
                            case "":
                                for fi13 := 0; fi13 < 9; fi13++ {
                                    p0 = 149
                                    v14 := GenS2{F0: -1}
                                    if p1 {
                                        var v15 int64 = -9223372036854775808
                                        _ = v15
                                    } else {
                                        v16 := []int{(-(fi9)), fi13}
                                        fmt.Println(p1, p0, p1)
                                        p1 = p1
                                        _ = v16
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                            case "foo bar":
                                p1 = p1
                            }
                            v17 := []int{fi9, (fi9 & (fi9 >> ((uint(p0)) & 63)))}
                            _ = fi9
                            _ = v17
                        }
                    case -9223372036854775808:
                        p0 -= p0
                    default:
                        p0 -= 0
                    }
                }
                _ = v7
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS4{F0: p1}
                    _ = fi19
                    _ = v20
                }
                var v21 int16 = 0
                if (-1 > v21) {
                    if false {
                        p0 -= (p0 + (p0 >> ((1) & 63)))
                    }
                    var v22 float32 = -2.25
                    _ = v22
                } else {
                    fmt.Println(v21, p0, p0)
                    var v23 float32 = float32(0.1)
                    _ = v23
                }
                _ = fi18
                _ = v21
            }
            var v24 float32 = 3.14159
            _ = v24
        }
    default:
        var v25 string = (("hello" + "hello") + ("hello" + "foo bar"))
        _ = v25
    }
    v26 := []int{284}
    if p1 {
        p0 += 2
        v27 := GenS1{F0: p1}
        _ = v27
    } else {
        p1 = (p1 || true)
        p1 = (false || p1)
        p0 -= 18446744073709551615
    }
    v28 := GenS5{F0: 2147483647, F1: ((((!(65535 <= 2)) || p1) && (1.5 != 1e10)) && (!(false && p1)))}
    _ = p0
    _ = p1
    _ = v26
    _ = v28
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    var v3 int32 = -1
    var v4 uint16 = 32767
    v5 := make(map[int]int)
    v2 -= v2
    v29 := []int{1, 42, 4294967295}
    v30 := GenS1{F0: true}
    var v31 string = "foo bar"
    fmt.Println(v4, v31)
    fmt.Println(v2, v31, v31)
    var v32 float32 = float32(0.1)
    fmt.Println(v31, v32, v3)
    v33 := GenS1{F0: (519 == 1)}
    for fi34 := 0; fi34 < 16; fi34++ {
        v2 = v2
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 uint16 = v4
            var v37 int16 = -32768
            _ = len(v29)
            v38 := v5[(fi35 ^ (fi35 / ((127) | 1)))]
            v36 = 0
            if ((^(v36)) > 2) {
                if (1 <= (fi34 / ((fi34) | 1))) {
                    fmt.Println(v36)
                } else {
                    fmt.Println(v38, v2)
                    if true {
                        v39 := GenS6{F0: 128, F1: v3, F2: 42}
                        var v40 uint64 = genFn0(4294967295, false)
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS7{F0: 209, F1: 42, F2: -9223372036854775808}
                        _ = v41
                    }
                    var v42 int8 = 100
                    _ = v42
                }
                if false {
                    var v43 uint32 = 882
                    v44 := GenS8{F0: 1, F1: 127}
                    _ = v43
                    _ = v44
                } else {
                    v45 := GenS7{F0: 255, F1: 979, F2: 1}
                    _ = v45
                }
            } else {
                var v46 uint8 = 1
                fmt.Println(v37, v31, v4)
                if (-9223372036854775808 >= 1) {
                    var v47 uint = 720
                    switch fi35 {
                    case 0:
                        v48 := GenS6{F0: (-(v38)), F1: (v3 | (v3 / ((v3) | 1))), F2: uint64(v36)}
                        var v49 uint8 = v46
                        _ = v48
                        _ = v49
                    default:
                        fmt.Println(v3, v3, v46)
                    }
                    v3 = v3
                    _ = v47
                }
                _ = v46
            }
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    if false {
        v50 := GenS9{F0: v31, F1: 32767, F2: false}
        var v51 string = v31
        if true {
            v2 = 0.0
            if (false || true) {
                for fi52 := 0; fi52 < 1; fi52++ {
                    if len(v29) > 0 {
                        v29[0] = fi52
                    }
                    v53 := GenS10{F0: -1, F1: 564}
                    var v54 int64 = 997
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                switch v31 {
                case "a":
                    fmt.Println(v31, v32)
                case "hello":
                    for fi55 := 0; fi55 < 5; fi55++ {
                        if (42 == 4294967295) {
                            var v56 uint = 135
                            v32 += v2
                            for fi57 := 0; fi57 < 20; fi57++ {
                                _ = len(v5)
                                v29 = append(v29, (fi55 % (((fi55 >> (((v56 % ((161) | 1))) & 63))) | 1)))
                                v3 += (v3 / (((-(v3))) | 1))
                                var v58 int8 = -128
                                _ = fi57
                                _ = v58
                            }
                            _ = v56
                        }
                        fmt.Println(fi55, v3)
                        _ = fi55
                    }
                }
                var v59 uint16 = v4
                _ = v59
            } else {
                v4 -= 65535
                if false {
                    v60 := v5[1]
                    _ = len(v29)
                    v61 := GenS2{F0: 1}
                    _ = v60
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 11; fi62++ {
                        v29 = append(v29, 2147483647)
                        v63 := GenS8{F0: -1, F1: 1}
                        _ = fi62
                        _ = v63
                    }
                    var v64 string = "foo bar"
                    _ = v64
                }
                fmt.Println(v51)
            }
            for fi65 := 0; fi65 < 15; fi65++ {
                var v66 float32 = (3.14159 / v32)
                var v67 int16 = 706
                if (0.0 >= 0.1) {
                    fmt.Println(v4, v32)
                    v68 := GenS10{F0: -1, F1: 778}
                    _ = v68
                }
                _ = fi65
                _ = v66
                _ = v67
            }
        } else {
            var v69 uint8 = 127
            v5[1] = 2147483647
            _ = v69
        }
        _ = v50
        _ = v51
    }
    var v70 uint16 = v4
    if len(v29) > 0 {
        v29[0] = 320
    }
    if (true || false) {
        switch "a" {
        case "foo bar":
            var v71 int16 = -1
            _ = v71
        case "a":
            if len(v29) > 0 {
                v29[0] = 9223372036854775807
            }
        }
        v72 := GenS11{F0: "hello", F1: -450}
        _ = v72
    } else {
        var v73 string = v31
        var v74 uint8 = 127
        _ = v73
        _ = v74
    }
    switch ("foo bar" + "foo bar") {
    case "":
        v2 -= 3.14159
    case "foo bar":
        var v75 int64 = int64(v3)
        _ = v75
    default:
        v76 := GenS12{F0: v70, F1: (v3 % ((v3) | 1))}
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v70
}
