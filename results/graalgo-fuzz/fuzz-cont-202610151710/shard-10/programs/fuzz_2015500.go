package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 string }

func main() {
    var v1 int64 = -1
    var v2 uint8 = 1
    var v3 uint64 = 2
    v4 := make(map[uint32]int)
    fmt.Println(v1, v1)
    switch "foo bar" {
    case "":
        v2 = v2
    case "hello":
        var v5 uint = 835
        _ = v5
    case "foo bar":
        switch "" {
        case "foo bar":
            v1 += 1
        case "a":
            v6 := GenS1{F0: -276, F1: 153, F2: 42}
            _ = v6
        case "hello":
            v2 = v2
        default:
            var v7 string = "\n\t"
            _ = v7
        }
    }
    var v8 uint = 822
    var v9 uint32 = 2147483648
    v4[(v9 + v9)] = 2147483647
    v10 := []int{4294967295, 42, 32768, 127}
    switch 32767 {
    case 128:
        switch -9223372036854775808 {
        case 393:
            for fi11 := 0; fi11 < 16; fi11++ {
                v9 = v9
                _ = fi11
            }
        case -190:
            v3 += v3
        case 9223372036854775807:
            switch 906 {
            case 42:
                if len(v10) > 0 {
                    v10[0] = 32768
                }
            case -168:
                v12 := GenS2{F0: -1, F1: -1, F2: v3}
                _ = v12
            }
        }
    case 42:
        v1 += v1
    }
    fmt.Println(v9)
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: 986, F1: 111, F2: uint64(v8)}
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: "a"}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 float32 = 1.5
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int32 = -443
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int32 = 359
            _ = fi20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    delete(v4, (v9 % ((0) | 1)))
    for fi22 := 0; fi22 < 2; fi22++ {
        switch fi22 {
        case 202:
            fmt.Println(v3, v3)
        case 128:
            var v23 int32 = 1
            _ = v23
        default:
            var v24 float64 = (0.0 + ((1.5 / 1e10) - 0.0))
            _ = v24
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := v4[v9]
            _ = fi25
            _ = v26
        }
        _ = fi22
    }
    var v27 int64 = (v1 / ((v1) | 1))
    var v28 bool = ((v1 + v27) > v27)
    var v29 uint64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
