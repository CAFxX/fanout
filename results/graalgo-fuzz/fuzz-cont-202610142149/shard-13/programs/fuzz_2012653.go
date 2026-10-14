package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0() int32 {
    switch 904 {
    case -1:
        for fi4 := 0; fi4 < 15; fi4++ {
            v5 := []int{(fi4 << ((255) & 63)), (fi4 % ((fi4) | 1))}
            var v6 uint16 = 2
            _ = fi4
            _ = v5
            _ = v6
        }
    case 1:
        var v7 int32 = 0
        _ = v7
    default:
        var v8 uint = 179
        _ = v8
    }
    var v9 float64 = 0.0
    v10 := GenS1{F0: true, F1: 42}
    var v11 int16 = 1000
    v11 -= (v11 * v11)
    _ = v10
    _ = v11
    _ = v9
    return -940
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint32 = 65535
    v3 := []uint32{v2, 2147483647, v2, 0}
    v2 += 2147483648
    switch "\n\t" {
    case "":
        var v12 uint16 = 207
        _ = v12
    case "hello":
        fmt.Println(v2)
    case "a":
        var v13 float64 = float64((((float32(-2.25) * 0.1) - 3.14159) / -2.25))
        _ = v13
    }
    var v14 bool = (42 > 0)
    v15 := GenS2{F0: false, F1: 255, F2: 1}
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            if false {
                fmt.Println(v2, v14, fi17)
            } else {
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi17)
                    fmt.Println(v2, v14, fi16)
                }
            }
            v18 := GenS3{F0: fi16}
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: true, F1: 139}
        _ = len(v3)
        _ = fi16
        _ = v19
    }
    var v20 string = "\n\t"
    for fi21 := 0; fi21 < 15; fi21++ {
        v2 = 2147483648
        _ = fi21
    }
    fmt.Println(v2)
    if v14 {
        v22 := GenS2{F0: (!v14), F1: 255, F2: 128}
        if true {
            var v23 uint64 = 923
            for fi24 := 0; fi24 < 13; fi24++ {
                if (1 >= 4294967295) {
                    v25 := GenS3{F0: 128}
                    var v26 int8 = -115
                    var v27 int = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
            _ = v23
        }
        var v28 float64 = 0.1
        _ = v22
        _ = v28
    } else {
        if false {
            fmt.Println(v2, v20)
        }
        v29 := GenS2{F0: (v20 == v20), F1: 1, F2: uint8(v2)}
        _ = v29
    }
    var v30 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v30
}
