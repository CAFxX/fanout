package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 0
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            fmt.Println(v2, v4)
        } else {
            var v7 float64 = (3.14159 - 1.5)
            v4 = (v4 + v4)
            _ = v7
        }
        var v8 uint8 = (v4 ^ (v4 & 37))
        _ = fi6
        _ = v8
    }
    var v9 int = v2
    v9 -= -104
    var v10 float32 = 3.14159
    v11 := GenS1{F0: "hello", F1: 154, F2: (v4 % (((^(v4))) | 1))}
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v2, v2)
        fmt.Println(v9)
        for fi13 := 0; fi13 < 20; fi13++ {
            if false {
                if (v10 >= v10) {
                    v10 = (0.1 - float32(0.0))
                }
            }
            v2 -= fi13
            _ = fi13
        }
        _ = fi12
    }
    var v14 int16 = -1
    switch v9 {
    case -610:
        var v15 int = v2
        _ = v15
    case 467:
        switch 127 {
        case 2147483647:
            var v16 uint16 = 158
            _ = v16
        case 32768:
            switch 9223372036854775807 {
            case 184:
                var v17 int64 = 0
                _ = v17
            default:
                v18 := GenS1{F0: ("" + "foo bar"), F1: 2, F2: 129}
                _ = v18
            }
        case -1:
            var v19 int32 = -505
            _ = v19
        default:
            if (!(!((float64(v10) / (((1.5 - float64(v10)) - 1e10) * ((float64((v10 / 3.14159)) / (-2.25 * (0.0 + 0.1))) / 3.14159))) >= 0.1))) {
                switch "\n\t" {
                case "foo bar":
                    v20 := GenS1{F0: "", F1: (^(v4)), F2: (v4 << ((575) & 63))}
                    _ = v20
                default:
                    v10 = (v10 - -2.25)
                }
                switch "hello" {
                case "a":
                    var v21 int8 = 100
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v23 := GenS1{F0: ((("\n\t" + "hello") + ("hello" + "\n\t")) + (("a" + (("" + ("\n\t" + ("a" + ("hello" + ("foo bar" + "\n\t"))))) + "foo bar")) + ("\n\t" + "a"))), F1: v4, F2: 54}
                        if false {
                            var v24 uint8 = 1
                            _ = v24
                        } else {
                            v25 := GenS1{F0: "hello", F1: v4, F2: (v4 * 64)}
                            var v26 uint = 1
                            _ = v25
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            } else {
                v27 := GenS1{F0: "", F1: v4, F2: (^(v4))}
                _ = v27
            }
        }
    case 32768:
        fmt.Println(v9, v2)
    default:
        var v28 float64 = (((0.1 + (float64(v10) / 0.1)) + 3.14159) - 3.14159)
        _ = v28
    }
    v4 += 245
    v14 -= -891
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
