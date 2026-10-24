package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string, p1 uint64) uint64 {
    v6 := []int{2147483647}
    for fi7 := 0; fi7 < 17; fi7++ {
        switch fi7 {
        case 4294967295:
            var v8 int64 = 754
            _ = v8
        default:
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint32 = 65535
                v6 = append(v6, fi7)
                _ = fi9
                _ = v10
            }
        }
        var v11 uint16 = 2
        fmt.Println(p0, v11, fi7)
        _ = fi7
        _ = v11
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 uint) uint16 {
    var v12 bool = false
    fmt.Println(p0)
    var v13 int16 = 577
    v12 = (v12 || ((^(p0)) == p0))
    fmt.Println(p0, v12)
    _ = p0
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 int16 = -746
    v2 := []string{"hello", ("\n\t" + "\n\t"), "foo bar", ""}
    v3 := []uint32{2147483648, 2147483647, 1, 2147483647}
    var v4 bool = (1.5 == 3.14159)
    var v5 uint8 = 42
    fmt.Println(v4, v1)
    if false {
        fmt.Println(v5)
    } else {
        var v14 uint = 18446744073709551615
        v4 = (!v4)
        v2 = append(v2, ("" + "hello"))
        _ = v14
    }
    switch "a" {
    case "a":
        switch "a" {
        case "":
            v15 := GenS1{F0: 42, F1: 0}
            _ = v15
        }
    case "hello":
        v5 = v5
    case "":
        if v4 {
            fmt.Println(v4)
            if (42 >= -433) {
                v16 := GenS1{F0: 65535, F1: 2147483647}
                _ = v16
            } else {
                v1 -= v1
            }
            v5 = v5
        }
    default:
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS2{F0: 1, F1: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS1{F0: 1, F1: 2147483648}
    if v4 {
        if v4 {
            switch 4294967295 {
            case 42:
                var v20 uint64 = genFn0("\n\t", 238)
                _ = v20
            default:
                fmt.Println(v5, v5)
            }
            v4 = (!v4)
        }
    } else {
        if v4 {
            fmt.Println(v5, v1)
            fmt.Println(v1, v5, v5)
        }
        v21 := GenS2{F0: 435, F1: -958}
        _ = v21
    }
    var v22 uint = 1
    v23 := GenS3{F0: 42}
    v3 = append(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
