package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16, p1 int) int8 {
    var v7 int16 = 0
    var v8 uint8 = 90
    if false {
        p1 += p1
    }
    v9 := GenS1{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -54
}

func main() {
    v1 := []uint32{1, 0, 42}
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 int16 = 493
    var v5 string = ("hello" + ("" + "a"))
    var v6 uint32 = 4294967295
    v6 += v6
    v6 -= v6
    v6 += 4294967295
    v10 := v3[32767]
    fmt.Println(v10, v6, v5)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint16 = 65535
        v13 := GenS2{F0: 42, F1: 42}
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 bool = true
            _ = fi14
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v16 := GenS1{F0: 18446744073709551615}
    fmt.Println(v6)
    switch -325 {
    case 4294967295:
        v17 := GenS3{F0: 637, F1: 371}
        _ = v17
    default:
        v18 := v3[208]
        _ = v18
    }
    v3[32768] = "foo bar"
    var v19 string = v10
    v20 := GenS2{F0: 127, F1: 65535}
    v21 := GenS4{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
