package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64, p1 int8) int16 {
    var v7 uint = 0
    var v8 int = -2147483648
    v9 := GenS1{F0: 1000, F1: 127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int16 = -908
    var v2 int16 = v1
    var v3 uint32 = 65535
    var v4 int64 = 1
    var v5 string = ""
    v6 := []uint32{v3, (v3 / (((v3 + v3)) | 1)), 904, 638}
    switch 42 {
    case 9223372036854775807:
        v2 = (v1 << ((uint(v3)) & 63))
    case 32767:
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 -= int16(v4)
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
    }
    switch 4294967295 {
    case 4294967295:
        v12 := GenS1{F0: (-(v2)), F1: 42}
        _ = v12
    case 422:
        _ = len(v6)
    case 9223372036854775807:
        var v13 string = ("foo bar" + (((("" + v5) + v5) + (v5 + v5)) + ((v5 + v5) + v5)))
        _ = v13
    default:
        var v14 int64 = 1
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        switch v5 {
        case "foo bar":
            var v16 int = -9223372036854775808
            _ = v16
        default:
            for fi17 := 0; fi17 < 11; fi17++ {
                if (466 >= 2) {
                    if ((v5 + "") < v5) {
                        switch fi17 {
                        case 9223372036854775807:
                            switch (fi17 - -1) {
                            case 32767:
                                v6 = append(v6, v3)
                            default:
                                if len(v6) > 0 {
                                    v6[0] = v3
                                }
                            }
                        case -767:
                            var v18 uint32 = (v3 - 2147483647)
                            _ = v18
                        case -111:
                            if (v4 < v4) {
                                v5 = (("\n\t" + v5) + v5)
                                fmt.Println(fi15)
                                var v19 int = fi15
                                var v20 int8 = 77
                                v21 := GenS1{F0: v1, F1: 128}
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v22 := GenS2{F0: 1, F1: 42, F2: 148}
                                _ = v22
                            }
                        default:
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                        }
                        var v23 uint64 = 4294967295
                        v4 -= (v4 + v4)
                        switch v5 {
                        case "a":
                            fmt.Println(fi15, v2)
                        case "hello":
                            var v24 string = "\n\t"
                            _ = v24
                        }
                        _ = v23
                    }
                    var v25 uint = 4294967295
                    v26 := GenS3{F0: 0}
                    v27 := GenS2{F0: 255, F1: 1, F2: 131}
                    v2 += 0
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v28 := GenS2{F0: 42, F1: 35, F2: 2}
                _ = fi17
                _ = v28
            }
        }
        v29 := GenS4{F0: 0}
        switch v5 {
        case "hello":
            var v30 uint16 = 982
            _ = v30
        case "a":
            v6 = append(v6, v3)
        default:
            fmt.Println(v5, fi15)
        }
        _ = fi15
        _ = v29
    }
    var v31 int = -1
    var v32 int8 = -1
    v31 = 128
    var v33 uint = 255
    for fi34 := 0; fi34 < 1; fi34++ {
        v5 = (("foo bar" + v5) + (v5 + v5))
        _ = fi34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint64 = 0
        if len(v6) > 0 {
            v6[0] = uint32(v36)
        }
        _ = fi35
        _ = v36
    }
    var v37 uint = 255
    v38 := GenS5{F0: uint(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
