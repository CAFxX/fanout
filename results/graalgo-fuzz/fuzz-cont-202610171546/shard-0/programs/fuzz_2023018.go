package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }

func main() {
    v1 := []uint32{586}
    var v2 uint16 = 2
    v3 := make(map[int]string)
    var v4 uint64 = 49
    v5 := make(map[int]int)
    v2 = (v2 ^ v2)
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = 527
        _ = fi6
    }
    v4 = (v4 - (v4 | 427))
    fmt.Println(v2, v4)
    var v7 uint8 = 42
    var v8 int8 = 0
    v2 -= (^(v2))
    v9 := GenS1{F0: v7, F1: 32768, F2: -1}
    v10 := GenS2{F0: 4294967295, F1: -1, F2: 995}
    switch int(v8) {
    case 335:
        for fi11 := 0; fi11 < 9; fi11++ {
            switch ("a" + "") {
            case "foo bar":
                var v12 uint64 = 4294967295
                _ = v12
            }
            if true {
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 uint8 = (v7 + (v7 / ((v7) | 1)))
                    var v15 uint = 0
                    v16 := GenS3{F0: v2, F1: v14, F2: -87}
                    var v17 float32 = 0.1
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 uint64 = (v4 * 42)
                _ = v18
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                if false {
                    v20 := GenS1{F0: 158, F1: fi19, F2: 695}
                    _ = v20
                }
                _ = fi19
            }
            _ = fi11
        }
    case 2147483647:
        var v21 uint16 = (v2 >> ((0) & 63))
        _ = v21
    }
    var v22 uint = 455
    var v23 uint8 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
