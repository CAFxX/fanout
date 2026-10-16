package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint16, p1 uint64, p2 uint64) bool {
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 9; fi7++ {
        switch 878 {
        case 127:
            v8 := []int{128, 1, fi7, 0}
            _ = v8
        case 1:
            v9 := GenS1{F0: -751}
            _ = v9
        case 0:
            var v10 uint32 = 2
            _ = v10
        default:
            v11 := []int{(fi7 * (fi7 % ((fi7) | 1))), 4294967295, (fi7 + fi7)}
            _ = v11
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int64 = -654
        switch "hello" {
        case "a":
            v14 := []int{(fi12 << ((4294967295) & 63)), fi12}
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int8 = -128
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS1{F0: 32767}
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS2{F0: 306, F1: 0}
        fmt.Println(p2, p0)
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v6
    return false
}
func genFn1(p0 bool, p1 int32) uint8 {
    p0 = (2 != 18446744073709551615)
    var v20 uint8 = 153
    if (1e10 > (1.5 / 1.0)) {
        var v21 uint64 = 266
        _ = v21
    } else {
        v22 := []int{}
        var v23 uint32 = 42
        _ = len(v22)
        _ = v22
        _ = v23
    }
    v24 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v20
    _ = v24
    return (v20 >> ((411) & 63))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 int16 = 42
    var v4 int8 = 1
    var v5 uint64 = 2
    delete(v2, 1)
    v4 += v4
    if (0 == 1) {
        v5 = (v5 | (v5 - (v5 ^ v5)))
        var v25 uint = 65535
        _ = v25
    } else {
        var v26 int = -2147483648
        var v27 int32 = -90
        if false {
            v28 := GenS3{F0: 1}
            v29 := v2[118]
            _ = v28
            _ = v29
        } else {
            if ((!true) || false) {
                v26 -= v26
            } else {
                v30 := []int{}
                v31 := GenS2{F0: -9223372036854775808, F1: (v4 << ((0) & 63))}
                switch (("" + "\n\t") + "") {
                case "hello":
                    fmt.Println(v4)
                }
                _ = v30
                _ = v31
            }
            var v32 int16 = v3
            v27 = v27
            _ = v32
        }
        _ = v26
        _ = v27
    }
    fmt.Println(v5, v3)
    v33 := v1[128]
    v34 := v1[32768]
    for fi35 := 0; fi35 < 13; fi35++ {
        if (255 > 2) {
            if (-9223372036854775808 < -1) {
                v3 += (^(v3))
                fmt.Println(v34, v33)
            } else {
                var v36 float32 = ((1.5 * 1e10) / 1.0)
                if (((0.0 + 1e10) - 1.5) != 1e10) {
                    for fi37 := 0; fi37 < 10; fi37++ {
                        var v38 int64 = int64(v4)
                        if (genFn1(v33, 2147483647) <= 128) {
                            var v39 uint16 = 42
                            _ = v39
                        } else {
                            if v33 {
                                var v40 uint64 = (^(v5))
                                v3 -= (v3 & v3)
                                v3 += v3
                                _ = v40
                            }
                            var v41 float32 = (-2.25 / 1.0)
                            fmt.Println(fi35, v36)
                            switch "a" {
                            case "\n\t":
                                delete(v1, fi35)
                            }
                            _ = v41
                        }
                        v42 := GenS1{F0: (-(fi37))}
                        _ = fi37
                        _ = v38
                        _ = v42
                    }
                }
                _ = v36
            }
            if (2 >= 189) {
                v43 := GenS2{F0: 1, F1: -85}
                _ = v43
            }
            v3 = v3
        }
        if v33 {
            var v44 string = ("" + ("foo bar" + "\n\t"))
            _ = v44
        } else {
            for fi45 := 0; fi45 < 11; fi45++ {
                if false {
                    var v46 uint64 = (v5 << ((4294967295) & 63))
                    _ = v46
                } else {
                    switch (fi45 | (fi45 >> ((0) & 63))) {
                    case -884:
                        v33 = ((v33 || (18446744073709551615 <= 2)) || (!false))
                    default:
                        v34 = ((v3 + -681) >= (v3 | -155))
                    }
                    v33 = (true && (!true))
                }
                var v47 uint16 = 42
                v48 := GenS1{F0: fi35}
                _ = fi45
                _ = v47
                _ = v48
            }
        }
        _ = fi35
    }
    v3 -= v3
    if false {
        fmt.Println(v4)
        switch "a" {
        case "foo bar":
            v3 = v3
        }
        v5 = v5
    }
    v49 := GenS4{F0: 1, F1: 42}
    v50 := GenS5{F0: v4, F1: 18446744073709551615}
    var v51 uint64 = (v5 % (((^(v5))) | 1))
    fmt.Println(v5, v3)
    v52 := GenS5{F0: (v4 ^ v4), F1: v5}
    for fi53 := 0; fi53 < 16; fi53++ {
        fmt.Println(v34, v3)
        switch (fi53 ^ (fi53 | 127)) {
        case 32768:
            for fi54 := 0; fi54 < 15; fi54++ {
                fmt.Println(v33, v33, v5)
                _ = fi54
            }
        case 9223372036854775807:
            fmt.Println(v33, v5)
        }
        var v55 float64 = float64(0.0)
        _ = fi53
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
