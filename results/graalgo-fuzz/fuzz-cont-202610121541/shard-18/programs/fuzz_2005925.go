package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0() bool {
    var v6 int32 = 2147483647
    var v7 int16 = 0
    v8 := GenS1{F0: -9223372036854775808, F1: 42, F2: 2}
    _ = v6
    _ = v7
    _ = v8
    return ((128 != -737) && (2147483648 <= 65535))
}
func genFn1(p0 uint64, p1 bool) uint8 {
    for fi9 := 0; fi9 < 2; fi9++ {
        p0 += 664
        var v10 string = (("" + "hello") + "foo bar")
        _ = fi9
        _ = v10
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 80
}

func main() {
    var v1 int32 = 458
    var v2 int = -2147483648
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    var v11 int64 = -9223372036854775808
    switch "hello" {
    case "hello":
        v12 := []int{v2}
        _ = v12
    case "":
        v2 -= v2
    case "foo bar":
        fmt.Println(v2, v11, v2)
    default:
        var v13 float64 = 1.5
        _ = v13
    }
    var v14 uint64 = 0
    if false {
        if false {
            v15 := []int{v2, v2, v2}
            _ = v15
        } else {
            v16 := GenS2{F0: (v14 ^ v14), F1: 546, F2: 1}
            _ = v16
        }
        v17 := GenS1{F0: 488, F1: v11, F2: 128}
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: 42, F1: v11, F2: 1}
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    v20 := []int{v2, v2, (v2 / ((1) | 1)), v2}
    for fi21 := 0; fi21 < 14; fi21++ {
        fmt.Println(v1, v14, fi21)
        if true {
            for fi22 := 0; fi22 < 17; fi22++ {
                if false {
                    v14 = v14
                    if (0 >= 65535) {
                        if true {
                            var v23 uint = 1
                            _ = v23
                        }
                    } else {
                        v24 := GenS2{F0: v14, F1: 163, F2: 0}
                        _ = v24
                    }
                } else {
                    var v25 uint8 = 0
                    _ = v25
                }
                _ = fi22
            }
            v26 := GenS1{F0: v11, F1: -9223372036854775808, F2: 169}
            if (42 <= 1000) {
                switch (fi21 / (((fi21 - fi21)) | 1)) {
                case 128:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        if (!true) {
                            if true {
                                if true {
                                    _ = len(v5)
                                    v28 := GenS3{F0: 128, F1: 42}
                                    var v29 string = "foo bar"
                                    _ = v28
                                    _ = v29
                                } else {
                                    if len(v20) > 0 {
                                        v20[0] = 0
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v30 uint8 = 42
                                    _ = v30
                                default:
                                    v31 := GenS1{F0: -1, F1: (^(v11)), F2: uint8(v14)}
                                    v14 = 18446744073709551615
                                    v11 = (v11 + v11)
                                    _ = v31
                                }
                            } else {
                                v14 -= 1
                                var v32 uint16 = 0
                                fmt.Println(v1)
                                _ = v32
                            }
                            _ = len(v20)
                        }
                        v20 = append(v20, 32768)
                        _ = fi27
                    }
                case 899:
                    v1 = int32(v14)
                case -9223372036854775808:
                    var v33 int32 = v1
                    _ = v33
                }
            }
            _ = v26
        } else {
            v34 := v5[2147483648]
            var v35 uint16 = 0
            _ = v34
            _ = v35
        }
        if (((!(0.1 >= 0.0)) || (!true)) && true) {
            v2 -= 32767
        }
        _ = fi21
    }
    if false {
        for fi36 := 0; fi36 < 4; fi36++ {
            v11 -= (v11 & v11)
            v1 -= v1
            _ = fi36
        }
    }
    if (!true) {
        delete(v4, 2)
        for fi37 := 0; fi37 < 1; fi37++ {
            fmt.Println(v2)
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "foo bar":
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v39 := GenS3{F0: 0, F1: 42}
                        var v40 uint16 = 255
                        if len(v20) > 0 {
                            v20[0] = (fi37 & 2147483647)
                        }
                        if true {
                            switch (fi37 ^ -2147483648) {
                            case 877:
                                var v41 int8 = 1
                                _ = v41
                            }
                        }
                        _ = len(v20)
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                default:
                    var v42 int64 = 42
                    _ = v42
                }
            case "a":
                if (!(false && (!(!false)))) {
                    v1 = v1
                    v20 = append(v20, (v2 - fi37))
                    var v43 uint16 = 65535
                    _ = v43
                } else {
                    v14 = v14
                }
            }
            _ = fi37
        }
    }
    var v44 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
