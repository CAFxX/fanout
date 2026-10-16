package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 int8, p1 int64, p2 uint) int {
    var v6 uint64 = 4294967295
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            if (int64(fi7) <= p1) {
                if false {
                    v6 += 4294967295
                } else {
                    var v8 uint16 = 2
                    _ = v8
                }
                if true {
                    v9 := GenS1{F0: 42, F1: (^(p0)), F2: 18446744073709551615}
                    _ = v9
                } else {
                    v10 := GenS1{F0: -697, F1: -95, F2: (^(p2))}
                    _ = v10
                }
            } else {
                for fi11 := 0; fi11 < 14; fi11++ {
                    p2 = (p2 >> ((65535) & 63))
                    _ = fi11
                }
                v12 := GenS2{F0: (^(p1)), F1: (true && false), F2: (fi7 & fi7)}
                var v13 int = 32768
                _ = v12
                _ = v13
            }
        } else {
            if false {
                p1 = (p1 % ((p1) | 1))
                if (uint16(p2) == 0) {
                    p2 -= 4294967295
                    v6 = (v6 - 453)
                    v14 := []int{2147483647, fi7, (fi7 << ((p2) & 63))}
                    _ = v14
                }
                if true {
                    v15 := GenS1{F0: -339, F1: 42, F2: (p2 >> (((^(p2))) & 63))}
                    var v16 int32 = 2147483647
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{fi7, (fi7 ^ (fi7 * fi7)), fi7, -617}
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 int = (-(fi7))
                        var v20 float64 = ((1e10 / 1e10) - 0.1)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v17
                }
            }
            v21 := GenS3{F0: p2, F1: (("\n\t" + "a") + "a")}
            _ = v21
        }
        if (!false) {
            p0 = (p0 % ((-1) | 1))
            v22 := GenS1{F0: -2147483648, F1: p0, F2: (p2 ^ (p2 >> ((p2) & 63)))}
            _ = v22
        } else {
            v23 := GenS2{F0: (p1 ^ p1), F1: false, F2: (^(fi7))}
            var v24 int8 = (p0 / (((p0 - (^(p0)))) | 1))
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 uint = (p2 - p2)
                fmt.Println(p0, v6, v6)
                var v27 uint = (p2 * (p2 & p2))
                var v28 string = "hello"
                v29 := []int{fi25, fi25}
                v28 = v28
                var v30 uint16 = 960
                fmt.Println(v28, v30, v24)
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = v23
            _ = v24
        }
        _ = fi7
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        p0 = (-(p0))
        v32 := []int{-735, fi31, (^(fi31)), 1}
        _ = fi31
        _ = v32
    }
    if true {
        v33 := []int{-9223372036854775808, -627, 746, -251}
        for fi34 := 0; fi34 < 12; fi34++ {
            if (((((-2.25 - -2.25) / float64(-2.25)) > (1e10 / (3.14159 * 1.5))) || (false && ((1e10 - 0.1) != 3.14159))) && false) {
                p2 = p2
                v35 := GenS4{F0: fi34, F1: ((p2 > p2) && (v6 == v6))}
                v6 = (v6 & (v6 & 42))
                _ = v35
            }
            var v36 float32 = (1.5 * (1.5 - (0.0 + 1.5)))
            var v37 int8 = p0
            _ = fi34
            _ = v36
            _ = v37
        }
        v33 = append(v33, -2147483648)
        _ = v33
    } else {
        var v38 uint64 = v6
        if (int16(p0) >= -1) {
            v38 = (v38 / ((v6) | 1))
        }
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 1; fi40++ {
                if true {
                    v38 = 1
                    v41 := GenS2{F0: (^(p1)), F1: true, F2: fi39}
                    v38 -= (v6 / ((721) | 1))
                    _ = v41
                }
                _ = fi40
            }
            v42 := GenS3{F0: uint(v38), F1: "\n\t"}
            _ = fi39
            _ = v42
        }
        _ = v38
    }
    p0 -= 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -661
}
func genFn1(p0 bool) uint32 {
    fmt.Println(p0)
    var v43 uint32 = 42
    _ = p0
    _ = v43
    return v43
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{1, 0, 2147483647, 0}
    v3 := []int{-2147483648, -45, 4294967295}
    var v4 int32 = -1
    var v5 int8 = 1
    fmt.Println(v1)
    var v44 bool = true
    v4 = (-(v4))
    v4 = v4
    fmt.Println(v1)
    v45 := GenS2{F0: -9223372036854775808, F1: (!(v4 <= (v4 | v4))), F2: 2147483647}
    v2 = append(v2, 1)
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 18; fi47++ {
            if len(v3) > 0 {
                v3[0] = (-(fi46))
            }
            _ = len(v2)
            v4 = -757
            _ = fi47
        }
        if len(v2) > 0 {
            v2[0] = 418
        }
        if len(v2) > 0 {
            v2[0] = genFn1(((!(((("" + "") < "hello") && v44) && v44)) && v44))
        }
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
