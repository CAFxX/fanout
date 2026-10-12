package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint = 2
    var v2 int = 2147483647
    v3 := []string{"hello", ("\n\t" + "hello"), "foo bar"}
    v4 := []string{}
    v5 := []string{"a"}
    v2 = (-(v2))
    v2 = (v2 >> ((v1) & 63))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v6 int8 = -100
    fmt.Println(v1)
    var v7 float32 = (float32(((1e10 + (3.14159 * 0.1)) / -2.25)) + 0.1)
    var v8 int64 = 117
    var v9 uint = (v1 % (((v1 - 4294967295)) | 1))
    v2 = v2
    if (2147483647 == 2147483647) {
        v5 = append(v5, (("hello" + ("" + "")) + ("" + "foo bar")))
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: 1, F1: 1}
            v8 += (^(v8))
            _ = len(v4)
            _ = fi10
            _ = v11
        }
    }
    if (!false) {
        if (42 <= 173) {
            fmt.Println(v7, v2)
            if ((1e10 + 1e10) == (3.14159 / 1.0)) {
                switch (v2 / (((v2 >> ((255) & 63))) | 1)) {
                case -1:
                    v2 = (v2 * v2)
                }
            }
            if (false || true) {
                for fi12 := 0; fi12 < 12; fi12++ {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 uint16 = 984
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                fmt.Println(v8, v6, v8)
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v15 := GenS1{F0: -1, F1: 42}
            _ = v15
        }
        var v16 int64 = v8
        _ = v16
    } else {
        fmt.Println(v6)
        for fi17 := 0; fi17 < 1; fi17++ {
            if ((!((((v6 - v6) <= 100) || (int32(v6) > 563)) || (-2147483648 < -2147483648))) || (1e10 >= (1.5 + 3.14159))) {
                var v18 uint8 = 84
                _ = v18
            } else {
                var v19 uint64 = 2
                switch -747 {
                case -188:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 uint = v1
                        v7 += v7
                        var v22 int64 = 236
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case 9223372036854775807:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v1 = 4294967295
                        _ = fi23
                    }
                case 4294967295:
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v9 = v1
                        _ = fi24
                    }
                default:
                    for fi25 := 0; fi25 < 13; fi25++ {
                        v6 -= 1
                        _ = fi25
                    }
                }
                _ = v19
            }
            v26 := GenS1{F0: 42, F1: 1000}
            v7 = v7
            _ = fi17
            _ = v26
        }
        v2 = (v2 + v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        v7 = 1.5
        _ = fi27
    }
    v9 = v9
    var v28 uint = v9
    v2 -= 9223372036854775807
    fmt.Println(v6, v28, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
