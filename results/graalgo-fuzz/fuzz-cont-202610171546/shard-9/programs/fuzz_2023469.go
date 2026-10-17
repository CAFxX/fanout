package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8) int16 {
    v7 := GenS1{F0: 166, F1: 1}
    var v8 uint32 = uint32(p0)
    var v9 uint64 = 4294967295
    var v10 uint64 = (v9 / (((v9 >> ((42) & 63))) | 1))
    if false {
        v8 = (v8 / (((v8 << ((1) & 63))) | 1))
        fmt.Println(v8)
        if false {
            v11 := []int{1, 4294967295}
            switch "hello" {
            case "hello":
                var v12 float32 = 0.0
                _ = v12
            case "foo bar":
                v9 = (v10 % (((v10 >> ((356) & 63))) | 1))
            case "":
                var v13 uint8 = 0
                _ = v13
            default:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            var v14 int64 = 42
            _ = v11
            _ = v14
        }
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int32, p1 int16, p2 uint8) uint64 {
    v15 := GenS2{F0: -1, F1: (p2 / (((p2 - (^(p2)))) | 1)), F2: uint(p2)}
    fmt.Println(p0)
    p1 = (p1 / ((genFn0(p2)) | 1))
    v16 := GenS1{F0: 2147483647, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 479
}

func main() {
    var v1 string = ((("a" + "a") + "\n\t") + ("hello" + "\n\t"))
    var v2 float64 = ((((3.14159 * 1.5) - 1.5) / (1.5 * (float64(1.5) / ((((float64(((3.14159 * 1e10) / 1e10)) * (0.1 * (1.5 * 0.0))) - ((1e10 / 1e10) / 3.14159)) - (1.5 + 1.5)) - (float64(0.0) - 0.1))))) + 3.14159)
    v3 := make(map[uint32]string)
    var v4 uint64 = 4294967295
    v5 := []int{127, 42, 195}
    var v6 int16 = -1
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS3{F0: (fi17 % ((2147483647) | 1)), F1: 127}
        _ = fi17
        _ = v18
    }
    v5 = append(v5, -9223372036854775808)
    switch v1 {
    case "":
        fmt.Println(v4, v4)
    case "hello":
        v1 = ("hello" + v1)
    default:
        if (false && (true && (true || false))) {
            v4 -= v4
        }
    }
    var v19 int16 = -467
    var v20 string = v1
    v20 = v1
    v21 := GenS4{F0: 42}
    switch -770 {
    case 42:
        fmt.Println(v6)
    case 128:
        fmt.Println(v4, v19)
    }
    var v22 uint32 = 42
    for fi23 := 0; fi23 < 16; fi23++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v19 = 42
        _ = fi23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v25 := GenS1{F0: (v22 >> ((2) & 63)), F1: -1}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v20, fi26)
        _ = fi26
    }
    fmt.Println(v19)
    var v27 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
