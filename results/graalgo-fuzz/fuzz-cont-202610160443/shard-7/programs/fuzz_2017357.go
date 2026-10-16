package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int; F1 int16; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint32 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        if (1 != 42) {
            v7 := []int{(fi6 ^ (fi6 * 400))}
            if ((!(!("hello" == "foo bar"))) || (false || true)) {
                var v8 uint16 = 255
                v8 += v8
                _ = v8
            } else {
                switch (fi6 % (((-(fi6))) | 1)) {
                case 2147483647:
                    if (1.5 != (-2.25 * float64((1e10 - -2.25)))) {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            fmt.Println(fi6)
                            var v10 uint8 = 1
                            _ = fi9
                            _ = v10
                        }
                        v11 := GenS1{F0: -101}
                        _ = v11
                    }
                case -1:
                    switch "foo bar" {
                    case "foo bar":
                        if (!(0 <= 2)) {
                            switch -674 {
                            case 1:
                                var v12 uint8 = 127
                                _ = v12
                            case -434:
                                _ = len(v7)
                            }
                        } else {
                            var v13 uint16 = 65535
                            var v14 int16 = -95
                            if ((("" + "\n\t") + ("a" + "a")) != "a") {
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v13 = (v13 % ((v13) | 1))
                                    var v16 uint16 = 255
                                    _ = len(v7)
                                    _ = fi15
                                    _ = v16
                                }
                            } else {
                                var v17 uint64 = 2
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    var v19 float64 = 0.1
                                    var v20 float32 = -2.25
                                    var v21 uint64 = 2
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = v17
                            }
                            _ = v13
                            _ = v14
                        }
                    }
                }
                var v22 float64 = 3.14159
                var v23 float64 = 1e10
                _ = v22
                _ = v23
            }
            _ = v7
        }
        var v24 int64 = 42
        _ = fi6
        _ = v24
    }
    var v25 string = "hello"
    v25 = (((("a" + "") + v25) + v25) + v25)
    v26 := GenS1{F0: -1}
    _ = v25
    _ = v26
    return 96
}
func genFn1(p0 int, p1 bool) uint8 {
    var v27 int = 32767
    v28 := GenS1{F0: v27}
    v29 := GenS1{F0: (p0 / (((v27 ^ 32768)) | 1))}
    var v30 uint16 = 42
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    return 220
}

func main() {
    var v1 float32 = (0.0 + 1.5)
    v2 := make(map[int]int)
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 0
    fmt.Println(v1)
    var v31 int64 = 9223372036854775807
    switch 592 {
    case -696:
        var v32 int8 = int8(v5)
        _ = v32
    case 2147483647:
        var v33 int16 = genFn0()
        _ = v33
    case 723:
        var v34 int64 = v31
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = 3.14159
        if ((1 < 0) && true) {
            v1 = v1
        } else {
            switch fi35 {
            case 32768:
                for fi36 := 0; fi36 < 8; fi36++ {
                    if ((false && false) && false) {
                        var v37 int8 = -26
                        v38 := GenS1{F0: (fi36 - fi36)}
                        if (!true) {
                            if (!(!false)) {
                                if ((42 > 42) && false) {
                                    var v39 uint16 = 65535
                                    var v40 uint64 = 4294967295
                                    v5 -= v5
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int32 = 862
                                    fmt.Println(v37)
                                    v31 = 1
                                    _ = v41
                                }
                                v5 -= v5
                            } else {
                                if true {
                                    var v42 int32 = -297
                                    if true {
                                        var v43 uint64 = 42
                                        _ = v43
                                    } else {
                                        if (42 == int16(v37)) {
                                            if false {
                                                v44 := GenS2{F0: 42}
                                                var v45 uint16 = 255
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v46 := GenS2{F0: 95}
                                                var v47 int16 = -547
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                        } else {
                                            if (-32768 < -605) {
                                                var v48 uint64 = 2
                                                v49 := GenS1{F0: fi36}
                                                v37 += (v37 * v37)
                                                _ = v48
                                                _ = v49
                                            }
                                            v50 := GenS1{F0: (fi35 + fi36)}
                                            _ = v50
                                        }
                                        _ = len(v3)
                                        v1 -= (float32(3.14159) + 3.14159)
                                    }
                                    delete(v2, fi35)
                                    _ = v42
                                } else {
                                    var v51 float64 = 3.14159
                                    _ = v51
                                }
                                fmt.Println(v37, v37, v5)
                            }
                            fmt.Println(fi35)
                        } else {
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v5)
                            } else {
                                if ((v5 < 1) && true) {
                                    if len(v3) > 0 {
                                        v3[0] = (fi35 | (^(fi36)))
                                    }
                                    if (42 < 0) {
                                        if false {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v1 -= (v1 * v1)
                                            var v52 uint = 528
                                            _ = v52
                                        }
                                        v5 = 2147483648
                                    }
                                } else {
                                    v1 = v1
                                    fmt.Println(v1, fi36, v31)
                                    _ = len(v3)
                                }
                                if (-32768 == -486) {
                                    v31 -= v31
                                } else {
                                    var v53 int32 = -906
                                    v5 = v5
                                    var v54 uint8 = 128
                                    _ = v53
                                    _ = v54
                                }
                                switch "foo bar" {
                                case "a":
                                    v55 := GenS3{F0: 2, F1: false, F2: 1}
                                    _ = len(v3)
                                    _ = v55
                                }
                            }
                            switch -659 {
                            case -48:
                                var v56 string = "a"
                                _ = v56
                            case 42:
                                v31 = 9223372036854775807
                                v37 = v37
                                if (1e10 <= 1.5) {
                                    var v57 uint64 = 1
                                    v37 = (v37 % ((v37) | 1))
                                    v31 += (v31 / ((v31) | 1))
                                    fmt.Println(fi35, v57, v31)
                                    _ = len(v3)
                                    _ = v57
                                } else {
                                    fmt.Println(v37)
                                }
                            }
                        }
                        _ = v37
                        _ = v38
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if false {
                        var v58 int64 = (v31 / (((v31 & v31)) | 1))
                        _ = v58
                    } else {
                        v31 = v31
                        v59 := GenS4{F0: 2, F1: 143}
                        _ = v59
                    }
                    _ = fi36
                }
            }
        }
        _ = fi35
    }
    for fi60 := 0; fi60 < 9; fi60++ {
        v61 := GenS3{F0: 181, F1: true, F2: genFn0()}
        v31 -= 42
        for fi62 := 0; fi62 < 17; fi62++ {
            var v63 uint = 463
            if true {
                fmt.Println(v63)
                v31 = (v31 | (v31 - 0))
            } else {
                v31 = v31
                if (v63 == 4294967295) {
                    v31 = v31
                    v4[v5] = (v5 * v5)
                    switch fi62 {
                    case 32768:
                        fmt.Println(v1)
                    case 30:
                        var v64 uint64 = 18446744073709551615
                        _ = v64
                    case -61:
                        var v65 float64 = 1e10
                        _ = v65
                    default:
                        var v66 int64 = (^(v31))
                        _ = v66
                    }
                }
                var v67 uint8 = 42
                _ = v67
            }
            _ = fi62
            _ = v63
        }
        _ = fi60
        _ = v61
    }
    switch 32767 {
    case 2147483647:
        v31 = v31
    case 32768:
        switch 0 {
        case 32768:
            if false {
                switch 4294967295 {
                case -135:
                    v1 -= (float32(-2.25) + -2.25)
                }
                if (!true) {
                    var v68 bool = (v31 < -9223372036854775808)
                    _ = v68
                }
                for fi69 := 0; fi69 < 13; fi69++ {
                    v70 := GenS5{F0: fi69, F1: 42, F2: 127}
                    _ = fi69
                    _ = v70
                }
            } else {
                switch -1 {
                case 0:
                    for fi71 := 0; fi71 < 1; fi71++ {
                        var v72 uint8 = 2
                        for fi73 := 0; fi73 < 12; fi73++ {
                            v31 = (v31 / (((v31 | (v31 | v31))) | 1))
                            var v74 uint16 = 80
                            _ = fi73
                            _ = v74
                        }
                        _ = fi71
                        _ = v72
                    }
                case -813:
                    v75 := GenS6{F0: (v31 + (v31 ^ (^(v31))))}
                    _ = v75
                default:
                    v31 = -9223372036854775808
                }
            }
        case 0:
            delete(v2, 0)
        }
    }
    for fi76 := 0; fi76 < 18; fi76++ {
        v77 := GenS7{F0: 42, F1: (v5 >> ((990) & 63))}
        var v78 uint = 196
        _ = fi76
        _ = v77
        _ = v78
    }
    v79 := GenS4{F0: 277, F1: 42}
    v1 -= (v1 * (1e10 + ((float32(1.5) - 0.0) * float32(3.14159))))
    var v80 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v79
    _ = v80
}
