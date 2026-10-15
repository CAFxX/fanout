package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int32, p1 int64, p2 uint16) bool {
    var v7 int32 = (p0 / (((^(p0))) | 1))
    var v8 int64 = -9223372036854775808
    v9 := GenS1{F0: 42, F1: true, F2: 1}
    if (!false) {
        p1 += p1
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int64 = (p1 << ((42) & 63))
                _ = fi11
                _ = v12
            }
            p0 = (p0 * (p0 + 66))
            v13 := []int{fi10, (fi10 >> ((2) & 63)), (fi10 & fi10), fi10}
            _ = fi10
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ((p2 % (((p2 - 42)) | 1)) != (p2 & (p2 & 255)))
}

func main() {
    var v1 int = 42
    var v2 int64 = 42
    v3 := make(map[int]int)
    var v4 uint = 2
    var v5 int32 = -1
    v6 := []string{"a"}
    _ = len(v3)
    if (128 > 0) {
        v14 := GenS2{F0: 125, F1: "a"}
        _ = v14
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS3{F0: v1}
        switch fi15 {
        case 32768:
            switch (fi15 << (((v4 | v4)) & 63)) {
            case -459:
                v17 := GenS4{F0: 42, F1: 69}
                _ = v17
            default:
                if (false && ("" == "")) {
                    var v18 bool = (!true)
                    _ = v18
                }
            }
        default:
            v19 := GenS2{F0: 268, F1: "\n\t"}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    var v20 int = v1
    v21 := GenS5{F0: -77, F1: 1, F2: 18446744073709551615}
    v22 := GenS5{F0: 0, F1: 811, F2: 18446744073709551615}
    v23 := v3[2147483647]
    v24 := GenS6{F0: 2}
    var v25 string = "hello"
    v26 := GenS4{F0: 42, F1: -128}
    v27 := GenS4{F0: 42, F1: 1}
    var v28 string = "\n\t"
    var v29 int32 = (v5 % (((v5 - v5)) | 1))
    for fi30 := 0; fi30 < 10; fi30++ {
        v25 = v25
        var v31 float64 = 0.0
        delete(v3, 32768)
        _ = fi30
        _ = v31
    }
    fmt.Println(v20, v20)
    var v32 float32 = (1.5 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
