package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 int32; F1 int }
func genFn0() int {
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(fi5)
        v6 := GenS1{F0: (!true), F1: 1}
        _ = fi5
        _ = v6
    }
    var v7 float64 = -2.25
    _ = v7
    return 316
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 42
    v3 := []uint32{597, 4294967295}
    var v4 float64 = -2.25
    var v8 int8 = -113
    if false {
        v9 := GenS2{F0: (v8 << (((v2 << ((18446744073709551615) & 63))) & 63)), F1: -1, F2: -806}
        if ((1 <= 2) || false) {
            v10 := GenS3{F0: 42, F1: 42}
            v8 = v8
            _ = v10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = (fi11 * (fi11 & fi11))
            if (0 > 42) {
                _ = len(v3)
            }
            _ = fi11
            _ = v12
        }
        _ = v9
    } else {
        switch -542 {
        case 2147483647:
            v13 := GenS3{F0: -373, F1: -861}
            _ = v13
        case 787:
            v14 := GenS1{F0: false, F1: 18446744073709551615}
            _ = v14
        case 0:
            v15 := GenS3{F0: 1, F1: -636}
            _ = v15
        default:
            v8 -= v8
        }
        var v16 int16 = -32768
        _ = v16
    }
    fmt.Println(v2, v8)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 += v8
    v3 = append(v3, 990)
    v2 = v2
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint16 = 255
        fmt.Println(v8, v8)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
        _ = v18
    }
    var v19 int16 = 0
    v1[42] = 61
    if true {
        if true {
            if false {
                v2 += v2
                if (!(true && (!(!((0.0 - 0.1) == 1e10))))) {
                    v2 = 65535
                }
                switch -2147483648 {
                case 127:
                    v20 := GenS1{F0: true, F1: 2}
                    _ = v20
                case 32767:
                    var v21 uint = (v2 / (((v2 >> ((v2) & 63))) | 1))
                    _ = v21
                case 727:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            if (v19 >= v19) {
                fmt.Println(v8)
                fmt.Println(v2)
            }
        }
    }
    fmt.Println(v4, v8)
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
