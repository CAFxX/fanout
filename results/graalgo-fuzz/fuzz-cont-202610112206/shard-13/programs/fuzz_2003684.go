package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64 }
func genFn0() bool {
    var v6 float32 = (1.5 + 1e10)
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 -= (((v6 * float32(float64(3.14159))) / v6) / v6)
        _ = fi7
    }
    v8 := []int{-1, -9223372036854775808, 128, 9223372036854775807}
    _ = v6
    _ = v8
    return (((v6 * ((float32(3.14159) - float32(1.5)) * v6)) == ((0.0 - (0.1 + float32((0.0 - (1.5 + 0.0))))) * 3.14159)) && (!false))
}
func genFn1(p0 uint) uint8 {
    if false {
        fmt.Println(p0)
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                p0 += (p0 - p0)
                p0 += p0
                _ = fi10
            }
            var v11 int16 = 0
            var v12 int64 = 1
            _ = fi9
            _ = v11
            _ = v12
        }
        var v13 uint32 = 2
        _ = v13
    } else {
        v14 := GenS1{F0: 65, F1: 0, F2: 1}
        v15 := GenS1{F0: 1, F1: (^(p0)), F2: 65535}
        p0 += 4294967295
        _ = v14
        _ = v15
    }
    p0 = p0
    _ = p0
    return 173
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 391
    var v3 uint64 = 977
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    v3 = v3
    v16 := GenS1{F0: 792, F1: (v2 * v2), F2: 65535}
    v17 := GenS1{F0: 163, F1: (v2 - (v2 ^ v2)), F2: 65535}
    v2 = 2
    v18 := GenS2{F0: 733, F1: true}
    var v19 float32 = float32(3.14159)
    fmt.Println(v19, v3, v3)
    fmt.Println(v2, v3, v19)
    if true {
        v2 -= 255
    } else {
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v3)
            _ = fi20
        }
    }
    var v21 bool = false
    fmt.Println(v21, v21, v19)
    var v22 string = "\n\t"
    for fi23 := 0; fi23 < 17; fi23++ {
        switch (fi23 << (((v2 % (((^(v2))) | 1))) & 63)) {
        case 4294967295:
            if false {
                var v24 uint32 = 289
                _ = v24
            } else {
                for fi25 := 0; fi25 < 20; fi25++ {
                    v26 := []int{fi23, (fi23 - (fi25 >> ((v2) & 63))), fi25, 2147483647}
                    v21 = (32767 != 138)
                    _ = fi25
                    _ = v26
                }
            }
        default:
            v27 := GenS3{F0: -9223372036854775808}
            _ = v27
        }
        var v28 int16 = -747
        v3 += (v3 - (v3 << ((v2) & 63)))
        _ = fi23
        _ = v28
    }
    var v29 uint = v2
    var v30 int32 = -1
    delete(v5, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
