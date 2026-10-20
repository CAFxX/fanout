package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = 864
    v3 := []string{"a", "foo bar", "", "foo bar"}
    var v4 int = 313
    v5 := []uint32{645}
    var v6 int16 = 42
    v7 := GenS1{F0: (v4 * (v4 | (-(v4)))), F1: 255, F2: 42}
    var v8 string = "foo bar"
    v6 += 42
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(v1, fi9)
            for fi10 := 0; fi10 < 1; fi10++ {
                fmt.Println(v2, v1)
                switch fi9 {
                case -557:
                    fmt.Println(v8)
                default:
                    switch v4 {
                    case 2147483647:
                        v5 = append(v5, 327)
                    default:
                        switch v8 {
                        case "hello":
                            var v11 float32 = -2.25
                            v11 += 0.1
                            v12 := GenS2{F0: 4294967295, F1: 251}
                            var v13 uint64 = uint64(v1)
                            v6 -= int16(v2)
                            var v14 uint16 = v1
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v14
                        case "a":
                            v1 -= (v1 % (((v1 | v1)) | 1))
                        default:
                            v5 = append(v5, 42)
                        }
                    }
                }
                fmt.Println(v2, v8, v6)
                _ = fi10
            }
            fmt.Println(fi9)
            _ = fi9
        }
        if len(v3) > 0 {
            v3[0] = ("foo bar" + v8)
        }
    }
    if len(v5) > 0 {
        v5[0] = 199
    }
    v8 = (v8 + (v8 + "foo bar"))
    v1 -= (v1 * (v1 / ((638) | 1)))
    fmt.Println(v6, v1)
    var v15 float64 = ((-2.25 - float64(0.1)) * (1.5 - float64(1.5)))
    var v16 uint = 882
    if true {
        switch -2147483648 {
        case 1:
            v6 = v6
        default:
            v4 -= (v4 & (v4 / ((int(v6)) | 1)))
        }
        var v17 string = (v8 + "\n\t")
        _ = v17
    }
    switch v8 {
    case "":
        v18 := GenS1{F0: 0, F1: (v16 << (((v16 >> (((^(v16))) & 63))) & 63)), F2: 81}
        _ = v18
    }
    v6 -= (v6 - v6)
    v8 = "a"
    var v19 uint = (v16 + (v16 ^ (v16 << ((v16) & 63))))
    switch v4 {
    case 4294967295:
        if false {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            _ = len(v5)
            var v20 string = v8
            _ = v20
        } else {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 int8 = -128
                if len(v5) > 0 {
                    v5[0] = 1
                }
                _ = fi21
                _ = v22
            }
            if (!false) {
                for fi23 := 0; fi23 < 5; fi23++ {
                    _ = len(v3)
                    if (!(!true)) {
                        v24 := GenS3{F0: v2, F1: false}
                        if false {
                            if len(v3) > 0 {
                                v3[0] = v8
                            }
                            var v25 int64 = (v2 << (((v16 >> ((v19) & 63))) & 63))
                            _ = v25
                        } else {
                            switch v4 {
                            case 42:
                                if (v6 != 0) {
                                    var v26 uint = (v16 * 1)
                                    _ = v26
                                } else {
                                    if ((v2 / ((v2) | 1)) <= (v2 | (v2 + (v2 % ((v2) | 1))))) {
                                        v2 += v2
                                        var v27 uint32 = 65535
                                        _ = v27
                                    }
                                    for fi28 := 0; fi28 < 1; fi28++ {
                                        v29 := GenS3{F0: (^(v2)), F1: false}
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            case 869:
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    if false {
                                        if (int16(v19) <= (v6 ^ (v6 ^ 572))) {
                                            var v31 string = v8
                                            _ = v31
                                        }
                                    }
                                    v19 = (v19 << ((v19) & 63))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v32 float64 = (v15 * v15)
                                    v16 += v16
                                    _ = fi30
                                    _ = v32
                                }
                            case 799:
                                v33 := GenS4{F0: 100, F1: uint16(v19)}
                                _ = v33
                            default:
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    v16 += 0
                                    _ = fi34
                                }
                            }
                        }
                        _ = v24
                    } else {
                        v35 := GenS5{F0: (^(v16))}
                        v36 := GenS6{F0: 42, F1: 818, F2: 2147483647}
                        _ = v35
                        _ = v36
                    }
                    _ = fi23
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 uint16 = v1
                    var v39 float32 = 1e10
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                if ((v15 - v15) < ((v15 - v15) - 1.5)) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    switch -2147483648 {
                    case -466:
                        v16 = (^(v16))
                    case -9223372036854775808:
                        for fi40 := 0; fi40 < 16; fi40++ {
                            var v41 uint32 = 1
                            _ = fi40
                            _ = v41
                        }
                    case 4294967295:
                        var v42 int32 = 2147483647
                        _ = v42
                    default:
                        v2 = (^(v2))
                    }
                    fmt.Println(v4, v19)
                }
            }
            var v43 bool = true
            _ = v43
        }
    case 1:
        for fi44 := 0; fi44 < 12; fi44++ {
            v45 := GenS4{F0: 127, F1: v1}
            if len(v5) > 0 {
                v5[0] = 1
            }
            v46 := GenS7{F0: 127, F1: 2147483647}
            _ = fi44
            _ = v45
            _ = v46
        }
    case 280:
        var v47 bool = ((v16 >= 885) || false)
        _ = v47
    default:
        v16 = (v19 % ((112) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
