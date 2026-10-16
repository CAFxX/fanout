package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 int16) uint64 {
    v6 := []int{}
    p0 = p0
    _ = p0
    _ = v6
    return 145
}
func genFn1(p0 uint8, p1 uint) uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch ("" + ("\n\t" + "foo bar")) {
        case "hello":
            if true {
                fmt.Println(p1, p0, p0)
                fmt.Println(p1, p1)
                if (3.14159 < 1e10) {
                    var v8 bool = false
                    v9 := GenS1{F0: 42, F1: 0, F2: 2}
                    _ = v8
                    _ = v9
                } else {
                    if (129 != 205) {
                        v10 := []int{2147483647}
                        v11 := GenS2{F0: 0}
                        var v12 int8 = 105
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 3; fi13++ {
                        v14 := []int{(fi7 >> (((^(p1))) & 63))}
                        _ = fi13
                        _ = v14
                    }
                    p0 = 0
                }
            } else {
                v15 := GenS2{F0: 0}
                _ = v15
            }
        case "a":
            p1 = (p1 / (((p1 - p1)) | 1))
        case "\n\t":
            var v16 uint16 = 0
            _ = v16
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                switch "hello" {
                case "a":
                    v18 := GenS3{F0: fi17}
                    _ = v18
                case "\n\t":
                    v19 := GenS3{F0: (fi17 ^ fi7)}
                    _ = v19
                case "foo bar":
                    var v20 uint16 = 65535
                    _ = v20
                }
                var v21 int = fi7
                var v22 float64 = float64((3.14159 - 1.5))
                _ = fi17
                _ = v21
                _ = v22
            }
        }
        _ = fi7
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS1{F0: int64(p0), F1: 65535, F2: p0}
        var v25 int16 = 423
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int64 = 42
    _ = p0
    _ = p1
    _ = v26
    return 720
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -410
    var v3 uint = 460
    var v4 bool = (0 > -1)
    var v5 int32 = -1
    fmt.Println(v1)
    var v27 uint8 = 42
    v28 := []int{-2147483648, 127, 42}
    var v29 uint64 = 934
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := GenS4{F0: "\n\t", F1: 137, F2: 1}
        _ = fi30
        _ = v31
    }
    switch 657 {
    case 0:
        v32 := GenS1{F0: 9223372036854775807, F1: 32767, F2: v27}
        _ = v32
    case 899:
        var v33 int = 1
        _ = v33
    default:
        v27 = (^(v27))
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v34 := GenS5{F0: v4, F1: 60}
    var v35 float64 = float64(v1)
    v27 -= (v27 ^ v27)
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
