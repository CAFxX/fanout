package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 string; F1 bool; F2 uint64 }

func main() {
    var v1 int64 = -1
    var v2 bool = (578 >= 2147483647)
    v3 := make(map[int]string)
    var v4 string = "\n\t"
    var v5 uint64 = 0
    v6 := make(map[uint32]int)
    var v7 float32 = ((0.0 / 0.1) / 1.0)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int64 = v1
        v10 := GenS1{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (v4 + ((v4 + v4) + v4)) {
    case "":
        var v11 int = -668
        _ = v11
    case "foo bar":
        if v2 {
            v12 := []int{-145, 42}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v12
        }
    case "a":
        var v13 float64 = 1.5
        _ = v13
    default:
        if (1 >= 0) {
            v4 = v4
        }
    }
    v4 = (v4 + "a")
    switch v4 {
    case "hello":
        for fi14 := 0; fi14 < 11; fi14++ {
            if true {
                if v2 {
                    v5 += 924
                }
                switch 9223372036854775807 {
                case 314:
                    v5 = 18446744073709551615
                case 128:
                    var v15 int64 = (v1 - (^(v1)))
                    _ = v15
                case 1:
                    v16 := GenS1{F0: 42}
                    _ = v16
                default:
                    if v2 {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            var v18 string = ("a" + ("foo bar" + "\n\t"))
                            var v19 uint16 = 32767
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi17 + fi17) {
                                case 0:
                                    var v20 int64 = (v1 & v1)
                                    _ = v20
                                case 128:
                                    var v21 int16 = 766
                                    _ = v21
                                }
                                if v2 {
                                    var v22 uint64 = 4294967295
                                    var v23 uint64 = v22
                                    v2 = true
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v24 int8 = 0
                                    v1 = v1
                                    _ = v24
                                }
                                var v25 uint64 = v5
                                _ = v25
                            case "hello":
                                switch fi14 {
                                case 4294967295:
                                    var v26 int = (^(fi14))
                                    _ = v26
                                }
                            case "":
                                fmt.Println(v19)
                            }
                            v5 -= v5
                            var v27 uint8 = 255
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v27
                        }
                    } else {
                        v28 := GenS2{F0: 9223372036854775807}
                        v2 = (0 <= (fi14 - -803))
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 int64 = (v1 - v1)
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                }
            } else {
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 int64 = v1
                    v3[1] = "a"
                    _ = fi31
                    _ = v32
                }
                if v2 {
                    v33 := GenS2{F0: v1}
                    _ = v33
                }
            }
            v34 := []int{32767, (fi14 ^ 9223372036854775807), fi14, 32768}
            if (false || false) {
                if true {
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v1 -= (v1 % (((v1 << ((445) & 63))) | 1))
                        fmt.Println(fi14, v7, fi14)
                        fmt.Println(v5)
                        _ = fi35
                    }
                    if false {
                        v4 = v4
                        for fi36 := 0; fi36 < 20; fi36++ {
                            v7 = (v7 - v7)
                            switch ("foo bar" + v4) {
                            case "a":
                                if (-2.25 <= 0.0) {
                                    if len(v34) > 0 {
                                        v34[0] = fi36
                                    }
                                    v5 = v5
                                }
                                if (-247 >= -508) {
                                    var v37 int32 = -600
                                    var v38 uint16 = 0
                                    v39 := GenS3{F0: (v38 >> ((0) & 63)), F1: v1, F2: (v4 + v4)}
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                v7 = v7
                                fmt.Println(v4, v2, v5)
                                v1 -= (^(v1))
                            case "foo bar":
                                var v40 bool = v2
                                _ = v40
                            case "":
                                v41 := GenS4{F0: 65535}
                                _ = v41
                            default:
                                _ = len(v3)
                            }
                            v42 := GenS4{F0: 0}
                            _ = len(v6)
                            _ = fi36
                            _ = v42
                        }
                    }
                    var v43 uint = 129
                    _ = v43
                } else {
                    fmt.Println(v7, v2, v4)
                }
                v5 += v5
            }
            _ = fi14
            _ = v34
        }
    }
    if (true && false) {
        var v44 int64 = (v1 % ((v1) | 1))
        _ = v44
    } else {
        v2 = false
        for fi45 := 0; fi45 < 19; fi45++ {
            var v46 uint64 = 18446744073709551615
            _ = fi45
            _ = v46
        }
    }
    v47 := []int{}
    var v48 uint16 = 1
    v49 := GenS5{F0: -944, F1: 1}
    v50 := v3[9223372036854775807]
    for fi51 := 0; fi51 < 8; fi51++ {
        v4 = "hello"
        for fi52 := 0; fi52 < 13; fi52++ {
            switch fi51 {
            case 1:
                fmt.Println(v50)
            }
            _ = fi52
        }
        _ = fi51
    }
    if v2 {
        v48 -= v48
    } else {
        for fi53 := 0; fi53 < 1; fi53++ {
            for fi54 := 0; fi54 < 12; fi54++ {
                switch ((v50 + (v50 + (v50 + "foo bar"))) + v50) {
                case "\n\t":
                    v2 = (42 > 42)
                case "":
                    v7 -= (v7 / v7)
                case "a":
                    v55 := v6[2147483647]
                    _ = v55
                }
                v56 := GenS6{F0: v50, F1: v2, F2: v5}
                v50 = ((v50 + (v50 + v50)) + v4)
                var v57 int = (fi54 >> ((756) & 63))
                _ = fi54
                _ = v56
                _ = v57
            }
            var v58 bool = v2
            _ = fi53
            _ = v58
        }
        for fi59 := 0; fi59 < 13; fi59++ {
            fmt.Println(v7)
            _ = fi59
        }
        switch 1 {
        case 0:
            v2 = true
        case 127:
            for fi60 := 0; fi60 < 14; fi60++ {
                v7 += v7
                var v61 int = 904
                for fi62 := 0; fi62 < 3; fi62++ {
                    var v63 uint64 = 42
                    var v64 uint32 = 0
                    _ = fi62
                    _ = v63
                    _ = v64
                }
                _ = fi60
                _ = v61
            }
        default:
            var v65 int32 = -1
            _ = v65
        }
    }
    v5 = v5
    for fi66 := 0; fi66 < 6; fi66++ {
        if (v2 || (float64(3.14159) >= -2.25)) {
            if len(v47) >= 2 {
                v47 = v47[1:len(v47)]
            }
            switch 42 {
            case 792:
                var v67 int = (fi66 ^ fi66)
                _ = v67
            }
        } else {
            var v68 uint32 = 42
            var v69 int32 = 1
            v70 := GenS3{F0: uint16(v68), F1: v1, F2: v50}
            _ = v68
            _ = v69
            _ = v70
        }
        var v71 int = -2147483648
        delete(v6, 65535)
        _ = fi66
        _ = v71
    }
    v5 += 18446744073709551615
    for fi72 := 0; fi72 < 7; fi72++ {
        _ = len(v47)
        _ = fi72
    }
    if v2 {
        for fi73 := 0; fi73 < 2; fi73++ {
            var v74 int16 = 42
            _ = fi73
            _ = v74
        }
        v50 = v4
        for fi75 := 0; fi75 < 10; fi75++ {
            v2 = v2
            v2 = (827 <= 42)
            _ = fi75
        }
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
