package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 bool }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 419
    var v3 bool = ((!false) || (!(2 < 0)))
    var v4 int8 = 29
    v5 := []int{-549, 896, 9223372036854775807}
    var v6 string = "hello"
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: v6, F1: 42}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            v3 = (v6 > "foo bar")
            _ = fi10
        }
        _ = fi7
    }
    v3 = v3
    if ((0 > 4294967295) || (v6 == v6)) {
        var v11 string = v6
        v11 = (v11 + "\n\t")
        if len(v5) > 0 {
            v5[0] = 32767
        }
        _ = v11
    }
    var v12 int8 = v4
    var v13 int = 4294967295
    v2 -= 616
    var v14 int = (v13 ^ -300)
    var v15 int16 = 32767
    v2 += 42
    switch v6 {
    case "a":
        if len(v5) > 0 {
            v5[0] = (v13 * -2147483648)
        }
    case "\n\t":
        var v16 int64 = -224
        _ = v16
    case "":
        if v3 {
            for fi17 := 0; fi17 < 12; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v2 += (v2 | (v2 - v2))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 += v12
                    _ = fi18
                }
                v19 := GenS2{F0: v3, F1: 2147483647, F2: (((float32(3.14159) * float32(-2.25)) / 1.0) != 3.14159)}
                _ = fi17
                _ = v19
            }
        } else {
            v15 = v15
            v20 := GenS3{F0: 255, F1: 2, F2: 42}
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 16; fi21++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 = v2
            var v22 uint64 = 2
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v2, v14)
    var v23 float32 = float32(3.14159)
    v24 := GenS2{F0: v3, F1: 963, F2: v3}
    if (v12 <= v12) {
        v25 := v1[2]
        _ = v25
    } else {
        if len(v5) > 0 {
            v5[0] = (v14 << ((255) & 63))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
