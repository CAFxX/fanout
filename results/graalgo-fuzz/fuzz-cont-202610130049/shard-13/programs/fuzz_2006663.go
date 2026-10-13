package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
func genFn0() int64 {
    v6 := GenS1{F0: 0, F1: 280, F2: 192}
    if false {
        var v7 bool = false
        _ = v7
    } else {
        switch "hello" {
        case "":
            var v8 uint64 = 439
            _ = v8
        case "hello":
            v9 := GenS1{F0: 1, F1: 1000, F2: 2147483647}
            _ = v9
        default:
            v10 := GenS1{F0: 18446744073709551615, F1: 1000, F2: -9223372036854775808}
            _ = v10
        }
    }
    var v11 bool = (1.5 > (((((1.5 / (((3.14159 / (3.14159 / 1e10)) - (float64(-2.25) / (3.14159 / float64(1.5)))) * 3.14159)) - 3.14159) + float64(1e10)) / 1e10) * 0.1))
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 float32 = -2.25
        _ = fi12
        _ = v13
    }
    _ = v11
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 int32 = 2147483647
    v4 := []uint32{656, 0, 1, 4294967295}
    v5 := []int{-208}
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 bool = true
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v3 += (-(v3))
    v3 += (v3 / (((v3 / (((v3 - v3)) | 1))) | 1))
    delete(v2, 764)
    v17 := GenS2{F0: (v3 << ((1) & 63)), F1: "a", F2: 0}
    for fi18 := 0; fi18 < 14; fi18++ {
        fmt.Println(fi18, fi18)
        for fi19 := 0; fi19 < 4; fi19++ {
            if false {
                v20 := GenS2{F0: 0, F1: "", F2: 32767}
                _ = v20
            }
            v5 = append(v5, fi18)
            var v21 uint64 = 611
            _ = fi19
            _ = v21
        }
        fmt.Println(fi18, fi18)
        _ = fi18
    }
    v3 -= (^(v3))
    if true {
        var v22 int16 = 1
        _ = v22
    }
    var v23 int64 = genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
