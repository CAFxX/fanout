package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 uint32; F2 int }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 string; F1 uint16; F2 int16 }
func genFn0() int16 {
    v4 := GenS1{F0: 65535}
    v5 := []int{}
    switch ("hello" + "foo bar") {
    case "":
        var v6 uint32 = 727
        _ = v6
    case "foo bar":
        var v7 uint16 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = 100
        _ = fi8
        _ = v9
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint32 {
    switch ("a" + "hello") {
    case "\n\t":
        var v10 bool = true
        _ = v10
    case "a":
        var v11 int32 = 1
        _ = v11
    }
    var v12 float64 = -2.25
    v13 := GenS1{F0: 0}
    _ = v12
    _ = v13
    return 299
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 int8 = 78
    if (0 > 65535) {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint64 = 42
            if v2 {
                v16 := GenS2{F0: 0, F1: 101, F2: 128}
                _ = v16
            } else {
                fmt.Println(v15, fi14, v3)
                v17 := v1["\n\t"]
                _ = v17
            }
            v18 := GenS2{F0: 861, F1: 161, F2: 127}
            _ = fi14
            _ = v15
            _ = v18
        }
        v2 = v2
    }
    switch "hello" {
    case "a":
        fmt.Println(v3, v3)
    case "":
        if (255 != 68) {
            v19 := []int{9223372036854775807, 1}
            v20 := GenS2{F0: 0, F1: 235, F2: 151}
            _ = v19
            _ = v20
        } else {
            if (127 <= 173) {
                var v21 float64 = 1e10
                var v22 uint = 608
                _ = v21
                _ = v22
            }
            v3 = v3
            if (0 < 0) {
                v23 := GenS3{F0: "hello", F1: 2}
                v2 = (!v2)
                _ = v23
            } else {
                v24 := v1[("foo bar" + "a")]
                if false {
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        v25 := GenS1{F0: v24}
                        _ = v25
                    case "hello":
                        v26 := []int{-597, 803}
                        _ = v26
                    case "":
                        var v27 uint = 18446744073709551615
                        _ = v27
                    }
                } else {
                    switch 32768 {
                    case -727:
                        if false {
                            switch ("a" + ("hello" + "a")) {
                            case "a":
                                for fi28 := 0; fi28 < 4; fi28++ {
                                    fmt.Println(v24)
                                    _ = fi28
                                }
                            case "":
                                v29 := []int{}
                                _ = v29
                            case "foo bar":
                                var v30 int = -354
                                _ = v30
                            }
                        } else {
                            v31 := []int{-1}
                            _ = v31
                        }
                    case -9223372036854775808:
                        v2 = (41 >= 1)
                    case 9223372036854775807:
                        v24 -= 2147483648
                    default:
                        v32 := GenS4{F0: ("a" + ("\n\t" + "\n\t")), F1: 42, F2: 1}
                        _ = v32
                    }
                }
                v33 := GenS2{F0: 0, F1: 1, F2: 1}
                _ = v24
                _ = v33
            }
        }
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        fmt.Println(v3)
    }
    switch -9223372036854775808 {
    case -883:
        for fi34 := 0; fi34 < 15; fi34++ {
            if v2 {
                v35 := []int{-186}
                for fi36 := 0; fi36 < 17; fi36++ {
                    v2 = v2
                    v2 = true
                    switch 2147483647 {
                    case -893:
                        var v37 int = (fi34 << ((189) & 63))
                        _ = v37
                    case -1:
                        var v38 int16 = 0
                        _ = v38
                    }
                    _ = fi36
                }
                v35 = append(v35, fi34)
                _ = v35
            } else {
                if v2 {
                    v1[(("" + "\n\t") + "")] = 4294967295
                } else {
                    fmt.Println(v2, fi34, v3)
                    v39 := GenS4{F0: ("hello" + "hello"), F1: 378, F2: 216}
                    _ = v39
                }
            }
            _ = fi34
        }
    }
    switch ("a" + "\n\t") {
    case "\n\t":
        if false {
            var v40 uint8 = 0
            fmt.Println(v40)
            _ = v40
        }
    case "":
        if (700 > 0) {
            v2 = false
            var v41 int32 = -742
            v42 := GenS2{F0: 984, F1: 42, F2: 1}
            _ = v41
            _ = v42
        } else {
            fmt.Println(v3)
        }
    }
    v43 := []int{0, -535, -900}
    fmt.Println(v2)
    fmt.Println(v3, v3)
    fmt.Println(v2, v3)
    var v44 uint16 = 2
    if ((v3 * (v3 / ((-1) | 1))) != -1) {
        v43 = append(v43, -9223372036854775808)
    } else {
        switch 32767 {
        case 32768:
            if (1 < -731) {
                _ = len(v43)
            } else {
                v2 = (true || true)
            }
        case 834:
            for fi45 := 0; fi45 < 1; fi45++ {
                if true {
                    fmt.Println(fi45, v3, v2)
                    v46 := GenS4{F0: "a", F1: 18446744073709551615, F2: 128}
                    _ = v46
                } else {
                    var v47 float64 = 0.1
                    var v48 uint = 65535
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        case 127:
            var v49 uint16 = v44
            _ = v49
        default:
            var v50 float32 = float32(((float64(float32(float64(0.0))) - (0.1 * 1.5)) + -2.25))
            _ = v50
        }
        var v51 float32 = -2.25
        var v52 int = 4294967295
        _ = v51
        _ = v52
    }
    var v53 float32 = 1e10
    if (((-1 != 1) && true) && v2) {
        var v54 int64 = -997
        switch "foo bar" {
        case "hello":
            for fi55 := 0; fi55 < 2; fi55++ {
                v56 := v1[("a" + "a")]
                if v2 {
                    v57 := GenS5{F0: "a", F1: 4294967295, F2: fi55}
                    fmt.Println(v2)
                    _ = v57
                }
                _ = fi55
                _ = v56
            }
        default:
            v54 = (v54 % ((v54) | 1))
        }
        _ = v54
    } else {
        for fi58 := 0; fi58 < 4; fi58++ {
            if v2 {
                for fi59 := 0; fi59 < 18; fi59++ {
                    fmt.Println(fi58, v3, v53)
                    _ = fi59
                }
                var v60 uint = 65535
                _ = v60
            } else {
                v3 = v3
            }
            v61 := GenS2{F0: 542, F1: 18446744073709551615, F2: 183}
            v62 := GenS6{F0: (v44 ^ v44), F1: fi58}
            _ = fi58
            _ = v61
            _ = v62
        }
        v63 := GenS7{F0: ((("" + "") + "") + ("hello" + "a")), F1: v44, F2: -32768}
        _ = v63
    }
    fmt.Println(v2, v44, v44)
    v2 = false
    var v64 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v53:", v53)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v53
    _ = v64
}
