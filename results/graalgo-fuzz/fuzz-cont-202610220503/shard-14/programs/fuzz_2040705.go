package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := []int{32768, 32768, 4294967295, -2147483648}
    var v2 string = "\n\t"
    var v3 uint64 = 4294967295
    var v4 int8 = 1
    var v5 bool = true
    v6 := make(map[uint32]string)
    if v5 {
        v7 := GenS1{F0: v5, F1: (v4 << ((42) & 63)), F2: 4294967295}
        _ = v7
    }
    var v8 uint64 = v3
    var v9 uint64 = 1
    if false {
        if false {
            if (!false) {
                if len(v1) > 0 {
                    v1[0] = 42
                }
            }
            switch 411 {
            case -1:
                var v10 int16 = -1
                _ = v10
            }
        }
        var v11 float64 = 1.5
        _ = v11
    }
    var v12 uint32 = 0
    v2 = (v2 + v2)
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 float64 = 0.1
        _ = fi13
        _ = v14
    }
    v3 = v3
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: v5, F1: 0, F2: uint32(v9)}
        _ = fi15
        _ = v16
    }
    var v17 int8 = (v4 + (v4 * -1))
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int = (fi18 << ((266) & 63))
        v17 = 42
        if v5 {
            var v20 float32 = float32(0.1)
            v21 := GenS1{F0: v5, F1: 0, F2: v12}
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v22 := GenS1{F0: true, F1: (v4 & (v17 ^ (-(v17)))), F2: (v12 ^ v12)}
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: 0}
        var v25 uint8 = 5
        switch v2 {
        case "\n\t":
            for fi26 := 0; fi26 < 11; fi26++ {
                fmt.Println(v25, fi26, v4)
                v5 = v5
                fmt.Println(v2, fi26, v3)
                _ = fi26
            }
        case "a":
            var v27 int32 = 2147483647
            _ = v27
        case "foo bar":
            var v28 int16 = int16(v17)
            _ = v28
        default:
            fmt.Println(v2, v2)
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v29 int = 128
    if v5 {
        switch v2 {
        case "a":
            v9 += v8
        case "hello":
            var v30 uint64 = v9
            _ = v30
        default:
            v8 += v9
        }
        switch (v29 / ((544) | 1)) {
        case 4294967295:
            _ = len(v6)
        default:
            _ = len(v1)
        }
        switch 127 {
        case -9223372036854775808:
            v2 = (v2 + v2)
        case 32767:
            if (v5 && (0.0 != float32(0.1))) {
                v31 := GenS3{F0: 351, F1: (^(v12))}
                v12 = 42
                var v32 int32 = -2147483648
                _ = v31
                _ = v32
            }
        }
    } else {
        var v33 uint32 = v12
        var v34 int64 = 1
        _ = v33
        _ = v34
    }
    var v35 int = (v29 + -751)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
