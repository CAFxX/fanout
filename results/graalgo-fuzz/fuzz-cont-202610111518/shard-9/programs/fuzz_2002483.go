package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int; F2 int16 }

func main() {
    var v1 uint16 = 42
    var v2 float32 = float32(0.0)
    var v3 uint16 = v1
    var v4 float64 = -2.25
    v5 := []string{"hello", ("foo bar" + "\n\t"), "\n\t"}
    var v6 string = "hello"
    fmt.Println(v2, v1, v1)
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v1, fi7, v3)
        for fi8 := 0; fi8 < 13; fi8++ {
            if (-57 <= -32768) {
                if true {
                    v9 := GenS1{F0: -1, F1: -9223372036854775808, F2: -286}
                    v3 += 0
                    if true {
                        if true {
                            switch fi8 {
                            case 42:
                                v10 := GenS1{F0: 308, F1: 0, F2: 1}
                                _ = v10
                            default:
                                switch "foo bar" {
                                case "a":
                                    v4 -= float64(((v2 / (float32(((v4 * (v4 + v4)) * float64(v2))) * v2)) - v2))
                                default:
                                    switch v6 {
                                    case "":
                                        fmt.Println(fi8, v2, v3)
                                    case "foo bar":
                                        if (-1 != -71) {
                                            switch "hello" {
                                            case "hello":
                                                v3 -= (v3 / (((v1 ^ v1)) | 1))
                                            case "":
                                                v11 := GenS1{F0: 796, F1: 9223372036854775807, F2: -456}
                                                _ = v11
                                            default:
                                                v12 := GenS1{F0: 1, F1: -622, F2: 0}
                                                _ = v12
                                            }
                                            var v13 uint64 = 4294967295
                                            v14 := GenS1{F0: -1, F1: 42, F2: 0}
                                            _ = v13
                                            _ = v14
                                        }
                                    case "hello":
                                        _ = len(v5)
                                    default:
                                        var v15 uint16 = v1
                                        _ = v15
                                    }
                                }
                            }
                        }
                        fmt.Println(v2)
                    } else {
                        v1 += (v1 % ((v1) | 1))
                        var v16 uint64 = 2
                        v17 := GenS1{F0: 1, F1: 42, F2: 2147483647}
                        _ = v16
                        _ = v17
                    }
                    _ = v9
                } else {
                    var v18 float64 = v4
                    fmt.Println(fi8, v18, v6)
                    _ = v18
                }
                v1 = v1
                if true {
                    switch fi8 {
                    case 127:
                        var v19 float32 = (v2 - float32((v4 / (0.0 - 1.5))))
                        _ = v19
                    case 626:
                        fmt.Println(v2, v6, fi7)
                    case 9223372036854775807:
                        if (true && false) {
                            if ((1 != 676) || true) {
                                v20 := GenS2{F0: -995, F1: v6}
                                _ = v20
                            }
                            var v21 int32 = 42
                            switch v6 {
                            case "foo bar":
                                switch ("a" + (v6 + "\n\t")) {
                                case "foo bar":
                                    if (fi7 < (fi8 & fi7)) {
                                        v2 = -2.25
                                        v1 += 32767
                                        v4 -= 1e10
                                    } else {
                                        switch (fi8 << ((42) & 63)) {
                                        case -126:
                                            var v22 int8 = 1
                                            _ = v22
                                        default:
                                            v6 = v6
                                        }
                                        var v23 float32 = 0.1
                                        _ = len(v5)
                                        _ = v23
                                    }
                                case "a":
                                    if false {
                                        var v24 float64 = (1e10 * v4)
                                        _ = v24
                                    } else {
                                        switch "a" {
                                        case "a":
                                            if (-32768 >= 32767) {
                                                v2 += v2
                                                if true {
                                                    v21 += v21
                                                    var v25 uint8 = 67
                                                    _ = v25
                                                }
                                                v26 := GenS2{F0: (v21 | v21), F1: ""}
                                                v2 = float32(v4)
                                                fmt.Println(v4, fi8, v2)
                                                fmt.Println(v6, v4, fi7)
                                                var v27 uint16 = v1
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v4 -= v4
                                            }
                                        case "hello":
                                            var v28 uint64 = 0
                                            _ = v28
                                        case "foo bar":
                                            var v29 int = (-(fi7))
                                            switch v6 {
                                            case "":
                                                if (!false) {
                                                    var v30 int8 = 1
                                                    v6 = v6
                                                    _ = v30
                                                }
                                            default:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            if true {
                                                v4 -= v4
                                            } else {
                                                v31 := GenS1{F0: (v21 << ((65535) & 63)), F1: 0, F2: v21}
                                                if true {
                                                    var v32 int8 = -128
                                                    var v33 int32 = -2147483648
                                                    switch (v29 % ((fi7) | 1)) {
                                                    case -1:
                                                        if (true || false) {
                                                            v34 := GenS3{F0: 65535, F1: 2147483647, F2: ""}
                                                            _ = len(v5)
                                                            var v35 int16 = int16(v1)
                                                            v35 = (v35 ^ -1)
                                                            v1 -= (v1 ^ v1)
                                                            _ = v34
                                                            _ = v35
                                                        } else {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    case 42:
                                                        var v36 int = (fi7 ^ fi8)
                                                        _ = v36
                                                    default:
                                                        v37 := GenS1{F0: v21, F1: 0, F2: v33}
                                                        _ = v37
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    var v38 uint64 = 4294967295
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v38
                                                }
                                                _ = v31
                                            }
                                            _ = v29
                                        }
                                        v21 += v21
                                        switch (v6 + "") {
                                        case "":
                                            var v39 int64 = -195
                                            _ = v39
                                        case "\n\t":
                                            var v40 uint32 = 0
                                            _ = v40
                                        case "a":
                                            v2 = v2
                                        }
                                    }
                                }
                            case "":
                                v1 -= (^(v3))
                            case "hello":
                                v1 = 255
                            }
                            _ = v21
                        } else {
                            fmt.Println(fi8, fi7)
                        }
                    default:
                        v5 = append(v5, v6)
                    }
                    v41 := GenS3{F0: 1, F1: 1, F2: ""}
                    _ = v41
                } else {
                    var v42 float64 = -2.25
                    _ = v42
                }
            } else {
                var v43 float32 = (v2 + (3.14159 / v2))
                var v44 bool = false
                if (v44 || (1 > 1)) {
                    switch -2147483648 {
                    case 128:
                        v1 = 4
                    case 586:
                        v5 = append(v5, ((v6 + v6) + v6))
                    }
                    fmt.Println(v6, v3, v3)
                    var v45 int16 = -32768
                    switch (fi8 << ((128) & 63)) {
                    case 128:
                        var v46 uint8 = 142
                        _ = v46
                    default:
                        fmt.Println(v3)
                        var v47 int8 = 100
                        fmt.Println(v1)
                        _ = v47
                    }
                    v4 -= 0.1
                    v3 = 42
                    fmt.Println(v6, fi8, v45)
                    v3 -= (^(v3))
                    _ = v45
                } else {
                    v48 := GenS4{F0: ((fi8 >> ((18446744073709551615) & 63)) >= (fi8 + fi7))}
                    var v49 int = (fi7 ^ (fi7 & fi7))
                    _ = v48
                    _ = v49
                }
                _ = v43
                _ = v44
            }
            if (!(false && true)) {
                var v50 uint = 629
                if ((443 >= -170) && true) {
                    switch fi7 {
                    case 127:
                        var v51 int8 = 81
                        _ = v51
                    case 2147483647:
                        v4 -= (-2.25 + v4)
                    }
                    v5 = append(v5, v6)
                } else {
                    v4 = (0.0 + v4)
                    var v52 uint = v50
                    if false {
                        v53 := GenS5{F0: 0}
                        fmt.Println(v4, fi8, v6)
                        v4 -= 3.14159
                        _ = v53
                    }
                    _ = v52
                }
                if true {
                    var v54 uint64 = 2
                    fmt.Println(v50, v4)
                    v55 := GenS6{F0: 929}
                    v56 := GenS3{F0: 4294967295, F1: 2, F2: v6}
                    _ = v54
                    _ = v55
                    _ = v56
                }
                _ = v50
            } else {
                v5 = append(v5, v6)
                v1 = v1
                fmt.Println(v3, v1)
            }
            _ = fi8
        }
        var v57 int16 = 0
        _ = fi7
        _ = v57
    }
    var v58 bool = true
    var v59 uint8 = 42
    v60 := GenS4{F0: (int32(v59) != -2147483648)}
    v61 := GenS5{F0: 863}
    fmt.Println(v58)
    if (42 == -1) {
        switch ("hello" + "foo bar") {
        case "":
            fmt.Println(v2, v59, v59)
        case "foo bar":
            v59 -= (v59 % ((uint8(v1)) | 1))
        }
    }
    for fi62 := 0; fi62 < 7; fi62++ {
        switch "a" {
        case "a":
            v58 = v58
        case "foo bar":
            v63 := GenS7{F0: 693, F1: (fi62 >> ((1) & 63)), F2: 0}
            _ = v63
        default:
            var v64 float64 = v4
            _ = v64
        }
        _ = fi62
    }
    v65 := GenS1{F0: 323, F1: 911, F2: 2147483647}
    var v66 int64 = 1
    v4 += 0.0
    v59 -= 2
    v66 = 92
    var v67 int = 502
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v65
    _ = v66
    _ = v67
}
