package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 string; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -128
    v3 := make(map[string]uint32)
    var v4 int = 127
    var v5 uint8 = 1
    var v6 string = (("foo bar" + ("\n\t" + "hello")) + ("" + ""))
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int32 = 0
            if true {
                delete(v3, "\n\t")
            } else {
                if true {
                    v3[""] = 2
                    v10 := GenS1{F0: 2}
                    if (-1 >= 0) {
                        if false {
                            _ = len(v1)
                        }
                        v2 += (-(v2))
                    }
                    _ = v10
                } else {
                    var v11 bool = (!(!(32767 <= 829)))
                    if (-2.25 >= (3.14159 + 3.14159)) {
                        var v12 bool = v11
                        _ = v12
                    }
                    switch v6 {
                    case "foo bar":
                        if ((true && v11) && (1e10 < 1.5)) {
                            _ = len(v1)
                            v13 := GenS1{F0: 402}
                            _ = v13
                        } else {
                            if true {
                                var v14 uint16 = 255
                                if false {
                                    if (385 == 42) {
                                        if v11 {
                                            fmt.Println(v11, v14, v6)
                                            if (v6 <= ("foo bar" + v6)) {
                                                v15 := []int{int(v2)}
                                                _ = v15
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    var v16 uint32 = 2147483648
                                                    if v11 {
                                                        var v17 int = fi8
                                                        _ = v17
                                                    } else {
                                                        v18 := GenS2{F0: v11, F1: v4}
                                                        if (0.0 < 1e10) {
                                                            v19 := GenS3{F0: v2}
                                                            var v20 uint8 = (v5 & 207)
                                                            v5 = (v20 ^ 0)
                                                            _ = v19
                                                            _ = v20
                                                        } else {
                                                            var v21 int8 = (v2 * (v2 / ((v2) | 1)))
                                                            if (3.14159 < float64(1e10)) {
                                                                var v22 float64 = (1.5 * (0.1 + -2.25))
                                                                var v23 int8 = v21
                                                                v24 := []int{-9223372036854775808, 756, v4}
                                                                var v25 int8 = v21
                                                                _ = v22
                                                                _ = v23
                                                                _ = v24
                                                                _ = v25
                                                            }
                                                            _ = v21
                                                        }
                                                        v26 := GenS4{F0: v9}
                                                        _ = v18
                                                        _ = v26
                                                    }
                                                    _ = v16
                                                case "\n\t":
                                                    v27 := GenS4{F0: (v9 << ((42) & 63))}
                                                    _ = v27
                                                default:
                                                    v28 := GenS5{F0: v2, F1: v2}
                                                    _ = v28
                                                }
                                                v29 := GenS3{F0: v2}
                                                v30 := GenS3{F0: v2}
                                                _ = v29
                                                _ = v30
                                            }
                                            v5 -= (^(v5))
                                        } else {
                                            var v31 int32 = 1
                                            v32 := []int{(v4 - 32767), fi8}
                                            var v33 int16 = 42
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        v34 := GenS3{F0: 0}
                                        v35 := GenS3{F0: v2}
                                        _ = v34
                                        _ = v35
                                    }
                                    v36 := GenS3{F0: v2}
                                    delete(v3, v6)
                                    _ = v36
                                } else {
                                    if (65535 < 11) {
                                        var v37 uint8 = (v5 & 2)
                                        switch v6 {
                                        case "foo bar":
                                            v6 = v6
                                        case "hello":
                                            switch "foo bar" {
                                            case "":
                                                v11 = v11
                                            case "a":
                                                v5 -= (v37 * 133)
                                            }
                                        }
                                        _ = v37
                                    }
                                }
                                var v38 string = v6
                                v4 -= fi8
                                _ = v14
                                _ = v38
                            }
                            v39 := GenS5{F0: (^(v2)), F1: -42}
                            _ = v39
                        }
                        var v40 int8 = 42
                        _ = v40
                    case "\n\t":
                        fmt.Println(v2)
                        v41 := []int{933, (fi8 & fi8), 1}
                        _ = len(v1)
                        switch "hello" {
                        case "foo bar":
                            v4 -= (fi8 | 68)
                        case "":
                            v42 := v1[(v6 + "\n\t")]
                            _ = v42
                        case "hello":
                            v2 -= 127
                        default:
                            var v43 uint32 = 4294967295
                            _ = v43
                        }
                        delete(v3, "")
                        _ = v41
                    case "hello":
                        var v44 int = fi8
                        _ = v44
                    default:
                        delete(v3, "\n\t")
                    }
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        if (false && (42 > 1)) {
            _ = len(v1)
            var v45 int16 = 42
            _ = v45
        } else {
            var v46 int32 = -1
            switch (v4 / ((fi7) | 1)) {
            case 947:
                v2 = (v2 - 0)
            case 128:
                if (!false) {
                    v1["hello"] = 2147483647
                    v6 = "a"
                }
            default:
                v3[v6] = 2147483648
            }
            _ = v46
        }
        var v47 int8 = (v2 >> ((2) & 63))
        _ = fi7
        _ = v47
    }
    switch v6 {
    case "\n\t":
        if (!false) {
            v48 := GenS3{F0: (v2 | v2)}
            if (!(42 >= 42)) {
                v49 := []int{v4, 415}
                if len(v49) >= 2 {
                    v49 = v49[1:len(v49)]
                }
                _ = v49
            } else {
                v50 := GenS2{F0: (1 <= (v4 / ((500) | 1))), F1: (v4 ^ v4)}
                v3[""] = 0
                for fi51 := 0; fi51 < 13; fi51++ {
                    v5 -= (v5 | v5)
                    _ = fi51
                }
                _ = v50
            }
            _ = v48
        } else {
            var v52 uint8 = v5
            _ = v52
        }
    case "":
        if (0.0 < (((0.0 * 0.1) * (3.14159 * 1e10)) / 3.14159)) {
            fmt.Println(v4)
        }
    case "foo bar":
        v5 = 0
    }
    v53 := []int{32768, v4, (^(v4))}
    var v54 string = v6
    var v55 uint8 = v5
    v6 = ""
    fmt.Println(v5, v4, v4)
    v56 := GenS1{F0: 579}
    for fi57 := 0; fi57 < 5; fi57++ {
        v58 := GenS3{F0: -128}
        var v59 int16 = -1
        _ = fi57
        _ = v58
        _ = v59
    }
    var v60 uint8 = (v5 ^ (v5 >> ((17) & 63)))
    switch v6 {
    case "":
        v6 = v6
    }
    for fi61 := 0; fi61 < 1; fi61++ {
        if false {
            v6 = v6
            v62 := GenS6{F0: v2, F1: 582}
            for fi63 := 0; fi63 < 13; fi63++ {
                v64 := GenS1{F0: 0}
                fmt.Println(fi63, v54, fi63)
                _ = fi63
                _ = v64
            }
            _ = v62
        }
        _ = fi61
    }
    var v65 int16 = 1000
    var v66 bool = true
    var v67 bool = v66
    v68 := GenS7{F0: v6, F1: v67}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
    _ = v60
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
