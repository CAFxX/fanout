package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{0, 455, 374}
    var v3 uint = 196
    var v4 uint16 = 452
    v4 = v4
    _ = len(v2)
    if true {
        var v5 int = 4294967295
        switch (v5 % (((^(v5))) | 1)) {
        case -788:
            var v6 uint16 = (^(v4))
            _ = v6
        case 145:
            var v7 int32 = -1
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            delete(v1, 885)
            v9 := v1[2147483648]
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 181}
        var v11 uint8 = 128
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 0}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v14 := GenS3{F0: (v3 ^ (v3 << ((uint(v4)) & 63))), F1: 225}
        fmt.Println(v3, v4)
        _ = fi13
        _ = v14
    }
    v15 := GenS4{F0: 1, F1: v4}
    if false {
        var v16 uint32 = 42
        v16 += (v16 * 2147483647)
        _ = v16
    } else {
        v4 = v4
        var v17 int16 = -32768
        v18 := GenS4{F0: 2, F1: 1}
        _ = v17
        _ = v18
    }
    if (false && false) {
        fmt.Println(v4)
        var v19 int64 = 9223372036854775807
        v20 := GenS4{F0: 2, F1: v4}
        _ = v19
        _ = v20
    }
    fmt.Println(v3)
    if true {
        if (4294967295 < 2) {
            if (0 <= int32(v3)) {
                v3 += 1
                v4 += (v4 - (v4 % ((v4) | 1)))
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 int = fi21
                v2 = append(v2, 1)
                v3 += v3
                _ = fi21
                _ = v22
            }
        }
        if true {
            v3 = 255
        }
        v23 := GenS4{F0: 4294967295, F1: (v4 * (v4 | v4))}
        _ = v23
    } else {
        fmt.Println(v4, v3)
        fmt.Println(v3, v3)
        _ = len(v2)
    }
    fmt.Println(v4, v3)
    v3 -= (v3 >> ((v3) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
