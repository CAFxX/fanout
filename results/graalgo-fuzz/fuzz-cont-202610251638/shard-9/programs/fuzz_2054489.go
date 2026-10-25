package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint; F2 bool }
type GenS7 struct { F0 bool; F1 string; F2 int64 }
func genFn0(p0 int, p1 uint16) int8 {
    if ((false && false) && true) {
        var v5 string = ("a" + ("\n\t" + ("a" + ("a" + "foo bar"))))
        v6 := []int{22, 42}
        v5 = v5
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v7
    return -41
}
func genFn1() uint64 {
    v8 := []int{883, 32767}
    var v9 int64 = 1
    switch 42 {
    case -561:
        switch "\n\t" {
        case "":
            v10 := GenS2{F0: 83, F1: 4294967295}
            _ = v10
        case "foo bar":
            v9 = (v9 / (((v9 * v9)) | 1))
        case "a":
            var v11 int64 = v9
            _ = v11
        default:
            v9 = (v9 / ((v9) | 1))
        }
    default:
        v9 = (^(v9))
    }
    v12 := GenS3{F0: 1, F1: 1}
    _ = v12
    _ = v8
    _ = v9
    return 851
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    var v3 float64 = float64(0.1)
    var v4 uint32 = 2
    fmt.Println(v3)
    for fi13 := 0; fi13 < 18; fi13++ {
        if true {
            fmt.Println(v4)
            var v14 uint = 42
            _ = v14
        } else {
            fmt.Println(v3, v3)
            if false {
                v3 -= float64((3.14159 / 1.0))
            }
            v15 := GenS1{F0: 179}
            _ = v15
        }
        if (((-9223372036854775808 == 479) || (!(-105 == -1))) || true) {
            var v16 int = fi13
            if false {
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int64 = 1
                    v19 := GenS3{F0: 329, F1: 1}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if (0 == 742) {
                    v4 += v4
                    var v20 int8 = -128
                    _ = v20
                } else {
                    if true {
                        switch ("foo bar" + "\n\t") {
                        case "":
                            var v21 string = "\n\t"
                            _ = v21
                        }
                        v3 = (((0.0 - (float64(0.1) - v3)) + 3.14159) + float64(((((((1.5 - (1.5 + (3.14159 + (1e10 + float32(0.0))))) - (1e10 / -2.25)) + (1e10 + float32(v3))) * ((1.5 - ((0.0 + (0.1 * (1.5 + (1.5 - (-2.25 + (3.14159 - 1e10)))))) + 1e10)) - 1.5)) / (1.5 / (((1.5 - float32(v3)) * (((0.1 * ((((3.14159 + (0.1 + 0.1)) / (0.1 * 0.1)) * float32((-2.25 / float64(float32(v3))))) / 1.5)) / 1.0) + float32(1e10))) - (1e10 - 0.1)))) / 1.0)))
                        v4 += (v4 / ((v4) | 1))
                    } else {
                        v16 -= 2147483647
                        for fi22 := 0; fi22 < 16; fi22++ {
                            var v23 uint64 = 2
                            _ = fi22
                            _ = v23
                        }
                    }
                    switch "" {
                    case "\n\t":
                        for fi24 := 0; fi24 < 5; fi24++ {
                            if false {
                                if true {
                                    v25 := GenS4{F0: "a"}
                                    v26 := GenS5{F0: 0}
                                    if ((v4 % ((v4) | 1)) <= (^(v4))) {
                                        var v27 int = 0
                                        var v28 int16 = 1000
                                        v28 -= v28
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v25
                                    _ = v26
                                }
                                if false {
                                    if true {
                                        v3 += (v3 * (((v3 - (float64(1e10) - v3)) - (1e10 / (v3 * v3))) / float64(1e10)))
                                        v3 = float64((1.5 * (1e10 * (1.5 / float32(v3)))))
                                        var v29 uint32 = (v4 + v4)
                                        _ = v29
                                    }
                                    if false {
                                        fmt.Println(v3, v4)
                                    }
                                    v16 += fi24
                                }
                            }
                            v4 = (v4 % (((^(v4))) | 1))
                            _ = fi24
                        }
                    default:
                        var v30 int8 = 127
                        _ = v30
                    }
                    switch (v16 | (fi13 - fi13)) {
                    case 231:
                        var v31 float64 = v3
                        _ = v31
                    case 128:
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v33 := GenS4{F0: "\n\t"}
                            v34 := GenS3{F0: (v4 + (v4 << ((255) & 63))), F1: -32768}
                            v35 := GenS4{F0: ("foo bar" + "a")}
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    default:
                        var v36 int = int(v4)
                        _ = v36
                    }
                }
                switch "\n\t" {
                case "hello":
                    v16 -= 788
                }
            } else {
                fmt.Println(fi13, v16, v4)
            }
            _ = v16
        } else {
            if len(v1) > 0 {
                v1[0] = (v4 * 4294967295)
            }
        }
        _ = fi13
    }
    var v37 string = ("" + "a")
    v38 := GenS6{F0: 4294967295, F1: 747, F2: false}
    if true {
        _ = len(v1)
        if false {
            var v39 string = v37
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 uint16 = 32767
                v3 += float64(1.5)
                _ = fi40
                _ = v41
            }
            var v42 float32 = 1e10
            _ = v39
            _ = v42
        } else {
            v43 := GenS7{F0: ((-71 > 880) && false), F1: "\n\t", F2: 42}
            for fi44 := 0; fi44 < 18; fi44++ {
                if (-1 != -1) {
                    v3 = 1e10
                    var v45 bool = false
                    _ = v45
                }
                fmt.Println(v4, v3, v3)
                _ = len(v2)
                _ = fi44
            }
            v46 := GenS3{F0: (v4 * v4), F1: -626}
            _ = v43
            _ = v46
        }
    }
    v37 = "foo bar"
    fmt.Println(v4, v4, v37)
    var v47 uint64 = 0
    v47 -= (v47 | v47)
    var v48 int = 177
    var v49 int16 = -32768
    var v50 int64 = 1
    v47 = (^(v47))
    if true {
        var v51 int64 = (v50 ^ int64(v49))
        v37 = "hello"
        _ = v51
    }
    v49 = v49
    if (false && (false || true)) {
        fmt.Println(v49, v48)
    } else {
        if true {
            v50 = 1
            fmt.Println(v49, v3)
        }
        var v52 int = int(v47)
        fmt.Println(v48, v3)
        _ = v52
    }
    fmt.Println(v48)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
