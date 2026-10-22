package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
func genFn0(p0 uint64, p1 int64) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{4294967295, (fi5 / (((fi5 & fi5)) | 1)), (fi5 / ((-582) | 1)), (fi5 / ((202) | 1))}
        var v7 uint8 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: -441, F1: (143 == -32768), F2: -2147483648}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -927, F1: true, F2: 42}
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (1 != -101)
    var v3 uint64 = 0
    v4 := make(map[int]int)
    if v2 {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int8 = 0
            var v13 float64 = 1.5
            var v14 string = "\n\t"
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        switch -318 {
        case -18:
            v2 = true
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            v3 = (^(v3))
            _ = fi15
        }
        v16 := []int{32767, 956, 1, 383}
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int16 = -489
        v19 := []int{fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if true {
            v21 := GenS1{F0: 32767, F1: true, F2: -2147483648}
            fmt.Println(v2, v3)
            var v22 int16 = 42
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    switch -1 {
    case -682:
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 int32 = -796
            _ = fi23
            _ = v24
        }
    case 9223372036854775807:
        var v25 uint64 = 2
        _ = v25
    case -476:
        v26 := []int{-390, 32767}
        _ = v26
    default:
        var v27 uint8 = 127
        _ = v27
    }
    v28 := []int{1, 42, 28}
    if v2 {
        switch 479 {
        case 127:
            switch "foo bar" {
            case "hello":
                v28 = append(v28, -116)
            }
        }
        var v29 uint = 962
        _ = v29
    }
    var v30 uint64 = 18446744073709551615
    var v31 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
