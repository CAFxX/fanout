package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int32 = -904
    var v2 bool = (!false)
    v3 := []int{9223372036854775807, 32767}
    var v4 string = "hello"
    v5 := make(map[string]string)
    v1 -= v1
    var v6 int = 490
    v2 = (!v2)
    v7 := GenS1{F0: -102}
    if len(v3) > 0 {
        v3[0] = v6
    }
    v1 -= -1
    var v8 int32 = (-(v1))
    v3 = append(v3, 128)
    if (622 != 32767) {
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                if (v2 || (!((-(fi9)) >= (fi10 ^ (-(fi10)))))) {
                    var v11 int32 = v1
                    fmt.Println(v2)
                    v3 = append(v3, -9223372036854775808)
                    _ = v11
                }
                v12 := GenS1{F0: 1}
                v13 := v5[("foo bar" + "hello")]
                _ = fi10
                _ = v12
                _ = v13
            }
            _ = fi9
        }
        switch 178 {
        case -948:
            v14 := GenS1{F0: 100}
            _ = v14
        case -2147483648:
            v2 = v2
        case -7:
            if (65535 == 65535) {
                v6 += v6
                for fi15 := 0; fi15 < 10; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint8 = 2
                        v17 -= v17
                        if (0 == 4294967295) {
                            if len(v3) > 0 {
                                v3[0] = 814
                            }
                            v18 := GenS1{F0: int8(v1)}
                            var v19 int32 = v8
                            v1 -= -2147483648
                            v20 := GenS1{F0: 100}
                            v21 := v5[(v4 + "foo bar")]
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            if (2147483648 > 2147483647) {
                                delete(v5, ("\n\t" + v4))
                            }
                            fmt.Println(v4)
                            _ = len(v3)
                        }
                        v22 := GenS1{F0: -1}
                        fmt.Println(v8, v1)
                        var v23 uint64 = 42
                        _ = fi16
                        _ = v17
                        _ = v22
                        _ = v23
                    }
                    _ = fi15
                }
                var v24 uint16 = 0
                _ = v24
            }
        default:
            var v25 float64 = (1e10 + float64(-2.25))
            _ = v25
        }
        v3 = append(v3, 808)
    } else {
        if false {
            if ((((!(!v2)) && (!v2)) && (v2 || ((3.14159 + 1.5) != 3.14159))) && true) {
                fmt.Println(v4)
            } else {
                v8 = (v8 | v8)
                v8 -= (v8 - (v8 & v8))
            }
            v26 := GenS2{F0: 0, F1: 2, F2: 2147483647}
            switch (v6 >> ((535) & 63)) {
            case -1:
                var v27 int64 = 0
                _ = v27
            case 2147483647:
                v28 := GenS3{F0: 219, F1: 358}
                _ = v28
            default:
                v6 -= (v6 >> ((327) & 63))
            }
            _ = v26
        } else {
            v6 = (v6 << ((18446744073709551615) & 63))
            var v29 uint64 = 399
            _ = v29
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                v1 = -959
                fmt.Println(fi31, v4)
                _ = fi31
            }
            _ = fi30
        }
        v8 += (v8 / (((v8 % ((v1) | 1))) | 1))
    }
    fmt.Println(v6, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
