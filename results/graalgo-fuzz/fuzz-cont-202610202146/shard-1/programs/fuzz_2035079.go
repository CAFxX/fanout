package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0() int16 {
    var v5 bool = false
    v5 = v5
    var v6 int32 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        v5 = v5
        fmt.Println(v5, v5, fi7)
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int8 = 100
            if v5 {
                var v10 string = "hello"
                fmt.Println(v9, v6)
                var v11 uint64 = 4294967295
                var v12 bool = true
                switch fi8 {
                case 42:
                    var v13 string = v10
                    _ = v13
                default:
                    var v14 string = ("a" + v10)
                    _ = v14
                }
                _ = v10
                _ = v11
                _ = v12
            } else {
                v15 := []int{fi7, 331}
                _ = v15
            }
            var v16 int8 = v9
            _ = fi8
            _ = v16
            _ = v9
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{42, 42, 42, 343}
    v2 := []int{771}
    var v3 int64 = -9223372036854775808
    var v4 float32 = 3.14159
    fmt.Println(v4)
    if ((true && (42 >= -613)) || true) {
        var v17 uint64 = 18446744073709551615
        switch "\n\t" {
        case "a":
            if true {
                fmt.Println(v4, v3)
            }
        case "hello":
            v4 += v4
        default:
            _ = len(v1)
        }
        _ = v17
    }
    if ((0 == 2) || (0.0 != 1.5)) {
        if false {
            v18 := GenS1{F0: 18446744073709551615, F1: 127, F2: 1}
            var v19 uint32 = 9
            v19 = v19
            _ = v18
            _ = v19
        }
        var v20 uint = 857
        _ = v20
    } else {
        if len(v2) > 0 {
            v2[0] = 711
        }
    }
    var v21 float32 = float32(((1e10 * 1e10) / 1.5))
    var v22 uint = 1
    v23 := GenS2{F0: 2147483647, F1: (v22 != v22), F2: 42}
    if true {
        v24 := GenS1{F0: (v22 * v22), F1: 0, F2: 695}
        v2 = append(v2, 4294967295)
        for fi25 := 0; fi25 < 17; fi25++ {
            if true {
                var v26 float64 = 1.5
                fmt.Println(v21, v26, v21)
                v27 := GenS3{F0: 691}
                _ = v26
                _ = v27
            }
            _ = fi25
        }
        _ = v24
    } else {
        v3 = v3
    }
    _ = len(v1)
    v3 += (v3 + (v3 ^ 1))
    v28 := GenS3{F0: -216}
    if true {
        v1 = append(v1, 65535)
    }
    v3 -= v3
    fmt.Println(v22, v3)
    for fi29 := 0; fi29 < 17; fi29++ {
        _ = len(v1)
        _ = len(v2)
        v30 := GenS3{F0: 1}
        _ = fi29
        _ = v30
    }
    fmt.Println(v22, v21)
    v31 := GenS4{F0: 4294967295, F1: 42}
    for fi32 := 0; fi32 < 4; fi32++ {
        if len(v2) > 0 {
            v2[0] = (^(fi32))
        }
        var v33 int16 = 1
        _ = fi32
        _ = v33
    }
    switch "" {
    case "a":
        v21 = (float32(1e10) + float32((3.14159 + float64(((v21 - v4) / ((3.14159 * v4) * 1e10))))))
    case "foo bar":
        v34 := GenS4{F0: 42, F1: 1}
        _ = v34
    default:
        for fi35 := 0; fi35 < 4; fi35++ {
            if len(v2) > 0 {
                v2[0] = fi35
            }
            var v36 float32 = (float32((((0.1 + 0.1) - (-2.25 + 0.1)) * -2.25)) + v4)
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
