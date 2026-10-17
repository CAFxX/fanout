package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 uint = 1
    var v5 int16 = 1
    v6 := make(map[string]int)
    if (4294967295 >= 2147483647) {
        var v7 uint = (v4 / (((v4 & 0)) | 1))
        if false {
            v8 := []int{}
            var v9 bool = false
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: -657, F1: (v5 + (v5 + v5)), F2: -239}
        _ = v10
        _ = v7
    } else {
        v4 = v4
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 1000, F1: (v5 ^ (v5 << ((v4) & 63))), F2: -243}
        if (65535 <= 1) {
            switch "\n\t" {
            case "a":
                v13 := GenS1{F0: -2147483648, F1: (-(v5)), F2: -995}
                _ = v13
            }
            v6[v1] = (fi11 | (fi11 * fi11))
            fmt.Println(v1)
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: true}
    var v15 bool = (false || false)
    var v16 uint8 = 128
    v1 = ("\n\t" + (v1 + ("a" + ("a" + ((v1 + v1) + "")))))
    var v17 float32 = 0.1
    v18 := v6[(v1 + ("a" + (v1 + v1)))]
    var v19 uint32 = 4294967295
    _ = len(v6)
    var v20 uint32 = v19
    v20 += (^(v19))
    v21 := GenS2{F0: v15}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
