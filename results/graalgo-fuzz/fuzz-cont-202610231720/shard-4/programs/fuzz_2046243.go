package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint) uint {
    if (-9223372036854775808 > 42) {
        switch -2147483648 {
        case 2147483647:
            var v5 int64 = 42
            _ = v5
        case 4294967295:
            v6 := []int{32767, 0}
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = 255
        v9 := []int{(fi7 | fi7)}
        var v10 uint8 = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int16 = 347
    var v12 int32 = 1
    _ = p0
    _ = v11
    _ = v12
    return p0
}
func genFn1(p0 uint64, p1 uint, p2 uint8) uint32 {
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float32 = ((float32(-2.25) + 1.5) * 3.14159)
        var v15 uint16 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v16 int8 = -128
        _ = v16
    case "a":
        fmt.Println(p2, p1)
    case "":
        v17 := GenS1{F0: 0, F1: 42}
        _ = v17
    default:
        if true {
            fmt.Println(p1, p0, p0)
        }
    }
    v18 := GenS1{F0: -476, F1: 127}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: 116, F1: int8(p1)}
        if ((^(p2)) != 0) {
            v21 := GenS1{F0: 0, F1: 37}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    switch "a" {
    case "a":
        p2 += (p2 | p2)
    case "foo bar":
        v22 := GenS2{F0: -32768, F1: -1, F2: 6}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return uint32(p1)
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 32767
    v3 := []uint32{4294967295, 4294967295}
    v4 := []uint32{273, 2, 2147483647, 848}
    var v23 uint16 = (v2 % (((v2 * v2)) | 1))
    if len(v4) > 0 {
        v4[0] = 2
    }
    v24 := GenS2{F0: v1, F1: v1, F2: -67}
    var v25 string = "\n\t"
    v2 -= 65535
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 bool = true
        _ = fi26
        _ = v27
    }
    if (65535 <= 255) {
        if len(v4) > 0 {
            v4[0] = 2
        }
    }
    switch (v25 + v25) {
    case "\n\t":
        v28 := GenS2{F0: 1000, F1: (v1 << ((65535) & 63)), F2: -1}
        _ = v28
    }
    if (127 == 0) {
        v29 := GenS1{F0: -1, F1: -1}
        v23 = (v2 & (v23 | v23))
        _ = v29
    } else {
        switch v25 {
        case "\n\t":
            if false {
                v2 += 42
            }
        default:
            v2 = 829
        }
        if (((v2 & v23) <= (v2 & (v2 * v23))) || true) {
            for fi30 := 0; fi30 < 16; fi30++ {
                v1 -= 0
                _ = fi30
            }
            v31 := GenS3{F0: v25, F1: true}
            _ = v31
        }
        switch (v25 + (((v25 + ((("a" + v25) + "a") + (("hello" + (("a" + "\n\t") + (v25 + v25))) + v25))) + (("\n\t" + v25) + (v25 + v25))) + (("foo bar" + "") + "\n\t"))) {
        case "\n\t":
            if ("foo bar" >= "") {
                if true {
                    v23 -= v2
                    var v32 int64 = -32
                    _ = v32
                }
                for fi33 := 0; fi33 < 18; fi33++ {
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v34 := GenS4{F0: (fi33 - (fi33 << ((genFn0(4294967295)) & 63)))}
                            _ = v34
                        } else {
                            if false {
                                if false {
                                    for fi35 := 0; fi35 < 12; fi35++ {
                                        v25 = v25
                                        fmt.Println(v2, v23, v2)
                                        _ = fi35
                                    }
                                    var v36 int32 = 750
                                    if (1 < 18446744073709551615) {
                                        switch 0 {
                                        case 0:
                                            var v37 int = fi33
                                            _ = v37
                                        case -2147483648:
                                            for fi38 := 0; fi38 < 11; fi38++ {
                                                v39 := GenS5{F0: 2147483647}
                                                _ = fi38
                                                _ = v39
                                            }
                                        default:
                                            v36 -= v36
                                        }
                                        var v40 float64 = 0.1
                                        _ = v40
                                    } else {
                                        for fi41 := 0; fi41 < 19; fi41++ {
                                            v42 := GenS5{F0: 65535}
                                            if (42 <= 685) {
                                                fmt.Println(v1)
                                            } else {
                                                v43 := GenS2{F0: (^(v1)), F1: v1, F2: -1}
                                                _ = v43
                                            }
                                            _ = fi41
                                            _ = v42
                                        }
                                    }
                                    _ = v36
                                }
                                if len(v4) > 0 {
                                    v4[0] = genFn1(0, 255, 223)
                                }
                            } else {
                                var v44 float32 = (0.0 * (0.1 + ((1.5 / 1.0) + -2.25)))
                                fmt.Println(v1, fi33, v23)
                                _ = v44
                            }
                        }
                    case "":
                        v23 = (v2 << ((2) & 63))
                    case "foo bar":
                        switch ((v25 + "") + "\n\t") {
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        default:
                            switch (v25 + v25) {
                            case "hello":
                                v45 := GenS6{F0: 2, F1: v1}
                                _ = v45
                            case "foo bar":
                                switch "hello" {
                                case "":
                                    if (4294967295 < 18446744073709551615) {
                                        var v46 float32 = ((0.0 + ((0.1 + (1.5 + 0.0)) + 0.1)) + float32(float64(float32(0.1))))
                                        v46 -= v46
                                        _ = v46
                                    }
                                }
                            case "\n\t":
                                for fi47 := 0; fi47 < 6; fi47++ {
                                    v48 := GenS7{F0: 1}
                                    v23 -= v23
                                    var v49 int16 = -712
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                    }
                    _ = fi33
                }
            }
        case "foo bar":
            v23 += (v2 | v2)
        case "a":
            var v50 uint16 = (v2 * (^(v2)))
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        if len(v3) > 0 {
            v3[0] = 284
        }
        v2 -= v23
        _ = fi51
    }
    v2 = 2
    switch 1 {
    case 2147483647:
        v25 = v25
    }
    v2 += v2
    v25 = (((v25 + (v25 + v25)) + "\n\t") + (v25 + (v25 + (("hello" + v25) + "foo bar"))))
    var v52 float32 = (-2.25 + (float32(((-2.25 - float64(0.1)) / (float64((1e10 / 0.1)) + 0.0))) / (1e10 - ((((0.1 * -2.25) + 1e10) + float32(3.14159)) - (0.1 * -2.25)))))
    var v53 uint32 = 0
    fmt.Println(v52, v53)
    var v54 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
