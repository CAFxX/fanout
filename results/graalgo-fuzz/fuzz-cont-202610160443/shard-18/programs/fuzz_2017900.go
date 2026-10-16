package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }
func genFn0(p0 uint64, p1 uint, p2 bool) int8 {
    if ((0.1 != float64(1.5)) || true) {
        var v6 uint = p1
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: 75, F1: 394, F2: 2}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    }
    var v10 int64 = 1
    if p2 {
        p0 += 1
    } else {
        p0 -= p0
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int32 = 299
        v13 := GenS2{F0: 24, F1: -1}
        var v14 uint16 = 32767
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            p1 += (p1 / ((p1) | 1))
            _ = fi16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            switch "hello" {
            case "":
                p1 -= 2
            case "foo bar":
                switch "\n\t" {
                case "a":
                    var v18 float32 = 1.5
                    _ = v18
                case "\n\t":
                    p1 -= p1
                case "":
                    var v19 uint32 = 2
                    _ = v19
                }
            default:
                v20 := []int{42, -566, -1}
                _ = v20
            }
            _ = fi17
        }
        v21 := GenS3{F0: 1000}
        _ = fi15
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return int8(v10)
}
func genFn1(p0 uint16, p1 int8, p2 int16) uint32 {
    v22 := []int{}
    v23 := GenS4{F0: 2147483647, F1: 42, F2: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 709
}

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 827
    var v3 string = "a"
    var v4 string = v3
    var v5 uint8 = 128
    var v24 int32 = -2147483648
    v5 = 82
    var v25 uint = 18446744073709551615
    for fi26 := 0; fi26 < 5; fi26++ {
        v25 = uint(v5)
        v25 = 269
        v27 := GenS3{F0: (v2 ^ -249)}
        _ = fi26
        _ = v27
    }
    var v28 uint8 = v5
    var v29 uint16 = 42
    v30 := GenS3{F0: -872}
    var v31 int64 = 0
    if false {
        var v32 uint16 = uint16(v25)
        var v33 uint8 = 221
        fmt.Println(v24, v5)
        _ = v32
        _ = v33
    } else {
        v29 = v29
        v29 = (v29 / (((v29 << ((v25) & 63))) | 1))
        var v34 int64 = -9223372036854775808
        _ = v34
    }
    v35 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
