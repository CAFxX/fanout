package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64, p1 int32) bool {
    v5 := GenS1{F0: 4294967295}
    var v6 uint32 = 707
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-128 < 100)
}
func genFn1(p0 int, p1 int, p2 uint) uint32 {
    v7 := GenS2{F0: (-9223372036854775808 == p0)}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]string)
    var v3 uint = 65535
    v4 := []uint32{1, 2147483648}
    v1 = v1
    var v8 int64 = 0
    fmt.Println(v3, v8, v1)
    for fi9 := 0; fi9 < 12; fi9++ {
        v8 -= v8
        var v10 int = (^(fi9))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint32 = 42
        v13 := GenS1{F0: (^(v12))}
        fmt.Println(v8)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v1)
    v8 = -1
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: 65535}
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (fi16 - fi14) {
            case -351:
                v2[42] = v1
                v1 = v1
                if (32768 < (fi16 + (^(fi14)))) {
                    v17 := GenS3{F0: v1}
                    switch v1 {
                    case "\n\t":
                        var v18 int16 = 567
                        _ = v18
                    default:
                        var v19 bool = false
                        _ = v19
                    }
                    v3 += v3
                    _ = len(v2)
                    var v20 int64 = v8
                    var v21 uint16 = 116
                    _ = v17
                    _ = v20
                    _ = v21
                }
                v22 := GenS4{F0: 100}
                _ = v22
            default:
                v23 := GenS2{F0: false}
                _ = v23
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v24 int = 288
    switch "\n\t" {
    case "\n\t":
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 uint8 = 128
            fmt.Println(v3, v3)
            v24 -= (v24 >> ((v3) & 63))
            _ = fi25
            _ = v26
        }
    case "":
        fmt.Println(v24, v3)
    default:
        v27 := GenS3{F0: "hello"}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
