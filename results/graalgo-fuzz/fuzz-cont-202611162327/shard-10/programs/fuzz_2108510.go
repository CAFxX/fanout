package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 int32 = 0
    var v2 int16 = -45
    var v3 uint = 65535
    var v4 int = 2147483647
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: (v4 - 32767), F1: -1, F2: 0}
    var v7 int16 = (v2 - (v2 & (v2 | v2)))
    v2 += v2
    v2 = v7
    fmt.Println(v4, v3)
    v2 = (v2 << ((472) & 63))
    var v8 uint16 = 249
    if (false || true) {
        v9 := GenS2{F0: -124, F1: 32767}
        var v10 int64 = 9223372036854775807
        _ = v10
        _ = v9
    } else {
        var v11 uint32 = 0
        if true {
            v12 := GenS1{F0: v4, F1: 670, F2: (v3 & 42)}
            if true {
                fmt.Println(v3, v1, v8)
                v13 := []int{v4, (v4 | int(v1))}
                if true {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = v13
            } else {
                var v14 int16 = v7
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint64 = 2
                    v16 -= v16
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS3{F0: true, F1: v1}
                        switch fi15 {
                        case 202:
                            delete(v5, (^(v11)))
                        }
                        v1 = int32(v16)
                        v5[v11] = ""
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
                var v19 int64 = -1
                _ = v14
                _ = v19
            }
            v20 := GenS2{F0: 0, F1: v4}
            _ = v12
            _ = v20
        } else {
            fmt.Println(v3, v4, v3)
            switch int(v3) {
            case 32767:
                var v21 int8 = -32
                _ = v21
            case 128:
                for fi22 := 0; fi22 < 1; fi22++ {
                    if ((1.5 - (3.14159 / 1.0)) > ((((-2.25 / 0.1) * ((float32(1.5) + -2.25) - 1e10)) * (1.5 / float32(1.5))) + (1e10 / 0.1))) {
                        v23 := []int{fi22, v4, 2147483647, int(v3)}
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v11 += v11
                            fmt.Println(v7, v1, v11)
                            _ = fi24
                        }
                        _ = v23
                    }
                    _ = fi22
                }
            default:
                for fi25 := 0; fi25 < 13; fi25++ {
                    v1 -= (v1 % (((v1 >> (((^(v3))) & 63))) | 1))
                    _ = fi25
                }
            }
        }
        _ = v11
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        fmt.Println(v7)
        for fi27 := 0; fi27 < 20; fi27++ {
            delete(v5, 1)
            _ = fi27
        }
        v1 = -864
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
