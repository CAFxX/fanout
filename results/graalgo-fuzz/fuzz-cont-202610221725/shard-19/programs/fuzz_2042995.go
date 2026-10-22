package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint16 }
type GenS5 struct { F0 string }

func main() {
    var v1 int8 = 8
    v2 := []string{("" + "hello"), "hello"}
    var v3 int32 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    v6 := []uint32{65535}
    delete(v4, 1)
    var v7 uint8 = 107
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int16 = 0
        v9 = v9
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -446}
    var v11 int64 = 9223372036854775807
    _ = len(v5)
    switch "foo bar" {
    case "\n\t":
        v12 := GenS1{F0: 1}
        _ = v12
    default:
        var v13 int32 = int32(v1)
        _ = v13
    }
    v1 += v1
    v14 := GenS2{F0: v7, F1: 65535}
    for fi15 := 0; fi15 < 16; fi15++ {
        if true {
            var v16 int16 = 1
            var v17 float64 = 3.14159
            v5[2147483648] = fi15
            _ = v16
            _ = v17
        } else {
            switch ("hello" + "a") {
            case "\n\t":
                var v18 uint8 = v7
                _ = v18
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v7, v11)
            case "\n\t":
                for fi19 := 0; fi19 < 2; fi19++ {
                    v20 := v4[fi15]
                    _ = fi19
                    _ = v20
                }
            }
            var v21 float64 = (-2.25 * 1e10)
            _ = v21
        }
        _ = fi15
    }
    var v22 uint64 = 2
    var v23 uint32 = 2147483647
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int16 = 32767
        for fi26 := 0; fi26 < 9; fi26++ {
            v25 += (v25 + 42)
            v27 := GenS3{F0: v23, F1: fi24, F2: 5}
            switch -9223372036854775808 {
            case 4294967295:
                var v28 string = "foo bar"
                _ = v28
            default:
                switch 32767 {
                case 32767:
                    fmt.Println(v11, v3)
                default:
                    v29 := GenS4{F0: "foo bar", F1: 65535, F2: 32767}
                    _ = v29
                }
                switch "" {
                case "a":
                    if true {
                        var v30 int = (fi26 << ((276) & 63))
                        switch (fi26 % ((457) | 1)) {
                        case -877:
                            var v31 int32 = v3
                            _ = v31
                        case 4294967295:
                            var v32 int64 = (v11 * (v11 ^ -1))
                            _ = v32
                        }
                        _ = v30
                    } else {
                        switch "a" {
                        case "hello":
                            if (false || (true || (2147483648 < v23))) {
                                if (false && ((false && (!((v22 != (v22 << ((2) & 63))) || (v7 < 255)))) && (v23 < (v23 ^ v23)))) {
                                    v11 -= (v11 + (^(v11)))
                                    if false {
                                        delete(v4, fi24)
                                    }
                                }
                                var v33 uint = 718
                                switch "" {
                                case "foo bar":
                                    v2 = append(v2, ("hello" + ("\n\t" + "\n\t")))
                                case "\n\t":
                                    v34 := GenS1{F0: fi24}
                                    if (!true) {
                                        switch (fi26 << (((v33 | 1)) & 63)) {
                                        case 0:
                                            v35 := GenS5{F0: ""}
                                            _ = v35
                                        default:
                                            v6 = append(v6, (^(v23)))
                                        }
                                        v11 = int64(v3)
                                        switch "foo bar" {
                                        case "hello":
                                            v2 = append(v2, ("a" + "a"))
                                        default:
                                            switch (fi24 >> (((v33 & v33)) & 63)) {
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "hello":
                                                    var v36 bool = ("\n\t" <= ("\n\t" + ("hello" + "a")))
                                                    _ = v36
                                                case "a":
                                                    switch ("a" + ("" + "\n\t")) {
                                                    case "":
                                                        var v37 uint = v33
                                                        _ = v37
                                                    default:
                                                        if (false && false) {
                                                            v23 = v23
                                                            var v38 int32 = (v3 ^ v3)
                                                            v38 = int32(v22)
                                                            v22 = (v22 * 1)
                                                            switch fi24 {
                                                            case 4294967295:
                                                                v39 := v5[42]
                                                                _ = v39
                                                            }
                                                            _ = v38
                                                        }
                                                    }
                                                case "foo bar":
                                                    v23 -= (v23 / ((272) | 1))
                                                default:
                                                    v40 := GenS3{F0: v23, F1: fi26, F2: 157}
                                                    _ = v40
                                                }
                                            }
                                        }
                                    } else {
                                        v22 += 495
                                        var v41 int32 = (v3 ^ v3)
                                        var v42 uint64 = v22
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v34
                                case "a":
                                    fmt.Println(v33, v33)
                                }
                                _ = v33
                            } else {
                                fmt.Println(v25, v23, v11)
                                delete(v4, (fi24 ^ 1))
                                v25 = (v25 * (v25 % ((1) | 1)))
                            }
                        default:
                            var v43 int32 = (v3 * (v3 >> ((255) & 63)))
                            _ = v43
                        }
                        v3 = v3
                    }
                default:
                    v44 := GenS4{F0: "foo bar", F1: 2147483647, F2: 1}
                    _ = v44
                }
            }
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    var v45 int64 = v11
    var v46 uint32 = v23
    var v47 int64 = v11
    switch 630 {
    case 32767:
        v48 := GenS3{F0: (v23 >> ((1) & 63)), F1: int(v3), F2: 0}
        _ = v48
    case -2147483648:
        v22 += (^(v22))
    default:
        v45 += (v11 << ((65535) & 63))
    }
    var v49 uint32 = (v23 & (v46 % ((2147483648) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
