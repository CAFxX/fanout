package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 42
    var v3 uint64 = v2
    var v4 uint32 = 2147483647
    if (v3 != v3) {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 int16 = 42
            var v7 int8 = 127
            if ((0.0 * 1e10) != 0.1) {
                fmt.Println(v7, v6)
            } else {
                v8 := []int{fi5, 32767, 42}
                _ = v8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        v1 = v1
        var v9 int16 = 552
        _ = v9
    }
    v10 := []int{474}
    v4 = (^(v4))
    var v11 int32 = 0
    _ = len(v10)
    if false {
        var v12 uint32 = v4
        _ = v12
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            if true {
                if len(v10) > 0 {
                    v10[0] = 113
                }
                var v14 uint32 = 2
                _ = v14
            }
            _ = fi13
        }
    }
    if true {
        v15 := GenS1{F0: 65535, F1: "a", F2: 65535}
        switch int(v2) {
        case 4294967295:
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int32 = v11
                v4 = (v4 | (v4 << ((255) & 63)))
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint8 = 0
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
        case 0:
            if false {
                if true {
                    if false {
                        v20 := GenS2{F0: 763, F1: 42}
                        _ = v20
                    }
                }
                var v21 uint64 = 693
                var v22 int32 = v11
                _ = v21
                _ = v22
            }
        case 233:
            v4 += v4
        default:
            var v23 int64 = 0
            _ = v23
        }
        _ = v15
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v4)
    v1 += v1
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v4, v2, v3)
        _ = fi24
    }
    v4 += v4
    var v25 uint16 = 1
    v10 = append(v10, -1)
    v3 -= v3
    v26 := GenS3{F0: (("a" > "") && ((true && (1e10 <= 0.1)) && true))}
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
