package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int }
type GenS4 struct { F0 string; F1 int; F2 int16 }
func genFn0(p0 int, p1 uint8, p2 uint64) int32 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: -1, F1: 9223372036854775807, F2: -32768}
            p1 -= p1
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint = 532
        _ = fi8
        _ = v9
    }
    p2 += p2
    if (true || (27 > 255)) {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int32 = -782
            v12 := GenS1{F0: -201, F1: 42, F2: 56}
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        var v13 int8 = -83
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 18446744073709551615
    v3 := make(map[string]string)
    var v4 uint8 = 255
    fmt.Println(v2)
    var v14 int16 = 0
    for fi15 := 0; fi15 < 13; fi15++ {
        if (!true) {
            var v16 int64 = 9223372036854775807
            v17 := GenS2{F0: uint32(v4)}
            v4 += 1
            _ = v16
            _ = v17
        } else {
            v3["hello"] = ""
        }
        v18 := GenS3{F0: int64(v1), F1: 655, F2: fi15}
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if (false || (3.14159 != 0.0)) {
            v2 = (v2 % (((v2 << (((v2 >> ((377) & 63))) & 63))) | 1))
        } else {
            var v20 uint = v2
            delete(v3, ("" + "foo bar"))
            _ = v20
        }
        var v21 uint8 = 54
        _ = fi19
        _ = v21
    }
    v22 := GenS4{F0: ("a" + "hello"), F1: -1, F2: 0}
    v2 = (v2 & 18446744073709551615)
    v1 += v1
    fmt.Println(v4, v2)
    var v23 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
