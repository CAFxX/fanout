package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 uint) int {
    v6 := []int{-442}
    var v7 int32 = 2147483647
    fmt.Println(p0)
    var v8 uint32 = 521
    v9 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 uint32) bool {
    v10 := []int{760}
    var v11 float64 = -2.25
    _ = p0
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 int16 = 42
    v2 := []int{0, -1}
    var v3 int64 = 1
    var v4 int32 = 2147483647
    var v5 int8 = -128
    var v12 int8 = v5
    v13 := GenS1{F0: 32767}
    var v14 uint16 = 2
    var v15 int8 = v5
    if false {
        v16 := GenS1{F0: 4294967295}
        switch "\n\t" {
        case "":
            switch -2147483648 {
            case 94:
                for fi17 := 0; fi17 < 16; fi17++ {
                    fmt.Println(fi17, fi17, v14)
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if len(v2) > 0 {
                            v2[0] = fi17
                        }
                        _ = fi18
                    }
                    fmt.Println(v12, v15)
                    _ = fi17
                }
            case -1:
                v4 = v4
            case -2147483648:
                v19 := GenS1{F0: 416}
                _ = v19
            default:
                if false {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 int32 = (^(v4))
                        _ = fi20
                        _ = v21
                    }
                }
            }
        }
        _ = v16
    }
    fmt.Println(v5, v1)
    var v22 bool = (((v5 << ((2) & 63)) > -128) || (!true))
    switch "hello" {
    case "hello":
        v23 := GenS1{F0: 0}
        _ = v23
    case "a":
        var v24 int32 = v4
        _ = v24
    case "foo bar":
        if v22 {
            var v25 uint64 = 1
            v26 := GenS1{F0: -628}
            v15 = (v15 | (v15 / ((v15) | 1)))
            _ = v25
            _ = v26
        } else {
            var v27 int64 = v3
            var v28 bool = v22
            if (false && v22) {
                v4 = v4
                fmt.Println(v5, v4)
                var v29 float64 = 1e10
                _ = v29
            } else {
                for fi30 := 0; fi30 < 11; fi30++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v14 += (^(v14))
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v27 = (v27 ^ (v27 >> ((2) & 63)))
                        v12 = v5
                    }
                    _ = fi30
                }
            }
            _ = v27
            _ = v28
        }
    }
    var v31 int64 = (v3 * v3)
    if (("foo bar" == ("\n\t" + "foo bar")) || true) {
        if v22 {
            v12 = v5
        } else {
            switch -1 {
            case 42:
                var v32 int16 = 645
                _ = v32
            default:
                v5 = (v15 ^ (v12 << ((18446744073709551615) & 63)))
            }
        }
        v4 = -1
    } else {
        v12 = v15
        v33 := GenS2{F0: 42, F1: 422, F2: true}
        if len(v2) > 0 {
            v2[0] = -1
        }
        _ = v33
    }
    if ((true || true) || v22) {
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 int32 = 981
            _ = fi34
            _ = v35
        }
        if v22 {
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS1{F0: (fi36 >> ((255) & 63))}
                v1 = v1
                v14 = (^(v14))
                _ = fi36
                _ = v37
            }
            if true {
                _ = len(v2)
            }
            v4 += (^(v4))
        }
        var v38 float64 = float64(0.1)
        _ = v38
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v39 uint8 = 192
    v31 += 0
    v31 = v3
    var v40 uint16 = v14
    v5 = v5
    if v22 {
        fmt.Println(v5)
        if (!(v22 || v22)) {
            fmt.Println(v15, v31, v39)
            if len(v2) > 0 {
                v2[0] = 32767
            }
            if v22 {
                if true {
                    if ((((float32(1e10) * 0.0) + 1.5) - (1e10 + -2.25)) < 0.0) {
                        if (float64(-2.25) > (1e10 / (1e10 * ((1e10 - 1.5) - 1e10)))) {
                            var v41 int8 = (-(v5))
                            v12 = (v5 & (v5 | (v5 << ((1) & 63))))
                            fmt.Println(v3)
                            _ = v41
                        } else {
                            var v42 bool = ((1.5 - float32((1e10 * 3.14159))) <= -2.25)
                            v43 := GenS3{F0: (v4 + (v4 / (((v4 & -517)) | 1))), F1: (v1 % ((-475) | 1))}
                            _ = v42
                            _ = v43
                        }
                        v2 = append(v2, genFn0(uint(v39)))
                    }
                } else {
                    switch int(v1) {
                    case 32767:
                        v22 = (v3 == 42)
                    case -9223372036854775808:
                        v44 := GenS2{F0: 2147483647, F1: 65535, F2: (42 > 1)}
                        _ = v44
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    fmt.Println(v40, v15)
                }
            }
        } else {
            v2 = append(v2, -2147483648)
            v39 += (^(v39))
        }
        v14 += v14
    } else {
        v31 -= v3
        v31 = (v3 % ((v3) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
