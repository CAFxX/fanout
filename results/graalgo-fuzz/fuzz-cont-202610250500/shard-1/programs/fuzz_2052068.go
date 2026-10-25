package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0(p0 uint16, p1 string, p2 uint8) int8 {
    v5 := GenS1{F0: (p2 + (p2 * 26)), F1: (p1 + ("\n\t" + ("a" + ((p1 + p1) + "")))), F2: 1000}
    var v6 float32 = 1.5
    p1 = p1
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 19; fi7++ {
        if (true || true) {
            v8 := GenS2{F0: 127, F1: p2, F2: 2147483647}
            p2 += 127
            _ = v8
        } else {
            p1 = p1
            switch 2147483647 {
            case 128:
                fmt.Println(v6)
            case 2147483647:
                var v9 uint64 = 833
                _ = v9
            case 32768:
                v6 -= v6
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 bool = (22 != -35)
    v4 := make(map[uint32]bool)
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            if v3 {
                v3 = (-670 != 379)
                var v11 uint = 65535
                _ = v11
            } else {
                v12 := GenS1{F0: 128, F1: "\n\t", F2: -956}
                for fi13 := 0; fi13 < 13; fi13++ {
                    switch fi13 {
                    case 42:
                        switch "hello" {
                        case "":
                            v3 = false
                        }
                    case -273:
                        _ = len(v4)
                    case 601:
                        v14 := GenS2{F0: 48, F1: 1, F2: 0}
                        _ = v14
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi13
                }
                _ = v12
            }
            _ = fi10
        }
        if (2 >= 42) {
            for fi15 := 0; fi15 < 17; fi15++ {
                if (42 != 32767) {
                    v16 := GenS2{F0: 0, F1: 0, F2: 2147483647}
                    for fi17 := 0; fi17 < 19; fi17++ {
                        delete(v2, 2147483648)
                        _ = fi17
                    }
                    _ = v16
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch ("a" + "hello") {
                    case "\n\t":
                        fmt.Println(v3, fi15, fi18)
                    }
                    _ = fi18
                }
                _ = fi15
            }
            switch "" {
            case "hello":
                v19 := GenS3{F0: 4294967295, F1: true}
                _ = v19
            default:
                v20 := v4[2]
                _ = v20
            }
            v21 := GenS4{F0: true}
            _ = v21
        } else {
            var v22 int16 = -32768
            var v23 uint8 = 128
            var v24 int16 = (-(v22))
            _ = v22
            _ = v23
            _ = v24
        }
        if true {
            for fi25 := 0; fi25 < 12; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v3 = true
                    var v27 int = (fi25 ^ (fi25 * -577))
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        } else {
            v28 := GenS5{F0: 1, F1: 860}
            var v29 string = "a"
            _ = v28
            _ = v29
        }
    } else {
        var v30 int32 = 42
        v3 = false
        _ = v30
    }
    switch "a" {
    case "foo bar":
        for fi31 := 0; fi31 < 13; fi31++ {
            v3 = v3
            fmt.Println(fi31, v3)
            _ = fi31
        }
    default:
        var v32 bool = (!true)
        _ = v32
    }
    fmt.Println(v3)
    fmt.Println(v3)
    var v33 int64 = -1
    if false {
        fmt.Println(v3, v33)
        if v3 {
            var v34 int16 = 1
            switch -2147483648 {
            case -210:
                v3 = true
            default:
                var v35 string = "foo bar"
                _ = v35
            }
            _ = v34
        }
    } else {
        v33 = (v33 << ((926) & 63))
        if true {
            var v36 int64 = -478
            _ = v36
        } else {
            if v3 {
                var v37 string = "foo bar"
                for fi38 := 0; fi38 < 1; fi38++ {
                    v3 = (53 > 122)
                    switch 32767 {
                    case -956:
                        if (-535 <= 1) {
                            for fi39 := 0; fi39 < 13; fi39++ {
                                v37 = (("a" + "foo bar") + "a")
                                switch fi39 {
                                case -79:
                                    var v40 float32 = 0.1
                                    _ = v40
                                case -9223372036854775808:
                                    if false {
                                        v41 := []int{fi38}
                                        var v42 bool = ((v37 + ((("foo bar" + (v37 + v37)) + v37) + v37)) != v37)
                                        if (0.1 < 1.5) {
                                            if v42 {
                                                v43 := GenS6{F0: 65535}
                                                v42 = v42
                                                _ = v43
                                            } else {
                                                v41 = append(v41, (fi39 / ((fi38) | 1)))
                                                switch fi39 {
                                                case 128:
                                                    v37 = (v37 + v37)
                                                case 32768:
                                                    v41 = append(v41, 42)
                                                case -745:
                                                    if true {
                                                        v3 = (v33 <= (-(v33)))
                                                    } else {
                                                        v44 := GenS3{F0: fi39, F1: (genFn0(32767, "foo bar", 204) == -128)}
                                                        v3 = v3
                                                        _ = v44
                                                    }
                                                }
                                            }
                                        } else {
                                            var v45 int = (fi38 << ((53) & 63))
                                            _ = v45
                                        }
                                        if (255 == 234) {
                                            _ = len(v4)
                                            v46 := GenS5{F0: 1, F1: 608}
                                            _ = v46
                                        }
                                        var v47 int64 = (v33 << ((65535) & 63))
                                        var v48 string = "a"
                                        _ = v41
                                        _ = v42
                                        _ = v47
                                        _ = v48
                                    }
                                case -1:
                                    if false {
                                        var v49 uint = 595
                                        v50 := v4[717]
                                        v51 := GenS1{F0: 2, F1: (v37 + v37), F2: 1}
                                        var v52 float64 = 1e10
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    } else {
                                        if (!(42 != 4294967295)) {
                                            if v3 {
                                                if v3 {
                                                    v53 := GenS1{F0: 0, F1: "foo bar", F2: -1}
                                                    if (140 != 20) {
                                                        switch (^(fi39)) {
                                                        case 4294967295:
                                                            var v54 string = "foo bar"
                                                            _ = v54
                                                        case -1:
                                                            v33 = (v33 + 42)
                                                        default:
                                                            v4[65535] = v3
                                                        }
                                                        var v55 bool = v3
                                                        v56 := []int{-1}
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    v37 = "\n\t"
                                                    _ = v53
                                                }
                                                v57 := []int{}
                                                _ = v57
                                            }
                                        } else {
                                            v4[42] = (138 >= 2)
                                        }
                                    }
                                default:
                                    v58 := GenS1{F0: 42, F1: v37, F2: 1}
                                    _ = v58
                                }
                                _ = fi39
                            }
                            v33 = (v33 + 173)
                        }
                    }
                    v59 := GenS1{F0: 2, F1: ((("foo bar" + v37) + (v37 + v37)) + (v37 + v37)), F2: 1}
                    _ = fi38
                    _ = v59
                }
                _ = v37
            } else {
                var v60 int16 = -1
                v33 += (v33 >> ((42) & 63))
                v60 = v60
                _ = v60
            }
        }
    }
    for fi61 := 0; fi61 < 13; fi61++ {
        fmt.Println(v3, fi61, v33)
        fmt.Println(fi61)
        _ = fi61
    }
    v62 := []int{-552}
    _ = len(v2)
    v63 := GenS7{F0: genFn0(284, ("\n\t" + "foo bar"), 7), F1: 0, F2: 1}
    var v64 bool = v3
    if v3 {
        if true {
            if len(v62) >= 2 {
                v62 = v62[1:len(v62)]
            }
        }
    }
    var v65 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
