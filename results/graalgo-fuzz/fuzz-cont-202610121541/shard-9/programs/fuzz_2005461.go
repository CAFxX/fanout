package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    var v3 uint = 65535
    v4 := make(map[uint32]bool)
    switch "a" {
    case "foo bar":
        for fi5 := 0; fi5 < 7; fi5++ {
            v3 = (v3 & (v3 | (v3 * 2)))
            switch (fi5 % (((fi5 - fi5)) | 1)) {
            case 0:
                v6 := GenS1{F0: -71, F1: v2, F2: "a"}
                _ = v6
            case -900:
                var v7 bool = false
                _ = v7
            case 42:
                v8 := v1[0]
                _ = v8
            default:
                var v9 int64 = -1
                _ = v9
            }
            v10 := GenS1{F0: 39, F1: (true && (152 <= 157)), F2: ("a" + "hello")}
            _ = fi5
            _ = v10
        }
    }
    switch "" {
    case "hello":
        v3 -= v3
    default:
        v11 := []int{}
        _ = v11
    }
    v12 := []int{}
    var v13 int16 = 42
    fmt.Println(v3)
    if len(v12) > 0 {
        v12[0] = int(v13)
    }
    v2 = (821 >= 123)
    var v14 int32 = 2147483647
    v4[2147483647] = true
    var v15 int16 = v13
    v2 = v2
    v16 := GenS1{F0: 42, F1: (65535 < 65535), F2: "a"}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
