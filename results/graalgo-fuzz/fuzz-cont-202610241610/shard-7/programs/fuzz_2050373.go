package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = (58 == -128)
    var v3 uint8 = 163
    v4 := []int{-9223372036854775808}
    for fi5 := 0; fi5 < 11; fi5++ {
        if (28 <= 921) {
            for fi6 := 0; fi6 < 20; fi6++ {
                v1 += (v1 + v1)
                var v7 int8 = 1
                _ = fi6
                _ = v7
            }
            var v8 uint8 = 2
            _ = v8
        }
        if (false && ((v1 >> ((v1) & 63)) != (v1 << ((1) & 63)))) {
            var v9 float32 = 0.0
            _ = v9
        } else {
            v1 += 0
            v3 -= 128
            if false {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS1{F0: (v3 | 127), F1: -176, F2: (fi5 % (((fi5 + (fi5 * (fi10 / (((fi10 & fi5)) | 1))))) | 1))}
                    v12 := GenS1{F0: v3, F1: 9223372036854775807, F2: int(v1)}
                    fmt.Println(v3, v2, v3)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v4 = append(v4, -9223372036854775808)
                v1 = v1
            }
        }
        v3 += (v3 & 42)
        _ = fi5
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v2 = v2
        _ = fi13
    }
    v4 = append(v4, -2147483648)
    var v14 float32 = -2.25
    if len(v4) > 0 {
        v4[0] = 113
    }
    if ((255 <= 562) || (true || (-2.25 >= 1e10))) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v15 uint32 = 2147483647
    if false {
        v2 = v2
        _ = len(v4)
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if ((2 != (v1 | 354)) || v2) {
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS1{F0: 177, F1: 42, F2: fi16}
                v19 := GenS2{F0: 881, F1: 100}
                _ = fi17
                _ = v18
                _ = v19
            }
            if (fi16 > fi16) {
                var v20 uint = v1
                _ = len(v4)
                v3 += (v3 | v3)
                _ = v20
            } else {
                for fi21 := 0; fi21 < 11; fi21++ {
                    switch "" {
                    case "foo bar":
                        v14 = -2.25
                    case "":
                        var v22 int = 127
                        _ = v22
                    }
                    _ = len(v4)
                    _ = fi21
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 int = fi16
                _ = fi23
                _ = v24
            }
        }
        _ = fi16
    }
    var v25 bool = true
    if (1 >= 32767) {
        v4 = append(v4, 0)
        v1 = v1
    } else {
        v4 = append(v4, 2147483647)
        v1 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
