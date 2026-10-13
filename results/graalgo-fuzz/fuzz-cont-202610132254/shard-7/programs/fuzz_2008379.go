package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint; F2 int8 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint32 }

func main() {
    var v1 int64 = -557
    var v2 bool = ("hello" < "\n\t")
    v3 := []uint32{992, 211}
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 32767, F1: -2147483648, F2: 2147483648}
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS2{F0: 65535}
            if v2 {
                var v8 string = "hello"
                switch -37 {
                case 42:
                    if false {
                        if (603 >= 501) {
                            fmt.Println(fi4, v8)
                            switch (fi6 / (((fi6 * fi4)) | 1)) {
                            case 9223372036854775807:
                                fmt.Println(v1, fi4, fi6)
                            case 928:
                                v9 := GenS3{F0: 65535, F1: (4294967295 > 18446744073709551615), F2: 952}
                                _ = v9
                            case 32768:
                                fmt.Println(v2, v8, v2)
                            default:
                                v10 := GenS4{F0: 0}
                                _ = v10
                            }
                            if len(v3) > 0 {
                                v3[0] = 488
                            }
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                        }
                        switch (fi4 << ((0) & 63)) {
                        case 9223372036854775807:
                            fmt.Println(v2, fi6, v1)
                            var v11 int32 = 42
                            _ = v11
                        }
                    } else {
                        v2 = v2
                        v12 := GenS4{F0: 42}
                        _ = v12
                    }
                default:
                    var v13 bool = true
                    _ = v13
                }
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                var v14 uint = 2
                _ = v14
                _ = v8
            }
            v15 := GenS2{F0: 1}
            _ = fi6
            _ = v15
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v16 := GenS2{F0: 255}
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int8 = 42
        var v19 int = (fi17 + (fi17 - 127))
        _ = fi17
        _ = v18
        _ = v19
    }
    v1 += 222
    v1 -= v1
    var v20 bool = (true || v2)
    v21 := GenS5{F0: 18446744073709551615, F1: 668, F2: -47}
    _ = len(v3)
    v22 := GenS6{F0: 248, F1: 100, F2: 257}
    v1 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
