package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int }
func genFn0(p0 int8, p1 int8) int32 {
    var v4 bool = true
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 = p0
        var v6 int16 = 507
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 828
}

func main() {
    v1 := []uint32{2}
    var v2 uint32 = 4294967295
    v3 := make(map[int]string)
    var v7 uint32 = 2147483647
    var v8 int8 = 100
    if false {
        v9 := GenS1{F0: 65535, F1: -1, F2: 0}
        switch ((("" + "") + (("a" + ((("foo bar" + ("a" + "\n\t")) + (("a" + "a") + "")) + "hello")) + ((("a" + ("a" + "")) + "hello") + "hello"))) + "foo bar") {
        case "\n\t":
            v10 := v3[1]
            _ = v10
        default:
            var v11 float32 = ((float32(0.1) * 1e10) + float32((3.14159 + ((1e10 + (((float64(0.1) * (3.14159 * 0.1)) + -2.25) / 1.5)) / 1.5))))
            _ = v11
        }
        v8 += (v8 + -124)
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if true {
            switch fi12 {
            case -58:
                if false {
                    fmt.Println(v8)
                    switch (("\n\t" + ("\n\t" + "a")) + ("hello" + "a")) {
                    case "\n\t":
                        v13 := GenS2{F0: (v8 | 100)}
                        _ = v13
                    case "hello":
                        if false {
                            var v14 uint16 = 116
                            var v15 uint16 = v14
                            for fi16 := 0; fi16 < 9; fi16++ {
                                if true {
                                    var v17 int32 = 0
                                    if len(v1) > 0 {
                                        v1[0] = (^(v7))
                                    }
                                    _ = v17
                                } else {
                                    v18 := GenS2{F0: (v8 % (((-(v8))) | 1))}
                                    var v19 uint8 = 0
                                    _ = v18
                                    _ = v19
                                }
                                v20 := v3[(^(fi16))]
                                _ = fi16
                                _ = v20
                            }
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(v7)
                        }
                    case "a":
                        fmt.Println(v8, v7, fi12)
                    default:
                        v21 := GenS2{F0: 127}
                        _ = v21
                    }
                    if true {
                        if (((v8 == v8) || (2 <= 4294967295)) && ((0 > 2) || (fi12 > fi12))) {
                            if false {
                                v8 = v8
                                v8 = -11
                            } else {
                                v2 += v7
                                if (4294967295 < 65535) {
                                    if len(v1) > 0 {
                                        v1[0] = v7
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                var v22 int16 = 390
                                _ = v22
                            }
                            v7 -= v7
                            if (v8 == (v8 * (v8 ^ 127))) {
                                var v23 string = "a"
                                fmt.Println(v2, v8)
                                _ = v23
                            } else {
                                v24 := GenS2{F0: -16}
                                fmt.Println(v7)
                                _ = v24
                            }
                        }
                        for fi25 := 0; fi25 < 18; fi25++ {
                            _ = len(v3)
                            if (-1 < 1) {
                                v26 := GenS1{F0: 559, F1: 0, F2: 42}
                                v8 -= v8
                                v8 -= (v8 ^ v8)
                                _ = v26
                            }
                            _ = fi25
                        }
                        v8 = (v8 + v8)
                    }
                }
            }
            v8 = v8
            v27 := GenS3{F0: 659, F1: (v7 * v7), F2: 0}
            _ = v27
        }
        _ = fi12
    }
    var v28 uint64 = 757
    delete(v3, -9223372036854775808)
    v1 = append(v1, (v2 >> ((2) & 63)))
    if true {
        switch "a" {
        case "foo bar":
            v28 += v28
        case "hello":
            v2 -= (v2 - 0)
        default:
            fmt.Println(v8, v7, v8)
        }
        var v29 string = "\n\t"
        _ = v29
    }
    v30 := GenS4{F0: 65535}
    v2 = 492
    switch (((("a" + (("\n\t" + ("" + "\n\t")) + (("foo bar" + "a") + ""))) + "\n\t") + "hello") + "foo bar") {
    case "\n\t":
        for fi31 := 0; fi31 < 18; fi31++ {
            v2 = v2
            switch "foo bar" {
            case "foo bar":
                var v32 float64 = 0.1
                _ = v32
            case "\n\t":
                v33 := GenS5{F0: 971, F1: 34, F2: fi31}
                _ = v33
            case "hello":
                switch fi31 {
                case -912:
                    v7 = v2
                case 0:
                    v7 = (v2 / (((v7 >> ((101) & 63))) | 1))
                case 128:
                    v2 = (v2 | (v7 >> ((917) & 63)))
                default:
                    v34 := GenS4{F0: 825}
                    _ = v34
                }
            }
            _ = fi31
        }
    }
    delete(v3, -544)
    var v35 uint64 = v28
    switch 1 {
    case -963:
        v28 += 2
    case 9223372036854775807:
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 5; fi37++ {
                fmt.Println(fi36, fi36)
                v8 += 32
                v8 = v8
                _ = fi37
            }
            for fi38 := 0; fi38 < 11; fi38++ {
                v39 := GenS2{F0: (v8 | v8)}
                _ = fi38
                _ = v39
            }
            _ = fi36
        }
    default:
        var v40 int32 = 400
        _ = v40
    }
    v2 += (v7 & (v7 - (v7 << ((4294967295) & 63))))
    v41 := GenS1{F0: 42, F1: 1000, F2: (v28 << ((65535) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v35
    _ = v41
    _ = v7
    _ = v8
}
