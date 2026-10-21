package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 float32 = 0.0
    v4 := make(map[string]string)
    v5 := []int{902, 32768}
    var v6 int = 4294967295
    v6 -= v6
    v6 += (v6 ^ v6)
    v5 = append(v5, 42)
    if (59 <= 18446744073709551615) {
        if (1 > -1) {
            var v7 int64 = 505
            _ = v7
        } else {
            _ = len(v5)
            v2[65535] = (v6 ^ (v6 ^ -239))
            if (((!true) && (true && (((255 != 640) && true) && (!true)))) && (1e10 < (v3 * 3.14159))) {
                v6 = (v6 << ((42) & 63))
                _ = len(v5)
                v3 -= v3
            }
        }
    }
    var v8 float64 = float64(v3)
    _ = len(v5)
    fmt.Println(v8)
    v4[""] = "\n\t"
    fmt.Println(v8, v8, v8)
    v9 := GenS1{F0: 65535, F1: 42, F2: 42}
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            switch 32768 {
            case 711:
                for fi11 := 0; fi11 < 9; fi11++ {
                    v6 += fi10
                    _ = fi11
                }
            case -1:
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := GenS2{F0: 42}
                    v8 += (v8 + (1.5 - ((v8 * 3.14159) + 3.14159)))
                    fmt.Println(v6)
                    v14 := v4[""]
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            var v15 bool = (!true)
            v6 = (fi10 - (v6 / ((fi10) | 1)))
            _ = v15
        }
        if true {
            v6 = v6
            fmt.Println(fi10, fi10, v6)
        } else {
            v6 = v6
        }
        var v16 uint = 42
        _ = fi10
        _ = v16
    }
    if true {
        v6 = v6
    }
    if (!false) {
        switch v6 {
        case -1:
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 int8 = 118
                    fmt.Println(v3, v8)
                    delete(v1, 42)
                    _ = fi18
                    _ = v19
                }
                v20 := GenS2{F0: 2147483647}
                _ = fi17
                _ = v20
            }
        case 1:
            var v21 int32 = -1
            _ = v21
        case 610:
            var v22 int = v6
            _ = v22
        default:
            v23 := GenS3{F0: 255, F1: false, F2: 42}
            _ = v23
        }
        _ = len(v5)
        var v24 uint16 = 42
        _ = v24
    }
    var v25 uint16 = 32767
    v25 -= v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
