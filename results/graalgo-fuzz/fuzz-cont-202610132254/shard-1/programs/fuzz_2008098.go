package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint; F1 string; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32, p1 uint, p2 uint32) int8 {
    p2 = (p2 * p0)
    var v4 uint = 599
    switch "hello" {
    case "":
        fmt.Println(p1, p2)
    default:
        v5 := []int{-9223372036854775808}
        _ = v5
    }
    var v6 int16 = -351
    v7 := GenS1{F0: 4294967295, F1: 255, F2: -528}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 uint8, p2 uint) int16 {
    for fi8 := 0; fi8 < 4; fi8++ {
        p1 -= (p1 * p1)
        if false {
            fmt.Println(p0, p2)
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                p1 += p1
                p0 = p0
                p0 = 100
                _ = fi9
            }
        }
        _ = fi8
    }
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 728
}

func main() {
    var v1 int32 = 42
    var v2 uint = 42
    var v3 uint32 = 813
    switch -9223372036854775808 {
    case -2147483648:
        switch -853 {
        case -826:
            var v10 float32 = (-2.25 + (0.1 + 1.5))
            _ = v10
        case 32768:
            fmt.Println(v2)
        default:
            v11 := GenS1{F0: 661, F1: 255, F2: 32767}
            _ = v11
        }
    default:
        fmt.Println(v1)
    }
    switch 0 {
    case -1:
        v1 = -850
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v1, v2, v2)
        if false {
            v13 := []int{(fi12 / (((fi12 ^ (fi12 & (-(fi12))))) | 1)), (fi12 - (fi12 << ((v2) & 63))), 0}
            _ = v13
        } else {
            var v14 int16 = 1000
            switch "foo bar" {
            case "\n\t":
                var v15 uint8 = 255
                _ = v15
            case "hello":
                switch "hello" {
                case "foo bar":
                    v16 := GenS2{F0: uint32(v2)}
                    _ = v16
                case "\n\t":
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if (false && (fi12 >= (fi12 ^ fi17))) {
                            fmt.Println(v14)
                            if (1 < 65535) {
                                v18 := GenS3{F0: false, F1: 42, F2: 0}
                                v19 := GenS4{F0: ((-(v1)) > v1), F1: v3}
                                _ = v18
                                _ = v19
                            } else {
                                v2 = v2
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v2, v3)
                                case "\n\t":
                                    if false {
                                        if (v14 >= (v14 * (v14 | v14))) {
                                            var v20 float64 = 0.0
                                            v3 = v3
                                            v21 := GenS5{F0: 4294967295, F1: "hello"}
                                            v22 := []int{(fi12 + (-(fi12)))}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        } else {
                                            v23 := []int{(fi17 % ((fi12) | 1)), fi12, fi12, fi17}
                                            var v24 uint8 = 42
                                            if ((fi17 | fi12) == (fi12 ^ (fi17 * 292))) {
                                                switch fi17 {
                                                case 128:
                                                    v25 := GenS6{F0: v2, F1: "a", F2: fi12}
                                                    _ = v25
                                                case 0:
                                                    fmt.Println(v1, v1)
                                                case 133:
                                                    fmt.Println(v24, v1, v14)
                                                default:
                                                    if len(v23) > 0 {
                                                        v23[0] = 127
                                                    }
                                                }
                                                if (((("hello" + "hello") + ("foo bar" + ("a" + ("\n\t" + ((("" + "\n\t") + "foo bar") + (("hello" + "foo bar") + "")))))) + ("" + (("hello" + "hello") + "foo bar"))) < "") {
                                                    var v26 int64 = 1
                                                    v27 := GenS7{F0: (v26 | (v26 % ((1) | 1)))}
                                                    switch (fi12 & int(v2)) {
                                                    case -873:
                                                        v1 -= 42
                                                    case 9223372036854775807:
                                                        v14 -= (v14 >> ((v2) & 63))
                                                    }
                                                    _ = v26
                                                    _ = v27
                                                }
                                                var v28 bool = (v14 == 1000)
                                                v29 := GenS1{F0: 4294967295, F1: (v24 << ((869) & 63)), F2: -1}
                                                var v30 uint64 = 2
                                                v23 = append(v23, 0)
                                                if len(v23) >= 2 {
                                                    v23 = v23[1:len(v23)]
                                                }
                                                switch fi17 {
                                                case 9223372036854775807:
                                                    v3 += 367
                                                default:
                                                    if v28 {
                                                        switch fi17 {
                                                        case 4294967295:
                                                            v30 = 1
                                                        case 1:
                                                            v30 += 2
                                                        }
                                                    }
                                                }
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v23
                                            _ = v24
                                        }
                                        fmt.Println(v3, v14)
                                        v2 = (v2 % ((v2) | 1))
                                    } else {
                                        v31 := []int{}
                                        v3 = v3
                                        _ = v31
                                    }
                                case "":
                                    fmt.Println(v1, fi12, v2)
                                }
                            }
                            v1 = v1
                        } else {
                            fmt.Println(v2, fi12, fi12)
                            var v32 uint8 = 255
                            _ = v32
                        }
                        switch fi12 {
                        case 772:
                            var v33 uint8 = 2
                            _ = v33
                        case 0:
                            v34 := []int{(fi17 | (fi17 - 32768)), 4294967295}
                            switch "foo bar" {
                            case "a":
                                if len(v34) > 0 {
                                    v34[0] = (^(fi12))
                                }
                            case "foo bar":
                                var v35 string = "foo bar"
                                _ = v35
                            }
                            _ = v34
                        }
                        v2 += uint(v3)
                        v14 = v14
                        v36 := []int{}
                        _ = fi17
                        _ = v36
                    }
                default:
                    for fi37 := 0; fi37 < 11; fi37++ {
                        v38 := []int{(fi37 ^ fi12), fi12, (^(fi37))}
                        v1 += (v1 & v1)
                        _ = fi37
                        _ = v38
                    }
                }
            case "a":
                fmt.Println(v1, v1)
            }
            v39 := GenS8{F0: 94, F1: 42}
            _ = v14
            _ = v39
        }
        _ = fi12
    }
    switch "\n\t" {
    case "\n\t":
        for fi40 := 0; fi40 < 8; fi40++ {
            var v41 uint64 = 1
            _ = fi40
            _ = v41
        }
    case "a":
        var v42 uint32 = v3
        _ = v42
    default:
        v3 = 1
    }
    v43 := []int{0, 32768}
    if len(v43) > 0 {
        v43[0] = 950
    }
    if ((false || true) || (true || (!false))) {
        var v44 int64 = 274
        _ = v44
    } else {
        var v45 int = 355
        _ = v45
    }
    var v46 uint32 = (v3 & (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v46
}
