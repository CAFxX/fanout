package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int8) int64 {
    v7 := GenS1{F0: -1}
    if (!(false && ("hello" != ("\n\t" + "a")))) {
        p0 += (p0 | p0)
    } else {
        var v8 float32 = 0.0
        _ = v8
    }
    p0 += p0
    for fi9 := 0; fi9 < 2; fi9++ {
        if (42 < 42) {
            v10 := GenS1{F0: 1}
            _ = v10
        }
        if (65535 == 1) {
            var v11 uint64 = 1
            p0 += p0
            var v12 uint = 65535
            _ = v11
            _ = v12
        } else {
            v13 := []int{fi9, -1, fi9, (fi9 + fi9)}
            var v14 int = fi9
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 int8, p2 bool) int8 {
    v15 := GenS1{F0: -1}
    if p2 {
        for fi16 := 0; fi16 < 2; fi16++ {
            if p2 {
                if ((p2 && p2) || (255 < 2)) {
                    var v17 float32 = 0.1
                    p2 = p2
                    _ = v17
                }
            } else {
                fmt.Println(p0)
            }
            _ = fi16
        }
        var v18 bool = false
        _ = v18
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: -1}
            fmt.Println(fi19, p2, p2)
            p0 += p0
            _ = fi19
            _ = v20
        }
        var v21 int64 = 42
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p1
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]uint32)
    var v3 int8 = -128
    var v4 float32 = 0.0
    var v5 float64 = 1e10
    v6 := make(map[int]int)
    if true {
        var v22 uint16 = 870
        _ = len(v2)
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        switch v1 {
        case 127:
            var v24 uint64 = 29
            _ = v24
        default:
            fmt.Println(v5, v1, v1)
        }
        v25 := GenS2{F0: 4294967295}
        _ = fi23
        _ = v25
    }
    v26 := []int{42, v1, 4294967295, v1}
    v3 -= (v3 << ((0) & 63))
    v26 = append(v26, v1)
    v6[(-(v1))] = 32767
    v6[(v1 * 166)] = (v1 % (((v1 * v1)) | 1))
    v27 := GenS3{F0: 987, F1: 172}
    v4 -= 1.5
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint16 = 65535
        v30 := GenS3{F0: v29, F1: 32767}
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v4, v1, v3)
    var v31 uint64 = 42
    var v32 int32 = -443
    v33 := GenS1{F0: (v32 % ((v32) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
