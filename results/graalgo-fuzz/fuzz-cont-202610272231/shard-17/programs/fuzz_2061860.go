package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8) int16 {
    var v4 float64 = 1.5
    v5 := GenS1{F0: -520}
    fmt.Println(v4, v4)
    v6 := []int{747, -2147483648, 280}
    if true {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                if (42 != 0) {
                    v4 = (3.14159 - v4)
                    switch fi7 {
                    case 175:
                        v4 = v4
                    }
                }
                _ = fi8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                v4 += (-2.25 / 1e10)
                if len(v6) > 0 {
                    v6[0] = fi9
                }
                _ = fi9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                _ = len(v6)
                _ = fi10
            }
            _ = fi7
        }
    } else {
        v4 += (v4 - (v4 * 1.5))
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 string, p1 string, p2 uint64) uint {
    fmt.Println(p1, p0, p2)
    v11 := GenS2{F0: -401}
    var v12 uint = 42
    if false {
        v13 := GenS3{F0: 0, F1: ("" + "a"), F2: ""}
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS3{F0: -9223372036854775808, F1: (p1 + p1), F2: (p1 + ("foo bar" + (p0 + "foo bar")))}
            switch -9223372036854775808 {
            case -840:
                v16 := GenS4{F0: 0, F1: (p0 + (p0 + ("" + p0)))}
                _ = v16
            case 127:
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS5{F0: 1000}
                    p0 = p1
                    var v19 uint = (v12 | (v12 ^ 18446744073709551615))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case -2147483648:
                var v20 float64 = 1e10
                _ = v20
            default:
                var v21 uint64 = p2
                _ = v21
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := []int{fi22}
            p1 = p1
            v12 = (v12 / ((992) | 1))
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := []int{(^(fi24)), (fi24 << ((v12) & 63)), fi24, fi24}
            if ((^(fi24)) >= (-(fi24))) {
                p2 += 0
                var v26 int16 = 1000
                _ = v26
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    if false {
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    }
                    _ = fi27
                }
                v28 := GenS3{F0: 365, F1: p1, F2: "\n\t"}
                _ = v28
            }
            fmt.Println(p1, v12, v12)
            _ = fi24
            _ = v25
        }
        p0 = "\n\t"
    }
    v29 := []int{32768, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v29
    return (v12 ^ (v12 * (v12 & 18446744073709551615)))
}

func main() {
    var v1 string = ("\n\t" + "a")
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v30 uint32 = 2
    var v31 int = -1
    var v32 int16 = -1
    v3[(v30 + (v30 % ((42) | 1)))] = false
    v33 := GenS6{F0: 42}
    v34 := v3[(v30 * (v30 | v30))]
    v32 = (-(v32))
    var v35 uint32 = (^(v30))
    if v34 {
        for fi36 := 0; fi36 < 7; fi36++ {
            fmt.Println(v30)
            v2[(^(v30))] = v31
            _ = fi36
        }
        v37 := []int{-999, v31, v31, v31}
        var v38 int16 = (v32 | 32767)
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
