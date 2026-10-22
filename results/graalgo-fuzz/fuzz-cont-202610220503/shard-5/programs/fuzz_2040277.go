package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS5 struct { F0 int64; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0() string {
    if true {
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        var v8 float64 = (0.1 - -2.25)
        v9 := GenS2{F0: 127, F1: 32767, F2: -1}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: 1}
        var v12 float32 = 0.1
        if (!true) {
            var v13 int32 = 0
            var v14 bool = (!(32767 == 32768))
            _ = v13
            _ = v14
        } else {
            fmt.Println(v12, fi10)
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v15 int32 = 42
    _ = v15
    return "foo bar"
}
func genFn1(p0 int64, p1 bool) string {
    for fi16 := 0; fi16 < 2; fi16++ {
        p0 += (p0 + (p0 & p0))
        var v17 uint16 = 2
        if ("hello" < "hello") {
            if (true || (41 > 255)) {
                v18 := []int{fi16}
                switch "hello" {
                case "":
                    var v19 uint16 = (v17 & v17)
                    _ = v19
                default:
                    var v20 uint16 = (^(v17))
                    _ = v20
                }
                _ = v18
            } else {
                v21 := []int{(fi16 & int(p0)), fi16}
                _ = v21
            }
            var v22 string = "\n\t"
            _ = v22
        }
        _ = fi16
        _ = v17
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        p0 = 1
        switch fi23 {
        case 0:
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 float32 = 0.0
                var v26 uint8 = 2
                v25 = v25
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = fi23
    }
    v27 := GenS1{F0: -1}
    v28 := GenS2{F0: 128, F1: 32768, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return "a"
}

func main() {
    var v1 int8 = 56
    v2 := make(map[int]uint32)
    var v3 string = ("hello" + ("hello" + ((("\n\t" + "\n\t") + (("hello" + "hello") + "\n\t")) + "foo bar")))
    v4 := []string{"a", v3, ("a" + (v3 + v3))}
    var v5 int8 = v1
    var v6 bool = false
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int8 = (v1 << ((151) & 63))
        if v6 {
            _ = len(v2)
        } else {
            var v31 uint16 = 65535
            fmt.Println(v3, v31)
            var v32 int64 = -94
            _ = v31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v6, v5)
    v3 = ((("" + v3) + "") + "")
    var v33 bool = v6
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v34 int8 = 54
    var v35 uint64 = 471
    v5 -= (v5 / ((100) | 1))
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 float64 = ((-2.25 - 3.14159) * 1e10)
        _ = fi36
        _ = v37
    }
    v5 -= v34
    _ = len(v4)
    switch "" {
    case "hello":
        if (!(!true)) {
            v38 := GenS2{F0: 128, F1: 797, F2: 4294967295}
            if (1e10 == (1.5 - -2.25)) {
                var v39 uint8 = 67
                if v6 {
                    var v40 float32 = 3.14159
                    switch 32768 {
                    case 4294967295:
                        v3 = "\n\t"
                    default:
                        switch -2147483648 {
                        case 780:
                            v4 = append(v4, v3)
                        case -622:
                            fmt.Println(v35, v34)
                        case -477:
                            v41 := GenS3{F0: -259, F1: 255, F2: (v34 & v5)}
                            _ = v41
                        default:
                            v2[0] = 2147483648
                        }
                    }
                    _ = v40
                }
                for fi42 := 0; fi42 < 5; fi42++ {
                    if len(v4) > 0 {
                        v4[0] = ("" + "\n\t")
                    }
                    v6 = v33
                    for fi43 := 0; fi43 < 16; fi43++ {
                        fmt.Println(v3)
                        _ = fi43
                    }
                    _ = fi42
                }
                _ = v39
            } else {
                if v33 {
                    v44 := GenS3{F0: -1, F1: 2, F2: (v1 / ((100) | 1))}
                    if (!(!(2 >= 42))) {
                        for fi45 := 0; fi45 < 7; fi45++ {
                            _ = len(v2)
                            if v6 {
                                v46 := GenS3{F0: 42, F1: 0, F2: 5}
                                var v47 int32 = 1
                                _ = v46
                                _ = v47
                            }
                            _ = fi45
                        }
                        for fi48 := 0; fi48 < 10; fi48++ {
                            if (255 < 84) {
                                var v49 int = (fi48 >> ((658) & 63))
                                _ = v49
                            }
                            _ = fi48
                        }
                    }
                    _ = v44
                }
                var v50 uint64 = v35
                if v6 {
                    v50 = v50
                    var v51 uint = 42
                    v3 = (v3 + (((v3 + ((v3 + v3) + (v3 + (v3 + ((((v3 + v3) + v3) + "foo bar") + ("" + "hello")))))) + "a") + v3))
                    _ = v51
                } else {
                    if v6 {
                        v52 := GenS4{F0: 42, F1: 1, F2: 152}
                        if false {
                            var v53 uint8 = 0
                            v54 := GenS2{F0: 0, F1: 42, F2: -914}
                            _ = v53
                            _ = v54
                        } else {
                            v55 := GenS2{F0: 255, F1: -1, F2: 9223372036854775807}
                            switch 190 {
                            case 32768:
                                if true {
                                    fmt.Println(v6, v6, v35)
                                    v5 = v1
                                } else {
                                    v35 = v35
                                }
                            }
                            _ = v55
                        }
                        v3 = v3
                        _ = v52
                    } else {
                        if v33 {
                            for fi56 := 0; fi56 < 18; fi56++ {
                                var v57 uint8 = 128
                                _ = fi56
                                _ = v57
                            }
                            var v58 int8 = (v5 % ((v5) | 1))
                            _ = v58
                        }
                        v33 = true
                        var v59 string = (("a" + v3) + "hello")
                        _ = v59
                    }
                    var v60 float32 = -2.25
                    _ = v60
                }
                _ = v50
            }
            v33 = (1e10 != (1.5 + -2.25))
            _ = v38
        }
    case "foo bar":
        v2[42] = 2
    case "":
        switch v3 {
        case "":
            for fi61 := 0; fi61 < 17; fi61++ {
                _ = len(v4)
                _ = fi61
            }
        }
    }
    for fi62 := 0; fi62 < 4; fi62++ {
        if (!(v3 != v3)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v63 int = 4294967295
            var v64 int64 = 0
            _ = v63
            _ = v64
        }
        v33 = true
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi62
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        fmt.Println(fi65, v3, v33)
        _ = len(v2)
        for fi66 := 0; fi66 < 11; fi66++ {
            v67 := GenS2{F0: 91, F1: (fi65 % ((fi65) | 1)), F2: fi65}
            _ = fi66
            _ = v67
        }
        _ = fi65
    }
    if v6 {
        fmt.Println(v33)
    } else {
        switch 462 {
        case 0:
            if (true || false) {
                v68 := GenS5{F0: -269, F1: -323, F2: 2147483647}
                v34 -= int8(v35)
                _ = v68
            } else {
                v69 := GenS2{F0: 186, F1: -803, F2: 127}
                _ = v69
            }
        case 9223372036854775807:
            switch v3 {
            case "a":
                v70 := GenS6{F0: 2, F1: 0, F2: -95}
                _ = v70
            default:
                if (!((v1 | v34) != (v34 % (((v34 - (v5 & v34))) | 1)))) {
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v35 -= 42
                        v35 = 2
                        var v72 uint64 = (v35 - (v35 / ((v35) | 1)))
                        _ = fi71
                        _ = v72
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
