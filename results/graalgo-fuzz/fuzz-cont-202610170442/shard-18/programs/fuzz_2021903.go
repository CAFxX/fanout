package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    var v3 uint16 = 2
    v4 := make(map[uint32]uint32)
    v2 = (-35 != 0)
    for fi5 := 0; fi5 < 7; fi5++ {
        if ("hello" < "foo bar") {
            v6 := GenS1{F0: 0, F1: v2}
            v3 = (v3 & v3)
            v2 = v2
            _ = v6
        }
        _ = fi5
    }
    var v7 uint16 = (v3 | v3)
    v7 = v3
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v8, v7, fi9)
        v10 := GenS1{F0: 0, F1: true}
        _ = fi9
        _ = v10
    }
    v11 := []int{1, -121, 42}
    if (v2 || (18446744073709551615 >= 225)) {
        v12 := GenS2{F0: 100, F1: 42}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(v3, v7, v8)
            if (!(false || v2)) {
                v14 := GenS3{F0: -9223372036854775808, F1: ("foo bar" + ("hello" + "\n\t"))}
                var v15 uint64 = 0
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    if len(v11) > 0 {
        v11[0] = 32767
    }
    v3 = v7
    switch 9223372036854775807 {
    case -293:
        for fi16 := 0; fi16 < 13; fi16++ {
            if (1 >= 255) {
                var v17 float64 = (float64(0.0) - 1e10)
                var v18 uint = 765
                v18 -= v18
                _ = v17
                _ = v18
            } else {
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int = (-(fi19))
                    _ = fi19
                    _ = v20
                }
                if len(v11) > 0 {
                    v11[0] = 2147483647
                }
                fmt.Println(v2, v7, v3)
            }
            _ = fi16
        }
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int16 = 32767
        v23 := GenS4{F0: 127, F1: (fi21 * (fi21 >> ((4294967295) & 63)))}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v7, v3)
    v7 += (v3 * (v7 << ((65535) & 63)))
    v8 += v8
    v24 := GenS1{F0: -128, F1: true}
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int8 = 21
        v2 = (193 == 127)
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
