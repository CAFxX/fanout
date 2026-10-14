package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{42, -259}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 int8 = 127
    var v6 uint64 = 735
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5, v6)
    _ = len(v4)
    if (((v6 < 42) || false) || ((true || (1 == uint32(v6))) && true)) {
        var v7 uint64 = 336
        _ = v7
    } else {
        switch "hello" {
        case "":
            switch 221 {
            case 900:
                for fi8 := 0; fi8 < 15; fi8++ {
                    v6 += (v6 / ((2) | 1))
                    switch fi8 {
                    case -126:
                        v6 += (v6 % ((1) | 1))
                    }
                    _ = fi8
                }
            case -2147483648:
                var v9 int64 = 1
                _ = v9
            }
        case "foo bar":
            if (0 >= 2) {
                fmt.Println(v6, v5)
            }
        case "\n\t":
            if (130 == 18446744073709551615) {
                if (-564 >= 1) {
                    v10 := GenS1{F0: 2, F1: (v6 ^ (v6 % (((v6 - 1)) | 1)))}
                    if true {
                        v2 = append(v2, -1)
                        v11 := GenS1{F0: 2, F1: (v6 << ((255) & 63))}
                        _ = v11
                    } else {
                        fmt.Println(v5, v6)
                    }
                    _ = v10
                } else {
                    v6 = (^(v6))
                    switch 42 {
                    case 49:
                        v12 := GenS1{F0: 42, F1: (v6 & v6)}
                        _ = v12
                    case 32767:
                        var v13 int32 = 2147483647
                        _ = v13
                    case 32768:
                        switch (("foo bar" + ("" + "a")) + ("\n\t" + (("" + "a") + ((("hello" + ("foo bar" + "")) + "a") + "")))) {
                        case "hello":
                            v14 := GenS2{F0: 255, F1: int64(v5)}
                            _ = v14
                        case "a":
                            var v15 uint64 = 345
                            _ = v15
                        case "\n\t":
                            v6 -= v6
                        default:
                            for fi16 := 0; fi16 < 7; fi16++ {
                                switch 0 {
                                case -2147483648:
                                    _ = len(v4)
                                case 127:
                                    if true {
                                        for fi17 := 0; fi17 < 8; fi17++ {
                                            v18 := GenS3{F0: -625}
                                            switch fi17 {
                                            case 32767:
                                                v19 := GenS1{F0: 42, F1: 1}
                                                _ = v19
                                            case 32768:
                                                _ = len(v1)
                                            case -2147483648:
                                                switch "foo bar" {
                                                case "a":
                                                    v20 := GenS1{F0: 0, F1: 18446744073709551615}
                                                    _ = v20
                                                }
                                            default:
                                                fmt.Println(v6, v6, fi17)
                                            }
                                            _ = fi17
                                            _ = v18
                                        }
                                        v21 := GenS2{F0: 0, F1: 42}
                                        v22 := GenS2{F0: 42, F1: 1}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                v5 = 1
                                if (false && true) {
                                    v2 = append(v2, 32768)
                                    if true {
                                        v23 := GenS4{F0: v6, F1: int(v6), F2: 848}
                                        var v24 int = 127
                                        v25 := GenS1{F0: 84, F1: v6}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    v26 := GenS5{F0: (^(v5)), F1: (!(!true)), F2: -595}
                                    _ = v26
                                }
                                _ = fi16
                            }
                        }
                    }
                    var v27 int16 = 32767
                    _ = v27
                }
            }
        default:
            if (-1 != 1) {
                var v28 bool = false
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v6 -= v6
                _ = v28
            }
        }
        fmt.Println(v6)
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS4{F0: 516, F1: (^(fi29)), F2: 1}
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v5, v6)
    var v31 bool = (true && false)
    for fi32 := 0; fi32 < 20; fi32++ {
        fmt.Println(v31, v5)
        _ = fi32
    }
    v1[42] = (!(40 >= 234))
    v33 := GenS6{F0: 0, F1: 128}
    v34 := GenS4{F0: (^(v6)), F1: 0, F2: 32767}
    var v35 float32 = 0.0
    switch -1 {
    case 127:
        var v36 uint8 = 1
        _ = v36
    }
    v6 = (v6 - v6)
    fmt.Println(v5, v5, v35)
    var v37 uint32 = 2
    v38 := GenS7{F0: 249, F1: (((1.5 - 0.0) + -2.25) < 3.14159)}
    var v39 float64 = (-2.25 * float64((1.5 * v35)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
