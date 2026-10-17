package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 int, p1 int8, p2 uint64) string {
    p2 += (p2 % (((p2 & (p2 | p2))) | 1))
    var v6 uint16 = 1
    v6 = v6
    v7 := []int{}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 uint16) uint32 {
    var v8 uint8 = 128
    var v9 uint = 875
    if true {
        switch ("foo bar" + "hello") {
        case "\n\t":
            if ((1e10 * 1.5) > (-2.25 / float32(1.5))) {
                switch 1 {
                case -726:
                    var v10 bool = false
                    _ = v10
                default:
                    v11 := GenS1{F0: ""}
                    _ = v11
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(v9, fi12, v8)
                    var v13 uint64 = 2
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 uint8 = v8
                v15 := GenS1{F0: "\n\t"}
                _ = v14
                _ = v15
            }
        case "hello":
            if (194 == 42) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 bool = (false && ((v9 + v9) > v9))
                    if (v17 && v17) {
                        fmt.Println(v17, p0)
                        v9 = 65535
                    } else {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            if (true || (0 > 42)) {
                                fmt.Println(p0, v8, p0)
                                v9 += v9
                            }
                            _ = fi18
                        }
                        if (((false || (!("" != "hello"))) || v17) || (!v17)) {
                            v17 = (!(0 > -1))
                        }
                        p0 = p0
                    }
                    _ = fi16
                    _ = v17
                }
            } else {
                v8 += (v8 - v8)
            }
        case "foo bar":
            v19 := GenS1{F0: ""}
            _ = v19
        default:
            if (true && ((v8 >> ((v9) & 63)) >= v8)) {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := []int{918, (fi20 << (((v9 + (^(v9)))) & 63)), fi20, fi20}
                    var v22 string = ""
                    switch fi20 {
                    case 4294967295:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            var v24 uint32 = 42
                            switch "" {
                            case "hello":
                                if (("" + v22) < v22) {
                                    v9 -= v9
                                    v25 := GenS1{F0: "foo bar"}
                                    fmt.Println(v24)
                                    _ = v25
                                }
                            }
                            fmt.Println(v9, p0, v8)
                            _ = fi23
                            _ = v24
                        }
                    default:
                        if (127 == (v8 ^ 255)) {
                            v8 = (v8 / ((127) | 1))
                            for fi26 := 0; fi26 < 16; fi26++ {
                                v27 := GenS2{F0: 65535, F1: 255}
                                var v28 uint32 = 98
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v9 = v9
                        }
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                p0 = (^(p0))
                switch ("foo bar" + "\n\t") {
                case "hello":
                    v29 := []int{-491, 32767, 128, 607}
                    _ = v29
                default:
                    v8 = (v8 - (v8 - 20))
                }
            }
        }
        v30 := GenS2{F0: 42, F1: (v9 | 847)}
        v31 := GenS3{F0: "foo bar", F1: 1}
        _ = v30
        _ = v31
    }
    _ = p0
    _ = v8
    _ = v9
    return 288
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 565
    v3 := []uint32{42, 2, 2147483648}
    v4 := make(map[uint32]bool)
    v5 := []int{-9223372036854775808}
    v32 := GenS3{F0: "", F1: 0}
    v2 -= v2
    var v33 int64 = 0
    delete(v4, 65535)
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    switch -2147483648 {
    case 4294967295:
        if (!true) {
            var v34 uint32 = 42
            var v35 int32 = 2147483647
            _ = v34
            _ = v35
        }
    case 904:
        if (!(47 == 1)) {
            switch "a" {
            case "hello":
                v33 = v2
            }
        }
    case 32767:
        for fi36 := 0; fi36 < 3; fi36++ {
            for fi37 := 0; fi37 < 1; fi37++ {
                fmt.Println(v33, fi37, fi37)
                _ = fi37
            }
            switch -815 {
            case 402:
                fmt.Println(fi36)
            }
            v2 += -1
            _ = fi36
        }
    }
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v33, v33)
    fmt.Println(v33)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
