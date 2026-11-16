package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint16 }

func main() {
    var v1 int = 1
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    fmt.Println(v1, v1)
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 18446744073709551615, F1: 33}
        if (!true) {
            var v6 uint64 = 812
            _ = v6
        } else {
            var v7 int64 = 276
            v8 := v3[4294967295]
            v8 = 1
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    var v9 uint8 = 42
    switch "" {
    case "hello":
        v10 := []int{(v1 % ((v1) | 1)), 1, v1, (^(v1))}
        _ = v10
    case "a":
        if (-2147483648 == 42) {
            v11 := GenS2{F0: -2147483648}
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS3{F0: -1, F1: true, F2: 0}
                fmt.Println(v2)
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    case "foo bar":
        fmt.Println(v1, v1, v2)
    }
    if false {
        v14 := []int{-264, (v1 + 451)}
        switch v1 {
        case -1:
            if len(v14) > 0 {
                v14[0] = 9223372036854775807
            }
        case 702:
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        _ = v14
    } else {
        var v15 bool = true
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float64 = (float64(0.0) * v2)
        _ = fi16
        _ = v17
    }
    if (0 <= -5) {
        v2 = v2
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            if false {
                if (255 == 930) {
                    switch "\n\t" {
                    case "\n\t":
                        v9 += v9
                    }
                    var v19 int64 = -554
                    v3[951] = 355
                    _ = v19
                }
            } else {
                if false {
                    v1 += fi18
                    fmt.Println(v1)
                    var v20 uint8 = (v9 >> ((uint(v9)) & 63))
                    _ = v20
                } else {
                    v21 := GenS1{F0: 36, F1: 2}
                    _ = v21
                }
                var v22 string = "foo bar"
                v23 := GenS3{F0: -518, F1: true, F2: 0}
                _ = v22
                _ = v23
            }
            _ = fi18
        }
        delete(v3, 2)
    }
    fmt.Println(v1, v1, v9)
    fmt.Println(v9)
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS4{F0: 749, F1: -1}
        for fi26 := 0; fi26 < 19; fi26++ {
            if false {
                switch "" {
                case "":
                    v27 := []int{fi26, fi26}
                    _ = v27
                }
                v9 += (^(v9))
            }
            var v28 uint = 0
            var v29 int = (fi24 / ((v1) | 1))
            _ = fi26
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v2, v9, v9)
    v9 = (v9 * (v9 - 128))
    var v30 uint32 = 2147483648
    var v31 string = ""
    var v32 uint = 42
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 4; fi34++ {
            switch (fi33 & fi33) {
            case 42:
                v9 = (^(v9))
            case 32768:
                v35 := GenS3{F0: (fi34 & v1), F1: (32767 < -265), F2: v30}
                switch (fi33 + (fi34 - fi34)) {
                case -2147483648:
                    v30 += 345
                default:
                    switch (("" + ("" + "")) + (("hello" + v31) + ((v31 + "") + v31))) {
                    case "foo bar":
                        fmt.Println(v2, v2, v9)
                    case "hello":
                        v1 += 2147483647
                    default:
                        switch (fi33 - fi33) {
                        case 386:
                            v31 = ("hello" + v31)
                            if true {
                                v9 += (v9 << ((v32) & 63))
                                if (v2 < 1e10) {
                                    var v36 int8 = -101
                                    v36 = v36
                                    fmt.Println(v32, v30)
                                    _ = v36
                                } else {
                                    fmt.Println(v32, v9)
                                    var v37 uint = 255
                                    v38 := GenS1{F0: (v37 - (v32 - v32)), F1: 591}
                                    _ = v37
                                    _ = v38
                                }
                            }
                        case -965:
                            _ = len(v3)
                        case -959:
                            fmt.Println(v30, v9, v1)
                        }
                    }
                }
                _ = v35
            case 2147483647:
                switch (fi34 + (fi34 ^ 128)) {
                case -2147483648:
                    var v39 float64 = ((1e10 + v2) + v2)
                    _ = v39
                case 32768:
                    switch (fi33 + 2147483647) {
                    case 128:
                        v40 := GenS5{F0: 2, F1: 0, F2: uint64(v9)}
                        _ = v40
                    case 32768:
                        var v41 uint32 = 2147483648
                        _ = v41
                    case -107:
                        delete(v3, v30)
                    default:
                        fmt.Println(v9, v1)
                    }
                default:
                    var v42 int32 = 1
                    _ = v42
                }
                fmt.Println(v9)
            }
            var v43 int64 = 0
            v44 := v3[uint32(v32)]
            _ = fi34
            _ = v43
            _ = v44
        }
        switch v31 {
        case "hello":
            for fi45 := 0; fi45 < 17; fi45++ {
                v46 := v3[v30]
                switch "\n\t" {
                case "a":
                    if true {
                        var v47 float64 = 0.0
                        _ = v47
                    }
                case "hello":
                    var v48 int16 = -351
                    _ = v48
                case "":
                    v49 := []int{}
                    _ = v49
                default:
                    fmt.Println(fi33)
                }
                _ = fi45
                _ = v46
            }
        case "\n\t":
            v50 := GenS6{F0: 42, F1: 42, F2: 23}
            _ = v50
        case "":
            if true {
                if ((v32 & v32) >= (^(v32))) {
                    var v51 int16 = -302
                    _ = v51
                } else {
                    v52 := GenS5{F0: (v30 % ((1) | 1)), F1: 2, F2: 42}
                    v32 = 4294967295
                    _ = v52
                }
            }
        default:
            if (0 != 1) {
                var v53 int = 4294967295
                _ = v53
            } else {
                v54 := []int{(v1 - (fi33 + fi33))}
                var v55 float64 = v2
                _ = v54
                _ = v55
            }
        }
        _ = fi33
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        for fi57 := 0; fi57 < 13; fi57++ {
            if ((1 < 803) || true) {
                var v58 int8 = -100
                v59 := GenS4{F0: 37, F1: (v58 << (((^(v32))) & 63))}
                v32 -= v32
                _ = v58
                _ = v59
            } else {
                v60 := []int{}
                var v61 float32 = -2.25
                _ = v60
                _ = v61
            }
            switch (v31 + (v31 + "\n\t")) {
            case "foo bar":
                var v62 int8 = 0
                _ = v62
            default:
                var v63 uint16 = 2
                _ = v63
            }
            v64 := []int{}
            var v65 uint = (v32 & (v32 | v32))
            _ = fi57
            _ = v64
            _ = v65
        }
        _ = fi56
    }
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v9
}
