package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }

func main() {
    var v1 uint64 = 899
    var v2 int32 = -604
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    var v6 uint32 = 332
    v2 = 42
    if true {
        v7 := []int{644}
        var v8 uint32 = 0
        fmt.Println(v2, v2, v5)
        _ = v7
        _ = v8
    }
    v6 += (v6 + v6)
    v9 := []int{}
    v2 = (v2 << ((400) & 63))
    var v10 uint64 = 4294967295
    v3[(v6 >> ((uint(v6)) & 63))] = (true && ((("foo bar" + ("foo bar" + ("a" + (("foo bar" + "\n\t") + ("" + "foo bar"))))) + "") >= "\n\t"))
    switch 42 {
    case 0:
        v11 := GenS1{F0: 9223372036854775807}
        _ = v11
    default:
        v12 := GenS1{F0: 32768}
        _ = v12
    }
    v13 := v4[v6]
    if (!false) {
        v9 = append(v9, 588)
        if false {
            v5 = v5
            if (255 >= 42) {
                v14 := GenS1{F0: 0}
                switch v13 {
                case "foo bar":
                    v10 -= 18446744073709551615
                case "hello":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        _ = len(v9)
                        _ = fi15
                    }
                }
                _ = v14
            } else {
                v2 += (v2 << ((18446744073709551615) & 63))
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 int64 = v5
                v5 -= v5
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS2{F0: 86, F1: int8(fi16), F2: true}
                    v4[(v6 & (v6 ^ uint32(v1)))] = ("a" + v13)
                    fmt.Println(fi18, fi16, v1)
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
        } else {
            var v20 string = (("foo bar" + v13) + (v13 + v13))
            _ = v20
        }
    }
    if (128 <= 42) {
        v21 := v3[v6]
        if ((true && v21) && (v2 >= (v2 << ((4294967295) & 63)))) {
            fmt.Println(v21, v2, v10)
            var v22 int = -1
            var v23 int64 = int64(v2)
            _ = v22
            _ = v23
        }
        var v24 int = 0
        _ = v21
        _ = v24
    } else {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v25 int64 = v5
        _ = v25
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
