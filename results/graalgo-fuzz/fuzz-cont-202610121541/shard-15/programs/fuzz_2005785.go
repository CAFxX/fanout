package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 1
    var v3 int64 = (v2 >> ((42) & 63))
    var v4 float64 = ((1.5 + 3.14159) / 0.1)
    var v5 float64 = v4
    fmt.Println(v5)
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v3)
        v4 -= (1.5 * v4)
        v3 += 0
        _ = fi6
    }
    v3 = 48
    v7 := GenS1{F0: 224, F1: -32768, F2: 711}
    _ = len(v1)
    v4 = v5
    v8 := GenS1{F0: 128, F1: -1, F2: 42}
    for fi9 := 0; fi9 < 6; fi9++ {
        if (127 >= 169) {
            v10 := GenS1{F0: 2, F1: 200, F2: 18446744073709551615}
            var v11 uint16 = 618
            _ = v10
            _ = v11
        }
        var v12 int64 = v3
        v13 := GenS2{F0: 515, F1: fi9}
        _ = fi9
        _ = v12
        _ = v13
    }
    var v14 float64 = (v4 / ((((v4 + float64(((3.14159 - (3.14159 * float32(v4))) * ((float32(v5) + float32(1.5)) + 0.1)))) * v4) / v4) - 1e10))
    var v15 uint8 = 0
    var v16 int64 = v3
    _ = len(v1)
    fmt.Println(v5)
    if (((v5 / 1.0) / (((float64((1.5 - (1e10 * 3.14159))) / 1e10) + v5) * v14)) == ((-2.25 / float64((-2.25 / 3.14159))) + v14)) {
        switch "" {
        case "a":
            v17 := GenS2{F0: 255, F1: 2147483647}
            _ = v17
        default:
            if true {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS2{F0: 1, F1: fi18}
                    var v20 int = fi18
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v2)
            } else {
                v21 := GenS3{F0: 0}
                v22 := GenS2{F0: 255, F1: 4294967295}
                _ = v21
                _ = v22
            }
        }
        if true {
            var v23 uint16 = 255
            var v24 int = 9223372036854775807
            _ = v23
            _ = v24
        }
    } else {
        var v25 float32 = 1e10
        v26 := GenS1{F0: v15, F1: 32767, F2: 2}
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
