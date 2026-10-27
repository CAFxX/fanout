package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 uint8) uint8 {
    if false {
        var v5 uint32 = 0
        v6 := []int{}
        _ = v5
        _ = v6
    }
    p1 = p1
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 562, F1: 0, F2: (p1 & (p1 * -2147483648))}
            v10 := GenS2{F0: fi8, F1: 4294967295, F2: -1}
            switch (fi7 / ((127) | 1)) {
            case 128:
                v11 := []int{(-(fi8)), fi7, p0, fi7}
                _ = v11
            case 4294967295:
                if true {
                    v12 := []int{127, (fi7 ^ 680), (^(p0)), -9223372036854775808}
                    v13 := GenS1{F0: 77, F1: 2, F2: (p1 + 0)}
                    _ = v12
                    _ = v13
                } else {
                    v14 := []int{(fi7 ^ fi7)}
                    _ = v14
                }
            case 1:
                p0 = fi8
            default:
                var v15 int16 = -966
                _ = v15
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v16 := []int{p0, p0, 42}
        _ = fi7
        _ = v16
    }
    v17 := []int{p0, p0, 32767, p0}
    var v18 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return (^(p2))
}

func main() {
    var v1 string = "a"
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 697
    var v4 bool = ((0.0 + 0.1) <= 0.1)
    switch 325 {
    case 552:
        var v19 float32 = -2.25
        _ = v19
    case 1:
        v2 -= v2
    case 9223372036854775807:
        if v4 {
            if true {
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 bool = (((0.1 * 1e10) <= 0.1) && (42 == 0))
                    v22 := []int{(fi20 >> ((18446744073709551615) & 63))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                v2 = -424
            }
        }
    }
    var v23 string = ("" + v1)
    var v24 bool = v4
    if false {
        v23 = (v23 + v1)
        v23 = v1
        v25 := []int{}
        _ = v25
    }
    v23 = "hello"
    var v26 uint16 = 846
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := []int{(fi27 + (^(fi27))), fi27, (fi27 % (((-(fi27))) | 1)), 42}
        v29 := GenS3{F0: 1}
        var v30 int = 354
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    v31 := GenS3{F0: uint64(v26)}
    v32 := GenS4{F0: -746}
    v33 := []int{375}
    var v34 string = ("\n\t" + v1)
    v35 := GenS2{F0: -9223372036854775808, F1: 42, F2: 100}
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            fmt.Println(v3, v34)
            var v38 bool = v24
            _ = fi37
            _ = v38
        }
        v4 = true
        _ = fi36
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 int32 = -296
        switch 9223372036854775807 {
        case -908:
            _ = len(v33)
        case -2147483648:
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        case 42:
            v34 = v23
        }
        var v41 uint8 = genFn0(int(v40), v40, 128)
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
