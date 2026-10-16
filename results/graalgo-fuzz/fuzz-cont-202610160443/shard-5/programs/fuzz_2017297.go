package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int) uint {
    p0 -= p0
    v7 := GenS1{F0: 18446744073709551615, F1: 247, F2: 772}
    if true {
        v8 := []int{(p0 ^ 68)}
        _ = v8
    } else {
        var v9 bool = true
        v10 := GenS1{F0: 65535, F1: 42, F2: 1}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 bool = true
    v2 := []int{127, 1, 2147483647, -1}
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    var v11 uint32 = 2
    v1 = (!(0 <= 1))
    v3 += v3
    var v12 int32 = -1
    for fi13 := 0; fi13 < 8; fi13++ {
        v11 = (v11 ^ (v11 << ((200) & 63)))
        var v14 uint64 = 938
        var v15 uint16 = uint16(v14)
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = append(v2, 4294967295)
    if len(v2) > 0 {
        v2[0] = 526
    }
    var v16 uint32 = (v11 | (v11 % ((2) | 1)))
    if (!false) {
        fmt.Println(v12, v1)
        v6 = (v6 + (("" + "") + ((v6 + "foo bar") + (v6 + (v6 + v6)))))
        var v17 uint = 976
        _ = v17
    } else {
        if ((v1 || true) || ((0 >= 4294967295) || v1)) {
            v18 := GenS2{F0: 255, F1: -472, F2: v6}
            switch (v6 + "hello") {
            case "\n\t":
                if v1 {
                    fmt.Println(v6, v12)
                    var v19 bool = true
                    var v20 int64 = -1
                    _ = v19
                    _ = v20
                }
            case "hello":
                var v21 bool = v1
                _ = v21
            }
            _ = v18
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := GenS3{F0: v12, F1: 464, F2: 4294967295}
                if v1 {
                    fmt.Println(fi22)
                } else {
                    v12 = 562
                }
                v24 := v5[fi22]
                _ = fi22
                _ = v23
                _ = v24
            }
            v11 = v16
        }
    }
    switch "hello" {
    case "hello":
        var v25 float64 = 3.14159
        _ = v25
    default:
        v26 := GenS1{F0: 2, F1: (v3 & 1), F2: 255}
        _ = v26
    }
    var v27 uint = 348
    v28 := v4[("" + "hello")]
    fmt.Println(v11, v16)
    _ = len(v4)
    fmt.Println(v11)
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS4{F0: 1, F1: 628, F2: 42}
        switch (fi29 - fi29) {
        case -9223372036854775808:
            if true {
                for fi31 := 0; fi31 < 11; fi31++ {
                    _ = len(v5)
                    v32 := GenS5{F0: v3, F1: 0}
                    var v33 uint64 = 813
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        default:
            for fi34 := 0; fi34 < 5; fi34++ {
                switch fi34 {
                case -2147483648:
                    _ = len(v2)
                case 75:
                    v35 := GenS2{F0: 42, F1: 42, F2: v6}
                    _ = v35
                case 960:
                    switch "hello" {
                    case "foo bar":
                        switch v6 {
                        case "a":
                            var v36 uint32 = (v16 & (^(v11)))
                            _ = v36
                        case "\n\t":
                            var v37 int = (-(fi34))
                            var v38 uint32 = (v16 * v16)
                            fmt.Println(v27, v16)
                            if (!v28) {
                                v27 = 936
                            }
                            _ = v37
                            _ = v38
                        case "":
                            var v39 int = fi29
                            _ = v39
                        default:
                            v6 = "\n\t"
                        }
                    case "a":
                        var v40 uint8 = v3
                        _ = v40
                    case "":
                        var v41 uint64 = uint64(v16)
                        _ = v41
                    }
                default:
                    v12 += v12
                }
                _ = fi34
            }
        }
        _ = fi29
        _ = v30
    }
    var v42 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
