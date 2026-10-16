package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 int64 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint64 = 1
    var v2 uint = 18446744073709551615
    var v3 uint64 = v1
    v4 := []int{42, 0, -9223372036854775808}
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v1, v3, v1)
        v6 := GenS1{F0: 56, F1: 2}
        _ = fi5
        _ = v6
    }
    fmt.Println(v2, v2)
    v3 = v3
    v7 := GenS2{F0: (v1 + (v1 * (v1 % (((v1 * (v1 / ((v1) | 1)))) | 1)))), F1: -32768, F2: -9223372036854775808}
    var v8 float64 = 1e10
    fmt.Println(v1, v8, v3)
    v3 = v1
    if (!false) {
        switch -2147483648 {
        case 9223372036854775807:
            switch ((("foo bar" + (("foo bar" + "\n\t") + "")) + "\n\t") + "\n\t") {
            case "foo bar":
                v2 += 739
            default:
                var v9 uint32 = 1
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if (!(-2.25 != 0.1)) {
                var v11 int32 = 0
                switch fi10 {
                case 9223372036854775807:
                    v12 := GenS2{F0: 2, F1: -320, F2: 9223372036854775807}
                    _ = v12
                case 0:
                    v1 += v3
                }
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 uint64 = (v3 + (v3 | (v1 / (((v1 + 18446744073709551615)) | 1))))
                    switch "\n\t" {
                    case "foo bar":
                        switch 2147483647 {
                        case 18:
                            var v15 uint16 = 715
                            _ = v15
                        case 9223372036854775807:
                            var v16 bool = true
                            _ = v16
                        default:
                            if true {
                                fmt.Println(v1, v3)
                                if (2147483647 <= 2147483647) {
                                    if (false && true) {
                                        var v17 uint16 = 255
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v2 = (^(v2))
                                        fmt.Println(v2)
                                        _ = v17
                                    }
                                }
                                var v18 int64 = 340
                                _ = v18
                            } else {
                                if true {
                                    switch "" {
                                    case "a":
                                        _ = len(v4)
                                    }
                                    fmt.Println(v1)
                                }
                                v19 := GenS3{F0: -1, F1: -9223372036854775808}
                                fmt.Println(v2, v8, v14)
                                _ = v19
                            }
                        }
                    case "":
                        if len(v4) > 0 {
                            v4[0] = (fi10 * fi10)
                        }
                        v4 = append(v4, fi13)
                        var v20 float32 = 1e10
                        fmt.Println(v14, fi10, fi10)
                        if false {
                            v2 += v2
                            switch "a" {
                            case "\n\t":
                                switch 200 {
                                case 127:
                                    if false {
                                        v8 = v8
                                    }
                                case 128:
                                    var v21 float32 = 0.1
                                    _ = v21
                                default:
                                    v22 := GenS4{F0: 1, F1: true, F2: 1}
                                    _ = v22
                                }
                            case "foo bar":
                                if (false && (true && false)) {
                                    var v23 uint8 = 1
                                    _ = v23
                                } else {
                                    if (!((908 == 1) || true)) {
                                        var v24 string = "a"
                                        v25 := GenS1{F0: 65535, F1: uint8(v3)}
                                        var v26 uint64 = (^(v1))
                                        v14 += v14
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    v27 := GenS4{F0: 1000, F1: true, F2: -1}
                                    _ = v27
                                }
                            default:
                                var v28 uint = (^(v2))
                                v4 = append(v4, (fi13 * (fi10 << (((v28 & 255)) & 63))))
                                _ = v28
                            }
                            var v29 uint32 = 1
                            _ = v29
                        } else {
                            fmt.Println(v20, v11)
                            v4 = append(v4, fi13)
                            v2 -= v2
                        }
                        _ = v20
                    case "hello":
                        var v30 int8 = -128
                        _ = v30
                    default:
                        switch (fi10 + (fi13 & fi13)) {
                        case 127:
                            var v31 float32 = (0.1 * -2.25)
                            _ = v31
                        default:
                            var v32 float32 = ((1.5 + 1.5) * float32((v8 - float64(0.0))))
                            _ = v32
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v11
            }
            var v33 uint16 = 66
            _ = len(v4)
            _ = fi10
            _ = v33
        }
        v2 = (^(v2))
    } else {
        switch 1 {
        case -1:
            if true {
                switch 0 {
                case 32768:
                    v1 += (v1 - 18446744073709551615)
                case -539:
                    fmt.Println(v8)
                }
                for fi34 := 0; fi34 < 3; fi34++ {
                    v1 -= (v1 % (((v1 - (v3 % ((1) | 1)))) | 1))
                    _ = fi34
                }
                switch -671 {
                case 32768:
                    v8 += v8
                case 128:
                    var v35 int = int(v2)
                    _ = v35
                }
            } else {
                if true {
                    var v36 int32 = 14
                    v37 := GenS1{F0: 1, F1: 46}
                    _ = v36
                    _ = v37
                } else {
                    fmt.Println(v3)
                    switch "a" {
                    case "\n\t":
                        if (!(true || true)) {
                            v2 = (v2 + (v2 / (((v2 ^ (^(v2)))) | 1)))
                            v3 += v3
                        }
                    case "a":
                        v3 += (v1 - v1)
                    default:
                        v38 := GenS5{F0: 1000}
                        _ = v38
                    }
                    v39 := GenS2{F0: (v3 << (((v2 ^ 42)) & 63)), F1: -820, F2: 42}
                    _ = v39
                }
            }
        case -322:
            for fi40 := 0; fi40 < 1; fi40++ {
                if (!(("\n\t" + "foo bar") != "a")) {
                    var v41 int64 = 207
                    if true {
                        v4 = append(v4, 505)
                        v41 = (v41 << ((943) & 63))
                    }
                    _ = v41
                } else {
                    _ = len(v4)
                    var v42 bool = false
                    _ = v42
                }
                v43 := GenS5{F0: 32767}
                _ = fi40
                _ = v43
            }
        }
        var v44 int32 = 313
        _ = v44
    }
    var v45 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v7
    _ = v8
}
