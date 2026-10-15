package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 uint8 = 0
    fmt.Println(v5)
    _ = len(v4)
    v1 = (0 > 4294967295)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = -9223372036854775808
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 string = ("\n\t" + "\n\t")
            if v1 {
                var v10 int = fi6
                var v11 string = v9
                _ = v10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        var v12 bool = (v1 && false)
        _ = fi6
        _ = v12
        _ = v7
    }
    v1 = v1
    switch -336 {
    case -2147483648:
        v5 -= v5
    case 4294967295:
        switch -544 {
        case 32768:
            v1 = ((false || v1) || (32767 <= 1))
        case 0:
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS1{F0: 93, F1: 637, F2: 32767}
                v1 = v1
                v15 := GenS2{F0: 42, F1: 628}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    case 127:
        delete(v2, -573)
    }
    v5 -= v5
    v16 := GenS1{F0: (v5 & v5), F1: 1, F2: 32767}
    var v17 uint32 = 1
    var v18 uint8 = 0
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := []int{}
        v21 := GenS2{F0: (^(v17)), F1: 4294967295}
        switch (fi19 ^ fi19) {
        case 128:
            var v22 float32 = (-2.25 - (1e10 + (0.0 * -2.25)))
            _ = v22
        case 1:
            var v23 uint = 851
            _ = v23
        case -2147483648:
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := GenS1{F0: v18, F1: 18446744073709551615, F2: 1000}
                v3["hello"] = ("hello" >= "")
                switch fi24 {
                case 0:
                    var v26 uint = 4294967295
                    _ = v26
                case -824:
                    fmt.Println(v5)
                    var v27 uint16 = 581
                    _ = v27
                case 1:
                    switch "a" {
                    case "a":
                        v5 += (v18 + v18)
                    default:
                        var v28 int16 = 42
                        _ = v28
                    }
                }
                _ = fi24
                _ = v25
            }
        default:
            v29 := GenS3{F0: 158}
            _ = v29
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v18 = v18
    if (544 < 180) {
        v5 = 61
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := v4[-988]
            var v32 uint = uint(v17)
            _ = fi30
            _ = v31
            _ = v32
        }
        v18 = 150
    } else {
        fmt.Println(v17, v1, v18)
    }
    v18 = (v18 % (((^(v18))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
