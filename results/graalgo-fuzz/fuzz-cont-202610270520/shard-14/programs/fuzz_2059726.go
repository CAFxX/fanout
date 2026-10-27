package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 384
    var v3 int16 = (v1 >> ((0) & 63))
    v4 := make(map[int]string)
    var v5 int64 = -1
    var v6 uint = 391
    switch ("\n\t" + "") {
    case "":
        v7 := GenS1{F0: 489, F1: 660, F2: v3}
        _ = v7
    case "\n\t":
        var v8 uint = v6
        _ = v8
    case "a":
        var v9 int16 = v1
        _ = v9
    }
    var v10 int8 = 127
    fmt.Println(v2, v1)
    if (2147483648 == 2) {
        if true {
            var v11 int16 = (v3 / ((v3) | 1))
            _ = v11
        }
        v12 := v4[-104]
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{int(v5), (-(fi13))}
        for fi15 := 0; fi15 < 4; fi15++ {
            fmt.Println(fi13, v6, v3)
            if (0 > 18) {
                var v16 int64 = 9223372036854775807
                switch ("a" + "\n\t") {
                case "a":
                    if false {
                        v1 = (^(v3))
                    }
                case "foo bar":
                    v1 += (v1 % (((v3 % (((v3 >> (((v6 ^ (^(v6)))) & 63))) | 1))) | 1))
                case "hello":
                    v16 = (v5 << (((v6 - 274)) & 63))
                default:
                    fmt.Println(v10, v16)
                }
                v17 := GenS2{F0: 351, F1: (-(v3))}
                _ = v16
                _ = v17
            } else {
                if (2147483648 >= 42) {
                    fmt.Println(v10, v6)
                }
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    var v18 uint64 = 18446744073709551615
    if false {
        delete(v4, 32767)
        v10 -= (v10 & (-(v10)))
    }
    v1 = (v1 | 716)
    v6 -= (v6 * (v6 % (((v6 % (((v6 ^ (v6 / ((uint(v18)) | 1)))) | 1))) | 1)))
    var v19 uint64 = (v18 >> (((v6 - (^(v6)))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
