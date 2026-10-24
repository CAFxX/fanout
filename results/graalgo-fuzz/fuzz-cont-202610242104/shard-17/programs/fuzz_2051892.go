package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        if (((0.0 < 3.14159) || (0 == 0)) || false) {
            var v5 uint = 255
            _ = v5
        }
        v6 := []int{}
        v7 := GenS1{F0: 42}
        _ = fi4
        _ = v6
        _ = v7
    }
    var v8 uint64 = 42
    v9 := GenS1{F0: 42}
    var v10 int64 = -1
    fmt.Println(v8, v8)
    _ = v10
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 bool, p1 int16, p2 uint64) uint {
    var v11 uint8 = 102
    v12 := []int{-9223372036854775808}
    switch -1 {
    case 4294967295:
        p0 = p0
    case 9223372036854775807:
        v13 := GenS1{F0: 2}
        _ = v13
    case 32768:
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 string = "a"
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    fmt.Println(v1, v2)
    fmt.Println(v1, v2)
    v15 := []int{4294967295, 127, int(v2), -582}
    var v16 uint64 = 802
    var v17 int = 128
    v17 += (v17 % (((^(v17))) | 1))
    v16 += (v16 >> (((v2 ^ (^(v2)))) & 63))
    v2 = v2
    fmt.Println(v17, v16)
    delete(v3, (v17 / (((v17 + 128)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
