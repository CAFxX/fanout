package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }

func main() {
    var v1 int32 = 1
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 int8 = v2
    v5 := make(map[string]string)
    fmt.Println(v4)
    for fi6 := 0; fi6 < 5; fi6++ {
        v4 = v4
        v4 = (v2 - 100)
        _ = fi6
    }
    switch -9223372036854775808 {
    case 1:
        v1 = (v1 << ((511) & 63))
    case 42:
        var v7 int64 = -554
        _ = v7
    default:
        var v8 bool = ((42 != 42) || (!false))
        _ = v8
    }
    var v9 int32 = int32(v2)
    v10 := GenS1{F0: 1, F1: 326, F2: 4294967295}
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 += v9
        fmt.Println(fi11, v1)
        _ = fi11
    }
    switch -173 {
    case 558:
        for fi12 := 0; fi12 < 3; fi12++ {
            switch (fi12 >> ((255) & 63)) {
            case -414:
                var v13 int64 = -9223372036854775808
                _ = v13
            default:
                if ((v4 >> ((18446744073709551615) & 63)) >= -128) {
                    v2 += (-(v2))
                    var v14 uint16 = 1
                    _ = len(v3)
                    _ = v14
                } else {
                    v1 -= v9
                    var v15 int8 = (v4 | v2)
                    _ = v15
                }
            }
            v16 := GenS1{F0: 2, F1: 42, F2: (fi12 ^ 42)}
            v17 := GenS2{F0: 626, F1: 984, F2: 0}
            _ = fi12
            _ = v16
            _ = v17
        }
    case -915:
        v4 = (v4 | (v2 * (-(v4))))
    default:
        v2 = v2
    }
    v18 := GenS1{F0: 0, F1: 934, F2: 0}
    var v19 uint8 = 2
    switch int(v19) {
    case 127:
        if (42 > 486) {
            v20 := []int{2147483647}
            if false {
                v19 -= v19
            } else {
                v19 = v19
                for fi21 := 0; fi21 < 15; fi21++ {
                    v19 = v19
                    var v22 int = fi21
                    _ = fi21
                    _ = v22
                }
            }
            _ = v20
        } else {
            fmt.Println(v1, v9)
        }
    case -2147483648:
        var v23 bool = (2147483648 != 1)
        _ = v23
    case -9223372036854775808:
        var v24 uint16 = 0
        _ = v24
    default:
        var v25 float64 = (0.0 + (3.14159 * (1e10 * -2.25)))
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if false {
            switch fi26 {
            case -800:
                v1 += 2147483647
            case 9223372036854775807:
                fmt.Println(v2, v4)
            default:
                switch (fi26 >> ((42) & 63)) {
                case 0:
                    if true {
                        v27 := []int{(fi26 & (fi26 % (((fi26 >> ((2) & 63))) | 1)))}
                        fmt.Println(v1, v19)
                        _ = v27
                    } else {
                        var v28 int8 = 42
                        _ = v28
                    }
                case 1:
                    delete(v5, "")
                }
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                if false {
                    if false {
                        fmt.Println(v9)
                    } else {
                        v19 += v19
                        var v30 float64 = 1.5
                        v1 += v9
                        _ = v30
                    }
                    v9 = 842
                }
                var v31 float32 = 3.14159
                _ = fi29
                _ = v31
            }
        } else {
            v1 -= (v1 % ((42) | 1))
            var v32 int32 = 0
            _ = v32
        }
        _ = fi26
    }
    v33 := GenS2{F0: 4294967295, F1: 2147483647, F2: (^(v19))}
    v34 := GenS2{F0: 102, F1: 39, F2: (v19 >> ((234) & 63))}
    if false {
        var v35 int = -785
        v1 += int32(v19)
        _ = v35
    } else {
        v36 := GenS1{F0: 0, F1: 4294967295, F2: 829}
        for fi37 := 0; fi37 < 17; fi37++ {
            var v38 float32 = (3.14159 + 1.5)
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    var v39 int16 = 1
    v1 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
