package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 1
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 2147483647}
    v5 := GenS1{F0: 1}
    var v6 int32 = 2147483647
    var v7 uint = 255
    var v8 int = 32768
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = false
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int = (-(v8))
                v13 := GenS2{F0: -1, F1: (v7 + (v7 + v7)), F2: 1}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        v14 := GenS1{F0: 4294967295}
        for fi15 := 0; fi15 < 5; fi15++ {
            delete(v3, 262)
            _ = fi15
        }
        _ = v14
    } else {
        var v16 string = "\n\t"
        _ = v16
    }
    v7 = (v7 / ((v7) | 1))
    var v17 uint16 = 868
    v18 := GenS2{F0: 53, F1: (^(v7)), F2: -410}
    var v19 uint64 = 2
    var v20 uint32 = 270
    if false {
        v21 := GenS3{F0: v6}
        var v22 int64 = 1
        _ = v21
        _ = v22
    }
    v2 = v2
    var v23 int16 = -567
    fmt.Println(v8)
    v7 -= (v7 - v7)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
