package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0() uint64 {
    if true {
        var v6 uint16 = 65535
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint8 = 150
        var v9 int16 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float32 = 0.0
    for fi11 := 0; fi11 < 7; fi11++ {
        v10 -= 3.14159
        switch "" {
        case "":
            v12 := []int{32768}
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 12; fi13++ {
                switch fi13 {
                case -43:
                    v14 := []int{fi13}
                    _ = v14
                case 456:
                    if (!(v10 >= v10)) {
                        var v15 uint8 = 42
                        var v16 uint64 = 4294967295
                        _ = v15
                        _ = v16
                    }
                case 42:
                    v17 := GenS1{F0: 4294967295, F1: (!false)}
                    var v18 uint64 = 4294967295
                    fmt.Println(v18, fi13)
                    _ = v17
                    _ = v18
                }
                if false {
                    var v19 float64 = -2.25
                    _ = v19
                }
                var v20 float64 = ((1.5 * float64(3.14159)) - 3.14159)
                _ = fi13
                _ = v20
            }
        case "foo bar":
            v10 = v10
        default:
            fmt.Println(v10, fi11)
        }
        v21 := GenS2{F0: -9223372036854775808, F1: (fi11 + (fi11 >> ((65535) & 63))), F2: 255}
        _ = fi11
        _ = v21
    }
    _ = v10
    return 511
}
func genFn1() int8 {
    var v22 string = ("\n\t" + "a")
    v23 := GenS2{F0: -1, F1: 424, F2: 2}
    if (!true) {
        var v24 int8 = -128
        v25 := []int{1, 2147483647, 32767}
        _ = v24
        _ = v25
    }
    _ = v22
    _ = v23
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 103
    v3 := make(map[string]string)
    v4 := []uint32{2147483647}
    var v5 uint16 = 32767
    v5 = 1
    var v26 uint64 = 18446744073709551615
    if false {
        if ("a" < (("\n\t" + (("" + "foo bar") + "\n\t")) + "\n\t")) {
            var v27 uint = 255
            _ = v27
        }
        v28 := v1["foo bar"]
        var v29 int = 9223372036854775807
        _ = v28
        _ = v29
    } else {
        var v30 int32 = -1
        _ = v30
    }
    var v31 float32 = 0.0
    if true {
        _ = len(v4)
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 int8 = 0
            _ = fi32
            _ = v33
        }
        var v34 int32 = -50
        _ = v34
    } else {
        v35 := v1["\n\t"]
        var v36 int32 = 798
        _ = v35
        _ = v36
    }
    var v37 string = ("foo bar" + "\n\t")
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 int = (fi38 * int(v5))
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 int64 = 789
            switch (^(fi41)) {
            case 4294967295:
                if true {
                    if true {
                        v43 := GenS1{F0: 0, F1: true}
                        _ = v43
                    } else {
                        v5 = (v5 >> ((1) & 63))
                        v42 -= v42
                        var v44 int8 = -1
                        _ = v44
                    }
                    var v45 string = v37
                    _ = v45
                }
            case 1:
                switch ((v37 + ("" + ("" + v37))) + (v37 + "")) {
                case "\n\t":
                    v31 = float32(1.5)
                case "foo bar":
                    v31 += float32(3.14159)
                default:
                    switch -9223372036854775808 {
                    case 594:
                        if true {
                            var v46 float32 = v31
                            v46 -= ((-2.25 * v31) * v46)
                            _ = v46
                        }
                    case 1:
                        var v47 uint8 = 42
                        _ = v47
                    default:
                        fmt.Println(v26, v5)
                    }
                }
            case -9223372036854775808:
                var v48 float64 = 0.0
                v1[v37] = fi41
                v49 := GenS1{F0: 0, F1: (!((!((!true) && ((v31 + (v31 * ((-2.25 * v31) / v31))) == v31))) && false))}
                var v50 int64 = v42
                v37 = v37
                _ = v48
                _ = v49
                _ = v50
            }
            _ = fi41
            _ = v42
        }
        var v51 int16 = 327
        _ = fi40
        _ = v51
    }
    switch v37 {
    case "\n\t":
        if ((("\n\t" + "") + ("foo bar" + "\n\t")) >= ((((v37 + (((v37 + "hello") + v37) + v37)) + "a") + (("" + "foo bar") + "")) + v37)) {
            fmt.Println(v31, v37)
        }
    }
    switch ((v37 + v37) + ("" + v37)) {
    case "hello":
        if (!true) {
            v26 = (v26 % ((525) | 1))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (-1 > int32(v2)) {
                switch "\n\t" {
                case "a":
                    v3[("" + v37)] = v37
                }
                switch 2147483647 {
                case 0:
                    if false {
                        if true {
                            if (true && false) {
                                if false {
                                    var v52 uint = 4294967295
                                    _ = v52
                                } else {
                                    if (!true) {
                                        if (!(1e10 != 0.0)) {
                                            var v53 uint32 = 65535
                                            v53 += (v53 ^ 42)
                                            _ = v53
                                        } else {
                                            switch ((v37 + ("" + "\n\t")) + ("a" + v37)) {
                                            case "\n\t":
                                                _ = len(v3)
                                            case "foo bar":
                                                var v54 float32 = v31
                                                _ = v54
                                            case "hello":
                                                fmt.Println(v37, v37)
                                            }
                                            var v55 int8 = genFn1()
                                            switch "a" {
                                            case "a":
                                                v37 = v37
                                            }
                                            _ = v55
                                        }
                                    } else {
                                        if (!false) {
                                            v56 := GenS3{F0: 916, F1: 865}
                                            _ = v56
                                        }
                                    }
                                    if (65535 == 65535) {
                                        v57 := GenS2{F0: 959, F1: 1, F2: 206}
                                        if (false && false) {
                                            for fi58 := 0; fi58 < 15; fi58++ {
                                                for fi59 := 0; fi59 < 4; fi59++ {
                                                    v60 := GenS4{F0: 4294967295, F1: -32768}
                                                    fmt.Println(fi59)
                                                    if (!false) {
                                                        _ = len(v4)
                                                        var v61 uint8 = 73
                                                        fmt.Println(fi59, v26, v26)
                                                        var v62 string = "hello"
                                                        var v63 uint = 4294967295
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = fi59
                                                    _ = v60
                                                }
                                                for fi64 := 0; fi64 < 3; fi64++ {
                                                    var v65 uint = 18446744073709551615
                                                    if (42 > -1) {
                                                        if ((v37 + (v37 + ("hello" + ((v37 + v37) + ((v37 + ("hello" + "hello")) + ("foo bar" + v37)))))) < v37) {
                                                            v31 = (1.5 * 3.14159)
                                                        }
                                                    } else {
                                                        var v66 uint32 = 1
                                                        v67 := GenS3{F0: (^(v66)), F1: (fi64 % ((fi58) | 1))}
                                                        switch 42 {
                                                        case -2147483648:
                                                            var v68 int16 = 42
                                                            _ = v68
                                                        case 128:
                                                            fmt.Println(fi64, fi58, v26)
                                                        }
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    _ = fi64
                                                    _ = v65
                                                }
                                                _ = fi58
                                            }
                                            for fi69 := 0; fi69 < 2; fi69++ {
                                                fmt.Println(v31)
                                                _ = fi69
                                            }
                                        }
                                        _ = v57
                                    }
                                }
                                v37 = v37
                            }
                            var v70 string = v37
                            var v71 uint16 = 255
                            _ = v70
                            _ = v71
                        }
                        v72 := GenS3{F0: 2, F1: -626}
                        _ = v72
                    }
                case 101:
                    if (v37 != (("foo bar" + v37) + "")) {
                        v37 = ("" + ((v37 + v37) + v37))
                    }
                default:
                    v31 += v31
                }
                v73 := GenS3{F0: 0, F1: 2147483647}
                _ = v73
            } else {
                if false {
                    for fi74 := 0; fi74 < 4; fi74++ {
                        var v75 int32 = 42
                        _ = fi74
                        _ = v75
                    }
                }
                for fi76 := 0; fi76 < 6; fi76++ {
                    if (2 < 2) {
                        v2 -= (v2 & (v2 | (^(v2))))
                    }
                    v26 += 57
                    _ = fi76
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
