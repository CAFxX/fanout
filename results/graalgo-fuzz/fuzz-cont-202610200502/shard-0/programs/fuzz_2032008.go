package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
func genFn0(p0 int8) int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{(fi7 & fi7), fi7, fi7, 922}
        var v9 int = fi7
        var v10 uint16 = 2
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{(fi11 / ((fi11) | 1)), (fi11 - fi11), 0}
        if (!(!false)) {
            p0 = p0
            p0 = (p0 << ((102) & 63))
            v13 := GenS1{F0: ((("a" + "a") + "hello") + "a"), F1: (3.14159 < -2.25)}
            _ = v13
        } else {
            var v14 uint32 = 2147483647
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS2{F0: 344, F1: 42}
    _ = p0
    _ = v15
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -90
    v3 := make(map[string]bool)
    var v4 uint64 = 102
    var v5 uint64 = 1
    var v6 uint8 = 127
    _ = len(v1)
    v5 += (v4 << ((1) & 63))
    v16 := GenS3{F0: v2, F1: 2, F2: 4294967295}
    var v17 uint16 = uint16(v6)
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := []int{(fi18 & fi18), -2147483648}
        var v20 bool = false
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: (v2 - (v2 % (((v2 + -44)) | 1))), F1: 467, F2: 243}
    if true {
        delete(v3, "hello")
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 int = fi22
            _ = fi22
            _ = v23
        }
    }
    v1[154] = 4294967295
    v24 := GenS3{F0: v2, F1: 239, F2: 222}
    if true {
        fmt.Println(v6, v6)
    }
    v5 = v4
    v4 = (^(v5))
    v17 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
