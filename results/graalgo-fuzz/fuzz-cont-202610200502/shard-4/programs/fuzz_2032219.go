package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint, p1 int64, p2 uint32) int32 {
    fmt.Println(p0, p2, p1)
    if true {
        p0 += (^(p0))
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 1, F1: 9223372036854775807, F2: 32767}
            _ = fi6
            _ = v7
        }
        var v8 float64 = 3.14159
        _ = v8
    } else {
        if (!(263 >= 938)) {
            v9 := GenS2{F0: 42, F1: 128, F2: p2}
            v10 := GenS2{F0: 42, F1: 0, F2: p2}
            v11 := []int{127, 42, 42}
            _ = v10
            _ = v11
            _ = v9
        }
        p0 = 65
        var v12 float32 = 0.1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return -998
}
func genFn1() int8 {
    v13 := GenS2{F0: -2147483648, F1: 90, F2: 65535}
    var v14 uint64 = 0
    v14 = (v14 & (v14 | (v14 | v14)))
    _ = v13
    _ = v14
    return 100
}

func main() {
    var v1 uint64 = 42
    var v2 uint8 = 0
    var v3 string = ""
    var v4 uint = 255
    v5 := make(map[string]uint32)
    var v15 string = ("\n\t" + (v3 + v3))
    switch 9223372036854775807 {
    case 9223372036854775807:
        v16 := GenS3{F0: 9223372036854775807, F1: 65535, F2: 4294967295}
        _ = v16
    case 32768:
        v17 := GenS3{F0: -538, F1: 0, F2: 1}
        _ = v17
    case 537:
        var v18 uint32 = 1
        _ = v18
    default:
        v4 += v4
    }
    v2 = v2
    for fi19 := 0; fi19 < 18; fi19++ {
        if true {
            delete(v5, v3)
            var v20 float32 = (1e10 - 1e10)
            switch 308 {
            case 128:
                v4 -= uint(v1)
            }
            _ = v20
        }
        _ = fi19
    }
    v15 = "a"
    switch "a" {
    case "":
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int32 = 2147483647
            v2 = v2
            v23 := GenS4{F0: 42, F1: -1, F2: 2147483647}
            _ = fi21
            _ = v22
            _ = v23
        }
    case "foo bar":
        var v24 int = 32767
        _ = v24
    case "hello":
        var v25 bool = (!(v1 > v1))
        _ = v25
    default:
        v26 := v5[(v3 + v15)]
        _ = v26
    }
    v27 := GenS5{F0: (!false), F1: 0}
    var v28 uint32 = 2147483647
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS6{F0: true}
        if false {
            v3 = v3
            var v31 string = ""
            _ = v31
        } else {
            var v32 int = fi29
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v3 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
