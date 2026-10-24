package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int; F2 string }
type GenS5 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint8 = 42
    var v2 uint = 440
    var v3 uint64 = 2
    var v4 uint = (v2 % ((v2) | 1))
    v5 := GenS1{F0: (0.1 >= 1e10), F1: 100, F2: -676}
    v4 += (^(v2))
    v1 = v1
    var v6 uint = (v2 ^ 834)
    fmt.Println(v4, v1, v2)
    var v7 uint16 = 534
    v3 = (v3 - 1)
    v8 := GenS1{F0: false, F1: -1, F2: -739}
    var v9 int8 = 100
    v10 := []int{}
    switch int(v7) {
    case 0:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case -1:
        for fi11 := 0; fi11 < 7; fi11++ {
            switch (fi11 & (-(fi11))) {
            case 32767:
                var v12 int64 = -9223372036854775808
                _ = v12
            case 1:
                var v13 int = (-(fi11))
                _ = v13
            default:
                v14 := GenS2{F0: false, F1: uint32(v7)}
                _ = v14
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                v9 = (v9 << ((v6) & 63))
                v3 += v3
                v10 = append(v10, (fi11 >> ((2) & 63)))
                v10 = append(v10, fi15)
                v9 += v9
                var v16 int32 = 754
                _ = fi15
                _ = v16
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi11
        }
    case 2147483647:
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(fi17)
                v7 -= 42
                switch fi18 {
                case 0:
                    if (true && (!false)) {
                        if true {
                            fmt.Println(fi18)
                            var v19 uint16 = (v7 * (^(v7)))
                            switch ("hello" + "hello") {
                            case "\n\t":
                                var v20 uint16 = (v7 << ((4294967295) & 63))
                                _ = v20
                            case "a":
                                var v21 float32 = 0.1
                                _ = v21
                            case "hello":
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            default:
                                switch "hello" {
                                case "":
                                    var v22 int32 = 644
                                    _ = v22
                                case "hello":
                                    v23 := GenS1{F0: true, F1: (v9 * (v9 * v9)), F2: -9223372036854775808}
                                    _ = v23
                                case "a":
                                    var v24 int64 = 9223372036854775807
                                    _ = v24
                                default:
                                    v7 += v19
                                }
                            }
                            _ = v19
                        } else {
                            v4 = (^(v2))
                            v1 = 106
                        }
                        _ = len(v10)
                        switch (("" + "\n\t") + "hello") {
                        case "hello":
                            v1 = v1
                        case "\n\t":
                            _ = len(v10)
                        case "":
                            if ((v9 & v9) < 100) {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                v2 += v6
                                v25 := GenS3{F0: v9}
                                _ = v25
                            }
                        }
                    } else {
                        v2 = (^(v2))
                        if true {
                            switch "hello" {
                            case "a":
                                switch 4294967295 {
                                case 128:
                                    fmt.Println(v2)
                                }
                            case "":
                                switch ("a" + "\n\t") {
                                case "":
                                    if (((0.1 - ((0.0 + (-2.25 + 0.0)) / float64(((1e10 / ((0.1 - 3.14159) + 0.0)) * 1.5)))) - 1e10) < 0.0) {
                                        if true {
                                            if true {
                                                var v26 int64 = 1
                                                v7 -= (v7 ^ uint16(v3))
                                                var v27 float64 = 1.5
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v2 -= v4
                                            }
                                            if false {
                                                if ("hello" == ("" + "foo bar")) {
                                                    switch fi17 {
                                                    case 2147483647:
                                                        v2 = (v2 | (v4 * (v2 + (v2 >> (((v2 | v6)) & 63)))))
                                                    case 4294967295:
                                                        v1 += 42
                                                    case 1:
                                                        if false {
                                                            v7 = v7
                                                        }
                                                    }
                                                } else {
                                                    if (!true) {
                                                        fmt.Println(v9, v1)
                                                    } else {
                                                        fmt.Println(v9, fi18, fi17)
                                                        var v28 uint = (v2 - (v4 / ((v2) | 1)))
                                                        v10 = append(v10, fi18)
                                                        switch fi18 {
                                                        case -2147483648:
                                                            if true {
                                                                v29 := GenS3{F0: v9}
                                                                _ = v29
                                                            } else {
                                                                v28 -= v4
                                                                v1 -= v1
                                                            }
                                                        default:
                                                            var v30 float32 = 0.1
                                                            _ = v30
                                                        }
                                                        fmt.Println(v4, v1, fi17)
                                                        fmt.Println(fi17, fi18)
                                                        _ = v28
                                                    }
                                                    var v31 int32 = -112
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    fmt.Println(fi18)
                                                    _ = v31
                                                }
                                                fmt.Println(v4, fi18)
                                                var v32 uint16 = (v7 + uint16(v3))
                                                _ = v32
                                            } else {
                                                v7 -= (v7 + (v7 % (((v7 | 35)) | 1)))
                                                v33 := GenS4{F0: int(v3), F1: -347, F2: ""}
                                                _ = len(v10)
                                                _ = v33
                                            }
                                        }
                                        v6 += (v6 % (((v2 + 1)) | 1))
                                        v34 := GenS5{F0: 453, F1: (v1 + v1), F2: -1}
                                        _ = v34
                                    } else {
                                        v6 -= (v6 ^ (v6 ^ 42))
                                        _ = len(v10)
                                        v7 -= v7
                                    }
                                case "\n\t":
                                    v9 += (^(v9))
                                }
                            }
                        } else {
                            v6 = (v6 - (v6 << ((1) & 63)))
                            switch fi17 {
                            case 1:
                                fmt.Println(v6, v3)
                            }
                        }
                    }
                default:
                    v35 := GenS6{F0: (v9 >> ((v2) & 63))}
                    _ = v35
                }
                _ = fi18
            }
            var v36 int32 = 0
            if (1e10 >= 1e10) {
                v9 = -66
            } else {
                if len(v10) > 0 {
                    v10[0] = fi17
                }
                v10 = append(v10, (fi17 * 128))
                var v37 int32 = v36
                _ = v37
            }
            _ = fi17
            _ = v36
        }
    }
    v9 -= v9
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
