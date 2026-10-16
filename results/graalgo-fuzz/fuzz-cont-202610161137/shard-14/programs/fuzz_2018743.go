package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 uint32 = 65535
    v4 := []uint32{v3, v3, (v3 | 65535), v3}
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int8 = -54
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v6)
        _ = fi5
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            _ = len(v2)
            v3 = v3
            v8 := GenS1{F0: -1, F1: false, F2: 1}
            _ = fi7
            _ = v8
        }
        v9 := GenS2{F0: 993, F1: 734, F2: 0}
        _ = v9
    }
    v3 -= v3
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = 42
        _ = len(v2)
        v11 -= 9
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "":
        v12 := v2[("a" + "")]
        _ = v12
    case "hello":
        v3 = 65535
    }
    v3 += v3
    v2["foo bar"] = ("\n\t" + "")
    var v13 int64 = 42
    var v14 uint16 = 65535
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v3, v3)
        v16 := v2["a"]
        delete(v2, v16)
        _ = fi15
        _ = v16
    }
    if (4294967295 != uint64(v3)) {
        v17 := GenS2{F0: v13, F1: 970, F2: -603}
        _ = v17
    }
    v14 = (v14 | v14)
    switch -9223372036854775808 {
    case 224:
        v18 := GenS2{F0: v13, F1: 255, F2: 0}
        _ = v18
    case 2147483647:
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(fi19, fi19, fi19)
            var v20 int32 = -1
            _ = fi19
            _ = v20
        }
    }
    if false {
        v21 := v1[(v3 ^ v3)]
        _ = v21
    } else {
        v13 = (v13 - v13)
        fmt.Println(v13, v3, v3)
        v22 := GenS3{F0: -938, F1: 1000, F2: 128}
        _ = v22
    }
    fmt.Println(v3, v3)
    if len(v4) > 0 {
        v4[0] = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
