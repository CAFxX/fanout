package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 uint8, p1 string) bool {
    v4 := []int{27}
    if false {
        p0 = p0
        if len(v4) > 0 {
            v4[0] = int(p0)
        }
    } else {
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 = p0
            p1 = p1
            for fi6 := 0; fi6 < 16; fi6++ {
                if (!false) {
                    v7 := GenS1{F0: true, F1: 487}
                    p1 = (p1 + ((((p1 + "\n\t") + "") + p1) + p1))
                    if false {
                        switch (p1 + p1) {
                        case "\n\t":
                            v8 := GenS2{F0: 0}
                            _ = v8
                        case "foo bar":
                            var v9 int = (fi6 | fi5)
                            _ = v9
                        case "hello":
                            fmt.Println(fi5, p1, p0)
                        default:
                            fmt.Println(fi6, fi6, fi6)
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v7
                } else {
                    p1 = "hello"
                    var v10 int = (fi6 * fi5)
                    switch ((p1 + (p1 + "\n\t")) + p1) {
                    case "":
                        if len(v4) > 0 {
                            v4[0] = fi5
                        }
                    case "hello":
                        _ = len(v4)
                    default:
                        if true {
                            var v11 int32 = 617
                            v11 = v11
                            _ = v11
                        } else {
                            if (false && false) {
                                var v12 float32 = -2.25
                                _ = v12
                            } else {
                                v10 += fi6
                                if true {
                                    if false {
                                        v10 = fi5
                                        fmt.Println(v10, fi6)
                                    } else {
                                        if true {
                                            fmt.Println(v10, v10, fi6)
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if false {
                                            v13 := GenS3{F0: fi6, F1: 214}
                                            fmt.Println(p0, p1)
                                            v10 += (fi5 / (((v10 & (fi5 | v10))) | 1))
                                            _ = v13
                                        } else {
                                            v14 := GenS2{F0: 42}
                                            _ = v14
                                        }
                                    }
                                    var v15 int = (fi6 + (v10 & fi6))
                                    var v16 int16 = -222
                                    _ = v15
                                    _ = v16
                                }
                                if (!true) {
                                    p0 += p0
                                    var v17 int64 = 42
                                    _ = v17
                                }
                            }
                        }
                    }
                    _ = v10
                }
                _ = fi6
            }
            _ = fi5
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v4 = append(v4, (fi18 + (^(fi18))))
            if len(v4) > 0 {
                v4[0] = fi18
            }
            _ = fi18
        }
        v19 := GenS1{F0: (false || (325 >= 2)), F1: 1000}
        _ = v19
    }
    _ = len(v4)
    var v20 uint64 = 478
    v20 = (^(v20))
    _ = p0
    _ = p1
    _ = v20
    _ = v4
    return (!true)
}

func main() {
    var v1 float64 = (((3.14159 + (-2.25 + 3.14159)) * (1.5 + ((-2.25 / -2.25) - 3.14159))) / (0.0 + 1e10))
    v2 := make(map[uint32]string)
    var v3 int16 = -297
    v1 = (float64(0.1) + float64(1e10))
    var v21 int = 32768
    for fi22 := 0; fi22 < 6; fi22++ {
        v2[2147483648] = "foo bar"
        v23 := []int{fi22}
        v24 := v2[2147483648]
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := []int{(v21 + (v21 / (((v21 << ((1) & 63))) | 1))), (-(v21)), 14, -2147483648}
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int64 = int64(v3)
            v3 -= v3
            v29 := GenS1{F0: false, F1: (v3 << ((4294967295) & 63))}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        v21 = (fi30 * (v21 & (fi30 * v21)))
        var v31 float64 = v1
        v31 += v31
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: (v3 % ((v3) | 1))}
    v1 = ((3.14159 - (v1 + (float64((3.14159 * 1e10)) / (v1 * v1)))) - float64(-2.25))
    var v33 bool = true
    switch v21 {
    case 315:
        if (!(18446744073709551615 > 4294967295)) {
            _ = len(v25)
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 uint32 = 2147483647
                _ = fi34
                _ = v35
            }
        } else {
            var v36 uint8 = 255
            fmt.Println(v3)
            v37 := GenS1{F0: (0 >= 42), F1: 42}
            _ = v36
            _ = v37
        }
    default:
        delete(v2, 0)
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v32
    _ = v33
}
