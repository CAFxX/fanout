package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint64 = 516
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 uint8 = 223
    var v5 uint = 0
    v6 := []uint32{2147483648, 156, uint32(v5)}
    if (1 <= uint32(v1)) {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(v5, v1)
            switch "hello" {
            case "\n\t":
                if (!false) {
                    v2 = (v2 + "a")
                } else {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v7 = v7
                        var v10 uint16 = 255
                        fmt.Println(v4, v2, v1)
                        var v11 int = 42
                        var v12 int64 = -718
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    switch "foo bar" {
                    case "a":
                        var v13 int64 = 42
                        _ = v13
                    case "hello":
                        if ((true && ((42 == -2147483648) || true)) || false) {
                            var v14 int = fi8
                            for fi15 := 0; fi15 < 14; fi15++ {
                                switch (v2 + v2) {
                                case "a":
                                    var v16 float64 = (1.5 + 3.14159)
                                    _ = v16
                                default:
                                    v7 -= (v7 % ((18446744073709551615) | 1))
                                }
                                v17 := GenS1{F0: 100, F1: fi15, F2: 0}
                                var v18 int = 650
                                _ = fi15
                                _ = v17
                                _ = v18
                            }
                            _ = v14
                        }
                    default:
                        v4 += v4
                    }
                }
            }
            if len(v6) > 0 {
                v6[0] = 1
            }
            _ = fi8
        }
        v19 := GenS1{F0: 0, F1: -9223372036854775808, F2: -516}
        _ = v19
        _ = v7
    }
    v20 := GenS2{F0: (v1 & (^(v1)))}
    v2 = ((v2 + v2) + "")
    v4 -= v4
    v5 += (^(v5))
    var v21 uint32 = 767
    fmt.Println(v4, v2, v5)
    switch v2 {
    case "hello":
        fmt.Println(v5, v1, v5)
    case "foo bar":
        if true {
            switch -9223372036854775808 {
            case -951:
                if (0.1 > 0.1) {
                    v22 := GenS1{F0: 560, F1: 4294967295, F2: 201}
                    _ = v22
                }
            case -977:
                v23 := v3[uint32(v5)]
                _ = v23
            }
        }
    default:
        v24 := GenS1{F0: 342, F1: 128, F2: 229}
        _ = v24
    }
    var v25 uint16 = 57
    v26 := GenS3{F0: "hello"}
    var v27 uint16 = (v25 / (((v25 ^ v25)) | 1))
    var v28 uint32 = 624
    var v29 uint64 = (v1 * v1)
    if false {
        v30 := GenS2{F0: 61}
        for fi31 := 0; fi31 < 1; fi31++ {
            if false {
                switch (fi31 ^ 0) {
                case 32767:
                    for fi32 := 0; fi32 < 20; fi32++ {
                        var v33 int16 = -32768
                        v34 := GenS4{F0: (!(float32(float64(-2.25)) <= 0.1)), F1: 32767, F2: (false && false)}
                        var v35 uint64 = (v1 * 42)
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                default:
                    fmt.Println(v28, v25, v1)
                }
                var v36 int8 = 127
                _ = v36
            } else {
                if (-1 <= -128) {
                    var v37 uint8 = v4
                    var v38 int8 = 1
                    _ = v37
                    _ = v38
                }
            }
            var v39 int16 = 1000
            _ = fi31
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            v28 -= v28
            if (true || false) {
                var v41 int = fi40
                switch 32768 {
                case 4294967295:
                    v42 := GenS5{F0: (^(v4))}
                    _ = v42
                default:
                    v5 = (v5 - v5)
                }
                _ = v41
            } else {
                for fi43 := 0; fi43 < 9; fi43++ {
                    v4 += 42
                    _ = fi43
                }
                if len(v6) > 0 {
                    v6[0] = (v28 << ((uint(v25)) & 63))
                }
            }
            var v44 int = fi40
            _ = fi40
            _ = v44
        }
        _ = v30
    } else {
        if true {
            switch -981 {
            case 1:
                var v45 int32 = -1
                _ = v45
            case 128:
                fmt.Println(v1, v4)
            case 524:
                switch -192 {
                case 4294967295:
                    switch 9223372036854775807 {
                    case 127:
                        switch int(v28) {
                        case -2147483648:
                            var v46 uint32 = (v21 + (v28 % (((^(v28))) | 1)))
                            _ = v46
                        default:
                            v25 = v27
                        }
                    }
                case 128:
                    if ((-517 >= -2147483648) && true) {
                        var v47 bool = ((v5 % ((v5) | 1)) == (v5 - v5))
                        _ = v47
                    }
                }
            default:
                switch v2 {
                case "hello":
                    fmt.Println(v29, v29)
                case "":
                    _ = len(v6)
                case "foo bar":
                    var v48 string = v2
                    _ = v48
                default:
                    v27 = uint16(v1)
                }
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
            v4 += v4
            var v49 int32 = -86
            _ = v49
        }
        for fi50 := 0; fi50 < 13; fi50++ {
            if (-1 <= -351) {
                var v51 int8 = 42
                v1 = v1
                for fi52 := 0; fi52 < 7; fi52++ {
                    if true {
                        v29 -= v29
                        if (true || (true && (int16(fi52) != 1000))) {
                            var v53 uint16 = v25
                            v54 := GenS6{F0: v5}
                            v55 := GenS5{F0: 255}
                            _ = v53
                            _ = v54
                            _ = v55
                        } else {
                            var v56 uint8 = (v4 - 127)
                            var v57 float32 = (((0.1 - float32((float64(0.0) - -2.25))) / float32(((1e10 * -2.25) + 1.5))) + 1.5)
                            _ = v56
                            _ = v57
                        }
                    } else {
                        switch fi52 {
                        case -1:
                            var v58 uint = 18446744073709551615
                            _ = v58
                        }
                        v21 += uint32(v25)
                    }
                    _ = fi52
                }
                _ = v51
            }
            _ = fi50
        }
        fmt.Println(v25, v5, v25)
    }
    v28 = (^(v28))
    fmt.Println(v5, v27)
    v3[(v28 ^ v21)] = ((v1 | v29) == 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
