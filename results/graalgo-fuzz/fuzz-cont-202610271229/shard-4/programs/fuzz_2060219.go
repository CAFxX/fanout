package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint; F2 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint16 }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 int32 = -1
    var v3 uint64 = 123
    var v4 uint = 0
    v2 = (v2 | (-(v2)))
    var v5 int64 = -9223372036854775808
    v6 := GenS1{F0: ((!(!(255 < v4))) && false)}
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (^(fi7)) {
        case 32768:
            v3 -= v3
        case 679:
            var v8 int = (fi7 / ((-1) | 1))
            _ = v8
        default:
            v9 := GenS1{F0: true}
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: ((-1 >= 1) || ((^(v3)) >= v3))}
    var v11 int32 = v2
    v12 := GenS2{F0: 1, F1: 42, F2: -754}
    fmt.Println(v3)
    if (v11 <= -247) {
        fmt.Println(v5, v5, v4)
    }
    v4 += 425
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            if (("\n\t" + "\n\t") <= "foo bar") {
                fmt.Println(v11, fi13, v4)
                v1 = append(v1, fi13)
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint32 = 380
                _ = fi14
                _ = v15
            }
            var v16 int = fi13
            _ = v16
        } else {
            for fi17 := 0; fi17 < 13; fi17++ {
                if len(v1) > 0 {
                    v1[0] = (fi13 >> ((v4) & 63))
                }
                _ = fi17
            }
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            if false {
                var v18 uint32 = 42
                _ = v18
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    switch fi19 {
                    case 42:
                        fmt.Println(v2, v11, fi19)
                    case 629:
                        var v20 int32 = (v2 | (v11 % ((619) | 1)))
                        _ = v20
                    case 322:
                        switch fi13 {
                        case 1:
                            var v21 uint8 = 127
                            _ = v21
                        case 9223372036854775807:
                            if false {
                                v1 = append(v1, 127)
                                v3 -= v3
                                var v22 uint32 = 2
                                _ = len(v1)
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (fi13 | (fi19 * fi13))
                                    }
                                }
                                _ = v22
                            }
                        default:
                            switch "hello" {
                            case "foo bar":
                                v23 := GenS1{F0: true}
                                _ = v23
                            }
                        }
                    }
                    v24 := GenS3{F0: 0, F1: 896, F2: 65535}
                    _ = fi19
                    _ = v24
                }
            }
        }
        v11 = v11
        fmt.Println(v5, v11, v3)
        _ = fi13
    }
    if ((!((-(v5)) == (-(v5)))) && false) {
        for fi25 := 0; fi25 < 10; fi25++ {
            if (!((v11 >> ((v4) & 63)) >= (v2 << (((v4 * v4)) & 63)))) {
                var v26 uint = 464
                _ = len(v1)
                _ = v26
            } else {
                v27 := GenS1{F0: (v2 > 219)}
                _ = v27
            }
            _ = fi25
        }
        var v28 float32 = float32(float64(3.14159))
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 uint = (^(v4))
            var v31 float64 = -2.25
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    if (2147483647 <= 487) {
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(v4, fi32, v5)
            _ = fi32
        }
    } else {
        v2 = (v11 & (v11 - v11))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
