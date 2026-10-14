package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{0}
    var v2 int64 = -77
    v3 := make(map[string]string)
    var v4 int64 = (v2 - (v2 - v2))
    v5 := make(map[uint32]uint32)
    var v6 string = ("\n\t" + ("" + "hello"))
    fmt.Println(v4, v2, v2)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v6, v6)
    for fi9 := 0; fi9 < 1; fi9++ {
        v2 = (v4 ^ (v2 ^ v2))
        _ = fi9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v2 = -110
        for fi11 := 0; fi11 < 13; fi11++ {
            if true {
                v12 := GenS1{F0: v2, F1: 0}
                if (127 < (fi10 >> ((598) & 63))) {
                    v13 := GenS1{F0: (v4 - (v4 & v4)), F1: 18446744073709551615}
                    _ = v13
                } else {
                    v6 = v6
                }
                switch v6 {
                case "":
                    v6 = v6
                case "foo bar":
                    v2 = v4
                case "a":
                    var v14 uint8 = 255
                    _ = v14
                default:
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                }
                v15 := GenS2{F0: 100}
                switch ("\n\t" + (v6 + "foo bar")) {
                case "foo bar":
                    var v16 int64 = 1
                    _ = v16
                }
                _ = v12
                _ = v15
            }
            _ = fi11
        }
        v3["a"] = "a"
        _ = fi10
    }
    var v17 float32 = 0.1
    _ = len(v1)
    switch "a" {
    case "a":
        var v18 int16 = 0
        _ = v18
    case "hello":
        var v19 int8 = -128
        _ = v19
    case "foo bar":
        var v20 uint16 = 65535
        _ = v20
    default:
        var v21 bool = false
        _ = v21
    }
    v2 = v2
    for fi22 := 0; fi22 < 5; fi22++ {
        v4 = (v4 << ((255) & 63))
        _ = fi22
    }
    fmt.Println(v17)
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS3{F0: v4}
            fmt.Println(v4)
            v2 -= (v4 & v4)
            _ = fi23
            _ = v24
        }
        v25 := v3[v6]
        var v26 uint8 = 42
        _ = v25
        _ = v26
    } else {
        v6 = v6
        v27 := GenS2{F0: -75}
        _ = v27
    }
    var v28 float32 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
