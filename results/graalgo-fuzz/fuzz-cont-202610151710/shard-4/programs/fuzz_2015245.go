package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 uint, p1 int64, p2 int16) uint32 {
    v7 := GenS1{F0: (!((843 >= 18446744073709551615) && false))}
    if true {
        var v8 int16 = -32768
        _ = v8
    } else {
        if (!false) {
            if (!(false && true)) {
                p2 = p2
            } else {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 int16 = -32768
                    v11 := GenS2{F0: 71, F1: (v10 ^ (p2 ^ -1)), F2: true}
                    v12 := []int{fi9, (fi9 | 649), fi9, fi9}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
        } else {
            p0 += (p0 ^ p0)
        }
        var v13 int = 450
        _ = v13
    }
    v14 := GenS3{F0: 0, F1: int16(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 string = ("foo bar" + "a")
    v6 := make(map[int]bool)
    v1[127] = -346
    for fi15 := 0; fi15 < 13; fi15++ {
        v5 = "\n\t"
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16, v5, fi15)
            v5 = "a"
            var v17 int8 = -90
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v6[(fi18 % (((fi18 << ((461) & 63))) | 1))] = (false || ((true || (218 < 255)) || (255 < 65535)))
        _ = fi18
    }
    fmt.Println(v5)
    for fi19 := 0; fi19 < 5; fi19++ {
        _ = len(v1)
        if true {
            var v20 bool = true
            v21 := []int{383, fi19, (fi19 % ((fi19) | 1)), -2147483648}
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    v5 = ((v5 + v5) + v5)
    var v22 uint64 = 18446744073709551615
    switch 4294967295 {
    case -109:
        var v23 int8 = 100
        _ = v23
    case 42:
        var v24 int = -9223372036854775808
        _ = v24
    case 32768:
        v25 := GenS4{F0: 1, F1: -797}
        _ = v25
    default:
        v5 = "hello"
    }
    v3[""] = 32768
    v22 = 4294967295
    v26 := []int{}
    delete(v6, 1)
    var v27 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
