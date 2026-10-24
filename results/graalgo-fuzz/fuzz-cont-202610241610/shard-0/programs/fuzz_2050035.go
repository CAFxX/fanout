package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int) int8 {
    if (!false) {
        v5 := GenS1{F0: 1, F1: 25}
        v6 := []int{p0}
        p0 -= p0
        _ = v5
        _ = v6
    }
    if false {
        fmt.Println(p0)
        p0 = p0
    } else {
        var v7 int8 = -1
        p0 += p0
        _ = v7
    }
    v8 := GenS2{F0: 131, F1: -1}
    v9 := []int{p0, p0, (p0 % (((p0 & p0)) | 1))}
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int = -425
    var v2 int8 = 42
    var v3 uint16 = 182
    v4 := make(map[string]bool)
    var v10 bool = (!false)
    var v11 uint16 = (v3 / (((v3 ^ 755)) | 1))
    for fi12 := 0; fi12 < 15; fi12++ {
        v4["foo bar"] = true
        v13 := v4["\n\t"]
        v1 += (fi12 % ((32767) | 1))
        _ = fi12
        _ = v13
    }
    _ = len(v4)
    v14 := GenS2{F0: 255, F1: -1}
    v11 += (v11 + 32767)
    v15 := GenS2{F0: 127, F1: 9223372036854775807}
    v1 += (v1 ^ -90)
    v16 := GenS2{F0: 147, F1: 1}
    for fi17 := 0; fi17 < 1; fi17++ {
        v11 += (v3 * v11)
        if v10 {
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 int32 = 1
                _ = fi18
                _ = v19
            }
            v20 := GenS3{F0: 0}
            _ = v20
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
