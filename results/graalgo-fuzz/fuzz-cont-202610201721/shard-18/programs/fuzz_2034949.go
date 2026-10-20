package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int; F2 string }

func main() {
    v1 := []uint32{1, 2}
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    var v4 uint = 42
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    _ = len(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    fmt.Println(v4, v2)
    var v7 int32 = -2147483648
    _ = len(v1)
    if false {
        if true {
            v7 = (^(v7))
            var v8 int16 = 1
            switch 127 {
            case 2147483647:
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: (((0.1 < 3.14159) || true) && false)}
                    if (v4 > (v4 | 42)) {
                        var v11 int = fi9
                        _ = v11
                    } else {
                        fmt.Println(v6, v8, v8)
                    }
                    _ = fi9
                    _ = v10
                }
            case 1:
                for fi12 := 0; fi12 < 4; fi12++ {
                    if true {
                        var v13 uint64 = 2
                        fmt.Println(v4, v8)
                        var v14 int32 = v7
                        _ = v13
                        _ = v14
                    }
                    var v15 uint8 = 2
                    _ = fi12
                    _ = v15
                }
            }
            _ = v8
        } else {
            v2 += (v2 ^ v2)
        }
        v16 := v3[0]
        v16 -= 4294967295
        _ = v16
    } else {
        var v17 uint64 = 18446744073709551615
        v1 = append(v1, 65535)
        var v18 string = v6
        _ = v17
        _ = v18
    }
    fmt.Println(v7)
    v19 := GenS1{F0: (v2 <= v2)}
    if true {
        switch (v6 + (v6 + v6)) {
        case "hello":
            var v20 uint = v4
            _ = v20
        default:
            switch 9223372036854775807 {
            case 153:
                var v21 uint64 = 212
                _ = v21
            case 2147483647:
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS2{F0: 468, F1: 18446744073709551615, F2: false}
                    _ = fi22
                    _ = v23
                }
            }
        }
        v24 := GenS3{F0: 0, F1: 496, F2: ("\n\t" + v6)}
        _ = v24
    } else {
        v7 = v7
        v2 -= 53
    }
    v25 := GenS1{F0: false}
    var v26 int16 = 1000
    v27 := GenS2{F0: v7, F1: 566, F2: (463 == uint64(v4))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
