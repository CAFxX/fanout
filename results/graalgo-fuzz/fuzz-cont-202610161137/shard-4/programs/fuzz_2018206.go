package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 9
    var v3 int16 = 933
    var v4 uint64 = 4294967295
    var v5 int16 = v3
    v6 := make(map[int]int)
    v1 -= -9223372036854775808
    v7 := []int{2147483647, -828, -977}
    v7 = append(v7, 4294967295)
    switch -75 {
    case -732:
        v3 -= (v3 >> ((65535) & 63))
    case 798:
        switch 0 {
        case 42:
            v8 := GenS1{F0: -1, F1: v2}
            _ = v8
        case 127:
            v9 := GenS1{F0: int8(v4), F1: 2}
            _ = v9
        case -768:
            v3 -= v5
        default:
            switch 4294967295 {
            case 127:
                v7 = append(v7, -137)
            case 432:
                for fi10 := 0; fi10 < 14; fi10++ {
                    fmt.Println(v2, v1)
                    var v11 uint32 = 833
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (836 >= 255) {
                            var v13 int = fi10
                            v14 := GenS2{F0: 821}
                            v6[32767] = v13
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(v5)
                        }
                        v15 := GenS1{F0: 115, F1: 987}
                        var v16 float64 = 1e10
                        _ = fi12
                        _ = v15
                        _ = v16
                    }
                    _ = fi10
                    _ = v11
                }
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        }
    default:
        v17 := GenS2{F0: 806}
        _ = v17
    }
    fmt.Println(v5, v1)
    for fi18 := 0; fi18 < 3; fi18++ {
        v4 = 42
        switch "a" {
        case "":
            switch 127 {
            case 32767:
                v6[(fi18 % (((fi18 | fi18)) | 1))] = fi18
            case 0:
                var v19 uint8 = 29
                _ = v19
            }
        case "a":
            switch 42 {
            case 19:
                var v20 int64 = int64(v4)
                _ = v20
            }
        default:
            v21 := GenS2{F0: 65535}
            _ = v21
        }
        var v22 int8 = 127
        _ = fi18
        _ = v22
    }
    var v23 bool = true
    v24 := GenS2{F0: 58}
    var v25 float32 = 1e10
    v26 := GenS3{F0: false}
    var v27 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
