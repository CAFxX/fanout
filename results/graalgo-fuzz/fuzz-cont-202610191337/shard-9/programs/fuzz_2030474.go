package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 int, p1 int64, p2 bool) int32 {
    p1 -= (p1 % ((9223372036854775807) | 1))
    v7 := []int{p0, (p0 * p0), (p0 * (p0 / (((-(p0))) | 1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 int64) uint16 {
    fmt.Println(p1, p0)
    switch ("hello" + "foo bar") {
    case "hello":
        var v8 string = (("a" + "") + "hello")
        _ = v8
    case "\n\t":
        switch int(p1) {
        case 1:
            if (127 <= 128) {
                if false {
                    if true {
                        v9 := GenS1{F0: 1, F1: p0, F2: p1}
                        _ = v9
                    } else {
                        var v10 uint8 = 1
                        _ = v10
                    }
                    if true {
                        var v11 int8 = (p0 * (p0 << ((0) & 63)))
                        var v12 string = (("a" + "\n\t") + "foo bar")
                        v11 = v11
                        _ = v11
                        _ = v12
                    }
                    if (p0 < 100) {
                        switch 4294967295 {
                        case 127:
                            p0 += (p0 + p0)
                        }
                        v13 := []int{}
                        if (407 > 32767) {
                            _ = len(v13)
                            fmt.Println(p0, p1)
                            if true {
                                fmt.Println(p1, p1)
                            } else {
                                switch 32768 {
                                case -850:
                                    p1 -= (p1 | (p1 - (p1 * (^(p1)))))
                                case 42:
                                    for fi14 := 0; fi14 < 14; fi14++ {
                                        if false {
                                            switch (fi14 - -177) {
                                            case 9223372036854775807:
                                                var v15 uint32 = 42
                                                _ = v15
                                            case 2147483647:
                                                var v16 uint16 = 568
                                                _ = v16
                                            default:
                                                var v17 int32 = 601
                                                _ = v17
                                            }
                                            var v18 uint32 = 1
                                            _ = v18
                                        }
                                        if (3.14159 <= 0.1) {
                                            fmt.Println(p0, p0, p1)
                                            v19 := GenS1{F0: 321, F1: -79, F2: (-(p1))}
                                            _ = v19
                                        } else {
                                            p1 += (p1 + (p1 * p1))
                                        }
                                        _ = fi14
                                    }
                                case -1:
                                    switch "" {
                                    case "":
                                        var v20 uint32 = 509
                                        _ = v20
                                    }
                                default:
                                    p1 = 9223372036854775807
                                }
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    p0 = 0
                                    p1 += (-(p1))
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    _ = fi21
                                }
                                var v22 uint8 = 42
                                _ = v22
                            }
                        }
                        _ = v13
                    }
                } else {
                    v23 := GenS1{F0: 322, F1: p0, F2: 42}
                    var v24 bool = (true || (1e10 >= 3.14159))
                    _ = v23
                    _ = v24
                }
            }
        }
    case "foo bar":
        fmt.Println(p0, p0)
    default:
        p0 = p0
    }
    v25 := []int{282, -789, 587}
    if true {
        p1 -= p1
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    } else {
        var v26 string = ("a" + (("" + "a") + "a"))
        if len(v25) > 0 {
            v25[0] = -2147483648
        }
        _ = v26
    }
    if len(v25) > 0 {
        v25[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = v25
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 0
    var v3 float64 = 1e10
    var v4 uint = 65535
    var v5 int32 = -1
    v6 := make(map[int]bool)
    if (21 <= 0) {
        switch 42 {
        case 9223372036854775807:
            v27 := GenS2{F0: 2147483647, F1: -27}
            _ = v27
        case 32768:
            var v28 bool = true
            _ = v28
        }
        if true {
            v5 = 2147483647
        } else {
            var v29 int = 151
            fmt.Println(v3, v29, v29)
            for fi30 := 0; fi30 < 15; fi30++ {
                switch (fi30 % (((fi30 << ((1) & 63))) | 1)) {
                case 1:
                    switch (v29 & 127) {
                    case 128:
                        v31 := []int{247, (fi30 | (fi30 ^ 698))}
                        _ = v31
                    case 9223372036854775807:
                        var v32 uint8 = 42
                        _ = v32
                    default:
                        fmt.Println(v2, fi30)
                    }
                case 128:
                    switch (^(fi30)) {
                    case 32768:
                        v33 := []int{(v29 + (fi30 % ((v29) | 1))), (fi30 - 957), v29, (v29 << ((65535) & 63))}
                        _ = v33
                    }
                case 2147483647:
                    var v34 uint16 = v2
                    _ = v34
                default:
                    var v35 int = (v29 % (((fi30 / (((fi30 >> (((^(v4))) & 63))) | 1))) | 1))
                    _ = v35
                }
                _ = fi30
            }
            _ = v29
        }
        fmt.Println(v4)
    } else {
        _ = len(v6)
    }
    v5 -= (v5 % ((1) | 1))
    v36 := GenS2{F0: v5, F1: 123}
    if false {
        fmt.Println(v3, v4)
        v37 := []int{}
        _ = v37
    } else {
        if true {
            if false {
                v38 := v6[-541]
                _ = v38
            } else {
                v1 = v1
                v4 += (v4 ^ v4)
            }
            var v39 uint8 = 76
            if ((v4 * 65535) != v4) {
                v40 := []int{-2147483648, 4294967295, 2147483647, 4294967295}
                _ = v40
            }
            _ = v39
        } else {
            v3 = v3
            var v41 int16 = 0
            if (v2 < v2) {
                var v42 uint64 = 18446744073709551615
                _ = v42
            }
            _ = v41
        }
    }
    v43 := GenS1{F0: v5, F1: 127, F2: -1}
    v44 := GenS3{F0: -421, F1: 1}
    for fi45 := 0; fi45 < 17; fi45++ {
        for fi46 := 0; fi46 < 3; fi46++ {
            v5 = 42
            fmt.Println(v2, v1, v3)
            _ = fi46
        }
        _ = fi45
    }
    var v47 float64 = (3.14159 - ((v3 * (v3 + float64((0.1 / float32(1e10))))) + (v3 * 1e10)))
    for fi48 := 0; fi48 < 7; fi48++ {
        var v49 int32 = v5
        v50 := GenS4{F0: 0, F1: 29}
        v51 := []int{int(v4), (^(fi48))}
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    v52 := []int{}
    fmt.Println(v5)
    v53 := GenS3{F0: 139, F1: v4}
    v54 := v6[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
