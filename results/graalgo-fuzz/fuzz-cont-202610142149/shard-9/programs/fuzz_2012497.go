package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = true
    v2 := []uint32{886, 624, 65535}
    var v3 uint8 = 42
    v4 := []uint32{2147483647}
    switch ("\n\t" + (("a" + "") + ("a" + ("foo bar" + "hello")))) {
    case "a":
        var v5 uint16 = 42
        _ = v5
    case "foo bar":
        var v6 int16 = 42
        _ = v6
    case "":
        if len(v4) > 0 {
            v4[0] = 9
        }
    default:
        v3 = (^(v3))
    }
    v4 = append(v4, 1)
    fmt.Println(v1)
    if v1 {
        fmt.Println(v3, v1)
    }
    var v7 uint64 = 734
    if (0 < 0) {
        var v8 uint8 = (v3 / (((v3 & 189)) | 1))
        for fi9 := 0; fi9 < 17; fi9++ {
            v4 = append(v4, 791)
            _ = len(v4)
            v10 := GenS1{F0: 690}
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        var v11 float64 = 0.0
        v12 := GenS1{F0: 81}
        _ = v11
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    if true {
        var v13 uint16 = 2
        if v1 {
            var v14 int64 = 9223372036854775807
            if v1 {
                v3 -= (v3 + 62)
                var v15 float64 = 0.1
                _ = v15
            }
            var v16 uint = 4294967295
            _ = v14
            _ = v16
        }
        var v17 uint16 = (v13 & (v13 * uint16(v3)))
        _ = v13
        _ = v17
    } else {
        var v18 uint = 665
        if v1 {
            if (v1 || (true || (-32768 > 1000))) {
                v18 = (^(v18))
            } else {
                v18 += (v18 & v18)
                var v19 int8 = -73
                v20 := GenS1{F0: 336}
                _ = v19
                _ = v20
            }
        } else {
            var v21 int32 = -922
            v22 := GenS1{F0: 42}
            v21 = (v21 << ((v18) & 63))
            _ = v21
            _ = v22
        }
        _ = v18
    }
    v23 := GenS2{F0: 42}
    v1 = (-821 < 0)
    var v24 uint16 = 2
    v3 = (v3 << ((42) & 63))
    v3 = v3
    var v25 uint32 = 384
    v7 = v7
    v26 := GenS2{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
}
