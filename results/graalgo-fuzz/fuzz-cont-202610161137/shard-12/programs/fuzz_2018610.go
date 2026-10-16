package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 string; F2 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int64 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    var v3 int32 = 818
    v4 := make(map[uint32]bool)
    if ((false || false) || false) {
        var v5 int16 = 42
        var v6 int32 = (v3 ^ v3)
        _ = v5
        _ = v6
    } else {
        var v7 uint = 4294967295
        v8 := GenS1{F0: 795, F1: -32768, F2: 168}
        _ = v7
        _ = v8
    }
    v1 = v1
    v9 := GenS2{F0: "", F1: ("a" + ("\n\t" + "foo bar")), F2: 239}
    if ((v3 | -120) == 330) {
        _ = len(v2)
        for fi10 := 0; fi10 < 13; fi10++ {
            switch "\n\t" {
            case "foo bar":
                if true {
                    v11 := GenS2{F0: (v1 + v1), F1: v1, F2: 978}
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(fi12, v3)
                var v13 string = "foo bar"
                _ = fi12
                _ = v13
            }
            fmt.Println(v3)
            _ = fi10
        }
        v14 := GenS3{F0: 42, F1: 368, F2: 42}
        _ = v14
    }
    v15 := GenS4{F0: 0, F1: 32767, F2: -1}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 bool = true
        fmt.Println(v17, v3, v3)
        var v18 uint32 = 2147483647
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: 2, F1: 42, F2: 0}
    if true {
        var v20 int64 = 152
        _ = v20
    } else {
        fmt.Println(v1, v3)
    }
    v3 += 0
    fmt.Println(v3, v1)
    v21 := GenS3{F0: 942, F1: 146, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
