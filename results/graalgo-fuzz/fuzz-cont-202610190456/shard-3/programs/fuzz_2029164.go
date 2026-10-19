package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 uint, p1 int8, p2 int) uint32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        p2 = (fi5 << (((p0 * (p0 * p0))) & 63))
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 uint16 = 255
            v8 := GenS1{F0: 42, F1: 0, F2: 42}
            _ = fi6
            _ = v7
            _ = v8
        }
        p2 = (^(fi5))
        _ = fi5
    }
    v9 := GenS1{F0: 1, F1: 60, F2: -2147483648}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 931
    var v3 uint8 = 2
    var v4 int16 = (v2 * -190)
    switch -1 {
    case 9223372036854775807:
        v10 := GenS1{F0: (v4 - v2), F1: 0, F2: -1}
        _ = v10
    }
    v2 += (^(v2))
    _ = len(v1)
    v11 := GenS1{F0: v2, F1: 127, F2: 42}
    if (42 != 2) {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint8 = v3
            var v14 uint8 = v3
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    var v15 uint = 4294967295
    var v16 bool = true
    switch "\n\t" {
    case "foo bar":
        var v17 uint16 = 65535
        _ = v17
    case "\n\t":
        v18 := GenS2{F0: 42, F1: (^(v2)), F2: -1}
        _ = v18
    case "":
        v3 = (v3 ^ v3)
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            if v16 {
                fmt.Println(v15)
            } else {
                fmt.Println(v2)
                v21 := GenS3{F0: 0}
                var v22 int8 = -75
                _ = v21
                _ = v22
            }
            fmt.Println(fi19, fi19, v2)
            v23 := GenS4{F0: 2147483647, F1: (v15 | v15)}
            _ = fi20
            _ = v23
        }
        v24 := GenS2{F0: 0, F1: v2, F2: -128}
        _ = fi19
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(v16, v2, v15)
        v16 = (!((false || v16) || false))
        _ = fi25
    }
    v26 := GenS3{F0: 1}
    v3 += v3
    v4 = (v4 - v2)
    fmt.Println(v2, v3, v16)
    v27 := GenS2{F0: 42, F1: 1, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
