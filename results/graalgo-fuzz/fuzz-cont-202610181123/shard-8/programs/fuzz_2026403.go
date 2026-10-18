package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 int, p1 string) int32 {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int64 = 465
        _ = fi5
        _ = v6
    }
    var v7 float32 = 1.5
    v8 := []int{876, p0, p0, (p0 - (p0 << ((988) & 63)))}
    v7 = (v7 + float32(0.0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int64, p1 bool) uint64 {
    p0 = -1
    if (!p1) {
        var v9 int64 = -668
        switch "a" {
        case "a":
            fmt.Println(p0)
        case "":
            for fi10 := 0; fi10 < 13; fi10++ {
                if (!p1) {
                    var v11 int16 = 42
                    _ = v11
                }
                v12 := []int{fi10}
                _ = fi10
                _ = v12
            }
        case "foo bar":
            var v13 uint64 = 751
            _ = v13
        }
        _ = v9
    } else {
        if (((4294967295 != 927) || true) && (false && (-128 <= 42))) {
            for fi14 := 0; fi14 < 19; fi14++ {
                p1 = p1
                _ = fi14
            }
            v15 := GenS1{F0: 399, F1: 1, F2: 448}
            _ = v15
        } else {
            v16 := []int{42}
            p1 = false
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int64 = -9223372036854775808
            if p1 {
                if (2147483647 < 1) {
                    switch fi17 {
                    case 4294967295:
                        p0 = (^(v18))
                    case 42:
                        p0 = (^(p0))
                    case -1:
                        v19 := GenS2{F0: 42, F1: p1}
                        _ = v19
                    default:
                        p0 = (^(p0))
                    }
                }
            }
            for fi20 := 0; fi20 < 5; fi20++ {
                if ((!((v18 * (v18 >> ((779) & 63))) < v18)) && true) {
                    var v21 uint16 = 972
                    _ = v21
                }
                _ = fi20
            }
            _ = fi17
            _ = v18
        }
    }
    var v22 int8 = 0
    v23 := []int{0, 32768, 2147483647}
    p0 = (p0 % ((-929) | 1))
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 845
}

func main() {
    var v1 string = "a"
    var v2 bool = (((3.14159 / 3.14159) / 3.14159) > ((0.1 * 3.14159) * ((-2.25 / 1e10) * (1e10 / (1.5 * (float64(-2.25) + float64(1.5)))))))
    var v3 int64 = -335
    var v4 uint8 = 1
    for fi24 := 0; fi24 < 5; fi24++ {
        v1 = v1
        v25 := GenS1{F0: v3, F1: 624, F2: 0}
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS1{F0: (-(v3)), F1: 975, F2: 65535}
            v1 = v1
            fmt.Println(v1)
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    v28 := GenS2{F0: 30, F1: ((!(!false)) || false)}
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := GenS2{F0: v4, F1: (true || false)}
            switch "" {
            case "\n\t":
                var v32 bool = false
                _ = v32
            default:
                var v33 uint = 4294967295
                _ = v33
            }
            if (v2 || v2) {
                fmt.Println(v2)
                v34 := []int{}
                _ = v34
            }
            _ = fi30
            _ = v31
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := []int{(^(fi29))}
            _ = fi35
            _ = v36
        }
        if v2 {
            var v37 int8 = -7
            v38 := []int{fi29}
            if v2 {
                fmt.Println(v2)
            } else {
                v2 = v2
                var v39 int8 = v37
                v39 -= (v39 + int8(v3))
                _ = v39
            }
            _ = v37
            _ = v38
        }
        _ = fi29
    }
    v40 := GenS1{F0: 0, F1: 42, F2: 32767}
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 float32 = 0.1
        _ = fi41
        _ = v42
    }
    if true {
        v43 := GenS1{F0: (v3 + (v3 ^ v3)), F1: 255, F2: 931}
        if (813 >= 2) {
            switch 9223372036854775807 {
            case -1:
                fmt.Println(v4, v4, v4)
            default:
                switch 127 {
                case -1:
                    v4 -= (v4 | (v4 + v4))
                default:
                    fmt.Println(v1, v2)
                }
            }
            var v44 int32 = int32(v3)
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 uint64 = 42
                var v47 int16 = -907
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = v44
        }
        var v48 uint8 = v4
        _ = v43
        _ = v48
    }
    fmt.Println(v2, v1)
    v4 += (v4 * 1)
    v49 := []int{2147483647, 1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v49
}
