package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]string)
    var v3 uint8 = 252
    v4 := []uint32{v1}
    if true {
        for fi5 := 0; fi5 < 13; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                switch "\n\t" {
                case "foo bar":
                    v1 = uint32(v3)
                case "hello":
                    v7 := GenS1{F0: 2, F1: v3, F2: 2}
                    _ = v7
                default:
                    var v8 bool = ((v3 - 1) != (v3 % ((v3) | 1)))
                    _ = v8
                }
                v2[(v1 >> ((573) & 63))] = (("a" + "a") + ("hello" + (("a" + "") + "a")))
                var v9 int16 = -108
                _ = fi6
                _ = v9
            }
            _ = fi5
        }
    } else {
        switch 1 {
        case 147:
            v1 = (v1 & (v1 - v1))
        case 721:
            switch 946 {
            case 693:
                if true {
                    v1 += v1
                    v1 = 688
                    switch 1 {
                    case -2147483648:
                        v3 += (v3 / ((v3) | 1))
                    case 1:
                        v3 = (v3 - 2)
                    }
                }
            case 1:
                var v10 uint16 = 351
                _ = v10
            default:
                fmt.Println(v3, v1)
            }
        case -1:
            v3 -= (v3 << ((926) & 63))
        default:
            if false {
                v1 = v1
                v3 = (v3 << ((314) & 63))
                v11 := GenS1{F0: 677, F1: (v3 * v3), F2: 866}
                _ = v11
            }
        }
        switch ("hello" + "a") {
        case "hello":
            var v12 uint64 = 497
            _ = v12
        case "\n\t":
            v13 := GenS1{F0: 1, F1: 2, F2: 80}
            _ = v13
        case "":
            v3 = (v3 << ((65535) & 63))
        }
        fmt.Println(v1, v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!(32767 == -9223372036854775808)) {
        if ((true && (-2147483648 == 1)) || (true && false)) {
            _ = len(v4)
            var v14 uint32 = (v1 % (((^(v1))) | 1))
            _ = v14
        }
        switch -9223372036854775808 {
        case 808:
            for fi15 := 0; fi15 < 17; fi15++ {
                switch fi15 {
                case 0:
                    fmt.Println(v1)
                case 1:
                    _ = len(v2)
                case 128:
                    v1 = 566
                default:
                    v16 := GenS2{F0: 127, F1: (fi15 / (((-(fi15))) | 1)), F2: (v3 | 255)}
                    _ = v16
                }
                v17 := v2[(v1 / (((v1 & v1)) | 1))]
                _ = fi15
                _ = v17
            }
        default:
            if (!true) {
                fmt.Println(v3, v1)
                var v18 float64 = (((float64(-2.25) * 0.0) / 1.5) - 0.0)
                v19 := GenS2{F0: 42, F1: 0, F2: v3}
                _ = v18
                _ = v19
            } else {
                v2[(v1 ^ v1)] = "\n\t"
                fmt.Println(v3, v1)
                v1 += 4294967295
            }
        }
        var v20 int = 184
        _ = v20
    }
    var v21 uint32 = 2147483648
    v1 = v21
    fmt.Println(v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS3{F0: 1, F1: 0, F2: (^(v1))}
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
