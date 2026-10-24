package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 int8 = -128
    var v2 int64 = -1
    var v3 int16 = -794
    v4 := []uint32{2}
    v5 := make(map[uint32]int)
    fmt.Println(v2, v2)
    if (0.0 == 0.1) {
        var v6 int64 = v2
        v1 = 127
        _ = v6
    }
    var v7 float32 = 3.14159
    v8 := GenS1{F0: "a", F1: 0}
    v9 := v5[900]
    if false {
        v2 = (-(v2))
        v7 += v7
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int64 = -1
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v1)
            v13 := GenS2{F0: 1, F1: 255, F2: ("" + "")}
            if len(v4) > 0 {
                v4[0] = 65535
            }
            _ = fi12
            _ = v13
        }
        v14 := GenS3{F0: (^(v3)), F1: v1}
        _ = fi10
        _ = v11
        _ = v14
    }
    switch "a" {
    case "a":
        for fi15 := 0; fi15 < 13; fi15++ {
            v7 -= (v7 * 1e10)
            fmt.Println(v7, v9, fi15)
            var v16 uint8 = 0
            _ = fi15
            _ = v16
        }
    case "foo bar":
        v17 := GenS4{F0: 0, F1: 1}
        _ = v17
    case "":
        switch ((("a" + (("hello" + ("hello" + "\n\t")) + "foo bar")) + "") + "a") {
        case "":
            var v18 float64 = 0.1
            _ = v18
        case "\n\t":
            if false {
                var v19 int8 = 12
                var v20 bool = (false && false)
                _ = v19
                _ = v20
            } else {
                if false {
                    v21 := GenS3{F0: 42, F1: (v1 % (((v1 - v1)) | 1))}
                    _ = v21
                } else {
                    v5[1] = 2147483647
                    v22 := GenS1{F0: "hello", F1: 1000}
                    v7 = (((((v7 + ((v7 + (1.5 + float32(float64(v7)))) * 0.0)) / v7) * 1e10) * ((v7 * ((((1e10 - float32(1e10)) + float32(-2.25)) / v7) - 1e10)) / (((v7 * 1.5) - float32((3.14159 + 0.1))) + float32(-2.25)))) * -2.25)
                    _ = v22
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 uint64 = 18446744073709551615
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    if false {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 int16 = (v3 & (v3 | (v3 << ((1) & 63))))
                            _ = fi26
                            _ = v27
                        }
                    }
                    if true {
                        v4 = append(v4, 1)
                        v1 += (v1 & 42)
                        v5[0] = (fi25 ^ int(v2))
                    } else {
                        if (1 < 261) {
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                            switch "foo bar" {
                            case "":
                                _ = len(v5)
                            case "hello":
                                switch "hello" {
                                case "\n\t":
                                    _ = len(v4)
                                case "a":
                                    delete(v5, 2147483647)
                                }
                            case "a":
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    v29 := GenS5{F0: 255}
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        } else {
                            switch fi25 {
                            case 42:
                                fmt.Println(v3, v1)
                            case -9223372036854775808:
                                v30 := GenS2{F0: 928, F1: 565, F2: "a"}
                                _ = v30
                            case 44:
                                switch ("\n\t" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) {
                                case "foo bar":
                                    v9 += 32767
                                }
                            default:
                                delete(v5, 726)
                            }
                            if false {
                                fmt.Println(fi25, v1)
                                var v31 int32 = -234
                                _ = v31
                            } else {
                                var v32 int64 = (v2 + v2)
                                _ = v32
                            }
                            v1 -= v1
                        }
                        var v33 int32 = 1
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 7; fi34++ {
                        v9 = fi25
                        var v35 uint16 = 0
                        _ = fi34
                        _ = v35
                    }
                    _ = fi25
                }
            }
        case "a":
            for fi36 := 0; fi36 < 1; fi36++ {
                v2 = v2
                v37 := GenS6{F0: (v3 % ((v3) | 1)), F1: 65535}
                _ = fi36
                _ = v37
            }
        default:
            var v38 uint = 1
            _ = v38
        }
    }
    if (4294967295 < 2147483647) {
        _ = len(v5)
        if ((v1 & v1) <= v1) {
            if (1.5 <= 1e10) {
                switch 0 {
                case 1:
                    if (false || false) {
                        if ("hello" <= "foo bar") {
                            v9 -= (v9 + v9)
                        }
                    }
                case -1:
                    switch 78 {
                    case 1:
                        v7 += (float32((float64(v7) / 0.1)) * float32(1.5))
                    default:
                        _ = len(v5)
                    }
                default:
                    var v39 float64 = (0.1 - float64(0.0))
                    _ = v39
                }
            } else {
                v3 = v3
                var v40 uint64 = 18446744073709551615
                _ = v40
            }
            v3 = (v3 | v3)
        } else {
            switch "a" {
            case "\n\t":
                _ = len(v5)
            default:
                v41 := v5[42]
                _ = v41
            }
        }
        v3 = (^(v3))
    }
    v42 := GenS6{F0: v3, F1: 57}
    for fi43 := 0; fi43 < 5; fi43++ {
        if len(v4) > 0 {
            v4[0] = 967
        }
        _ = len(v5)
        _ = fi43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 int32 = 1
            v47 := GenS7{F0: 65535, F1: 482, F2: 207}
            _ = fi45
            _ = v46
            _ = v47
        }
        v48 := GenS8{F0: 255}
        v2 = v2
        _ = fi44
        _ = v48
    }
    if true {
        var v49 uint16 = 0
        fmt.Println(v49, v9, v9)
        delete(v5, 696)
        _ = v49
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
