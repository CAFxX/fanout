package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 string, p1 int16, p2 uint8) int16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int = fi6
        if true {
            p2 -= p2
        } else {
            switch fi6 {
            case 446:
                switch p0 {
                case "":
                    if true {
                        v8 := []int{(v7 & v7)}
                        var v9 int64 = 0
                        var v10 int32 = 0
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                default:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := []int{(fi11 * 32768), (v7 + (fi6 - 128)), fi11}
                        var v13 int16 = p1
                        v14 := GenS1{F0: int64(fi11), F1: 4294967295}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
            case 781:
                var v15 uint16 = 32767
                _ = v15
            default:
                switch v7 {
                case 9223372036854775807:
                    p1 += (p1 - (^(p1)))
                case 128:
                    for fi16 := 0; fi16 < 12; fi16++ {
                        switch -9223372036854775808 {
                        case 32768:
                            switch fi6 {
                            case 1:
                                var v17 string = p0
                                _ = v17
                            default:
                                v7 = (fi16 & 32767)
                            }
                        case 32767:
                            if true {
                                v18 := GenS1{F0: -1, F1: 42}
                                if (((true && true) && true) && false) {
                                    v19 := []int{}
                                    var v20 uint16 = 65535
                                    _ = v19
                                    _ = v20
                                }
                                if (((p0 + p0) != ("foo bar" + (p0 + p0))) && (!false)) {
                                    p2 += (p2 * p2)
                                    v21 := GenS1{F0: 0, F1: 1}
                                    v22 := []int{-506, fi6, 2147483647}
                                    _ = v21
                                    _ = v22
                                }
                                _ = v18
                            }
                        default:
                            p2 -= p2
                        }
                        _ = fi16
                    }
                case -1:
                    var v23 int64 = 42
                    _ = v23
                default:
                    var v24 int16 = (p1 + 32767)
                    _ = v24
                }
            }
        }
        _ = fi6
        _ = v7
    }
    var v25 bool = true
    v26 := GenS1{F0: -9223372036854775808, F1: 2147483647}
    var v27 int32 = -1
    var v28 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    return 108
}

func main() {
    var v1 int32 = -371
    v2 := []uint32{2147483647, 4294967295, 2, 65535}
    v3 := make(map[string]bool)
    v4 := []string{}
    var v5 uint64 = 0
    if len(v2) > 0 {
        v2[0] = 481
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            if ((!(v1 < 2147483647)) && true) {
                var v31 float64 = (0.1 * float64(1e10))
                v31 -= (v31 / float64(float32((v31 * v31))))
                v32 := GenS2{F0: 0, F1: 904}
                v3["hello"] = (-1 < 0)
                _ = v31
                _ = v32
            } else {
                v5 += (v5 & v5)
                var v33 uint64 = v5
                var v34 bool = (42 < 2)
                _ = v33
                _ = v34
            }
            var v35 float32 = 3.14159
            _ = fi30
            _ = v35
        }
        _ = fi29
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        for fi38 := 0; fi38 < 13; fi38++ {
            switch fi38 {
            case 9223372036854775807:
                if (-1 != -678) {
                    var v39 uint = 18446744073709551615
                    _ = v39
                } else {
                    switch (-(fi38)) {
                    case 32767:
                        v5 -= (v5 ^ 42)
                        var v40 uint64 = v5
                        _ = v40
                    case -1:
                        switch fi38 {
                        case 9223372036854775807:
                            v41 := GenS3{F0: 2, F1: fi36, F2: 4294967295}
                            _ = v41
                        case 2147483647:
                            v5 -= (v5 % ((v5) | 1))
                        }
                    case 0:
                        if true {
                            var v42 int8 = -8
                            _ = v42
                        }
                    default:
                        v2 = append(v2, 750)
                    }
                    switch (fi38 % ((int(v5)) | 1)) {
                    case 9223372036854775807:
                        if true {
                            var v43 float64 = (float64(1e10) - 0.1)
                            v44 := v3["hello"]
                            v1 -= 42
                            _ = v43
                            _ = v44
                        }
                    case -2147483648:
                        switch "foo bar" {
                        case "":
                            v1 -= v1
                        case "a":
                            var v45 uint8 = 128
                            _ = v45
                        default:
                            var v46 uint32 = 4294967295
                            _ = v46
                        }
                    default:
                        var v47 uint32 = 0
                        _ = v47
                    }
                }
            }
            _ = fi38
        }
        v48 := GenS3{F0: 4, F1: fi36, F2: v5}
        _ = fi36
        _ = v37
        _ = v48
    }
    var v49 int16 = 1000
    if (1 < -9223372036854775808) {
        v2 = append(v2, 65535)
    } else {
        v1 = v1
        var v50 uint16 = 2
        _ = v50
    }
    var v51 uint8 = 127
    _ = len(v3)
    fmt.Println(v5, v5)
    var v52 uint32 = 59
    delete(v3, "")
    fmt.Println(v1, v5, v49)
    v51 += v51
    for fi53 := 0; fi53 < 6; fi53++ {
        v3["\n\t"] = true
        switch "hello" {
        case "":
            var v54 uint16 = 1
            _ = v54
        case "a":
            v5 = (v5 & (^(v5)))
        case "foo bar":
            v52 -= (v52 ^ (v52 * (^(v52))))
        }
        v55 := GenS2{F0: 56, F1: 698}
        _ = fi53
        _ = v55
    }
    fmt.Println(v51, v51)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v52
}
