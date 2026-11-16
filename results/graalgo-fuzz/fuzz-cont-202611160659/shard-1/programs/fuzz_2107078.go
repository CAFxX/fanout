package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 uint32, p1 string, p2 int) uint64 {
    fmt.Println(p2, p0, p1)
    v6 := []int{}
    if true {
        if (320 < 1) {
            fmt.Println(p2)
        }
        var v7 int32 = 352
        _ = v7
    } else {
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 323
}
func genFn1(p0 int) int16 {
    p0 = -9223372036854775808
    fmt.Println(p0)
    if false {
        fmt.Println(p0)
        v8 := GenS1{F0: -2147483648, F1: 1, F2: 1}
        switch p0 {
        case 1:
            var v9 float32 = 3.14159
            _ = v9
        case -765:
            v10 := []int{(^(p0)), (^(p0)), -662, (-(p0))}
            _ = v10
        }
        _ = v8
    } else {
        var v11 uint8 = 128
        p0 += 1000
        fmt.Println(v11, p0)
        _ = v11
    }
    _ = p0
    return 32767
}

func main() {
    v1 := []string{"hello", "hello", ""}
    var v2 uint = 65535
    var v3 uint64 = 2
    v4 := []string{""}
    var v5 float32 = 1e10
    fmt.Println(v2, v3, v5)
    var v12 int16 = 437
    v2 = 18446744073709551615
    if (-267 <= (^(v12))) {
        if true {
            if true {
                v2 -= v2
            }
            var v13 int64 = -840
            v14 := GenS1{F0: 0, F1: 1, F2: 1}
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: -1, F1: 1, F2: 157}
        _ = v15
    } else {
        fmt.Println(v12, v12, v2)
        if ((1 == -902) && false) {
            for fi16 := 0; fi16 < 12; fi16++ {
                v2 = 0
                v3 = (^(v3))
                v2 -= (v2 | 2)
                _ = fi16
            }
            var v17 float32 = ((1e10 / float32(float64(v5))) + ((-2.25 + (0.1 + 0.0)) / (float32(3.14159) / v5)))
            _ = v17
        }
        switch "\n\t" {
        case "":
            var v18 uint8 = 86
            _ = v18
        default:
            if true {
                v2 = (v2 - v2)
                v2 += 42
            }
        }
    }
    var v19 int32 = 2147483647
    if true {
        var v20 int16 = v12
        var v21 int = 774
        var v22 uint32 = 4294967295
        _ = v20
        _ = v21
        _ = v22
    } else {
        if true {
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: (^(v12)), F1: "a", F2: 903}
                _ = fi23
                _ = v24
            }
            var v25 uint32 = 42
            v26 := GenS2{F0: v12, F1: "\n\t", F2: v25}
            _ = v25
            _ = v26
        } else {
            v2 = (v2 << (((v2 - (^(v2)))) & 63))
        }
        var v27 string = (("a" + "hello") + "\n\t")
        _ = v27
    }
    v28 := GenS1{F0: v19, F1: (v12 % ((v12) | 1)), F2: 0}
    var v29 string = "a"
    var v30 int32 = (v19 * 414)
    if true {
        _ = len(v1)
    } else {
        if len(v4) > 0 {
            v4[0] = v29
        }
        _ = len(v1)
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 uint = (v2 | v2)
        for fi33 := 0; fi33 < 18; fi33++ {
            if true {
                v32 -= v32
            }
            fmt.Println(v2)
            v19 = (v30 >> (((v32 >> ((v2) & 63))) & 63))
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    v5 = v5
    v2 += (v2 << ((v2) & 63))
    _ = len(v1)
    var v34 string = (v29 + v29)
    v35 := GenS3{F0: "\n\t", F1: (-(v19))}
    v2 = (v2 | v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
