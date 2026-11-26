package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 float32 = 1e10
    v2 := []string{}
    var v3 int = 1
    var v4 int64 = -9223372036854775808
    v5 := []string{"", ((("hello" + "\n\t") + ((("" + "foo bar") + "foo bar") + "foo bar")) + "foo bar")}
    v6 := []int{(v3 & (v3 + (^(v3))))}
    switch "hello" {
    case "hello":
        var v7 string = ""
        _ = v7
    case "\n\t":
        fmt.Println(v3)
    case "":
        if len(v6) > 0 {
            v6[0] = (^(v3))
        }
    default:
        v5 = append(v5, ("foo bar" + ""))
    }
    fmt.Println(v3)
    var v8 uint = 1
    v4 = v4
    fmt.Println(v3)
    var v9 uint8 = 2
    for fi10 := 0; fi10 < 3; fi10++ {
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        var v11 int8 = 36
        v12 := GenS1{F0: (v8 | (v8 / ((uint(v9)) | 1)))}
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: v8}
        v3 = (fi13 << (((v8 << ((65535) & 63))) & 63))
        v3 -= (v3 >> (((v8 ^ v8)) & 63))
        _ = fi13
        _ = v14
    }
    if (32767 == 42) {
        v1 = v1
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int = (v3 + 32768)
            var v17 int16 = -1
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        switch v3 {
        case -9223372036854775808:
            v1 -= float32((1.5 / (((-2.25 / (((1.5 + 3.14159) * 1.5) - 0.0)) / -2.25) - 0.1)))
        case 1:
            v2 = append(v2, "a")
        }
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            if (int16(v4) < -1) {
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 uint32 = 0
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (!true) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v8, v1)
                v20 := GenS2{F0: -1}
                _ = v20
            }
            var v21 int32 = 42
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := GenS2{F0: (v4 ^ (v4 << (((v8 * v8)) & 63)))}
                _ = fi22
                _ = v23
            }
            _ = v21
        }
    }
    var v24 string = "hello"
    if ((3.14159 >= 1e10) && (!true)) {
        fmt.Println(v3)
        v8 = (v8 / (((^(v8))) | 1))
        v5 = append(v5, "")
    }
    if false {
        if len(v2) > 0 {
            v2[0] = (v24 + "")
        }
        fmt.Println(v4, v4, v4)
        var v25 uint32 = 483
        _ = v25
    }
    v9 = (v9 % ((v9) | 1))
    var v26 uint32 = 0
    if (uint64(v26) < 0) {
        for fi27 := 0; fi27 < 10; fi27++ {
            if (v26 < v26) {
                var v28 int8 = 93
                _ = v28
            } else {
                v29 := GenS2{F0: v4}
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := GenS3{F0: 2, F1: (v24 + v24)}
                    if true {
                        switch -9223372036854775808 {
                        case -1:
                            v32 := GenS3{F0: uint32(v9), F1: v24}
                            _ = v32
                        }
                        switch fi27 {
                        case 2147483647:
                            v26 = v26
                        case -2147483648:
                            switch v24 {
                            case "foo bar":
                                if len(v6) > 0 {
                                    v6[0] = int(v8)
                                }
                            case "a":
                                if ((v3 | (fi27 + fi30)) != 127) {
                                    v33 := GenS4{F0: 100, F1: 0}
                                    _ = v33
                                } else {
                                    v1 -= (((v1 - 1e10) / ((v1 - v1) * (((v1 + (v1 / v1)) - (1.5 * v1)) / v1))) - 1e10)
                                }
                            case "\n\t":
                                v9 = v9
                            default:
                                if (62 != 42) {
                                    var v34 int8 = -63
                                    v1 = v1
                                    fmt.Println(fi30, v24, v26)
                                    v1 -= (3.14159 + 0.0)
                                    _ = v34
                                }
                            }
                        case 716:
                            v35 := GenS5{F0: 42}
                            _ = v35
                        default:
                            fmt.Println(fi30, v24)
                        }
                    } else {
                        var v36 bool = (false || (0 >= 344))
                        v24 = (v24 + v24)
                        v3 = 42
                        _ = v36
                    }
                    v1 += float32(1.5)
                    _ = fi30
                    _ = v31
                }
                v8 -= (v8 << (((v8 ^ (v8 * (v8 << ((uint(v9)) & 63))))) & 63))
                _ = v29
            }
            _ = fi27
        }
    } else {
        var v37 uint = (v8 - v8)
        if len(v5) > 0 {
            v5[0] = "a"
        }
        if true {
            if len(v2) > 0 {
                v2[0] = v24
            }
        } else {
            var v38 int = v3
            for fi39 := 0; fi39 < 12; fi39++ {
                if true {
                    if false {
                        var v40 int16 = 466
                        v1 = float32((float64((((float32(1e10) - v1) + float32(-2.25)) + 1.5)) + (0.1 + 1e10)))
                        _ = v40
                    }
                    var v41 int = (v38 - -34)
                    _ = v41
                }
                var v42 string = (v24 + v24)
                v6 = append(v6, 628)
                _ = fi39
                _ = v42
            }
            v4 -= 16
            _ = v38
        }
        _ = v37
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        var v44 int64 = (v4 >> (((^(v8))) & 63))
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
