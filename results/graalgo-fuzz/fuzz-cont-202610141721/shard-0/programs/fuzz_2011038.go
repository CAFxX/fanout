package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
func genFn0(p0 uint64, p1 uint8, p2 uint8) int8 {
    v5 := []int{940, 32767}
    var v6 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint32 {
    if ((false && true) && (!true)) {
        var v7 int64 = -9223372036854775808
        switch 128 {
        case -181:
            var v8 uint16 = 65535
            _ = v8
        case 2147483647:
            v9 := GenS1{F0: 4294967295, F1: v7}
            _ = v9
        default:
            if true {
                fmt.Println(v7)
            }
        }
        fmt.Println(v7)
        _ = v7
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint32 = 2147483647
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: 311, F1: 1}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    return 227
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (3.14159 + -2.25)
    v3 := make(map[uint32]bool)
    var v4 int = 861
    if false {
        v4 = 32767
        if (!false) {
            switch v4 {
            case 32767:
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 uint32 = 0
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS2{F0: 1, F1: 0, F2: 0}
                        var v18 float32 = -2.25
                        var v19 int32 = -2147483648
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
            }
            fmt.Println(v2, v4)
        }
    } else {
        var v20 uint16 = 2
        fmt.Println(v20, v4, v2)
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 float64 = v2
        v2 = ((v22 + float64(-2.25)) + float64(float32((v22 + v22))))
        _ = fi21
        _ = v22
    }
    fmt.Println(v2, v2)
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 float64 = (-2.25 / v2)
            fmt.Println(v2, fi23, v2)
            _ = fi23
            _ = v24
        }
    }
    var v25 uint = 309
    fmt.Println(v25, v4, v2)
    v26 := GenS1{F0: (v25 + 797), F1: 771}
    v27 := v1[15]
    v2 -= float64(1.5)
    var v28 bool = true
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 uint8 = 63
        v25 = v25
        _ = fi29
        _ = v30
    }
    fmt.Println(v25)
    var v31 uint16 = 32767
    var v32 string = v27
    v33 := GenS3{F0: v4, F1: 0, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
