package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32) uint {
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int = (^(fi6))
        _ = fi6
        _ = v7
    }
    if (!false) {
        if (false && false) {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: fi8, F1: 2}
                _ = fi8
                _ = v9
            }
            if (false && true) {
                fmt.Println(v5, v5)
                var v10 int8 = -128
                if (0 <= 1) {
                    if (-167 > 569) {
                        if true {
                            p0 = 0
                            v11 := []int{32768, 95, 0}
                            _ = len(v11)
                            _ = v11
                        }
                    } else {
                        v12 := []int{1, -763}
                        _ = len(v12)
                        if (0 != -649) {
                            switch -9223372036854775808 {
                            case 2147483647:
                                var v13 string = "\n\t"
                                _ = v13
                            }
                        }
                        _ = v12
                    }
                    var v14 int16 = -1
                    if (((183 > 2) || false) || true) {
                        var v15 float64 = -2.25
                        _ = v15
                    }
                    _ = v14
                } else {
                    if (792 >= (p0 & p0)) {
                        v10 = (v10 ^ 42)
                    } else {
                        p0 = (p0 ^ (p0 << ((18446744073709551615) & 63)))
                        v16 := GenS2{F0: 0, F1: "a", F2: 458}
                        v10 = -23
                        _ = v16
                    }
                    var v17 bool = (!true)
                    v18 := []int{-2147483648}
                    _ = v17
                    _ = v18
                }
                _ = v10
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 uint8 = 172
                    _ = fi19
                    _ = v20
                }
            }
        }
    } else {
        v21 := []int{}
        p0 += (p0 / ((p0) | 1))
        _ = v21
    }
    _ = p0
    _ = v5
    return 769
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) int {
    var v22 bool = false
    v23 := GenS3{F0: 312}
    for fi24 := 0; fi24 < 9; fi24++ {
        p2 -= p2
        var v25 int32 = 107
        _ = fi24
        _ = v25
    }
    fmt.Println(v22)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 4294967295
}

func main() {
    var v1 int16 = 16
    var v2 int = 32767
    v3 := make(map[int]string)
    var v4 float32 = 1.5
    _ = len(v3)
    v26 := GenS2{F0: 282, F1: "foo bar", F2: 265}
    fmt.Println(v1)
    v27 := GenS4{F0: 4294967295}
    _ = len(v3)
    v28 := []int{v2, 128, genFn1(2, 2, 4294967295)}
    for fi29 := 0; fi29 < 6; fi29++ {
        if true {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 int64 = 9223372036854775807
                if ((100 > 51) && false) {
                    switch (fi30 + fi30) {
                    case 2147483647:
                        var v32 uint16 = 255
                        _ = v32
                    default:
                        v4 -= v4
                    }
                    if (-1 < -128) {
                        if (!false) {
                            if (128 != (fi29 & fi30)) {
                                v33 := GenS4{F0: 42}
                                delete(v3, genFn1(187, 317, 708))
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                var v34 uint32 = 2147483648
                                _ = v33
                                _ = v34
                            }
                            v35 := GenS4{F0: 2147483647}
                            _ = v35
                        } else {
                            fmt.Println(v31, v4)
                            var v36 uint = 1
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            _ = v36
                        }
                    }
                    fmt.Println(v1, v1)
                } else {
                    var v37 uint8 = 1
                    if (4294967295 != 4294967295) {
                        fmt.Println(fi29, v4, v1)
                        var v38 int8 = -1
                        fmt.Println(v38, v1)
                        _ = v38
                    }
                    if (true && false) {
                        fmt.Println(fi30, v31, v1)
                    }
                    _ = v37
                }
                _ = fi30
                _ = v31
            }
            for fi39 := 0; fi39 < 15; fi39++ {
                if (true && true) {
                    fmt.Println(fi39)
                }
                _ = fi39
            }
            var v40 int16 = v1
            _ = v40
        } else {
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 int32 = 0
                v42 = (v42 & v42)
                _ = fi41
                _ = v42
            }
        }
        _ = fi29
    }
    if len(v28) > 0 {
        v28[0] = v2
    }
    var v43 int = 1
    v44 := GenS3{F0: (v1 * (v1 % (((v1 & v1)) | 1)))}
    v4 -= (v4 + (0.1 / v4))
    var v45 int32 = -2147483648
    if (1 >= -1) {
        switch "\n\t" {
        case "":
            fmt.Println(v4, v45)
        default:
            v46 := GenS3{F0: 3}
            _ = v46
        }
        fmt.Println(v4, v2, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
