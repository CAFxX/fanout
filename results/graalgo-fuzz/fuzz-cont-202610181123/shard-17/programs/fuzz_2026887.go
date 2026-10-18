package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32) uint64 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 65}
            p0 = 2147483648
            var v9 int64 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        if (0 >= 42) {
            p0 = 681
        } else {
            v10 := []int{(-(fi6))}
            for fi11 := 0; fi11 < 10; fi11++ {
                p0 = 0
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi11
            }
            p0 -= (p0 / ((p0) | 1))
            _ = v10
        }
        if (65535 > 0) {
            if (!true) {
                var v12 float32 = (1.5 * float32((0.0 - 1e10)))
                if (-128 >= -72) {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        p0 = p0
                        p0 -= 2147483648
                        var v14 uint8 = 239
                        _ = fi13
                        _ = v14
                    }
                    v12 += (v12 + (float32(-2.25) + (v12 - -2.25)))
                }
                _ = v12
            }
            var v15 float32 = 1.5
            p0 = p0
            _ = v15
        } else {
            var v16 uint = 2
            _ = v16
        }
        _ = fi6
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        if true {
            if (!true) {
                if true {
                    p0 = 1
                }
                var v18 int64 = -855
                _ = v18
            } else {
                p0 = (p0 * 206)
            }
            v19 := []int{-623, fi17, fi17}
            _ = v19
        }
        p0 = p0
        _ = fi17
    }
    var v20 uint32 = (p0 % (((p0 << ((0) & 63))) | 1))
    _ = p0
    _ = v20
    return 639
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 1
    var v3 int8 = (v2 >> ((4294967295) & 63))
    v21 := []int{127}
    v2 += 91
    v22 := GenS1{F0: 1}
    if (false || (0 == 345)) {
        var v23 uint8 = 129
        _ = v23
    }
    if len(v21) > 0 {
        v21[0] = 32767
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if (-128 == v3) {
            var v25 uint16 = 718
            _ = v25
        }
        v1 = (v1 & 42)
        v2 = -1
        _ = fi24
    }
    var v26 int = 127
    v2 += v3
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int32 = 2147483647
        for fi29 := 0; fi29 < 13; fi29++ {
            v1 = (v1 ^ (v1 + v1))
            v26 -= -2147483648
            var v30 uint16 = 845
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    if true {
        var v31 bool = true
        for fi32 := 0; fi32 < 20; fi32++ {
            switch "a" {
            case "":
                for fi33 := 0; fi33 < 19; fi33++ {
                    if false {
                        switch (fi32 << ((18446744073709551615) & 63)) {
                        case 2147483647:
                            if v31 {
                                v3 = (-(v2))
                                var v34 uint64 = 1
                                var v35 int16 = -1
                                _ = v34
                                _ = v35
                            } else {
                                v1 += (v1 ^ (v1 ^ 282))
                            }
                        case 32767:
                            fmt.Println(v2)
                        }
                    } else {
                        v36 := GenS2{F0: 42, F1: (v3 << ((4294967295) & 63))}
                        if len(v21) > 0 {
                            v21[0] = (fi32 + -172)
                        }
                        _ = v36
                    }
                    _ = fi33
                }
            case "\n\t":
                v26 += (fi32 + (v26 * 42))
            default:
                var v37 uint = 0
                _ = v37
            }
            var v38 int = v26
            _ = fi32
            _ = v38
        }
        _ = v31
    }
    v2 = (v3 ^ v3)
    v39 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v39
}
