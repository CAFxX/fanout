package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32767, 4294967295, -73}
    var v3 uint32 = 2147483647
    if len(v2) > 0 {
        v2[0] = -591
    }
    v4 := GenS1{F0: 42, F1: 2, F2: -67}
    switch ("a" + ("a" + "foo bar")) {
    case "a":
        var v5 int = 42
        _ = v5
    default:
        v6 := GenS2{F0: true}
        _ = v6
    }
    if true {
        v7 := GenS3{F0: 182}
        _ = v7
    }
    if ((((1e10 * float64((((1e10 + 0.1) / 1.0) + 0.1))) + 3.14159) == 3.14159) || false) {
        v3 -= v3
        if false {
            var v8 uint64 = 42
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int64 = -1
                switch (fi9 % ((721) | 1)) {
                case 42:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v11 float64 = ((((1.5 * 3.14159) - 0.1) - -2.25) / 3.14159)
                    _ = v11
                }
                var v12 int16 = 42
                _ = fi9
                _ = v10
                _ = v12
            }
            _ = v8
        }
        fmt.Println(v3)
    } else {
        v1[(v3 % (((^(v3))) | 1))] = v3
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 uint = 845
    v15 := GenS3{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
