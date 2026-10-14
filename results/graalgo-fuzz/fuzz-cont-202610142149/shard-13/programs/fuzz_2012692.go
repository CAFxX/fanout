package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint32 = 1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: v4, F1: 255, F2: 1}
    var v8 float64 = 0.1
    fmt.Println(v8, v8)
    v9 := []int{260, 1, -9223372036854775808}
    v3 = -1
    for fi10 := 0; fi10 < 14; fi10++ {
        switch (fi10 >> ((42) & 63)) {
        case -782:
            switch fi10 {
            case 1:
                _ = len(v1)
            case 128:
                v8 -= 0.0
            case -9223372036854775808:
                var v11 uint16 = 42
                _ = v11
            default:
                var v12 bool = false
                _ = v12
            }
        default:
            v2 -= (v2 ^ v2)
        }
        switch fi10 {
        case -894:
            v13 := GenS1{F0: 2, F1: 2, F2: v4}
            _ = v13
        case 32768:
            v1[""] = ("foo bar" + "foo bar")
        case 128:
            v4 += v4
        }
        _ = fi10
    }
    var v14 uint64 = 488
    switch -9223372036854775808 {
    case -587:
        var v15 uint = 2
        _ = v15
    case -1:
        var v16 uint32 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v2 = -2147483648
            _ = fi18
        }
        var v19 int64 = (v3 ^ v3)
        var v20 int8 = -76
        _ = fi17
        _ = v19
        _ = v20
    }
    switch 42 {
    case 0:
        var v21 uint64 = v14
        _ = v21
    }
    var v22 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
