package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int8 }

func main() {
    var v1 int16 = -1
    v2 := []uint32{2, 2, 1, 2}
    var v3 int64 = 1
    var v4 string = ("foo bar" + "\n\t")
    v5 := []uint32{2147483648}
    var v6 uint = 922
    v7 := GenS1{F0: 484, F1: 65535}
    _ = len(v5)
    v6 = v6
    v6 = (v6 - v6)
    if (true || true) {
        v4 = ("hello" + "\n\t")
        if (853 <= 65535) {
            fmt.Println(v3, v6, v6)
            var v8 int32 = 42
            var v9 int = 0
            _ = v8
            _ = v9
        } else {
            v5 = append(v5, 4294967295)
            v10 := GenS2{F0: 115}
            _ = v10
        }
    }
    var v11 uint8 = 2
    v1 -= v1
    if (2 > 32767) {
        v12 := GenS1{F0: 65535, F1: 255}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS2{F0: 515}
            _ = fi13
            _ = v14
        }
        _ = len(v2)
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 string = (v4 + "foo bar")
            _ = fi15
            _ = v16
        }
    }
    if (true || false) {
        v17 := GenS3{F0: true, F1: 42}
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        var v18 uint8 = (v11 * (v11 % ((v11) | 1)))
        _ = v17
        _ = v18
    }
    fmt.Println(v4)
    v19 := GenS1{F0: 255, F1: 1}
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 float64 = 1.5
        _ = fi20
        _ = v21
    }
    fmt.Println(v6, v11, v1)
    v2 = append(v2, 0)
    for fi22 := 0; fi22 < 12; fi22++ {
        switch 42 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 16; fi23++ {
                switch (fi23 << ((v6) & 63)) {
                case -1:
                    _ = len(v5)
                case 32768:
                    var v24 int64 = (-(v3))
                    v25 := GenS4{F0: 42}
                    _ = v24
                    _ = v25
                case -945:
                    v1 = v1
                }
                _ = fi23
            }
        case 334:
            switch fi22 {
            case 42:
                if (((fi22 / (((fi22 & -439)) | 1)) > (fi22 + 4294967295)) || false) {
                    var v26 uint16 = 7
                    _ = v26
                }
            case -1:
                switch (fi22 >> ((uint(v11)) & 63)) {
                case -9223372036854775808:
                    if (!(!false)) {
                        fmt.Println(v1)
                        if (fi22 != 128) {
                            var v27 int8 = 0
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                fmt.Println(fi28, v11)
                                v11 -= v11
                                _ = fi28
                            }
                            fmt.Println(fi22)
                            switch (fi22 | (-(fi22))) {
                            case 32767:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 2147483647:
                                if (uint16(v6) != 0) {
                                    v29 := GenS4{F0: 2}
                                    _ = v29
                                } else {
                                    v4 = v4
                                }
                            default:
                                v3 = 9223372036854775807
                            }
                        }
                    } else {
                        v30 := GenS4{F0: 1}
                        _ = v30
                    }
                case 42:
                    v31 := GenS5{F0: (v3 + 216), F1: 11}
                    _ = v31
                }
            case 4294967295:
                v2 = append(v2, 0)
            }
        default:
            v3 = -266
        }
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi22
    }
    v32 := GenS1{F0: 368, F1: 42}
    switch v4 {
    case "hello":
        v5 = append(v5, 2147483647)
    case "\n\t":
        v3 += v3
    case "a":
        for fi33 := 0; fi33 < 5; fi33++ {
            if len(v5) > 0 {
                v5[0] = 2
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
