package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint64; F1 bool; F2 int64 }
func genFn0(p0 uint16, p1 bool, p2 uint64) int64 {
    fmt.Println(p2, p1)
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: p2, F1: p2}
        _ = fi5
        _ = v6
    }
    p0 = p0
    var v7 int64 = -1
    if ((2 <= 0) || (436 <= -2147483648)) {
        var v8 string = "hello"
        for fi9 := 0; fi9 < 18; fi9++ {
            if (p1 && ((!true) && p1)) {
                v10 := GenS2{F0: v7, F1: p1}
                _ = v10
            } else {
                v11 := []int{fi9, (fi9 * fi9), (fi9 - (fi9 ^ 128))}
                fmt.Println(v8, p2, p1)
                fmt.Println(p1, fi9)
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(v7))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 29
    v3 := []string{"foo bar", "hello", ("a" + ("\n\t" + "a")), "hello"}
    var v4 bool = ((3.14159 / -2.25) <= 0.0)
    for fi12 := 0; fi12 < 16; fi12++ {
        if v4 {
            fmt.Println(v4, fi12)
            if true {
                var v13 string = ("\n\t" + "")
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 float64 = ((0.0 / (1e10 + 0.1)) / float64(3.14159))
                    _ = len(v3)
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: 259, F1: 2}
                _ = v13
                _ = v16
            }
            var v17 uint64 = 2
            _ = v17
        } else {
            v18 := v1[fi12]
            if false {
                fmt.Println(v4)
                v19 := GenS3{F0: (v2 >> ((369) & 63))}
                _ = v19
            } else {
                fmt.Println(fi12, v2, v2)
                fmt.Println(v2, v2)
                v20 := GenS1{F0: 0, F1: 18446744073709551615}
                _ = v20
            }
            if (183 == 624) {
                v21 := GenS3{F0: (v2 | 100)}
                switch fi12 {
                case 903:
                    if ((!(v18 < v18)) || (-1 <= 0)) {
                        v1[fi12] = (v18 + 753)
                    }
                }
                var v22 uint = 859
                _ = v21
                _ = v22
            }
            _ = v18
        }
        switch "" {
        case "hello":
            for fi23 := 0; fi23 < 8; fi23++ {
                v2 -= (v2 & v2)
                var v24 uint = 42
                v2 -= v2
                v24 -= v24
                switch "foo bar" {
                case "\n\t":
                    v24 -= (^(v24))
                default:
                    fmt.Println(v24, fi12, fi23)
                }
                v2 = v2
                var v25 uint8 = 1
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            var v26 int16 = -991
            _ = v26
        }
        v2 += (v2 << ((855) & 63))
        _ = fi12
    }
    _ = len(v1)
    switch -843 {
    case 4294967295:
        switch "a" {
        case "foo bar":
            var v27 uint16 = 2
            _ = v27
        case "a":
            switch 42 {
            case -640:
                var v28 uint = 449
                _ = v28
            default:
                fmt.Println(v2, v2)
            }
        default:
            var v29 int = 4294967295
            _ = v29
        }
    case -895:
        var v30 uint32 = 1
        _ = v30
    case 2147483647:
        if v4 {
            if (10 > 1) {
                v31 := GenS4{F0: 0}
                _ = v31
            } else {
                var v32 int16 = -32768
                _ = v32
            }
            v4 = true
            v33 := GenS5{F0: (("\n\t" + "hello") + "a"), F1: 46}
            _ = v33
        } else {
            if v4 {
                if (4294967295 > 2147483647) {
                    switch "\n\t" {
                    case "":
                        for fi34 := 0; fi34 < 10; fi34++ {
                            switch fi34 {
                            case -888:
                                switch "foo bar" {
                                case "hello":
                                    var v35 uint8 = 0
                                    _ = v35
                                case "":
                                    v3 = append(v3, ("hello" + "foo bar"))
                                case "\n\t":
                                    var v36 uint32 = 1
                                    _ = v36
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v37 := GenS3{F0: -128}
                            _ = fi34
                            _ = v37
                        }
                    default:
                        v38 := v1[-987]
                        _ = v38
                    }
                    v2 = v2
                    v39 := GenS3{F0: v2}
                    _ = v39
                } else {
                    if v4 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        var v40 int64 = -945
                        fmt.Println(v40, v4, v2)
                        _ = v40
                    }
                }
                if true {
                    var v41 float32 = 0.0
                    var v42 float32 = v41
                    _ = v41
                    _ = v42
                } else {
                    delete(v1, 723)
                    var v43 string = ("foo bar" + "hello")
                    fmt.Println(v2)
                    _ = v43
                }
            } else {
                var v44 int32 = 22
                _ = v44
            }
            if (v4 && (v2 <= -128)) {
                v4 = false
                for fi45 := 0; fi45 < 3; fi45++ {
                    fmt.Println(v4)
                    var v46 bool = v4
                    _ = fi45
                    _ = v46
                }
                _ = len(v1)
            } else {
                v47 := GenS5{F0: ("\n\t" + "hello"), F1: 32767}
                _ = v47
            }
            v4 = (!(13 <= 1))
        }
    }
    v2 = 1
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 int32 = -1
        v2 -= (v2 ^ 0)
        delete(v1, -514)
        _ = fi48
        _ = v49
    }
    fmt.Println(v2)
    if false {
        var v50 float32 = (0.0 + 0.1)
        _ = v50
    } else {
        var v51 uint64 = 2
        _ = v51
    }
    v2 += -73
    if len(v3) > 0 {
        v3[0] = ("a" + "\n\t")
    }
    v3 = append(v3, "foo bar")
    fmt.Println(v2)
    v52 := GenS5{F0: ((("" + "\n\t") + "") + "a"), F1: 255}
    v53 := GenS3{F0: (v2 | 1)}
    v54 := GenS6{F0: 724, F1: true, F2: 9223372036854775807}
    if true {
        v55 := GenS6{F0: 18446744073709551615, F1: (4294967295 != -720), F2: 0}
        var v56 string = ("\n\t" + (("\n\t" + ("foo bar" + ("\n\t" + "a"))) + "a"))
        var v57 string = v56
        _ = v55
        _ = v56
        _ = v57
    } else {
        v58 := GenS2{F0: 9223372036854775807, F1: true}
        _ = v58
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        fmt.Println(v4, v2)
        fmt.Println(v2, v2)
        var v60 uint16 = 510
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 18; fi61++ {
        fmt.Println(v4, v2)
        _ = fi61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
