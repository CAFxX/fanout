package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0() uint64 {
    if true {
        var v6 int32 = -188
        v7 := []int{-1, 135}
        _ = v6
        _ = v7
    }
    var v8 float64 = float64((0.1 - 0.1))
    _ = v8
    return 4294967295
}

func main() {
    var v1 int64 = 53
    v2 := []string{"\n\t", "hello", "", ""}
    v3 := make(map[uint32]bool)
    var v4 int = 2147483647
    var v5 uint64 = 4294967295
    var v9 float64 = 0.0
    v5 = (v5 << ((585) & 63))
    fmt.Println(v9, v1, v1)
    fmt.Println(v4)
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS1{F0: (true || false), F1: 260}
            v9 = (v9 - v9)
            var v13 int8 = -1
            var v14 string = ""
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: ((false || true) && ((32768 < (fi10 / ((int(v5)) | 1))) && true)), F1: 4294967295}
        _ = fi10
        _ = v15
    }
    v2 = append(v2, ("foo bar" + "foo bar"))
    var v16 int64 = 42
    fmt.Println(v1, v16)
    v17 := GenS2{F0: (false || (1 <= 32767))}
    switch v4 {
    case 84:
        if true {
            var v18 uint64 = (v5 + 506)
            var v19 int16 = -32768
            _ = v18
            _ = v19
        } else {
            v3[2] = false
            v20 := GenS3{F0: 128, F1: 325, F2: (v16 < (^(v1)))}
            _ = v20
        }
    case 127:
        v4 += v4
    }
    v2 = append(v2, "foo bar")
    var v21 float32 = 1.5
    var v22 int8 = 116
    var v23 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
