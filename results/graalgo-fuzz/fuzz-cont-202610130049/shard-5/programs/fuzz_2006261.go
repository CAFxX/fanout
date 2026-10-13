package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }

func main() {
    var v1 uint8 = 27
    v2 := make(map[int]int)
    var v3 float32 = ((1.5 - (float32(1.5) / (0.0 - 1e10))) * float32(3.14159))
    var v4 uint8 = v1
    var v5 bool = true
    v6 := GenS1{F0: 9223372036854775807}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 0
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 float32 = v3
            if v5 {
                if (true && v5) {
                    v8 = v8
                    if v5 {
                        var v11 int32 = 461
                        v12 := GenS1{F0: -9223372036854775808}
                        _ = v11
                        _ = v12
                    }
                    v13 := []int{fi7, 186, fi7}
                    _ = v13
                }
                var v14 float32 = v10
                switch "a" {
                case "":
                    fmt.Println(fi9)
                default:
                    fmt.Println(v10)
                }
                var v15 uint = 589
                v16 := GenS1{F0: -9223372036854775808}
                switch "" {
                case "\n\t":
                    var v17 int = -9223372036854775808
                    _ = v17
                }
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v18 := GenS1{F0: 299}
    fmt.Println(v4, v5, v5)
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{(fi19 - (fi19 << ((18446744073709551615) & 63)))}
        _ = fi19
        _ = v20
    }
    v21 := []int{201}
    v5 = v5
    v4 = 2
    v22 := v2[1]
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v5, fi23, v1)
        v22 += (v22 | v22)
        _ = fi23
    }
    v3 += v3
    var v24 float32 = v3
    var v25 float64 = 1.5
    var v26 int16 = 42
    var v27 string = ""
    v28 := GenS2{F0: 42, F1: v4, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
