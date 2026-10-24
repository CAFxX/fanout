package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []string{""}
    var v2 bool = false
    var v3 float64 = 0.0
    var v4 int16 = 1000
    v5 := make(map[string]int)
    var v6 float32 = (0.1 - (0.0 / 0.1))
    var v7 uint64 = 516
    var v8 int16 = 1000
    if (100 == -1) {
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 98, F1: 348}
                var v12 float64 = v3
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    _ = len(v5)
    v5["hello"] = -2147483648
    if v2 {
        v3 += v3
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint = 0
            _ = fi13
            _ = v14
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            switch 4294967295 {
            case 607:
                fmt.Println(v8, v6, v8)
            case -9223372036854775808:
                if (100 == int8(v4)) {
                    v5["\n\t"] = 1
                    v7 = 42
                } else {
                    v7 -= 228
                    switch 32767 {
                    case -802:
                        v15 := GenS1{F0: 42, F1: 1}
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v8 += v8
                            _ = fi16
                        }
                    }
                }
            default:
                switch -9223372036854775808 {
                case -587:
                    fmt.Println(v8)
                }
            }
        default:
            v17 := GenS2{F0: 32767}
            _ = v17
        }
    }
    v18 := GenS3{F0: 2147483648}
    v19 := GenS4{F0: (v7 >> ((2) & 63))}
    if (118 >= 2) {
        delete(v5, "")
    }
    var v20 int = -415
    v21 := GenS3{F0: 65535}
    v22 := GenS3{F0: 2}
    for fi23 := 0; fi23 < 12; fi23++ {
        v8 += v8
        _ = fi23
    }
    var v24 string = "a"
    v25 := v5[v24]
    v26 := GenS5{F0: v7}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
