package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
func genFn0(p0 uint8, p1 uint16) uint32 {
    p0 = (p0 - 138)
    fmt.Println(p0, p0)
    var v7 uint8 = (p0 - (p0 ^ (p0 >> ((42) & 63))))
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = fi8
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return uint32(p0)
}
func genFn1(p0 int8) uint64 {
    fmt.Println(p0)
    if false {
        for fi10 := 0; fi10 < 4; fi10++ {
            p0 += (p0 / (((p0 + 100)) | 1))
            _ = fi10
        }
        v11 := GenS1{F0: 2147483648, F1: -9223372036854775808}
        _ = v11
    } else {
        var v12 int = 42
        p0 += p0
        _ = v12
    }
    v13 := []int{32767, -9223372036854775808, 2147483647}
    p0 -= 127
    _ = p0
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 1
    v3 := []uint32{1, 42}
    var v4 uint32 = 2147483648
    var v5 int16 = -357
    var v6 uint64 = 994
    v14 := v1[v4]
    if false {
        var v15 int32 = 0
        var v16 bool = ("" <= "foo bar")
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint8 = 2
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 uint64 = (^(v6))
            v4 += v4
            fmt.Println(v2, fi19, v20)
            var v21 int16 = 32767
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    v22 := GenS2{F0: (v2 | v2), F1: v6, F2: -13}
    v23 := GenS1{F0: v14, F1: 188}
    for fi24 := 0; fi24 < 10; fi24++ {
        switch fi24 {
        case 2147483647:
            v25 := GenS3{F0: v2, F1: "\n\t", F2: v2}
            _ = v25
        case -1:
            v26 := GenS1{F0: v4, F1: 0}
            _ = v26
        }
        _ = fi24
    }
    delete(v1, (v4 << ((2) & 63)))
    if (v5 != (v5 + v5)) {
        if false {
            for fi27 := 0; fi27 < 19; fi27++ {
                v14 = v14
                var v28 int64 = -9223372036854775808
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: (v14 % ((1) | 1)), F1: -796}
            _ = v29
        } else {
            switch -527 {
            case 128:
                if true {
                    v3 = append(v3, genFn0(255, 0))
                }
            }
        }
        v5 = -267
        switch 42 {
        case 42:
            v2 -= (v2 * v2)
        }
    } else {
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 19; fi31++ {
                v5 -= -515
                v2 = (-(v2))
                _ = fi31
            }
            v32 := GenS1{F0: v4, F1: 42}
            _ = fi30
            _ = v32
        }
    }
    fmt.Println(v2)
    switch 2147483647 {
    case 0:
        v33 := GenS1{F0: (v4 + v4), F1: 591}
        _ = v33
    case 161:
        var v34 float64 = (3.14159 + float64(-2.25))
        _ = v34
    case -29:
        var v35 uint8 = 42
        _ = v35
    default:
        if false {
            var v36 uint = 12
            var v37 int64 = -1
            _ = v36
            _ = v37
        }
    }
    if true {
        var v38 uint16 = uint16(v6)
        fmt.Println(v4)
        _ = v38
    } else {
        if (int(v14) <= -2147483648) {
            fmt.Println(v4, v5, v5)
        }
    }
    v4 += (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
