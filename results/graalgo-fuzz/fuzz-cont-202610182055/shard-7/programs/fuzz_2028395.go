package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 uint16, p1 int16) int16 {
    v6 := GenS1{F0: 65535}
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{(fi8 / ((fi8) | 1)), fi8}
            v9 = append(v9, (fi8 * fi7))
            _ = fi8
            _ = v9
        }
        var v10 float32 = 0.1
        v11 := []int{fi7, 42, 748, fi7}
        _ = fi7
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return (p1 * 1000)
}
func genFn1(p0 uint32, p1 bool) int32 {
    var v12 uint16 = 255
    p1 = (p1 || p1)
    v13 := []int{1, 922, -1, 0}
    var v14 uint64 = 0
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 988
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 int32 = -2147483648
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    switch 2147483647 {
    case 42:
        delete(v1, "foo bar")
    default:
        v15 := GenS2{F0: 255, F1: 0}
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint8 = 236
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 42
    if true {
        if false {
            switch int(v3) {
            case -679:
                if true {
                    var v19 float32 = ((3.14159 + -2.25) - 1e10)
                    _ = v19
                }
            case 1:
                var v20 uint = 18446744073709551615
                _ = v20
            case 32768:
                switch -9223372036854775808 {
                case 2147483647:
                    fmt.Println(v3, v18)
                default:
                    fmt.Println(v3, v3)
                }
            default:
                v18 = v18
            }
            var v21 int16 = genFn0(255, 1)
            var v22 uint32 = 1
            _ = v21
            _ = v22
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v23 int32 = v3
            v24 := GenS2{F0: 32767, F1: 1}
            _ = v23
            _ = v24
        }
    } else {
        v18 = v18
        v18 += (v18 - 255)
        v4["hello"] = false
    }
    v18 = (^(v18))
    switch "a" {
    case "\n\t":
        var v25 int64 = 105
        _ = v25
    }
    v18 = 0
    if true {
        var v26 uint32 = 0
        _ = v26
    }
    if (0.0 == -2.25) {
        v18 = (v18 << ((1) & 63))
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v27 := GenS3{F0: uint64(v18), F1: genFn1(4294967295, true), F2: -908}
        _ = v27
    }
    var v28 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
