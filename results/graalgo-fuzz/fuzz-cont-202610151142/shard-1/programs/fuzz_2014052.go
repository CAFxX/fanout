package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    if true {
        var v4 uint = 4294967295
        _ = v4
    } else {
        v5 := GenS1{F0: 93, F1: 1, F2: "\n\t"}
        if (false || false) {
            if true {
                var v6 string = "\n\t"
                fmt.Println(v6)
                _ = v6
            }
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int8 = -128
                v9 := GenS2{F0: 0, F1: 1, F2: 42}
                _ = fi7
                _ = v8
                _ = v9
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v10 int32 = -320
            _ = v10
        }
        _ = v5
    }
    v1[4294967295] = -706
    var v11 uint16 = 338
    var v12 uint64 = 983
    fmt.Println(v12)
    if true {
        if (((585 < 255) || false) || true) {
            v12 = (v12 >> ((uint(v11)) & 63))
        }
        fmt.Println(v12)
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int8 = 116
            _ = fi13
            _ = v14
        }
    }
    if (v12 <= (v12 & (^(v12)))) {
        fmt.Println(v12)
        if false {
            v15 := v2["hello"]
            fmt.Println(v15)
            _ = v15
        }
        _ = len(v3)
    } else {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        v11 -= (v11 | v11)
        v16 := GenS3{F0: (("hello" + "") + "foo bar")}
        _ = v16
    }
    switch ("" + ("hello" + "hello")) {
    case "foo bar":
        v17 := GenS4{F0: 128}
        _ = v17
    case "hello":
        fmt.Println(v11, v11)
    default:
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                v12 += v12
                v11 = (v11 ^ v11)
                v20 := GenS5{F0: 422}
                v12 += 4294967295
                _ = fi19
                _ = v20
            }
            delete(v2, "a")
            v11 -= v11
            _ = fi18
        }
    }
    var v21 uint = 574
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
}
