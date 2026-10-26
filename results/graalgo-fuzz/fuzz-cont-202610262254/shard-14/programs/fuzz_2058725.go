package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int64, p1 int32, p2 int64) string {
    var v5 float32 = 1e10
    if (0 < 42) {
        var v6 int8 = -1
        v7 := GenS1{F0: p1}
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int32 = -401
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS2{F0: 42, F1: 42, F2: 4294967295}
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 42
            _ = fi8
            _ = v12
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        v13 := GenS1{F0: 42}
        p0 -= -982
        var v14 int64 = p0
        _ = v13
        _ = v14
    }
    p2 += p0
    switch -910 {
    case 472:
        v15 := GenS1{F0: (p1 | (p1 - p1))}
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        p0 = p0
        for fi17 := 0; fi17 < 11; fi17++ {
            if (!true) {
                if true {
                    v18 := GenS1{F0: (-(p1))}
                    _ = v18
                }
                v19 := GenS1{F0: (p1 << ((4294967295) & 63))}
                _ = v19
            } else {
                switch "\n\t" {
                case "a":
                    var v20 bool = (0.1 == float64(v5))
                    _ = v20
                case "hello":
                    v21 := GenS2{F0: 214, F1: 128, F2: 395}
                    _ = v21
                case "foo bar":
                    v22 := []int{0}
                    var v23 uint8 = 225
                    switch "foo bar" {
                    case "a":
                        switch ((("" + "") + ((("hello" + "hello") + "") + "hello")) + ("" + "hello")) {
                        case "\n\t":
                            var v24 string = "a"
                            _ = v24
                        case "a":
                            var v25 bool = (true || (2 <= 0))
                            _ = v25
                        case "hello":
                            v26 := GenS1{F0: -2147483648}
                            v27 := GenS1{F0: -97}
                            switch "\n\t" {
                            case "":
                                var v28 uint16 = 65535
                                _ = v28
                            default:
                                v23 += 47
                            }
                            _ = v26
                            _ = v27
                        default:
                            var v29 uint = 1
                            _ = v29
                        }
                    case "hello":
                        var v30 int8 = 127
                        _ = v30
                    default:
                        v23 -= v23
                    }
                    var v31 int8 = 42
                    _ = v22
                    _ = v23
                    _ = v31
                }
                var v32 uint16 = 414
                _ = v32
            }
            switch (^(fi17)) {
            case -2147483648:
                var v33 int32 = (p1 + (p1 ^ (p1 - (p1 + p1))))
                _ = v33
            default:
                if false {
                    var v34 int64 = (-(p0))
                    switch 42 {
                    case 32768:
                        fmt.Println(fi16, p0, fi17)
                    }
                    var v35 uint8 = 221
                    _ = v34
                    _ = v35
                }
            }
            switch (fi17 + fi17) {
            case 1:
                fmt.Println(v5, p2, fi16)
            case -2147483648:
                v36 := GenS2{F0: 127, F1: 1, F2: 1}
                _ = v36
            case 0:
                var v37 string = ""
                _ = v37
            default:
                p1 = 584
            }
            _ = fi17
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}
func genFn1() int16 {
    var v38 uint = 788
    var v39 float32 = (0.1 / -2.25)
    _ = v38
    _ = v39
    return -1
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = 1
    var v3 uint32 = 431
    var v4 int32 = -481
    v1 -= v1
    v1 += v1
    if false {
        var v40 int64 = 1
        v41 := GenS3{F0: 127, F1: uint32(v1), F2: 9223372036854775807}
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        switch fi42 {
        case -697:
            if true {
                fmt.Println(v4)
                var v43 uint64 = 1
                _ = v43
            } else {
                v44 := GenS4{F0: (v4 * 57), F1: 1}
                for fi45 := 0; fi45 < 6; fi45++ {
                    v3 = (v3 << ((255) & 63))
                    _ = fi45
                }
                _ = v44
            }
        case 107:
            v3 += 668
        case 509:
            var v46 int16 = int16(v1)
            _ = v46
        default:
            v47 := GenS4{F0: v4, F1: 255}
            _ = v47
        }
        if (((v4 - (v4 + int32(v2))) >= (-(v4))) && (true || false)) {
            v48 := GenS5{F0: 0, F1: 0}
            _ = v48
        } else {
            v4 += (^(v4))
            if true {
                v49 := GenS6{F0: 0, F1: 1, F2: -1}
                v1 = 2
                if ((2 == 42) && (false && false)) {
                    v50 := GenS7{F0: 4294967295}
                    _ = v50
                }
                _ = v49
            }
            var v51 uint32 = 1
            _ = v51
        }
        v4 += v4
        _ = fi42
    }
    v52 := []int{127, 1}
    var v53 float32 = float32(3.14159)
    if (90 < v2) {
        _ = len(v52)
        v54 := GenS8{F0: 255}
        v2 += int8(v4)
        _ = v54
    } else {
        v52 = append(v52, 420)
    }
    var v55 uint64 = 215
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v55
}
