package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 float32 = 0.1
    var v2 uint = 65535
    v3 := []int{4294967295, 775, -1, -1}
    v4 := []uint32{262, 2147483648, 564, 2147483647}
    var v5 uint32 = 0
    var v6 float32 = v1
    switch -1 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint16 = 479
                var v10 int32 = -1
                if len(v4) > 0 {
                    v4[0] = v5
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            switch -98 {
            case -367:
                if ((("hello" + ((("" + (("" + "\n\t") + (("hello" + ("foo bar" + "")) + ("hello" + ((("a" + "hello") + "hello") + "\n\t"))))) + ("hello" + "\n\t")) + "hello")) > "foo bar") && (2 > 255)) {
                    if false {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            switch "a" {
                            case "hello":
                                if (uint8(v2) < 201) {
                                    if len(v3) > 0 {
                                        v3[0] = int(v2)
                                    }
                                }
                            case "":
                                var v12 uint8 = 128
                                _ = v12
                            case "foo bar":
                                v1 += ((1.5 - (v6 - v1)) * 0.0)
                            default:
                                v6 = (0.1 - float32(0.0))
                            }
                            _ = fi11
                        }
                    } else {
                        v6 = (v6 / v1)
                    }
                    switch "hello" {
                    case "":
                        v13 := GenS1{F0: 792, F1: 2147483647}
                        _ = v13
                    case "hello":
                        var v14 float64 = (1.5 + (((3.14159 / (3.14159 * 3.14159)) * float64(v6)) / 1e10))
                        _ = v14
                    }
                    var v15 int16 = 565
                    _ = v15
                }
            case 32768:
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 uint32 = (v5 << (((v2 - 255)) & 63))
                    var v18 uint32 = v5
                    v18 = 42
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case 407:
                switch fi7 {
                case 32768:
                    v4 = append(v4, v5)
                case 9223372036854775807:
                    v19 := GenS1{F0: 1, F1: 609}
                    _ = v19
                case 283:
                    v5 = v5
                default:
                    _ = len(v4)
                }
            }
            _ = fi7
        }
    default:
        v6 += float32(1e10)
    }
    v20 := GenS1{F0: 32767, F1: -544}
    var v21 uint = 781
    v6 = v6
    var v22 int8 = 1
    if len(v4) > 0 {
        v4[0] = 972
    }
    switch "foo bar" {
    case "foo bar":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 float32 = v1
            switch fi23 {
            case 32767:
                v25 := GenS1{F0: 300, F1: -2147483648}
                _ = v25
            default:
                if (uint64(v2) == 12) {
                    if ((!true) || true) {
                        var v26 uint32 = v5
                        for fi27 := 0; fi27 < 5; fi27++ {
                            switch fi27 {
                            case -912:
                                v24 += (((v6 + v1) + float32((-2.25 - ((1.5 + (((1e10 - 3.14159) / 0.1) + -2.25)) * 0.0)))) + (((v24 / 1.5) / 1.0) - (v24 * v24)))
                            }
                            _ = fi27
                        }
                        _ = v26
                    }
                }
            }
            _ = fi23
            _ = v24
        }
    case "\n\t":
        var v28 uint8 = 42
        _ = v28
    default:
        var v29 string = ("" + (("foo bar" + ("\n\t" + ((("" + "\n\t") + "a") + "foo bar"))) + "a"))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
