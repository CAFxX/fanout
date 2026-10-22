package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }

func main() {
    v1 := []uint32{905}
    var v2 int16 = 93
    var v3 bool = false
    v4 := []string{}
    var v5 uint8 = 1
    v6 := make(map[uint32]bool)
    switch 9223372036854775807 {
    case -2147483648:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 107:
        if (v3 && v3) {
            v1 = append(v1, 2147483647)
        } else {
            if (true && (42 == 18446744073709551615)) {
                v7 := GenS1{F0: 128, F1: -762, F2: 1}
                v5 -= v5
                if v3 {
                    var v8 bool = v3
                    if (0.0 <= -2.25) {
                        v3 = v8
                        if v8 {
                            var v9 uint64 = 18446744073709551615
                            v5 = (v5 >> ((613) & 63))
                            _ = v9
                        } else {
                            fmt.Println(v3)
                        }
                    } else {
                        v5 = 110
                        v2 = v2
                        var v10 uint16 = 697
                        _ = v10
                    }
                    _ = v8
                }
                _ = v7
            }
        }
    }
    v3 = ((true && (42 == -92)) && (v2 < v2))
    v11 := GenS1{F0: 32768, F1: 32768, F2: 813}
    v3 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (105 < 4294967295) {
        _ = len(v1)
        var v12 bool = (!v3)
        _ = v12
    } else {
        if false {
            fmt.Println(v3)
            v4 = append(v4, "a")
        }
        v3 = v3
    }
    fmt.Println(v3)
    var v13 bool = true
    var v14 uint32 = 4294967295
    var v15 float32 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
