package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0() int64 {
    switch "foo bar" {
    case "a":
        var v4 uint8 = 127
        _ = v4
    case "\n\t":
        var v5 int = -904
        _ = v5
    case "foo bar":
        v6 := []int{}
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 995}
    v8 := GenS1{F0: 2, F1: 0}
    var v9 bool = true
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    v10 := GenS1{F0: 1, F1: 18446744073709551615}
    fmt.Println(v1)
    fmt.Println(v1)
    if false {
        if (!true) {
            var v11 uint = 562
            _ = v11
        }
        fmt.Println(v1)
        if (42 <= 18446744073709551615) {
            v12 := GenS2{F0: (2 >= 2147483647), F1: 0, F2: 4294967295}
            _ = len(v2)
            v1 -= 0.1
            _ = v12
        }
    }
    var v13 string = ""
    if (1 <= 127) {
        switch "\n\t" {
        case "":
            var v14 int16 = -1
            _ = v14
        }
        var v15 float32 = v1
        for fi16 := 0; fi16 < 18; fi16++ {
            v1 = v15
            _ = fi16
        }
        _ = v15
    } else {
        switch "a" {
        case "":
            v17 := GenS1{F0: 177, F1: 573}
            _ = v17
        case "\n\t":
            v18 := GenS2{F0: true, F1: 0, F2: 42}
            _ = v18
        default:
            var v19 int8 = -128
            _ = v19
        }
        var v20 uint16 = 65535
        switch -980 {
        case -96:
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint16 = v20
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch 1 {
                    case 0:
                        v24 := GenS2{F0: (v22 >= v20), F1: 0, F2: 42}
                        _ = v24
                    case 32767:
                        var v25 uint64 = 18446744073709551615
                        _ = v25
                    default:
                        var v26 float64 = 1.5
                        _ = v26
                    }
                    v3[v13] = 65535
                    _ = fi23
                }
                _ = fi21
                _ = v22
            }
        case 2147483647:
            for fi27 := 0; fi27 < 7; fi27++ {
                v20 = v20
                _ = fi27
            }
        default:
            var v28 int16 = 42
            _ = v28
        }
        _ = v20
    }
    v29 := GenS3{F0: 1, F1: -104}
    delete(v2, v13)
    v30 := GenS2{F0: false, F1: 1000, F2: 42}
    var v31 uint32 = 643
    var v32 uint = 42
    var v33 int64 = 0
    fmt.Println(v33, v1)
    fmt.Println(v33, v1, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
