package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0() uint64 {
    if true {
        var v7 uint64 = 0
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int8 = -1
            v10 := []int{(fi8 * fi8), 172, (fi8 >> ((834) & 63))}
            v11 := GenS1{F0: 0, F1: fi8}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(v7)
        _ = v7
    } else {
        var v12 bool = false
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(fi13, fi13)
            _ = fi13
        }
        switch 256 {
        case 706:
            for fi14 := 0; fi14 < 12; fi14++ {
                if false {
                    if v12 {
                        v15 := GenS1{F0: 0, F1: 32767}
                        _ = v15
                    } else {
                        v12 = ((4294967295 <= 65535) && v12)
                        for fi16 := 0; fi16 < 10; fi16++ {
                            fmt.Println(v12, fi16, v12)
                            var v17 uint8 = 246
                            var v18 float32 = float32(1.5)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    switch "hello" {
                    case "hello":
                        v12 = false
                    case "\n\t":
                        var v19 uint32 = 794
                        _ = v19
                    case "foo bar":
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v12 = (v12 || (2147483647 != 166))
                            v21 := []int{}
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 int32 = 42
                    _ = v22
                }
                v12 = v12
                _ = fi14
            }
        case 127:
            for fi23 := 0; fi23 < 18; fi23++ {
                switch fi23 {
                case 32767:
                    v24 := []int{(fi23 & 32768), fi23}
                    _ = v24
                default:
                    v25 := GenS1{F0: 772, F1: fi23}
                    _ = v25
                }
                if (!(!("a" > "foo bar"))) {
                    v12 = (!((0.0 + -2.25) == (1.5 + -2.25)))
                    v26 := []int{127, (fi23 - (fi23 ^ 32767)), (fi23 % ((fi23) | 1))}
                    v27 := GenS1{F0: 65535, F1: fi23}
                    _ = v26
                    _ = v27
                }
                _ = fi23
            }
        default:
            v28 := []int{-2147483648}
            _ = v28
        }
        _ = v12
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        if (3 != 2) {
            var v30 string = "a"
            v30 = v30
            _ = v30
        } else {
            if false {
                fmt.Println(fi29)
                v31 := GenS1{F0: 2, F1: fi29}
                var v32 uint16 = 1
                _ = v31
                _ = v32
            }
            var v33 uint16 = 42
            _ = v33
        }
        _ = fi29
    }
    return 42
}
func genFn1(p0 bool, p1 uint16, p2 int64) int16 {
    p2 += p2
    var v34 int64 = 1
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 int8 = -128
        _ = fi35
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 972
}

func main() {
    var v1 uint16 = 32767
    var v2 int16 = -32
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.1
    var v5 int = 127
    var v6 uint = 42
    v1 -= (v1 * (v1 >> (((^(v6))) & 63)))
    switch (v5 << ((4294967295) & 63)) {
    case 32767:
        v2 -= -122
    default:
        v2 -= v2
    }
    var v37 uint = (v6 % (((v6 << (((v6 << (((v6 & 306)) & 63))) & 63))) | 1))
    v38 := v3[828]
    for fi39 := 0; fi39 < 12; fi39++ {
        v40 := GenS1{F0: 255, F1: -1}
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 uint32 = 65535
        var v43 uint8 = 2
        v44 := []int{(v5 * v5), 576, (fi41 - (-(v5))), (^(v5))}
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v46 := GenS2{F0: 175, F1: (v1 % ((v1) | 1)), F2: v2}
        _ = fi45
        _ = v46
    }
    var v47 float32 = (3.14159 * (0.1 / 1.5))
    for fi48 := 0; fi48 < 10; fi48++ {
        fmt.Println(v2)
        if v38 {
            v2 += -32768
        } else {
            v4 = v4
            v49 := GenS1{F0: (v37 + 255), F1: fi48}
            _ = v49
        }
        _ = fi48
    }
    var v50 uint16 = v1
    v51 := []int{}
    fmt.Println(v38)
    for fi52 := 0; fi52 < 15; fi52++ {
        v37 = v6
        _ = fi52
    }
    v53 := GenS1{F0: v6, F1: (v5 * v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v6
}
