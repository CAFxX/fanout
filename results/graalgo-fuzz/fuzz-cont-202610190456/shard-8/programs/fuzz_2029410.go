package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS3 struct { F0 int }
func genFn0() bool {
    var v7 uint32 = 4294967295
    fmt.Println(v7)
    v8 := GenS1{F0: -1, F1: true, F2: "hello"}
    v9 := GenS2{F0: 4294967295, F1: 0, F2: 1000}
    var v10 float64 = -2.25
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 uint64, p1 uint32) int32 {
    p0 = p0
    var v11 int16 = -473
    p1 = p1
    if (!(true || true)) {
        p1 = p1
    } else {
        var v12 int8 = -1
        switch "foo bar" {
        case "hello":
            var v13 float64 = 1.5
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 12; fi14++ {
                p0 = (p0 + p0)
                _ = fi14
            }
        case "":
            var v15 int = -1
            _ = v15
        }
        p0 = uint64(p1)
        _ = v12
    }
    v16 := []int{0, 127, 32767, 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v16
    return 0
}

func main() {
    var v1 uint8 = 1
    var v2 int = 32767
    v3 := make(map[int]bool)
    var v4 int64 = 0
    var v5 float32 = (1e10 + float32((1e10 - 0.1)))
    v6 := make(map[string]bool)
    var v17 uint16 = 634
    if true {
        v4 = (^(v4))
        var v18 int16 = 790
        v18 = -1
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            switch 1 {
            case 2147483647:
                var v21 int16 = 42
                _ = v21
            case -594:
                v22 := GenS1{F0: -53, F1: true, F2: "a"}
                _ = v22
            default:
                v23 := v6["a"]
                _ = v23
            }
            _ = fi20
        }
        fmt.Println(fi19, v4, v17)
        _ = fi19
    }
    v24 := v6[""]
    v4 = -1
    if v24 {
        v5 = 1.5
        var v25 string = "hello"
        var v26 int = -410
        _ = v25
        _ = v26
    } else {
        v4 -= v4
        v17 = v17
        var v27 uint = 2
        _ = v27
    }
    v28 := GenS3{F0: (v2 << ((1) & 63))}
    if (genFn1(2, 568) != 42) {
        var v29 uint8 = 2
        _ = v29
    }
    var v30 uint32 = 737
    v31 := GenS2{F0: 0, F1: 127, F2: -390}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
