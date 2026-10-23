package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint64, p2 uint64) uint32 {
    if ((42 != -278) || (9223372036854775807 != -529)) {
        for fi5 := 0; fi5 < 13; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 int64 = 889
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]int)
    var v3 int64 = 1
    v4 := []uint32{65535, 2, 837}
    v1 = v1
    _ = len(v4)
    switch 575 {
    case 303:
        v4 = append(v4, genFn0(1, (^(v1)), (^(v1))))
    case -9223372036854775808:
        for fi8 := 0; fi8 < 6; fi8++ {
            v3 -= (v3 % (((^(v3))) | 1))
            var v9 uint8 = 84
            _ = fi8
            _ = v9
        }
    }
    if true {
        for fi10 := 0; fi10 < 15; fi10++ {
            switch 127 {
            case -9223372036854775808:
                if (v3 != v3) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        switch fi11 {
                        case -2147483648:
                            v3 -= (v3 + (v3 & -54))
                        default:
                            if (1 <= -108) {
                                var v12 uint32 = 42
                                _ = v12
                            }
                        }
                        _ = fi11
                    }
                    for fi13 := 0; fi13 < 14; fi13++ {
                        _ = len(v4)
                        _ = fi13
                    }
                }
            case -2147483648:
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 uint16 = 1
                    _ = fi14
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                v4 = append(v4, 598)
                if true {
                    var v17 int64 = (v3 & v3)
                    v3 = (v3 / (((^(v17))) | 1))
                    if (false && (((false && false) && (true && (true || false))) && false)) {
                        var v18 uint8 = 241
                        var v19 uint16 = 0
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                v3 -= v3
                _ = fi16
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi10
        }
        v20 := v2[("" + "\n\t")]
        _ = v20
    } else {
        var v21 string = "a"
        v3 = (v3 << ((1) & 63))
        var v22 uint64 = v1
        _ = v21
        _ = v22
    }
    v1 = v1
    if ((!(false || (-2.25 >= 3.14159))) || false) {
        switch ("" + "foo bar") {
        case "\n\t":
            fmt.Println(v3, v3)
        case "hello":
            var v23 uint32 = 869
            _ = v23
        case "":
            v24 := GenS1{F0: 17, F1: 0, F2: (v3 << ((1) & 63))}
            _ = v24
        }
    } else {
        var v25 uint32 = 0
        var v26 int32 = 375
        if len(v4) > 0 {
            v4[0] = 319
        }
        _ = v25
        _ = v26
    }
    v1 = 4294967295
    v27 := GenS2{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
