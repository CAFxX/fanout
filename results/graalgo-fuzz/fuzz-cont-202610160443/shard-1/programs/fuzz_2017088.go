package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 uint64; F2 int16 }
func genFn0(p0 uint16, p1 uint32) uint {
    p1 = 2
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint32 = (p1 | p1)
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return 639
}

func main() {
    var v1 int16 = 1
    var v2 uint = 1
    v3 := make(map[uint32]string)
    var v6 string = ("hello" + ("\n\t" + "hello"))
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v1, fi7)
        var v8 int32 = 153
        if (false && true) {
            var v9 uint = (v2 ^ (v2 << ((v2) & 63)))
            _ = v9
        } else {
            v10 := []int{int(v8), -1}
            if (0.1 >= 1e10) {
                v11 := GenS1{F0: 0, F1: 2, F2: ("" + v6)}
                if false {
                    v1 = (v1 / (((v1 - (v1 & 0))) | 1))
                    var v12 string = (v6 + "\n\t")
                    _ = v12
                } else {
                    fmt.Println(v1, v8, v8)
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := v3[305]
                    _ = fi13
                    _ = v14
                }
                _ = v11
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 int32 = (v8 >> (((v2 / ((1) | 1))) & 63))
                    v17 := GenS1{F0: 165, F1: 797, F2: ("a" + (v6 + (((v6 + v6) + "foo bar") + (v6 + v6))))}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v2 = (v2 ^ (v2 >> ((genFn0(1, 580)) & 63)))
    v18 := GenS2{F0: v6, F1: 2, F2: (v1 << (((v2 >> ((v2) & 63))) & 63))}
    v19 := []int{-123}
    var v20 int16 = 32767
    v2 -= (v2 << (((^(v2))) & 63))
    v2 -= (v2 / (((v2 % ((v2) | 1))) | 1))
    switch v6 {
    case "hello":
        v20 = (v1 - -1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
