package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32) int64 {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6, fi6}
        if len(v7) > 0 {
            v7[0] = (fi6 % ((fi6) | 1))
        }
        switch ("a" + ("a" + "")) {
        case "foo bar":
            if false {
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS1{F0: fi6, F1: (^(fi6))}
                    p0 += p0
                    _ = fi8
                    _ = v9
                }
            }
        case "hello":
            v7 = append(v7, (fi6 * fi6))
        }
        _ = fi6
        _ = v7
    }
    var v10 int8 = 58
    var v11 uint = 320
    v12 := GenS2{F0: 2, F1: -541, F2: "foo bar"}
    var v13 uint16 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return -149
}

func main() {
    var v1 int32 = -282
    v2 := make(map[uint32]bool)
    var v3 int32 = (v1 ^ (^(v1)))
    var v4 uint64 = 620
    var v5 int64 = 0
    var v14 int64 = (v5 % ((v5) | 1))
    v15 := GenS3{F0: 171}
    v5 += v5
    fmt.Println(v14)
    var v16 int64 = (v5 << ((2) & 63))
    v17 := v2[4294967295]
    v3 = v1
    var v18 uint32 = 916
    if true {
        v5 = v5
        var v19 int16 = -32768
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
