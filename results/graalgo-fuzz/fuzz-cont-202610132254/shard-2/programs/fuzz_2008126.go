package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 bool) uint32 {
    fmt.Println(p0, p1)
    p0 -= 713
    for fi4 := 0; fi4 < 19; fi4++ {
        if p1 {
            p1 = (1 != 0)
        } else {
            p0 = (p0 | p0)
            var v5 int = (fi4 % ((int(p0)) | 1))
            _ = v5
        }
        v6 := []int{}
        var v7 int16 = 32767
        _ = fi4
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 uint = 555
    v3 := make(map[int]uint32)
    var v8 uint8 = 45
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{-414, fi10, fi9, -9223372036854775808}
            v12 := GenS1{F0: 34, F1: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int32 = 42
        if false {
            fmt.Println(v2)
        }
        _ = fi9
        _ = v13
    }
    v14 := v3[127]
    v15 := GenS2{F0: 18446744073709551615}
    v1 += v1
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint64 = 667
        v18 := GenS1{F0: 100, F1: genFn0(int64(v1), true)}
        _ = fi16
        _ = v17
        _ = v18
    }
    delete(v3, 0)
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := []int{(fi19 * fi19), (fi19 & fi19), 958}
        var v21 bool = false
        delete(v3, fi19)
        _ = fi19
        _ = v20
        _ = v21
    }
    v8 -= v8
    v1 += (v1 >> ((4294967295) & 63))
    v1 -= (v1 ^ v1)
    var v22 uint16 = 255
    var v23 int = int(v1)
    switch 2147483647 {
    case 128:
        fmt.Println(v2)
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v23 += (v23 >> ((189) & 63))
        _ = fi24
    }
    delete(v3, (v23 | -66))
    v25 := v3[876]
    if false {
        var v26 float64 = 0.1
        _ = v26
    } else {
        v3[-9223372036854775808] = uint32(v22)
        if false {
            var v27 float32 = (-2.25 * 1.5)
            _ = v27
        } else {
            v23 = v23
            v2 = (^(v2))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v8
}
