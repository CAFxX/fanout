package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint = 42
    var v2 uint16 = 65535
    v3 := make(map[uint32]string)
    v4 := []int{-242, 4294967295, 14, -115}
    v5 := GenS1{F0: 65535, F1: ("a" > "")}
    if (9223372036854775807 > 1) {
        v2 = (v2 * v2)
        v1 = (v1 + v1)
        for fi6 := 0; fi6 < 10; fi6++ {
            v1 = 1
            _ = fi6
        }
    }
    switch -9223372036854775808 {
    case 42:
        v1 = v1
    }
    v7 := GenS1{F0: 283, F1: true}
    v1 = 18446744073709551615
    var v8 int = 2147483647
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 = 4294967295
        for fi10 := 0; fi10 < 20; fi10++ {
            if true {
                switch (v8 % ((32768) | 1)) {
                case -290:
                    if len(v4) > 0 {
                        v4[0] = (fi9 << ((v1) & 63))
                    }
                case 4294967295:
                    fmt.Println(v8, v1)
                }
                var v11 uint32 = 42
                _ = v11
            } else {
                v8 += 32767
            }
            fmt.Println(v8)
            _ = fi10
        }
        _ = fi9
    }
    var v12 uint8 = 128
    v2 = (^(v2))
    v13 := GenS1{F0: 230, F1: false}
    var v14 int32 = 0
    v15 := GenS2{F0: 100}
    v16 := GenS2{F0: 27}
    for fi17 := 0; fi17 < 7; fi17++ {
        v2 += (v2 % ((v2) | 1))
        _ = fi17
    }
    _ = len(v3)
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
