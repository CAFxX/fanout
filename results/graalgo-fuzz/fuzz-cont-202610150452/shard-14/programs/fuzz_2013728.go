package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int8) uint32 {
    var v7 uint8 = 8
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 += (^(v7))
        var v9 int32 = -1
        v10 := GenS1{F0: 182}
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        var v11 int16 = -1
        v12 := []int{179, 728, 128, 128}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = 42
    var v2 bool = false
    v3 := make(map[string]bool)
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]bool)
    var v6 float64 = 0.1
    v13 := []int{-9223372036854775808, int(v4), 32768}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int = (fi14 | fi14)
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 65535, F1: 1}
    v4 -= v4
    if v2 {
        var v17 uint32 = 2147483648
        _ = v17
    }
    fmt.Println(v1)
    v5[2147483647] = v2
    switch 32767 {
    case 1:
        v18 := GenS2{F0: 1, F1: 9223372036854775807}
        _ = v18
    case -1:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v2 = false
                var v21 int8 = (v1 - v1)
                switch (^(fi20)) {
                case 568:
                    v22 := GenS1{F0: 255}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    case -656:
        if (!((128 <= 127) && v2)) {
            for fi23 := 0; fi23 < 2; fi23++ {
                if (!(!(4294967295 <= v4))) {
                    v4 = v4
                    var v24 uint8 = 0
                    _ = v24
                } else {
                    var v25 bool = (-217 > 9223372036854775807)
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v4 += v4
                        _ = fi26
                    }
                    _ = v25
                }
                if len(v13) > 0 {
                    v13[0] = 128
                }
                _ = fi23
            }
        }
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 bool = v2
        switch (fi27 % ((fi27) | 1)) {
        case 4294967295:
            v2 = false
        }
        if v2 {
            var v29 bool = (18446744073709551615 <= 557)
            var v30 int8 = v1
            _ = v29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
