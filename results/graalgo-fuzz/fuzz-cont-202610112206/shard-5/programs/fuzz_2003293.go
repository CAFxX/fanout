package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 string) uint {
    v7 := GenS1{F0: -2147483648, F1: 42, F2: 128}
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 = p0
        var v9 uint32 = 4294967295
        v9 += v9
        _ = fi8
        _ = v9
    }
    p0 = "a"
    var v10 uint = 489
    _ = p0
    _ = v10
    _ = v7
    return v10
}

func main() {
    var v1 uint8 = 131
    var v2 uint = 350
    var v3 uint8 = (v1 | v1)
    v4 := []string{"a", "foo bar", ("hello" + "foo bar")}
    v5 := make(map[uint32]string)
    var v6 uint = v2
    v3 = (v1 - 127)
    v3 = (v3 - 0)
    var v11 uint64 = 1
    v12 := GenS2{F0: v1, F1: -635}
    switch -2147483648 {
    case -2147483648:
        var v13 uint8 = (v3 & v3)
        _ = v13
    case -92:
        var v14 int = -1
        _ = v14
    default:
        v15 := GenS3{F0: 1}
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v1)
        _ = fi16
    }
    v11 -= (v11 / ((v11) | 1))
    var v17 float32 = (-2.25 * 0.0)
    v18 := GenS1{F0: -1, F1: 32767, F2: (v3 + (v1 / ((1) | 1)))}
    switch "foo bar" {
    case "foo bar":
        var v19 uint16 = 255
        _ = v19
    default:
        for fi20 := 0; fi20 < 1; fi20++ {
            v2 = uint(v1)
            if true {
                fmt.Println(v2)
                if (100 != 1) {
                    var v21 int = (fi20 + (fi20 + fi20))
                    _ = len(v4)
                    _ = v21
                }
            } else {
                var v22 int64 = 939
                _ = v22
            }
            if false {
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 string = "a"
                    v4 = append(v4, (((v24 + v24) + ("" + "foo bar")) + (v24 + ("" + v24))))
                    v2 += v6
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 uint32 = 65535
                _ = v25
            }
            _ = fi20
        }
    }
    var v26 int64 = -9223372036854775808
    v1 = (v1 << ((515) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
