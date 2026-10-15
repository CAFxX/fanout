package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int32; F1 string; F2 string }
func genFn0() bool {
    if true {
        var v7 bool = (113 <= 42)
        _ = v7
    }
    v8 := GenS1{F0: -2147483648, F1: 0, F2: -2147483648}
    _ = v8
    return (-1 > -1)
}
func genFn1(p0 string) bool {
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = 9223372036854775807
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 float64 = 1e10
                if false {
                    var v13 int32 = 0
                    var v14 uint = 4294967295
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(v12, fi11, v12)
                    v15 := []int{-2147483648, (-(fi9))}
                    _ = v15
                }
                if (false && (false || false)) {
                    p0 = p0
                    v16 := GenS1{F0: 42, F1: 930, F2: -1}
                    v17 := []int{fi9, (fi11 << ((18446744073709551615) & 63)), fi11}
                    fmt.Println(p0, fi9, v10)
                    _ = v16
                    _ = v17
                } else {
                    v18 := []int{(^(fi11)), fi11, -1}
                    _ = v18
                }
                _ = fi11
                _ = v12
            }
            for fi19 := 0; fi19 < 12; fi19++ {
                switch "" {
                case "\n\t":
                    v20 := GenS1{F0: 734, F1: 926, F2: 1}
                    _ = v20
                case "foo bar":
                    var v21 bool = (0 <= 65535)
                    _ = v21
                case "a":
                    fmt.Println(fi19)
                default:
                    var v22 int64 = v10
                    fmt.Println(v10, fi19)
                    v23 := GenS1{F0: 442, F1: 520, F2: 1}
                    _ = v22
                    _ = v23
                }
                switch p0 {
                case "":
                    p0 = p0
                case "a":
                    var v24 int = (^(fi9))
                    _ = v24
                default:
                    switch (-(fi9)) {
                    case -9223372036854775808:
                        if false {
                            var v25 uint64 = 609
                            v26 := []int{(fi9 | 9223372036854775807), (fi19 >> ((837) & 63)), (fi9 & (-(fi19))), (fi9 / (((fi19 * -9223372036854775808)) | 1))}
                            v10 = 9223372036854775807
                            _ = v25
                            _ = v26
                        }
                    case 940:
                        v10 += (v10 + (v10 | (v10 | 42)))
                    }
                }
                _ = fi19
            }
            _ = fi9
            _ = v10
        }
    default:
        p0 = ("foo bar" + ((((p0 + ((p0 + ((((p0 + p0) + (p0 + p0)) + p0) + "")) + (p0 + ("\n\t" + p0)))) + (((((p0 + "\n\t") + p0) + p0) + (p0 + p0)) + "hello")) + ("a" + p0)) + p0))
    }
    fmt.Println(p0)
    var v27 int32 = 42
    var v28 uint8 = 10
    for fi29 := 0; fi29 < 10; fi29++ {
        p0 = "hello"
        _ = fi29
    }
    _ = p0
    _ = v27
    _ = v28
    return true
}

func main() {
    var v1 uint64 = 0
    v2 := []int{0, 42}
    v3 := []string{}
    var v4 uint = 42
    v5 := []int{}
    var v6 int = 128
    var v30 float64 = (1e10 * float64(-2.25))
    var v31 uint8 = 93
    if true {
        v32 := GenS1{F0: -726, F1: 255, F2: -681}
        if true {
            v33 := GenS2{F0: 935, F1: 0, F2: v1}
            v34 := GenS3{F0: v31, F1: 65535}
            _ = v33
            _ = v34
        } else {
            for fi35 := 0; fi35 < 2; fi35++ {
                switch "\n\t" {
                case "hello":
                    switch -2147483648 {
                    case 4294967295:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 32767:
                        var v36 int16 = 0
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 9; fi37++ {
                            v6 -= fi37
                            fmt.Println(v30, fi35, v31)
                            v38 := GenS2{F0: 42, F1: 0, F2: 42}
                            _ = fi37
                            _ = v38
                        }
                    }
                }
                _ = fi35
            }
        }
        _ = v32
    }
    var v39 uint = 255
    v40 := GenS4{F0: 0, F1: "", F2: "hello"}
    fmt.Println(v30, v1, v4)
    _ = len(v3)
    var v41 int16 = 36
    v39 += (v4 ^ (v4 * v4))
    _ = len(v2)
    var v42 uint32 = 413
    var v43 uint32 = (^(v42))
    _ = len(v3)
    var v44 int8 = 1
    switch v6 {
    case 32768:
        if ((v4 + v4) > 856) {
            var v45 float32 = 1.5
            _ = v45
        }
    case 397:
        if len(v5) > 0 {
            v5[0] = v6
        }
    default:
        v41 -= v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
