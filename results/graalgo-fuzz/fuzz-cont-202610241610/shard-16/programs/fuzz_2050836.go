package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 int16 }
type GenS4 struct { F0 int; F1 int16; F2 uint }
type GenS5 struct { F0 bool; F1 int32; F2 int32 }
func genFn0() int {
    var v4 uint64 = 873
    v5 := []int{int(v4)}
    v4 += v4
    v6 := GenS1{F0: -998, F1: 9223372036854775807, F2: -1}
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1() uint32 {
    v7 := GenS2{F0: 32767, F1: 143}
    v8 := GenS3{F0: ("" + "a"), F1: 18446744073709551615, F2: -1}
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 989
    v2 := []int{127, 32767}
    var v3 bool = false
    v3 = ((42 > 2) && (0 <= v1))
    v1 -= v1
    if ((1.5 * 1.5) >= 0.1) {
        v3 = true
        _ = len(v2)
        v9 := GenS4{F0: 337, F1: -32768, F2: 2}
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = 32767
    }
    fmt.Println(v3, v1)
    if len(v2) > 0 {
        v2[0] = -26
    }
    v10 := GenS5{F0: v3, F1: 847, F2: -1}
    var v11 int64 = -9223372036854775808
    switch -788 {
    case 9223372036854775807:
        switch genFn0() {
        case -268:
            var v12 int32 = -1
            _ = v12
        }
    case 1:
        switch 639 {
        case 2147483647:
            var v13 float64 = 0.0
            _ = v13
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = (!v3)
    v14 := GenS5{F0: (v3 && v3), F1: 2147483647, F2: -2147483648}
    var v15 uint64 = 589
    switch "" {
    case "\n\t":
        v1 = v1
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch fi16 {
        case 233:
            var v17 int = (fi16 * (fi16 & (^(fi16))))
            _ = v17
        default:
            var v18 uint8 = 1
            _ = v18
        }
        _ = fi16
    }
    if true {
        switch -869 {
        case 128:
            fmt.Println(v11, v11, v15)
        case 271:
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 string = "hello"
                fmt.Println(v3)
                _ = fi19
                _ = v20
            }
        case 674:
            var v21 uint = 0
            _ = v21
        default:
            fmt.Println(v1)
        }
        var v22 uint8 = uint8(v15)
        var v23 int = 9223372036854775807
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
