package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int32 = -602
    var v4 int32 = -2147483648
    var v5 int64 = 9223372036854775807
    var v6 float64 = (0.0 - 0.1)
    v5 -= (v5 & (v5 - v5))
    v7 := GenS1{F0: 2147483648, F1: 128, F2: (v5 & (v5 | v5))}
    for fi8 := 0; fi8 < 3; fi8++ {
        v2[fi8] = (4294967295 <= 0)
        v3 -= v3
        var v9 uint8 = 128
        _ = fi8
        _ = v9
    }
    if (!true) {
        var v10 float32 = 0.1
        _ = v10
    } else {
        if (!(!true)) {
            if (((v3 - (v3 >> ((1) & 63))) <= v3) && true) {
                v4 -= (v4 + 192)
                v11 := GenS1{F0: 114, F1: -2147483648, F2: 9223372036854775807}
                for fi12 := 0; fi12 < 6; fi12++ {
                    v6 = float64(1.5)
                    _ = fi12
                }
                _ = v11
            } else {
                if false {
                    v1["\n\t"] = ("a" + ((("\n\t" + "") + "foo bar") + "foo bar"))
                    for fi13 := 0; fi13 < 5; fi13++ {
                        if true {
                            if (false || true) {
                                v14 := GenS2{F0: fi13, F1: fi13}
                                _ = v14
                            }
                            v15 := GenS2{F0: -1, F1: fi13}
                            var v16 uint8 = 255
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(v4)
                        }
                        _ = fi13
                    }
                    v1[(("" + "hello") + (("\n\t" + "\n\t") + ("a" + "foo bar")))] = ((("a" + "foo bar") + (("a" + ("" + "\n\t")) + "")) + "\n\t")
                } else {
                    v17 := GenS1{F0: 42, F1: 2147483647, F2: (v5 / (((v5 * (-(v5)))) | 1))}
                    fmt.Println(v3)
                    if false {
                        switch 59 {
                        case 128:
                            for fi18 := 0; fi18 < 8; fi18++ {
                                delete(v1, "a")
                                v4 = 42
                                if true {
                                    var v19 int64 = v5
                                    _ = v19
                                }
                                _ = fi18
                            }
                        }
                        if ((447 >= 0) && true) {
                            v6 = 1.5
                            var v20 string = ("\n\t" + ("\n\t" + ""))
                            v20 = (("hello" + "a") + v20)
                            _ = v20
                        } else {
                            if false {
                                v3 = (v4 & (-(v3)))
                                if (((1e10 / 3.14159) + (((0.1 + float32(0.0)) + (1.5 / 1.5)) / 1e10)) > 0.1) {
                                    v6 = (0.1 * (-2.25 * v6))
                                    if true {
                                        for fi21 := 0; fi21 < 9; fi21++ {
                                            for fi22 := 0; fi22 < 11; fi22++ {
                                                v6 = float64(0.1)
                                                _ = fi22
                                            }
                                            _ = fi21
                                        }
                                    } else {
                                        v4 = 2147483647
                                    }
                                    v5 = -9223372036854775808
                                } else {
                                    v23 := GenS3{F0: (v4 / ((v3) | 1)), F1: -1}
                                    var v24 float64 = v6
                                    _ = v23
                                    _ = v24
                                }
                            }
                        }
                    } else {
                        if (2 <= 298) {
                            switch 1 {
                            case -2147483648:
                                var v25 uint64 = 2
                                _ = v25
                            default:
                                var v26 int64 = 42
                                _ = v26
                            }
                            v3 = v4
                            v3 -= (-(v4))
                        } else {
                            v27 := GenS4{F0: 2147483647}
                            if true {
                                v28 := []int{128, 537}
                                _ = v28
                            } else {
                                var v29 int8 = -24
                                v30 := []int{-15}
                                _ = v29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 19; fi31++ {
                                if true {
                                    v32 := GenS4{F0: v4}
                                    _ = v32
                                }
                                fmt.Println(v3)
                                _ = fi31
                            }
                            _ = v27
                        }
                    }
                    _ = v17
                }
            }
            v5 += v5
        }
    }
    var v33 uint32 = 1
    if (((18446744073709551615 >= 0) && (false && (4294967295 < 466))) || (9223372036854775807 >= -971)) {
        if (true && true) {
            v34 := []int{-9223372036854775808, 4294967295, -9223372036854775808, 32767}
            fmt.Println(v5, v4)
            v4 = (v3 | 2147483647)
            _ = v34
        } else {
            for fi35 := 0; fi35 < 17; fi35++ {
                fmt.Println(v3)
                v36 := []int{}
                _ = fi35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            v4 += (v3 * (v4 ^ v4))
            v38 := v1[""]
            if ((^(v5)) == (v5 ^ (v5 | v5))) {
                v39 := GenS5{F0: 1, F1: 95, F2: v5}
                _ = v39
            } else {
                v40 := []int{}
                fmt.Println(v4, v38)
                _ = v40
            }
            _ = fi37
            _ = v38
        }
        v41 := []int{9223372036854775807, 32768, -702}
        _ = v41
    } else {
        var v42 int = int(v5)
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        if true {
            v33 = v33
            var v44 uint = 42
            _ = v44
        } else {
            v45 := []int{375, (fi43 | (fi43 - fi43)), fi43}
            _ = v45
        }
        _ = fi43
    }
    var v46 uint32 = 42
    v46 -= 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
