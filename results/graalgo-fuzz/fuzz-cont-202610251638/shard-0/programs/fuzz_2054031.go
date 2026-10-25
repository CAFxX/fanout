package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 int16 }
func genFn0(p0 int32) uint32 {
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 17; fi6++ {
            if true {
                for fi7 := 0; fi7 < 2; fi7++ {
                    v8 := GenS1{F0: 2, F1: 4294967295}
                    v9 := GenS1{F0: 1, F1: 42}
                    p0 = (p0 + (p0 - (p0 * p0)))
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                if (!true) {
                    v10 := []int{(fi6 | 267), fi6, 16, fi6}
                    _ = v10
                }
                v11 := []int{(fi6 << ((2) & 63))}
                _ = v11
            }
            p0 += p0
            _ = fi6
        }
    case "":
        v12 := []int{}
        _ = v12
    case "\n\t":
        p0 -= (p0 + (p0 | -1))
    }
    if (!((1 != 393) || (2147483647 != 32767))) {
        if ((-(p0)) == (-(p0))) {
            v13 := GenS2{F0: -128, F1: ("hello" >= ("\n\t" + ("a" + (("\n\t" + ("\n\t" + ("\n\t" + ("a" + "hello")))) + ""))))}
            _ = v13
        } else {
            v14 := []int{9223372036854775807}
            fmt.Println(p0)
            if (47 != 7) {
                switch "foo bar" {
                case "a":
                    p0 += (p0 % ((906) | 1))
                }
                v15 := GenS3{F0: 586, F1: 150}
                var v16 uint16 = 19
                _ = v15
                _ = v16
            }
            _ = v14
        }
    }
    v17 := []int{127, 32768, -1}
    _ = p0
    _ = v17
    return 65535
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", (("\n\t" + "foo bar") + (("a" + "\n\t") + "\n\t"))}
    v4 := make(map[uint32]uint32)
    var v5 float64 = 0.1
    switch "" {
    case "\n\t":
        fmt.Println(v1, v1)
    case "foo bar":
        v1 = v1
    case "hello":
        var v18 int32 = 2147483647
        _ = v18
    }
    v4[uint32(v1)] = 1
    var v19 uint32 = 2147483648
    v20 := GenS1{F0: 274, F1: 42}
    for fi21 := 0; fi21 < 1; fi21++ {
        if true {
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := v2["a"]
                v24 := GenS2{F0: 42, F1: false}
                fmt.Println(fi22, v23, v5)
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(v5, fi21)
        } else {
            if false {
                if false {
                    var v25 int16 = -32768
                    if true {
                        v5 = ((v5 + 3.14159) / (float64(1.5) - v5))
                        v26 := GenS4{F0: v19, F1: (^(v25)), F2: 127}
                        switch "\n\t" {
                        case "":
                            v27 := GenS1{F0: 0, F1: 42}
                            _ = v27
                        case "\n\t":
                            fmt.Println(fi21, v5)
                        default:
                            var v28 bool = false
                            _ = v28
                        }
                        _ = v26
                    }
                    _ = v25
                }
            } else {
                v1 += (v1 & v1)
                if false {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        if (!(false && (767 > 0))) {
                            var v30 int8 = 1
                            v5 -= (v5 - v5)
                            v4[v19] = (v19 % ((65535) | 1))
                            v31 := GenS5{F0: 2, F1: 32767, F2: true}
                            _ = v30
                            _ = v31
                        } else {
                            v5 -= v5
                        }
                        var v32 uint16 = 65535
                        fmt.Println(v1, fi29)
                        if (0.0 != (3.14159 / (v5 - ((v5 + v5) * v5)))) {
                            v32 = 311
                            if ((true || true) || true) {
                                switch (("" + "\n\t") + ("\n\t" + "a")) {
                                case "":
                                    v33 := GenS6{F0: 973}
                                    _ = v33
                                }
                                var v34 uint16 = v32
                                _ = v34
                            }
                            fmt.Println(fi29, v32)
                        } else {
                            v35 := GenS7{F0: (^(v32))}
                            _ = v35
                        }
                        _ = fi29
                        _ = v32
                    }
                    var v36 uint16 = 780
                    fmt.Println(v5, v19, v1)
                    _ = v36
                } else {
                    var v37 int = -330
                    _ = v37
                }
                _ = len(v3)
            }
            v19 = 552
        }
        if ((float64(3.14159) + (((-2.25 / v5) * 0.1) - 0.1)) == v5) {
            var v38 uint64 = 42
            switch fi21 {
            case 127:
                v19 += (^(v19))
            case 42:
                switch ("a" + "") {
                case "foo bar":
                    switch (fi21 >> ((42) & 63)) {
                    case 42:
                        var v39 uint64 = v38
                        _ = v39
                    case 128:
                        fmt.Println(v19)
                    default:
                        v19 = v19
                    }
                case "hello":
                    v40 := GenS8{F0: 1, F1: 1, F2: 2}
                    _ = v40
                case "\n\t":
                    v38 -= 4294967295
                }
            case -2147483648:
                delete(v2, "")
            }
            _ = v38
        }
        var v41 int8 = 63
        _ = fi21
        _ = v41
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        var v43 uint = (v1 >> (((^(v1))) & 63))
        if ((v5 / v5) >= ((v5 / v5) - float64(0.1))) {
            fmt.Println(fi42, v5)
            v44 := GenS9{F0: 0}
            _ = v44
        } else {
            v45 := GenS2{F0: -128, F1: (false && (v5 <= v5))}
            v46 := GenS7{F0: 65535}
            if (true || (genFn0(995) != v19)) {
                fmt.Println(v5)
            } else {
                delete(v2, ("" + ("\n\t" + "hello")))
                v47 := GenS10{F0: (fi42 * fi42), F1: 42}
                _ = len(v4)
                _ = v47
            }
            _ = v45
            _ = v46
        }
        v43 = 470
        _ = fi42
        _ = v43
    }
    switch "hello" {
    case "":
        var v48 float64 = ((v5 - float64(3.14159)) - v5)
        _ = v48
    }
    v1 = (v1 | (v1 * 497))
    fmt.Println(v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
