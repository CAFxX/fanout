package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint, p2 uint64) uint {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: -958}
        v6 := GenS1{F0: 531}
        _ = fi4
        _ = v5
        _ = v6
    }
    switch 128 {
    case -2147483648:
        p2 -= p2
    case 42:
        p0 = (((p2 % ((p2) | 1)) == (p2 / ((p2) | 1))) || p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 uint64) int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 42}
        if false {
            v9 := GenS1{F0: -834}
            _ = v9
        }
        if (false || true) {
            var v10 int8 = 127
            _ = v10
        } else {
            fmt.Println(p0)
            v11 := []int{fi7, fi7}
            v12 := GenS2{F0: ""}
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    v13 := []int{int(p0)}
    _ = p0
    _ = v13
    return 339
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 709
    var v3 float64 = 1.5
    var v14 uint64 = 2
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS2{F0: "\n\t"}
        _ = fi15
        _ = v16
    }
    var v17 int32 = (v2 + v2)
    v3 = float64(0.1)
    fmt.Println(v1, v3, v3)
    v18 := GenS1{F0: -524}
    fmt.Println(v1)
    var v19 string = "hello"
    v20 := []int{4294967295, -175, 128, -592}
    var v21 uint32 = 938
    var v22 uint64 = (v14 | (v14 - v14))
    fmt.Println(v19, v21)
    if len(v20) > 0 {
        v20[0] = 127
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            v19 = v19
            v25 := GenS3{F0: 255}
            switch fi23 {
            case 620:
                v19 = v19
            case -207:
                var v26 uint = 255
                _ = v26
            default:
                var v27 uint = 65535
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        _ = len(v20)
        _ = fi23
    }
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
