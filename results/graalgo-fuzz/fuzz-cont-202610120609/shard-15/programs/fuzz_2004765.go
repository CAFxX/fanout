package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 int8, p1 int16, p2 uint8) int32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = p2
        p1 = 1000
        _ = fi6
        _ = v7
    }
    p1 -= 42
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint16, p1 uint, p2 uint64) uint64 {
    p0 = 601
    if (-1 == -32768) {
        v8 := []int{1}
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 1}
            _ = fi9
            _ = v10
        }
        var v11 uint8 = 0
        _ = v11
        _ = v8
    } else {
        v12 := GenS1{F0: 42}
        fmt.Println(p1, p2, p1)
        fmt.Println(p0, p1, p0)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 | (p2 << ((1) & 63)))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 uint16 = 926
    var v4 uint16 = v3
    var v5 uint16 = 0
    fmt.Println(v3, v5)
    for fi13 := 0; fi13 < 14; fi13++ {
        v5 += (v3 << ((255) & 63))
        v3 -= v3
        _ = fi13
    }
    delete(v2, 127)
    v4 += 1
    fmt.Println(v3, v5)
    v14 := GenS2{F0: 9223372036854775807, F1: 0, F2: 778}
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint64 = 18446744073709551615
        var v17 uint32 = uint32(v16)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: (((("foo bar" + "hello") + ("\n\t" + "hello")) + "hello") >= "foo bar")}
    if true {
        if (false || (-1 <= 42)) {
            var v19 int64 = -318
            _ = v19
        } else {
            v2[432] = true
            fmt.Println(v3, v4)
            v3 -= (v5 + v5)
        }
    } else {
        var v20 int = 32767
        if false {
            var v21 int8 = 100
            switch v20 {
            case 128:
                v5 = (v4 + v3)
            case 4294967295:
                if true {
                    delete(v2, (v20 ^ (v20 & 127)))
                    switch v20 {
                    case -1:
                        v20 = v20
                    case 2147483647:
                        for fi22 := 0; fi22 < 6; fi22++ {
                            var v23 int = v20
                            v20 -= 2147483647
                            var v24 uint32 = 2
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    v4 += v4
                }
            case 1:
                var v25 uint16 = (v5 + (v5 * 0))
                _ = v25
            default:
                switch "\n\t" {
                case "a":
                    var v26 uint16 = 1
                    _ = v26
                }
            }
            var v27 uint8 = 127
            _ = v21
            _ = v27
        }
        _ = v20
    }
    v3 += v4
    fmt.Println(v5, v4)
    if ((127 == 32768) && false) {
        v28 := GenS4{F0: 2, F1: 589, F2: 0}
        var v29 int8 = 111
        _ = v28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
