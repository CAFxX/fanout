package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    var v1 int32 = 42
    v2 := []uint32{42, 0}
    var v3 uint16 = 32767
    v4 := []uint32{0}
    fmt.Println(v3)
    v3 -= v3
    if len(v4) > 0 {
        v4[0] = 519
    }
    if ((2 == 255) || (651 <= 515)) {
        fmt.Println(v1, v3)
        v3 = v3
        v5 := GenS1{F0: 2147483648, F1: 0}
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        switch 709 {
        case 127:
            var v7 bool = true
            _ = v7
        case -513:
            if false {
                switch fi6 {
                case 454:
                    v3 -= (v3 * (v3 % (((v3 << ((875) & 63))) | 1)))
                }
            } else {
                var v8 uint64 = 4294967295
                v8 = (^(v8))
                _ = v8
            }
        }
        v9 := GenS2{F0: 0, F1: 42, F2: -128}
        v3 = (^(v3))
        _ = fi6
        _ = v9
    }
    v3 = v3
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            fmt.Println(fi10, v3, v1)
            for fi11 := 0; fi11 < 11; fi11++ {
                switch "hello" {
                case "\n\t":
                    var v12 float32 = float32(-2.25)
                    _ = v12
                case "foo bar":
                    var v13 bool = true
                    _ = v13
                case "":
                    fmt.Println(v1, v1, fi11)
                }
                _ = fi11
            }
        } else {
            var v14 uint = 65535
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := GenS3{F0: 222, F1: 18446744073709551615}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v17 int32 = v1
        if true {
            v18 := GenS4{F0: 19, F1: v1, F2: v3}
            for fi19 := 0; fi19 < 5; fi19++ {
                v17 -= v17
                switch -588 {
                case 0:
                    v20 := GenS5{F0: -1, F1: (v17 % ((v17) | 1))}
                    _ = v20
                default:
                    switch 32767 {
                    case 9223372036854775807:
                        if false {
                            if true {
                                v21 := GenS1{F0: 2, F1: -837}
                                var v22 int8 = 100
                                v22 += -1
                                _ = v21
                                _ = v22
                            }
                        } else {
                            var v23 int32 = 2147483647
                            _ = v23
                        }
                    case 127:
                        v3 -= v3
                    case -285:
                        var v24 string = ("a" + "")
                        _ = v24
                    default:
                        fmt.Println(fi19, fi19)
                    }
                    v3 -= 996
                    if true {
                        if (1 >= 2) {
                            var v25 int8 = 42
                            var v26 int16 = 1
                            var v27 string = "hello"
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        if ((-1 < 127) || (42 >= 166)) {
                            if len(v4) > 0 {
                                v4[0] = 684
                            }
                            if false {
                                if (!false) {
                                    v1 += (v17 / (((v17 / ((v1) | 1))) | 1))
                                }
                            } else {
                                switch "" {
                                case "hello":
                                    var v28 int = (fi19 + fi10)
                                    _ = v28
                                }
                            }
                            v1 = (v1 - v17)
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                switch ("\n\t" + ("a" + "a")) {
                                case "":
                                    var v29 int32 = v1
                                    _ = v29
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    var v30 string = ""
                                    _ = v30
                                }
                            case "\n\t":
                                fmt.Println(v3, v17)
                            }
                            v1 = v17
                            fmt.Println(fi10)
                        }
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            var v31 int = fi10
                            _ = v31
                        }
                    }
                }
                _ = fi19
            }
            v32 := GenS4{F0: 0, F1: v17, F2: (v3 >> ((600) & 63))}
            _ = v18
            _ = v32
        } else {
            if ((18446744073709551615 > 1) || false) {
                v33 := GenS3{F0: 42, F1: 0}
                _ = v33
            }
        }
        _ = fi10
        _ = v17
    }
    v34 := GenS6{F0: 1, F1: uint(v3), F2: 1}
    var v35 uint16 = v3
    var v36 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
