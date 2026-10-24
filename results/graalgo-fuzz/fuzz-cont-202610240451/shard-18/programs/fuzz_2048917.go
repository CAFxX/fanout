package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32) int32 {
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    var v4 uint8 = 2
    v5 := []int{890, -706}
    var v6 bool = (2147483647 >= 128)
    v6 = v6
    v6 = (316 >= -667)
    var v7 uint8 = v4
    v8 := GenS1{F0: 32768, F1: 18446744073709551615, F2: ("" + ("hello" + "a"))}
    v6 = (18446744073709551615 > 2)
    v2 += (v2 >> ((255) & 63))
    for fi9 := 0; fi9 < 19; fi9++ {
        if (v2 == 345) {
            v2 = (v2 * 42)
        }
        var v10 uint = 0
        v11 := GenS1{F0: 32768, F1: uint(v7), F2: "foo bar"}
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 -= 42
    _ = len(v3)
    v6 = (((0.0 + 0.1) * (1.5 + (1e10 + float32(1e10)))) >= ((((1e10 / ((3.14159 + (float32(0.1) / 1e10)) + float32(-2.25))) / ((((1e10 / 1e10) * 1.5) / 1e10) * -2.25)) + 0.0) / 3.14159))
    _ = len(v1)
    v7 += (v7 * v7)
    v12 := GenS1{F0: 42, F1: 494, F2: ("\n\t" + "\n\t")}
    fmt.Println(v6, v7)
    v13 := v1[65535]
    switch 32768 {
    case -2147483648:
        if v6 {
            for fi14 := 0; fi14 < 15; fi14++ {
                v2 = v2
                _ = fi14
            }
            _ = len(v1)
        } else {
            v15 := GenS2{F0: ((255 >= 42) || false), F1: 1, F2: 153}
            switch 0 {
            case 0:
                if ((v2 * 562) <= 255) {
                    fmt.Println(v13, v6)
                    v13 -= 2
                    if (!(0.1 == 1e10)) {
                        switch "\n\t" {
                        case "a":
                            if true {
                                switch 4294967295 {
                                case 1:
                                    var v16 int = 32767
                                    _ = v16
                                case 825:
                                    if v6 {
                                        for fi17 := 0; fi17 < 17; fi17++ {
                                            v18 := GenS2{F0: ((-1 >= 41) && (!true)), F1: (fi17 >> ((935) & 63)), F2: 1}
                                            v19 := GenS3{F0: -1, F1: (fi17 / (((fi17 + (fi17 & fi17))) | 1))}
                                            if len(v5) > 0 {
                                                v5[0] = fi17
                                            }
                                            _ = fi17
                                            _ = v18
                                            _ = v19
                                        }
                                        switch "a" {
                                        case "a":
                                            switch 9223372036854775807 {
                                            case 42:
                                                v20 := GenS3{F0: 16, F1: 1}
                                                _ = v20
                                            case 424:
                                                v13 = v13
                                            case -276:
                                                _ = len(v5)
                                            }
                                        case "\n\t":
                                            v21 := GenS2{F0: v6, F1: 42, F2: 18446744073709551615}
                                            _ = v21
                                        default:
                                            var v22 bool = ((false && (v7 <= v7)) && false)
                                            _ = v22
                                        }
                                        var v23 bool = (!(!v6))
                                        _ = v23
                                    }
                                default:
                                    v7 = v7
                                }
                                switch (("hello" + ("foo bar" + "foo bar")) + ("" + ("" + "hello"))) {
                                case "a":
                                    if v6 {
                                        v6 = (uint8(v13) != v4)
                                    }
                                default:
                                    var v24 uint64 = 757
                                    _ = v24
                                }
                                var v25 uint64 = 566
                                _ = v25
                            }
                        }
                    }
                }
            case 4294967295:
                v4 = v7
            default:
                v3[v13] = "a"
            }
            _ = v15
        }
    case 344:
        v26 := GenS2{F0: v6, F1: -9223372036854775808, F2: 0}
        _ = v26
    }
    v27 := GenS4{F0: v6}
    v28 := GenS4{F0: (!false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
