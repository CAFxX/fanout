package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 string; F1 uint16; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 uint32 }
func genFn0(p0 int16, p1 uint8, p2 uint8) uint32 {
    var v5 int16 = -32768
    p0 += p0
    var v6 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 978
}
func genFn1() uint {
    if (false || false) {
        var v7 int64 = 42
        _ = v7
    } else {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        v9 := []int{}
        var v10 bool = (-629 != 1)
        _ = v10
        _ = v9
    }
    var v11 float64 = (-2.25 / 1.0)
    var v12 uint64 = 1
    v13 := GenS1{F0: 11, F1: -32768}
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []uint32{2, 0, 2147483647}
    var v2 float64 = 0.1
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 -= v2
    for fi14 := 0; fi14 < 1; fi14++ {
        v2 -= (v2 / ((float64(3.14159) + v2) / (3.14159 * 1.5)))
        switch -254 {
        case 42:
            if (!(genFn1() <= 418)) {
                v15 := GenS1{F0: 197, F1: 1000}
                _ = v15
            } else {
                var v16 int16 = 32767
                v2 = v2
                v17 := GenS2{F0: 1, F1: v16}
                _ = v16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: -1, F1: 1000}
            delete(v4, fi14)
            if false {
                v20 := GenS2{F0: -1, F1: 42}
                v21 := GenS3{F0: -1, F1: 223}
                v1 = append(v1, 837)
                v2 -= ((((v2 / 3.14159) - ((-2.25 / float64(1e10)) - float64(0.0))) - (-2.25 + (v2 * (v2 / -2.25)))) + v2)
                v2 -= (0.1 - ((0.1 - v2) - v2))
                _ = v20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = fi14
    }
    if (("hello" + "") > "a") {
        switch 127 {
        case -2147483648:
            v22 := GenS2{F0: 57, F1: 1000}
            _ = v22
        case 32767:
            switch "hello" {
            case "hello":
                v2 = v2
            case "a":
                var v23 float64 = float64(0.1)
                _ = v23
            default:
                v24 := v4[32768]
                _ = v24
            }
        default:
            var v25 string = "foo bar"
            _ = v25
        }
        switch "a" {
        case "foo bar":
            v26 := GenS2{F0: 1, F1: 460}
            _ = v26
        }
        fmt.Println(v2)
    }
    switch "hello" {
    case "foo bar":
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := GenS4{F0: true, F1: -426}
            _ = fi27
            _ = v28
        }
    default:
        v29 := GenS3{F0: 356, F1: 42}
        _ = v29
    }
    fmt.Println(v2)
    v2 += ((v2 * v2) + v2)
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 uint = 770
        _ = fi30
        _ = v31
    }
    if true {
        v32 := GenS5{F0: 2, F1: (42 <= 128)}
        if len(v1) > 0 {
            v1[0] = 2
        }
        switch -2147483648 {
        case -2147483648:
            if false {
                var v33 uint16 = 711
                _ = v33
            } else {
                v2 -= 3.14159
                var v34 uint8 = 0
                _ = v34
            }
        case 369:
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 float32 = 1.5
                for fi37 := 0; fi37 < 6; fi37++ {
                    fmt.Println(fi37, fi35, v36)
                    var v38 int64 = 628
                    if true {
                        _ = len(v4)
                        delete(v4, (fi35 | fi35))
                        _ = len(v1)
                    }
                    _ = fi37
                    _ = v38
                }
                v2 = float64(v36)
                _ = fi35
                _ = v36
            }
        case 32767:
            v39 := GenS4{F0: (127 >= 4294967295), F1: -1}
            _ = v39
        }
        _ = v32
    } else {
        v2 += v2
        v40 := GenS3{F0: -32768, F1: 18446744073709551615}
        var v41 int16 = 42
        _ = v40
        _ = v41
    }
    if (!false) {
        var v42 int = 128
        v4[v42] = 42
        v4[v42] = 729
        _ = v42
    } else {
        v43 := GenS4{F0: (!false), F1: -614}
        v2 -= 0.0
        if false {
            switch "" {
            case "":
                if (-2.25 <= (3.14159 - (1e10 * 1.5))) {
                    v2 -= 0.0
                    v44 := v4[-9223372036854775808]
                    _ = v44
                } else {
                    v45 := GenS6{F0: -684, F1: 128}
                    _ = v45
                }
            }
            v2 += (((v2 * v2) / -2.25) / -2.25)
        } else {
            switch -938 {
            case -613:
                for fi46 := 0; fi46 < 1; fi46++ {
                    v47 := GenS7{F0: "foo bar", F1: 963, F2: 966}
                    _ = fi46
                    _ = v47
                }
            case 231:
                if (0 >= 42) {
                    fmt.Println(v2)
                    var v48 uint32 = 0
                    _ = v48
                }
            }
            var v49 string = ""
            _ = v49
        }
        _ = v43
    }
    switch -2147483648 {
    case -639:
        v2 = -2.25
    }
    v50 := GenS8{F0: 18446744073709551615}
    var v51 string = "hello"
    for fi52 := 0; fi52 < 20; fi52++ {
        fmt.Println(v2, v2, v51)
        for fi53 := 0; fi53 < 14; fi53++ {
            v54 := GenS9{F0: (fi52 + fi52), F1: 1}
            if false {
                v55 := GenS8{F0: 2}
                switch (fi52 % ((32768) | 1)) {
                case 9223372036854775807:
                    _ = len(v4)
                case 386:
                    v51 = "\n\t"
                default:
                    v56 := GenS9{F0: (-(fi53)), F1: 2147483648}
                    _ = v56
                }
                _ = v55
            }
            if false {
                fmt.Println(fi52, v51, fi52)
                v1 = append(v1, 2147483647)
                var v57 bool = (115 >= 27)
                _ = v57
            } else {
                v51 = ((v51 + (("\n\t" + v51) + "")) + v51)
                if true {
                    switch fi52 {
                    case 170:
                        v1 = append(v1, 2147483647)
                    default:
                        v51 = v51
                    }
                } else {
                    var v58 uint32 = 2
                    var v59 string = "hello"
                    _ = v58
                    _ = v59
                }
            }
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
