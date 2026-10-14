package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 uint32, p1 uint8) int {
    if true {
        v5 := GenS1{F0: -316}
        var v6 int64 = -9223372036854775808
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := []int{}
                p0 = 2
                v6 = v6
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                p0 = 4294967295
                _ = fi10
            }
            fmt.Println(fi7)
            _ = fi7
        }
        _ = v5
        _ = v6
    } else {
        var v11 string = ("\n\t" + "")
        _ = v11
    }
    v12 := GenS2{F0: -32768}
    _ = p0
    _ = p1
    _ = v12
    return 270
}

func main() {
    var v1 uint64 = 0
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: ("\n\t" + ""), F1: true, F2: 42}
    v4 = (v4 % (((v4 ^ (v4 % ((v4) | 1)))) | 1))
    var v16 int16 = 598
    v16 += (v16 % ((v16) | 1))
    v17 := v3[1]
    if (((0.0 / 1.0) - ((0.0 / (1.5 / -2.25)) / (0.1 * (1e10 / 0.1)))) > 1e10) {
        for fi18 := 0; fi18 < 15; fi18++ {
            if ((fi18 / (((fi18 >> ((2) & 63))) | 1)) > fi18) {
                delete(v3, 42)
            }
            _ = fi18
        }
    }
    v19 := GenS4{F0: (v4 & v4), F1: false}
    var v20 uint8 = 255
    fmt.Println(v20, v2, v20)
    v4 = v4
    switch -238 {
    case 32767:
        var v21 int64 = (^(v4))
        _ = v21
    case 127:
        v17 = (v17 + v17)
    case 42:
        var v22 float32 = 3.14159
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        if false {
            v2 = (^(v2))
            fmt.Println(v20)
            v24 := GenS2{F0: v16}
            _ = v24
        }
        _ = fi23
    }
    if ("a" >= v17) {
        _ = len(v3)
        v4 += -939
        v17 = ""
    } else {
        if true {
            v25 := GenS1{F0: (v16 / (((v16 + v16)) | 1))}
            if false {
                var v26 uint8 = v20
                switch v17 {
                case "a":
                    v27 := GenS1{F0: int16(v2)}
                    _ = v27
                default:
                    v2 = (v2 << ((269) & 63))
                }
                _ = v26
            }
            v2 = -1
            _ = v25
        } else {
            var v28 uint64 = (v1 << ((0) & 63))
            var v29 int = genFn0(42, (v20 >> ((569) & 63)))
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
