package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 bool; F2 int8 }
type GenS7 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 string = "hello"
    var v4 uint32 = 616
    var v5 int32 = 2147483647
    var v6 int = int(v5)
    var v7 string = ("" + v3)
    var v8 bool = false
    var v9 int16 = 1000
    _ = len(v2)
    switch (v6 - int(v5)) {
    case 32767:
        v10 := []int{}
        _ = v10
    case 528:
        if v8 {
            var v11 uint16 = 0
            v12 := GenS1{F0: v8, F1: int(v9), F2: 200}
            v11 = v11
            _ = v11
            _ = v12
        }
    case -725:
        v4 += v4
    default:
        delete(v2, (v4 ^ v4))
    }
    v9 -= (v9 / ((v9) | 1))
    if ((1e10 >= 0.1) && (!v8)) {
        if v8 {
            switch -2147483648 {
            case -351:
                v9 -= v9
            case 4294967295:
                for fi13 := 0; fi13 < 14; fi13++ {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v15 := []int{v6, fi13, fi14, fi14}
                        if (!true) {
                            if false {
                                v2[(v4 + v4)] = (-(fi14))
                                switch v7 {
                                case "foo bar":
                                    v16 := v1[(v6 << ((164) & 63))]
                                    _ = v16
                                default:
                                    switch (("hello" + (v3 + v7)) + (v7 + ((v3 + (v7 + "hello")) + "foo bar"))) {
                                    case "":
                                        fmt.Println(v5, v3, v3)
                                    }
                                }
                                var v17 int16 = 0
                                _ = v17
                            }
                            switch (v7 + v7) {
                            case "":
                                switch v3 {
                                case "hello":
                                    v18 := GenS2{F0: 1, F1: 1, F2: 608}
                                    _ = v18
                                case "":
                                    fmt.Println(v4, v6)
                                case "a":
                                    if true {
                                        var v19 uint32 = v4
                                        _ = v19
                                    }
                                default:
                                    v6 = 0
                                }
                            case "a":
                                v20 := GenS2{F0: 133, F1: 1, F2: 42}
                                _ = v20
                            case "\n\t":
                                if v8 {
                                    if v8 {
                                        v3 = v3
                                        var v21 int = 195
                                        _ = v21
                                    }
                                    v22 := GenS1{F0: false, F1: (v6 ^ (fi14 + (fi14 | -172))), F2: 18446744073709551615}
                                    _ = v22
                                }
                            }
                            v23 := v2[2147483647]
                            _ = v23
                        } else {
                            v9 -= (v9 * 32767)
                        }
                        _ = fi14
                        _ = v15
                    }
                    var v24 uint32 = v4
                    v25 := GenS3{F0: (v9 * v9), F1: ((100 < 127) && v8), F2: 255}
                    _ = fi13
                    _ = v24
                    _ = v25
                }
            default:
                if false {
                    switch v7 {
                    case "":
                        var v26 uint32 = v4
                        _ = v26
                    case "\n\t":
                        v27 := []int{}
                        _ = v27
                    default:
                        v4 -= 65535
                    }
                    v28 := GenS1{F0: (-1 != 1), F1: (^(v6)), F2: 2}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 15; fi29++ {
                        if v8 {
                            v6 -= 804
                            v3 = "hello"
                        } else {
                            var v30 bool = ((42 < 481) || false)
                            _ = v30
                        }
                        _ = fi29
                    }
                    fmt.Println(v8)
                    var v31 float64 = float64((0.0 - 0.1))
                    _ = v31
                }
            }
            v32 := []int{v6, 32}
            if ((false && (!v8)) && v8) {
                v4 = (v4 | v4)
            } else {
                fmt.Println(v9, v8, v9)
                var v33 uint = 233
                v34 := GenS1{F0: ((1.5 * (3.14159 + (0.1 - 0.1))) > 1e10), F1: -534, F2: 2}
                _ = v33
                _ = v34
            }
            _ = v32
        }
        switch v7 {
        case "foo bar":
            switch (v6 & (v6 * -140)) {
            case 637:
                if (3.14159 < 0.1) {
                    v4 = 2147483648
                    v4 -= (v4 % (((v4 & 42)) | 1))
                }
            case 0:
                v35 := GenS4{F0: v5, F1: 128}
                _ = v35
            default:
                for fi36 := 0; fi36 < 7; fi36++ {
                    v8 = (3.14159 < 0.1)
                    v3 = ""
                    _ = fi36
                }
            }
        case "hello":
            v3 = v3
        }
        var v37 int32 = (v5 >> ((255) & 63))
        _ = v37
    } else {
        v4 = (v4 * (v4 | v4))
    }
    var v38 uint32 = (v4 / (((v4 ^ v4)) | 1))
    v3 = v7
    v8 = (9223372036854775807 != -9223372036854775808)
    v39 := GenS5{F0: uint8(v38)}
    v40 := []int{(v6 ^ (v6 & 4294967295))}
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS6{F0: "a", F1: (v8 || true), F2: -73}
        fmt.Println(v5, v5, v38)
        _ = fi41
        _ = v42
    }
    if true {
        v43 := GenS4{F0: (-(v5)), F1: (v6 * v6)}
        v5 -= (v5 + -108)
        var v44 string = v3
        _ = v43
        _ = v44
    }
    switch v7 {
    case "":
        for fi45 := 0; fi45 < 10; fi45++ {
            var v46 uint = 4294967295
            for fi47 := 0; fi47 < 2; fi47++ {
                v48 := GenS3{F0: (v9 + v9), F1: (-2.25 <= (((0.0 - 1e10) - 3.14159) + ((float64(((0.1 + (3.14159 + (-2.25 + 0.1))) - (3.14159 / 1e10))) / 1.5) - 0.0))), F2: 45}
                if (!(v8 || true)) {
                    switch (("a" + (v7 + (((v7 + v3) + "a") + v3))) + "foo bar") {
                    case "":
                        v49 := GenS3{F0: v9, F1: v8, F2: 73}
                        _ = v49
                    case "\n\t":
                        var v50 float64 = 3.14159
                        _ = v50
                    default:
                        var v51 int64 = 9223372036854775807
                        _ = v51
                    }
                    delete(v2, 65535)
                    switch (v7 + v3) {
                    case "\n\t":
                        v7 = v3
                    default:
                        switch fi47 {
                        case 73:
                            var v52 int32 = v5
                            _ = v52
                        case 2147483647:
                            if (!v8) {
                                fmt.Println(v3, v8)
                            }
                        case 128:
                            switch v7 {
                            case "foo bar":
                                if (!(!v8)) {
                                    fmt.Println(v7)
                                    switch (fi45 >> ((42) & 63)) {
                                    case -538:
                                        v53 := GenS4{F0: int32(v6), F1: (fi45 + 1)}
                                        _ = v53
                                    case 314:
                                        if len(v40) > 0 {
                                            v40[0] = -9223372036854775808
                                        }
                                    case 32767:
                                        var v54 uint64 = 2
                                        _ = v54
                                    }
                                } else {
                                    v40 = append(v40, fi47)
                                    v55 := GenS1{F0: (42 >= 42), F1: (^(fi47)), F2: 0}
                                    var v56 uint8 = 93
                                    _ = v55
                                    _ = v56
                                }
                            case "hello":
                                var v57 string = (v3 + v3)
                                switch v6 {
                                case -9223372036854775808:
                                    v8 = v8
                                }
                                v58 := GenS7{F0: 299, F1: (fi45 + 641)}
                                _ = v57
                                _ = v58
                            case "\n\t":
                                var v59 float32 = 1e10
                                _ = v59
                            }
                        }
                    }
                }
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
                _ = fi47
                _ = v48
            }
            v60 := GenS5{F0: 112}
            _ = fi45
            _ = v46
            _ = v60
        }
    case "foo bar":
        if true {
            fmt.Println(v3, v4)
        }
    }
    if v8 {
        v9 = (v9 & v9)
        v4 -= v4
        v61 := GenS3{F0: v9, F1: (v9 < (v9 % (((v9 | v9)) | 1))), F2: 2}
        _ = v61
    }
    if v8 {
        v62 := GenS3{F0: (^(v9)), F1: (v6 < (v6 & v6)), F2: 127}
        var v63 uint8 = 248
        _ = v62
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
