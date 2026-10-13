package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 4294967295
    var v3 string = "a"
    if (v3 >= v3) {
        v3 = v3
    } else {
        v2 += v2
        var v4 uint = 700
        _ = v4
    }
    var v5 bool = (0 != -383)
    v6 := []int{32767, 2147483647}
    fmt.Println(v1, v1)
    v6 = append(v6, 360)
    v5 = (!v5)
    fmt.Println(v5, v5)
    var v7 uint64 = (v2 & (^(v2)))
    fmt.Println(v2)
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 2, F1: v5, F2: 516}
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10)
            if len(v6) > 0 {
                v6[0] = -908
            }
            var v11 uint8 = 127
            var v12 bool = ((v11 << ((2) & 63)) != (v11 | 0))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v2 = v2
    _ = len(v6)
    var v13 bool = v5
    switch 790 {
    case 32768:
        v14 := GenS2{F0: 255, F1: 2, F2: -332}
        _ = v14
    case -1:
        var v15 uint8 = 128
        _ = v15
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v3 = ""
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            if (v1 < float64(1e10)) {
                v1 = (float64(3.14159) * v1)
                if (((v3 + v3) + (("hello" + v3) + ("a" + (v3 + v3)))) < "foo bar") {
                    v3 = v3
                    var v18 int64 = 749
                    switch (^(fi17)) {
                    case 2147483647:
                        var v19 bool = (127 <= 1)
                        _ = v19
                    case 442:
                        fmt.Println(v5)
                    case -883:
                        v20 := GenS2{F0: 437, F1: 1, F2: 2147483647}
                        _ = v20
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v21 := GenS2{F0: 43, F1: 30, F2: -919}
                    _ = v18
                    _ = v21
                } else {
                    if v13 {
                        switch 0 {
                        case 470:
                            if ((32767 <= 32767) && v13) {
                                v7 -= (v7 >> ((255) & 63))
                                switch (v3 + ("a" + v3)) {
                                case "":
                                    var v22 uint8 = uint8(v7)
                                    _ = v22
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                                v23 := GenS1{F0: 42, F1: (true && ((1e10 - 1.5) > -2.25)), F2: 512}
                                fmt.Println(v13, fi17)
                                _ = v23
                            }
                        default:
                            switch 1 {
                            case -9223372036854775808:
                                v6 = append(v6, (fi17 | (-(fi16))))
                            case 42:
                                v3 = v3
                            case 32767:
                                v24 := GenS3{F0: 288, F1: 127, F2: 0}
                                _ = v24
                            }
                        }
                        fmt.Println(v13, v13, fi17)
                        v3 = v3
                    }
                }
            } else {
                v25 := GenS4{F0: 0, F1: 32767, F2: 255}
                switch v3 {
                case "foo bar":
                    v6 = append(v6, 127)
                case "\n\t":
                    v13 = true
                }
                _ = v25
            }
            _ = fi17
        }
        _ = fi16
    }
    v3 = "a"
    var v26 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
