package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 bool; F2 string }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint; F1 int16; F2 int16 }

func main() {
    var v1 uint = 0
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    var v4 uint8 = uint8(v2)
    v5 := []int{95, 1}
    v6 := make(map[int]bool)
    var v7 uint8 = 128
    if false {
        v8 := v6[9223372036854775807]
        v8 = true
        var v9 int32 = -650
        _ = v8
        _ = v9
    } else {
        var v10 uint64 = (v2 + v2)
        var v11 uint64 = v10
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: false}
    v1 = 4294967295
    if true {
        v13 := GenS1{F0: true}
        _ = v13
    } else {
        v14 := GenS1{F0: false}
        for fi15 := 0; fi15 < 18; fi15++ {
            v4 -= v7
            switch ("" + "a") {
            case "foo bar":
                var v16 float32 = 1e10
                _ = v16
            case "hello":
                for fi17 := 0; fi17 < 20; fi17++ {
                    fmt.Println(v7)
                    delete(v3, 2)
                    v18 := v6[292]
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 uint8 = v4
                _ = v19
            }
            _ = fi15
        }
        if true {
            v7 = v7
            v3[2] = (("\n\t" + "a") + "")
        } else {
            var v20 uint16 = 65535
            if false {
                for fi21 := 0; fi21 < 7; fi21++ {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        fmt.Println(v1, fi21)
                        _ = fi22
                    }
                    _ = fi21
                }
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 uint8 = v4
                    delete(v3, 659)
                    v5 = append(v5, (fi23 + 127))
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 int64 = 0
                if (float32(-2.25) >= 3.14159) {
                    v3[1] = ("hello" + "hello")
                    v5 = append(v5, -9223372036854775808)
                } else {
                    v26 := v6[128]
                    _ = v26
                }
                _ = v25
            }
            if true {
                for fi27 := 0; fi27 < 15; fi27++ {
                    if ("\n\t" != "foo bar") {
                        _ = len(v3)
                        v1 = (v1 / (((v1 * 255)) | 1))
                        v1 = v1
                    } else {
                        if true {
                            for fi28 := 0; fi28 < 13; fi28++ {
                                var v29 bool = (65535 > v1)
                                _ = fi28
                                _ = v29
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            var v30 string = "a"
                            v31 := v3[832]
                            var v32 int64 = 1
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        if false {
                            if (505 >= -1) {
                                v33 := GenS2{F0: (v2 ^ 0)}
                                var v34 int8 = 37
                                var v35 uint = (v1 - (v1 >> (((v1 | (^(v1)))) & 63)))
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            v36 := GenS3{F0: 575, F1: "a"}
                            v37 := GenS4{F0: v7, F1: 1, F2: v2}
                            _ = v36
                            _ = v37
                        } else {
                            fmt.Println(v4)
                        }
                    }
                    v38 := GenS1{F0: false}
                    _ = fi27
                    _ = v38
                }
                if (0 <= 875) {
                    fmt.Println(v7, v1, v7)
                }
                if true {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v40 := GenS5{F0: (!("hello" < ("foo bar" + "\n\t")))}
                        switch ("\n\t" + "hello") {
                        case "hello":
                            switch ((((("a" + "a") + ("foo bar" + "hello")) + "\n\t") + "a") + "foo bar") {
                            case "a":
                                var v41 uint8 = uint8(v2)
                                _ = v41
                            default:
                                v42 := GenS2{F0: (v2 >> ((v1) & 63))}
                                _ = v42
                            }
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = (fi39 | fi39)
                            }
                        }
                        for fi43 := 0; fi43 < 2; fi43++ {
                            if ((fi43 ^ fi39) > fi43) {
                                if len(v5) > 0 {
                                    v5[0] = (fi43 - (fi39 ^ fi39))
                                }
                                if (1e10 >= (3.14159 + (((0.1 + (3.14159 + -2.25)) - ((-2.25 / (1.5 - (float64(0.0) - -2.25))) * -2.25)) / (1e10 * -2.25)))) {
                                    var v44 uint16 = v20
                                    _ = v44
                                }
                                var v45 float64 = (1.5 * ((0.1 * ((1.5 + ((0.1 + 0.1) * float64(3.14159))) + (0.0 + ((3.14159 + (float64(0.0) * 0.1)) - 0.0)))) + float64(0.1)))
                                _ = v45
                            }
                            v46 := GenS1{F0: false}
                            switch (fi39 & (-(fi43))) {
                            case -2147483648:
                                v7 = uint8(v1)
                            default:
                                var v47 int8 = 100
                                _ = v47
                            }
                            _ = fi43
                            _ = v46
                        }
                        _ = fi39
                        _ = v40
                    }
                    v2 = (v2 / (((^(v2))) | 1))
                    if ((!((0.0 / 3.14159) < -2.25)) && false) {
                        fmt.Println(v4, v20, v2)
                    }
                }
            } else {
                if false {
                    for fi48 := 0; fi48 < 19; fi48++ {
                        v2 = v2
                        _ = fi48
                    }
                    v49 := GenS3{F0: -404, F1: ""}
                    v50 := GenS6{F0: 42}
                    _ = v49
                    _ = v50
                } else {
                    if false {
                        if false {
                            v7 -= (^(v4))
                        }
                        v51 := GenS2{F0: (v2 + 0)}
                        if ((v2 >> (((v1 * v1)) & 63)) == (v2 | (v2 + v2))) {
                            if true {
                                for fi52 := 0; fi52 < 3; fi52++ {
                                    if true {
                                        fmt.Println(v20)
                                    }
                                    for fi53 := 0; fi53 < 18; fi53++ {
                                        v54 := GenS6{F0: (v2 / ((0) | 1))}
                                        switch (("hello" + "foo bar") + "\n\t") {
                                        case "":
                                            v20 = (^(v20))
                                        case "hello":
                                            var v55 int64 = -529
                                            _ = v55
                                        case "a":
                                            if (false && true) {
                                                var v56 uint = v1
                                                v1 = (v1 >> ((v56) & 63))
                                                _ = v56
                                            }
                                        }
                                        _ = fi53
                                        _ = v54
                                    }
                                    fmt.Println(fi52)
                                    _ = fi52
                                }
                                if true {
                                    for fi57 := 0; fi57 < 13; fi57++ {
                                        if (false || (((0.0 / 1.0) / 1.0) <= (0.0 * -2.25))) {
                                            var v58 int32 = 765
                                            v1 = (v1 - (v1 - (^(v1))))
                                            _ = v58
                                        }
                                        _ = fi57
                                    }
                                } else {
                                    switch "a" {
                                    case "a":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "hello":
                                        v59 := GenS6{F0: uint64(v7)}
                                        _ = v59
                                    default:
                                        var v60 bool = ((v2 >> ((v1) & 63)) >= v2)
                                        _ = v60
                                    }
                                    fmt.Println(v4, v4, v4)
                                }
                                v61 := GenS7{F0: 724, F1: (int32(v1) >= 2147483647), F2: "a"}
                                _ = v61
                            }
                            _ = len(v6)
                            var v62 uint16 = (v20 | 0)
                            _ = v62
                        }
                        _ = v51
                    }
                }
            }
            _ = v20
        }
        _ = v14
    }
    fmt.Println(v4)
    v1 -= 255
    _ = len(v6)
    v7 = (v7 << ((uint(v4)) & 63))
    v63 := GenS8{F0: 41, F1: 0}
    fmt.Println(v7)
    if ((true && (float32(1e10) != float32(1.5))) || ("" != "\n\t")) {
        var v64 int16 = -1
        v2 = v2
        var v65 float64 = float64(1.5)
        _ = v64
        _ = v65
    }
    if true {
        if ((127 >= -23) || (false || false)) {
            v7 += (v4 >> ((uint(v2)) & 63))
            v66 := GenS9{F0: -1}
            _ = v66
        }
    }
    v67 := v3[65535]
    v68 := GenS10{F0: -417, F1: true}
    if ((v7 % ((2) | 1)) > (^(v4))) {
        for fi69 := 0; fi69 < 6; fi69++ {
            if ((29 != 32767) || true) {
                var v70 uint16 = 1
                if (false || (!false)) {
                    for fi71 := 0; fi71 < 13; fi71++ {
                        var v72 float64 = 3.14159
                        _ = fi71
                        _ = v72
                    }
                    var v73 int = (fi69 - fi69)
                    _ = v73
                } else {
                    v67 = v67
                    var v74 uint16 = (v70 >> ((uint(v4)) & 63))
                    switch fi69 {
                    case 32768:
                        var v75 uint32 = 327
                        _ = v75
                    case 2147483647:
                        v4 = (^(v7))
                    }
                    _ = v74
                }
                var v76 int32 = -55
                _ = v70
                _ = v76
            } else {
                if false {
                    v77 := GenS11{F0: (v1 | (^(v1))), F1: 0, F2: 517}
                    _ = v77
                } else {
                    delete(v3, 2147483648)
                    var v78 int8 = 39
                    for fi79 := 0; fi79 < 16; fi79++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v80 := GenS3{F0: 163, F1: ("foo bar" + "foo bar")}
                        v1 = 42
                        _ = fi79
                        _ = v80
                    }
                    _ = v78
                }
            }
            _ = fi69
        }
        fmt.Println(v7, v2, v2)
        v4 = (v7 + (v4 / (((^(v7))) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v67
    _ = v68
    _ = v7
}
