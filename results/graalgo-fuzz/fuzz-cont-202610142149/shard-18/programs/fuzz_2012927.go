package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint }
func genFn0(p0 int64, p1 bool, p2 int64) int16 {
    v6 := []int{263}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 162
        fmt.Println(p0, v8)
        _ = fi7
        _ = v8
    }
    switch 746 {
    case 128:
        switch ("a" + "") {
        case "\n\t":
            for fi9 := 0; fi9 < 19; fi9++ {
                p0 = -1
                _ = fi9
            }
        case "":
            if p1 {
                var v10 int32 = 0
                p0 = 1
                if (!p1) {
                    _ = len(v6)
                    var v11 uint = 2
                    if p1 {
                        var v12 uint32 = 0
                        _ = v12
                    }
                    _ = v11
                }
                _ = v10
            }
        default:
            var v13 bool = true
            _ = v13
        }
    default:
        fmt.Println(p1, p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -746
}

func main() {
    v1 := []uint32{187, 679, 827}
    var v2 int8 = -128
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    v5 := make(map[int]uint32)
    v14 := GenS1{F0: -9223372036854775808}
    var v15 int64 = 1
    v3 += (v3 & (v3 ^ v3))
    fmt.Println(v15, v3, v3)
    fmt.Println(v3, v2, v3)
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(fi16, v15, fi16)
        fmt.Println(v2, fi16)
        _ = len(v1)
        _ = fi16
    }
    if true {
        v1 = append(v1, 2147483648)
    } else {
        _ = len(v1)
    }
    v17 := GenS2{F0: v3, F1: v2, F2: ((("\n\t" + "foo bar") + "a") + (("\n\t" + "") + ""))}
    v15 += (v15 | v15)
    v18 := GenS3{F0: 732, F1: (("\n\t" + ("foo bar" + "hello")) + ("\n\t" + "a"))}
    _ = len(v1)
    var v19 uint16 = 0
    v20 := GenS4{F0: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
