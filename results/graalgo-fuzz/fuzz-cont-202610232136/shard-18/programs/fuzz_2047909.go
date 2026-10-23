package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 bool, p1 uint8) int16 {
    var v5 int32 = -836
    var v6 uint64 = uint64(p1)
    v7 := []int{0, -849, 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{}
    var v3 bool = (((v1 >> ((v1) & 63)) < (^(v1))) && ((true && false) || (!((128 == 74) || false))))
    var v4 int64 = 42
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int8 = -1
        v11 := GenS1{F0: fi9, F1: 806}
        var v12 string = "foo bar"
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if v3 {
        v4 -= -1
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint = 65535
                _ = len(v2)
                v3 = (!(v8 == v8))
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        fmt.Println(v3, v8)
    } else {
        v8 = v8
        v1 -= 309
    }
    v16 := GenS1{F0: 0, F1: -239}
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint16 = 255
        for fi19 := 0; fi19 < 16; fi19++ {
            v18 = (^(v18))
            v20 := GenS1{F0: (fi19 << (((v1 % ((416) | 1))) & 63)), F1: -586}
            switch fi19 {
            case 128:
                var v21 float32 = (((-2.25 + (0.0 - (-2.25 - 1.5))) - 0.1) / 1.5)
                _ = v21
            case 887:
                fmt.Println(fi19, v4)
            case -2147483648:
                v22 := GenS1{F0: 886, F1: 42}
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    var v23 string = "foo bar"
    v24 := GenS2{F0: (1e10 == 0.1), F1: (v4 << ((v1) & 63))}
    if len(v2) > 0 {
        v2[0] = 32768
    }
    var v25 uint = 1
    v8 -= (v8 | (^(v8)))
    v8 = v8
    var v26 uint8 = 211
    fmt.Println(v8, v3, v8)
    v1 = (v25 - (v1 >> ((v1) & 63)))
    fmt.Println(v23, v26)
    var v27 int16 = 1
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
