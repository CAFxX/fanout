package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 int8 = 60
    v2 := make(map[int]int)
    var v3 uint = 512
    var v4 uint32 = 2147483648
    if (-2147483648 > 127) {
        v4 = v4
        var v5 uint32 = v4
        v6 := GenS1{F0: false, F1: -9223372036854775808}
        _ = v5
        _ = v6
    }
    var v7 uint16 = 373
    fmt.Println(v3)
    fmt.Println(v1, v3)
    var v8 uint64 = 494
    fmt.Println(v7, v3)
    fmt.Println(v1, v8)
    v9 := v2[int(v7)]
    fmt.Println(v1, v3)
    switch (("hello" + "") + "\n\t") {
    case "hello":
        var v10 uint8 = 211
        _ = v10
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            delete(v2, (v9 * (v9 << (((v3 % ((v3) | 1))) & 63))))
            _ = fi11
        }
    case "\n\t":
        var v12 uint32 = 774
        _ = v12
    default:
        switch "foo bar" {
        case "\n\t":
            if (!true) {
                for fi13 := 0; fi13 < 17; fi13++ {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        fmt.Println(v8, v4)
                        if (0.0 == float32(0.0)) {
                            fmt.Println(v3)
                        } else {
                            var v15 uint64 = 1
                            v15 = (v15 ^ (^(v8)))
                            var v16 uint = v3
                            _ = v15
                            _ = v16
                        }
                        _ = fi14
                    }
                    _ = fi13
                }
                v17 := []int{(v9 * int(v3)), (v9 % (((v9 - 9223372036854775807)) | 1)), v9}
                v18 := GenS1{F0: (false || true), F1: int64(v9)}
                _ = v17
                _ = v18
            }
        case "hello":
            fmt.Println(v3)
        case "":
            fmt.Println(v4)
        }
    }
    if false {
        v19 := GenS1{F0: (v1 == (v1 % ((v1) | 1))), F1: 1}
        var v20 uint8 = 128
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: -1, F1: v7, F2: (v1 * (v1 % ((v1) | 1)))}
    var v22 bool = true
    v22 = v22
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS1{F0: (fi23 >= fi23), F1: -9223372036854775808}
        var v25 uint32 = (^(v4))
        v9 -= v9
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
