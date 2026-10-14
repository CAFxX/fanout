package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0(p0 int, p1 uint8, p2 int16) uint32 {
    switch "\n\t" {
    case "foo bar":
        for fi7 := 0; fi7 < 15; fi7++ {
            p2 += (^(p2))
            var v8 uint8 = p1
            v9 := GenS1{F0: (v8 ^ (p1 * 1))}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "a":
        fmt.Println(p1)
    case "":
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint64 = 4294967295
            v12 := []int{(fi10 - (p0 + p0)), 32768}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if true {
        if (-51 <= 0) {
            for fi13 := 0; fi13 < 1; fi13++ {
                p1 = (p1 ^ 128)
                _ = fi13
            }
            if true {
                var v14 uint64 = 4294967295
                p1 = p1
                var v15 string = ""
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                if true {
                    if (((-2.25 - (1.5 / 1e10)) - 1e10) < 1e10) {
                        p2 = p2
                        p1 = p1
                    }
                    v17 := GenS1{F0: 127}
                    _ = v17
                } else {
                    p2 -= p2
                    v18 := GenS1{F0: p1}
                    _ = v18
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    p1 -= (p1 + (^(p1)))
                    p2 -= -222
                    _ = fi19
                }
                switch fi16 {
                case 32768:
                    if false {
                        p1 += p1
                        v20 := GenS1{F0: p1}
                        v21 := []int{}
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi16
            }
        }
        switch -1 {
        case 32768:
            p2 = (p2 | p2)
        case 851:
            v22 := []int{(^(p0))}
            _ = v22
        }
        p2 += (p2 + (p2 * (p2 & p2)))
    }
    v23 := []int{}
    var v24 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 1
    v3 := make(map[int]uint32)
    var v4 float64 = float64((1e10 * 1e10))
    var v5 bool = (18446744073709551615 < 42)
    var v6 bool = (!true)
    v25 := GenS1{F0: 227}
    var v26 uint64 = 320
    v26 -= v26
    var v27 int8 = 1
    var v28 int64 = -9223372036854775808
    for fi29 := 0; fi29 < 13; fi29++ {
        switch "" {
        case "a":
            v28 = 0
        case "":
            if true {
                fmt.Println(v6, v27, v4)
                fmt.Println(v27)
                v1[2147483647] = v5
            }
        case "\n\t":
            v4 -= -2.25
        default:
            v28 = v28
        }
        v30 := GenS2{F0: ((!v5) || (42 == 1)), F1: 2147483647}
        v31 := []int{fi29, -510, (^(fi29)), (-(fi29))}
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS1{F0: 1}
    v28 -= -9223372036854775808
    _ = len(v1)
    v6 = v6
    if (!v6) {
        if (255 < 18446744073709551615) {
            v33 := []int{548, -1}
            _ = v33
        } else {
            fmt.Println(v28)
        }
    } else {
        v34 := GenS3{F0: 2, F1: v28, F2: 302}
        v35 := GenS3{F0: 157, F1: (v28 * v28), F2: -2147483648}
        _ = v34
        _ = v35
    }
    if true {
        for fi36 := 0; fi36 < 10; fi36++ {
            v37 := GenS2{F0: (((3.14159 * v4) == (v4 + -2.25)) || (false || (v6 || (100 <= int8(v2))))), F1: -209}
            _ = fi36
            _ = v37
        }
    } else {
        v5 = (1.5 >= -2.25)
        if (!(42 <= 2147483647)) {
            if v6 {
                v38 := GenS4{F0: v26, F1: 212, F2: 2}
                v26 += v26
                v6 = (((((1e10 + 0.0) + 1e10) + 0.0) > 0.1) || false)
                _ = v38
            }
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 uint = 65535
                v41 := GenS4{F0: v26, F1: 206, F2: (v40 + (^(v40)))}
                _ = fi39
                _ = v40
                _ = v41
            }
        }
    }
    switch -170 {
    case 1:
        v42 := []int{128, 2147483647, -857}
        _ = v42
    case -50:
        v2 = (v2 >> ((42) & 63))
    default:
        v43 := GenS5{F0: 1000, F1: 0}
        _ = v43
    }
    var v44 uint32 = 503
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
