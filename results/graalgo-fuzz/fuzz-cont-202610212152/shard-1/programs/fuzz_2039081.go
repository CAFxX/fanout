package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint8) int {
    var v7 uint8 = 255
    fmt.Println(v7)
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]string)
    var v3 float64 = 1e10
    var v4 int8 = -1
    var v5 uint64 = 608
    v6 := []int{}
    fmt.Println(v4, v4, v1)
    var v8 int32 = 96
    for fi9 := 0; fi9 < 20; fi9++ {
        if (v1 >= v1) {
            switch (fi9 << ((65535) & 63)) {
            case 4294967295:
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int32 = 360
                    fmt.Println(fi10, fi9, fi10)
                    _ = fi10
                    _ = v11
                }
            }
            var v12 uint = 255
            _ = v12
        }
        delete(v2, fi9)
        _ = fi9
    }
    v13 := GenS1{F0: v5, F1: 18446744073709551615}
    var v14 int16 = 168
    fmt.Println(v5)
    var v15 int = 484
    var v16 string = "\n\t"
    var v17 uint16 = 2
    if (((-(v15)) < genFn0(233)) && true) {
        fmt.Println(v15, v17, v4)
        v1 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
