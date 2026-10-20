package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0() uint64 {
    var v5 uint = 18446744073709551615
    v5 += (v5 * 112)
    if true {
        fmt.Println(v5)
        v6 := []int{127, -9223372036854775808}
        _ = v6
    }
    v5 -= 219
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := []int{652, 32768}
    v2 := make(map[int]string)
    var v3 uint16 = 918
    var v4 bool = false
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 255, F1: 42, F2: -9223372036854775808}
        var v9 string = ""
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = 42
    v4 = v4
    var v11 int8 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            if (!v4) {
                v4 = (1000 > 2147483647)
            }
            if true {
                v1 = append(v1, (fi13 >> ((2) & 63)))
                if v4 {
                    switch (fi13 | 32767) {
                    case 1:
                        v11 -= (v11 & (^(v11)))
                    case 128:
                        var v14 int16 = v10
                        _ = v14
                    case -1:
                        v15 := GenS1{F0: 65535, F1: 1, F2: 42}
                        _ = v15
                    }
                    v16 := GenS2{F0: 609, F1: fi12, F2: 9223372036854775807}
                    _ = v16
                } else {
                    fmt.Println(v4)
                    var v17 int16 = 1000
                    _ = v17
                }
                var v18 string = "a"
                _ = v18
            } else {
                var v19 int8 = 1
                _ = len(v2)
                _ = v19
            }
            v11 -= (v11 ^ v11)
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v4, v11)
    v3 = (v3 << ((4294967295) & 63))
    var v20 uint16 = 42
    v21 := GenS3{F0: false, F1: 42}
    if v4 {
        v22 := GenS2{F0: 259, F1: 32768, F2: 42}
        var v23 uint8 = 128
        _ = v22
        _ = v23
    } else {
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
