package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 bool; F1 int8; F2 uint }

func main() {
    var v1 int32 = -853
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    var v4 uint32 = 4294967295
    v5 := []int{-2147483648, -318, 4294967295, 104}
    v4 += (^(v4))
    v2[v3] = (^(v4))
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint8 = 42
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v3 = (v3 + v3)
        v4 -= (v4 << ((4294967295) & 63))
        var v9 uint = 65535
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int64 = 1
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 string = "a"
            v11 = (v11 ^ int64(fi10))
            var v14 int = (fi12 - (fi12 << ((65535) & 63)))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    var v15 int32 = v1
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint8 = 47
        if (false && (128 < 0)) {
            fmt.Println(v1, v4)
            for fi18 := 0; fi18 < 8; fi18++ {
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                var v19 uint16 = 106
                var v20 uint8 = v17
                _ = fi18
                _ = v19
                _ = v20
            }
            switch "\n\t" {
            case "a":
                _ = len(v5)
            case "\n\t":
                switch fi16 {
                case 2147483647:
                    v21 := GenS1{F0: v17, F1: (fi16 & (fi16 % ((fi16) | 1)))}
                    _ = v21
                case 791:
                    v22 := GenS1{F0: v17, F1: (fi16 & fi16)}
                    _ = v22
                default:
                    if (552 >= 4294967295) {
                        switch fi16 {
                        case 1:
                            var v23 uint8 = (v17 + (v17 | (v17 + (v17 ^ (^(v17))))))
                            _ = v23
                        }
                    } else {
                        v3 = v3
                        v17 += v17
                    }
                }
            }
        } else {
            v24 := GenS2{F0: 4294967295, F1: v4, F2: 510}
            fmt.Println(fi16, v15, v17)
            fmt.Println(v17, v15, fi16)
            _ = v24
        }
        fmt.Println(v1)
        _ = fi16
        _ = v17
    }
    var v25 float64 = 3.14159
    _ = len(v2)
    v26 := GenS3{F0: "a", F1: 255}
    v27 := GenS4{F0: false, F1: 0, F2: 65535}
    var v28 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
