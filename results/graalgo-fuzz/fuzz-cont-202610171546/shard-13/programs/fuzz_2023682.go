package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []uint32{1, 2147483648}
    v2 := make(map[uint32]int)
    var v3 int = 127
    v4 := make(map[int]string)
    var v5 int = v3
    v6 := []int{(v3 - 652), 32768, v5, v5}
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(v5, v3, fi7)
            _ = fi7
        }
        v3 = (v5 - v3)
    }
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v3, v5)
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v3)
        default:
            switch v3 {
            case -474:
                v10 := v2[0]
                _ = v10
            case -2147483648:
                if len(v6) > 0 {
                    v6[0] = v5
                }
            case 32768:
                v2[307] = v5
            default:
                v5 -= (v3 / ((v5) | 1))
            }
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            v5 += 2147483647
            _ = fi11
        }
        _ = fi9
    }
    v3 += (v5 | v5)
    var v12 int16 = 164
    delete(v2, 2147483647)
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v3)
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        _ = fi13
    }
    switch (v3 % ((-9223372036854775808) | 1)) {
    case 42:
        for fi14 := 0; fi14 < 9; fi14++ {
            v3 -= (^(fi14))
            if false {
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(fi14)
                    v5 += 1
                    _ = fi15
                }
                v3 = -692
            }
            _ = fi14
        }
    case 346:
        fmt.Println(v3, v12, v3)
    case -9223372036854775808:
        v16 := v2[0]
        _ = v16
    }
    if (!(false || true)) {
        for fi17 := 0; fi17 < 9; fi17++ {
            v12 = (v12 % ((0) | 1))
            _ = fi17
        }
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS1{F0: 42}
            v20 := GenS2{F0: 903}
            var v21 uint32 = 656
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        switch (^(v3)) {
        case 1:
            v22 := GenS2{F0: 0}
            _ = v22
        case 4294967295:
            for fi23 := 0; fi23 < 2; fi23++ {
                for fi24 := 0; fi24 < 19; fi24++ {
                    _ = len(v1)
                    v25 := GenS3{F0: v12, F1: 308}
                    v26 := GenS4{F0: 947}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        }
    }
    v3 += (v3 | 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
