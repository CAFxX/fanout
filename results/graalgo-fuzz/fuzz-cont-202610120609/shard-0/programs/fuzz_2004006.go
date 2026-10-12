package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() int16 {
    var v5 int8 = 2
    v6 := GenS1{F0: 171, F1: 0, F2: 1}
    _ = v5
    _ = v6
    return -265
}
func genFn1(p0 uint64, p1 uint8) int32 {
    v7 := []int{}
    var v8 int64 = -9223372036854775808
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 531
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 uint32 = 212
    v4 -= 218
    v4 -= 2147483647
    if true {
        fmt.Println(v4)
        fmt.Println(v4)
        v4 += (^(v4))
    } else {
        var v9 uint = 485
        _ = v9
    }
    v10 := GenS1{F0: -845, F1: 41, F2: 638}
    if true {
        v4 = (v4 << ((255) & 63))
        fmt.Println(v4)
        fmt.Println(v4)
    } else {
        var v11 bool = true
        _ = v11
    }
    v12 := GenS1{F0: genFn1(175, 255), F1: 920, F2: -578}
    v13 := GenS2{F0: 1}
    var v14 bool = (("a" + "hello") > "")
    delete(v1, (v4 | v4))
    fmt.Println(v4, v14)
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: 0, F1: 1, F2: -2147483648}
        _ = fi15
        _ = v16
    }
    if v14 {
        if (!v14) {
            for fi17 := 0; fi17 < 16; fi17++ {
                v2[2] = "a"
                var v18 string = "a"
                fmt.Println(v4)
                _ = fi17
                _ = v18
            }
        }
    } else {
        v2[v4] = ("\n\t" + "")
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int8 = 0
        v14 = true
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
