package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
type GenS3 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 int16, p1 uint8, p2 uint8) uint64 {
    switch -866 {
    case 32767:
        v5 := []int{127, -308, int(p0), 32768}
        _ = v5
    case 128:
        v6 := GenS1{F0: p1, F1: (p0 % ((p0) | 1)), F2: 0}
        _ = v6
    default:
        p0 = (-(p0))
    }
    var v7 uint32 = 65535
    v8 := GenS1{F0: (^(p2)), F1: p0, F2: 127}
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{fi9, -9223372036854775808, (fi9 << ((18446744073709551615) & 63)), fi9}
        _ = fi9
        _ = v10
    }
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 616
}

func main() {
    var v1 uint8 = 94
    var v2 bool = false
    var v3 bool = (18446744073709551615 > 340)
    var v4 float64 = float64(3.14159)
    v11 := []int{}
    v3 = false
    for fi12 := 0; fi12 < 10; fi12++ {
        switch "hello" {
        case "hello":
            switch fi12 {
            case 2147483647:
                var v13 int32 = 0
                _ = v13
            case -532:
                if true {
                    fmt.Println(v2, v1)
                    v14 := GenS2{F0: 42, F1: 65535, F2: -1}
                    _ = v14
                } else {
                    v15 := GenS2{F0: 65535, F1: 2, F2: fi12}
                    if len(v11) > 0 {
                        v11[0] = (fi12 + (fi12 % ((fi12) | 1)))
                    }
                    v16 := GenS1{F0: v1, F1: 32767, F2: 42}
                    _ = v15
                    _ = v16
                }
            }
        }
        _ = fi12
    }
    if (!true) {
        v4 -= float64((1e10 - -2.25))
        v4 -= v4
    }
    fmt.Println(v3, v4)
    if len(v11) > 0 {
        v11[0] = 314
    }
    var v17 int32 = 2147483647
    v17 = -2147483648
    var v18 int8 = 1
    for fi19 := 0; fi19 < 10; fi19++ {
        if (!v2) {
            v4 = -2.25
        }
        _ = fi19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS3{F0: v17, F1: (-(fi20)), F2: (v18 * v18)}
        v1 = v1
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: -1, F1: -747, F2: v18}
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        if false {
            var v24 int64 = -9223372036854775808
            fmt.Println(v1, v2)
            _ = v24
        }
        _ = fi23
    }
    v25 := GenS3{F0: 1, F1: -2147483648, F2: v18}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
