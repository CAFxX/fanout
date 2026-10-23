package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 uint8 = 61
    var v2 uint = 421
    var v3 uint64 = 0
    v4 := make(map[int]bool)
    var v5 uint8 = (^(v1))
    var v6 uint16 = 426
    v7 := GenS1{F0: v3, F1: v6, F2: 42}
    v8 := GenS2{F0: v1, F1: 2147483648, F2: false}
    v9 := GenS3{F0: "", F1: 473, F2: v3}
    var v10 float32 = 0.1
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v5, v3, v2)
    case "\n\t":
        var v11 int8 = -128
        _ = v11
    case "":
        fmt.Println(v2)
    default:
        switch ("\n\t" + ("foo bar" + (("" + "\n\t") + ""))) {
        case "\n\t":
            for fi12 := 0; fi12 < 6; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int16 = 1000
                    switch ("hello" + "\n\t") {
                    case "hello":
                        var v15 uint = 65535
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                v10 += v10
                _ = fi12
            }
        default:
            v2 += v2
        }
    }
    var v16 int8 = int8(v6)
    v1 -= v5
    v2 = (v2 - uint(v6))
    switch 581 {
    case -613:
        if (-679 <= 42) {
            var v17 bool = false
            if ((v5 == v1) || (!false)) {
                var v18 int16 = -1
                _ = v18
            } else {
                for fi19 := 0; fi19 < 17; fi19++ {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        if false {
                            if ((v5 << ((2) & 63)) >= (v1 << (((v2 / ((2) | 1))) & 63))) {
                                var v21 int64 = 892
                                v22 := GenS4{F0: v6}
                                var v23 uint = v2
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            if false {
                                var v24 bool = v17
                                _ = v24
                            } else {
                                v4[-9223372036854775808] = (v6 < 0)
                                var v25 uint8 = (v5 - 1)
                                v26 := GenS5{F0: 437, F1: 42}
                                _ = v25
                                _ = v26
                            }
                            v17 = v17
                        } else {
                            var v27 int32 = 2147483647
                            _ = len(v4)
                            _ = v27
                        }
                        v28 := GenS3{F0: ("a" + ("\n\t" + "foo bar")), F1: 42, F2: (v3 * (^(v3)))}
                        v5 += 33
                        _ = fi20
                        _ = v28
                    }
                    _ = fi19
                }
                var v29 uint64 = (v3 + v3)
                _ = v29
            }
            _ = v17
        } else {
            switch "a" {
            case "\n\t":
                v30 := GenS6{F0: int32(v16), F1: v16}
                _ = v30
            case "hello":
                v6 += (^(v6))
            }
        }
    case 1:
        for fi31 := 0; fi31 < 14; fi31++ {
            if true {
                v32 := GenS7{F0: 42}
                if (!true) {
                    switch fi31 {
                    case 128:
                        fmt.Println(v10, v10, v5)
                    case -9223372036854775808:
                        var v33 float64 = 1.5
                        _ = v33
                    default:
                        var v34 string = "\n\t"
                        _ = v34
                    }
                    v6 = (v6 + v6)
                } else {
                    var v35 uint32 = 825
                    v36 := GenS1{F0: v3, F1: (v6 % (((^(v6))) | 1)), F2: -32768}
                    _ = v35
                    _ = v36
                }
                v6 = v6
                _ = v32
            } else {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v38 := []int{(fi37 << (((v2 & v2)) & 63)), fi37, fi37}
                    _ = fi37
                    _ = v38
                }
                v39 := GenS4{F0: (^(v6))}
                v16 = v16
                _ = v39
            }
            switch "hello" {
            case "foo bar":
                for fi40 := 0; fi40 < 7; fi40++ {
                    v41 := GenS1{F0: (v3 | v3), F1: (v6 << ((v2) & 63)), F2: 1}
                    var v42 uint16 = (^(v6))
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            var v43 uint16 = 32767
            _ = fi31
            _ = v43
        }
    case 52:
        var v44 int = 9223372036854775807
        _ = v44
    default:
        switch 549 {
        case 4294967295:
            for fi45 := 0; fi45 < 1; fi45++ {
                var v46 float64 = (0.0 / (-2.25 - (-2.25 / 1e10)))
                v10 = v10
                var v47 int64 = -159
                _ = fi45
                _ = v46
                _ = v47
            }
        case -2147483648:
            var v48 uint = v2
            _ = v48
        default:
            v49 := v4[-2147483648]
            _ = v49
        }
    }
    v5 = (^(v1))
    switch "" {
    case "foo bar":
        v3 = v3
    case "a":
        v50 := []int{9223372036854775807}
        _ = v50
    default:
        if true {
            var v51 string = ""
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
