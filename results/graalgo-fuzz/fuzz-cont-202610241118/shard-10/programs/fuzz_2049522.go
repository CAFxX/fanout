package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 int16) int16 {
    v7 := []int{}
    p0 = ("" == "\n\t")
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 string = ""
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 int16, p1 int, p2 uint64) int64 {
    v10 := []int{(p1 >> ((uint(p2)) & 63)), p1, (p1 ^ p1), (p1 & p1)}
    if (0 >= -1) {
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = v2
    var v6 bool = false
    v1 = (v1 / (((v1 / ((genFn1(287, 4294967295, 0)) | 1))) | 1))
    _ = len(v4)
    v11 := GenS1{F0: -1}
    v2 += (v5 % ((v5) | 1))
    var v12 int16 = genFn0(v6, 0)
    v13 := []int{}
    v3[1] = true
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: 100}
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 0}
    fmt.Println(v12, v12, v1)
    for fi17 := 0; fi17 < 17; fi17++ {
        v6 = (591 != 1)
        if len(v13) > 0 {
            v13[0] = fi17
        }
        _ = fi17
    }
    fmt.Println(v2, v5, v2)
    v2 = (v2 >> ((18446744073709551615) & 63))
    for fi18 := 0; fi18 < 15; fi18++ {
        _ = len(v4)
        if (!v6) {
            fmt.Println(v2, v6, v5)
            for fi19 := 0; fi19 < 15; fi19++ {
                v3[4294967295] = (!(42 >= -2147483648))
                var v20 int32 = -1
                _ = fi19
                _ = v20
            }
            fmt.Println(v2, v12, v1)
        } else {
            v1 += (v1 | v1)
            var v21 int = fi18
            if (float64((3.14159 + 0.0)) > 0.1) {
                delete(v4, uint32(v5))
                fmt.Println(v1)
                v2 = (v5 * (v5 * 0))
            }
            _ = v21
        }
        _ = fi18
    }
    if len(v13) > 0 {
        v13[0] = 127
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        delete(v4, 337)
        _ = fi22
    }
    var v23 int64 = v1
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
