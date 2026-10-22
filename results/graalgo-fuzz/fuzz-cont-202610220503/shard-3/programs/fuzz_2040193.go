package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 int16; F1 int64; F2 uint }
type GenS6 struct { F0 uint32; F1 bool; F2 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 string = ("hello" + v1)
    var v3 int8 = -128
    var v4 bool = (false && false)
    var v5 uint64 = 4294967295
    v5 = v5
    var v6 uint32 = 203
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(v4, v5, v4)
        _ = fi7
    }
    var v8 uint64 = 18446744073709551615
    var v9 uint32 = 294
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = v1
        v12 := GenS1{F0: (v3 % (((v3 / (((v3 % (((v3 + 64)) | 1))) | 1))) | 1)), F1: v3}
        var v13 int16 = 1
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v6 += (v6 & v9)
    switch 4294967295 {
    case 1:
        if (false || (((("a" + (v2 + v2)) + ("" + "foo bar")) + v2) < v1)) {
            if true {
                var v14 uint16 = 65535
                _ = v14
            }
            if true {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := []int{(fi15 & fi15), int(v9), fi15, fi15}
                    for fi17 := 0; fi17 < 7; fi17++ {
                        if v4 {
                            v5 += v8
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        } else {
                            var v19 uint16 = 32767
                            switch (fi15 / ((fi17) | 1)) {
                            case -1:
                                v9 = v6
                            case -2147483648:
                                v20 := GenS1{F0: -69, F1: (v3 * (v3 | -98))}
                                _ = v20
                            case -576:
                                if (false && true) {
                                    v9 = uint32(v19)
                                    switch "hello" {
                                    case "":
                                        v3 = v3
                                    }
                                    if true {
                                        switch v1 {
                                        case "a":
                                            v21 := GenS1{F0: (v3 >> ((2) & 63)), F1: (v3 % ((127) | 1))}
                                            _ = v21
                                        case "\n\t":
                                            var v22 int64 = 130
                                            _ = v22
                                        default:
                                            var v23 bool = v4
                                            _ = v23
                                        }
                                        fmt.Println(fi15, v8, v2)
                                        fmt.Println(v6)
                                    } else {
                                        if ((v6 % (((v6 << ((18446744073709551615) & 63))) | 1)) == (^(v9))) {
                                            var v24 uint64 = 256
                                            _ = v24
                                        }
                                        if true {
                                            v25 := GenS2{F0: 1, F1: (v3 & (v3 * -128))}
                                            v26 := GenS3{F0: int(v19), F1: v5}
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 int8 = v3
                                            _ = v27
                                        }
                                        if len(v16) > 0 {
                                            v16[0] = (fi15 & 32767)
                                        }
                                    }
                                }
                                if v4 {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    var v28 int64 = 1
                                    var v29 uint16 = (v19 + 65535)
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch 127 {
                                    case 375:
                                        var v30 float64 = 1e10
                                        _ = v30
                                    case -9223372036854775808:
                                        v31 := GenS1{F0: 1, F1: v3}
                                        _ = v31
                                    default:
                                        v32 := GenS1{F0: (v3 - 42), F1: v3}
                                        _ = v32
                                    }
                                }
                            default:
                                v16 = append(v16, 4294967295)
                            }
                            _ = v19
                        }
                        v4 = v4
                        var v33 float64 = 0.1
                        _ = fi17
                        _ = v33
                    }
                    switch fi15 {
                    case 286:
                        var v34 uint8 = 0
                        _ = v34
                    default:
                        var v35 uint = 42
                        _ = v35
                    }
                    _ = fi15
                    _ = v16
                }
                v36 := []int{42, -852, 1}
                var v37 float64 = 0.0
                _ = v36
                _ = v37
            } else {
                switch v2 {
                case "hello":
                    var v38 int32 = 42
                    _ = v38
                }
                if v4 {
                    var v39 float64 = 0.1
                    v40 := []int{128}
                    _ = v39
                    _ = v40
                } else {
                    var v41 bool = v4
                    switch 376 {
                    case 508:
                        v42 := GenS4{F0: 32768, F1: "\n\t", F2: 32767}
                        _ = v42
                    case 1:
                        v2 = v1
                    case 32767:
                        var v43 uint8 = 2
                        _ = v43
                    default:
                        if ((!v41) && (1 == 2147483647)) {
                            var v44 uint32 = 0
                            var v45 uint32 = v9
                            if false {
                                if v4 {
                                    if v4 {
                                        v1 = (("hello" + v1) + v2)
                                        v5 += 2
                                    }
                                } else {
                                    if false {
                                        var v46 int32 = -162
                                        _ = v46
                                    }
                                }
                                var v47 uint8 = 1
                                if v4 {
                                    var v48 string = ""
                                    switch "a" {
                                    case "":
                                        var v49 uint16 = 255
                                        _ = v49
                                    default:
                                        v50 := []int{int(v5)}
                                        _ = v50
                                    }
                                    _ = v48
                                } else {
                                    v41 = ((4294967295 > 0) || (!(!true)))
                                    fmt.Println(v2, v3)
                                    fmt.Println(v41)
                                }
                                _ = v47
                            } else {
                                v51 := GenS5{F0: -32768, F1: -50, F2: 255}
                                v52 := GenS3{F0: 145, F1: v5}
                                var v53 bool = (false || (1 == (v3 | v3)))
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            _ = v44
                            _ = v45
                        }
                    }
                    var v54 uint8 = 13
                    _ = v41
                    _ = v54
                }
            }
            if false {
                fmt.Println(v1, v3)
                v55 := GenS6{F0: v6, F1: v4, F2: 42}
                v56 := GenS1{F0: (v3 | (^(v3))), F1: -128}
                _ = v55
                _ = v56
            } else {
                fmt.Println(v5, v6)
                v3 = (^(v3))
            }
        }
    case 361:
        v57 := []int{-9223372036854775808, -1, 4294967295}
        _ = v57
    default:
        var v58 uint32 = (v6 ^ v9)
        _ = v58
    }
    fmt.Println(v2, v3, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
