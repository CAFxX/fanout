package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 string; F1 uint8; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int8; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint32; F1 int32; F2 int8 }

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 uint = 1
    var v3 uint = 18446744073709551615
    v4 := GenS1{F0: 4294967295}
    var v5 int8 = 0
    v5 -= (-(v5))
    fmt.Println(v3)
    v1 = append(v1, "")
    v6 := GenS1{F0: 714}
    var v7 uint16 = 42
    var v8 uint = v3
    v9 := GenS2{F0: 2147483647, F1: "foo bar", F2: 4294967295}
    if (true && false) {
        var v10 int16 = 1
        var v11 float32 = 0.1
        _ = v10
        _ = v11
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v12 float64 = 0.0
            _ = v12
        case 32768:
            for fi13 := 0; fi13 < 18; fi13++ {
                if false {
                    v14 := GenS3{F0: "hello", F1: 127, F2: (2 <= uint64(v7))}
                    v3 -= v3
                    var v15 bool = false
                    _ = v14
                    _ = v15
                } else {
                    if (1.5 < 3.14159) {
                        var v16 string = ""
                        var v17 float32 = 0.0
                        v18 := GenS1{F0: 4294967295}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        var v19 int32 = 2147483647
                        for fi20 := 0; fi20 < 11; fi20++ {
                            _ = len(v1)
                            v19 += (v19 * v19)
                            switch ("hello" + ("foo bar" + "")) {
                            case "":
                                v21 := GenS4{F0: 2147483647}
                                _ = v21
                            case "foo bar":
                                v22 := GenS5{F0: (v3 << ((v8) & 63)), F1: 91}
                                _ = v22
                            case "a":
                                v7 = (v7 % (((^(v7))) | 1))
                            default:
                                switch fi13 {
                                case 9223372036854775807:
                                    var v23 float32 = 3.14159
                                    v24 := GenS6{F0: -1, F1: 2}
                                    switch "" {
                                    case "":
                                        var v25 int64 = 805
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (("hello" + "a") + "")
                                        }
                                        _ = v25
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                    }
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = fi20
                        }
                        v26 := GenS1{F0: 4294967295}
                        _ = v19
                        _ = v26
                    }
                }
                _ = fi13
            }
        case 128:
            _ = len(v1)
        default:
            if false {
                if (2 != 2) {
                    v8 = v3
                    v7 += v7
                    for fi27 := 0; fi27 < 19; fi27++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi28 := 0; fi28 < 20; fi28++ {
                            if len(v1) > 0 {
                                v1[0] = (("\n\t" + "\n\t") + "a")
                            }
                            _ = fi28
                        }
                        v29 := GenS4{F0: 1}
                        _ = fi27
                        _ = v29
                    }
                }
                var v30 uint16 = v7
                var v31 int32 = 562
                _ = v30
                _ = v31
            }
        }
        fmt.Println(v8, v7)
    }
    var v32 string = "\n\t"
    fmt.Println(v32)
    v8 -= 4294967295
    if false {
        switch 32768 {
        case 122:
            for fi33 := 0; fi33 < 7; fi33++ {
                for fi34 := 0; fi34 < 7; fi34++ {
                    fmt.Println(v8)
                    v35 := GenS7{F0: (v5 | (v5 % ((0) | 1))), F1: v5, F2: 255}
                    if (false || ((int8(fi33) > v5) || false)) {
                        if (v2 == (v8 ^ 669)) {
                            fmt.Println(fi34, v5, v5)
                            fmt.Println(fi34)
                            var v36 int64 = 9223372036854775807
                            _ = v36
                        }
                        if false {
                            v37 := GenS6{F0: 1000, F1: (^(v7))}
                            var v38 uint32 = uint32(v7)
                            v39 := GenS2{F0: -744, F1: "hello", F2: (fi33 / ((int(v7)) | 1))}
                            v1 = append(v1, "a")
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    } else {
                        v40 := GenS6{F0: 42, F1: v7}
                        switch fi33 {
                        case 118:
                            if (4294967295 > 1) {
                                switch (-(fi33)) {
                                case -9223372036854775808:
                                    if false {
                                        var v41 uint16 = (v7 * (v7 % ((uint16(v3)) | 1)))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v41
                                    } else {
                                        v42 := GenS1{F0: 0}
                                        var v43 uint16 = (v7 ^ v7)
                                        v2 += 4294967295
                                        _ = v42
                                        _ = v43
                                    }
                                case 42:
                                    v44 := GenS5{F0: v8, F1: 1}
                                    _ = v44
                                default:
                                    v1 = append(v1, v32)
                                }
                                if (2 < v7) {
                                    v1 = append(v1, "")
                                } else {
                                    var v45 int32 = -2147483648
                                    fmt.Println(v7, v32)
                                    v3 = v3
                                    _ = v45
                                }
                                v46 := GenS5{F0: 613, F1: 128}
                                _ = v46
                            } else {
                                var v47 uint = v3
                                if (!(!true)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v48 bool = ((v5 | (v5 | (v5 * (v5 / (((v5 & (v5 * 127))) | 1))))) > v5)
                                    v8 = 2
                                    _ = v48
                                }
                                v2 += (v2 % ((v2) | 1))
                                _ = v47
                            }
                        case -9223372036854775808:
                            var v49 uint8 = 74
                            _ = v49
                        case 128:
                            v1 = append(v1, v32)
                        default:
                            v50 := GenS1{F0: 366}
                            if false {
                                switch v32 {
                                case "a":
                                    v32 = v32
                                case "":
                                    v32 = v32
                                case "foo bar":
                                    fmt.Println(v3)
                                }
                                fmt.Println(v32)
                            } else {
                                if true {
                                    v7 -= v7
                                } else {
                                    fmt.Println(v5, v7)
                                    switch ("\n\t" + (v32 + v32)) {
                                    case "a":
                                        if (((v32 + (((((v32 + "\n\t") + v32) + (("\n\t" + (v32 + "\n\t")) + v32)) + v32) + "foo bar")) + "a") >= "") {
                                            v51 := GenS4{F0: 42}
                                            var v52 float32 = 1e10
                                            _ = v51
                                            _ = v52
                                        } else {
                                            var v53 uint16 = uint16(v2)
                                            _ = len(v1)
                                            _ = v53
                                        }
                                    case "hello":
                                        _ = len(v1)
                                    }
                                    fmt.Println(v3)
                                }
                            }
                            _ = v50
                        }
                        fmt.Println(fi34, v3, v8)
                        _ = v40
                    }
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        case -2147483648:
            v54 := GenS8{F0: 584, F1: -472, F2: v5}
            _ = v54
        case -885:
            v55 := GenS1{F0: 0}
            _ = v55
        }
        if true {
            var v56 int32 = 1
            _ = v56
        }
    } else {
        v1 = append(v1, ((v32 + "foo bar") + v32))
        if true {
            v57 := GenS6{F0: -32768, F1: (v7 >> ((v3) & 63))}
            v3 -= (v3 & v3)
            for fi58 := 0; fi58 < 16; fi58++ {
                fmt.Println(v32, v5, v5)
                for fi59 := 0; fi59 < 2; fi59++ {
                    switch (("\n\t" + (v32 + "\n\t")) + (v32 + v32)) {
                    case "\n\t":
                        switch (v32 + "foo bar") {
                        case "":
                            v8 += v3
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    _ = fi59
                }
                _ = fi58
            }
            _ = v57
        }
        var v60 float64 = 1.5
        _ = v60
    }
    v1 = append(v1, ("hello" + "a"))
    v7 -= v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
