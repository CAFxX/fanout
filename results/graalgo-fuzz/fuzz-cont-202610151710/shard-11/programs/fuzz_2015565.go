package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{4294967295}
    v3 := make(map[uint32]int)
    var v4 int = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 += (fi5 & (fi5 % ((fi5) | 1)))
        _ = fi5
    }
    v6 := GenS1{F0: 2, F1: 343}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if true {
                switch "" {
                case "\n\t":
                    var v9 float64 = (-2.25 - 3.14159)
                    _ = v9
                }
                if (!false) {
                    v10 := GenS1{F0: 4294967295, F1: 880}
                    _ = v10
                }
                v4 -= (fi7 << ((42) & 63))
                if (!(-439 < 42)) {
                    switch "foo bar" {
                    case "foo bar":
                        v4 = v4
                    case "":
                        var v11 uint32 = 65535
                        _ = v11
                    default:
                        fmt.Println(fi7, fi8, v4)
                    }
                }
            } else {
                v12 := GenS2{F0: false, F1: 701}
                _ = v12
            }
            var v13 float32 = (((1e10 + ((-2.25 * -2.25) * ((float32((float64((1.5 + (-2.25 / (0.1 * 1.5)))) + 3.14159)) - ((1e10 - 1e10) / 0.1)) - ((3.14159 / float32(0.1)) - float32(1e10))))) - 3.14159) * 0.1)
            _ = fi8
            _ = v13
        }
        if true {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int16 = -1
                v4 = 32767
                _ = fi14
                _ = v15
            }
        }
        _ = fi7
    }
    v4 += v4
    v4 = v4
    switch (v4 / (((v4 << ((961) & 63))) | 1)) {
    case 1:
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v4 = (v4 / ((fi16) | 1))
                v18 := GenS3{F0: 2, F1: 0}
                switch "hello" {
                case "":
                    fmt.Println(fi17, fi16)
                case "a":
                    switch 4294967295 {
                    case -9223372036854775808:
                        switch 32768 {
                        case 1:
                            fmt.Println(fi17)
                        case 0:
                            switch ("" + "\n\t") {
                            case "a":
                                v1[626] = ""
                                v4 += -84
                            default:
                                fmt.Println(v4)
                            }
                        case -1:
                            v4 += fi17
                        }
                    default:
                        v19 := GenS3{F0: 165, F1: -222}
                        _ = v19
                    }
                case "\n\t":
                    v20 := GenS4{F0: ((("hello" + "\n\t") + ("foo bar" + ("a" + (("" + "") + ("a" + ("\n\t" + "")))))) + ""), F1: 1}
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    case -2147483648:
        v21 := GenS5{F0: 128}
        _ = v21
    case 2147483647:
        v1[2147483647] = (("" + (("foo bar" + "foo bar") + ("a" + ""))) + (("" + ((("hello" + "foo bar") + ("a" + "hello")) + "")) + (("foo bar" + "\n\t") + "")))
    default:
        _ = len(v2)
    }
    var v22 uint8 = 2
    v4 = (v4 | 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v6
}
