package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 2
    var v3 uint8 = 1
    v4 := make(map[string]string)
    v2 -= v2
    v5 := GenS1{F0: 127, F1: 421}
    v6 := []int{}
    if false {
        v1 -= v1
        for fi7 := 0; fi7 < 3; fi7++ {
            v2 += v2
            var v8 int = fi7
            _ = fi7
            _ = v8
        }
    }
    var v9 bool = true
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 bool = false
        if v11 {
            var v12 float32 = 1.5
            var v13 uint8 = v3
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint = 42
    var v15 uint64 = (^(v2))
    switch "a" {
    case "a":
        if ((65535 != 0) || (!v9)) {
            var v16 int64 = 359
            v9 = v9
            v3 = 127
            _ = v16
        }
    default:
        var v17 uint = (^(v14))
        _ = v17
    }
    var v18 uint = 2
    _ = len(v6)
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint16 = 42
        var v21 int64 = -1000
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(v20, v14, v14)
            _ = fi22
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v1 -= (v1 >> ((0) & 63))
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v18, v2)
            _ = fi24
        }
        if ((!(0.0 == ((0.1 / 1.0) * 0.0))) && v9) {
            v3 -= uint8(v2)
            fmt.Println(v1, v15)
            v1 += -2147483648
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            _ = len(v6)
            v9 = v9
            if (v14 > (v18 << (((v14 / ((255) | 1))) & 63))) {
                fmt.Println(v9)
                v6 = append(v6, fi26)
                var v27 string = ""
                _ = v27
            }
            v28 := GenS1{F0: 1, F1: 65535}
            _ = fi26
            _ = v28
        }
        var v29 uint64 = 109
        var v30 float32 = (float32(3.14159) * 0.1)
        _ = fi25
        _ = v29
        _ = v30
    }
    var v31 int64 = -1
    v9 = (v31 >= v31)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
