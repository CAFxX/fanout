package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
func genFn0(p0 int64, p1 uint, p2 int64) uint32 {
    switch "foo bar" {
    case "a":
        if false {
            v7 := []int{32768, -2147483648, -9223372036854775808, 32767}
            _ = v7
        } else {
            if (true && (!false)) {
                var v8 uint32 = 2147483647
                var v9 int64 = p0
                switch -717 {
                case 127:
                    v10 := []int{-1, -73, -2147483648}
                    _ = v10
                case 32768:
                    var v11 uint64 = 2
                    _ = v11
                default:
                    v12 := GenS1{F0: 4294967295, F1: 1, F2: 65535}
                    _ = v12
                }
                _ = v8
                _ = v9
            } else {
                fmt.Println(p1)
            }
            v13 := GenS1{F0: 0, F1: 2147483647, F2: 840}
            _ = v13
        }
    }
    var v14 float64 = 1e10
    if (65535 >= 521) {
        var v15 uint64 = 964
        var v16 int32 = -1
        _ = v15
        _ = v16
    } else {
        p2 = p2
        var v17 uint64 = 2
        _ = v17
    }
    switch -2147483648 {
    case 42:
        switch 294 {
        case -2147483648:
            v18 := []int{32768, 0, 356, -533}
            _ = v18
        case 42:
            var v19 int8 = 1
            _ = v19
        }
    default:
        p0 -= (p2 - (p2 % ((p2) | 1)))
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 2147483648
}
func genFn1(p0 bool, p1 int64) int16 {
    var v20 int = 1
    switch int(p1) {
    case 9223372036854775807:
        p0 = p0
    }
    v21 := []int{v20, (v20 >> ((65535) & 63)), (v20 & v20), (v20 << ((207) & 63))}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 1
}

func main() {
    v1 := []int{-627, 9223372036854775807, 4294967295, 42}
    var v2 uint = 0
    var v3 int32 = -1
    var v4 int8 = 100
    var v5 uint = (^(v2))
    var v6 uint32 = 773
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    v5 = (v2 / (((v5 | 892)) | 1))
    var v22 uint16 = 255
    var v23 float32 = 1.5
    var v24 int = 784
    v4 = 0
    var v25 int8 = -1
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v26 := GenS1{F0: 1, F1: 1, F2: v22}
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v2, v3)
        _ = fi27
    }
    switch "a" {
    case "a":
        v28 := GenS1{F0: 236, F1: (v3 << (((v2 << (((v2 ^ 752)) & 63))) & 63)), F2: v22}
        _ = v28
    case "\n\t":
        if (false && true) {
            var v29 uint = v2
            _ = v29
        }
    }
    switch "" {
    case "":
        var v30 int = (v24 << (((^(v5))) & 63))
        _ = v30
    case "hello":
        if false {
            for fi31 := 0; fi31 < 20; fi31++ {
                for fi32 := 0; fi32 < 19; fi32++ {
                    _ = len(v1)
                    v5 -= (v2 << (((v5 ^ 813)) & 63))
                    fmt.Println(v25, v25, v5)
                    _ = fi32
                }
                _ = fi31
            }
        }
    case "foo bar":
        var v33 int = -913
        _ = v33
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
