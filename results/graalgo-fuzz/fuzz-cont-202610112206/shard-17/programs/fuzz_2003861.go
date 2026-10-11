package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    var v1 bool = (-1 > -1)
    var v2 uint8 = 128
    var v3 int16 = 42
    var v4 int = 128
    var v5 uint = 4294967295
    switch ("hello" + ("a" + ((("" + (("" + "hello") + "")) + ("" + "")) + "\n\t"))) {
    case "a":
        v6 := []int{}
        _ = v6
    }
    v7 := GenS1{F0: v4}
    v8 := GenS1{F0: v4}
    var v9 uint8 = (v2 >> (((v5 << (((v5 + 42)) & 63))) & 63))
    v10 := GenS1{F0: (v4 ^ v4)}
    if (!(1 <= 2147483647)) {
        var v11 uint = (v5 + v5)
        v12 := GenS1{F0: int(v11)}
        if (-670 <= 2147483647) {
            var v13 int8 = 1
            _ = v13
        } else {
            if v1 {
                v14 := GenS1{F0: (v4 ^ v4)}
                var v15 uint = (^(v5))
                _ = v14
                _ = v15
            } else {
                if (0 <= -1) {
                    v5 -= v5
                } else {
                    var v16 int8 = -33
                    _ = v16
                }
                var v17 uint32 = 2
                _ = v17
            }
            v18 := GenS2{F0: (v3 >> ((v5) & 63))}
            v19 := GenS1{F0: v4}
            _ = v18
            _ = v19
        }
        _ = v11
        _ = v12
    }
    v20 := []int{(v4 + (v4 % ((-1) | 1))), (v4 - v4), v4, v4}
    if true {
        var v21 int8 = -1
        v5 = (v5 + v5)
        _ = v21
    }
    v22 := GenS3{F0: 42, F1: false}
    var v23 uint64 = 2
    var v24 uint64 = (v23 / ((v23) | 1))
    v20 = append(v20, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
