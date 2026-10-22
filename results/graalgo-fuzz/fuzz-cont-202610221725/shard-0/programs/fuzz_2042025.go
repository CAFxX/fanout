package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    v3 := make(map[string]string)
    var v4 bool = (255 <= 255)
    var v5 uint32 = 2147483648
    v6 := make(map[uint32]string)
    var v7 uint = 658
    fmt.Println(v4, v7)
    v8 := GenS1{F0: 32767}
    if true {
        var v9 float64 = v2
        if (!v4) {
            _ = len(v1)
        } else {
            fmt.Println(v9, v9)
            v10 := GenS2{F0: 2}
            _ = v10
        }
        _ = v9
    } else {
        var v11 int8 = 39
        fmt.Println(v11, v4, v7)
        var v12 int = 4294967295
        _ = v11
        _ = v12
    }
    v7 = v7
    if (0 != 2147483647) {
        if (!false) {
            v2 += v2
            v7 += (^(v7))
            v4 = (0 != 103)
        } else {
            v1[(^(v5))] = (((3.14159 * (0.1 - 0.1)) / 1.5) >= 1e10)
            v6[v5] = "\n\t"
            v13 := []int{}
            _ = v13
        }
        if v4 {
            v1[(v5 ^ 0)] = v4
            fmt.Println(v4, v2)
            if v4 {
                if v4 {
                    v7 = uint(v5)
                }
                v14 := GenS2{F0: 76}
                for fi15 := 0; fi15 < 4; fi15++ {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v17 := GenS2{F0: 255}
                        var v18 uint64 = 42
                        fmt.Println(v7, v18)
                        var v19 bool = v4
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
                _ = v14
            }
        }
    } else {
        fmt.Println(v2, v5)
        fmt.Println(v4, v7)
        var v20 uint32 = v5
        _ = v20
    }
    fmt.Println(v7, v5)
    v4 = true
    var v21 int32 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS3{F0: 355}
        _ = fi22
        _ = v23
    }
    fmt.Println(v7, v21, v4)
    var v24 int16 = -805
    v25 := GenS1{F0: 372}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
