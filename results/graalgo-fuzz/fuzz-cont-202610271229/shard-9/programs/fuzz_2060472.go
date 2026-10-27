package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 bool }
func genFn0(p0 uint32, p1 string) int {
    var v6 uint8 = 53
    if true {
        p0 -= (p0 ^ (p0 >> ((uint(p0)) & 63)))
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int16 = 1
        v9 := []int{fi7, (^(fi7))}
        if (true && false) {
            var v10 int16 = v8
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        var v11 int = -207
        _ = v11
    } else {
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}
func genFn1() uint {
    switch -541 {
    case 1:
        var v12 int64 = -710
        _ = v12
    case 36:
        v13 := GenS1{F0: 238, F1: 286}
        _ = v13
    case -9223372036854775808:
        var v14 uint32 = 2147483648
        _ = v14
    default:
        if (false || false) {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS1{F0: 1000, F1: -339}
                var v17 int8 = 100
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            v18 := GenS2{F0: 42, F1: 125, F2: -2147483648}
            var v19 int64 = -1
            _ = v18
            _ = v19
        }
    }
    var v20 uint8 = 1
    _ = v20
    return 2
}

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 625
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    var v5 uint64 = (v1 << ((255) & 63))
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: 42, F1: 1}
        v3[0] = false
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 1000, F1: 0}
    v24 := GenS3{F0: 255, F1: "a", F2: ((0.0 + (float32(0.0) * 3.14159)) <= 1.5)}
    v2 = v2
    v5 -= 2
    v25 := v3[v2]
    var v26 int64 = v4
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 uint = 255
        v29 := v3[v2]
        var v30 uint32 = v2
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 uint = 2
    var v32 int64 = (^(v26))
    v33 := v3[v2]
    if (-1 >= 42) {
        var v34 uint16 = 32767
        var v35 string = "a"
        var v36 int = genFn0(v2, ((v35 + ("a" + v35)) + (v35 + v35)))
        _ = v34
        _ = v35
        _ = v36
    }
    v37 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
