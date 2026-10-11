package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 string) int64 {
    v7 := GenS1{F0: 0, F1: 2, F2: 108}
    v8 := []int{-1, -9223372036854775808, -1, 14}
    v9 := GenS2{F0: 42}
    p0 -= 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() int16 {
    if (223 != 0) {
        switch 0 {
        case 42:
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int = -267
                if (true || (!true)) {
                    v12 := GenS1{F0: 255, F1: 1, F2: 2}
                    if (1 == 872) {
                        if (42 <= 42) {
                            if true {
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        v11 = 0
                                    case "a":
                                        fmt.Println(v11)
                                    case "hello":
                                        var v14 uint32 = 506
                                        _ = v14
                                    }
                                    switch ("hello" + "hello") {
                                    case "foo bar":
                                        v11 += v11
                                    case "hello":
                                        var v15 int32 = -754
                                        _ = v15
                                    }
                                    _ = fi13
                                }
                            } else {
                                var v16 uint8 = 2
                                _ = v16
                            }
                            var v17 string = (("hello" + "\n\t") + "foo bar")
                            for fi18 := 0; fi18 < 4; fi18++ {
                                var v19 uint8 = 48
                                if true {
                                    v19 = (^(v19))
                                    if true {
                                        var v20 int = v11
                                        _ = v20
                                    } else {
                                        fmt.Println(v19)
                                        if false {
                                            fmt.Println(v11)
                                            var v21 float32 = -2.25
                                            fmt.Println(v21, v19, v21)
                                            _ = v21
                                        } else {
                                            v22 := GenS3{F0: -293}
                                            _ = v22
                                        }
                                        v23 := []int{2147483647}
                                        _ = v23
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        }
                        v11 = (^(v11))
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v25 := GenS4{F0: 633, F1: 42, F2: 4294967295}
                            v11 -= (fi10 & v11)
                            if (2 < 1) {
                                v26 := GenS5{F0: 1, F1: 80, F2: 2147483647}
                                v11 -= v11
                                var v27 int8 = 127
                                _ = v26
                                _ = v27
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = v12
                }
                fmt.Println(v11)
                _ = fi10
                _ = v11
            }
        case 1:
            var v28 uint32 = 0
            _ = v28
        case -1:
            for fi29 := 0; fi29 < 15; fi29++ {
                if (2 == 127) {
                    v30 := []int{fi29, fi29}
                    _ = v30
                } else {
                    var v31 int = (fi29 * 366)
                    v31 += fi29
                    _ = v31
                }
                switch fi29 {
                case -97:
                    var v32 float32 = float32(((1.5 / -2.25) - -2.25))
                    _ = v32
                case 127:
                    if (1.5 < 1.5) {
                        switch "foo bar" {
                        case "hello":
                            var v33 uint = 255
                            _ = v33
                        }
                        v34 := GenS3{F0: -80}
                        var v35 float64 = float64((3.14159 - 1e10))
                        _ = v34
                        _ = v35
                    }
                }
                _ = fi29
            }
        }
        var v36 uint32 = 1
        for fi37 := 0; fi37 < 1; fi37++ {
            var v38 int16 = -32768
            _ = fi37
            _ = v38
        }
        _ = v36
    } else {
        v39 := GenS2{F0: 901}
        var v40 uint64 = 129
        _ = v39
        _ = v40
    }
    if true {
        var v41 float32 = 3.14159
        _ = v41
    } else {
        var v42 string = "\n\t"
        _ = v42
    }
    var v43 string = (("foo bar" + "") + "\n\t")
    _ = v43
    return -1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 float32 = (1.5 / 1.5)
    v5 := []string{"foo bar", "foo bar"}
    v6 := []string{"hello", "\n\t"}
    v1 -= 42
    if false {
        v44 := GenS6{F0: 65535, F1: -1, F2: 1}
        _ = v44
    }
    v45 := GenS7{F0: 0}
    switch -253 {
    case 9223372036854775807:
        var v46 uint32 = v1
        _ = v46
    }
    v5 = append(v5, "a")
    var v47 int16 = 943
    delete(v2, 276)
    for fi48 := 0; fi48 < 18; fi48++ {
        if (v3 && v3) {
            switch 9223372036854775807 {
            case -1:
                v47 -= (v47 * 868)
            default:
                switch fi48 {
                case 0:
                    v1 += (^(v1))
                case 32767:
                    fmt.Println(v3, v47)
                case 88:
                    v3 = v3
                default:
                    var v49 string = "\n\t"
                    _ = v49
                }
            }
        } else {
            var v50 uint64 = 4294967295
            v5 = append(v5, "a")
            v1 += (v1 << ((255) & 63))
            _ = v50
        }
        _ = fi48
    }
    for fi51 := 0; fi51 < 4; fi51++ {
        switch (fi51 | (fi51 / ((2147483647) | 1))) {
        case 375:
            var v52 uint = 255
            _ = v52
        case -1:
            for fi53 := 0; fi53 < 8; fi53++ {
                v54 := v2[v1]
                v55 := GenS5{F0: 2147483647, F1: 0, F2: (v1 >> ((489) & 63))}
                _ = fi53
                _ = v54
                _ = v55
            }
        default:
            v3 = v3
        }
        v56 := GenS2{F0: 0}
        var v57 uint8 = 20
        _ = fi51
        _ = v56
        _ = v57
    }
    v58 := GenS3{F0: -1}
    v59 := GenS8{F0: -952, F1: "a"}
    var v60 bool = true
    for fi61 := 0; fi61 < 1; fi61++ {
        switch fi61 {
        case 882:
            _ = len(v2)
        case 2147483647:
            if (3.14159 < (1e10 + 0.0)) {
                v62 := GenS5{F0: 0, F1: 2, F2: v1}
                for fi63 := 0; fi63 < 1; fi63++ {
                    switch ("hello" + "hello") {
                    case "\n\t":
                        delete(v2, (v1 & v1))
                    case "hello":
                        v3 = v3
                    default:
                        var v64 uint64 = 18446744073709551615
                        _ = v64
                    }
                    var v65 int = fi61
                    _ = fi63
                    _ = v65
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v62
            }
        case -36:
            var v66 int64 = 1
            _ = v66
        default:
            if (genFn1() == v47) {
                v4 -= v4
            } else {
                var v67 float32 = v4
                _ = v67
            }
        }
        _ = fi61
    }
    var v68 float64 = 1.5
    v4 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v68
}
