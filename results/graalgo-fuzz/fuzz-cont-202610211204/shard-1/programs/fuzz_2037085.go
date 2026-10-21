package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    fmt.Println(v4)
    v5 := GenS1{F0: 65535, F1: 42, F2: (!true)}
    v6 := GenS1{F0: 255, F1: -1000, F2: false}
    var v7 bool = (true || (32767 >= 1))
    v8 := GenS1{F0: 90, F1: 782, F2: (!v7)}
    var v9 int64 = -950
    switch 442 {
    case 743:
        var v10 bool = true
        _ = v10
    default:
        fmt.Println(v7)
    }
    v11 := GenS1{F0: 510, F1: -11, F2: (!(!(v7 && v7)))}
    var v12 uint64 = 1
    if false {
        var v13 string = "\n\t"
        v14 := GenS2{F0: 4294967295}
        _ = v13
        _ = v14
    } else {
        var v15 float32 = (1e10 * (3.14159 / 1.0))
        if (float32((1.5 * 1e10)) > 3.14159) {
            if (!false) {
                v7 = (v7 && v7)
            }
            v7 = v7
        } else {
            v15 -= v15
            for fi16 := 0; fi16 < 2; fi16++ {
                if v7 {
                    v17 := GenS2{F0: (v4 % (((v4 & (^(v4)))) | 1))}
                    _ = v17
                }
                _ = fi16
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v7 = (42 <= v12)
                    _ = fi19
                }
                fmt.Println(fi18)
                var v20 int64 = -569
                _ = fi18
                _ = v20
            }
        }
        _ = v15
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS3{F0: v12, F1: (42 != 255)}
        _ = fi21
        _ = v22
    }
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
