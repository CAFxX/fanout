package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 728}
    var v2 int64 = 1
    v3 := []uint32{0, 974}
    var v4 int8 = 30
    switch 42 {
    case -40:
        switch ("a" + "") {
        case "\n\t":
            fmt.Println(v4, v2)
        case "":
            fmt.Println(v4, v2)
        case "a":
            fmt.Println(v2, v2)
        }
    }
    switch 1 {
    case 127:
        switch -99 {
        case -958:
            if true {
                v5 := GenS1{F0: 2, F1: 431}
                v3 = append(v3, 4294967295)
                fmt.Println(v4, v2)
                _ = v5
            }
        case -1:
            v2 = v2
        case 128:
            if (((!true) || true) || true) {
                var v6 int16 = -32768
                v7 := GenS1{F0: 42, F1: 55}
                for fi8 := 0; fi8 < 4; fi8++ {
                    v6 = v6
                    var v9 string = ("a" + "hello")
                    var v10 int16 = (v6 - v6)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                var v11 uint8 = 2
                _ = v11
            }
        default:
            v4 = v4
        }
    case 0:
        v12 := GenS1{F0: 42, F1: 65535}
        _ = v12
    default:
        v4 += (v4 - v4)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = (-(v4))
        switch "a" {
        case "":
            var v14 int8 = (v4 + v4)
            _ = v14
        }
        switch -781 {
        case -2147483648:
            v15 := GenS2{F0: -32768, F1: 2147483647}
            _ = v15
        case -9223372036854775808:
            fmt.Println(v4, fi13)
        case -913:
            v2 += 1
        }
        _ = fi13
    }
    v2 = v2
    v2 -= 9223372036854775807
    var v16 float32 = -2.25
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint = 0
        _ = fi17
        _ = v18
    }
    _ = len(v1)
    switch ("\n\t" + ((("a" + "") + "foo bar") + "a")) {
    case "foo bar":
        var v19 string = "hello"
        _ = v19
    }
    var v20 uint8 = 42
    v21 := GenS3{F0: v4, F1: -304}
    if true {
        var v22 int16 = 0
        _ = v22
    } else {
        v23 := GenS4{F0: int8(v2), F1: (v20 | 1)}
        if (true && (false && (v4 != v4))) {
            if ((-419 == 2147483647) && false) {
                for fi24 := 0; fi24 < 13; fi24++ {
                    _ = len(v1)
                    _ = fi24
                }
                if true {
                    var v25 int16 = -788
                    _ = v25
                } else {
                    v2 = -1
                    v20 += (v20 | 42)
                }
                v20 = v20
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                if (!true) {
                    v20 += 127
                    var v27 int64 = (v2 + (^(v2)))
                    if false {
                        for fi28 := 0; fi28 < 20; fi28++ {
                            fmt.Println(v20, v2, fi26)
                            _ = fi28
                        }
                        var v29 int64 = v27
                        if ((!(true && false)) && (!((fi26 + (fi26 & fi26)) > fi26))) {
                            for fi30 := 0; fi30 < 12; fi30++ {
                                v20 += (v20 / ((46) | 1))
                                v29 += v2
                                var v31 uint8 = v20
                                v2 -= (v29 - (v2 | v27))
                                v29 -= v2
                                if false {
                                    var v32 uint32 = 2
                                    var v33 uint16 = 994
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi30
                                _ = v31
                            }
                            for fi34 := 0; fi34 < 15; fi34++ {
                                var v35 int64 = (v29 - (v27 & (v27 | (v27 / (((-(v29))) | 1)))))
                                switch (fi34 + 9223372036854775807) {
                                case 42:
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                default:
                                    var v36 string = "\n\t"
                                    _ = v36
                                }
                                switch "\n\t" {
                                case "hello":
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                case "\n\t":
                                    v37 := GenS3{F0: 42, F1: v35}
                                    _ = v37
                                default:
                                    _ = len(v1)
                                }
                                _ = fi34
                                _ = v35
                            }
                        }
                        _ = v29
                    } else {
                        var v38 uint16 = 320
                        var v39 uint = 304
                        _ = v38
                        _ = v39
                    }
                    _ = v27
                } else {
                    v20 += v20
                }
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi26
            }
        }
        _ = v23
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        for fi41 := 0; fi41 < 8; fi41++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi41
        }
        _ = fi40
    }
    v42 := GenS5{F0: (v4 + v4)}
    var v43 uint32 = 42
    for fi44 := 0; fi44 < 14; fi44++ {
        if ((!(true || (int32(fi44) < -1))) || true) {
            fmt.Println(v2, v43)
        }
        switch (fi44 + fi44) {
        case 42:
            fmt.Println(v16)
        }
        _ = fi44
    }
    fmt.Println(v16)
    var v45 uint32 = v43
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v45
}
