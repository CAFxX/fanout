package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 int64, p1 uint) string {
    if true {
        if (0 > 2) {
            var v6 uint64 = 42
            v7 := []int{-63}
            p0 = (p0 - 9223372036854775807)
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint64 = 978
            _ = fi8
            _ = v9
        }
    } else {
        p0 = (p0 >> (((p1 & p1)) & 63))
        var v10 bool = false
        _ = v10
    }
    v11 := []int{4294967295}
    _ = p0
    _ = p1
    _ = v11
    return ("a" + "foo bar")
}
func genFn1(p0 uint64, p1 int64) int16 {
    v12 := []int{0}
    if (-691 < (^(p1))) {
        fmt.Println(p1)
        if true {
            if false {
                p1 = (p1 % ((p1) | 1))
            } else {
                if true {
                    v13 := GenS1{F0: -9223372036854775808, F1: 0}
                    var v14 int32 = -146
                    _ = v13
                    _ = v14
                } else {
                    switch 0 {
                    case 42:
                        fmt.Println(p1, p0)
                    case 397:
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v16 := GenS1{F0: (p1 & (^(p1))), F1: 271}
                            if (float32((1.5 / 1.0)) < float32(1e10)) {
                                v17 := GenS1{F0: p1, F1: 0}
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    if false {
                                        p0 = p0
                                    }
                                    _ = fi18
                                }
                                v19 := GenS1{F0: p1, F1: 114}
                                _ = v17
                                _ = v19
                            }
                            _ = fi15
                            _ = v16
                        }
                    case 32767:
                        fmt.Println(p0, p1)
                    }
                }
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                if len(v12) > 0 {
                    v12[0] = (fi20 ^ (fi20 % (((fi20 / ((-147) | 1))) | 1)))
                }
                if ("a" < "") {
                    fmt.Println(fi20, p0)
                    if false {
                        var v21 int64 = -9223372036854775808
                        for fi22 := 0; fi22 < 10; fi22++ {
                            v23 := GenS1{F0: p1, F1: 2147483647}
                            var v24 int = 0
                            if len(v12) > 0 {
                                v12[0] = 42
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    } else {
                        v25 := GenS1{F0: p1, F1: -2147483648}
                        if true {
                            for fi26 := 0; fi26 < 16; fi26++ {
                                if len(v12) > 0 {
                                    v12[0] = fi20
                                }
                                var v27 uint = 986
                                _ = fi26
                                _ = v27
                            }
                        } else {
                            for fi28 := 0; fi28 < 14; fi28++ {
                                var v29 uint = 65535
                                _ = fi28
                                _ = v29
                            }
                        }
                        _ = v25
                    }
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 float64 = -2.25
                        fmt.Println(p1, p0, p1)
                        var v32 uint64 = 42
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                var v33 bool = ((!(127 == 0)) || true)
                _ = fi20
                _ = v33
            }
            var v34 int32 = 1
            _ = v34
        }
        v35 := GenS2{F0: 100, F1: 9223372036854775807}
        _ = v35
    } else {
        p1 = 194
    }
    p0 += 2
    _ = p0
    _ = p1
    _ = v12
    return 32767
}

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = 1.5
    var v3 uint64 = 370
    var v4 uint8 = 127
    var v5 float32 = ((1e10 * 0.1) + 0.1)
    switch -952 {
    case -9223372036854775808:
        if true {
            v36 := []int{v1, (v1 / ((v1) | 1)), (v1 | (v1 * -837)), (v1 / (((v1 & 32767)) | 1))}
            _ = v36
        }
    case 9223372036854775807:
        var v37 uint32 = 0
        _ = v37
    default:
        v1 += v1
    }
    v38 := GenS3{F0: 0}
    v1 = (v1 << ((2) & 63))
    switch (v1 - (^(v1))) {
    case -2147483648:
        v5 = v5
    }
    if false {
        v39 := []int{}
        var v40 uint16 = 42
        var v41 float64 = 1e10
        _ = v39
        _ = v40
        _ = v41
    }
    v42 := []int{16}
    var v43 string = "\n\t"
    var v44 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
