package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
func genFn0(p0 int64) int8 {
    if false {
        var v4 int = 4294967295
        var v5 uint8 = 128
        if true {
            v6 := []int{}
            _ = len(v6)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(p0, p0, v4)
                v8 := GenS1{F0: 42, F1: 255}
                _ = fi7
                _ = v8
            }
        }
        _ = v4
        _ = v5
    }
    v9 := GenS1{F0: 100, F1: 317}
    _ = p0
    _ = v9
    return -25
}
func genFn1() int {
    var v10 uint8 = 255
    var v11 float64 = float64(1.5)
    var v12 uint8 = v10
    var v13 uint8 = (v10 % ((v12) | 1))
    switch "foo bar" {
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int32 = 1
            v16 := GenS1{F0: 100, F1: 42}
            v17 := []int{(fi14 + (fi14 | (fi14 * fi14)))}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    case "a":
        if (117 >= 2147483648) {
            v18 := GenS1{F0: -1, F1: 108}
            v19 := GenS1{F0: 100, F1: 65535}
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int16 = 0
                fmt.Println(v13, v10)
                var v22 int32 = -784
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v18
            _ = v19
        } else {
            v23 := []int{1, -478, 1, -9223372036854775808}
            v24 := GenS1{F0: -1, F1: 1}
            var v25 int16 = -74
            _ = v23
            _ = v24
            _ = v25
        }
    case "\n\t":
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := []int{fi26}
            v28 := GenS2{F0: 100}
            v12 = v12
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    switch (("" + "hello") + "hello") {
    case "foo bar":
        var v29 string = "hello"
        _ = v29
    default:
        if (true && ((2147483647 >= 462) || (578 < 636))) {
            v1 = (^(v1))
            v30 := []int{32767, -376}
            _ = v30
        } else {
            var v31 bool = false
            var v32 int = 281
            v2 += v2
            _ = v31
            _ = v32
        }
    }
    if (("hello" + "foo bar") <= "hello") {
        v1 -= (v1 | v1)
        if (128 <= 4294967295) {
            v33 := GenS1{F0: 0, F1: 795}
            _ = v33
        }
    } else {
        var v34 uint64 = 42
        var v35 uint64 = 4294967295
        _ = v34
        _ = v35
    }
    v2 -= 0
    v2 = v2
    for fi36 := 0; fi36 < 4; fi36++ {
        switch (-(fi36)) {
        case 128:
            if true {
                var v37 bool = true
                switch fi36 {
                case 42:
                    fmt.Println(fi36, v37, v37)
                case -123:
                    if (!(0.1 <= 1e10)) {
                        v38 := GenS2{F0: genFn0(v2)}
                        _ = v38
                    }
                case 1:
                    v37 = v37
                }
                _ = v37
            }
        default:
            for fi39 := 0; fi39 < 15; fi39++ {
                v40 := v3[4294967295]
                v40 = (fi39 >> ((v1) & 63))
                _ = fi39
                _ = v40
            }
        }
        _ = len(v3)
        _ = fi36
    }
    switch "\n\t" {
    case "a":
        v41 := GenS1{F0: 42, F1: 32767}
        _ = v41
    case "":
        v42 := []int{261}
        _ = v42
    case "hello":
        if true {
            v43 := v3[182]
            if (!(1 > 10)) {
                var v44 bool = false
                _ = v44
            } else {
                var v45 uint8 = 21
                v3[416] = (v43 - (v43 & (v43 & v43)))
                _ = v45
            }
            _ = v43
        }
    default:
        v3[uint32(v1)] = genFn1()
    }
    var v46 uint16 = 65535
    for fi47 := 0; fi47 < 20; fi47++ {
        switch -1 {
        case 9223372036854775807:
            var v48 uint8 = 42
            _ = v48
        case 32768:
            v49 := GenS1{F0: 127, F1: (v46 / (((v46 << ((65535) & 63))) | 1))}
            _ = v49
        default:
            v50 := GenS1{F0: -1, F1: v46}
            _ = v50
        }
        v51 := GenS3{F0: -2147483648, F1: true, F2: v2}
        _ = fi47
        _ = v51
    }
    switch "a" {
    case "":
        for fi52 := 0; fi52 < 19; fi52++ {
            v53 := GenS4{F0: 2, F1: -2147483648}
            if true {
                var v54 uint16 = 415
                var v55 float32 = 1e10
                _ = v54
                _ = v55
            }
            _ = fi52
            _ = v53
        }
    case "a":
        if false {
            _ = len(v3)
        }
    case "hello":
        switch 32767 {
        case 127:
            for fi56 := 0; fi56 < 8; fi56++ {
                v1 -= 4294967295
                var v57 float32 = 3.14159
                v58 := GenS1{F0: 127, F1: 893}
                _ = fi56
                _ = v57
                _ = v58
            }
        default:
            if true {
                v2 = v2
                fmt.Println(v1, v46, v2)
            }
        }
    }
    v59 := GenS2{F0: 77}
    v60 := GenS5{F0: 42, F1: ((false && ((^(v1)) < v1)) && (((^(v1)) <= v1) || (1 == 0)))}
    v61 := GenS4{F0: 941, F1: 1}
    for fi62 := 0; fi62 < 1; fi62++ {
        var v63 uint16 = (v46 / ((v46) | 1))
        fmt.Println(v46, v2, v1)
        switch ((("\n\t" + "a") + "foo bar") + "") {
        case "\n\t":
            v63 += v46
        }
        _ = fi62
        _ = v63
    }
    if ((3.14159 / 1.5) < -2.25) {
        var v64 int32 = -1
        _ = v64
    } else {
        for fi65 := 0; fi65 < 5; fi65++ {
            v66 := GenS3{F0: 42, F1: (!true), F2: 0}
            for fi67 := 0; fi67 < 9; fi67++ {
                var v68 int = genFn1()
                var v69 uint16 = (v46 - v46)
                switch genFn1() {
                case -2147483648:
                    switch "a" {
                    case "a":
                        var v70 uint = v1
                        _ = v70
                    case "foo bar":
                        v69 = (v46 + 645)
                    default:
                        if false {
                            v71 := GenS3{F0: int32(fi67), F1: (0 != 18446744073709551615), F2: v2}
                            fmt.Println(v2)
                            fmt.Println(fi67, fi65, v2)
                            _ = v71
                        } else {
                            v72 := []int{}
                            v73 := GenS6{F0: 42}
                            switch v68 {
                            case 67:
                                v74 := GenS7{F0: 1000}
                                _ = v74
                            }
                            _ = v72
                            _ = v73
                        }
                    }
                default:
                    fmt.Println(v68, v2, v69)
                }
                _ = fi67
                _ = v68
                _ = v69
            }
            _ = fi65
            _ = v66
        }
    }
    v75 := []int{32767, 9223372036854775807}
    fmt.Println(v46, v2)
    v2 = v2
    if true {
        v76 := v3[2147483647]
        fmt.Println(v76, v76)
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v59
    _ = v60
    _ = v61
    _ = v75
}
