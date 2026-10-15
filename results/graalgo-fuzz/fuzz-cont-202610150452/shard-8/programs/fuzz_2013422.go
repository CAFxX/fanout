package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 int8, p1 int16) int32 {
    p0 = -1
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint16 = 184
        p0 = p0
        p1 -= (p1 | 42)
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 0
    p1 = p1
    switch -9223372036854775808 {
    case 0:
        v10 := GenS1{F0: v9}
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 241}
            _ = fi11
            _ = v12
        }
    case -1:
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{(fi13 - int(p0)), fi13}
            _ = len(v14)
            _ = fi13
            _ = v14
        }
    default:
        var v15 int32 = 42
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v9
    return 267
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 65535
    var v3 int64 = 1
    var v4 bool = (42 < 139)
    var v5 string = ""
    var v6 int = 882
    switch v5 {
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := []int{v6}
            var v18 string = "hello"
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        v19 := []int{(v6 << ((42) & 63)), (^(v6)), v6}
        _ = v19
    }
    v5 = "foo bar"
    v5 = v5
    var v20 float32 = 0.0
    fmt.Println(v20, v20)
    v2 -= 1
    v21 := []int{}
    fmt.Println(v6)
    if ((v20 / (v20 - v20)) <= -2.25) {
        var v22 uint32 = (v2 % ((539) | 1))
        _ = v22
    } else {
        for fi23 := 0; fi23 < 16; fi23++ {
            switch fi23 {
            case 127:
                for fi24 := 0; fi24 < 5; fi24++ {
                    if (((0.0 + float64((v20 - (v20 - v20)))) != ((1.5 - -2.25) + float64((v20 - ((float32((1.5 * float64(v20))) - v20) / v20))))) && ((-221 != 42) || v4)) {
                        v4 = v4
                        switch (fi24 | v6) {
                        case 74:
                            fmt.Println(v6, v2, v20)
                            switch ("a" + "foo bar") {
                            case "a":
                                var v25 int64 = (v3 >> ((255) & 63))
                                _ = v25
                            case "":
                                _ = len(v21)
                            }
                            var v26 float32 = (1.5 - 0.0)
                            _ = v26
                        case -9223372036854775808:
                            var v27 int = v6
                            _ = v27
                        case 127:
                            switch fi24 {
                            case -663:
                                var v28 uint8 = 42
                                _ = v28
                            case 282:
                                if true {
                                    var v29 int = (-(fi23))
                                    _ = v29
                                } else {
                                    var v30 string = (((("a" + "foo bar") + "foo bar") + v5) + v5)
                                    _ = v30
                                }
                            case 9223372036854775807:
                                v31 := GenS1{F0: 128}
                                _ = v31
                            default:
                                fmt.Println(v2)
                            }
                        }
                    } else {
                        var v32 int = fi24
                        var v33 uint = 270
                        if false {
                            fmt.Println(v3, v5, v33)
                        } else {
                            var v34 uint = (v33 | (v33 - v33))
                            if true {
                                v35 := GenS1{F0: 58}
                                _ = v35
                            }
                            if v4 {
                                v36 := GenS2{F0: -577, F1: 42, F2: (^(v2))}
                                fmt.Println(fi23, fi23)
                                _ = v36
                            }
                            _ = v34
                        }
                        fmt.Println(v3, v2, v20)
                        var v37 uint8 = 215
                        if (v33 <= uint(v37)) {
                            v5 = v5
                            v38 := GenS3{F0: (v37 & (v37 % ((v37) | 1))), F1: (^(v33))}
                            _ = v38
                        } else {
                            fmt.Println(v32, v2)
                            if false {
                                var v39 bool = false
                                _ = v39
                            } else {
                                if v4 {
                                    var v40 uint32 = (^(v2))
                                    v41 := GenS4{F0: 2, F1: 18446744073709551615}
                                    _ = v40
                                    _ = v41
                                }
                                v2 += (^(v2))
                                v5 = ((v5 + "\n\t") + ("a" + "\n\t"))
                            }
                            fmt.Println(v37, v2, v37)
                        }
                        _ = v32
                        _ = v33
                        _ = v37
                    }
                    v20 = v20
                    var v42 int16 = 1000
                    _ = fi24
                    _ = v42
                }
            default:
                v43 := GenS2{F0: 334, F1: 42, F2: (v2 % ((1) | 1))}
                _ = v43
            }
            var v44 int8 = 100
            _ = fi23
            _ = v44
        }
        v45 := GenS5{F0: genFn0(85, -250), F1: 769}
        var v46 float32 = v20
        _ = v45
        _ = v46
    }
    _ = len(v21)
    v47 := GenS4{F0: 255, F1: 42}
    var v48 float32 = (-2.25 / 1.0)
    switch v5 {
    case "":
        if true {
            for fi49 := 0; fi49 < 20; fi49++ {
                v5 = "\n\t"
                fmt.Println(fi49, v6, v20)
                _ = fi49
            }
            for fi50 := 0; fi50 < 3; fi50++ {
                v51 := GenS3{F0: 0, F1: 0}
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 float32 = v48
                    if v4 {
                        var v54 int32 = 1
                        _ = v54
                    } else {
                        if (32767 < 0) {
                            v55 := GenS5{F0: 1, F1: 239}
                            v6 -= -2147483648
                            _ = v55
                        } else {
                            v5 = (v5 + (("hello" + "foo bar") + v5))
                            v53 -= 0.1
                            v6 = 42
                            var v56 uint32 = 779
                            _ = v56
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                _ = fi50
                _ = v51
            }
            for fi57 := 0; fi57 < 17; fi57++ {
                switch (v6 / ((v6) | 1)) {
                case -1:
                    v2 -= 2147483647
                case -2147483648:
                    var v58 uint32 = v2
                    _ = v58
                case 9223372036854775807:
                    v59 := GenS5{F0: 0, F1: 2}
                    _ = v59
                default:
                    var v60 string = ("" + v5)
                    _ = v60
                }
                if true {
                    _ = len(v1)
                    var v61 float64 = ((3.14159 - (1e10 - (1.5 - 1e10))) / ((-2.25 + 1.5) + 0.0))
                    _ = v61
                }
                var v62 int32 = 1
                _ = fi57
                _ = v62
            }
        }
    case "\n\t":
        fmt.Println(v20, v3)
    case "hello":
        switch (v6 - -941) {
        case 127:
            switch v6 {
            case 0:
                v6 = (v6 % ((128) | 1))
            case -886:
                var v63 int8 = 0
                _ = v63
            case -101:
                if (uint64(v2) <= 18446744073709551615) {
                    var v64 int32 = 0
                    for fi65 := 0; fi65 < 14; fi65++ {
                        v66 := GenS5{F0: (v64 + 42), F1: 2}
                        var v67 int64 = (-(v3))
                        var v68 float64 = 3.14159
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    _ = len(v21)
                    _ = v64
                }
            default:
                v69 := GenS2{F0: int16(v2), F1: int8(v6), F2: v2}
                _ = v69
            }
        default:
            var v70 bool = ((v6 / (((v6 << ((2) & 63))) | 1)) <= v6)
            _ = v70
        }
    default:
        var v71 uint64 = 13
        _ = v71
    }
    var v72 int = 128
    v4 = true
    v48 = v48
    var v73 uint64 = 825
    v20 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v72
    _ = v73
}
