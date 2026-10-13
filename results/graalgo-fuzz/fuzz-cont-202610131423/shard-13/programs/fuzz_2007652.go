package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0() int16 {
    var v7 int64 = -1
    v7 = v7
    if false {
        v7 = v7
    }
    v7 = -1
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{fi8, (-(fi8)), fi8, fi8}
        v7 += (^(v7))
        _ = fi8
        _ = v9
    }
    _ = v7
    return 32767
}

func main() {
    v1 := []int{32767, 32768, -1, 42}
    var v2 int16 = 1000
    var v3 string = "\n\t"
    var v4 uint64 = 42
    v5 := []int{0, 127, 641}
    var v6 int64 = 0
    if len(v5) > 0 {
        v5[0] = 32768
    }
    v2 -= -358
    var v10 float64 = (1e10 * 0.0)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: (^(v2)), F1: 0}
        var v13 uint32 = 42
        var v14 float32 = 1e10
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v3 = (v3 + v3)
    v15 := GenS1{F0: genFn0(), F1: 218}
    if false {
        var v16 float32 = 0.0
        v16 -= ((v16 / ((v16 / (((v16 / (v16 * (0.0 + (v16 / float32(1e10))))) - (0.1 - 0.0)) + 1e10)) - v16)) - v16)
        _ = v16
    }
    v5 = append(v5, 32767)
    if (((v2 + 1) < (v2 * genFn0())) || false) {
        v17 := GenS1{F0: v2, F1: 2}
        _ = v17
    } else {
        v18 := GenS1{F0: genFn0(), F1: 72}
        _ = v18
    }
    fmt.Println(v6, v2)
    var v19 uint = 18446744073709551615
    if true {
        if len(v1) > 0 {
            v1[0] = -1
        }
    }
    switch ((v3 + "a") + v3) {
    case "\n\t":
        fmt.Println(v3)
    case "foo bar":
        var v20 float64 = v10
        _ = v20
    case "a":
        v2 += v2
    default:
        switch "foo bar" {
        case "foo bar":
            switch (v3 + v3) {
            case "foo bar":
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(v6, v10)
                    if ((^(v19)) < 18446744073709551615) {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            if true {
                                v19 += (v19 ^ v19)
                                v1 = append(v1, -2147483648)
                                fmt.Println(v3)
                            }
                            if false {
                                fmt.Println(fi21)
                            }
                            _ = fi22
                        }
                        v23 := GenS1{F0: -574, F1: 42}
                        _ = v23
                    } else {
                        var v24 float64 = v10
                        v10 -= 1.5
                        _ = v24
                    }
                    if ((("" <= v3) || true) || true) {
                        var v25 uint16 = 88
                        v26 := GenS1{F0: (v2 % (((v2 & (v2 | genFn0()))) | 1)), F1: 7}
                        _ = v25
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v28 := GenS2{F0: 65535, F1: 976}
                            if (2147483647 != 1) {
                                if ((v19 + v19) <= v19) {
                                    _ = len(v5)
                                }
                                switch v3 {
                                case "":
                                    if ((v19 & (v19 >> ((147) & 63))) >= (v19 << (((^(v19))) & 63))) {
                                        v3 = v3
                                        var v29 int64 = v6
                                        _ = v29
                                    }
                                case "hello":
                                    fmt.Println(fi21)
                                case "a":
                                    if false {
                                        if true {
                                            var v30 uint16 = 65535
                                            v3 = ""
                                            v31 := GenS3{F0: 9223372036854775807, F1: 4294967295, F2: v30}
                                            _ = v30
                                            _ = v31
                                        }
                                        if true {
                                            var v32 int64 = (v6 % ((int64(fi21)) | 1))
                                            var v33 uint16 = 65535
                                            v10 -= -2.25
                                            v19 = 402
                                            _ = v32
                                            _ = v33
                                        }
                                        fmt.Println(v19, v3)
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            v19 = (v19 ^ v19)
                                        case "a":
                                            if len(v1) > 0 {
                                                v1[0] = (^(fi27))
                                            }
                                        case "foo bar":
                                            switch "a" {
                                            case "hello":
                                                var v34 int32 = -1
                                                if true {
                                                    var v35 float32 = 0.1
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v34)
                                                    var v36 int8 = 127
                                                    _ = len(v1)
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v36
                                                }
                                                _ = v34
                                            }
                                        default:
                                            switch (v3 + "") {
                                            case "a":
                                                if len(v5) > 0 {
                                                    v5[0] = -2147483648
                                                }
                                            case "":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        }
                                        v37 := GenS1{F0: v2, F1: 42}
                                        _ = v37
                                    }
                                }
                                v3 = v3
                                _ = len(v5)
                            }
                            _ = len(v5)
                            _ = fi27
                            _ = v28
                        }
                        v38 := GenS2{F0: 858, F1: 936}
                        v4 -= 1
                        _ = v38
                    }
                    _ = fi21
                }
            case "hello":
                v6 = (v6 - -1)
            case "\n\t":
                var v39 bool = (true || (42 != v4))
                _ = v39
            default:
                if true {
                    v4 += 42
                    var v40 int8 = 1
                    _ = v40
                } else {
                    var v41 int32 = -375
                    _ = v41
                }
            }
        default:
            var v42 int64 = v6
            _ = v42
        }
    }
    var v43 uint16 = 42
    var v44 int16 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
