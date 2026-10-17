package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 690
    var v2 uint32 = 818
    var v3 uint8 = 209
    v4 := []uint32{uint32(v1)}
    v5 := []uint32{v2}
    var v6 uint16 = 713
    var v7 uint = 1
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    v6 += 0
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: (v2 - uint32(v7)), F1: (v7 ^ uint(v2))}
        for fi10 := 0; fi10 < 7; fi10++ {
            v6 = (v6 | v6)
            v6 -= (^(v6))
            _ = fi10
        }
        v11 := GenS1{F0: v2, F1: (^(v7))}
        _ = fi8
        _ = v11
        _ = v9
    }
    v3 = 42
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            _ = len(v4)
            if len(v4) > 0 {
                v4[0] = v2
            }
            _ = fi13
        }
        fmt.Println(v2, v2, v6)
        var v14 int64 = 0
        _ = fi12
        _ = v14
    }
    v2 = (v2 - (v2 - (v2 % (((v2 & (v2 - v2))) | 1))))
    v7 = v7
    if len(v5) > 0 {
        v5[0] = (^(v2))
    }
    if false {
        v1 = v1
        v6 += v6
        v6 -= 370
    } else {
        var v15 uint32 = v2
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch (^(fi16)) {
        case 577:
            if false {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 int32 = 1
                    v6 = (v6 << ((v7) & 63))
                    if true {
                        var v19 uint = v7
                        v20 := GenS1{F0: v2, F1: (v19 | v7)}
                        v21 := GenS1{F0: v2, F1: v7}
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v22 := GenS2{F0: 42}
                        v23 := GenS1{F0: uint32(v3), F1: uint(v2)}
                        _ = v22
                        _ = v23
                    }
                    _ = fi17
                    _ = v18
                }
                v1 = 594
                v24 := GenS1{F0: 65535, F1: (v7 & v7)}
                _ = v24
            } else {
                if true {
                    var v25 uint16 = (v6 % (((^(v6))) | 1))
                    if (!(!(true || true))) {
                        var v26 uint64 = (v1 % (((v1 + (v1 << (((v7 >> ((v7) & 63))) & 63)))) | 1))
                        _ = v26
                    }
                    _ = v25
                } else {
                    v7 = v7
                    v4 = append(v4, (v2 | 0))
                    v7 += 255
                }
            }
        case 42:
            if len(v4) > 0 {
                v4[0] = (v2 + 42)
            }
        case 2147483647:
            var v27 int32 = int32(v6)
            _ = v27
        }
        v28 := GenS1{F0: v2, F1: (v7 | (v7 >> ((522) & 63)))}
        if ((v3 >> (((v7 >> (((v7 & uint(v2))) & 63))) & 63)) >= v3) {
            _ = len(v4)
            if (((3.14159 - 1.5) * ((-2.25 * ((0.0 + 0.0) - (0.1 / 3.14159))) * (-2.25 / 0.1))) != 1.5) {
                v1 = (v1 - v1)
            }
            v1 += uint64(v7)
        }
        _ = fi16
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
