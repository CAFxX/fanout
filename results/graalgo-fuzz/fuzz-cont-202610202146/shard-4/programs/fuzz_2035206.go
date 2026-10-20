package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = "a"
    var v2 int32 = 39
    v3 := make(map[uint32]int)
    var v4 uint32 = 65535
    var v5 bool = true
    v4 = (v4 + (v4 ^ (^(v4))))
    fmt.Println(v2, v1)
    var v6 int16 = int16(v4)
    v5 = (9223372036854775807 == 4294967295)
    if v5 {
        fmt.Println(v2, v2, v2)
    } else {
        var v7 uint32 = (^(v4))
        _ = v7
    }
    var v8 int64 = int64(v4)
    if (float64(0.0) == ((float64((-2.25 / 3.14159)) * 0.0) - 1.5)) {
        delete(v3, v4)
        if (1 == 42) {
            v9 := GenS1{F0: 945, F1: 32767}
            if v5 {
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch 128 {
                    case 42:
                        var v11 int = (fi10 % ((fi10) | 1))
                        _ = v11
                    }
                    _ = fi10
                }
                var v12 int16 = (-(v6))
                _ = v12
            } else {
                v5 = true
                v13 := []int{9223372036854775807, 51, 42}
                var v14 int64 = v8
                _ = v13
                _ = v14
            }
            fmt.Println(v2, v5, v6)
            _ = v9
        } else {
            var v15 uint64 = 1
            _ = v15
        }
        v16 := []int{42}
        _ = v16
    } else {
        v17 := GenS2{F0: (^(v4))}
        _ = v17
    }
    var v18 int = int(v2)
    switch (v18 % (((v18 | v18)) | 1)) {
    case 2147483647:
        v19 := []int{}
        _ = v19
    }
    v20 := []int{v18, (v18 * (v18 ^ int(v2))), int(v4), (v18 % (((v18 ^ 239)) | 1))}
    v21 := GenS1{F0: v18, F1: 65535}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
