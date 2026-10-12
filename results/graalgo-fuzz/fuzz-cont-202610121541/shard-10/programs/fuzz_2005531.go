package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    v3 -= (v3 + 528)
    v3 += v3
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 1000, F1: "hello", F2: v3}
    v6 := v1[127]
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if (((4294967295 > 42) && true) && true) {
                v9 := GenS1{F0: 743, F1: "foo bar", F2: 0}
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v4)
                default:
                    v6 = (v6 + v6)
                }
                v10 := GenS1{F0: 42, F1: ((v6 + (v6 + ("foo bar" + v6))) + v6), F2: (v3 % (((v3 / ((163) | 1))) | 1))}
                _ = v10
                _ = v9
            }
            if false {
                v11 := GenS2{F0: (!(1 == 728))}
                v6 = v6
                fmt.Println(v4, v6)
                _ = v11
            } else {
                var v12 int = (-(fi7))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v13 := GenS1{F0: 950, F1: v6, F2: (v3 | v3)}
                _ = v12
                _ = v13
            }
            _ = fi8
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi7
    }
    fmt.Println(v4)
    v2 = append(v2, v3)
    v4 = v4
    v2 = append(v2, (v3 | v3))
    v4 = v4
    v6 = (v6 + "a")
    var v14 uint16 = 32767
    switch -860 {
    case 42:
        fmt.Println(v3)
    }
    var v15 uint64 = uint64(v14)
    fmt.Println(v15, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
