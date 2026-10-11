package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int8 = 1
    var v4 int64 = -9223372036854775808
    var v5 uint = 65535
    fmt.Println(v5, v3, v5)
    v6 := GenS1{F0: 9223372036854775807, F1: 430}
    if ((2 >= 2) && (-1 < 32768)) {
        fmt.Println(v3)
        v5 += 18446744073709551615
    } else {
        v7 := GenS1{F0: 32767, F1: 18446744073709551615}
        _ = v7
    }
    v5 -= v5
    var v8 float64 = float64(-2.25)
    var v9 uint16 = 42
    var v10 int = 32768
    switch (^(v10)) {
    case 127:
        v3 -= -53
    case 32767:
        v3 = (v3 % (((v3 / ((v3) | 1))) | 1))
    default:
        if ((255 <= 32767) && (true || ((true || true) || ((v4 - (v4 / (((v4 / ((int64(v10)) | 1))) | 1))) <= v4)))) {
            var v11 uint64 = 18446744073709551615
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v12 uint = v5
                    _ = v12
                case "":
                    if ("hello" < "\n\t") {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 int32 = -2147483648
                            _ = fi13
                            _ = v14
                        }
                        v15 := []int{(v10 | v10), (v10 % ((v10) | 1)), 32767}
                        _ = v15
                    } else {
                        var v16 int16 = -399
                        _ = v16
                    }
                }
                var v17 uint32 = 0
                _ = v17
            } else {
                delete(v1, 2)
            }
            v18 := GenS1{F0: 426, F1: 2}
            _ = v11
            _ = v18
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 float64 = (float64((3.14159 * 3.14159)) + (v8 / 0.1))
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                for fi22 := 0; fi22 < 19; fi22++ {
                    if (false || false) {
                        fmt.Println(fi21)
                        switch v10 {
                        case 127:
                            if (true || true) {
                                fmt.Println(v5, fi21, v4)
                            } else {
                                var v23 float32 = 1e10
                                v5 += 2
                                switch "foo bar" {
                                case "a":
                                    v24 := GenS2{F0: (v5 & (v5 & 2))}
                                    _ = v24
                                default:
                                    fmt.Println(v3)
                                }
                                _ = v23
                            }
                        case 9223372036854775807:
                            if false {
                                _ = len(v2)
                            } else {
                                var v25 int = (fi22 | 811)
                                var v26 int32 = 2147483647
                                v25 = (fi21 & v25)
                                _ = v25
                                _ = v26
                            }
                            if true {
                                if true {
                                    v3 += v3
                                    v9 += (v9 | v9)
                                    v27 := v2["a"]
                                    _ = v27
                                }
                                var v28 uint8 = 128
                                v29 := []int{v10, int(v28), int(v28), -855}
                                switch "\n\t" {
                                case "hello":
                                    var v30 uint = v5
                                    _ = v30
                                case "a":
                                    var v31 int8 = v3
                                    _ = v31
                                case "foo bar":
                                    var v32 int32 = -1
                                    _ = v32
                                }
                                var v33 string = ("foo bar" + "")
                                delete(v2, v33)
                                _ = v28
                                _ = v29
                                _ = v33
                            } else {
                                v34 := GenS2{F0: (v5 | (v5 & (^(v5))))}
                                fmt.Println(fi22, v3)
                                _ = v34
                            }
                            fmt.Println(fi22, v8, v3)
                        case 2147483647:
                            var v35 uint32 = 291
                            _ = v35
                        }
                        var v36 int = -51
                        _ = v36
                    }
                    _ = fi22
                }
                var v37 uint16 = v9
                _ = fi21
                _ = v37
            }
            v38 := GenS2{F0: v5}
            _ = v38
        }
    }
    if (false || (1e10 >= (((v8 / v8) * ((1.5 + v8) + ((((v8 + 1e10) / 0.1) / 1.5) * (v8 + ((v8 / v8) * 1.5))))) - (v8 + float64(0.1))))) {
        if ((((((("a" + ("hello" + (("a" + "foo bar") + ("foo bar" + ("hello" + "foo bar"))))) + "foo bar") + ("hello" + "hello")) + "a") + "") + "") > ("" + "foo bar")) {
            var v39 int16 = 1
            var v40 int8 = v3
            v41 := GenS3{F0: v9}
            _ = v39
            _ = v40
            _ = v41
        } else {
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 float64 = 0.1
                if ((!true) || true) {
                    var v44 string = "\n\t"
                    v45 := GenS4{F0: v3, F1: (^(v3))}
                    _ = v44
                    _ = v45
                } else {
                    v46 := []int{int(v9), int(v5), -466, (fi42 | 484)}
                    var v47 int64 = -9223372036854775808
                    _ = v46
                    _ = v47
                }
                _ = fi42
                _ = v43
            }
            _ = len(v1)
        }
        fmt.Println(v10, v3)
    }
    v3 -= (v3 << (((v5 ^ v5)) & 63))
    v4 = (v4 / ((-9223372036854775808) | 1))
    v48 := []int{(v10 & (v10 - (v10 + 32767))), (^(v10))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
