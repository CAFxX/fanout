package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }
type GenS5 struct { F0 int8; F1 bool; F2 uint32 }
type GenS6 struct { F0 int8; F1 int8; F2 int8 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint16; F1 int; F2 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int16; F2 int8 }

func main() {
    var v1 int8 = -1
    var v2 uint8 = 255
    var v3 uint16 = 572
    v3 = (v3 * v3)
    fmt.Println(v3, v3, v2)
    v2 = v2
    var v4 string = "foo bar"
    v5 := GenS1{F0: -51, F1: 0}
    var v6 float64 = (((0.0 * float64(-2.25)) - float64(float32(3.14159))) * 0.0)
    if false {
        if false {
            if (v6 != v6) {
                v3 = v3
                v7 := GenS1{F0: 870, F1: -1}
                var v8 uint64 = 278
                _ = v7
                _ = v8
            } else {
                v6 -= (v6 - 1e10)
                fmt.Println(v2, v3, v3)
                var v9 int64 = 1
                _ = v9
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: 0, F1: 1}
                v12 := []int{}
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            if false {
                if false {
                    var v13 uint32 = 2147483647
                    fmt.Println(v3, v3)
                    v13 += 2147483648
                    _ = v13
                } else {
                    var v14 float64 = 1.5
                    v15 := []int{-1}
                    switch ("a" + "foo bar") {
                    case "":
                        _ = len(v15)
                    case "a":
                        if false {
                            var v16 float64 = v14
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            for fi17 := 0; fi17 < 17; fi17++ {
                                fmt.Println(v14, v1)
                                _ = fi17
                            }
                            _ = v16
                        }
                    default:
                        var v18 float32 = float32(v14)
                        _ = v18
                    }
                    _ = v14
                    _ = v15
                }
                switch "hello" {
                case "":
                    v19 := []int{42, 128, 32767, 859}
                    _ = v19
                default:
                    v3 = (^(v3))
                }
            } else {
                if (v3 != 255) {
                    if true {
                        v20 := GenS2{F0: uint16(v2), F1: -128, F2: 128}
                        _ = v20
                    }
                }
                switch "" {
                case "":
                    var v21 uint8 = (v2 ^ 42)
                    _ = v21
                case "\n\t":
                    switch v4 {
                    case "hello":
                        var v22 uint32 = 1
                        _ = v22
                    case "":
                        var v23 int16 = -611
                        _ = v23
                    default:
                        v24 := GenS3{F0: 0}
                        _ = v24
                    }
                default:
                    if false {
                        var v25 int64 = -9223372036854775808
                        _ = v25
                    }
                }
                v3 = v3
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                switch int(v2) {
                case 1:
                    if (0 >= 571) {
                        var v27 uint = 772
                        v27 = v27
                        if true {
                            if (false && (-537 <= (fi26 + fi26))) {
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    var v29 int = 0
                                    v2 += (v2 / (((v2 << (((v27 | 65535)) & 63))) | 1))
                                    v30 := GenS4{F0: (v27 << ((v27) & 63)), F1: v1, F2: 431}
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            }
                        } else {
                            switch fi26 {
                            case 32767:
                                fmt.Println(v3, v6, v4)
                            default:
                                v31 := GenS2{F0: (^(v3)), F1: (v1 / ((-109) | 1)), F2: 127}
                                _ = v31
                            }
                            var v32 bool = true
                            switch (fi26 | fi26) {
                            case 127:
                                v33 := []int{(fi26 & fi26), fi26, fi26, int(v3)}
                                _ = v33
                            }
                            _ = v32
                        }
                        _ = v27
                    }
                case 0:
                    if ((0.0 - (((-2.25 + ((1.5 - -2.25) * float32(v6))) * 0.1) / (-2.25 + float32(v6)))) == (((float32((-2.25 * float64(1.5))) - 3.14159) - -2.25) / 0.1)) {
                        var v34 float64 = 0.1
                        _ = v34
                    } else {
                        v35 := GenS3{F0: -305}
                        fmt.Println(fi26, v3)
                        _ = v35
                    }
                case -1:
                    switch (fi26 - (fi26 + 32767)) {
                    case 2147483647:
                        v2 = uint8(v3)
                    case 127:
                        var v36 uint8 = v2
                        _ = v36
                    }
                default:
                    var v37 float64 = v6
                    _ = v37
                }
                _ = fi26
            }
        }
        switch v4 {
        case "\n\t":
            v38 := GenS5{F0: (^(v1)), F1: (-2.25 >= (-2.25 / float32(v6))), F2: 2}
            _ = v38
        case "foo bar":
            var v39 int8 = (v1 % ((v1) | 1))
            _ = v39
        case "hello":
            var v40 int64 = -969
            _ = v40
        default:
            var v41 float64 = v6
            _ = v41
        }
    }
    v1 = (v1 / ((v1) | 1))
    v6 += 0.0
    v42 := []int{127, 958}
    v2 = 1
    if false {
        v43 := GenS4{F0: 457, F1: (v1 ^ (v1 >> ((23) & 63))), F2: 0}
        fmt.Println(v6, v3, v6)
        var v44 uint16 = (v3 * 65535)
        _ = v43
        _ = v44
    } else {
        for fi45 := 0; fi45 < 10; fi45++ {
            v1 += (v1 << ((42) & 63))
            var v46 uint8 = v2
            for fi47 := 0; fi47 < 15; fi47++ {
                v48 := GenS6{F0: 0, F1: v1, F2: v1}
                var v49 float32 = 3.14159
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi45
            _ = v46
        }
        if true {
            v42 = append(v42, 1)
            if true {
                v1 = int8(v3)
                if true {
                    v50 := GenS7{F0: 901, F1: -966}
                    var v51 uint16 = uint16(v2)
                    v52 := GenS8{F0: (^(v3)), F1: -149, F2: v3}
                    _ = v50
                    _ = v51
                    _ = v52
                }
            }
            v2 += v2
        } else {
            v53 := GenS7{F0: 1, F1: 0}
            v54 := GenS9{F0: 205}
            _ = v53
            _ = v54
        }
        var v55 int32 = 2147483647
        _ = v55
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        if true {
            var v57 int8 = (^(v1))
            if (false || (-2.25 < 0.0)) {
                for fi58 := 0; fi58 < 19; fi58++ {
                    fmt.Println(v2, v6)
                    _ = len(v42)
                    v57 += (v1 % ((v57) | 1))
                    _ = fi58
                }
                var v59 bool = false
                var v60 float64 = v6
                _ = v59
                _ = v60
            }
            switch "\n\t" {
            case "hello":
                var v61 float64 = 1e10
                _ = v61
            default:
                if len(v42) > 0 {
                    v42[0] = int(v57)
                }
            }
            _ = v57
        } else {
            v3 = 0
            if false {
                var v62 int32 = int32(fi56)
                var v63 int = fi56
                v64 := GenS10{F0: (v2 << ((255) & 63)), F1: 928, F2: (v1 ^ v1)}
                _ = v62
                _ = v63
                _ = v64
            }
        }
        v4 = v4
        _ = fi56
    }
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    v65 := GenS6{F0: (v1 % ((v1) | 1)), F1: 0, F2: v1}
    v2 = (v2 | (v2 - 2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v65
}
