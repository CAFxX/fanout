package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
func genFn0(p0 string, p1 uint8, p2 uint32) bool {
    p1 = (p1 | (p1 - 171))
    var v6 int32 = -2147483648
    p0 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1(p0 int64, p1 int64, p2 uint8) uint64 {
    var v7 string = ""
    for fi8 := 0; fi8 < 1; fi8++ {
        if false {
            var v9 uint32 = 42
            if true {
                if (255 <= 0) {
                    p1 = 870
                    switch (fi8 >> ((662) & 63)) {
                    case 32767:
                        switch int(p1) {
                        case 32767:
                            p0 -= (p1 << ((42) & 63))
                        default:
                            p2 = (^(p2))
                        }
                    case 9223372036854775807:
                        for fi10 := 0; fi10 < 18; fi10++ {
                            fmt.Println(v7, p2, fi10)
                            fmt.Println(p2)
                            _ = fi10
                        }
                    case -2147483648:
                        var v11 uint = 0
                        _ = v11
                    }
                }
            } else {
                var v12 string = ((("hello" + v7) + "foo bar") + (((v7 + "\n\t") + v7) + v7))
                var v13 int32 = 42
                v14 := GenS1{F0: 165, F1: 424, F2: v12}
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int = 32768
                _ = fi15
                _ = v16
            }
            _ = v9
        }
        if true {
            for fi17 := 0; fi17 < 13; fi17++ {
                switch (^(fi17)) {
                case 108:
                    fmt.Println(p0, p0)
                case -201:
                    v18 := []int{(fi8 >> ((1) & 63)), -9223372036854775808, (-(fi8))}
                    _ = v18
                case 1:
                    v19 := []int{-2147483648, (-(fi17)), 9223372036854775807, (fi8 % ((fi8) | 1))}
                    _ = v19
                default:
                    switch v7 {
                    case "foo bar":
                        v20 := []int{(^(fi8))}
                        _ = v20
                    case "\n\t":
                        fmt.Println(p2, p2)
                    case "":
                        fmt.Println(p2, p2)
                    default:
                        if (!true) {
                            v21 := []int{(-(fi8)), (fi17 / (((-(fi17))) | 1)), (fi8 & (fi17 * (fi17 << ((18446744073709551615) & 63)))), (fi17 % ((-9223372036854775808) | 1))}
                            var v22 int16 = -1
                            fmt.Println(p1)
                            _ = v21
                            _ = v22
                        }
                    }
                }
                var v23 bool = false
                _ = fi17
                _ = v23
            }
            fmt.Println(fi8)
        }
        _ = fi8
    }
    switch 1 {
    case 42:
        p0 -= (p1 ^ (p0 - (p0 ^ (p1 / ((p1) | 1)))))
    }
    v24 := GenS1{F0: p2, F1: 2, F2: (v7 + v7)}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v7
    return 487
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int8 = 11
    v5 := make(map[string]string)
    v25 := GenS1{F0: 64, F1: 65535, F2: ("hello" + "")}
    var v26 uint16 = 0
    var v27 int8 = -128
    fmt.Println(v27, v27, v4)
    var v28 int16 = -274
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := []int{fi29, -9223372036854775808}
        var v31 int64 = 9223372036854775807
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int16 = v28
    var v33 float32 = 1e10
    v27 = (v27 | v4)
    switch 940 {
    case 127:
        var v34 uint32 = 1
        _ = v34
    case 32767:
        v35 := []int{4294967295}
        _ = v35
    case 32768:
        if false {
            var v36 uint16 = v26
            v37 := []int{42, -917, -595}
            _ = v36
            _ = v37
        }
    default:
        switch ("a" + "") {
        case "":
            var v38 uint32 = 734
            _ = v38
        case "a":
            fmt.Println(v28, v27)
        }
    }
    v32 = 1000
    var v39 uint16 = 32767
    fmt.Println(v26, v32, v28)
    v40 := []int{127, 385, 32768, -9223372036854775808}
    fmt.Println(v27, v28)
    v28 += (v32 & (^(v28)))
    var v41 uint16 = v39
    v27 -= (v4 % ((v4) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
