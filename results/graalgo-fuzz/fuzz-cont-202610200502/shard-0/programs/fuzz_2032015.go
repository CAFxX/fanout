package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{-9223372036854775808, fi8, 2147483647}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{(-(fi10)), fi10}
        _ = fi10
        _ = v11
    }
    _ = v7
    return 254
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int64 = 1
    v4 := make(map[string]string)
    var v5 uint8 = 139
    var v6 uint64 = 0
    fmt.Println(v5, v6)
    v12 := []int{32768, -1}
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 -= (v5 / (((v5 | v5)) | 1))
        for fi14 := 0; fi14 < 5; fi14++ {
            switch "" {
            case "":
                fmt.Println(v6, v5, v3)
            case "\n\t":
                var v15 int64 = (v3 >> ((42) & 63))
                _ = v15
            case "foo bar":
                v16 := GenS1{F0: 0, F1: uint(v6), F2: (v6 / (((^(v6))) | 1))}
                _ = v16
            }
            if (v6 > v6) {
                var v17 uint8 = 1
                _ = v17
            } else {
                v5 += (v5 | (v5 | uint8(v6)))
                if (0 > -1) {
                    switch 32768 {
                    case 9223372036854775807:
                        v5 -= 255
                    case 42:
                        var v18 uint16 = 691
                        _ = v18
                    case 881:
                        fmt.Println(fi13, fi14, v3)
                    default:
                        _ = len(v12)
                    }
                } else {
                    v3 -= v3
                    if false {
                        if false {
                            v3 -= (v3 << ((31) & 63))
                            v5 += 127
                        }
                        fmt.Println(fi13, fi13)
                    } else {
                        if ((v6 / (((v6 & 42)) | 1)) >= 18446744073709551615) {
                            var v19 int64 = (v3 - 287)
                            switch ("foo bar" + "a") {
                            case "a":
                                v3 = v3
                            default:
                                v6 = (v6 - 62)
                            }
                            _ = len(v2)
                            v6 += v6
                            v5 += 137
                            _ = v19
                        } else {
                            fmt.Println(v6, v3, fi14)
                            switch fi13 {
                            case 0:
                                var v20 float64 = 1.5
                                _ = v20
                            default:
                                switch ("a" + "\n\t") {
                                case "a":
                                    v6 -= (v6 << ((18446744073709551615) & 63))
                                default:
                                    var v21 bool = true
                                    _ = v21
                                }
                            }
                            v3 -= (v3 % (((v3 >> ((94) & 63))) | 1))
                        }
                        var v22 uint32 = 197
                        v23 := GenS1{F0: -2147483648, F1: 459, F2: (^(v6))}
                        _ = v22
                        _ = v23
                    }
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    v3 += (v3 - v3)
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint32 = 0
        _ = fi24
        _ = v25
    }
    var v26 uint = 18446744073709551615
    if true {
        var v27 uint16 = 255
        var v28 string = "a"
        var v29 int16 = 32767
        _ = v27
        _ = v28
        _ = v29
    } else {
        v26 -= (v26 + v26)
    }
    if (0.0 < 0.0) {
        v3 = v3
        var v30 int = -1
        _ = v30
    } else {
        var v31 uint16 = 0
        _ = v31
    }
    fmt.Println(v6, v26)
    var v32 int32 = -2147483648
    for fi33 := 0; fi33 < 9; fi33++ {
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 string = ""
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    v5 += (v5 ^ 255)
    if true {
        var v36 int = -1
        _ = v36
    } else {
        var v37 int = 42
        v38 := GenS2{F0: v3, F1: (v32 | (v32 - v32)), F2: v32}
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        fmt.Println(fi39, v3)
        if (uint16(v26) != 32767) {
            if len(v12) > 0 {
                v12[0] = (fi39 << ((0) & 63))
            }
        } else {
            var v40 int8 = 127
            _ = v40
        }
        v41 := GenS3{F0: v32}
        _ = fi39
        _ = v41
    }
    if true {
        var v42 string = "foo bar"
        _ = v42
    } else {
        if true {
            v26 += 4294967295
        }
        if len(v12) > 0 {
            v12[0] = 1
        }
    }
    delete(v4, "\n\t")
    var v43 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
