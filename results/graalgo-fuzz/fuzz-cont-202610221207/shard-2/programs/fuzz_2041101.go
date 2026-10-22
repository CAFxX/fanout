package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int64 }

func main() {
    var v1 string = "hello"
    var v2 string = v1
    var v3 uint32 = 0
    v4 := []string{v1, v1}
    v5 := []int{-831, -1, 128, 4294967295}
    v6 := make(map[int]bool)
    if false {
        var v7 uint16 = 1
        _ = v7
    } else {
        if true {
            var v8 uint = 18446744073709551615
            var v9 int32 = 2147483647
            v8 = (v8 * (^(v8)))
            _ = v8
            _ = v9
        }
        var v10 string = v2
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 float32 = (3.14159 - (3.14159 / 1.0))
        _ = fi11
        _ = v12
    }
    if true {
        v3 = v3
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 float32 = (1e10 / -2.25)
                v2 = (("foo bar" + (("a" + ((v2 + v2) + "foo bar")) + v2)) + v2)
                _ = len(v6)
                _ = fi14
                _ = v15
            }
            v2 = v2
            v16 := GenS1{F0: ((v2 + v2) + "\n\t")}
            _ = fi13
            _ = v16
        }
        v2 = "\n\t"
    }
    v2 = v1
    switch v2 {
    case "hello":
        var v17 uint64 = 4294967295
        _ = v17
    case "a":
        if false {
            var v18 uint8 = 42
            v18 -= 127
            _ = v18
        } else {
            var v19 uint16 = 804
            v6[-854] = false
            switch "hello" {
            case "a":
                _ = len(v5)
            case "hello":
                v20 := GenS1{F0: ""}
                _ = v20
            }
            _ = v19
        }
    }
    if true {
        fmt.Println(v3, v1)
    } else {
        if (!(!(float32((0.1 / float64(1e10))) < 1.5))) {
            var v21 int8 = int8(v3)
            v1 = "\n\t"
            _ = v21
        }
        switch ("foo bar" + v1) {
        case "hello":
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: "hello"}
                _ = fi22
                _ = v23
            }
        case "":
            switch 642 {
            case 914:
                if (false && true) {
                    fmt.Println(v2, v3)
                    v2 = v1
                }
            default:
                if (136 < 1) {
                    fmt.Println(v3)
                    v2 = ("" + ((v1 + v1) + "\n\t"))
                } else {
                    var v24 uint16 = 0
                    switch "\n\t" {
                    case "":
                        v25 := GenS1{F0: v2}
                        _ = v25
                    case "\n\t":
                        if ((0.0 - (3.14159 + 1.5)) == -2.25) {
                            fmt.Println(v1, v3)
                            if (!false) {
                                v3 = v3
                            } else {
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    v27 := v6[fi26]
                                    _ = fi26
                                    _ = v27
                                }
                                var v28 uint16 = 843
                                _ = v28
                            }
                            var v29 int8 = 42
                            _ = v29
                        } else {
                            var v30 uint = 0
                            _ = v30
                        }
                    case "foo bar":
                        switch v2 {
                        case "a":
                            v24 = (^(v24))
                        case "\n\t":
                            for fi31 := 0; fi31 < 18; fi31++ {
                                fmt.Println(fi31)
                                for fi32 := 0; fi32 < 9; fi32++ {
                                    v33 := GenS2{F0: (true && true), F1: (fi32 >> ((255) & 63)), F2: (v24 ^ v24)}
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v24)
                                    case "foo bar":
                                        delete(v6, fi31)
                                    case "":
                                        _ = len(v6)
                                    }
                                    switch (fi32 * 617) {
                                    case 1:
                                        v34 := GenS2{F0: true, F1: (^(fi32)), F2: v24}
                                        _ = v34
                                    case 547:
                                        switch -9223372036854775808 {
                                        case 631:
                                            v2 = v1
                                        case 0:
                                            v3 = (v3 * v3)
                                        case -2147483648:
                                            v35 := GenS3{F0: (v3 % (((v3 >> ((701) & 63))) | 1)), F1: 42}
                                            _ = v35
                                        }
                                    default:
                                        v24 = v24
                                        switch ((("a" + ((v2 + (v1 + v2)) + "foo bar")) + v2) + v2) {
                                        case "\n\t":
                                            v2 = ("a" + "hello")
                                        case "a":
                                            fmt.Println(fi32, v2)
                                        }
                                        var v36 int64 = -1
                                        _ = v36
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                v37 := GenS4{F0: 0, F1: 988}
                                _ = fi31
                                _ = v37
                            }
                        case "hello":
                            v38 := GenS5{F0: 4294967295, F1: 0, F2: 421}
                            _ = v38
                        default:
                            var v39 uint = 158
                            _ = v39
                        }
                    default:
                        for fi40 := 0; fi40 < 10; fi40++ {
                            var v41 float64 = (0.1 + 1e10)
                            for fi42 := 0; fi42 < 13; fi42++ {
                                v3 += v3
                                var v43 int16 = -1
                                v24 -= v24
                                _ = fi42
                                _ = v43
                            }
                            fmt.Println(v24, v41)
                            _ = fi40
                            _ = v41
                        }
                    }
                    switch (v2 + (v2 + v1)) {
                    case "":
                        v24 = v24
                    case "hello":
                        var v44 uint32 = (v3 % ((uint32(v24)) | 1))
                        _ = v44
                    default:
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            var v45 int = 127
                            _ = v45
                        case 2147483647:
                            var v46 int32 = -317
                            _ = v46
                        }
                    }
                    _ = v24
                }
            }
        case "\n\t":
            v4 = append(v4, v2)
        }
    }
    var v47 int32 = 0
    var v48 int = 9223372036854775807
    v49 := GenS5{F0: 4294967295, F1: (v47 & (^(v47))), F2: 42}
    v50 := GenS6{F0: 232, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
