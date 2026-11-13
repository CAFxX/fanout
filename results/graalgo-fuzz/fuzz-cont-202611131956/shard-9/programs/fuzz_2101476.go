package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0() uint32 {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = -210
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    var v10 uint64 = 965
    _ = v10
    _ = v7
    return 2
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 65535
    var v3 float64 = 1e10
    var v4 int = 0
    v5 := []uint32{4294967295, uint32(v2), 0}
    var v6 uint8 = 42
    v4 = (v4 % ((4294967295) | 1))
    switch v1 {
    case -878:
        var v11 uint16 = (v2 / (((v2 ^ (v2 >> ((341) & 63)))) | 1))
        _ = v11
    }
    fmt.Println(v2, v6, v2)
    var v12 float32 = 0.1
    if false {
        v13 := GenS1{F0: -393, F1: int(v6)}
        v14 := GenS2{F0: false, F1: 0}
        v15 := GenS2{F0: true, F1: 0}
        _ = v13
        _ = v14
        _ = v15
    } else {
        fmt.Println(v6, v1)
    }
    v16 := GenS1{F0: int32(v1), F1: v1}
    if len(v5) > 0 {
        v5[0] = 276
    }
    var v17 int32 = -1
    switch "hello" {
    case "hello":
        fmt.Println(v17)
    case "foo bar":
        fmt.Println(v17, v6, v6)
    case "\n\t":
        for fi18 := 0; fi18 < 18; fi18++ {
            switch (v1 / (((^(fi18))) | 1)) {
            case 9223372036854775807:
                v19 := GenS2{F0: (!false), F1: (-(v17))}
                _ = v19
            case 182:
                v3 -= 1.5
            }
            if (!false) {
                v20 := GenS3{F0: 42, F1: v2, F2: (-(v17))}
                if (1 >= 42) {
                    if false {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            var v22 uint16 = v2
                            v2 += 950
                            var v23 int8 = 1
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        if (v3 > float64((((v12 / v12) + (v12 * v12)) / (1e10 / -2.25)))) {
                            var v24 string = ""
                            _ = v24
                        }
                        var v25 uint16 = (^(v2))
                        _ = v25
                    }
                }
                _ = v20
            } else {
                v2 -= 32767
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            fmt.Println(v3, v6)
            _ = fi18
        }
    }
    v12 -= ((0.1 + v12) / 3.14159)
    var v26 int32 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
