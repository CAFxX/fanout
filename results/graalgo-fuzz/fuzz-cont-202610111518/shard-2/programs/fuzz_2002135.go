package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0() int16 {
    var v5 int32 = 2147483647
    var v6 bool = false
    var v7 uint32 = 686
    switch 32767 {
    case 879:
        v8 := GenS1{F0: 4294967295}
        _ = v8
    case 2147483647:
        var v9 bool = v6
        _ = v9
    case -111:
        var v10 int8 = 1
        _ = v10
    default:
        v7 = v7
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]bool)
    var v3 uint16 = 509
    var v4 int8 = -91
    var v11 int16 = genFn0()
    for fi12 := 0; fi12 < 6; fi12++ {
        switch 128 {
        case 2147483647:
            if (!(int64(v4) >= 42)) {
                switch (fi12 * (fi12 + (-(fi12)))) {
                case -157:
                    v11 = 990
                default:
                    var v13 uint32 = (v1 * (^(v1)))
                    _ = v13
                }
                v14 := []int{(fi12 & -193), fi12, fi12}
                if true {
                    fmt.Println(v11, v11, v4)
                    var v15 int64 = 328
                    if true {
                        if true {
                            fmt.Println(v4)
                        }
                    }
                    _ = v15
                } else {
                    var v16 uint16 = 65535
                    _ = v16
                }
                _ = v14
            }
        case -537:
            for fi17 := 0; fi17 < 2; fi17++ {
                if (4294967295 > 18446744073709551615) {
                    var v18 uint16 = 32767
                    fmt.Println(fi17)
                    v2[9223372036854775807] = false
                    _ = v18
                }
                _ = fi17
            }
        }
        _ = fi12
    }
    v11 = 42
    switch -123 {
    case 42:
        v19 := GenS1{F0: -46}
        _ = v19
    default:
        fmt.Println(v3, v3, v3)
    }
    switch 2147483647 {
    case -2147483648:
        if (!true) {
            v20 := GenS2{F0: -1}
            v21 := []int{-2147483648, -158, 9223372036854775807}
            v22 := GenS1{F0: -2147483648}
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v23 := []int{652}
    if false {
        var v24 uint16 = 32767
        if (1 == 206) {
            v23 = append(v23, -1)
        }
        _ = len(v2)
        _ = v24
    } else {
        v11 -= v11
        var v25 int32 = 1
        _ = v25
    }
    v26 := GenS2{F0: -2147483648}
    if (!((42 > 1) || true)) {
        if ((v4 << ((255) & 63)) >= -128) {
            v27 := GenS3{F0: 288, F1: (v1 * 42), F2: -1}
            fmt.Println(v11)
            fmt.Println(v3, v11, v1)
            _ = v27
        } else {
            v4 -= (^(v4))
        }
        var v28 string = "\n\t"
        for fi29 := 0; fi29 < 6; fi29++ {
            fmt.Println(v4, v3)
            _ = fi29
        }
        _ = v28
    }
    switch -714 {
    case 871:
        v30 := GenS3{F0: 65535, F1: (v1 + 4294967295), F2: (^(v4))}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
