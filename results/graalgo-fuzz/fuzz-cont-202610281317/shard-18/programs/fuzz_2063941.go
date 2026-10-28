package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 2
    v3 := []int{128, 128, 32767}
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    var v6 int64 = -1
    if (uint(v1) < 65535) {
        for fi7 := 0; fi7 < 5; fi7++ {
            if false {
                v8 := v4["hello"]
                var v9 int64 = (v6 & int64(fi7))
                _ = v8
                _ = v9
            } else {
                v10 := v5[2]
                _ = v10
            }
            fmt.Println(v1, fi7, v6)
            if (!((v6 & -1) <= v6)) {
                fmt.Println(v2, v2, fi7)
                var v11 string = "\n\t"
                _ = v11
            }
            _ = fi7
        }
    } else {
        var v12 int = -570
        v13 := GenS1{F0: uint16(v2), F1: 255, F2: -39}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 1, F1: 499, F2: -1}
    v15 := GenS1{F0: 245, F1: uint16(v2), F2: -86}
    fmt.Println(v2)
    v3 = append(v3, 42)
    v16 := v4["foo bar"]
    v17 := GenS2{F0: 9223372036854775807}
    v18 := GenS3{F0: 42, F1: v2}
    for fi19 := 0; fi19 < 7; fi19++ {
        v6 += (v6 | 0)
        _ = fi19
    }
    v5[65535] = true
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
