package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (3.14159 + (0.1 + 0.1))
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v6 uint8 = 0
    switch ((("\n\t" + "") + ("\n\t" + "\n\t")) + ("hello" + "")) {
    case "\n\t":
        for fi7 := 0; fi7 < 7; fi7++ {
            v1 = 1e10
            _ = fi7
        }
    case "":
        v2 -= (((v2 + 1e10) * v2) + (v2 + v2))
    case "a":
        if true {
            switch -999 {
            case 4294967295:
                for fi8 := 0; fi8 < 8; fi8++ {
                    if false {
                        var v9 uint8 = (v6 << ((1) & 63))
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v11 := GenS1{F0: true, F1: "foo bar", F2: -92}
                            _ = fi10
                            _ = v11
                        }
                        _ = v9
                    } else {
                        v2 = (v2 - 0.1)
                        var v12 float32 = (v1 / v1)
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v12 -= 1.5
                            switch (fi13 + (fi13 & fi13)) {
                            case -215:
                                v14 := []int{}
                                _ = v14
                            }
                            switch fi13 {
                            case 1:
                                v12 -= v1
                            case -223:
                                v6 = v6
                            case -1:
                                v15 := GenS1{F0: false, F1: "", F2: 42}
                                _ = v15
                            }
                            _ = fi13
                        }
                        _ = v12
                    }
                    _ = fi8
                }
            case -2147483648:
                switch ((((("" + ("a" + ("" + "a"))) + "") + (("" + "\n\t") + "")) + "foo bar") + ((((("foo bar" + ("" + "\n\t")) + ("foo bar" + "\n\t")) + "") + "foo bar") + "")) {
                case "":
                    var v16 uint64 = 707
                    _ = v16
                }
            case 128:
                switch "foo bar" {
                case "a":
                    var v17 int8 = 1
                    _ = v17
                case "":
                    fmt.Println(v6, v1)
                default:
                    _ = len(v4)
                }
            }
            var v18 uint32 = 573
            v18 += (v18 << ((2) & 63))
            _ = v18
        }
    }
    v1 = 0.1
    switch "foo bar" {
    case "a":
        var v19 uint32 = 4294967295
        _ = v19
    case "":
        var v20 int32 = 42
        _ = v20
    default:
        var v21 int8 = 102
        _ = v21
    }
    fmt.Println(v2, v2, v6)
    var v22 int8 = 100
    v6 = v6
    var v23 uint16 = 255
    v6 = (v6 - v6)
    if true {
        v2 -= 0.1
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int32 = 1
            _ = fi24
            _ = v25
        }
    } else {
        v26 := GenS1{F0: (false || (!true)), F1: ("foo bar" + ("foo bar" + "a")), F2: (-(v22))}
        for fi27 := 0; fi27 < 9; fi27++ {
            if ((v2 - v2) == -2.25) {
                switch "" {
                case "":
                    fmt.Println(v2, v1)
                case "foo bar":
                    if false {
                        v2 = 1.5
                        var v28 int32 = -1
                        _ = v28
                    }
                }
            } else {
                for fi29 := 0; fi29 < 1; fi29++ {
                    switch 266 {
                    case 129:
                        v3[(fi27 & fi29)] = 42
                    case 32767:
                        v30 := GenS2{F0: (v6 >> ((2) & 63)), F1: -32768, F2: -379}
                        switch fi27 {
                        case -215:
                            fmt.Println(v2, v2)
                        }
                        v2 = v2
                        var v31 string = "\n\t"
                        if false {
                            switch (v31 + "foo bar") {
                            case "hello":
                                var v32 uint16 = v23
                                _ = v32
                            case "\n\t":
                                v33 := GenS1{F0: (v6 >= 255), F1: v31, F2: -109}
                                _ = v33
                            default:
                                v34 := v4[""]
                                _ = v34
                            }
                            if ((155 == 497) && true) {
                                var v35 float32 = 0.1
                                v35 = 1e10
                                _ = v35
                            } else {
                                switch v31 {
                                case "foo bar":
                                    v31 = "hello"
                                    _ = len(v5)
                                }
                                var v36 uint16 = (v23 / ((v23) | 1))
                                if false {
                                    var v37 int8 = v22
                                    var v38 float32 = ((v1 * v1) * 0.0)
                                    _ = v37
                                    _ = v38
                                }
                                switch fi27 {
                                case 2147483647:
                                    var v39 uint16 = (^(v23))
                                    _ = v39
                                case 32768:
                                    var v40 string = "a"
                                    _ = v40
                                case -1:
                                    v41 := GenS3{F0: uint16(v6)}
                                    _ = v41
                                }
                                v22 = v22
                                _ = v36
                            }
                            v42 := GenS4{F0: 0, F1: -96, F2: 720}
                            _ = v42
                        } else {
                            var v43 uint64 = 18446744073709551615
                            var v44 bool = false
                            v45 := []int{(-(fi27))}
                            fmt.Println(v22)
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        if ((v6 / ((128) | 1)) == 2) {
                            v2 += float64(v1)
                            if (714 >= 9223372036854775807) {
                                if (!(!(-923 < -62))) {
                                    v3[fi27] = 644
                                    fmt.Println(v6, v23)
                                    v23 -= v23
                                    v2 += 1.5
                                } else {
                                    v46 := v5[(v31 + "")]
                                    fmt.Println(v6, fi29, fi27)
                                    v1 = (v1 / 1.0)
                                    _ = v46
                                }
                            }
                            var v47 int64 = -9223372036854775808
                            _ = v47
                        } else {
                            if (!(((-(fi29)) >= fi29) && false)) {
                                var v48 uint64 = 42
                                v31 = ("\n\t" + v31)
                                _ = v48
                            }
                            if (false || true) {
                                if false {
                                    v23 += 255
                                } else {
                                    if false {
                                        fmt.Println(v1)
                                        delete(v5, "")
                                        v31 = v31
                                    } else {
                                        v22 = v22
                                    }
                                    v22 += 1
                                    fmt.Println(v23)
                                }
                                v49 := GenS5{F0: (v22 / ((127) | 1)), F1: 2, F2: 42}
                                var v50 uint32 = 0
                                _ = v49
                                _ = v50
                            } else {
                                v51 := v4["hello"]
                                _ = v51
                            }
                            var v52 int64 = -1
                            _ = v52
                        }
                        _ = v30
                        _ = v31
                    }
                    v2 = (v2 + 0.0)
                    v53 := v3[(fi27 & (fi27 << ((491) & 63)))]
                    v53 = (v53 * (v53 | v53))
                    var v54 uint32 = 947
                    v55 := GenS5{F0: v22, F1: 191, F2: (v6 / (((^(v6))) | 1))}
                    v56 := v4[("a" + "a")]
                    _ = fi29
                    _ = v53
                    _ = v54
                    _ = v55
                    _ = v56
                }
            }
            fmt.Println(v22, v22)
            for fi57 := 0; fi57 < 6; fi57++ {
                if false {
                    var v58 float64 = v2
                    var v59 uint8 = v6
                    _ = v58
                    _ = v59
                }
                if false {
                    v60 := []int{int(v23), 2147483647, (fi27 + 42)}
                    _ = v60
                } else {
                    v22 = (^(v22))
                    v61 := GenS6{F0: 864}
                    _ = v61
                }
                _ = fi57
            }
            _ = fi27
        }
        _ = v26
    }
    v23 -= v23
    for fi62 := 0; fi62 < 2; fi62++ {
        for fi63 := 0; fi63 < 6; fi63++ {
            fmt.Println(v6, v1)
            if false {
                v64 := GenS7{F0: true, F1: 42}
                _ = v64
            }
            var v65 int64 = 1
            _ = fi63
            _ = v65
        }
        _ = len(v4)
        v66 := GenS8{F0: 2, F1: 128, F2: false}
        _ = fi62
        _ = v66
    }
    _ = len(v4)
    var v67 string = "a"
    if true {
        v2 = (((v2 / -2.25) - 0.0) / ((0.0 * ((float64(0.1) + (((-2.25 / (v2 + v2)) + float64((v1 + float32((0.0 / ((v2 - v2) - v2)))))) - v2)) / 1.0)) * v2))
    }
    var v68 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
