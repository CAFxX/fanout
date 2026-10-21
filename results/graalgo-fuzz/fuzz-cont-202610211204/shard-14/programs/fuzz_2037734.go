package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := []uint32{42, 185, 42}
    var v2 float32 = (float32(3.14159) * 3.14159)
    var v3 uint = 65535
    var v4 string = "\n\t"
    v5 := make(map[string]int)
    var v6 int32 = 0
    v1 = append(v1, 0)
    v6 += v6
    if false {
        fmt.Println(v6, v4)
        v2 = 0.1
    } else {
        switch (v4 + "hello") {
        case "":
            fmt.Println(v2, v2, v2)
        case "a":
            v7 := GenS1{F0: -1, F1: 2, F2: 4294967295}
            _ = v7
        }
    }
    v4 = "hello"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v6)
    for fi8 := 0; fi8 < 15; fi8++ {
        v5["foo bar"] = fi8
        if false {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint16 = uint16(v3)
                _ = fi9
                _ = v10
            }
        }
        v11 := GenS2{F0: (v4 + (v4 + v4)), F1: 36}
        _ = fi8
        _ = v11
    }
    var v12 uint16 = 505
    var v13 string = "\n\t"
    switch v4 {
    case "":
        v14 := GenS2{F0: v13, F1: 18446744073709551615}
        _ = v14
    case "a":
        if (0 >= (v12 & (^(v12)))) {
            fmt.Println(v3, v13)
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint32 = 618
                v12 = 2
                var v17 int8 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case "hello":
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 float64 = 1.5
            var v20 int = 258
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        v21 := GenS1{F0: -1, F1: 0, F2: 914}
        _ = v21
    }
    v2 = 0.0
    var v22 int64 = 42
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 float32 = v2
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        fmt.Println(v4)
        var v26 int = fi25
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
