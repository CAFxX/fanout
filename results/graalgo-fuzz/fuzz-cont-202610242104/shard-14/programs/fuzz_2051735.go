package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint) int64 {
    v4 := []int{}
    v5 := GenS1{F0: 80, F1: 275, F2: 364}
    _ = p0
    _ = v4
    _ = v5
    return 175
}
func genFn1(p0 uint, p1 uint, p2 int) uint {
    for fi6 := 0; fi6 < 13; fi6++ {
        switch "hello" {
        case "a":
            p2 = (fi6 - -9223372036854775808)
        case "":
            for fi7 := 0; fi7 < 15; fi7++ {
                p0 = p0
                p2 = fi6
                _ = fi7
            }
        case "\n\t":
            var v8 int = 127
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS2{F0: (^(p1)), F1: 583, F2: 2}
        v11 := GenS2{F0: p1, F1: 99, F2: 65535}
        fmt.Println(p1, p0)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 2
    v13 := []int{(p2 - 2147483647), 32767, (^(p2)), (p2 - p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p0
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]uint32)
    v3 := []string{""}
    _ = len(v3)
    v1 -= (^(v1))
    v3 = append(v3, ("a" + ("hello" + "foo bar")))
    var v14 float64 = 0.1
    v14 = v14
    v15 := GenS1{F0: -128, F1: 65535, F2: 0}
    v16 := v2[-422]
    v1 = v1
    fmt.Println(v14, v1)
    var v17 int64 = 1
    v18 := GenS1{F0: -128, F1: v16, F2: 1}
    fmt.Println(v1, v16, v16)
    v2[42] = 1
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS3{F0: fi19}
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int64 = v17
            if (!false) {
                v23 := GenS3{F0: fi19}
                fmt.Println(v22)
                var v24 uint16 = 255
                fmt.Println(v17, fi19, v17)
                v24 = v24
                _ = v23
                _ = v24
            } else {
                v25 := v2[2147483647]
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        var v26 uint32 = v16
        _ = fi19
        _ = v20
        _ = v26
    }
    switch "\n\t" {
    case "foo bar":
        _ = len(v2)
    case "hello":
        v27 := GenS2{F0: 607, F1: 42, F2: 32767}
        _ = v27
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int8 = 100
            v14 = v14
            _ = fi28
            _ = v29
        }
    }
    _ = len(v3)
    var v30 int16 = 386
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
}
