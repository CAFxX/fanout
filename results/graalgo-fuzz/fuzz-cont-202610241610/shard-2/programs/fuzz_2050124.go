package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int32; F1 string; F2 int32 }
func genFn0(p0 uint, p1 int32) int32 {
    var v6 int32 = (p1 * -114)
    switch 892 {
    case 31:
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: -1, F1: 42, F2: (p1 + 0)}
            v9 := []int{fi7, (fi7 / (((fi7 / ((fi7) | 1))) | 1)), (^(fi7))}
            if ((!true) && (255 >= 32767)) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(fi10, v6, fi7)
                    }
                    _ = fi10
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if false {
                    if len(v9) > 0 {
                        v9[0] = (^(fi7))
                    }
                    if (!(2 < 1)) {
                        var v11 uint8 = 2
                        _ = v11
                    }
                } else {
                    fmt.Println(fi7, fi7, v6)
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                if (29 < -1) {
                    p1 = p1
                    p0 += 255
                }
                var v13 uint8 = 0
                switch "\n\t" {
                case "a":
                    var v14 uint = 0
                    _ = v14
                case "":
                    v15 := []int{(fi12 & -1)}
                    _ = v15
                default:
                    v6 -= (p1 & v6)
                }
                _ = v13
            } else {
                fmt.Println(p0, p0)
                for fi16 := 0; fi16 < 17; fi16++ {
                    if true {
                        p0 += 1
                        v17 := []int{207}
                        _ = v17
                    } else {
                        v18 := []int{fi16, (fi16 >> (((p0 + 705)) & 63)), 2147483647, -2147483648}
                        v19 := GenS2{F0: 42, F1: 1, F2: true}
                        if true {
                            var v20 float64 = 1.5
                            _ = v20
                        }
                        _ = v18
                        _ = v19
                    }
                    if false {
                        var v21 uint16 = 1
                        switch fi12 {
                        case 940:
                            v22 := []int{fi16, (fi16 * fi16)}
                            _ = v22
                        case 32768:
                            var v23 int64 = 201
                            _ = v23
                        case 2147483647:
                            v24 := GenS1{F0: 911, F1: 63, F2: v6}
                            if false {
                                v25 := []int{(^(fi12)), fi16, fi12}
                                _ = v25
                            } else {
                                var v26 int16 = 32767
                                v6 = (p1 | (-(p1)))
                                _ = v26
                            }
                            _ = v24
                        default:
                            v21 += v21
                        }
                        _ = v21
                    }
                    fmt.Println(p1, v6, fi12)
                    _ = fi16
                }
            }
            var v27 int8 = 127
            if (true || ((65535 <= 1) && true)) {
                v6 = -1
                v6 += (v6 * (p1 << ((p0) & 63)))
                v28 := GenS2{F0: 1, F1: 42, F2: (("\n\t" + "") >= "hello")}
                _ = v28
            }
            _ = fi12
            _ = v27
        }
    case 32768:
        p0 = p0
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 bool = (true && (fi29 < fi29))
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v6
    return (p1 - (v6 ^ -324))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 float64 = float64(0.0)
    v5 := make(map[int]int)
    fmt.Println(v4)
    v4 -= v4
    var v31 int32 = 1
    var v32 int32 = genFn0(4294967295, (v31 << ((18446744073709551615) & 63)))
    var v33 float32 = 1e10
    v34 := GenS2{F0: 0, F1: -378, F2: true}
    v35 := GenS3{F0: 855, F1: 0}
    v36 := GenS2{F0: 255, F1: 0, F2: false}
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 int16 = 666
        v32 = (^(v32))
        _ = fi37
        _ = v38
    }
    fmt.Println(v32)
    var v39 int8 = 0
    if true {
        if true {
            var v40 int32 = v32
            v2[127] = ("\n\t" + "\n\t")
            _ = v40
        } else {
            fmt.Println(v31, v4, v33)
            v41 := []int{127, 636, 703, -2147483648}
            _ = v41
        }
        v42 := GenS4{F0: -761, F1: "", F2: v32}
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
