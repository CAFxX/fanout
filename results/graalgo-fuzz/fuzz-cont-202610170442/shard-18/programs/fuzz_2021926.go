package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{4294967295, 4294967295, 1, 2147483647}
    var v3 float32 = ((((1e10 - 3.14159) - -2.25) - 1.5) - (0.0 * (0.1 / 1.0)))
    var v4 uint = 677
    v5 := GenS1{F0: -1}
    if false {
        v4 = 1
    } else {
        v6 := GenS2{F0: 4294967295, F1: 2147483647, F2: 1}
        switch (("hello" + "\n\t") + "\n\t") {
        case "hello":
            if (!(v3 == v3)) {
                v2 = append(v2, 2147483648)
                v3 = v3
                v7 := GenS3{F0: -106, F1: 1}
                _ = v7
            }
        }
        _ = v6
    }
    _ = len(v1)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS3{F0: -537, F1: 0}
        _ = fi8
        _ = v9
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v3)
    v3 = v3
    v1[4294967295] = "hello"
    v10 := GenS4{F0: -1, F1: -1}
    delete(v1, 32767)
    v11 := GenS4{F0: 0, F1: -53}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
