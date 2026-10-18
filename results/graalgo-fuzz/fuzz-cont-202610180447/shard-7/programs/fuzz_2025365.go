package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 string; F1 int16; F2 uint32 }
type GenS5 struct { F0 int32; F1 int32; F2 uint }
func genFn0(p0 int64, p1 uint, p2 uint32) uint32 {
    p2 += (^(p2))
    p1 -= (p1 + p1)
    fmt.Println(p1, p0, p1)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 362
}
func genFn1(p0 int32, p1 uint32, p2 int64) uint64 {
    switch "" {
    case "\n\t":
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint32 = 4294967295
            _ = fi6
            _ = v7
        }
    case "foo bar":
        switch -996 {
        case 32768:
            v8 := GenS1{F0: (p2 << ((4294967295) & 63)), F1: 1}
            _ = v8
        case -59:
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint = 255
                _ = fi9
                _ = v10
            }
        case -942:
            v11 := []int{32767, 1, 2147483647, 32767}
            _ = v11
        }
    }
    if (true && ((495 == 42) || false)) {
        p0 = -2147483648
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (-(fi12)) {
            case 136:
                if (genFn0(p2, 4294967295, p1) != 675) {
                    var v13 int16 = -574
                    v14 := []int{fi12}
                    _ = v13
                    _ = v14
                } else {
                    var v15 uint64 = 925
                    v16 := GenS1{F0: (-(p2)), F1: v15}
                    _ = v15
                    _ = v16
                }
            default:
                p1 += (p1 * p1)
            }
            if (-756 >= 249) {
                fmt.Println(p2, fi12)
                for fi17 := 0; fi17 < 2; fi17++ {
                    if true {
                        if true {
                            fmt.Println(fi12)
                        }
                        v18 := GenS1{F0: p2, F1: 4294967295}
                        _ = v18
                    }
                    _ = fi17
                }
                for fi19 := 0; fi19 < 3; fi19++ {
                    switch "" {
                    case "a":
                        v20 := GenS1{F0: p2, F1: 626}
                        _ = v20
                    case "\n\t":
                        v21 := []int{fi12, fi19, fi12, 32768}
                        _ = v21
                    case "hello":
                        v22 := []int{-1, (fi12 & fi19), 127}
                        _ = v22
                    default:
                        var v23 float32 = 0.0
                        _ = v23
                    }
                    _ = fi19
                }
            }
            _ = fi12
        }
    } else {
        p0 = (p0 + p0)
        var v24 uint = 1
        v25 := []int{-368, 1, 2147483647, 42}
        _ = v24
        _ = v25
    }
    var v26 float64 = (float64(float32((0.0 / (3.14159 * (3.14159 - 1e10))))) / 1e10)
    var v27 uint = 1
    v28 := GenS2{F0: -366}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v28
    return 42
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 1
    v3 := []int{2147483647, 128, 2147483647}
    var v4 int64 = 552
    var v5 int = -9223372036854775808
    fmt.Println(v5)
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := GenS2{F0: (^(v5))}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v2 += 401
    for fi32 := 0; fi32 < 14; fi32++ {
        if true {
            v33 := GenS2{F0: (v5 % ((9223372036854775807) | 1))}
            _ = v33
        } else {
            var v34 uint8 = 127
            switch fi32 {
            case 32767:
                var v35 uint64 = genFn1(0, 2, (-(v4)))
                _ = v35
            case -404:
                v36 := GenS3{F0: 273, F1: 127}
                _ = v36
            }
            _ = v34
        }
        _ = len(v3)
        _ = fi32
    }
    v37 := GenS4{F0: "a", F1: 32767, F2: 2147483647}
    v5 = (v5 << ((42) & 63))
    if false {
        var v38 float64 = 0.0
        _ = v38
    } else {
        var v39 int8 = 1
        _ = v39
    }
    var v40 string = "\n\t"
    for fi41 := 0; fi41 < 20; fi41++ {
        v3 = append(v3, (v5 / ((fi41) | 1)))
        _ = fi41
    }
    v5 -= (v5 >> (((^(v1))) & 63))
    fmt.Println(v5, v4)
    fmt.Println(v40, v4, v4)
    if (18446744073709551615 <= 4294967295) {
        if true {
            v42 := GenS5{F0: 2147483647, F1: 42, F2: uint(v2)}
            var v43 uint64 = uint64(v1)
            _ = v42
            _ = v43
        } else {
            v40 = (v40 + (((v40 + ((v40 + ((v40 + v40) + "hello")) + ("hello" + ("hello" + v40)))) + "\n\t") + ((("foo bar" + v40) + v40) + (v40 + ("" + ("foo bar" + ("a" + v40)))))))
        }
        v44 := GenS4{F0: (v40 + v40), F1: 42, F2: 1}
        _ = v44
    }
    v45 := GenS3{F0: (v1 | v1), F1: 0}
    for fi46 := 0; fi46 < 1; fi46++ {
        var v47 uint16 = v2
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 uint8 = 128
            _ = fi48
            _ = v49
        }
        v1 -= v1
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v45
    _ = v5
}
