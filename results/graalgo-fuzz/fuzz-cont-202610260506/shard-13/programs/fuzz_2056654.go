package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 int64 = -1
    var v2 string = "\n\t"
    var v3 uint64 = 295
    v4 := []uint32{2147483647, 468, 68, 412}
    v5 := make(map[int]int)
    v1 += (v1 >> ((1) & 63))
    v5[2147483647] = 4294967295
    var v6 float32 = float32((0.0 * 1.5))
    if false {
        if len(v4) > 0 {
            v4[0] = 1
        }
        var v7 float32 = v6
        _ = v7
    } else {
        fmt.Println(v6, v2, v2)
        delete(v5, 610)
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            _ = len(v5)
            _ = fi9
        }
        fmt.Println(v6, v2, v3)
        _ = fi8
    }
    if (-294 < -1) {
        fmt.Println(v3, v3, v6)
    } else {
        var v10 uint32 = 42
        _ = v10
    }
    v2 = v2
    var v11 int = 127
    v12 := GenS1{F0: 965, F1: int32(v11), F2: 17}
    var v13 uint8 = 155
    if (32767 <= 32767) {
        var v14 uint16 = 65535
        var v15 uint8 = uint8(v3)
        _ = len(v4)
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: (v3 / (((^(v3))) | 1)), F1: 1, F2: int32(v13)}
    if (!false) {
        for fi17 := 0; fi17 < 2; fi17++ {
            if len(v4) > 0 {
                v4[0] = 353
            }
            fmt.Println(v11)
            _ = fi17
        }
    }
    v4 = append(v4, 1)
    switch v2 {
    case "foo bar":
        switch v11 {
        case -9223372036854775808:
            if false {
                switch v2 {
                case "a":
                    var v18 uint32 = 1
                    _ = v18
                case "":
                    if (true && true) {
                        if ((65535 < 65535) && true) {
                            var v19 int8 = -10
                            _ = v19
                        } else {
                            if false {
                                v11 = 2147483647
                                v20 := GenS3{F0: v3, F1: 994}
                                _ = v20
                            }
                        }
                    }
                case "hello":
                    delete(v5, (v11 - -2147483648))
                }
            }
        case -2147483648:
            if ((v3 << ((4294967295) & 63)) < (v3 % (((^(v3))) | 1))) {
                v21 := v5[-382]
                var v22 bool = false
                _ = v21
                _ = v22
            }
        case 4294967295:
            v23 := GenS1{F0: 1, F1: 2147483647, F2: -128}
            _ = v23
        default:
            var v24 int64 = 523
            _ = v24
        }
    case "":
        for fi25 := 0; fi25 < 8; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                switch (fi26 & 128) {
                case 9223372036854775807:
                    v2 = (v2 + v2)
                default:
                    var v27 int = v11
                    _ = v27
                }
                var v28 int16 = 32767
                _ = fi26
                _ = v28
            }
            _ = fi25
        }
    case "hello":
        fmt.Println(v6)
    default:
        v11 = v11
    }
    v11 += v11
    var v29 int64 = (v1 - (v1 >> ((uint(v13)) & 63)))
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := GenS1{F0: 18446744073709551615, F1: -1, F2: 127}
        for fi32 := 0; fi32 < 14; fi32++ {
            switch (v2 + v2) {
            case "hello":
                switch 32767 {
                case 128:
                    var v33 int64 = int64(v13)
                    _ = v33
                case -2147483648:
                    v2 = ("\n\t" + v2)
                }
                v3 += v3
            default:
                var v34 int32 = 0
                _ = v34
            }
            v35 := GenS2{F0: 89, F1: 463, F2: -2147483648}
            if (v3 >= (v3 + v3)) {
                var v36 uint8 = 255
                v13 = v13
                _ = v36
            }
            _ = fi32
            _ = v35
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
