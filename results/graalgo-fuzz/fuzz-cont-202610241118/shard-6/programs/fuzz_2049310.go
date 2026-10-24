package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 float64 = (((float64(1.5) - 1e10) / 3.14159) / 0.1)
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 0, F1: 0, F2: 0}
    v6 := GenS1{F0: 32767, F1: 42, F2: 100}
    v7 := []int{2147483647, -881, -703, -1}
    if false {
        v8 := GenS2{F0: 1, F1: 4294967295}
        _ = v8
    }
    _ = len(v4)
    _ = len(v4)
    var v9 int32 = 1
    v9 -= v9
    switch ("\n\t" + "a") {
    case "":
        v9 = -1
    case "hello":
        fmt.Println(v9)
    default:
        fmt.Println(v9)
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        switch (fi10 << ((832) & 63)) {
        case 683:
            switch fi10 {
            case 2147483647:
                var v11 int64 = 793
                _ = v11
            case -2147483648:
                v2[fi10] = 2147483648
            }
        case 534:
            v3 = (v3 * v3)
        default:
            v12 := GenS3{F0: ""}
            _ = v12
        }
        _ = fi10
    }
    if (32767 >= 4294967295) {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint8 = 1
            var v15 int8 = 100
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    _ = len(v7)
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            var v17 float32 = 3.14159
            _ = len(v4)
            fmt.Println(v17, fi16, v9)
            _ = v17
        } else {
            var v18 int32 = v9
            _ = v18
        }
        _ = fi16
    }
    var v19 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
