package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    v1 := []int{42, -122, 0, 0}
    var v2 int16 = -1
    var v3 uint32 = 733
    var v4 uint8 = 1
    var v5 int = 127
    v6 := make(map[string]string)
    fmt.Println(v3)
    v7 := GenS1{F0: -1, F1: ""}
    var v8 uint32 = (v3 & v3)
    v9 := v6["foo bar"]
    var v10 int64 = 9223372036854775807
    v10 = -1
    var v11 uint64 = 18446744073709551615
    v12 := GenS1{F0: int8(v10), F1: v9}
    if true {
        if (100 == 100) {
            var v13 int64 = (v10 ^ (v10 ^ (v10 ^ (v10 % (((v10 ^ (v10 + v10))) | 1)))))
            var v14 int = (v5 << ((912) & 63))
            for fi15 := 0; fi15 < 1; fi15++ {
                v5 = 9223372036854775807
                v4 += 1
                _ = fi15
            }
            _ = v13
            _ = v14
        } else {
            v16 := GenS2{F0: (v2 & v2), F1: 850, F2: 32767}
            v2 = (v2 / ((v2) | 1))
            _ = v16
        }
    } else {
        v17 := GenS3{F0: v9, F1: 0}
        v3 = v3
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
