package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (true && (1.5 != (0.1 - (1.5 + (-2.25 + 1.5)))))
    v3 := []uint32{65535, 594}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    var v6 string = ""
    v7 := GenS1{F0: 4294967295, F1: -1}
    v6 = "\n\t"
    if (-1 < -1) {
        v8 := GenS1{F0: 2, F1: 634}
        _ = v8
    } else {
        v9 := v5[0]
        fmt.Println(v6, v6, v9)
        _ = v9
    }
    fmt.Println(v2, v2)
    v10 := v5[298]
    if (v6 < "foo bar") {
        delete(v4, v10)
    }
    v10 = v10
    fmt.Println(v2, v2)
    v11 := GenS2{F0: 128, F1: -985, F2: -9223372036854775808}
    var v12 string = "a"
    for fi13 := 0; fi13 < 20; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    if v2 {
        v6 = (v12 + "foo bar")
        if (!v2) {
            if ((639 > 42) && true) {
                switch "foo bar" {
                case "foo bar":
                    var v14 uint16 = 0
                    _ = v14
                default:
                    v6 = v12
                }
                v15 := GenS3{F0: -119}
                v16 := GenS4{F0: (1.5 == 0.1), F1: -2147483648}
                _ = v15
                _ = v16
            }
        } else {
            v17 := GenS3{F0: 42}
            fmt.Println(v6, v2, v12)
            fmt.Println(v6, v2)
            _ = v17
        }
    }
    v18 := GenS1{F0: 255, F1: -1}
    var v19 string = v6
    var v20 uint8 = 255
    v21 := GenS5{F0: 0}
    v22 := GenS3{F0: -106}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
