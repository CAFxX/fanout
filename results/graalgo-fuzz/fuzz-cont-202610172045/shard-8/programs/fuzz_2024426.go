package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
func genFn0() int {
    var v4 uint64 = 216
    v4 += v4
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{fi5, fi5, -1, (fi5 % (((fi5 * (fi5 - (fi5 % ((int(v4)) | 1))))) | 1))}
        var v7 int8 = 55
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = v4
    _ = v4
    _ = v8
    return 128
}
func genFn1(p0 string) int8 {
    v9 := GenS1{F0: "hello", F1: (0 != 65535), F2: -2147483648}
    v10 := []int{889, -1}
    v11 := GenS1{F0: p0, F1: false, F2: -142}
    if false {
        var v12 uint16 = 65535
        _ = v12
    } else {
        if true {
            v13 := GenS1{F0: (((p0 + ("a" + p0)) + ((p0 + p0) + "hello")) + p0), F1: false, F2: 127}
            _ = v13
        } else {
            p0 = p0
            p0 = (p0 + "\n\t")
        }
        var v14 int8 = 105
        v14 -= (v14 >> ((42) & 63))
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{554, 52, 65535, 65535}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    for fi15 := 0; fi15 < 2; fi15++ {
        v3 += v3
        var v16 float64 = (3.14159 + (0.0 + (3.14159 + ((-2.25 + 3.14159) - 1.5))))
        _ = fi15
        _ = v16
    }
    var v17 int64 = -631
    if false {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    }
    v1 = append(v1, 4294967295)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v18 := GenS1{F0: "hello", F1: (!true), F2: 32768}
    fmt.Println(v17, v3)
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 12; fi19++ {
        v3 = (v3 >> ((0) & 63))
        _ = fi19
    }
    if (65535 > 2) {
        var v20 float64 = 1e10
        fmt.Println(v3, v17, v20)
        _ = v20
    }
    fmt.Println(v17, v17)
    delete(v2, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
