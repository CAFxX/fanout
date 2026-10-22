package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 bool; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint16; F1 int16; F2 int32 }
func genFn0(p0 uint8, p1 int64) uint16 {
    var v5 int64 = 0
    if true {
        v6 := []int{32768, -810}
        v7 := GenS1{F0: ("\n\t" + ""), F1: 1, F2: 2147483647}
        _ = v6
        _ = v7
    } else {
        v8 := GenS2{F0: (v5 & (v5 << ((18446744073709551615) & 63))), F1: 924, F2: -1}
        v5 = (p1 - p1)
        _ = v8
    }
    var v9 string = ""
    if true {
        v10 := []int{-231, 285}
        switch v9 {
        case "foo bar":
            var v11 string = v9
            _ = v11
        case "":
            v12 := GenS3{F0: 1}
            _ = v12
        case "hello":
            var v13 uint16 = 96
            _ = v13
        }
        switch 128 {
        case 32767:
            p1 += p1
        case -9223372036854775808:
            var v14 uint = 1
            _ = v14
        case 943:
            switch 575 {
            case -993:
                v10 = append(v10, 4294967295)
            case 128:
                var v15 int8 = 127
                _ = v15
            case 32767:
                _ = len(v10)
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    _ = len(v10)
                    _ = fi16
                }
            }
        default:
            var v17 bool = true
            _ = v17
        }
        _ = v10
    } else {
        if ((!(!(91 <= 4294967295))) && true) {
            switch 4294967295 {
            case 9223372036854775807:
                fmt.Println(v5, p1, p0)
            case -1:
                v18 := GenS2{F0: v5, F1: -44, F2: -1}
                _ = v18
            default:
                p1 = p1
            }
            switch "\n\t" {
            case "\n\t":
                var v19 int32 = 2147483647
                _ = v19
            case "a":
                var v20 string = "a"
                _ = v20
            case "hello":
                v21 := GenS4{F0: 258, F1: true, F2: 42}
                _ = v21
            }
            fmt.Println(p1, v9, v9)
        }
        v9 = ("a" + (v9 + v9))
        v9 = ((("a" + "a") + "hello") + "hello")
    }
    v22 := GenS2{F0: -869, F1: 42, F2: v5}
    _ = p0
    _ = p1
    _ = v22
    _ = v5
    _ = v9
    return 909
}
func genFn1(p0 int8) string {
    p0 = p0
    if false {
        for fi23 := 0; fi23 < 15; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int = 32768
                if (225 > 2) {
                    v26 := []int{(fi24 | (fi23 * (fi23 ^ -267))), (fi23 << ((18446744073709551615) & 63)), v25}
                    v25 = (v25 >> ((42) & 63))
                    fmt.Println(v25)
                    var v27 int8 = 42
                    _ = v26
                    _ = v27
                } else {
                    if true {
                        p0 = (p0 | p0)
                    }
                    if false {
                        var v28 bool = (1000 < 0)
                        p0 = (^(p0))
                        _ = v28
                    }
                    if false {
                        v29 := []int{}
                        _ = v29
                    } else {
                        var v30 uint8 = 128
                        v30 += (v30 | (v30 - v30))
                        if true {
                            var v31 int16 = 126
                            if false {
                                v25 = 1
                                v32 := []int{}
                                _ = v32
                            }
                            v33 := []int{-9223372036854775808, (fi24 ^ v25), 0, fi23}
                            _ = v31
                            _ = v33
                        }
                        if (0.1 >= float32((1e10 + 1.5))) {
                            v34 := []int{fi23, fi24, v25, (v25 >> ((255) & 63))}
                            _ = v34
                        } else {
                            var v35 float32 = (0.0 / 0.1)
                            fmt.Println(p0, fi24, p0)
                            var v36 int32 = 2147483647
                            p0 += (^(p0))
                            v37 := []int{v25, fi23}
                            v35 += v35
                            if ("" <= ("hello" + ("foo bar" + "a"))) {
                                v30 = 42
                            }
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = v30
                    }
                }
                _ = fi24
                _ = v25
            }
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 int32 = 1
                if true {
                    v40 := GenS2{F0: 0, F1: -32, F2: -1}
                    _ = v40
                } else {
                    if ((p0 >> ((1) & 63)) > (p0 * -116)) {
                        v41 := GenS2{F0: -307, F1: 32767, F2: 286}
                        _ = v41
                    }
                    fmt.Println(fi38, v39, p0)
                    p0 -= p0
                }
                v42 := GenS5{F0: false}
                _ = fi38
                _ = v39
                _ = v42
            }
            _ = fi23
        }
    } else {
        var v43 int8 = (p0 * -128)
        p0 -= (p0 & (p0 % (((^(p0))) | 1)))
        _ = v43
    }
    p0 = p0
    _ = p0
    return ""
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 int8 = 102
    v4 = (v4 | (v4 ^ v4))
    var v44 uint8 = 42
    v4 -= int8(v44)
    var v45 uint = 865
    var v46 uint16 = 524
    if false {
        switch 32767 {
        case -382:
            if (!false) {
                v47 := GenS6{F0: (v45 % ((v45) | 1)), F1: (v4 | (v4 % ((v4) | 1)))}
                for fi48 := 0; fi48 < 12; fi48++ {
                    v49 := GenS3{F0: 221}
                    _ = fi48
                    _ = v49
                }
                fmt.Println(v46, v46, v46)
                _ = v47
            }
        case -1:
            _ = len(v1)
        default:
            v4 = v4
        }
        v50 := GenS7{F0: v46, F1: 99, F2: 2147483647}
        _ = v50
    }
    v51 := GenS4{F0: 18446744073709551615, F1: false, F2: -1}
    var v52 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v51
    _ = v52
}
