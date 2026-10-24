package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0(p0 uint64, p1 uint) int8 {
    v4 := GenS1{F0: 32768, F1: -901, F2: 42}
    v5 := GenS2{F0: 42, F1: 127}
    v6 := GenS3{F0: 2147483647, F1: 4294967295}
    p1 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint64 = 2
    v2 := []uint32{2147483647, 2147483647, 290, 1}
    v3 := make(map[uint32]bool)
    v2 = append(v2, 4294967295)
    switch "\n\t" {
    case "a":
        if (42 < 1) {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS2{F0: -48, F1: 127}
                _ = fi7
                _ = v8
            }
            var v9 uint64 = v1
            var v10 int32 = -1
            _ = v10
            _ = v9
        }
    case "\n\t":
        v1 += v1
    default:
        var v11 int = 0
        _ = v11
    }
    fmt.Println(v1)
    v1 -= 974
    switch "hello" {
    case "\n\t":
        if false {
            v1 = v1
            var v12 uint64 = (v1 - (^(v1)))
            _ = v12
        } else {
            v1 -= 18446744073709551615
            fmt.Println(v1)
        }
    }
    v13 := GenS4{F0: -565, F1: "\n\t"}
    for fi14 := 0; fi14 < 11; fi14++ {
        if true {
            switch 127 {
            case 42:
                switch fi14 {
                case 2147483647:
                    v1 += 1
                case 215:
                    v1 -= (v1 * (v1 + (v1 * 18446744073709551615)))
                default:
                    fmt.Println(v1)
                }
            case -9223372036854775808:
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(v1, v1, v1)
                    _ = fi15
                }
            default:
                fmt.Println(v1, fi14)
            }
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 float32 = (3.14159 / -2.25)
            fmt.Println(fi16, v17, v17)
            if false {
                fmt.Println(fi16, fi16)
                v18 := GenS2{F0: int8(v1), F1: -1}
                _ = v18
            } else {
                fmt.Println(fi16, v1)
                v3[4294967295] = false
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    if true {
        fmt.Println(v1)
        if ((((v1 << ((18446744073709551615) & 63)) <= v1) || (100 != 69)) && (false && true)) {
            switch "hello" {
            case "a":
                var v19 uint16 = 255
                _ = v19
            case "foo bar":
                if true {
                    if false {
                        var v20 string = ("hello" + ("a" + "a"))
                        v20 = v20
                        var v21 float32 = 0.0
                        _ = v20
                        _ = v21
                    }
                    v1 += (v1 % (((v1 + (v1 + v1))) | 1))
                    fmt.Println(v1)
                }
            case "hello":
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    fmt.Println(v1)
                    if (!(65535 <= 255)) {
                        v22 := GenS3{F0: 427, F1: -190}
                        var v23 uint16 = 2
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        _ = v22
                        _ = v23
                    } else {
                        v1 = v1
                        for fi24 := 0; fi24 < 8; fi24++ {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                var v26 uint16 = 1
                                v27 := GenS5{F0: fi25, F1: -32768}
                                var v28 int32 = 0
                                _ = fi25
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            if (-2147483648 != 42) {
                                v1 -= v1
                                v1 = v1
                                var v29 uint16 = 488
                                _ = v29
                            }
                            var v30 int32 = 42
                            _ = fi24
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v1 -= (v1 * (^(v1)))
                            v32 := GenS6{F0: (-85 <= 100), F1: 544, F2: 730}
                            _ = fi31
                            _ = v32
                        }
                    }
                    v1 += (^(v1))
                }
            }
            if (!false) {
                for fi33 := 0; fi33 < 6; fi33++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi33
                }
                if (935 == 0) {
                    fmt.Println(v1)
                    v1 -= (v1 / (((v1 * v1)) | 1))
                }
            }
            _ = len(v3)
        }
    } else {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 int64 = 0
            fmt.Println(fi34)
            _ = fi34
            _ = v35
        }
        var v36 bool = (false || false)
        _ = v36
    }
    v37 := GenS1{F0: 32767, F1: 42, F2: 1000}
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 uint32 = 0
        _ = fi38
        _ = v39
    }
    v1 -= (v1 + (v1 | v1))
    v40 := GenS4{F0: -2147483648, F1: "foo bar"}
    var v41 int32 = 42
    if (-561 != 0) {
        v1 = (^(v1))
        v41 = -1
    }
    v1 -= v1
    var v42 int = 128
    v43 := GenS5{F0: (v42 | (v42 + (v42 / ((v42) | 1)))), F1: -899}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
