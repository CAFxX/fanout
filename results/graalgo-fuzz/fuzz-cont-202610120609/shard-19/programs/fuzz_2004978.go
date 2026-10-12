package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 uint64, p1 int32, p2 int64) int16 {
    var v6 uint = 18446744073709551615
    v7 := []int{127, 9223372036854775807, -340}
    switch int(p1) {
    case -1:
        v8 := GenS1{F0: (!false), F1: -1}
        _ = v8
    case 32767:
        v9 := GenS2{F0: 2, F1: "hello", F2: 353}
        _ = v9
    case 9223372036854775807:
        var v10 int8 = 100
        _ = v10
    }
    p2 += (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 uint) int {
    v11 := []int{0}
    v12 := GenS2{F0: 42, F1: "a", F2: 276}
    var v13 float32 = 1e10
    var v14 string = ""
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -349
}

func main() {
    var v1 uint64 = 2
    var v2 int8 = 42
    v3 := make(map[string]bool)
    var v4 string = "a"
    v5 := []uint32{2, 287}
    v1 -= 1
    if true {
        v1 = (v1 * (v1 | (^(v1))))
        v15 := v3[(v4 + (((v4 + v4) + v4) + (v4 + v4)))]
        v16 := GenS2{F0: 42, F1: v4, F2: 65535}
        _ = v15
        _ = v16
    }
    fmt.Println(v1, v1, v1)
    v4 = "foo bar"
    var v17 uint16 = 0
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int16 = 32767
        var v20 int32 = -654
        v21 := GenS3{F0: "hello", F1: v1}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch 127 {
    case 678:
        var v22 int = -494
        _ = v22
    case -1:
        v17 = (v17 * 255)
    default:
        if (!true) {
            var v23 uint8 = 128
            delete(v3, v4)
            _ = v23
        }
    }
    v24 := v3[v4]
    v25 := GenS2{F0: (-(v2)), F1: v4, F2: 823}
    v2 = v2
    v2 = v2
    var v26 int64 = -142
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
