package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint32 = 65535
    var v2 float64 = 0.1
    var v3 uint64 = 610
    v4 := []string{("" + "foo bar"), ""}
    var v5 uint16 = 742
    var v6 uint = 457
    v7 := GenS1{F0: (v5 + 255), F1: ((42 != 622) && (!(!false))), F2: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = ("foo bar" + "hello")
    }
    _ = len(v4)
    v3 = (v3 ^ uint64(v6))
    v8 := GenS1{F0: 2, F1: true, F2: v3}
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v3, v1)
        var v10 int32 = -373
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(fi11, v5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch (fi11 - fi11) {
        case 32768:
            v5 = 910
        default:
            if true {
                if (v1 <= 490) {
                    var v12 float32 = ((0.1 / 1.0) + -2.25)
                    var v13 int16 = 1
                    _ = v12
                    _ = v13
                }
                fmt.Println(v6)
                switch (-(fi11)) {
                case 9223372036854775807:
                    v5 = (v5 % ((v5) | 1))
                default:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        fmt.Println(v2)
                        _ = fi14
                    }
                }
            } else {
                if (0 == 0) {
                    v1 = (v1 - v1)
                }
                switch fi11 {
                case 2147483647:
                    v15 := GenS1{F0: (v5 | (v5 * (v5 / ((42) | 1)))), F1: (!true), F2: 42}
                    _ = v15
                case 32767:
                    switch -891 {
                    case 128:
                        if false {
                            var v16 uint8 = 231
                            _ = v16
                        }
                    case 4294967295:
                        if (!((-(fi11)) != -1)) {
                            v17 := GenS2{F0: 18446744073709551615}
                            v6 -= (v6 ^ v6)
                            v3 -= (v3 % ((v3) | 1))
                            _ = v17
                        } else {
                            switch int(v5) {
                            case 615:
                                var v18 string = (("foo bar" + "") + "\n\t")
                                _ = v18
                            case 277:
                                fmt.Println(fi11)
                            }
                            v19 := GenS2{F0: v6}
                            _ = v19
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            var v20 uint16 = (v5 | v5)
                            _ = v20
                        case "\n\t":
                            v21 := GenS2{F0: 18446744073709551615}
                            _ = v21
                        default:
                            switch 1 {
                            case 42:
                                v1 = v1
                            default:
                                v22 := GenS1{F0: (v5 % (((v5 * 264)) | 1)), F1: true, F2: (^(v3))}
                                _ = v22
                            }
                        }
                    }
                case 167:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    var v23 int64 = -9223372036854775808
                    _ = v23
                }
                var v24 int8 = -1
                _ = v24
            }
        }
        _ = fi11
    }
    fmt.Println(v1, v6, v3)
    for fi25 := 0; fi25 < 19; fi25++ {
        v2 -= (0.0 * v2)
        _ = fi25
    }
    v6 += v6
    v2 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
