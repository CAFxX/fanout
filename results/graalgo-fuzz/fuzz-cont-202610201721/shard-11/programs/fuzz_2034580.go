package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8, p1 uint16) uint16 {
    if true {
        if (((42 != 74) && (!true)) && (65535 <= 2)) {
            v6 := GenS1{F0: false, F1: p0}
            _ = v6
        } else {
            var v7 uint16 = (^(p1))
            fmt.Println(p1)
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(fi8, p1, p1)
                switch "hello" {
                case "a":
                    for fi9 := 0; fi9 < 10; fi9++ {
                        if (498 > 420) {
                            fmt.Println(p0, fi8)
                            v10 := GenS1{F0: ("foo bar" < "foo bar"), F1: 128}
                            var v11 bool = true
                            _ = v10
                            _ = v11
                        }
                        if (true && (1 < 760)) {
                            v7 += 65535
                            fmt.Println(fi8, v7, v7)
                        } else {
                            var v12 uint8 = p0
                            if (!true) {
                                v13 := []int{}
                                _ = v13
                            } else {
                                if (776 != -451) {
                                    var v14 int32 = -2147483648
                                    v15 := []int{fi9, (fi9 / ((42) | 1)), fi8}
                                    var v16 int16 = 0
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    v17 := []int{(fi9 ^ 9223372036854775807)}
                                    var v18 int16 = 1
                                    _ = v17
                                    _ = v18
                                }
                            }
                            p1 -= (p1 - (p1 << ((42) & 63)))
                            _ = v12
                        }
                        var v19 float32 = 0.1
                        _ = fi9
                        _ = v19
                    }
                case "\n\t":
                    switch fi8 {
                    case -1:
                        v20 := GenS1{F0: (1e10 == (-2.25 / 1.0)), F1: (p0 - (^(p0)))}
                        _ = v20
                    case 32768:
                        fmt.Println(fi8)
                    }
                default:
                    switch (("foo bar" + "foo bar") + "\n\t") {
                    case "\n\t":
                        v7 = 1
                    default:
                        var v21 float64 = float64((((0.0 - 0.1) - (0.0 * (0.1 - float32(((1e10 - 0.0) - 1.5))))) - -2.25))
                        _ = v21
                    }
                }
                _ = fi8
            }
            _ = v7
        }
    } else {
        if (1 > 1) {
            var v22 int64 = 1
            fmt.Println(v22, p1)
            _ = v22
        }
        v23 := GenS2{F0: 18446744073709551615, F1: 1}
        _ = v23
    }
    var v24 uint = 862
    _ = p0
    _ = p1
    _ = v24
    return p1
}

func main() {
    var v1 int8 = -116
    v2 := make(map[string]int)
    var v3 uint16 = 231
    var v4 bool = (0.1 != 0.0)
    var v5 float64 = float64(-2.25)
    v25 := GenS2{F0: 834, F1: genFn0(0, (v3 + (v3 | 255)))}
    var v26 bool = (!(!(v3 >= genFn0(42, genFn0(172, (^(v3)))))))
    delete(v2, ("\n\t" + ("" + "foo bar")))
    var v27 int32 = 2147483647
    v28 := []int{-2147483648, 127}
    if ((111 > 186) || ((3.14159 / 1e10) == v5)) {
        fmt.Println(v27)
        v2["hello"] = -1
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS2{F0: 18446744073709551615, F1: (v3 & (v3 & v3))}
            var v31 int16 = 1000
            switch 93 {
            case 2147483647:
                var v32 uint64 = 280
                _ = v32
            case 42:
                v33 := GenS3{F0: (v1 | v1)}
                _ = v33
            case 32767:
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 int32 = (v27 << ((511) & 63))
                    switch "" {
                    case "\n\t":
                        switch fi29 {
                        case 584:
                            if (42 <= 808) {
                                fmt.Println(fi29, fi29, fi34)
                                var v36 string = "\n\t"
                                _ = len(v2)
                                v4 = false
                                v35 -= (v27 & (^(v35)))
                                _ = v36
                            } else {
                                v27 += 214
                            }
                        case 9223372036854775807:
                            var v37 string = "foo bar"
                            _ = v37
                        default:
                            v4 = (((^(v3)) <= 42) && (42 >= 42))
                        }
                    case "foo bar":
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v5, v1, v3)
                        case "":
                            v38 := GenS1{F0: ((^(v31)) > (v31 - v31)), F1: 42}
                            _ = v38
                        case "foo bar":
                            v3 = (v3 % (((v3 / (((^(v3))) | 1))) | 1))
                        default:
                            if true {
                                v27 -= -2147483648
                            } else {
                                v39 := GenS4{F0: 42, F1: 1}
                                if true {
                                    if false {
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        v35 = int32(v3)
                                        v31 = -1
                                    } else {
                                        if v4 {
                                            if v4 {
                                                v26 = v4
                                                v5 = (v5 / v5)
                                            } else {
                                                v5 = v5
                                                fmt.Println(v4, v3, fi29)
                                            }
                                        }
                                        if false {
                                            fmt.Println(v3, fi34, v5)
                                            if (42 >= 1) {
                                                var v40 int = fi29
                                                fmt.Println(v35, v3, v27)
                                                var v41 int = v40
                                                _ = v40
                                                _ = v41
                                            }
                                        }
                                    }
                                    v42 := GenS3{F0: 1}
                                    fmt.Println(v31)
                                    _ = v42
                                } else {
                                    v4 = false
                                }
                                _ = v39
                            }
                        }
                    }
                    var v43 uint = 0
                    _ = fi34
                    _ = v35
                    _ = v43
                }
            default:
                if (!v4) {
                    switch fi29 {
                    case -421:
                        var v44 float32 = 1.5
                        _ = v44
                    case 42:
                        v1 -= (v1 * (v1 - 69))
                    }
                    var v45 uint32 = 148
                    _ = v45
                } else {
                    fmt.Println(v31, v26)
                }
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        switch "" {
        case "a":
            v5 += float64(0.0)
        case "\n\t":
            for fi46 := 0; fi46 < 5; fi46++ {
                switch "a" {
                case "\n\t":
                    var v47 float32 = -2.25
                    _ = v47
                case "a":
                    var v48 int64 = -478
                    _ = v48
                default:
                    v4 = true
                }
                for fi49 := 0; fi49 < 13; fi49++ {
                    var v50 int32 = (v27 << ((255) & 63))
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    _ = fi49
                    _ = v50
                }
                _ = fi46
            }
        case "foo bar":
            v1 = 100
        }
    }
    v4 = (2 <= 310)
    fmt.Println(v1, v27, v26)
    var v51 string = ("hello" + ("\n\t" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
