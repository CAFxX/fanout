package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -1
    var v3 int16 = -1
    var v4 uint32 = 4294967295
    var v5 int32 = 0
    var v6 int32 = v5
    v7 := GenS1{F0: -20, F1: 128}
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(v3, v3)
        var v9 bool = true
        for fi10 := 0; fi10 < 14; fi10++ {
            if v9 {
                var v11 uint32 = v4
                _ = v11
            } else {
                switch (fi10 % ((fi8) | 1)) {
                case -2147483648:
                    if (1 >= (v3 + 1000)) {
                        switch "a" {
                        case "":
                            v4 += v4
                        default:
                            var v12 bool = (v4 > 2147483648)
                            _ = v12
                        }
                        if v9 {
                            v4 += (v4 << ((65535) & 63))
                            v1[(v4 / (((v4 << ((93) & 63))) | 1))] = (v4 << ((65535) & 63))
                        }
                    }
                case -9223372036854775808:
                    var v13 uint32 = 4294967295
                    v14 := []int{-865, (fi8 ^ fi10), 721, fi8}
                    _ = v13
                    _ = v14
                case 32767:
                    v6 = 1
                default:
                    v5 = (v5 << ((18446744073709551615) & 63))
                }
            }
            v4 -= (v4 >> ((748) & 63))
            var v15 float32 = (3.14159 * (float32(1e10) / -2.25))
            _ = fi10
            _ = v15
        }
        _ = fi8
        _ = v9
    }
    switch "hello" {
    case "":
        v16 := []int{32768}
        _ = v16
    case "a":
        fmt.Println(v6, v5, v4)
    case "\n\t":
        var v17 float32 = 1e10
        _ = v17
    }
    v5 = v5
    v18 := GenS1{F0: 75, F1: 0}
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint = 255
        v4 -= 157
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: true, F1: 233}
    v6 -= -1
    switch -807 {
    case -1:
        v22 := GenS3{F0: 0, F1: v6, F2: 65535}
        _ = v22
    }
    v5 += (v5 / ((v6) | 1))
    v23 := v1[1]
    var v24 int32 = int32(v2)
    var v25 uint64 = 694
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
