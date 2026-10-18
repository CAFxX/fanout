package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint = 65535
    v2 := []uint32{}
    var v3 uint = (v1 | (v1 - v1))
    v1 = 262
    if false {
        v3 = 42
    }
    v4 := GenS1{F0: 42, F1: 1, F2: 793}
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    switch 213 {
    case 524:
        v2 = append(v2, 0)
    case -9223372036854775808:
        var v5 bool = (42 == 65535)
        _ = v5
    default:
        fmt.Println(v1, v3)
    }
    v6 := GenS1{F0: 763, F1: -128, F2: 1}
    switch 2147483647 {
    case 32768:
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint32 = 2147483648
                switch (fi8 * (fi8 ^ (fi8 * (fi7 * (fi7 << ((v1) & 63)))))) {
                case 42:
                    if (true || (true || (42 < 581))) {
                        var v10 float32 = 0.1
                        if (!(-1 < -32768)) {
                            var v11 uint = 861
                            _ = v11
                        }
                        var v12 uint = (v3 - 4294967295)
                        switch "" {
                        case "foo bar":
                            if true {
                                if true {
                                    v13 := GenS2{F0: 1}
                                    _ = v13
                                } else {
                                    v14 := GenS1{F0: 629, F1: -1, F2: 1}
                                    _ = v14
                                }
                                if true {
                                    v12 -= (v3 ^ v12)
                                    fmt.Println(v12)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    v9 += (^(v9))
                                    v15 := GenS3{F0: false, F1: 255}
                                    _ = v15
                                }
                                v9 = (v9 >> ((v3) & 63))
                            } else {
                                var v16 int16 = -1
                                _ = v16
                            }
                        }
                        _ = v10
                        _ = v12
                    }
                case 0:
                    if (42 == 42) {
                        var v17 int32 = 814
                        _ = v17
                    }
                case -1:
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                }
                var v18 uint32 = v9
                _ = fi8
                _ = v18
                _ = v9
            }
            var v19 int16 = -873
            _ = fi7
            _ = v19
        }
    case 1:
        fmt.Println(v3, v1)
    }
    v1 = (^(v1))
    v3 += v3
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v3
    var v20 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
