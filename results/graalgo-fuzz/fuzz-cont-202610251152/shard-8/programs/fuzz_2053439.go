package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    v3 := []int{32767, -996, 2147483647, 4294967295}
    v2 -= -23
    v4 := GenS1{F0: 128}
    v5 := GenS2{F0: 0, F1: 181}
    v6 := GenS1{F0: 128}
    v2 = (v2 << ((65535) & 63))
    v2 = (v2 << ((665) & 63))
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 1}
        _ = fi7
        _ = v8
    }
    switch "" {
    case "":
        if (255 != 1) {
            fmt.Println(v2)
            v9 := GenS1{F0: 106}
            var v10 int32 = 1
            _ = v10
            _ = v9
        }
    case "foo bar":
        fmt.Println(v2)
    }
    if (false || false) {
        v11 := GenS3{F0: 32767}
        var v12 uint8 = 42
        if false {
            var v13 string = ""
            if (377 >= 2147483648) {
                _ = len(v1)
                var v14 int8 = v2
                v13 = (v13 + (v13 + v13))
                _ = v14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                v12 += 127
                if (((((fi15 / ((fi15) | 1)) <= (fi15 * fi15)) || false) && true) || true) {
                    switch fi15 {
                    case 0:
                        v16 := v1[4294967295]
                        _ = v16
                    case -367:
                        if true {
                            v12 -= v12
                            for fi17 := 0; fi17 < 13; fi17++ {
                                v2 = (v2 & (v2 - int8(fi17)))
                                delete(v1, uint32(v12))
                                var v18 string = "hello"
                                _ = fi17
                                _ = v18
                            }
                        } else {
                            var v19 int32 = -693
                            v2 -= v2
                            _ = v19
                        }
                    case -9223372036854775808:
                        v13 = (v13 + "\n\t")
                    }
                } else {
                    var v20 bool = ((v2 | v2) < 86)
                    v21 := GenS4{F0: -2147483648, F1: (-(fi15)), F2: 2147483647}
                    v20 = false
                    _ = v20
                    _ = v21
                }
                v22 := v1[2]
                _ = fi15
                _ = v22
            }
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        var v23 uint32 = 848
        if (-1 < 0) {
            var v24 float32 = 0.1
            var v25 int16 = 1
            v26 := GenS2{F0: 4294967295, F1: 280}
            _ = v24
            _ = v25
            _ = v26
        }
        v27 := GenS4{F0: 514, F1: -1, F2: (v23 ^ (v23 - (v23 << ((0) & 63))))}
        _ = v23
        _ = v27
    }
    v28 := GenS4{F0: 0, F1: -1, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
