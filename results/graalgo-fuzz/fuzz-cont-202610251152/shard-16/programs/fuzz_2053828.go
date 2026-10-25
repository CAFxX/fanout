package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint8; F1 int16 }
func genFn0(p0 int, p1 uint8, p2 int16) uint {
    if true {
        v7 := []int{(p0 * p0), p0, (p0 | p0)}
        _ = v7
    } else {
        switch -603 {
        case -9223372036854775808:
            var v8 string = "foo bar"
            _ = v8
        case 128:
            if false {
                v9 := GenS1{F0: 2147483647, F1: -125, F2: (p0 | (p0 | 0))}
                _ = v9
            }
        default:
            if true {
                for fi10 := 0; fi10 < 9; fi10++ {
                    p1 -= (^(p1))
                    v11 := GenS2{F0: 0}
                    var v12 bool = (!(-887 >= 1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                var v13 int8 = 1
                _ = v13
            }
        }
        var v14 uint32 = 42
        for fi15 := 0; fi15 < 1; fi15++ {
            v14 -= v14
            _ = fi15
        }
        _ = v14
    }
    if false {
        p0 -= p0
    } else {
        v16 := []int{(p0 % (((p0 + -9223372036854775808)) | 1))}
        _ = v16
    }
    p1 = (^(p1))
    switch "\n\t" {
    case "hello":
        if false {
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := GenS3{F0: (-9223372036854775808 <= -1)}
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := GenS4{F0: (p1 << ((802) & 63))}
                _ = fi19
                _ = v20
            }
        } else {
            var v21 int32 = -1
            fmt.Println(p0, p0)
            var v22 float64 = 1.5
            _ = v21
            _ = v22
        }
    case "foo bar":
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS3{F0: false}
            for fi25 := 0; fi25 < 1; fi25++ {
                if (0.0 > 0.1) {
                    v26 := []int{fi25}
                    p2 -= (-(p2))
                    _ = v26
                } else {
                    if ((1 != 65535) || false) {
                        if true {
                            v27 := []int{fi23, -2147483648, fi25, (fi25 / ((fi23) | 1))}
                            _ = v27
                        } else {
                            switch p0 {
                            case 9223372036854775807:
                                if false {
                                    if false {
                                        var v28 int32 = 0
                                        _ = v28
                                    } else {
                                        v29 := []int{(fi25 ^ -1), (fi23 & (p0 & p0)), -9223372036854775808, -2147483648}
                                        v30 := GenS2{F0: -52}
                                        var v31 int8 = -1
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    var v32 float32 = 1.5
                                    _ = v32
                                } else {
                                    fmt.Println(p1)
                                    p2 += -649
                                }
                            case 0:
                                v33 := GenS4{F0: p1}
                                _ = v33
                            case 128:
                                v34 := []int{fi23}
                                if (1 != 454) {
                                    p0 = p0
                                    p0 = 1
                                    p0 = (fi25 - p0)
                                }
                                _ = v34
                            default:
                                v35 := GenS4{F0: p1}
                                _ = v35
                            }
                        }
                        var v36 string = "foo bar"
                        if false {
                            fmt.Println(p0, v36, fi23)
                        } else {
                            v37 := GenS3{F0: false}
                            switch (fi23 + p0) {
                            case -382:
                                v38 := []int{}
                                _ = v38
                            case -2147483648:
                                switch p0 {
                                case -820:
                                    p2 += (p2 / (((p2 << ((42) & 63))) | 1))
                                case 620:
                                    var v39 int8 = -108
                                    _ = v39
                                case -1:
                                    if (!false) {
                                        if (42 < -128) {
                                            fmt.Println(p0)
                                            v40 := []int{(p0 << ((255) & 63)), fi23}
                                            fmt.Println(fi25, fi25)
                                            _ = v40
                                        }
                                        var v41 uint64 = 2
                                        _ = v41
                                    }
                                default:
                                    v42 := []int{(^(fi23))}
                                    _ = v42
                                }
                            case -372:
                                p1 = p1
                                var v43 int = (-(fi23))
                                p0 -= p0
                                _ = v43
                            }
                            switch (("\n\t" + v36) + "hello") {
                            case "hello":
                                var v44 int = fi25
                                _ = v44
                            case "a":
                                v45 := GenS3{F0: ((!((p1 >> ((920) & 63)) != 127)) || true)}
                                _ = v45
                            }
                            v46 := GenS3{F0: ((p2 & 1000) == (p2 - p2))}
                            _ = v37
                            _ = v46
                        }
                        fmt.Println(v36)
                        fmt.Println(p1, fi23)
                        _ = v36
                    } else {
                        p2 += 32767
                        v47 := GenS2{F0: 1}
                        _ = v47
                    }
                    fmt.Println(fi25, p1, p2)
                    v48 := []int{(fi23 - (-(fi23))), -172}
                    _ = v48
                }
                fmt.Println(p1, p2, p1)
                _ = fi25
            }
            v49 := GenS5{F0: 0, F1: true}
            _ = fi23
            _ = v24
            _ = v49
        }
    default:
        var v50 int64 = -358
        _ = v50
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 uint8, p1 uint64) uint {
    var v51 int8 = -1
    fmt.Println(p0)
    for fi52 := 0; fi52 < 3; fi52++ {
        v53 := []int{(fi52 & fi52), (^(fi52))}
        _ = fi52
        _ = v53
    }
    _ = p0
    _ = p1
    _ = v51
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 2
    var v3 bool = false
    v4 := []uint32{v1}
    v5 := []string{"", ("\n\t" + ""), "a"}
    var v6 int = 32768
    v54 := GenS4{F0: 42}
    switch (v6 * v6) {
    case 9223372036854775807:
        if (0 >= 1) {
            v55 := GenS6{F0: 127, F1: 1000}
            _ = v55
        }
    case -2147483648:
        v56 := GenS3{F0: (v3 && (v3 && false))}
        _ = v56
    case -776:
        for fi57 := 0; fi57 < 18; fi57++ {
            fmt.Println(v3)
            _ = fi57
        }
    default:
        var v58 uint8 = 128
        _ = v58
    }
    v2 -= (^(v2))
    fmt.Println(v6, v2, v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v59 uint8 = 127
    for fi60 := 0; fi60 < 1; fi60++ {
        v61 := GenS2{F0: -5}
        _ = fi60
        _ = v61
    }
    switch v6 {
    case 127:
        v62 := GenS4{F0: 128}
        _ = v62
    default:
        fmt.Println(v59, v2)
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        if (!(!(4294967295 >= 4294967295))) {
            var v64 uint32 = (v1 >> ((0) & 63))
            v5 = append(v5, (((("\n\t" + "foo bar") + (("foo bar" + "hello") + "a")) + "hello") + ("" + "hello")))
            _ = v64
        } else {
            for fi65 := 0; fi65 < 3; fi65++ {
                v3 = v3
                v3 = ((472 == 255) && true)
                var v66 uint64 = 42
                _ = fi65
                _ = v66
            }
            var v67 string = ("" + "a")
            _ = v67
        }
        _ = fi63
    }
    var v68 uint8 = v59
    if true {
        var v69 bool = v3
        _ = v69
    }
    fmt.Println(v3, v1, v6)
    var v70 float32 = 0.0
    if (314 <= 1000) {
        if v3 {
            var v71 int16 = 32767
            _ = v71
        } else {
            if v3 {
                var v72 float32 = float32(1e10)
                _ = v72
            }
            switch v6 {
            case -284:
                for fi73 := 0; fi73 < 2; fi73++ {
                    v74 := GenS2{F0: 127}
                    switch (("a" + "\n\t") + ("" + ("\n\t" + ("\n\t" + "foo bar")))) {
                    case "a":
                        var v75 int = v6
                        _ = v75
                    case "\n\t":
                        if v3 {
                            if (!(-976 < 1000)) {
                                fmt.Println(fi73, v1, v3)
                                var v76 int64 = 0
                                _ = v76
                            } else {
                                for fi77 := 0; fi77 < 16; fi77++ {
                                    fmt.Println(v70, v59, fi77)
                                    _ = fi77
                                }
                                v2 += v2
                            }
                        }
                    case "foo bar":
                        var v78 bool = true
                        _ = v78
                    }
                    _ = fi73
                    _ = v74
                }
            }
            var v79 int8 = -128
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v59
    _ = v6
    _ = v68
    _ = v70
}
