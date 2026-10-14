package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 int16 }

func main() {
    var v1 uint16 = 0
    var v2 string = (("\n\t" + "a") + (("" + "a") + "foo bar"))
    var v3 bool = (!false)
    v4 := make(map[uint32]int)
    var v5 float64 = ((((0.1 + float64(3.14159)) - 3.14159) + 1.5) + ((float64(((float32(1e10) * 1.5) - 0.0)) * 0.1) + 3.14159))
    var v6 uint64 = 2
    if v3 {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int64 = -71
            _ = fi7
            _ = v8
        }
        var v9 int16 = 32767
        v9 += v9
        _ = v9
    }
    var v10 uint = uint(v6)
    v10 += 330
    v11 := GenS1{F0: (v2 + v2), F1: 745, F2: -32768}
    v2 = (v2 + v2)
    v12 := GenS1{F0: v2, F1: 0, F2: 540}
    if false {
        v13 := v4[42]
        v14 := GenS1{F0: v2, F1: 2147483647, F2: -1}
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: "a", F1: int32(v6), F2: 42}
    for fi16 := 0; fi16 < 8; fi16++ {
        switch (-(fi16)) {
        case 127:
            fmt.Println(v10, v2, fi16)
        default:
            v1 += uint16(v6)
        }
        var v17 bool = v3
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 254
    switch v2 {
    case "hello":
        var v19 uint16 = v1
        _ = v19
    case "\n\t":
        var v20 uint32 = 1
        _ = v20
    case "foo bar":
        fmt.Println(v5)
    default:
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = int32(v1)
            _ = fi21
            _ = v22
        }
    }
    switch "\n\t" {
    case "":
        v10 = 4294967295
    default:
        fmt.Println(v2, v18)
    }
    v23 := GenS2{F0: -662, F1: v3, F2: 42}
    var v24 uint64 = 2
    v18 -= (v18 | 527)
    fmt.Println(v1)
    for fi25 := 0; fi25 < 18; fi25++ {
        v3 = (v3 && v3)
        _ = len(v4)
        v26 := GenS3{F0: -844, F1: -1, F2: 0}
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
