package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0() uint8 {
    v7 := GenS1{F0: 0, F1: 133}
    v8 := GenS1{F0: 2147483648, F1: 58}
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 6; fi10++ {
        if false {
            fmt.Println(fi10, v9)
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := []int{}
                v9 = (v9 << ((255) & 63))
                v9 -= v9
                switch "hello" {
                case "hello":
                    if (207 >= 4294967295) {
                        if false {
                            if (!((-42 >= -1) && false)) {
                                if true {
                                    var v13 int16 = 32767
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    _ = v13
                                } else {
                                    v14 := GenS2{F0: (-(v9)), F1: 2}
                                    if (!false) {
                                        v9 = 1
                                    }
                                    var v15 int16 = 846
                                    if true {
                                        var v16 bool = false
                                        v15 -= (v15 | (v15 * v15))
                                        v17 := GenS2{F0: (v9 << ((4294967295) & 63)), F1: 18446744073709551615}
                                        _ = v16
                                        _ = v17
                                    } else {
                                        if true {
                                            v9 -= (v9 - 42)
                                            v15 = v15
                                        }
                                        var v18 uint64 = 1
                                        v15 = 1
                                        _ = v18
                                    }
                                    v15 = v15
                                    if true {
                                        var v19 int = fi10
                                        var v20 int16 = 539
                                        var v21 uint64 = 561
                                        if (100 <= 1) {
                                            v22 := GenS3{F0: -128, F1: 255, F2: (1 <= 983)}
                                            _ = v22
                                        }
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                    fmt.Println(fi11)
                                    _ = v14
                                    _ = v15
                                }
                                v23 := GenS4{F0: 255, F1: ("hello" + (("hello" + (("" + (("a" + "\n\t") + ("hello" + "\n\t"))) + ("hello" + "a"))) + "\n\t"))}
                                var v24 uint64 = 1
                                v25 := GenS4{F0: 238, F1: ""}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v9 = -1
                        } else {
                            var v26 uint16 = 255
                            if (true && false) {
                                var v27 string = ("a" + ("hello" + "a"))
                                _ = v27
                            }
                            var v28 int64 = (v9 * v9)
                            _ = v26
                            _ = v28
                        }
                    } else {
                        v9 += 0
                        var v29 int64 = (v9 / ((v9) | 1))
                        _ = v29
                    }
                case "a":
                    fmt.Println(fi11)
                default:
                    var v30 uint8 = 0
                    _ = v30
                }
                _ = fi11
                _ = v12
            }
            v31 := []int{fi10, fi10}
            _ = v31
        }
        v32 := []int{}
        _ = fi10
        _ = v32
    }
    var v33 string = "hello"
    _ = v33
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = -1
    var v2 int8 = 100
    v3 := make(map[int]string)
    var v4 uint16 = 88
    v5 := make(map[uint32]int)
    v6 := []string{("\n\t" + "a"), "a"}
    if false {
        fmt.Println(v1, v2)
        v2 -= 0
        fmt.Println(v4, v2, v2)
    }
    if true {
        var v34 uint16 = 212
        _ = v34
    } else {
        var v35 uint64 = uint64(v4)
        v2 = (v2 ^ 1)
        if (!true) {
            v1 = (-(v1))
        } else {
            var v36 int8 = -128
            _ = v36
        }
        _ = v35
    }
    var v37 float32 = (((-2.25 / -2.25) + -2.25) + 0.0)
    delete(v5, 4294967295)
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            delete(v5, 4294967295)
            v40 := GenS1{F0: 2, F1: genFn0()}
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    if (1e10 == v37) {
        v2 = v2
    } else {
        v4 = (v4 >> ((856) & 63))
        if (v1 != (v1 ^ v1)) {
            v41 := GenS3{F0: v2, F1: v4, F2: ((v4 & (v4 | v4)) <= (v4 | 1))}
            var v42 uint8 = 42
            _ = v41
            _ = v42
        } else {
            switch "\n\t" {
            case "\n\t":
                v4 = v4
            case "a":
                v6 = append(v6, "")
            }
            v2 = v2
        }
    }
    if true {
        v4 = (v4 | 116)
        var v43 float32 = (v37 - v37)
        _ = v43
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v4)
    v4 -= 32767
    if true {
        if (int(v1) != 248) {
            v44 := GenS3{F0: 100, F1: v4, F2: false}
            _ = v44
        } else {
            v6 = append(v6, "foo bar")
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if true {
                for fi45 := 0; fi45 < 6; fi45++ {
                    v46 := GenS2{F0: 1, F1: 1}
                    fmt.Println(v1, v37)
                    v4 += 462
                    _ = fi45
                    _ = v46
                }
                v37 = 0.1
                if (v37 != 0.0) {
                    v2 = 42
                } else {
                    v37 = 1e10
                    if false {
                        _ = len(v6)
                        v2 += (-(v2))
                        v1 = (v1 ^ -92)
                    } else {
                        for fi47 := 0; fi47 < 16; fi47++ {
                            for fi48 := 0; fi48 < 1; fi48++ {
                                v2 -= (v2 ^ -1)
                                var v49 int8 = (v2 >> ((42) & 63))
                                var v50 int64 = int64(v1)
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v37 -= ((v37 - v37) + v37)
                            var v51 int32 = v1
                            _ = fi47
                            _ = v51
                        }
                    }
                }
            } else {
                var v52 uint16 = 523
                if false {
                    v37 -= 1e10
                    v1 -= (v1 + v1)
                    if true {
                        var v53 float64 = (-2.25 + 0.0)
                        v3[59] = "\n\t"
                        _ = v53
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi54 := 0; fi54 < 13; fi54++ {
                        v55 := GenS3{F0: (v2 - (v2 / ((-1) | 1))), F1: (v4 - 65535), F2: true}
                        _ = fi54
                        _ = v55
                    }
                }
                _ = v52
            }
        }
        v37 = v37
    }
    if ((v1 / (((^(v1))) | 1)) == v1) {
        var v56 uint64 = 677
        var v57 uint8 = 128
        if true {
            fmt.Println(v57, v57, v1)
        } else {
            switch "foo bar" {
            case "foo bar":
                v56 = v56
            case "":
                if len(v6) > 0 {
                    v6[0] = "a"
                }
            case "hello":
                _ = len(v3)
            default:
                v58 := GenS5{F0: (v1 ^ -919), F1: (v57 / ((1) | 1)), F2: -32768}
                _ = v58
            }
        }
        _ = v56
        _ = v57
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        var v60 float32 = 0.1
        v4 += (v4 & v4)
        switch fi59 {
        case 32767:
            v60 = ((0.1 * v60) * v60)
        default:
            var v61 uint16 = 255
            _ = v61
        }
        _ = fi59
        _ = v60
    }
    var v62 int = 634
    v2 += v2
    v4 += v4
    switch v62 {
    case 0:
        var v63 int16 = -32768
        _ = v63
    default:
        v64 := GenS2{F0: 1, F1: 1}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
