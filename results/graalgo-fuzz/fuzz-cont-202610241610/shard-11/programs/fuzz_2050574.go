package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    var v4 uint = 1
    v4 -= v4
    var v5 uint64 = 42
    fmt.Println(v5, v5)
    v4 = (^(v4))
    v6 := v3[""]
    v5 = 951
    _ = len(v3)
    v5 = (v5 >> ((0) & 63))
    v7 := GenS1{F0: 255}
    var v8 uint8 = 31
    v9 := GenS2{F0: 127, F1: 9223372036854775807, F2: ""}
    if true {
        v10 := GenS1{F0: 42}
        _ = v10
    }
    fmt.Println(v8, v6, v4)
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: uint16(v4)}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            if true {
                var v14 uint = uint(v8)
                var v15 float32 = 3.14159
                _ = v14
                _ = v15
            } else {
                v8 = (^(v8))
                switch ("foo bar" + ("a" + ("hello" + "foo bar"))) {
                case "a":
                    v16 := []int{fi13, -9223372036854775808, (fi13 * (fi13 - (fi13 - -798)))}
                    _ = v16
                case "foo bar":
                    v4 += (v4 << ((v4) & 63))
                case "\n\t":
                    var v17 float32 = 3.14159
                    _ = v17
                }
            }
            fmt.Println(fi13, v4)
            _ = fi13
        }
        switch "" {
        case "\n\t":
            if true {
                var v18 uint32 = (v6 % ((1) | 1))
                v8 = uint8(v18)
                v19 := GenS3{F0: (^(v18)), F1: (v8 / (((v8 ^ v8)) | 1))}
                _ = v18
                _ = v19
            }
        case "foo bar":
            var v20 uint8 = 42
            _ = v20
        default:
            v6 += (v6 / (((v6 >> (((v4 / (((^(v4))) | 1))) & 63))) | 1))
        }
    }
    switch (("foo bar" + ("hello" + (("foo bar" + "foo bar") + "a"))) + "a") {
    case "a":
        if true {
            var v21 uint = (v4 ^ uint(v6))
            _ = v21
        }
    default:
        v22 := GenS2{F0: 100, F1: 128, F2: "hello"}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
