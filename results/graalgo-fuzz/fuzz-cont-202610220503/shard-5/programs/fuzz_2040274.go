package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 2
    v3 := make(map[uint32]string)
    var v4 int = -572
    var v5 uint32 = 4294967295
    if (true || (!(4294967295 < uint64(v5)))) {
        v2 = (v2 << ((815) & 63))
        v5 += (v5 / ((1) | 1))
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint = 127
            var v8 uint8 = 255
            v2 += (v2 + v2)
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := []int{(v4 + v4)}
    var v10 int64 = -239
    v5 = (v5 | (v5 >> ((460) & 63)))
    var v11 uint32 = v5
    v4 = -682
    switch ("" + "\n\t") {
    case "a":
        if true {
            fmt.Println(v4, v5, v2)
            if true {
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 float32 = (3.14159 * 3.14159)
                    _ = fi12
                    _ = v13
                }
            }
        }
    default:
        v5 = v5
    }
    v14 := GenS1{F0: 592, F1: 1000}
    if true {
        fmt.Println(v10, v11)
        fmt.Println(v11)
    }
    if (42 > -24) {
        if (!true) {
            v9 = append(v9, (v4 >> ((v2) & 63)))
            v15 := v3[(^(v11))]
            v1[v15] = v15
            _ = v15
        }
    } else {
        v5 = (^(v11))
        switch "hello" {
        case "\n\t":
            var v16 string = ("" + "a")
            _ = v16
        default:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    v11 = 1
    switch (v4 >> ((v2) & 63)) {
    case 0:
        v17 := GenS2{F0: -119, F1: 59, F2: 1}
        _ = v17
    }
    if len(v9) > 0 {
        v9[0] = (v4 + (^(v4)))
    }
    var v18 int = -117
    v3[(v11 / ((v5) | 1))] = "foo bar"
    v19 := GenS1{F0: 0, F1: -243}
    var v20 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
