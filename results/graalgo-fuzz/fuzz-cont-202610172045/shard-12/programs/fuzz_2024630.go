package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    var v5 int16 = 42
    var v6 uint = 42
    v7 := GenS1{F0: (^(v3))}
    switch ("" + "") {
    case "\n\t":
        var v8 string = "foo bar"
        _ = v8
    case "hello":
        if (true && false) {
            var v9 uint = (v6 + v6)
            v9 = (v9 - v6)
            _ = v9
        } else {
            v5 = (v5 / ((-992) | 1))
            var v10 int64 = 0
            v10 = v10
            _ = v10
        }
    default:
        fmt.Println(v6, v3, v6)
    }
    v11 := GenS1{F0: (v3 | (v3 << ((534) & 63)))}
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            if (false && false) {
                var v13 int32 = -2147483648
                v14 := v1[2147483647]
                var v15 int16 = (-(v5))
                _ = v13
                _ = v14
                _ = v15
            } else {
                fmt.Println(v6, v3, v3)
                v5 = (v5 / (((v5 ^ v5)) | 1))
            }
        }
        _ = fi12
    }
    fmt.Println(v5, v5)
    var v16 int64 = -209
    v17 := GenS1{F0: (v3 / ((40) | 1))}
    v16 -= int64(v6)
    if ((v16 - v16) == (v16 >> ((v6) & 63))) {
        v16 -= 1
        v3 = (v3 | v3)
        if false {
            fmt.Println(v3)
            v16 += v16
        } else {
            v18 := []int{-248}
            _ = v18
        }
    }
    v6 = v6
    v19 := GenS1{F0: v3}
    v20 := []int{289, -9223372036854775808, 0}
    var v21 float64 = -2.25
    switch -915 {
    case 9223372036854775807:
        for fi22 := 0; fi22 < 13; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 uint64 = 2
                v5 = v5
                switch fi23 {
                case 4294967295:
                    v5 += v5
                    var v25 int = fi22
                    _ = v25
                case 402:
                    v26 := GenS1{F0: (v3 | (^(v3)))}
                    _ = v26
                case -2147483648:
                    v27 := GenS2{F0: 42}
                    _ = v27
                }
                v5 = v5
                _ = fi23
                _ = v24
            }
            fmt.Println(v21)
            if true {
                v3 = v3
            }
            _ = fi22
        }
    default:
        v20 = append(v20, -2147483648)
    }
    v28 := v1[1]
    v16 -= -9223372036854775808
    v29 := GenS2{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
