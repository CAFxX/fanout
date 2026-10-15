package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint }
type GenS7 struct { F0 uint32; F1 uint64; F2 string }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint8 }

func main() {
    var v1 bool = ((!true) || true)
    var v2 float64 = 1e10
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    v5 := []int{127}
    v6 := make(map[string]int)
    if (4294967295 <= 0) {
        switch 128 {
        case 32768:
            v3 -= (v3 ^ v3)
        default:
            fmt.Println(v3)
        }
        v7 := GenS1{F0: "foo bar"}
        _ = v7
    }
    var v8 int64 = 0
    var v9 uint8 = 39
    switch "" {
    case "":
        if (1.5 <= -2.25) {
            if (1 >= 42) {
                var v10 uint16 = v3
                switch (("a" + "\n\t") + "\n\t") {
                case "foo bar":
                    if ("hello" != "\n\t") {
                        var v11 uint32 = 17
                        if (((665 <= 18446744073709551615) || v1) || v1) {
                            v4[2147483647] = (^(v11))
                            for fi12 := 0; fi12 < 3; fi12++ {
                                fmt.Println(v8, v8)
                                var v13 float32 = (0.0 * float32(v2))
                                v11 -= (^(v11))
                                _ = fi12
                                _ = v13
                            }
                            if v1 {
                                switch "foo bar" {
                                case "\n\t":
                                    v3 = v3
                                default:
                                    var v14 uint32 = (v11 % ((42) | 1))
                                    _ = v14
                                }
                                if (128 >= 32767) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        }
                        _ = v11
                    } else {
                        v8 = -20
                        if v1 {
                            var v15 float64 = v2
                            _ = v15
                        } else {
                            if (632 != 82) {
                                if (82 < v9) {
                                    var v16 int64 = (v8 - (v8 / ((v8) | 1)))
                                    v3 = v10
                                    _ = v16
                                } else {
                                    var v17 bool = (107 == 88)
                                    v3 = v10
                                    _ = v17
                                }
                            } else {
                                v18 := GenS1{F0: ("" + ("\n\t" + ""))}
                                v8 += v8
                                if ((-796 < 4294967295) || (4294967295 >= 2147483647)) {
                                    if (!v1) {
                                        var v19 int = 2147483647
                                        for fi20 := 0; fi20 < 20; fi20++ {
                                            var v21 float64 = v2
                                            var v22 int32 = -422
                                            var v23 bool = true
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = v19
                                    } else {
                                        _ = len(v4)
                                        v24 := GenS2{F0: 1000}
                                        _ = v24
                                    }
                                    v9 = 127
                                } else {
                                    fmt.Println(v8, v3)
                                }
                                _ = v18
                            }
                            v2 = 0.0
                            delete(v6, ("" + ""))
                        }
                        v25 := GenS1{F0: ("hello" + "a")}
                        _ = v25
                    }
                case "hello":
                    delete(v4, -1)
                default:
                    if (v3 > 65535) {
                        switch "foo bar" {
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                v26 := GenS1{F0: ""}
                                _ = v26
                            }
                        case "":
                            v2 -= v2
                        }
                    }
                }
                _ = v10
            }
        }
    case "foo bar":
        var v27 uint32 = 2
        _ = v27
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            v6["hello"] = fi28
            var v29 int64 = v8
            _ = fi28
            _ = v29
        }
    }
    if len(v5) > 0 {
        v5[0] = -875
    }
    var v30 int64 = v8
    v8 = -41
    var v31 uint = 18446744073709551615
    if len(v5) > 0 {
        v5[0] = 426
    }
    switch "foo bar" {
    case "a":
        if v1 {
            for fi32 := 0; fi32 < 4; fi32++ {
                if (1.5 >= (1.5 / (1e10 + ((float32((0.0 + (-2.25 - ((1.5 / v2) * v2)))) * -2.25) - (1e10 * 1.5))))) {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v3 += 42
                        if v1 {
                            var v34 uint32 = 2147483647
                            _ = v34
                        }
                        if v1 {
                            v35 := v6[("\n\t" + ("" + ""))]
                            if false {
                                if (!(693 != -1)) {
                                    var v36 float32 = 0.1
                                    fmt.Println(v36)
                                    if len(v5) > 0 {
                                        v5[0] = -295
                                    }
                                    _ = v36
                                } else {
                                    if false {
                                        v37 := GenS3{F0: 1, F1: ("" + (("foo bar" + "") + "foo bar"))}
                                        _ = v37
                                    } else {
                                        var v38 int8 = 69
                                        switch (v35 - v35) {
                                        case 4294967295:
                                            var v39 int8 = (v38 >> ((v31) & 63))
                                            _ = v39
                                        case 1:
                                            v5 = append(v5, (^(v35)))
                                        default:
                                            v40 := GenS4{F0: 2147483647}
                                            _ = v40
                                        }
                                        if (0 > 594) {
                                            v38 = (v38 / ((42) | 1))
                                            _ = len(v5)
                                            _ = len(v4)
                                        }
                                        switch "" {
                                        case "a":
                                            v41 := GenS5{F0: -32768, F1: 0}
                                            _ = v41
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v42 := GenS4{F0: 1}
                                        if (v1 || false) {
                                            if true {
                                                v5 = append(v5, v35)
                                                var v43 uint32 = 1
                                                if len(v5) > 0 {
                                                    v5[0] = v35
                                                }
                                                var v44 int = -2147483648
                                                _ = v43
                                                _ = v44
                                            }
                                            v31 = (v31 << (((v31 % ((v31) | 1))) & 63))
                                        }
                                        switch (fi33 + (fi32 & (fi32 + v35))) {
                                        case 32768:
                                            switch v35 {
                                            case -119:
                                                fmt.Println(v38, v38)
                                            case 0:
                                                v9 -= 42
                                            }
                                        default:
                                            var v45 uint64 = 316
                                            _ = v45
                                        }
                                        _ = v38
                                        _ = v42
                                    }
                                }
                                var v46 int16 = 1000
                                fmt.Println(v30)
                                _ = len(v5)
                                var v47 uint64 = 4294967295
                                _ = v46
                                _ = v47
                            } else {
                                switch (^(fi33)) {
                                case 395:
                                    var v48 int8 = 100
                                    _ = v48
                                case 1:
                                    v30 += v30
                                case 723:
                                    if v1 {
                                        v49 := GenS6{F0: 2, F1: (v30 ^ v30), F2: 42}
                                        v50 := GenS6{F0: 0, F1: v8, F2: 855}
                                        switch v35 {
                                        case 128:
                                            fmt.Println(v9, v3, fi32)
                                        default:
                                            if v1 {
                                                var v51 uint16 = (v3 - v3)
                                                v51 = v51
                                                v52 := GenS3{F0: 0, F1: "\n\t"}
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                        var v53 int64 = (v30 * 648)
                                        _ = v49
                                        _ = v50
                                        _ = v53
                                    } else {
                                        if (!(1 >= 2147483647)) {
                                            var v54 int = fi32
                                            var v55 bool = true
                                            v2 = float64((1.5 / 1e10))
                                            if (v55 && v55) {
                                                switch (v54 >> (((^(v31))) & 63)) {
                                                case 32768:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            } else {
                                                var v56 bool = v1
                                                fmt.Println(v9, fi32, v31)
                                                _ = v56
                                            }
                                            v57 := GenS7{F0: 42, F1: 4294967295, F2: ("" + "\n\t")}
                                            _ = v54
                                            _ = v55
                                            _ = v57
                                        }
                                        fmt.Println(fi33, v31, v9)
                                        v58 := GenS8{F0: (v9 >> (((v31 ^ uint(v9))) & 63)), F1: 2, F2: v9}
                                        var v59 int64 = v8
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                v5 = append(v5, (fi32 << ((v31) & 63)))
                                var v60 bool = (!false)
                                _ = v60
                            }
                            v30 = 95
                            _ = v35
                        }
                        _ = fi33
                    }
                    if ((v9 << (((v31 % (((v31 / (((v31 - v31)) | 1))) | 1))) & 63)) <= (v9 % (((^(v9))) | 1))) {
                        v3 = 1
                    } else {
                        for fi61 := 0; fi61 < 19; fi61++ {
                            var v62 int32 = 42
                            _ = fi61
                            _ = v62
                        }
                        var v63 string = "a"
                        switch "foo bar" {
                        case "hello":
                            v30 += v30
                        case "a":
                            v64 := v4[fi32]
                            _ = v64
                        case "\n\t":
                            var v65 int64 = 0
                            _ = v65
                        default:
                            v30 -= (v30 >> ((v31) & 63))
                        }
                        _ = v63
                    }
                }
                v31 = (^(v31))
                _ = fi32
            }
            var v66 uint32 = 4294967295
            _ = v66
        }
    default:
        v67 := GenS2{F0: -32768}
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
