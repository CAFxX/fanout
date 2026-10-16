package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 int32 }
func genFn0() int32 {
    var v6 int64 = -677
    fmt.Println(v6)
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 9; fi8++ {
        switch fi8 {
        case 4294967295:
            v9 := []int{fi8}
            _ = v9
        case 9223372036854775807:
            var v10 int = fi8
            _ = v10
        default:
            v11 := GenS1{F0: 4294967295, F1: 32767}
            _ = v11
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return (v7 >> ((2) & 63))
}
func genFn1(p0 string) int {
    switch p0 {
    case "hello":
        if true {
            p0 = "a"
            for fi12 := 0; fi12 < 12; fi12++ {
                p0 = p0
                for fi13 := 0; fi13 < 14; fi13++ {
                    p0 = ((p0 + (p0 + "")) + "foo bar")
                    _ = fi13
                }
                p0 = "\n\t"
                _ = fi12
            }
        }
    default:
        v14 := GenS1{F0: 42, F1: -404}
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: 481, F2: -44}
    var v16 uint64 = 956
    _ = p0
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    var v5 uint8 = 2
    var v17 int32 = 147
    delete(v3, 837)
    if true {
        for fi18 := 0; fi18 < 6; fi18++ {
            v4 = v4
            _ = fi18
        }
        v19 := GenS3{F0: v17, F1: v17, F2: (true || false)}
        if (241 < 2147483647) {
            v1 = v1
            v17 -= v17
        }
        _ = v19
    }
    switch genFn1(("\n\t" + ("\n\t" + "foo bar"))) {
    case -2147483648:
        var v20 bool = ((2 == 2147483648) && true)
        _ = v20
    case -956:
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS4{F0: v17}
            var v23 int16 = 32767
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v1, fi24)
            _ = fi24
        }
    }
    var v25 uint8 = (^(v5))
    var v26 int = genFn1("")
    v26 += (v26 % ((v26) | 1))
    var v27 uint16 = (^(v1))
    v27 = v1
    v28 := v2[2147483648]
    var v29 int16 = 1000
    v29 = -1
    var v30 uint64 = 18446744073709551615
    var v31 uint16 = (^(v1))
    v32 := []int{}
    var v33 int8 = -128
    delete(v3, v28)
    if false {
        fmt.Println(v4)
        v34 := GenS2{F0: v33, F1: (v29 * v29), F2: int8(v30)}
        var v35 int32 = v17
        _ = v34
        _ = v35
    } else {
        var v36 int64 = (v4 + (v4 | (v4 ^ -1)))
        var v37 int16 = v29
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
