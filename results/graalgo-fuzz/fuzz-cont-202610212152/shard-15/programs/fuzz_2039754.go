package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint = 42
    v2 := []string{"foo bar"}
    var v3 int32 = 444
    for fi4 := 0; fi4 < 8; fi4++ {
        v1 -= v1
        fmt.Println(fi4, fi4, v3)
        _ = fi4
    }
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
        }
        if (("a" < "\n\t") || true) {
            for fi7 := 0; fi7 < 17; fi7++ {
                v3 += v3
                _ = fi7
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 string = "hello"
                v2 = append(v2, v9)
                switch "foo bar" {
                case "hello":
                    v1 = v1
                    v1 = (v1 ^ (v1 >> ((772) & 63)))
                    v1 -= v1
                case "foo bar":
                    v1 += v1
                case "a":
                    var v10 uint32 = 4294967295
                    _ = v10
                default:
                    v11 := GenS1{F0: 477, F1: 0}
                    var v12 string = v9
                    _ = v11
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                _ = len(v2)
                _ = fi13
            }
        } else {
            if (!true) {
                v3 += (v3 & (v3 % ((v3) | 1)))
                fmt.Println(v1, v1)
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint = (v1 / (((v1 / (((v1 & (v1 * v1))) | 1))) | 1))
                    _ = fi14
                    _ = v15
                }
            }
            fmt.Println(fi5, v1)
            _ = len(v2)
        }
        _ = fi5
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint8 = 33
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: "hello", F2: v3}
    var v19 uint8 = 127
    if (v1 >= (v1 - v1)) {
        v19 = (v19 >> (((^(v1))) & 63))
        for fi20 := 0; fi20 < 19; fi20++ {
            switch (^(fi20)) {
            case 128:
                switch "foo bar" {
                case "":
                    v21 := GenS3{F0: -889}
                    _ = v21
                default:
                    v19 += v19
                }
            case -9223372036854775808:
                v22 := GenS2{F0: -128, F1: "foo bar", F2: v3}
                _ = v22
            default:
                _ = len(v2)
            }
            if false {
                if true {
                    v1 -= 111
                    switch 0 {
                    case 235:
                        switch 32767 {
                        case 32768:
                            switch fi20 {
                            case 577:
                                v1 = 0
                            default:
                                fmt.Println(v1, fi20, v1)
                            }
                        case 127:
                            var v23 uint64 = 0
                            _ = v23
                        default:
                            fmt.Println(fi20, v19, v3)
                        }
                    case 1:
                        var v24 float64 = -2.25
                        _ = v24
                    case 32767:
                        v1 -= (^(v1))
                    default:
                        v19 = v19
                    }
                }
            }
            _ = fi20
        }
    }
    fmt.Println(v1)
    _ = len(v2)
    v25 := GenS1{F0: 65535, F1: 65535}
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
