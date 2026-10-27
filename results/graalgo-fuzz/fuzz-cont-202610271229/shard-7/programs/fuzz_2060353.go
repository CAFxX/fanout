package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 string; F1 bool; F2 int }

func main() {
    var v1 bool = false
    var v2 uint32 = 65535
    var v3 int8 = 42
    var v4 int32 = -2147483648
    var v5 uint32 = 432
    for fi6 := 0; fi6 < 14; fi6++ {
        v1 = v1
        switch fi6 {
        case -2147483648:
            if true {
                var v7 int8 = -128
                _ = v7
            } else {
                v2 = v5
                v5 = 371
                var v8 uint16 = 993
                _ = v8
            }
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 1}
        switch fi9 {
        case 320:
            v11 := []int{fi9, (fi9 / ((fi9) | 1))}
            _ = v11
        case 128:
            v12 := GenS2{F0: "a"}
            _ = v12
        case -2147483648:
            var v13 int64 = -9223372036854775808
            _ = v13
        default:
            v14 := []int{fi9}
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 int32 = v4
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int64 = -872
            switch 32767 {
            case 4294967295:
                v19 := GenS3{F0: 979, F1: true}
                _ = v19
            case -9223372036854775808:
                v18 -= -9223372036854775808
            case 315:
                v20 := GenS3{F0: int16(v4), F1: v1}
                _ = v20
            }
            var v21 uint = 18446744073709551615
            _ = fi17
            _ = v18
            _ = v21
        }
        v22 := GenS3{F0: 827, F1: v1}
        _ = fi16
        _ = v22
    }
    v4 += v15
    v23 := []int{-748, -173, 2147483647}
    switch -9223372036854775808 {
    case 313:
        v24 := GenS3{F0: 347, F1: true}
        _ = v24
    case 32767:
        var v25 int64 = int64(v15)
        _ = v25
    }
    var v26 int32 = v4
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            v15 += v4
            v4 = v26
            _ = fi28
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS4{F0: -682, F1: 65535}
            switch "" {
            case "foo bar":
                v2 -= 785
            case "":
                if v1 {
                    switch (fi29 ^ fi29) {
                    case 42:
                        _ = len(v23)
                    case -1:
                        v26 -= (^(v26))
                    }
                }
            }
            _ = fi29
            _ = v30
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            v5 = (^(v2))
            if v1 {
                if (!v1) {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                } else {
                    v33 := GenS1{F0: (v2 | (v5 + (v2 + (v2 * v5))))}
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1)
                    case "\n\t":
                        switch 400 {
                        case 928:
                            switch "hello" {
                            case "foo bar":
                                v4 += (v26 << ((906) & 63))
                            }
                        case -195:
                            var v34 int32 = 755
                            _ = v34
                        case -595:
                            v3 = (v3 & v3)
                        }
                    default:
                        if false {
                            fmt.Println(v15, v26)
                            fmt.Println(fi31)
                            if len(v23) > 0 {
                                v23[0] = -758
                            }
                            var v35 uint64 = 936
                            var v36 uint32 = (v2 - v2)
                            _ = v35
                            _ = v36
                        } else {
                            if v1 {
                                var v37 int64 = 9223372036854775807
                                var v38 uint16 = 0
                                _ = v37
                                _ = v38
                            } else {
                                var v39 bool = (fi31 < (-(fi31)))
                                _ = v39
                            }
                            v26 -= (v4 ^ v15)
                            if false {
                                if (2 != 2) {
                                    v3 -= 1
                                    v26 = v26
                                    fmt.Println(v1)
                                    v23 = append(v23, fi32)
                                    var v40 int16 = 0
                                    v41 := GenS4{F0: (v40 / ((v40) | 1)), F1: v2}
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    _ = v40
                                    _ = v41
                                } else {
                                    v26 = v4
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    var v42 float32 = -2.25
                                    _ = v42
                                }
                            } else {
                                if (v1 || v1) {
                                    v2 = 1
                                }
                                fmt.Println(fi32, fi32)
                            }
                            _ = len(v23)
                            v1 = (4294967295 != 18446744073709551615)
                        }
                    }
                    fmt.Println(v3, fi32)
                    _ = v33
                }
                v3 = (-(v3))
                if true {
                    v1 = v1
                    if v1 {
                        v43 := GenS5{F0: v3, F1: 42}
                        var v44 int16 = -32768
                        _ = v43
                        _ = v44
                    }
                }
            } else {
                v45 := GenS6{F0: "hello", F1: v1, F2: 2147483647}
                _ = v45
            }
            _ = fi32
        }
        if false {
            v3 -= 35
            v5 = v2
            var v46 uint16 = 2
            _ = v46
        }
        v2 = 648
        _ = fi31
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
