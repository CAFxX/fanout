package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint, p1 uint16) uint32 {
    v7 := GenS1{F0: p0}
    var v8 uint16 = p1
    for fi9 := 0; fi9 < 13; fi9++ {
        switch ("" + (("\n\t" + "a") + "")) {
        case "a":
            var v10 int8 = 0
            _ = v10
        default:
            v11 := GenS1{F0: (^(p0))}
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 bool) int64 {
    if true {
        var v12 int32 = 42
        fmt.Println(v12, p0)
        _ = v12
    }
    v13 := []int{127, 128}
    var v14 bool = p0
    v15 := GenS1{F0: 1}
    v16 := GenS1{F0: 2}
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    v3 := make(map[int]uint32)
    var v4 int8 = 42
    var v5 bool = true
    v6 := []uint32{}
    if true {
        switch 786 {
        case 0:
            var v17 float64 = 1.5
            _ = v17
        default:
            var v18 float64 = 3.14159
            _ = v18
        }
    }
    var v19 uint64 = 0
    var v20 int16 = 68
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v21 := GenS1{F0: 0}
        _ = v21
    }
    v5 = v5
    v5 = false
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            switch 2147483647 {
            case 128:
                var v24 uint16 = 42
                _ = v24
            case 1:
                var v25 float32 = 3.14159
                _ = v25
            case 2147483647:
                var v26 uint64 = v19
                _ = v26
            default:
                var v27 uint32 = 2
                var v28 uint = 915
                v29 := v3[32768]
                v2 = v2
                v19 += 18446744073709551615
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi23
        }
        if (0 >= 1) {
            var v30 int16 = v20
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 uint = 1
                var v33 int64 = 1
                var v34 int64 = 900
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = v30
        }
        if (42 != -1) {
            v35 := GenS2{F0: "hello"}
            _ = v35
        }
        _ = fi22
    }
    v19 = v19
    v3[2147483647] = 2147483648
    v36 := GenS3{F0: 840, F1: 1}
    _ = len(v6)
    v2 = (v2 | 30)
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
