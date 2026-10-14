package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0(p0 uint8, p1 int64) int32 {
    v7 := []int{32768, 32768}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 42
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: (-696 != -1), F1: 228, F2: ((65535 == 709) || (!(127 == 1)))}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = 1
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 0
    var v3 uint16 = v2
    var v4 uint16 = (v2 / ((v3) | 1))
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    switch 128 {
    case 284:
        v12 := GenS1{F0: (-128 >= -128), F1: 682, F2: (false || true)}
        _ = v12
    default:
        var v13 int64 = -651
        _ = v13
    }
    v5["foo bar"] = 9223372036854775807
    v1 = 1
    v2 = (v4 | (v4 / (((v4 * 937)) | 1)))
    v2 -= (v3 - v4)
    v4 -= v2
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                for fi15 := 0; fi15 < 5; fi15++ {
                    fmt.Println(fi15, v1)
                    v1 += v1
                    _ = fi15
                }
                var v16 uint64 = (v1 - 639)
                v17 := GenS2{F0: 777}
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: (v1 ^ (^(v1))), F1: uint(v2), F2: 42}
            v3 += v3
            _ = fi14
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int16 = 158
            v1 += (v1 >> ((4294967295) & 63))
            _ = fi19
            _ = v20
        }
        if (-340 >= -675) {
            v21 := GenS1{F0: (1 == -1), F1: 0, F2: true}
            fmt.Println(v4, v1, v4)
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 float32 = 1.5
                var v24 uint16 = 65535
                var v25 uint = 255
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        }
    } else {
        v3 -= v3
        var v26 uint64 = v1
        _ = v26
    }
    fmt.Println(v4, v2)
    v27 := GenS1{F0: true, F1: v1, F2: true}
    var v28 uint16 = (v2 - 32767)
    var v29 float32 = 1e10
    v2 += 1
    v28 = v2
    var v30 bool = (2 > 42)
    switch "\n\t" {
    case "":
        var v31 int64 = -875
        _ = v31
    case "a":
        fmt.Println(v1, v1, v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
