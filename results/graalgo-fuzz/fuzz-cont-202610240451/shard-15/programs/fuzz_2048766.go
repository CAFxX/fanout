package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 733
    var v3 uint = 368
    var v4 float32 = 0.1
    switch "" {
    case "hello":
        v2 += (v2 - (v2 + (v2 - v2)))
    case "foo bar":
        v3 -= 255
    case "\n\t":
        var v5 float64 = float64(3.14159)
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 127, F2: 423}
    var v7 uint64 = 619
    v8 := GenS2{F0: 100, F1: uint8(v2)}
    if false {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{629, fi9}
            v11 := GenS1{F0: v7, F1: 100, F2: 1}
            switch -85 {
            case 127:
                _ = len(v1)
            case 476:
                switch (^(fi9)) {
                case -895:
                    switch (fi9 - (fi9 >> ((4294967295) & 63))) {
                    case 1:
                        switch ("\n\t" + ("" + "foo bar")) {
                        case "a":
                            v4 = v4
                        case "":
                            var v12 float32 = v4
                            _ = v12
                        }
                    }
                case 4294967295:
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v4 = float32(float64(0.0))
                        v14 := GenS3{F0: 1, F1: 2147483647, F2: (fi9 & fi9)}
                        _ = fi13
                        _ = v14
                    }
                case -1:
                    v15 := GenS1{F0: (^(v7)), F1: -38, F2: -913}
                    _ = v15
                default:
                    fmt.Println(fi9)
                }
            case -9223372036854775808:
                var v16 uint16 = 1
                _ = v16
            default:
                var v17 float64 = (float64(0.0) + 1e10)
                _ = v17
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v4, v3, v2)
    }
    var v18 bool = (100 > -128)
    v19 := GenS4{F0: (v3 + (v3 >> (((^(v3))) & 63)))}
    var v20 uint16 = 42
    v21 := []int{}
    if len(v21) > 0 {
        v21[0] = -1
    }
    if ((v3 >> ((v3) & 63)) <= (^(v3))) {
        v22 := GenS5{F0: 42}
        if (42 > -1) {
            delete(v1, "a")
            var v23 uint = (v3 ^ v3)
            _ = v23
        } else {
            switch 577 {
            case 747:
                if true {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = fi24
                    }
                    var v25 uint64 = (v7 % ((0) | 1))
                    _ = v25
                }
            case 128:
                v18 = (127 >= 42)
            default:
                v21 = append(v21, 32767)
            }
            switch 32767 {
            case -541:
                if len(v21) > 0 {
                    v21[0] = 128
                }
            default:
                var v26 int = 32767
                _ = v26
            }
        }
        fmt.Println(v3, v20)
        _ = v22
    }
    v27 := GenS4{F0: v3}
    v20 -= 680
    var v28 int64 = int64(v3)
    var v29 uint32 = 2147483647
    if (v28 <= v28) {
        switch "a" {
        case "hello":
            if ((v29 << ((uint(v20)) & 63)) >= uint32(v3)) {
                for fi30 := 0; fi30 < 3; fi30++ {
                    if ("foo bar" == "a") {
                        v31 := GenS6{F0: (fi30 * (fi30 % ((128) | 1)))}
                        fmt.Println(v18)
                        var v32 int64 = v28
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(v18, v18, v28)
                    v33 := GenS7{F0: 0}
                    _ = fi30
                    _ = v33
                }
                if v18 {
                    v3 -= v3
                    v21 = append(v21, -435)
                }
            } else {
                v3 -= v3
                switch "foo bar" {
                case "":
                    var v34 int64 = (v28 | v28)
                    _ = v34
                case "hello":
                    if v18 {
                        v35 := v1["\n\t"]
                        fmt.Println(v29)
                        _ = v35
                    }
                case "foo bar":
                    if v18 {
                        if true {
                            var v36 bool = ((^(v7)) == (v7 + v7))
                            _ = len(v21)
                            for fi37 := 0; fi37 < 10; fi37++ {
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    var v39 bool = v18
                                    _ = fi38
                                    _ = v39
                                }
                                switch fi37 {
                                case -1:
                                    v7 += uint64(v29)
                                default:
                                    v40 := GenS5{F0: 888}
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 16; fi41++ {
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    if ((v7 - 556) <= v7) {
                                        var v42 uint8 = 1
                                        var v43 uint32 = (v29 >> ((uint(v29)) & 63))
                                        v44 := GenS8{F0: ("" + ("hello" + ("\n\t" + "a"))), F1: v29}
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    var v45 uint32 = (v29 - 0)
                                    var v46 int8 = int8(fi37)
                                    _ = fi41
                                    _ = v45
                                    _ = v46
                                }
                                _ = fi37
                            }
                            _ = v36
                        } else {
                            for fi47 := 0; fi47 < 18; fi47++ {
                                delete(v1, "a")
                                if ((uint(v7) <= (v3 / ((4294967295) | 1))) || ((v28 & v28) == 305)) {
                                    var v48 bool = v18
                                    for fi49 := 0; fi49 < 11; fi49++ {
                                        var v50 int16 = -1
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = v48
                                }
                                _ = fi47
                            }
                            var v51 float64 = float64(v4)
                            if len(v21) > 0 {
                                v21[0] = 128
                            }
                            _ = v51
                        }
                        v52 := v1[("" + ("a" + "hello"))]
                        _ = v52
                    } else {
                        var v53 string = "hello"
                        var v54 uint8 = 255
                        var v55 float32 = v4
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                default:
                    switch -2147483648 {
                    case -9223372036854775808:
                        for fi56 := 0; fi56 < 18; fi56++ {
                            if false {
                                var v57 uint = uint(v20)
                                v58 := GenS9{F0: -1}
                                _ = v57
                                _ = v58
                            } else {
                                delete(v1, ("foo bar" + "a"))
                            }
                            v28 += v28
                            fmt.Println(fi56)
                            _ = fi56
                        }
                    }
                }
            }
        case "":
            if (v20 < uint16(v7)) {
                v2 = 2
                switch "" {
                case "a":
                    v59 := GenS6{F0: 0}
                    _ = v59
                }
                v60 := GenS1{F0: 645, F1: 100, F2: 0}
                _ = v60
            } else {
                switch 869 {
                case 32767:
                    var v61 uint = (^(v3))
                    _ = v61
                }
                fmt.Println(v28, v3, v2)
                v62 := GenS9{F0: -852}
                _ = v62
            }
        default:
            if false {
                var v63 string = "hello"
                v64 := GenS10{F0: (v28 & (v28 | v28)), F1: 127}
                switch ((v63 + "hello") + "foo bar") {
                case "\n\t":
                    v4 = (((1.5 * v4) - (float32(((3.14159 * 0.1) * float64(float32(1e10)))) / float32((1e10 + (float64(v4) - 0.0))))) + 0.0)
                case "":
                    if ((-1 == 952) && (1000 > 1000)) {
                        v2 -= 1
                        var v65 int = 266
                        switch "foo bar" {
                        case "":
                            var v66 uint64 = v7
                            _ = v66
                        case "\n\t":
                            if ((1 > 0) && true) {
                                v67 := GenS11{F0: v2, F1: v20, F2: (v29 & uint32(v7))}
                                delete(v1, v63)
                                v68 := GenS1{F0: 42, F1: -128, F2: -1}
                                _ = v67
                                _ = v68
                            }
                        default:
                            _ = len(v21)
                        }
                        _ = v65
                    }
                }
                _ = v63
                _ = v64
            }
        }
        var v69 uint8 = 127
        _ = v69
    } else {
        v7 = 0
        var v70 string = ""
        _ = v70
    }
    v29 += v29
    v71 := GenS9{F0: 157}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v71
    _ = v8
}
