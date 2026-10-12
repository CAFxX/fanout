package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    v1 := []uint32{772, 2, 0, 2147483648}
    var v2 int = 127
    var v3 uint8 = 210
    v4 := []uint32{1}
    if (false || (84 == 0)) {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 uint8 = 0
            switch (v2 - 931) {
            case 176:
                v2 += (v2 - (v2 / ((v2) | 1)))
            default:
                for fi7 := 0; fi7 < 5; fi7++ {
                    if ((1.5 / 0.1) < 1e10) {
                        switch "hello" {
                        case "hello":
                            v8 := GenS1{F0: 2147483647}
                            _ = v8
                        default:
                            var v9 uint = 519
                            _ = v9
                        }
                        v3 += (v3 + (^(v6)))
                        if ((v6 | v3) != 42) {
                            switch fi5 {
                            case -9223372036854775808:
                                v3 += v6
                                if false {
                                    v2 = fi7
                                    v6 = v3
                                    v2 -= (v2 << ((42) & 63))
                                }
                            case 32767:
                                v10 := GenS1{F0: 0}
                                _ = v10
                            default:
                                var v11 string = (("" + ("" + "foo bar")) + "\n\t")
                                _ = v11
                            }
                            v3 = (v6 >> ((65535) & 63))
                            v2 = (fi5 / ((-476) | 1))
                        } else {
                            var v12 uint8 = (v6 * 0)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v1 = append(v1, 4294967295)
                            v1 = append(v1, 2)
                            _ = v12
                        }
                        if (1 < 9223372036854775807) {
                            v3 -= v6
                            v6 -= 0
                        } else {
                            v2 = fi7
                            var v13 int16 = 249
                            if ((1 > 65) && true) {
                                v6 -= (v6 % (((^(v3))) | 1))
                            } else {
                                v13 = v13
                                v3 -= 0
                            }
                            _ = v13
                        }
                        _ = len(v4)
                    }
                    _ = fi7
                }
            }
            switch "\n\t" {
            case "foo bar":
                v3 = 128
            default:
                fmt.Println(v2)
            }
            _ = fi5
            _ = v6
        }
    }
    v14 := GenS1{F0: -2147483648}
    _ = len(v4)
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = (^(v3))
        _ = fi15
    }
    v16 := GenS2{F0: 203, F1: 4294967295}
    fmt.Println(v2)
    v1 = append(v1, 2147483647)
    v3 = (v3 ^ v3)
    if true {
        _ = len(v1)
        if len(v4) > 0 {
            v4[0] = 331
        }
        v17 := GenS2{F0: 255, F1: 65535}
        _ = v17
    } else {
        if (!((0 <= 0) || (0 > 255))) {
            _ = len(v1)
            var v18 int16 = 426
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v2)
        _ = fi19
    }
    v20 := GenS1{F0: 695}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
