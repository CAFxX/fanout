package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 bool; F1 int64 }
func genFn0(p0 string, p1 uint64, p2 uint16) int8 {
    var v7 bool = false
    if v7 {
        p0 = p0
        p0 = (("\n\t" + p0) + "hello")
    } else {
        v8 := []int{-51, 32767, -9223372036854775808, 32768}
        if ((^(p2)) < (p2 << ((844) & 63))) {
            var v9 int64 = -104
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v9
        } else {
            var v10 uint32 = 65535
            if false {
                var v11 int32 = -2147483648
                for fi12 := 0; fi12 < 13; fi12++ {
                    _ = len(v8)
                    v13 := GenS1{F0: p1}
                    if true {
                        v14 := GenS2{F0: p0, F1: (p0 + (p0 + "foo bar"))}
                        _ = v14
                    } else {
                        fmt.Println(p1, p0)
                        p1 = 0
                        for fi15 := 0; fi15 < 14; fi15++ {
                            if v7 {
                                fmt.Println(v10, fi12)
                                var v16 uint8 = 193
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v16
                            }
                            var v17 int = 127
                            fmt.Println(v7, p2, v11)
                            _ = fi15
                            _ = v17
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                if len(v8) > 0 {
                    v8[0] = -3
                }
                _ = v11
            }
            p0 = "\n\t"
            _ = v10
        }
        var v18 int64 = int64(p2)
        _ = v18
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -128
}
func genFn1(p0 bool) uint32 {
    var v19 uint8 = 128
    for fi20 := 0; fi20 < 2; fi20++ {
        switch "foo bar" {
        case "a":
            if p0 {
                v21 := []int{fi20}
                if p0 {
                    p0 = (252 < 1)
                    fmt.Println(v19, v19, fi20)
                }
                v19 = v19
                _ = v21
            } else {
                var v22 int16 = 1
                v22 -= (-(v22))
                var v23 int16 = (^(v22))
                _ = v22
                _ = v23
            }
        case "hello":
            for fi24 := 0; fi24 < 20; fi24++ {
                p0 = p0
                switch fi20 {
                case 0:
                    fmt.Println(p0)
                }
                _ = fi24
            }
        case "foo bar":
            if false {
                switch (("a" + ("a" + "a")) + "foo bar") {
                case "hello":
                    if p0 {
                        if false {
                            fmt.Println(p0, v19, v19)
                            if p0 {
                                var v25 int64 = 9223372036854775807
                                fmt.Println(p0)
                                v19 -= (^(v19))
                                _ = v25
                            } else {
                                v26 := []int{(-(fi20))}
                                _ = v26
                            }
                            v27 := GenS3{F0: 535}
                            _ = v27
                        }
                        switch fi20 {
                        case 2147483647:
                            v28 := GenS4{F0: 9223372036854775807, F1: 9223372036854775807, F2: 255}
                            _ = v28
                        case -323:
                            var v29 int64 = 9223372036854775807
                            _ = v29
                        case 128:
                            if p0 {
                                switch "a" {
                                case "":
                                    v30 := GenS5{F0: -9223372036854775808, F1: 367}
                                    _ = v30
                                case "a":
                                    v19 -= (v19 * (v19 & (v19 & 88)))
                                default:
                                    v19 += v19
                                }
                                var v31 uint16 = 127
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    p0 = (!true)
                                    v33 := GenS5{F0: fi20, F1: uint64(v31)}
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v31
                            } else {
                                p0 = true
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    v19 -= (v19 * (^(v19)))
                                    var v35 string = "hello"
                                    var v36 int32 = -370
                                    if ((1.5 < float32(1e10)) && p0) {
                                        v37 := []int{}
                                        v36 = v36
                                        _ = v37
                                    }
                                    var v38 int8 = -128
                                    v35 = v35
                                    v39 := GenS3{F0: 42}
                                    v40 := GenS6{F0: 4294967295, F1: 255}
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            }
                        }
                    }
                case "a":
                    v41 := GenS7{F0: p0, F1: 1}
                    _ = v41
                default:
                    v42 := GenS6{F0: 942, F1: 65535}
                    _ = v42
                }
                v19 = (v19 + v19)
                var v43 uint32 = 2147483648
                _ = v43
            } else {
                v44 := []int{fi20, (-(fi20)), fi20}
                _ = v44
            }
        default:
            v45 := []int{-889, (^(fi20))}
            _ = v45
        }
        p0 = p0
        _ = fi20
    }
    _ = p0
    _ = v19
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 string = "a"
    var v4 bool = ((v3 <= v3) || (275 != 199))
    var v5 int32 = 1
    v6 := make(map[uint32]bool)
    var v46 uint16 = 42
    for fi47 := 0; fi47 < 15; fi47++ {
        if (v46 != (v46 * uint16(v2))) {
            switch fi47 {
            case 1:
                if (v5 >= v5) {
                    if v4 {
                        v5 += v5
                    }
                    var v48 int8 = 3
                    var v49 float64 = (((float64(1e10) - -2.25) / 1.5) * ((((1.5 + (float64(1e10) * float64(0.0))) - 1.5) + 1e10) * 1e10))
                    _ = v48
                    _ = v49
                }
            case 484:
                if v4 {
                    v5 = 1
                }
            default:
                var v50 float64 = float64(1e10)
                _ = v50
            }
            var v51 uint8 = 66
            if true {
                v51 = v51
            } else {
                v2 = v2
                if (v3 == "") {
                    fmt.Println(v4, fi47)
                    switch v3 {
                    case "hello":
                        var v52 float32 = (-2.25 / ((((1.5 - 3.14159) - (3.14159 / 1.0)) + 0.0) * (1.5 / 1.0)))
                        _ = v52
                    case "\n\t":
                        var v53 string = v3
                        _ = v53
                    case "":
                        v3 = ((v3 + "foo bar") + (v3 + (v3 + ((v3 + "a") + v3))))
                    }
                } else {
                    v54 := []int{}
                    v55 := GenS6{F0: (v2 << ((42) & 63)), F1: (v46 >> ((100) & 63))}
                    _ = v54
                    _ = v55
                }
                var v56 uint8 = v51
                _ = v56
            }
            _ = v51
        } else {
            var v57 float64 = 1.5
            _ = v57
        }
        v58 := []int{}
        v2 = (v2 % ((v2) | 1))
        _ = fi47
        _ = v58
    }
    var v59 uint8 = 128
    v4 = v4
    v5 = -2147483648
    var v60 uint32 = 42
    _ = len(v1)
    v6[(v60 | (v60 % (((^(v60))) | 1)))] = v4
    v3 = v3
    if v4 {
        var v61 int = 32767
        _ = v61
    }
    var v62 int32 = v5
    var v63 int16 = 823
    if ((v60 << ((610) & 63)) != (v60 >> ((255) & 63))) {
        v63 = (^(v63))
        v64 := []int{}
        _ = v64
    }
    v6[genFn1(false)] = ((42 > -1) || true)
    if (((v3 + (("a" + v3) + "a")) + (v3 + v3)) != "hello") {
        v4 = false
    }
    v65 := []int{32767, 1, int(v63), -901}
    fmt.Println(v46, v63, v4)
    switch v3 {
    case "":
        var v66 uint32 = (v60 % (((v60 * v60)) | 1))
        _ = v66
    default:
        v5 = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v62
    _ = v63
    _ = v65
}
