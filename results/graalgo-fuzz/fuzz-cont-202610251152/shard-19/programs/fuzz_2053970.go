package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0(p0 uint, p1 string, p2 string) int16 {
    v6 := []int{-600, 42, int(p0), 42}
    var v7 int32 = -2147483648
    v8 := GenS1{F0: 4294967295, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -32768
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 394, 294}
    v2 = (509 != 2147483647)
    for fi9 := 0; fi9 < 6; fi9++ {
        if (850 < 42) {
            var v10 int32 = 0
            _ = v10
        } else {
            fmt.Println(fi9, v2)
            var v11 int32 = -1
            _ = v11
        }
        _ = fi9
    }
    if (!false) {
        if true {
            v2 = v2
            v12 := v3[4294967295]
            _ = v12
        } else {
            var v13 int64 = -1
            v2 = (v2 || v2)
            var v14 int8 = 42
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: 0, F1: 42}
    if ((3.14159 * 3.14159) >= (3.14159 / 3.14159)) {
        switch 32768 {
        case 9223372036854775807:
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int8 = -1
                delete(v4, 99)
                if false {
                    if (0.1 != 3.14159) {
                        var v18 int32 = 0
                        fmt.Println(v2)
                        _ = v18
                    }
                }
                _ = fi16
                _ = v17
            }
        }
    }
    switch "a" {
    case "foo bar":
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS2{F0: 0, F1: 0, F2: -32768}
            for fi21 := 0; fi21 < 19; fi21++ {
                v2 = (42 < -32768)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    fmt.Println(v2)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int32 = -2147483648
        v24 := GenS1{F0: fi22, F1: 252}
        _ = fi22
        _ = v23
        _ = v24
    }
    delete(v4, 63)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
