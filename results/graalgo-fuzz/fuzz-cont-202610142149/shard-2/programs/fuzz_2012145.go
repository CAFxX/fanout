package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int16; F1 bool; F2 string }
type GenS6 struct { F0 uint32; F1 int8; F2 string }
type GenS7 struct { F0 int8 }
func genFn0(p0 int32) int64 {
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: 716, F1: 42, F2: 41}
    var v9 int16 = int16(p0)
    v9 = (v9 | (v9 << ((2) & 63)))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint) int8 {
    for fi10 := 0; fi10 < 3; fi10++ {
        if (!(genFn0(305) <= 42)) {
            if (351 >= 0) {
                if true {
                    p0 = (p0 >> (((^(p0))) & 63))
                    var v11 uint = p0
                    _ = v11
                }
                p0 += p0
                var v12 uint64 = 18446744073709551615
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    fmt.Println(fi10)
                    var v14 uint16 = 112
                    switch 201 {
                    case -9223372036854775808:
                        fmt.Println(v14, fi10)
                    case 4294967295:
                        v15 := GenS1{F0: 42, F1: 42, F2: 159}
                        _ = v15
                    }
                    var v16 uint16 = (v14 / ((v14) | 1))
                    v14 -= v16
                    if (!true) {
                        v17 := GenS2{F0: -2147483648, F1: 42, F2: 4294967295}
                        _ = v17
                    } else {
                        if (false || false) {
                            var v18 uint32 = 493
                            _ = v18
                        } else {
                            v19 := []int{fi10, fi10}
                            var v20 float32 = 3.14159
                            _ = v19
                            _ = v20
                        }
                    }
                    switch fi10 {
                    case -1:
                        v16 = v14
                    default:
                        var v21 int8 = 100
                        _ = v21
                    }
                    _ = fi13
                    _ = v14
                    _ = v16
                }
                var v22 float64 = 3.14159
                _ = v22
            }
        } else {
            v23 := []int{2147483647, (fi10 << (((^(p0))) & 63)), fi10, -1}
            p0 = 2
            var v24 int = (fi10 + fi10)
            _ = v23
            _ = v24
        }
        _ = fi10
    }
    var v25 uint64 = 2
    if false {
        var v26 uint16 = 780
        _ = v26
    } else {
        var v27 uint16 = 255
        if false {
            fmt.Println(v25, v27)
            fmt.Println(v27)
        }
        _ = v27
    }
    v28 := []int{}
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            if false {
                var v31 int8 = 1
                v32 := GenS3{F0: "\n\t", F1: -1}
                _ = v31
                _ = v32
            }
            switch "foo bar" {
            case "\n\t":
                v33 := GenS1{F0: -168, F1: 227, F2: 175}
                _ = v33
            case "":
                v25 = (v25 | 2)
            case "hello":
                fmt.Println(v25, fi29, v25)
            default:
                fmt.Println(fi29, v25, fi30)
            }
            var v34 int64 = -117
            switch (fi30 & fi29) {
            case 128:
                p0 = (p0 + p0)
            case -9223372036854775808:
                switch ("hello" + ("hello" + "hello")) {
                case "foo bar":
                    v34 = (v34 ^ (v34 & (v34 + v34)))
                default:
                    switch (("foo bar" + "hello") + "foo bar") {
                    case "\n\t":
                        if len(v28) > 0 {
                            v28[0] = 32767
                        }
                    case "foo bar":
                        var v35 uint8 = 42
                        _ = v35
                    default:
                        var v36 int8 = -1
                        _ = v36
                    }
                }
            case 1:
                v34 += (v34 - int64(v25))
            default:
                var v37 uint8 = 128
                _ = v37
            }
            v38 := GenS4{F0: "foo bar", F1: (p0 + p0)}
            _ = fi30
            _ = v34
            _ = v38
        }
        if len(v28) > 0 {
            v28[0] = fi29
        }
        v25 += (v25 + 0)
        _ = fi29
    }
    _ = p0
    _ = v25
    _ = v28
    return 1
}

func main() {
    var v1 float64 = ((0.0 / -2.25) / 1e10)
    var v2 int = -2147483648
    var v3 uint32 = 2147483647
    var v4 float64 = v1
    var v5 int64 = -450
    v6 := make(map[uint32]bool)
    v39 := GenS2{F0: 42, F1: v5, F2: (^(v3))}
    v40 := v6[v3]
    v1 += (v4 + 0.0)
    v41 := GenS4{F0: "a", F1: 1}
    v42 := GenS5{F0: 42, F1: false, F2: "foo bar"}
    v6[v3] = true
    if (0 == 935) {
        delete(v6, (v3 ^ v3))
        if (!false) {
            v5 = (v5 * v5)
        }
    }
    v4 = (1.5 + 0.0)
    v43 := []int{v2, (v2 + 0), v2, (v2 % ((v2) | 1))}
    fmt.Println(v5, v1, v5)
    v4 = v1
    v2 = v2
    for fi44 := 0; fi44 < 6; fi44++ {
        switch "foo bar" {
        case "foo bar":
            v4 -= (v4 * v4)
        case "hello":
            if false {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v3, fi44, v1)
                default:
                    var v45 int32 = 679
                    _ = v45
                }
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 int64 = 9223372036854775807
                    _ = fi46
                    _ = v47
                }
                _ = len(v43)
            }
        }
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint64 = 644
            if (1 == fi44) {
                var v50 int64 = v5
                v5 -= genFn0(0)
                var v51 int32 = 42
                _ = v50
                _ = v51
            }
            delete(v6, (v3 - uint32(v49)))
            _ = fi48
            _ = v49
        }
        _ = fi44
    }
    switch (v2 << ((42) & 63)) {
    case 127:
        var v52 bool = v40
        _ = v52
    case 9223372036854775807:
        var v53 int16 = 1
        _ = v53
    case 812:
        var v54 float64 = v4
        _ = v54
    default:
        v55 := GenS6{F0: 4294967295, F1: 127, F2: "hello"}
        _ = v55
    }
    var v56 int = -296
    for fi57 := 0; fi57 < 15; fi57++ {
        for fi58 := 0; fi58 < 1; fi58++ {
            if v40 {
                v1 += (3.14159 * -2.25)
                fmt.Println(v5, v2, fi57)
            } else {
                v59 := GenS3{F0: "foo bar", F1: 32767}
                var v60 bool = ((v40 || v40) || (v3 <= 2147483647))
                var v61 int64 = genFn0(2147483647)
                _ = v59
                _ = v60
                _ = v61
            }
            v62 := GenS7{F0: int8(v5)}
            _ = fi58
            _ = v62
        }
        _ = fi57
    }
    v3 += (v3 % (((v3 >> ((787) & 63))) | 1))
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v56
    _ = v6
}
