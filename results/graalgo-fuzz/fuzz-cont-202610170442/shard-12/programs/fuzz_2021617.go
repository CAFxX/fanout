package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 127
    var v2 float64 = (1.5 - 1.5)
    var v3 int = 2147483647
    v2 = 3.14159
    var v4 int32 = 0
    if false {
        for fi5 := 0; fi5 < 1; fi5++ {
            if false {
                switch int(v1) {
                case 32768:
                    var v6 int = (v3 << ((65535) & 63))
                    _ = v6
                }
                v7 := []int{4294967295, (v3 ^ (v3 - fi5)), (fi5 - v3), (fi5 + -9223372036854775808)}
                v3 = 404
                _ = v7
            }
            _ = fi5
        }
        var v8 uint8 = (^(v1))
        _ = v8
    } else {
        fmt.Println(v3, v2, v4)
    }
    v9 := []int{}
    switch (v3 * (-(v3))) {
    case 1:
        if true {
            v10 := GenS1{F0: -1}
            if (float64((1.5 / (1e10 - 0.1))) != v2) {
                if len(v9) > 0 {
                    v9[0] = v3
                }
                fmt.Println(v4, v2, v2)
                v1 = 130
            } else {
                v9 = append(v9, v3)
            }
            v11 := GenS1{F0: 34}
            _ = v10
            _ = v11
        } else {
            var v12 int64 = 9223372036854775807
            _ = v12
        }
    case 42:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    case 9223372036854775807:
        var v13 float32 = -2.25
        _ = v13
    }
    var v14 int64 = 0
    var v15 uint64 = 18446744073709551615
    v14 -= v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
