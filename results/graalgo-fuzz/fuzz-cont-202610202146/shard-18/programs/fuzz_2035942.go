package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 string) uint {
    if (!true) {
        var v4 int = (p0 >> ((4294967295) & 63))
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: 2147483648, F1: -128, F2: 149}
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    v7 := GenS1{F0: 528, F1: -6, F2: 4294967295}
    v8 := GenS2{F0: 9223372036854775807}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 924
}

func main() {
    var v1 uint = 659
    var v2 int64 = 0
    v3 := []uint32{2147483648}
    v2 = v2
    fmt.Println(v2, v1)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    _ = len(v3)
    v1 -= v1
    fmt.Println(v1, v2)
    v9 := GenS2{F0: (v2 / (((-(v2))) | 1))}
    var v10 int8 = 0
    if (!(1 == (v2 * -291))) {
        for fi11 := 0; fi11 < 12; fi11++ {
            v10 += v10
            for fi12 := 0; fi12 < 11; fi12++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi12
            }
            _ = fi11
        }
        var v13 int = -9223372036854775808
        _ = v13
    } else {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    if (!false) {
        if (v10 <= 108) {
            v1 = (^(v1))
            v1 += v1
            switch 725 {
            case 2147483647:
                for fi14 := 0; fi14 < 14; fi14++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi14
                }
            case -660:
                if true {
                    if (false || false) {
                        fmt.Println(v1, v1)
                        var v15 int64 = (^(v2))
                        v15 -= v2
                        _ = v15
                    } else {
                        var v16 float32 = 0.1
                        _ = len(v3)
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 uint32 = 2147483647
                        v10 = 0
                        v3 = append(v3, (v18 ^ 901))
                        _ = fi17
                        _ = v18
                    }
                    v10 = (v10 + -128)
                }
            case -454:
                v1 = v1
            }
        }
        v10 += (v10 * 1)
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            if (false && false) {
                v20 := GenS2{F0: -64}
                _ = v20
            }
            switch fi19 {
            case 1:
                if true {
                    v2 = int64(v10)
                    v21 := GenS3{F0: v2, F1: 0}
                    _ = v21
                } else {
                    if len(v3) > 0 {
                        v3[0] = 557
                    }
                    switch (-(fi19)) {
                    case -1:
                        fmt.Println(v10)
                    case 32767:
                        switch ("\n\t" + (("a" + (("a" + "foo bar") + ("\n\t" + ""))) + "\n\t")) {
                        case "\n\t":
                            var v22 uint8 = 42
                            _ = v22
                        default:
                            var v23 int8 = v10
                            _ = v23
                        }
                    }
                }
            case -1:
                switch 9223372036854775807 {
                case -415:
                    v24 := GenS3{F0: (v2 / ((v2) | 1)), F1: 32768}
                    _ = v24
                case 128:
                    v10 -= v10
                default:
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v10 += v10
                        fmt.Println(v10)
                        _ = fi25
                    }
                }
            default:
                v3 = append(v3, 1)
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS2{F0: (v2 ^ 42)}
                _ = fi26
                _ = v27
            }
            _ = fi19
        }
    }
    var v28 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v9
}
