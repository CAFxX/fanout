package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int; F2 string }

func main() {
    v1 := []uint32{65535, 566}
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    v4 := []string{}
    v5 := []uint32{}
    v6 := make(map[string]int)
    var v7 float64 = ((0.0 / 1.0) * 3.14159)
    if (v3 > v3) {
        delete(v6, "a")
        for fi8 := 0; fi8 < 20; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v9 := v2[2147483648]
            var v10 uint64 = 168
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                if (933 >= 1) {
                    var v13 int = fi11
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v13
                }
                var v14 bool = (127 >= 520)
                v14 = (float32(0.0) >= float32(1e10))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi12
                _ = v14
            }
            _ = fi11
        }
    }
    v7 += ((v7 - ((((((1e10 / 1e10) * (v7 / (v7 * (((v7 / (v7 - 3.14159)) / v7) - v7)))) - float64(1e10)) / v7) - v7) + 0.0)) + (3.14159 - v7))
    v15 := GenS1{F0: 42}
    switch 32768 {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 15; fi16++ {
            v3 = (v3 + (((v3 + "\n\t") + ("foo bar" + v3)) + v3))
            _ = fi16
        }
    case 1:
        if (!(v3 < v3)) {
            v3 = v3
            v17 := GenS2{F0: 934}
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: 221}
                var v20 int = fi18
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
    default:
        v21 := GenS2{F0: 42}
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v3 = v3
        v23 := GenS2{F0: 1}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 42, F1: 2147483647, F2: (v3 + "a")}
    var v25 float32 = float32((v7 - (((1.5 - (v7 + float64(0.0))) / (0.1 / 1.0)) / v7)))
    switch 32767 {
    case -450:
        v3 = v3
    }
    switch v3 {
    case "":
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
