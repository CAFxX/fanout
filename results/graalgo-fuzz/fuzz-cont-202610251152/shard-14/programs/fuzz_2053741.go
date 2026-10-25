package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 string, p1 uint64) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{fi7, fi7}
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(fi9, p0)
            _ = fi9
        }
        var v10 int64 = -765
        _ = fi7
        _ = v10
        _ = v8
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (p0 + ((("" + ("foo bar" + "hello")) + p0) + ("" + p0))) {
        case "foo bar":
            p1 = (p1 % ((p1) | 1))
        case "hello":
            switch (^(fi12)) {
            case 32768:
                for fi13 := 0; fi13 < 12; fi13++ {
                    fmt.Println(fi12, p1)
                    v14 := GenS1{F0: 0, F1: (("foo bar" + "hello") + p0), F2: 255}
                    v15 := GenS2{F0: (fi13 & (fi13 - 128))}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case -9223372036854775808:
                p0 = (p0 + ("a" + p0))
            case -632:
                p1 -= (p1 ^ (^(p1)))
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    if true {
                        if false {
                            var v17 uint16 = 2
                            var v18 float32 = (0.0 - 0.1)
                            _ = v17
                            _ = v18
                        }
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = fi16
                }
            }
        }
        _ = fi12
    }
    if (127 == 1) {
        var v19 float32 = 0.1
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v11
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = ((-939 <= 128) || (false && true))
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2}
    v5 := make(map[int]bool)
    var v6 bool = (607 > 785)
    var v20 int16 = 1
    switch 613 {
    case 9223372036854775807:
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    v20 += 323
    fmt.Println(v20, v2, v20)
    v21 := GenS3{F0: 744}
    v4 = append(v4, 2)
    v20 += v20
    var v22 uint16 = 255
    if false {
        var v23 string = ("hello" + ("a" + "\n\t"))
        _ = v23
    }
    var v24 uint8 = 5
    v25 := GenS4{F0: 2147483647, F1: 32768, F2: 743}
    var v26 int32 = -35
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
