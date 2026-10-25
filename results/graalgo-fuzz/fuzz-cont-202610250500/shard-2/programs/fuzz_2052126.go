package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 int16) bool {
    p0 -= p0
    v7 := GenS1{F0: 2}
    var v8 string = ((("\n\t" + "foo bar") + ("foo bar" + (("\n\t" + "hello") + "a"))) + ("" + "\n\t"))
    v8 = v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    var v3 int64 = -1
    var v4 int16 = 107
    var v5 int16 = v4
    v6 := make(map[int]string)
    if (920 == 4294967295) {
        delete(v6, 0)
    }
    var v9 uint32 = 844
    var v10 uint8 = 1
    var v11 int = 127
    v12 := GenS2{F0: 127, F1: (v10 ^ v10)}
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v3, v11, v9)
        _ = fi13
    }
    v14 := GenS3{F0: true, F1: v9, F2: 255}
    v5 += v4
    v11 -= v11
    if ((v1 | 127) == 190) {
        var v15 uint32 = v9
        if (false && (!("foo bar" != "foo bar"))) {
            if (1 == 663) {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 uint16 = 1
                    _ = fi16
                    _ = v17
                }
            } else {
                v5 -= (v5 << ((0) & 63))
                v18 := GenS1{F0: (v9 & v15)}
                _ = v18
            }
            v19 := GenS4{F0: 18446744073709551615, F1: -1, F2: v1}
            _ = v19
        }
        _ = v15
    }
    var v20 int = (v11 ^ v11)
    switch v11 {
    case -9223372036854775808:
        fmt.Println(v5, v10)
    case 1:
        var v21 int8 = -128
        _ = v21
    }
    v6[(v11 | v20)] = ""
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
