package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int16; F2 bool }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 string; F2 int64 }
func genFn0(p0 int, p1 uint, p2 uint32) uint8 {
    v4 := GenS1{F0: 0, F1: 57}
    var v5 int32 = -993
    v6 := GenS2{F0: 32767, F1: false, F2: -587}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    if true {
        var v7 int8 = -48
        var v8 uint64 = 42
        var v9 bool = true
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{1, 4294967295}
            switch "\n\t" {
            case "hello":
                var v12 uint = 65
                _ = v12
            case "foo bar":
                for fi13 := 0; fi13 < 6; fi13++ {
                    if false {
                        v14 := v1["foo bar"]
                        v14 = v14
                        _ = v14
                    } else {
                        var v15 int32 = -2147483648
                        fmt.Println(fi10, fi10, v15)
                        v11 = append(v11, fi10)
                        _ = v15
                    }
                    _ = fi13
                }
            default:
                var v16 uint = 42
                _ = v16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 int32 = 1
                v18 = v18
                _ = fi17
                _ = v18
            }
            _ = fi10
            _ = v11
        }
        var v19 bool = ((("a" + (("\n\t" + "foo bar") + "a")) + "hello") >= (("hello" + "\n\t") + "\n\t"))
        _ = v19
    }
    if false {
        var v20 int32 = 1
        _ = v20
    }
    var v21 uint8 = 42
    for fi22 := 0; fi22 < 8; fi22++ {
        v21 += (v21 << ((18446744073709551615) & 63))
        var v23 string = ("foo bar" + "foo bar")
        _ = fi22
        _ = v23
    }
    var v24 int16 = 0
    v25 := []int{4294967295, 32767, 628}
    var v26 int8 = 0
    fmt.Println(v26)
    var v27 int32 = 2147483647
    var v28 bool = true
    if (255 <= 65535) {
        if (!(4294967295 >= 4294967295)) {
            if (v28 || (v28 && ((255 < 162) && v28))) {
                var v29 uint8 = (v21 + v21)
                v30 := GenS3{F0: 51}
                _ = v29
                _ = v30
            } else {
                v31 := GenS4{F0: ("a" + "a"), F1: v24, F2: v28}
                v2["\n\t"] = "hello"
                _ = v31
            }
            var v32 uint64 = 18446744073709551615
            _ = v32
        }
    } else {
        switch 32768 {
        case -661:
            var v33 uint8 = (v21 - genFn0(-1, 839, 0))
            _ = v33
        case 32767:
            var v34 int64 = -9223372036854775808
            _ = v34
        default:
            for fi35 := 0; fi35 < 17; fi35++ {
                if true {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        var v37 float32 = (3.14159 - 1.5)
                        var v38 uint16 = 2
                        var v39 int64 = -9223372036854775808
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                _ = fi35
            }
        }
        v40 := v3[uint32(v21)]
        _ = v40
    }
    v21 = (v21 / ((2) | 1))
    if (560 == 293) {
        var v41 int32 = v27
        switch 32767 {
        case 42:
            var v42 uint32 = 1
            _ = v42
        case -9223372036854775808:
            fmt.Println(v41, v26, v28)
        }
        _ = v41
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        if false {
            v44 := GenS1{F0: v27, F1: 255}
            switch 127 {
            case -716:
                v45 := GenS1{F0: (v27 | v27), F1: genFn0(32767, 588, 2147483648)}
                _ = v45
            case -27:
                for fi46 := 0; fi46 < 1; fi46++ {
                    _ = len(v25)
                    if v28 {
                        if (v28 && (int(v26) >= fi46)) {
                            if v28 {
                                var v47 int = fi43
                                v24 = (v24 + int16(v26))
                                switch "\n\t" {
                                case "hello":
                                    v28 = (18446744073709551615 > 1)
                                case "a":
                                    v48 := v3[65535]
                                    v27 = 845
                                    if len(v25) > 0 {
                                        v25[0] = (v47 ^ (v47 % ((v47) | 1)))
                                    }
                                    _ = v48
                                case "foo bar":
                                    v47 += v47
                                default:
                                    v26 = 100
                                }
                                _ = v47
                            } else {
                                v28 = v28
                                var v49 uint64 = 18446744073709551615
                                _ = v49
                            }
                            v26 -= v26
                            v28 = ((v28 || v28) || false)
                        } else {
                            if (false && v28) {
                                var v50 int64 = -9223372036854775808
                                _ = v50
                            } else {
                                switch "" {
                                case "\n\t":
                                    v28 = v28
                                case "hello":
                                    v28 = false
                                }
                                var v51 bool = (!(v28 && (int32(v24) > v27)))
                                var v52 uint32 = 662
                                _ = v51
                                _ = v52
                            }
                            v24 += (v24 - v24)
                            v28 = (4294967295 >= 845)
                            if ((2 > 18446744073709551615) && (v21 < 128)) {
                                if (!v28) {
                                    switch "a" {
                                    case "hello":
                                        v53 := GenS5{F0: (v21 + v21), F1: "hello"}
                                        v26 += (v26 * v26)
                                        var v54 uint64 = 42
                                        var v55 uint64 = (v54 / ((42) | 1))
                                        v54 = (v54 / ((v54) | 1))
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    case "a":
                                        if (v28 || v28) {
                                            fmt.Println(fi46, v24)
                                            v56 := GenS6{F0: 551}
                                            v24 += -145
                                            if v28 {
                                                var v57 int8 = -1
                                                _ = v57
                                            }
                                            _ = v56
                                        }
                                    case "":
                                        if (!v28) {
                                            fmt.Println(v27)
                                            v58 := GenS7{F0: 860, F1: 18446744073709551615}
                                            fmt.Println(v28)
                                            _ = v58
                                        }
                                    default:
                                        switch "a" {
                                        case "foo bar":
                                            var v59 float64 = 1e10
                                            _ = v59
                                        case "hello":
                                            v60 := GenS8{F0: 162}
                                            var v61 uint = 255
                                            var v62 int = fi43
                                            var v63 uint8 = (v21 ^ 2)
                                            v27 += v27
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        default:
                                            fmt.Println(fi46)
                                        }
                                    }
                                    v28 = (v21 != (^(v21)))
                                    fmt.Println(v21, fi46, v21)
                                }
                                var v64 float64 = 0.1
                                v26 = 127
                                _ = v64
                            } else {
                                v28 = (-2.25 >= -2.25)
                            }
                            var v65 uint = 37
                            v66 := GenS8{F0: 18446744073709551615}
                            switch (("\n\t" + ("hello" + "foo bar")) + "foo bar") {
                            case "hello":
                                _ = len(v2)
                            case "a":
                                var v67 uint16 = 360
                                _ = v67
                            }
                            _ = v65
                            _ = v66
                        }
                        var v68 string = ""
                        _ = v68
                    } else {
                        v69 := GenS7{F0: 42, F1: 18446744073709551615}
                        v26 -= (v26 << ((858) & 63))
                        v24 += (-(v24))
                        _ = v69
                    }
                    _ = fi46
                }
            case 32768:
                v70 := GenS9{F0: "a", F1: ((("" + "\n\t") + ("" + ("a" + ("hello" + ("a" + "hello"))))) + "hello"), F2: -9223372036854775808}
                _ = v70
            }
            _ = len(v2)
            _ = v44
        } else {
            delete(v2, "hello")
        }
        _ = fi43
    }
    v27 += (^(v27))
    v71 := GenS5{F0: 0, F1: ("foo bar" + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v71
}
