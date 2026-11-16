package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint64 }

func main() {
    var v1 int32 = -340
    var v2 int = -143
    var v3 uint16 = 2
    v4 := []string{}
    var v5 int32 = (v1 - v1)
    if (false && (!false)) {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: fi6, F1: (^(v5)), F2: -1}
            if false {
                var v8 uint64 = 164
                if false {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v1 += v5
                        var v10 float64 = 1e10
                        var v11 int = (v2 + (fi6 % ((791) | 1)))
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := GenS1{F0: fi12, F1: (^(v1)), F2: (v5 % (((-(v5))) | 1))}
                        if true {
                            var v14 int32 = -748
                            v2 -= (fi6 + fi12)
                            v4 = append(v4, (("a" + (("hello" + "a") + ("a" + (("" + "a") + ("\n\t" + "foo bar"))))) + ""))
                            _ = v14
                        }
                        _ = fi12
                        _ = v13
                    }
                } else {
                    _ = len(v4)
                    v3 = (v3 % (((^(v3))) | 1))
                }
                _ = v8
            } else {
                if (((v3 * v3) == v3) || true) {
                    if false {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            fmt.Println(fi15)
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v16 float32 = 3.14159
                                _ = v16
                            case -2147483648:
                                var v17 uint8 = 1
                                _ = v17
                            case 9223372036854775807:
                                v2 = (fi15 - -9223372036854775808)
                            default:
                                var v18 float64 = 1e10
                                _ = v18
                            }
                            _ = fi15
                        }
                        var v19 int64 = 329
                        _ = v19
                    }
                    v4 = append(v4, "\n\t")
                }
            }
            _ = fi6
            _ = v7
        }
    } else {
        switch (v2 % ((v2) | 1)) {
        case 32767:
            v3 -= (v3 >> ((42) & 63))
        case -9223372036854775808:
            v20 := GenS2{F0: -9223372036854775808, F1: ((true || true) || (((^(v2)) <= v2) && ((!true) || (-32768 < -32768)))), F2: 42}
            _ = v20
        default:
            var v21 int8 = -108
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        switch "" {
        case "hello":
            for fi23 := 0; fi23 < 2; fi23++ {
                v2 += int(v1)
                _ = fi23
            }
        case "":
            var v24 uint8 = 0
            _ = v24
        case "\n\t":
            v4 = append(v4, "foo bar")
        default:
            fmt.Println(v2, v2)
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        v4 = append(v4, "a")
        _ = fi25
    }
    v1 = 0
    for fi26 := 0; fi26 < 2; fi26++ {
        v3 = (v3 | v3)
        _ = fi26
    }
    v2 = (v2 + (v2 - 167))
    v4 = append(v4, "hello")
    for fi27 := 0; fi27 < 3; fi27++ {
        v2 += fi27
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
