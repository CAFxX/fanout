package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }

func main() {
    var v1 bool = (!(65535 > 4294967295))
    var v2 uint8 = 133
    var v3 uint = 431
    v4 := []int{2147483647, -9223372036854775808, 2147483647, 294}
    var v5 int32 = 2147483647
    v6 := GenS1{F0: 18446744073709551615, F1: v2}
    v7 := GenS2{F0: (((("foo bar" > (("hello" + "a") + "a")) && (!((v1 || (32767 > 1)) || ((!(926 >= -352)) || ((4294967295 < 4294967295) && v1))))) && ((!(true || v1)) && v1)) || v1)}
    v4 = append(v4, 4294967295)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint8 = (v2 ^ v2)
            _ = fi8
            _ = v9
        }
    case "a":
        var v10 uint64 = 18446744073709551615
        _ = v10
    case "hello":
        if (v1 || (!(-775 > -950))) {
            v4 = append(v4, -85)
            v4 = append(v4, 32768)
        }
    default:
        if len(v4) > 0 {
            v4[0] = 907
        }
    }
    var v11 float32 = 3.14159
    v12 := GenS1{F0: 2, F1: v2}
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int32 = v5
        if (-9223372036854775808 == 42) {
            switch (fi13 & (fi13 ^ fi13)) {
            case 127:
                var v15 uint16 = 686
                _ = v15
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    switch (fi16 << (((^(v3))) & 63)) {
                    case 1:
                        switch ("\n\t" + "foo bar") {
                        case "":
                            v17 := GenS3{F0: 0, F1: v5, F2: int64(fi16)}
                            _ = v17
                        case "a":
                            if (!false) {
                                if len(v4) > 0 {
                                    v4[0] = (^(fi16))
                                }
                                if len(v4) > 0 {
                                    v4[0] = int(v2)
                                }
                                v5 = (v14 << ((v3) & 63))
                            } else {
                                if (!v1) {
                                    v1 = (2147483648 != 0)
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi16
                                    }
                                    if true {
                                        if ((1 <= 9223372036854775807) || v1) {
                                            var v18 uint32 = 4294967295
                                            var v19 float64 = 1e10
                                            _ = v18
                                            _ = v19
                                        }
                                    }
                                    fmt.Println(v2, v2, v14)
                                }
                            }
                        }
                    }
                    var v20 uint16 = 255
                    v1 = ((167 >= 65535) && v1)
                    _ = fi16
                    _ = v20
                }
            }
        } else {
            var v21 uint64 = 4294967295
            _ = v21
        }
        var v22 int32 = (v14 / (((^(v14))) | 1))
        _ = fi13
        _ = v14
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
