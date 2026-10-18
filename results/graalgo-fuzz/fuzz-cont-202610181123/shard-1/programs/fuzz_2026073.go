package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 bool, p1 uint8, p2 uint32) uint16 {
    p2 -= p2
    p2 += p2
    p2 = (^(p2))
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 string = "foo bar"
        v7 := []int{fi5, (fi5 << ((768) & 63)), (fi5 - fi5), (fi5 - (fi5 ^ fi5))}
        var v8 uint8 = (p1 * p1)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 235
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = 1000
    v3 := []int{-1, 0, 32768, -883}
    v4 := make(map[int]bool)
    v9 := GenS1{F0: false}
    v10 := GenS2{F0: 359, F1: (("hello" + "a") + ""), F2: v1}
    if true {
        v2 = v2
        fmt.Println(v2, v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 := GenS1{F0: true}
    if false {
        v12 := GenS1{F0: (false && (1.5 > (1.5 * (0.0 * (1e10 / float64(float32(float64(1e10))))))))}
        _ = v12
    } else {
        v3 = append(v3, int(v2))
        var v13 uint32 = 2147483647
        v1 = v1
        _ = v13
    }
    var v14 string = "foo bar"
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            v2 = (v2 >> ((1) & 63))
            var v16 int16 = v2
            _ = v16
        } else {
            v1 -= 127
        }
        var v17 int64 = 7
        _ = fi15
        _ = v17
    }
    v1 = (v1 << ((18446744073709551615) & 63))
    v2 = -32768
    v1 = v1
    var v18 int16 = v2
    var v19 float32 = -2.25
    fmt.Println(v1)
    v4[32768] = (42 < (v2 >> ((229) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
