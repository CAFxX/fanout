package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16; F1 int; F2 uint }

func main() {
    v1 := []int{2147483647, 128, 2147483647, 2147483647}
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 float32 = float32((1.5 / -2.25))
    if true {
        v4 += -2.25
        var v5 uint32 = 4294967295
        _ = v5
    } else {
        v6 := GenS1{F0: 514, F1: false, F2: 2}
        _ = v6
    }
    if (-87 < -128) {
        v2[4294967295] = -589
    } else {
        switch 2147483647 {
        case 1:
            fmt.Println(v4)
        default:
            if false {
                _ = len(v1)
                switch "\n\t" {
                case "a":
                    var v7 uint = 390
                    _ = v7
                default:
                    v4 = 3.14159
                }
                fmt.Println(v4)
            } else {
                var v8 bool = (9223372036854775807 != 64)
                v1 = append(v1, 351)
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                delete(v2, 2147483648)
                if (32767 > 4294967295) {
                    v3[(("" + ("a" + "")) + "foo bar")] = (fi9 % ((128) | 1))
                }
                _ = fi10
            }
            v11 := GenS1{F0: 2147483648, F1: (0 > -128), F2: 0}
            _ = fi9
            _ = v11
        }
        var v12 int8 = 1
        _ = v12
    }
    v4 = 3.14159
    fmt.Println(v4)
    _ = len(v1)
    if (-683 > 1) {
        if (642 != 42) {
            if (-2.25 <= 1e10) {
                v13 := GenS2{F0: 32767, F1: 0, F2: 610}
                fmt.Println(v4)
                var v14 uint32 = 583
                _ = v13
                _ = v14
            }
            var v15 uint64 = 2
            _ = v15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            if ((4294967295 > 0) || true) {
                v17 := v3["hello"]
                var v18 int = v17
                _ = v17
                _ = v18
            } else {
                var v19 uint64 = 2
                if true {
                    if (379 <= 0) {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            var v21 float32 = v4
                            if true {
                                v19 -= v19
                                v1 = append(v1, (fi20 << ((0) & 63)))
                            }
                            v19 -= 274
                            _ = fi20
                            _ = v21
                        }
                        v19 = v19
                        v4 += 1e10
                    } else {
                        if (!false) {
                            var v22 int64 = 9223372036854775807
                            v22 -= v22
                            _ = v22
                        } else {
                            if len(v1) > 0 {
                                v1[0] = fi16
                            }
                            v19 = 4294967295
                            var v23 uint = 4294967295
                            _ = v23
                        }
                    }
                }
                v4 = v4
                _ = v19
            }
            _ = len(v1)
            switch "a" {
            case "a":
                var v24 uint16 = 528
                _ = v24
            }
            _ = fi16
        }
    }
    var v25 uint16 = 1
    v26 := GenS2{F0: 1, F1: 42, F2: 18446744073709551615}
    v27 := GenS1{F0: uint32(v25), F1: ((v25 << ((255) & 63)) == 42), F2: 2}
    if false {
        var v28 uint64 = 2
        _ = v28
    }
    v29 := GenS1{F0: 448, F1: true, F2: 39}
    var v30 uint8 = 42
    if true {
        v30 = (v30 % ((127) | 1))
        for fi31 := 0; fi31 < 8; fi31++ {
            v30 = (^(v30))
            _ = len(v1)
            v1 = append(v1, (^(fi31)))
            _ = fi31
        }
        v30 = (v30 * (v30 + (v30 / ((v30) | 1))))
    }
    switch "\n\t" {
    case "":
        var v32 uint = 18446744073709551615
        _ = v32
    default:
        v25 = (v25 >> ((65535) & 63))
    }
    v25 = (v25 & (v25 << ((42) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
