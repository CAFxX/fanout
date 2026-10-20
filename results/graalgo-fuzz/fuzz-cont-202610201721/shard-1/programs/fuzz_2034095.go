package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8; F2 int32 }

func main() {
    var v1 int64 = 1
    var v2 int = -2147483648
    var v3 int64 = -73
    v4 := []uint32{2, 4294967295, 65535}
    v5 := make(map[int]int)
    if (1 < 10) {
        var v6 uint = 65535
        _ = len(v5)
        _ = v6
    }
    if true {
        var v7 uint64 = 18446744073709551615
        v8 := GenS1{F0: v2, F1: 65535, F2: int8(v1)}
        for fi9 := 0; fi9 < 4; fi9++ {
            v3 = v3
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    switch v2 {
    case 127:
        if (!(128 >= 11)) {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int8 = 127
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS2{F0: 156, F1: v3, F2: int32(fi10)}
                    if len(v4) > 0 {
                        v4[0] = 196
                    }
                    v14 := GenS1{F0: fi10, F1: 4294967295, F2: v11}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
            if false {
                v15 := GenS3{F0: 0}
                _ = v15
            }
            v16 := GenS1{F0: -1, F1: 659, F2: 127}
            _ = v16
        }
    case -111:
        v2 -= (v2 % ((-1) | 1))
    case 128:
        var v17 int16 = -929
        _ = v17
    default:
        v2 -= (v2 + v2)
    }
    v4 = append(v4, 42)
    v3 = (v3 - -9223372036854775808)
    switch (v2 >> ((4294967295) & 63)) {
    case 0:
        var v18 uint = 4294967295
        _ = v18
    }
    switch v2 {
    case -655:
        var v19 string = ""
        _ = v19
    case 2147483647:
        v5[v2] = v2
    case -513:
        fmt.Println(v3, v1, v1)
    }
    fmt.Println(v1, v2)
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int8 = -1
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 bool = (!(true && ((464 <= 165) && (!false))))
            if v23 {
                var v24 bool = (-1 <= -2147483648)
                v25 := GenS4{F0: "a", F1: 255}
                var v26 uint64 = 641
                _ = v24
                _ = v25
                _ = v26
            }
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v27 := GenS5{F0: 2147483647, F1: 0, F2: 42}
        _ = v27
    }
    var v28 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
