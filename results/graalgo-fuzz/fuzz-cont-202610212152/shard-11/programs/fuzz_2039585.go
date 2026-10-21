package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 uint) uint8 {
    v4 := GenS1{F0: 0, F1: 128}
    v5 := []int{32767, 829}
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 12; fi7++ {
        p1 = 1
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if false {
        var v9 int8 = 127
        var v10 bool = (42 >= 2147483647)
        _ = v10
        _ = v9
    } else {
        if ((0.0 == (1.5 + (((0.1 * (1.5 / (1e10 * ((-2.25 / 0.1) / 3.14159)))) - ((0.0 / ((float64(-2.25) / -2.25) / 1.0)) - float64(-2.25))) * 1.5))) && false) {
            fmt.Println(v6)
            var v11 int = 127
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 bool, p1 int16, p2 int) int32 {
    switch (p2 & (p2 & -1)) {
    case 81:
        p1 -= p1
    case 9223372036854775807:
        for fi12 := 0; fi12 < 2; fi12++ {
            switch p2 {
            case 2147483647:
                var v13 uint16 = 1
                _ = v13
            default:
                if (p0 || false) {
                    fmt.Println(fi12, fi12, p0)
                } else {
                    var v14 int8 = -103
                    _ = v14
                }
            }
            p1 -= (-(p1))
            var v15 int32 = 42
            _ = fi12
            _ = v15
        }
    case 685:
        p2 = -1
    }
    fmt.Println(p0, p2)
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return 231
}

func main() {
    var v1 bool = (1 < 65535)
    v2 := make(map[uint32]int)
    var v3 string = ""
    for fi16 := 0; fi16 < 8; fi16++ {
        if v1 {
            v17 := GenS2{F0: (fi16 ^ (fi16 + fi16)), F1: int64(fi16), F2: 0}
            switch fi16 {
            case -67:
                switch 32767 {
                case 32768:
                    var v18 uint = 0
                    _ = v18
                case 127:
                    fmt.Println(fi16)
                default:
                    v19 := GenS1{F0: 0, F1: fi16}
                    _ = v19
                }
            case 2147483647:
                var v20 uint64 = 4294967295
                _ = v20
            case 984:
                var v21 bool = (18446744073709551615 < 218)
                _ = v21
            default:
                var v22 uint8 = 1
                _ = v22
            }
            _ = v17
        }
        v23 := GenS3{F0: 42, F1: 686}
        _ = fi16
        _ = v23
    }
    var v24 bool = false
    v25 := GenS4{F0: -128}
    var v26 float32 = float32(1.5)
    var v27 int32 = 0
    var v28 int32 = genFn1(false, 1000, -9223372036854775808)
    var v29 uint64 = 113
    fmt.Println(v1, v1, v1)
    v27 += v28
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS4{F0: 41}
        _ = len(v2)
        _ = fi30
        _ = v31
    }
    switch -782 {
    case 584:
        v32 := v2[0]
        _ = v32
    case 2147483647:
        if v24 {
            v33 := GenS4{F0: 95}
            v2[2147483648] = 32768
            _ = v33
        } else {
            var v34 bool = (int64(v27) >= -799)
            _ = v34
        }
    case 4294967295:
        if (!((((float32(0.1) * v26) * v26) == 1.5) && (3.14159 != 0.0))) {
            v35 := []int{}
            for fi36 := 0; fi36 < 10; fi36++ {
                for fi37 := 0; fi37 < 3; fi37++ {
                    v3 = "a"
                    _ = len(v35)
                    _ = fi37
                }
                _ = fi36
            }
            var v38 uint32 = 1
            _ = v35
            _ = v38
        }
    }
    var v39 int16 = 32767
    fmt.Println(v29)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
}
