package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint64) int {
    p0 += p0
    v6 := GenS1{F0: (true || (32768 != 32768)), F1: 1}
    if true {
        p0 += p0
        var v7 int64 = 0
        var v8 bool = (554 > (p0 / (((p0 - p0)) | 1)))
        _ = v7
        _ = v8
    } else {
        switch 0 {
        case 2147483647:
            p0 = p0
        case -140:
            p0 = 0
        default:
            fmt.Println(p0)
        }
        p0 -= (p0 ^ p0)
        p0 = (p0 >> ((255) & 63))
    }
    p0 = (p0 * 2)
    _ = p0
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 1000
    var v3 uint32 = 250
    v4 := make(map[string]bool)
    var v5 bool = true
    v9 := []int{433}
    v5 = v5
    if false {
        var v10 uint8 = 114
        var v11 float64 = (1e10 - 0.0)
        v2 = int16(v3)
        _ = v10
        _ = v11
    } else {
        var v12 uint64 = 0
        _ = v12
    }
    var v13 int = 4294967295
    if false {
        v1 = v1
        v5 = v5
    } else {
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    fmt.Println(v2, v1)
    if (42 == 42) {
        if v5 {
            for fi15 := 0; fi15 < 14; fi15++ {
                v4[((v1 + v1) + v1)] = false
                _ = fi15
            }
            v1 = (v1 + "hello")
            var v16 int32 = -114
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v2 += (-(v2))
            v9 = append(v9, (v13 >> ((18446744073709551615) & 63)))
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS2{F0: (^(v13)), F1: v3, F2: -128}
                var v20 float64 = -2.25
                if false {
                    if v5 {
                        if len(v9) > 0 {
                            v9[0] = (fi18 + (fi18 % ((genFn0(563)) | 1)))
                        }
                        var v21 uint8 = 255
                        v22 := GenS2{F0: fi18, F1: (^(v3)), F2: 42}
                        _ = v21
                        _ = v22
                    } else {
                        switch genFn0(18446744073709551615) {
                        case 832:
                            v1 = ("" + "")
                        case 127:
                            var v23 int = (^(v13))
                            _ = v23
                        case -9223372036854775808:
                            if (v1 > "hello") {
                                switch v1 {
                                case "foo bar":
                                    v24 := v4[v1]
                                    _ = v24
                                case "":
                                    if len(v9) > 0 {
                                        v9[0] = 32767
                                    }
                                case "hello":
                                    fmt.Println(v20, v5)
                                }
                                v25 := GenS1{F0: ((v20 * ((v20 + 0.0) / v20)) > v20), F1: 114}
                                var v26 uint64 = 42
                                _ = len(v4)
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(v20, v13, v5)
                            }
                        }
                        var v27 uint = 2
                        _ = v27
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    } else {
        _ = len(v9)
        var v28 int64 = 42
        _ = v28
    }
    v29 := GenS3{F0: (v2 * (v2 * (^(v2)))), F1: ("\n\t" + v1), F2: 0}
    switch v13 {
    case 576:
        var v30 float32 = float32((float64(1.5) - 1e10))
        _ = v30
    case 0:
        var v31 uint16 = 2
        _ = v31
    }
    if true {
        delete(v4, (v1 + ("\n\t" + v1)))
    }
    delete(v4, (v1 + ""))
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 uint32 = v3
        var v34 uint8 = 23
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: true, F1: 18446744073709551615, F2: 65535}
    switch v1 {
    case "hello":
        v36 := GenS2{F0: (v13 & (v13 << ((369) & 63))), F1: 2, F2: -79}
        _ = v36
    case "a":
        v5 = false
    default:
        switch "" {
        case "hello":
            v2 = (v2 - -1)
        case "":
            v37 := GenS5{F0: 1, F1: 32767}
            _ = v37
        case "foo bar":
            v38 := v4["hello"]
            _ = v38
        }
    }
    v39 := GenS3{F0: (v2 + (v2 & v2)), F1: (("hello" + (v1 + v1)) + "a"), F2: 127}
    v1 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
