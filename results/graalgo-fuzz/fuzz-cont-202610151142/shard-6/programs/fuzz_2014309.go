package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int16; F1 bool; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int16, p1 int16, p2 uint16) string {
    p2 += p2
    fmt.Println(p0, p0)
    v4 := []int{-580, 32768, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 1
    var v3 float32 = 1.5
    var v5 int = 4294967295
    switch "foo bar" {
    case "hello":
        _ = len(v1)
    case "\n\t":
        v2 += (^(v2))
    }
    var v6 uint = (^(v2))
    v7 := GenS1{F0: v6, F1: v5}
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{-2147483648}
        if (1 < 2) {
            v5 = -621
            v3 = (v3 - v3)
            var v10 uint = (v2 << ((v6) & 63))
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if (!(!((-32768 <= 42) && true))) {
        v2 = (v2 / ((1) | 1))
        var v11 float32 = 1.5
        _ = v11
    }
    v12 := GenS1{F0: v6, F1: v5}
    v13 := GenS1{F0: 577, F1: (v5 ^ (v5 << (((^(v6))) & 63)))}
    if (v5 <= v5) {
        if false {
            if ((0.1 >= (1.5 / float64(v3))) && true) {
                if true {
                    v1[1] = "a"
                    v6 = (v2 & v6)
                    v14 := []int{v5, v5, (v5 % ((-1) | 1)), (v5 ^ v5)}
                    _ = v14
                }
                var v15 int8 = -128
                for fi16 := 0; fi16 < 3; fi16++ {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v18 := []int{(^(v5))}
                        if (true || true) {
                            var v19 int64 = -21
                            var v20 float32 = v3
                            var v21 string = "foo bar"
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            v2 -= v2
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                _ = v15
            } else {
                v2 = (v6 ^ (v6 ^ (v6 >> ((v2) & 63))))
            }
            v2 = v2
        }
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint64 = 42
                v25 := GenS2{F0: 42, F1: 9223372036854775807, F2: -1}
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS3{F0: 1, F1: 0}
                _ = fi26
                _ = v27
            }
            switch v5 {
            case -9223372036854775808:
                if false {
                    switch (fi22 & (fi22 & -1)) {
                    case -971:
                        v28 := GenS4{F0: -956, F1: ("\n\t" > "\n\t"), F2: ((("" + "foo bar") + "\n\t") + "hello")}
                        _ = v28
                    default:
                        if (1 >= 928) {
                            v3 += v3
                            fmt.Println(fi22)
                            v29 := GenS5{F0: 1}
                            _ = v29
                        } else {
                            var v30 float64 = 0.0
                            _ = v30
                        }
                    }
                    switch "hello" {
                    case "hello":
                        v1[1] = "\n\t"
                    case "":
                        var v31 int8 = 0
                        _ = v31
                    default:
                        v5 = (v5 + fi22)
                    }
                    if ((4294967295 >= 2) || true) {
                        if (true || (2147483647 < 2147483647)) {
                            v32 := GenS5{F0: 2147483647}
                            v1[126] = ("a" + "foo bar")
                            if false {
                                v6 = (v2 + (v2 & v2))
                                switch "a" {
                                case "hello":
                                    v33 := v1[669]
                                    _ = v33
                                case "":
                                    var v34 string = "hello"
                                    _ = v34
                                }
                                for fi35 := 0; fi35 < 18; fi35++ {
                                    v36 := GenS6{F0: 7, F1: 1}
                                    if (255 > 0) {
                                        var v37 float32 = (v3 + v3)
                                        v38 := GenS3{F0: 42, F1: 40}
                                        v39 := v1[0]
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    } else {
                                        fmt.Println(v5, fi35)
                                        v40 := []int{4294967295, 32767, fi22}
                                        v2 -= v6
                                        _ = v40
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            }
                            _ = v32
                        }
                        if ((!false) && (!false)) {
                            v5 -= v5
                            v2 -= (v2 * v6)
                        }
                        switch "" {
                        case "\n\t":
                            v2 = 572
                        case "foo bar":
                            var v41 int16 = 859
                            _ = v41
                        }
                    }
                }
            case 9223372036854775807:
                v3 = -2.25
            case 127:
                v2 = 30
            }
            _ = fi22
        }
        v42 := GenS2{F0: 62, F1: 0, F2: -128}
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
