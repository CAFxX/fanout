package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 string, p1 uint8, p2 bool) uint {
    v4 := GenS1{F0: -851, F1: 128}
    var v5 int16 = 616
    p1 -= p1
    var v6 uint32 = 42
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 uint64, p2 uint8) uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, p1, fi7)
        var v8 int16 = p0
        if ((p2 - 0) == (p2 | (p2 << ((genFn0("\n\t", (p2 - p2), false)) & 63)))) {
            var v9 int64 = 9223372036854775807
            if true {
                v8 += 920
                switch "\n\t" {
                case "":
                    var v10 int32 = -376
                    _ = v10
                case "\n\t":
                    var v11 bool = (p1 <= (p1 ^ p1))
                    _ = v11
                case "hello":
                    if ("hello" > "\n\t") {
                        v12 := []int{}
                        _ = v12
                    } else {
                        v9 -= v9
                        for fi13 := 0; fi13 < 17; fi13++ {
                            v14 := GenS2{F0: 579, F1: -2147483648, F2: fi13}
                            _ = fi13
                            _ = v14
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v15 uint64 = 42
                            _ = v15
                        }
                    }
                }
                v16 := []int{(fi7 & 0), 232, (fi7 << ((65535) & 63))}
                _ = v16
            } else {
                v17 := GenS3{F0: 0, F1: (v9 / ((42) | 1))}
                v18 := GenS3{F0: int32(v9), F1: (v9 / (((-(v9))) | 1))}
                v9 += v9
                _ = v17
                _ = v18
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v19 int8 = 56
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 122
    var v3 uint16 = 1
    var v20 string = "hello"
    var v21 float64 = (0.0 * 0.0)
    fmt.Println(v21)
    var v22 int8 = 1
    v2 = genFn1(-32768, 2, v2)
    v23 := GenS2{F0: 801, F1: int32(v22), F2: -2147483648}
    for fi24 := 0; fi24 < 11; fi24++ {
        v21 -= ((v21 - 0.0) / 0.1)
        v21 += float64(((1e10 / 1e10) / (1.5 + -2.25)))
        _ = fi24
    }
    v25 := GenS3{F0: 935, F1: 42}
    if false {
        var v26 float64 = v21
        for fi27 := 0; fi27 < 20; fi27++ {
            v1[2147483647] = (^(fi27))
            _ = fi27
        }
        _ = v26
    }
    switch ((("a" + (v20 + v20)) + "a") + v20) {
    case "hello":
        v28 := v1[4294967295]
        _ = v28
    case "foo bar":
        for fi29 := 0; fi29 < 7; fi29++ {
            for fi30 := 0; fi30 < 8; fi30++ {
                v1[2147483647] = 4294967295
                switch 0 {
                case 4294967295:
                    if true {
                        v31 := []int{(fi30 & -1)}
                        _ = v31
                    } else {
                        if false {
                            v32 := []int{fi29, (fi29 << ((0) & 63)), int(v3)}
                            _ = v32
                        } else {
                            switch fi30 {
                            case 9223372036854775807:
                                v21 = v21
                            case -9223372036854775808:
                                v33 := GenS2{F0: 892, F1: -2147483648, F2: fi29}
                                _ = v33
                            default:
                                if false {
                                    _ = len(v1)
                                    v34 := GenS4{F0: 876}
                                    _ = v34
                                } else {
                                    var v35 uint16 = (v3 + (v3 | 1))
                                    if true {
                                        delete(v1, 1)
                                        v22 = (v22 * -1)
                                        var v36 float64 = v21
                                        _ = v36
                                    }
                                    v37 := GenS1{F0: (fi29 >> ((65535) & 63)), F1: genFn0((((v20 + (v20 + "hello")) + v20) + "foo bar"), v2, true)}
                                    _ = v35
                                    _ = v37
                                }
                            }
                            v22 -= (v22 % (((v22 / ((v22) | 1))) | 1))
                        }
                        v38 := GenS3{F0: 2147483647, F1: -459}
                        v39 := GenS5{F0: genFn0("hello", 78, (245 <= 42)), F1: (fi30 - fi30), F2: 42}
                        _ = v38
                        _ = v39
                    }
                default:
                    fmt.Println(v22, fi29, v2)
                    v40 := GenS5{F0: 42, F1: fi30, F2: -9223372036854775808}
                    var v41 int32 = 0
                    v42 := []int{fi29, fi30, (-(fi30))}
                    _ = v40
                    _ = v41
                    _ = v42
                }
                _ = fi30
            }
            var v43 uint32 = 65535
            v44 := GenS2{F0: 737, F1: 0, F2: fi29}
            _ = fi29
            _ = v43
            _ = v44
        }
    default:
        var v45 uint16 = (^(v3))
        _ = v45
    }
    fmt.Println(v21, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
