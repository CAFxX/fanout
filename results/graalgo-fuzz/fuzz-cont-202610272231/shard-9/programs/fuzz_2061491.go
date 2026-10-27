package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }

func main() {
    var v1 int8 = 113
    v2 := make(map[string]int)
    var v3 uint = 960
    var v4 bool = true
    var v5 uint16 = 65535
    v6 := []int{1, -198}
    var v7 int8 = v1
    var v8 uint16 = 134
    v9 := v2["\n\t"]
    if (((-(v9)) > (v9 / ((v9) | 1))) || (4294967295 <= 42)) {
        var v10 int = int(v3)
        _ = v10
    }
    if (0.1 != 3.14159) {
        if v4 {
            v11 := GenS1{F0: -32768, F1: (v9 / (((v9 | v9)) | 1)), F2: (^(v5))}
            _ = v11
        } else {
            var v12 int64 = -1
            v5 += uint16(v3)
            var v13 uint8 = 42
            _ = v12
            _ = v13
        }
    } else {
        v7 += v1
        v14 := v2["hello"]
        _ = len(v6)
        _ = v14
    }
    v6 = append(v6, v9)
    v15 := GenS1{F0: 1000, F1: (^(v9)), F2: (v8 ^ (v8 % ((v5) | 1)))}
    if false {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint32 = 1
            var v18 float64 = (1.5 + 0.0)
            _ = fi16
            _ = v17
            _ = v18
        }
        v8 += v5
        v4 = (false && (int8(v3) > v1))
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        if false {
            var v20 int = (fi19 | 519)
            var v21 float64 = ((float64(0.1) - 1e10) - 3.14159)
            _ = v20
            _ = v21
        }
        var v22 float64 = 0.1
        _ = fi19
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
