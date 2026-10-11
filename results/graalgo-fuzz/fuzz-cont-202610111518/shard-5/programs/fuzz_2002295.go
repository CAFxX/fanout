package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 string = "\n\t"
    v2 := []int{32768, -873, -9223372036854775808, -1}
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 uint8 = 255
    v6 := make(map[int]bool)
    v5 = v5
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float64 = float64((3.14159 + float32(1e10)))
        if false {
            var v9 int64 = -93
            for fi10 := 0; fi10 < 10; fi10++ {
                v1 = v1
                _ = fi10
            }
            _ = v9
        } else {
            fmt.Println(v5, fi7)
            v8 = (v8 / (0.0 - v8))
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 44, F1: false}
    v5 = v5
    switch v1 {
    case "a":
        if len(v2) > 0 {
            v2[0] = 581
        }
    case "foo bar":
        var v12 int32 = 42
        _ = v12
    case "":
        v13 := GenS2{F0: 65535, F1: (((3.14159 + float32(0.1)) / 1.5) == ((3.14159 - 3.14159) * float32(1.5))), F2: 42}
        _ = v13
    }
    v5 = 1
    v1 = ""
    v4[0] = 2
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 bool = false
        switch (fi14 * (fi14 / (((-(fi14))) | 1))) {
        case 0:
            _ = len(v6)
        case 32768:
            for fi16 := 0; fi16 < 17; fi16++ {
                v2 = append(v2, 32768)
                v15 = v15
                _ = fi16
            }
        case 9223372036854775807:
            var v17 int8 = -1
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    switch -693 {
    case 1:
        fmt.Println(v5)
    case 32768:
        for fi18 := 0; fi18 < 1; fi18++ {
            v5 = 5
            var v19 uint64 = 319
            var v20 int8 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v5 = (v5 | 255)
    var v21 uint = 263
    v22 := GenS3{F0: (v5 - v5), F1: 1}
    _ = len(v4)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
