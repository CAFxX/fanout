package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    var v3 int8 = -51
    var v4 string = ("hello" + "a")
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 982}
    v7 := []int{-745, 32768, 1, 164}
    fmt.Println(v3)
    switch "\n\t" {
    case "hello":
        if (!(false || (765 >= 2147483647))) {
            v1 -= v1
        }
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int16 = 1000
        if (883 < 65535) {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int64 = v1
                v12 := GenS2{F0: 42, F1: 128, F2: 127}
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                switch "a" {
                case "foo bar":
                    var v14 int32 = 822
                    _ = v14
                case "a":
                    switch 0 {
                    case 42:
                        v4 = v4
                    default:
                        v15 := GenS3{F0: 2, F1: v1, F2: 473}
                        _ = v15
                    }
                }
                var v16 int = (-(fi13))
                if ((false && false) || (87 <= 181)) {
                    v4 = v4
                    if len(v7) > 0 {
                        v7[0] = (-(v16))
                    }
                    fmt.Println(fi13)
                }
                _ = fi13
                _ = v16
            }
            var v17 float32 = 0.0
            _ = v17
        }
        switch (v4 + v4) {
        case "a":
            fmt.Println(v4)
        default:
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS2{F0: 1, F1: 42, F2: 127}
                var v20 string = v4
                var v21 uint16 = 1
                var v22 uint32 = 663
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = fi8
        _ = v9
    }
    _ = len(v5)
    v1 = (v1 * 1)
    var v23 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
