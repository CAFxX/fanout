package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    v4 := GenS1{F0: 615, F1: 0}
    v3 += (v3 | (v3 >> ((18446744073709551615) & 63)))
    v5 := GenS2{F0: -9223372036854775808}
    fmt.Println(v3)
    v6 := GenS1{F0: 32767, F1: 127}
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            var v8 bool = false
            var v9 float32 = (0.0 * 0.0)
            _ = v8
            _ = v9
        } else {
            switch "" {
            case "hello":
                v10 := GenS2{F0: fi7}
                _ = v10
            default:
                var v11 uint16 = 2
                _ = v11
            }
            v3 += -1
            v3 += v3
        }
        fmt.Println(fi7, v3)
        _ = fi7
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v3 = v3
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint = 344
            var v15 int8 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v3 = v3
    v3 = 0
    var v16 uint8 = 247
    if true {
        v16 = 42
    } else {
        var v17 int16 = 32767
        var v18 uint8 = 127
        if false {
            v19 := []int{297, 535, 0, 4294967295}
            v17 = v17
            _ = v19
        }
        _ = v17
        _ = v18
    }
    if false {
        v20 := GenS1{F0: 1, F1: v16}
        _ = v20
    }
    v21 := GenS2{F0: -1}
    var v22 uint16 = 596
    switch "foo bar" {
    case "a":
        v23 := GenS3{F0: -817}
        _ = v23
    case "hello":
        var v24 uint = 943
        _ = v24
    case "foo bar":
        var v25 uint8 = 1
        _ = v25
    default:
        var v26 uint = 255
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
