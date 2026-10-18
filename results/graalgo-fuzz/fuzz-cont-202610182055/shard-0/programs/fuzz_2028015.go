package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint64 }

func main() {
    v1 := []string{("\n\t" + "hello"), "hello"}
    v2 := make(map[int]int)
    v3 := []int{2147483647}
    v4 := []uint32{}
    if true {
        var v5 int16 = 42
        var v6 uint16 = 0
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        v8 := GenS1{F0: 462, F1: 569}
        v9 := GenS2{F0: 255, F1: 42, F2: -2147483648}
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 127, F1: 5, F2: 42}
    _ = len(v3)
    var v11 float32 = (0.0 * 0.1)
    var v12 bool = false
    delete(v2, 1)
    v13 := GenS3{F0: 1, F1: 4294967295}
    v12 = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (740 >= 953) {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int = (fi14 & fi14)
            for fi16 := 0; fi16 < 8; fi16++ {
                if true {
                    var v17 uint64 = 821
                    fmt.Println(v17, v17, v11)
                    v15 += 80
                    _ = v17
                } else {
                    v18 := GenS3{F0: 472, F1: 65535}
                    v19 := GenS4{F0: 32767, F1: 0, F2: 2}
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            v20 := GenS3{F0: 0, F1: 2147483648}
            _ = fi14
            _ = v15
            _ = v20
        }
        v11 += v11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
