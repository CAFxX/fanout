package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int8 }
func genFn0(p0 int16, p1 uint, p2 uint) uint8 {
    var v6 int16 = (-(p0))
    for fi7 := 0; fi7 < 3; fi7++ {
        p2 = (^(p1))
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint = (p1 | (^(p1)))
            if true {
                if false {
                    v6 = v6
                    if (2147483648 != 1) {
                        if (18446744073709551615 != 18446744073709551615) {
                            v10 := GenS1{F0: 222}
                            fmt.Println(v6, fi8, fi7)
                            _ = v10
                        } else {
                            var v11 uint64 = 18446744073709551615
                            _ = v11
                        }
                        if true {
                            var v12 int = (fi7 * fi7)
                            if false {
                                var v13 int64 = 837
                                v14 := GenS1{F0: 13}
                                _ = v13
                                _ = v14
                            } else {
                                v15 := GenS1{F0: 120}
                                if (659 == 1) {
                                    v16 := []int{2147483647, (v12 | (fi7 | 127)), fi7}
                                    var v17 string = ""
                                    _ = v16
                                    _ = v17
                                }
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    var v18 int32 = 227
                                    _ = v18
                                case "hello":
                                    v9 = (p1 ^ (p1 - (^(p2))))
                                default:
                                    var v19 uint32 = 2147483648
                                    _ = v19
                                }
                                _ = v15
                            }
                            if (1 != 375) {
                                v20 := GenS1{F0: 89}
                                v21 := []int{fi8, fi7}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 float32 = (0.0 / 3.14159)
                                _ = v22
                            }
                            _ = v12
                        } else {
                            var v23 uint = v9
                            var v24 uint32 = 889
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS1{F0: 1}
                        switch (fi8 % ((-732) | 1)) {
                        case 32767:
                            if (p1 > (p2 >> ((p1) & 63))) {
                                if ((1.5 >= ((1e10 - 1.5) + -2.25)) && true) {
                                    switch fi8 {
                                    case 1:
                                        var v26 int = (fi8 - fi7)
                                        fmt.Println(p0, fi7)
                                        _ = v26
                                    case -364:
                                        v9 = (^(v9))
                                    case 399:
                                        fmt.Println(fi8)
                                    default:
                                        fmt.Println(p1, p0)
                                    }
                                } else {
                                    p1 += 255
                                    v27 := GenS1{F0: 42}
                                    v28 := GenS2{F0: v6}
                                    _ = v27
                                    _ = v28
                                }
                            }
                        case -729:
                            p1 = (p1 + 42)
                        }
                        switch (fi7 * (fi8 * (fi8 / (((fi7 ^ fi7)) | 1)))) {
                        case 127:
                            p1 = v9
                        case -2147483648:
                            v29 := []int{(fi7 * fi7), (fi8 >> ((p2) & 63))}
                            _ = v29
                        }
                        _ = v25
                    }
                    v30 := []int{fi7, (fi7 / (((fi7 % ((fi7) | 1))) | 1)), fi7}
                    _ = v30
                } else {
                    v9 = 65535
                    switch "" {
                    case "":
                        v31 := GenS2{F0: p0}
                        v32 := GenS3{F0: "a"}
                        v33 := []int{}
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                v34 := []int{}
                _ = v34
            } else {
                p0 -= (v6 * 32767)
            }
            v9 = (^(p1))
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := []int{}
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            fmt.Println(fi35, p2)
            switch "hello" {
            case "":
                if false {
                    p1 += p1
                    p2 = 285
                }
            case "foo bar":
                v38 := GenS4{F0: 42, F1: 42}
                _ = v38
            case "\n\t":
                p1 = (^(p1))
            default:
                switch "foo bar" {
                case "foo bar":
                    v39 := GenS3{F0: ("hello" + ("a" + "\n\t"))}
                    _ = v39
                case "":
                    v6 = v6
                }
            }
            v40 := GenS3{F0: ((("\n\t" + "foo bar") + ("foo bar" + (("\n\t" + "a") + ("" + ((((((("" + ("" + (("foo bar" + "hello") + ("foo bar" + "hello")))) + (("a" + "a") + "foo bar")) + ((("a" + "foo bar") + "foo bar") + (("hello" + "hello") + "hello"))) + (("\n\t" + "\n\t") + ((("\n\t" + (("" + "\n\t") + "foo bar")) + ("hello" + (((("hello" + ("foo bar" + (("hello" + "\n\t") + ("foo bar" + "hello")))) + "a") + "a") + ("a" + "\n\t")))) + ("" + "a")))) + "hello") + "\n\t") + "foo bar"))))) + ("foo bar" + ("a" + "foo bar")))}
            fmt.Println(p1, p0)
            p2 = p2
            _ = fi37
            _ = v40
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v6, p2)
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := GenS5{F0: (p1 | (p1 ^ p2)), F1: 516}
        v43 := GenS6{F0: (false && (true || ((255 == 301) && false)))}
        v6 += p0
        _ = fi41
        _ = v42
        _ = v43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}
func genFn1(p0 uint, p1 uint64, p2 string) uint32 {
    v44 := []int{-2147483648, -9223372036854775808, 2147483647}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    var v45 uint8 = 255
    var v46 float64 = (0.1 * 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v45
    _ = v46
    return 339
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 2147483647
    var v3 int8 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    v4["foo bar"] = v2
    switch "a" {
    case "\n\t":
        v47 := GenS7{F0: (v2 & (v2 - genFn1(870, 42, "\n\t"))), F1: 65535}
        _ = v47
    }
    v48 := []int{0, -1, -47, 531}
    var v49 float64 = 1.5
    var v50 int8 = (v3 ^ v3)
    v51 := GenS8{F0: int8(v2)}
    v52 := GenS6{F0: true}
    fmt.Println(v1, v1, v49)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
