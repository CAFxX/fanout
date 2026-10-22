package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int64 = -832
    v2 := make(map[uint32]int)
    var v3 uint32 = 0
    var v4 float32 = (0.0 * (-2.25 / 1.5))
    var v5 uint16 = 807
    v6 := GenS1{F0: 676, F1: 592, F2: 251}
    var v7 string = "a"
    var v8 uint32 = (v3 - v3)
    v9 := GenS2{F0: 1, F1: (^(v1)), F2: 82}
    var v10 int32 = 2147483647
    fmt.Println(v5, v10, v10)
    var v11 string = ("a" + "\n\t")
    v12 := GenS2{F0: 0, F1: (v1 % ((v1) | 1)), F2: -9223372036854775808}
    v13 := GenS2{F0: 1, F1: (-(v1)), F2: 4294967295}
    v2[uint32(v5)] = 127
    v14 := v2[1]
    v1 = -1
    if false {
        fmt.Println(v4)
        v2[v3] = -798
        if true {
            if true {
                var v15 int32 = (v10 & v10)
                if false {
                    var v16 int32 = (v15 / ((v15) | 1))
                    _ = v16
                } else {
                    var v17 bool = true
                    _ = v17
                }
                fmt.Println(v10)
                _ = v15
            }
            var v18 int32 = (v10 | (v10 >> ((4294967295) & 63)))
            _ = v18
        }
    } else {
        var v19 uint32 = v8
        if (2147483647 <= v10) {
            v20 := GenS3{F0: 1, F1: (^(v3))}
            _ = v20
        } else {
            v14 -= 0
            var v21 int32 = -2147483648
            var v22 int = v14
            _ = v21
            _ = v22
        }
        _ = v19
    }
    v23 := GenS1{F0: 0, F1: 18446744073709551615, F2: 0}
    v24 := []int{(v14 - (v14 / ((v14) | 1))), v14, 680}
    if (!(true || (!(2 < 37)))) {
        v14 = (v14 >> ((35) & 63))
        switch 15 {
        case -462:
            fmt.Println(v3, v4, v1)
        }
    }
    if ((0 >= 711) && (v11 == v11)) {
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
