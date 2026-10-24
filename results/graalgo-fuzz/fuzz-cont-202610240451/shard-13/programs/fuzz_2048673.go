package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int64 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{4294967295}
    var v3 uint64 = 423
    var v4 uint64 = (v3 * v3)
    v5 := []uint32{2147483647, 359, 2147483648}
    v6 := make(map[uint32]string)
    if ((0.1 + float32(1.5)) >= (1.5 + 1e10)) {
        var v7 float32 = 0.1
        v8 := GenS1{F0: -1, F1: 90, F2: "hello"}
        _ = v7
        _ = v8
    }
    var v9 int = 42
    switch (v9 % ((v9) | 1)) {
    case 2147483647:
        if ("\n\t" == "foo bar") {
            if ((true && false) || true) {
                if false {
                    v10 := GenS2{F0: 749, F1: "foo bar"}
                    _ = v10
                } else {
                    switch v9 {
                    case 32767:
                        var v11 uint8 = 128
                        _ = v11
                    case 9223372036854775807:
                        var v12 float64 = (3.14159 / 1.5)
                        _ = v12
                    case -234:
                        v1[0] = 127
                    }
                }
                switch (("a" + "a") + ("hello" + (("foo bar" + "") + "a"))) {
                case "foo bar":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint8 = 255
                            if (42 < 42) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                switch (fi14 << ((1) & 63)) {
                                case 128:
                                    switch (fi14 << ((2) & 63)) {
                                    case -1:
                                        v16 := GenS1{F0: 127, F1: -27, F2: "hello"}
                                        v17 := v6[2147483648]
                                        var v18 bool = (!(!true))
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    case 42:
                                        if ((fi14 ^ v9) != (v9 * fi13)) {
                                            var v19 uint64 = (v4 << ((1) & 63))
                                            v19 = (v19 >> ((65535) & 63))
                                            _ = v19
                                        }
                                        _ = len(v2)
                                        var v20 uint64 = (v4 / ((v4) | 1))
                                        _ = v20
                                    case 517:
                                        var v21 int8 = -6
                                        _ = v21
                                    }
                                case 1:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 127:
                                    switch fi14 {
                                    case -2147483648:
                                        switch 32768 {
                                        case 1:
                                            switch 127 {
                                            case 4294967295:
                                                if (0 != -199) {
                                                    v2 = append(v2, 473)
                                                    var v22 int8 = 1
                                                    switch ("hello" + ("\n\t" + ("a" + "\n\t"))) {
                                                    case "\n\t":
                                                        v23 := v6[352]
                                                        _ = v23
                                                    case "a":
                                                        if true {
                                                            v24 := GenS3{F0: 159, F1: 394}
                                                            if true {
                                                                if false {
                                                                    v25 := GenS3{F0: 2, F1: -779}
                                                                    _ = v25
                                                                }
                                                                v4 = v4
                                                            } else {
                                                                switch ("\n\t" + "\n\t") {
                                                                case "foo bar":
                                                                    v15 += 255
                                                                }
                                                                v4 = (v4 - 4294967295)
                                                            }
                                                            _ = v24
                                                        }
                                                    case "":
                                                        if ("foo bar" <= "") {
                                                            var v26 int = 959
                                                            v9 = -9223372036854775808
                                                            _ = v26
                                                        }
                                                    }
                                                    _ = v22
                                                }
                                            case 42:
                                                var v27 int16 = 1
                                                _ = v27
                                            default:
                                                v9 = (v9 | 0)
                                            }
                                        case 935:
                                            fmt.Println(v15, fi13)
                                        case -1:
                                            v4 = (v3 * v3)
                                        default:
                                            var v28 int8 = 0
                                            _ = v28
                                        }
                                    default:
                                        fmt.Println(v15, fi14, v3)
                                        _ = len(v1)
                                        if len(v5) > 0 {
                                            v5[0] = 42
                                        }
                                        v4 = 42
                                        var v29 bool = (128 < (fi13 >> ((0) & 63)))
                                        _ = v29
                                    }
                                default:
                                    v30 := GenS1{F0: 127, F1: 51, F2: "\n\t"}
                                    _ = v30
                                }
                                v3 -= v3
                            } else {
                                var v31 string = "a"
                                _ = v31
                            }
                            v32 := GenS3{F0: 0, F1: int16(v15)}
                            v33 := GenS2{F0: 1, F1: "foo bar"}
                            var v34 uint = 846
                            _ = fi14
                            _ = v15
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v4 -= (v3 / ((v3) | 1))
                        _ = fi13
                    }
                case "hello":
                    if (!(int8(v9) <= 71)) {
                        v9 = v9
                    }
                case "a":
                    if false {
                        var v35 uint = 18446744073709551615
                        _ = len(v5)
                        _ = v35
                    }
                }
                v3 = v4
            } else {
                if true {
                    var v36 int8 = -128
                    _ = v36
                } else {
                    fmt.Println(v3, v3, v3)
                }
                var v37 float64 = (1.5 / 1.0)
                _ = v37
            }
            v38 := GenS1{F0: 1, F1: -65, F2: "hello"}
            _ = v38
        } else {
            v3 = 625
        }
    case 0:
        v4 = (v4 >> ((1) & 63))
    case 9223372036854775807:
        v39 := GenS2{F0: 255, F1: ""}
        _ = v39
    }
    if (2147483648 > 533) {
        if false {
            if true {
                v40 := v1[2]
                for fi41 := 0; fi41 < 3; fi41++ {
                    var v42 float32 = (1e10 * 0.1)
                    _ = fi41
                    _ = v42
                }
                if true {
                    delete(v1, 4294967295)
                }
                _ = v40
            } else {
                for fi43 := 0; fi43 < 16; fi43++ {
                    v4 = (v3 * v4)
                    _ = fi43
                }
            }
            switch (^(v9)) {
            case -2147483648:
                for fi44 := 0; fi44 < 5; fi44++ {
                    if true {
                        v45 := GenS4{F0: 2147483648, F1: 226, F2: -268}
                        _ = v45
                    }
                    v3 = (v3 & (^(v3)))
                    switch (v9 << ((2) & 63)) {
                    case -1:
                        v46 := GenS3{F0: 2, F1: 32767}
                        _ = v46
                    case -2147483648:
                        for fi47 := 0; fi47 < 9; fi47++ {
                            var v48 float32 = (float32((1e10 / 1.0)) / ((1e10 / -2.25) + 3.14159))
                            switch v9 {
                            case -1:
                                fmt.Println(v4, v3)
                                if (("\n\t" + "hello") >= (("" + "hello") + "hello")) {
                                    switch "a" {
                                    case "":
                                        v49 := GenS1{F0: 0, F1: 64, F2: "hello"}
                                        _ = v49
                                    case "hello":
                                        v50 := v1[262]
                                        _ = v50
                                    case "a":
                                        var v51 int16 = -1
                                        _ = v51
                                    }
                                    fmt.Println(fi47)
                                } else {
                                    v4 += 0
                                }
                            case -218:
                                v9 = 9223372036854775807
                            }
                            v5 = append(v5, 42)
                            _ = fi47
                            _ = v48
                        }
                    case 31:
                        switch "\n\t" {
                        case "\n\t":
                            if ((("a" + "a") + ("hello" + "a")) == "foo bar") {
                                var v52 int8 = 100
                                v4 = 0
                                _ = v52
                            } else {
                                _ = len(v6)
                                v53 := GenS1{F0: 100, F1: -52, F2: ""}
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        var v54 int16 = -32768
                                        _ = v54
                                    } else {
                                        if (true || (!false)) {
                                            for fi55 := 0; fi55 < 1; fi55++ {
                                                v4 = v4
                                                _ = fi55
                                            }
                                        } else {
                                            delete(v1, 65535)
                                        }
                                        v56 := GenS1{F0: 100, F1: 42, F2: "a"}
                                        _ = v56
                                    }
                                }
                                _ = v53
                            }
                        case "foo bar":
                            var v57 float64 = -2.25
                            _ = v57
                        case "":
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                        default:
                            fmt.Println(v4, v9, v9)
                        }
                    default:
                        fmt.Println(v9, fi44)
                    }
                    _ = fi44
                }
            case 42:
                for fi58 := 0; fi58 < 8; fi58++ {
                    v59 := GenS2{F0: 255, F1: "hello"}
                    v9 -= (v9 * (v9 + fi58))
                    var v60 int32 = -2147483648
                    _ = fi58
                    _ = v59
                    _ = v60
                }
            }
            for fi61 := 0; fi61 < 3; fi61++ {
                v62 := GenS5{F0: v9}
                _ = fi61
                _ = v62
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v9, v3, v4)
    var v63 bool = (343 < 1)
    switch ((("hello" + "hello") + "\n\t") + "\n\t") {
    case "hello":
        var v64 uint16 = 2
        _ = v64
    case "":
        if false {
            var v65 int8 = 1
            v2 = append(v2, 42)
            v63 = v63
            _ = v65
        }
    default:
        for fi66 := 0; fi66 < 2; fi66++ {
            fmt.Println(v63)
            v4 -= 2
            v4 -= 2
            _ = fi66
        }
    }
    var v67 uint8 = 174
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v67
    _ = v9
}
