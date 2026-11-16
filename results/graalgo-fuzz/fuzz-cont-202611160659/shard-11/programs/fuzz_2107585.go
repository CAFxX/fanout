package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 string }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = v1
    var v3 uint = 886
    fmt.Println(v3, v1, v3)
    v4 := []int{-2147483648}
    v1 = v2
    v1 = v1
    switch ("hello" + "\n\t") {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        var v5 string = "hello"
        _ = v5
    default:
        fmt.Println(v3)
    }
    if false {
        v2 = v1
        if ((128 != 253) && (true && (42 < -32768))) {
            v1 -= (v1 ^ (v1 ^ v1))
            _ = len(v4)
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 bool = false
                for fi8 := 0; fi8 < 18; fi8++ {
                    switch fi8 {
                    case 51:
                        if true {
                            var v9 uint64 = 4294967295
                            switch (fi8 % ((32768) | 1)) {
                            case -1:
                                v10 := GenS1{F0: v9}
                                _ = v10
                            case 9223372036854775807:
                                switch 9223372036854775807 {
                                case 42:
                                    switch (fi8 / (((fi8 - -203)) | 1)) {
                                    case 2147483647:
                                        switch ("\n\t" + "a") {
                                        case "":
                                            v4 = append(v4, fi6)
                                        case "hello":
                                            fmt.Println(v3, v9)
                                            fmt.Println(v2, v3)
                                        case "\n\t":
                                            switch int(v1) {
                                            case -412:
                                                var v11 bool = false
                                                v7 = v7
                                                switch "" {
                                                case "a":
                                                    var v12 uint16 = 340
                                                    _ = v12
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "":
                                                        var v13 bool = (v1 != (^(v2)))
                                                        _ = v13
                                                    case "a":
                                                        v7 = (!(false || (-9223372036854775808 == -108)))
                                                    case "foo bar":
                                                        fmt.Println(v3, v7, v1)
                                                    default:
                                                        var v14 int64 = -9223372036854775808
                                                        _ = v14
                                                    }
                                                default:
                                                    switch fi6 {
                                                    case 9223372036854775807:
                                                        if (v3 < 127) {
                                                            var v15 float32 = 1.5
                                                            _ = v15
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    case 1:
                                                        switch fi6 {
                                                        case 810:
                                                            v2 = (^(v2))
                                                        case -2147483648:
                                                            v16 := GenS2{F0: 555, F1: -1, F2: ""}
                                                            _ = v16
                                                        case 42:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    case -9223372036854775808:
                                                        v4 = append(v4, 623)
                                                    default:
                                                        v17 := GenS2{F0: v2, F1: 100, F2: ("foo bar" + ("a" + "foo bar"))}
                                                        _ = v17
                                                    }
                                                }
                                                _ = v11
                                            case 32767:
                                                v7 = ((^(v3)) == v3)
                                            case 152:
                                                fmt.Println(v1)
                                            default:
                                                if (!(42 <= 510)) {
                                                    _ = len(v4)
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                } else {
                                                    var v18 int8 = 127
                                                    var v19 int16 = 1000
                                                    var v20 uint16 = (v1 + v2)
                                                    var v21 uint = 42
                                                    switch (fi8 * fi6) {
                                                    case 590:
                                                        v1 -= v2
                                                    case -1:
                                                        var v22 uint = v21
                                                        _ = v22
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                    _ = v20
                                                    _ = v21
                                                }
                                            }
                                        }
                                    case -2147483648:
                                        if ((-(fi6)) != 568) {
                                            var v23 bool = (v3 != uint(v1))
                                            v24 := GenS3{F0: fi8, F1: -172}
                                            _ = v23
                                            _ = v24
                                        } else {
                                            fmt.Println(v3, v1, fi8)
                                            switch -1 {
                                            case -253:
                                                v2 = (v2 | 32767)
                                            case 144:
                                                fmt.Println(v7)
                                                fmt.Println(v3, v7)
                                            case 32767:
                                                var v25 uint = 0
                                                _ = v25
                                            }
                                        }
                                    }
                                case -366:
                                    v26 := GenS3{F0: (fi8 - -271), F1: 371}
                                    fmt.Println(fi6)
                                    _ = v26
                                case -382:
                                    var v27 int64 = 42
                                    v9 += (v9 * v9)
                                    fmt.Println(fi8, fi6)
                                    var v28 float32 = 3.14159
                                    v4 = append(v4, -60)
                                    _ = v27
                                    _ = v28
                                default:
                                    v3 -= v3
                                }
                            }
                            _ = v9
                        } else {
                            var v29 uint16 = (v1 | (v1 % (((v1 + v1)) | 1)))
                            _ = v29
                        }
                        v1 += (v1 & v2)
                        var v30 float32 = 3.14159
                        _ = v30
                    case -2147483648:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 4294967295:
                        v7 = v7
                    default:
                        if (fi8 <= 2147483647) {
                            if v7 {
                                fmt.Println(v7, v2)
                            } else {
                                var v31 uint32 = 2
                                var v32 uint8 = 136
                                v33 := GenS2{F0: 65535, F1: 127, F2: "foo bar"}
                                switch (fi6 / ((fi8) | 1)) {
                                case 641:
                                    v2 = v1
                                case 42:
                                    v32 = (^(v32))
                                }
                                v2 += 255
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            v34 := GenS3{F0: 2147483647, F1: 42}
                            _ = v34
                        }
                    }
                    var v35 uint32 = 2147483648
                    _ = fi8
                    _ = v35
                }
                _ = fi6
                _ = v7
            }
        }
    }
    if ((3.14159 + 1.5) == 0.1) {
        v36 := GenS3{F0: 128, F1: 43}
        var v37 uint = 18446744073709551615
        _ = v36
        _ = v37
    } else {
        var v38 float32 = -2.25
        _ = v38
    }
    switch 1 {
    case 0:
        v39 := GenS4{F0: -9, F1: v2, F2: 32767}
        _ = v39
    case 681:
        var v40 int64 = -155
        _ = v40
    default:
        v2 -= 65535
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        fmt.Println(fi41, v1)
        var v42 string = ("hello" + ("foo bar" + "hello"))
        _ = fi41
        _ = v42
    }
    if (127 != 128) {
        for fi43 := 0; fi43 < 18; fi43++ {
            for fi44 := 0; fi44 < 6; fi44++ {
                var v45 int8 = -1
                _ = fi44
                _ = v45
            }
            for fi46 := 0; fi46 < 2; fi46++ {
                v1 -= 255
                v3 = 65535
                v2 -= (v2 & 2)
                _ = fi46
            }
            v47 := GenS1{F0: 1}
            _ = fi43
            _ = v47
        }
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 int8 = 100
            for fi50 := 0; fi50 < 8; fi50++ {
                v51 := GenS2{F0: 255, F1: v49, F2: "hello"}
                _ = fi50
                _ = v51
            }
            v3 -= (v3 + v3)
            _ = fi48
            _ = v49
        }
        var v52 uint32 = 2147483647
        _ = v52
    }
    var v53 uint32 = 2147483648
    switch ("a" + "foo bar") {
    case "\n\t":
        v53 -= (v53 | v53)
    case "hello":
        v54 := GenS1{F0: 101}
        _ = v54
    default:
        switch (("foo bar" + "foo bar") + (("\n\t" + "a") + "a")) {
        case "hello":
            if true {
                for fi55 := 0; fi55 < 19; fi55++ {
                    _ = len(v4)
                    fmt.Println(fi55, v1)
                    var v56 int16 = -795
                    _ = fi55
                    _ = v56
                }
                var v57 int32 = -640
                if true {
                    var v58 string = ""
                    if len(v4) > 0 {
                        v4[0] = 128
                    }
                    var v59 uint8 = 127
                    _ = v58
                    _ = v59
                }
                _ = v57
            }
        case "foo bar":
            var v60 string = ("a" + "\n\t")
            _ = v60
        }
    }
    var v61 float32 = 1.5
    v62 := GenS5{F0: 152}
    v2 = 0
    if true {
        if len(v4) > 0 {
            v4[0] = 32767
        }
        for fi63 := 0; fi63 < 13; fi63++ {
            if true {
                v2 += v1
                v2 += 494
            } else {
                var v64 string = "foo bar"
                _ = v64
            }
            _ = fi63
        }
    } else {
        v65 := GenS2{F0: uint16(v3), F1: -1, F2: ""}
        _ = v65
    }
    fmt.Println(v3, v1)
    var v66 float32 = float32((0.0 * (1e10 + 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v61
    _ = v62
    _ = v66
}
