package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int32; F2 int32 }

func main() {
    var v1 int16 = 1
    v2 := []uint32{912, 739}
    var v3 uint = 718
    var v4 int = -2147483648
    v5 := make(map[string]int)
    v6 := []string{}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint64 = 369
        v4 = (^(fi7))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: v1, F2: 4294967295}
    var v10 int = -2147483648
    switch (v4 % (((^(v4))) | 1)) {
    case 2147483647:
        v3 -= 814
    case -2147483648:
        switch v4 {
        case 0:
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(v4, v10)
                _ = len(v5)
                for fi12 := 0; fi12 < 7; fi12++ {
                    var v13 int32 = 768
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        case -851:
            if false {
                v3 += 18446744073709551615
                var v14 float64 = 1e10
                _ = v14
            }
        case 443:
            var v15 uint64 = 4294967295
            _ = v15
        default:
            switch v4 {
            case 127:
                v10 = 2147483647
            case -9223372036854775808:
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (!false) {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            if ((v3 >= 1) || (42 >= -408)) {
                                fmt.Println(v4, v4)
                            } else {
                                var v18 int16 = v1
                                v4 = (fi16 << ((v3) & 63))
                                _ = v18
                            }
                            _ = fi17
                        }
                    }
                    _ = fi16
                }
            }
        }
    case -637:
        fmt.Println(v10, v1, v10)
    }
    if ((1 > 164) && (v3 < v3)) {
        v19 := GenS1{F0: 968, F1: v1, F2: 163}
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int8 = 0
            _ = fi20
            _ = v21
        }
        var v22 uint = v3
        _ = v19
        _ = v22
    }
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS1{F0: 0, F1: v1, F2: 909}
        var v26 string = "\n\t"
        var v27 bool = true
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if ((v3 & (v3 & v3)) >= (v3 % ((v3) | 1))) {
        v28 := GenS1{F0: 716, F1: (v1 * 32767), F2: 0}
        _ = v28
    }
    fmt.Println(v23, v23)
    v10 = (v4 >> (((v3 * (v3 - (v3 >> (((v3 ^ v3)) & 63))))) & 63))
    delete(v5, "\n\t")
    if (v4 <= (v4 * (v10 | (^(v10))))) {
        v29 := GenS2{F0: 42}
        _ = v29
    }
    v10 = v4
    for fi30 := 0; fi30 < 2; fi30++ {
        fmt.Println(v1)
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 bool = true
            var v33 int64 = -9223372036854775808
            _ = fi31
            _ = v32
            _ = v33
        }
        v34 := GenS3{F0: (fi30 | (-(v10))), F1: -283, F2: 544}
        _ = fi30
        _ = v34
    }
    if ((2147483647 > 1) && false) {
        v35 := GenS2{F0: 127}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
