package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int = 4294967295
    v2 := []uint32{}
    v3 := []int{685, (v1 << ((4294967295) & 63))}
    v4 := make(map[int]string)
    var v5 bool = true
    var v6 int32 = 0
    fmt.Println(v6)
    switch (-(v1)) {
    case 32767:
        if (!(1000 < -831)) {
            v7 := GenS1{F0: 9223372036854775807}
            v8 := GenS1{F0: 1}
            _ = v7
            _ = v8
        }
    case 42:
        v9 := GenS1{F0: 1}
        _ = v9
    case -499:
        if (true && v5) {
            var v10 int64 = 131
            _ = v10
        } else {
            v2 = append(v2, 2147483647)
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 int = (fi11 * 0)
                _ = fi11
                _ = v12
            }
            fmt.Println(v5, v6)
        }
    }
    var v13 uint8 = 1
    var v14 string = ("foo bar" + ("a" + "hello"))
    v15 := GenS2{F0: 2147483648, F1: 0}
    v1 += 127
    if (72 < 0) {
        if false {
            var v16 uint8 = 128
            _ = v16
        }
        v17 := GenS2{F0: 0, F1: -274}
        _ = v17
    }
    var v18 int16 = 928
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
