package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    v1 := []uint32{1, 0, 69, 240}
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    v4 = ""
    var v5 uint16 = 65535
    _ = len(v1)
    v4 = v4
    var v6 uint8 = 128
    v6 = 0
    v7 := GenS1{F0: 4294967295, F1: v6}
    fmt.Println(v5)
    v8 := GenS2{F0: 0, F1: 341}
    if (true || (610 > 127)) {
        switch -1 {
        case -9223372036854775808:
            switch 420 {
            case 9223372036854775807:
                switch 32768 {
                case 42:
                    v6 = 0
                }
            case 0:
                var v9 int32 = 2147483647
                _ = v9
            default:
                var v10 uint16 = 606
                _ = v10
            }
        case -611:
            v11 := GenS2{F0: 859, F1: 4294967295}
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            v4 = (((v4 + v4) + "a") + ("hello" + v4))
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v5)
                var v14 float64 = 0.0
                fmt.Println(v4)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        v15 := GenS3{F0: 1, F1: 594}
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
