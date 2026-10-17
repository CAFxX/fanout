package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint64; F1 int16; F2 int8 }

func main() {
    var v1 uint64 = 390
    var v2 int32 = 1
    var v3 bool = (!false)
    var v4 int = 127
    v5 := GenS1{F0: 255, F1: 255}
    v6 := []int{-756}
    if (v3 || v3) {
        var v7 string = "foo bar"
        _ = v7
    }
    v8 := GenS1{F0: 128, F1: 19}
    switch "\n\t" {
    case "a":
        var v9 uint = uint(v1)
        _ = v9
    case "":
        switch "a" {
        case "\n\t":
            v10 := GenS2{F0: 132, F1: -1, F2: (v2 % ((v2) | 1))}
            _ = v10
        case "foo bar":
            switch "hello" {
            case "a":
                if v3 {
                    fmt.Println(v2, v4)
                } else {
                    var v11 uint64 = v1
                    _ = v11
                }
            case "":
                var v12 int64 = -335
                _ = v12
            }
        case "":
            if false {
                for fi13 := 0; fi13 < 3; fi13++ {
                    _ = len(v6)
                    _ = fi13
                }
            }
        }
    default:
        _ = len(v6)
    }
    if v3 {
        switch (("a" + "") + "a") {
        case "hello":
            var v14 uint8 = 1
            _ = v14
        case "a":
            if true {
                for fi15 := 0; fi15 < 16; fi15++ {
                    if v3 {
                        if false {
                            v2 += v2
                            if (((-2.25 + 1e10) != 1.5) && (0 != 156)) {
                                v2 += v2
                                if (!false) {
                                    fmt.Println(v3, fi15)
                                }
                            }
                        } else {
                            v2 -= v2
                        }
                        v3 = (4294967295 <= 42)
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v1 -= (v1 * v1)
                            v17 := GenS2{F0: 497, F1: 13, F2: (v2 << ((42) & 63))}
                            var v18 float64 = -2.25
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = fi15
                }
            } else {
                fmt.Println(v4, v3, v4)
                var v19 uint = 42
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v19
            }
        }
        v20 := GenS2{F0: -469, F1: -128, F2: v2}
        var v21 int = v4
        _ = v20
        _ = v21
    }
    if (!(94 < 255)) {
        for fi22 := 0; fi22 < 11; fi22++ {
            v4 = (v4 % (((fi22 + 32768)) | 1))
            v4 = v4
            _ = fi22
        }
        if v3 {
            v1 = (v1 * v1)
            v6 = append(v6, (v4 * (-(v4))))
            if ((v3 || v3) && (true || (32767 != -9223372036854775808))) {
                if len(v6) > 0 {
                    v6[0] = v4
                }
                if v3 {
                    v3 = true
                } else {
                    v6 = append(v6, v4)
                }
            }
        } else {
            v6 = append(v6, (v4 % ((42) | 1)))
        }
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            if false {
                v25 := GenS2{F0: 424, F1: 127, F2: v2}
                _ = v25
            } else {
                var v26 float64 = 3.14159
                _ = len(v6)
                _ = v26
            }
            if true {
                var v27 float64 = 3.14159
                if (!v3) {
                    _ = len(v6)
                    v28 := GenS3{F0: 862}
                    _ = v28
                } else {
                    v1 = (v1 + v1)
                    fmt.Println(v27, v4)
                    if ((-(v2)) >= (-(v2))) {
                        v2 -= (v2 >> ((0) & 63))
                        var v29 uint8 = 128
                        _ = v29
                    }
                    var v30 int = v4
                    v6 = append(v6, (fi24 + fi23))
                    _ = v30
                }
                _ = v27
            } else {
                v6 = append(v6, v4)
                switch "\n\t" {
                case "hello":
                    switch 128 {
                    case 9223372036854775807:
                        fmt.Println(fi24, v3, fi23)
                    case -251:
                        var v31 bool = true
                        _ = v31
                    case -948:
                        var v32 string = "\n\t"
                        _ = v32
                    }
                default:
                    switch "hello" {
                    case "":
                        v33 := GenS4{F0: 2, F1: 42}
                        _ = v33
                    default:
                        v3 = (255 == 32767)
                    }
                }
                var v34 uint32 = 65535
                _ = v34
            }
            _ = fi24
        }
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 uint32 = 42
            var v37 uint = 0
            var v38 int64 = 1
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        var v39 float64 = float64(0.0)
        _ = fi23
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        v2 += (v2 % (((v2 ^ v2)) | 1))
        for fi41 := 0; fi41 < 7; fi41++ {
            switch "\n\t" {
            case "hello":
                v3 = (!v3)
            case "a":
                var v42 int8 = 100
                _ = v42
            case "\n\t":
                v3 = (370 > 65535)
            default:
                v1 -= v1
            }
            switch "a" {
            case "hello":
                v43 := GenS1{F0: 128, F1: 151}
                _ = v43
            case "foo bar":
                var v44 string = "foo bar"
                _ = v44
            }
            _ = fi41
        }
        v45 := GenS5{F0: (v1 ^ v1), F1: -32768, F2: 100}
        _ = fi40
        _ = v45
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
