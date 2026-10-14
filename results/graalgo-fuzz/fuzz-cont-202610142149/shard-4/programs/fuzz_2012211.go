package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 int8 = -128
    if ((true && (true && true)) || true) {
        if (65535 > 42) {
            var v6 int32 = -37
            var v7 string = (("hello" + "a") + ("\n\t" + "hello"))
            var v8 int32 = v6
            _ = v6
            _ = v7
            _ = v8
        } else {
            v4 = v4
        }
        var v9 bool = true
        switch -61 {
        case 1:
            v10 := []int{128, 1}
            _ = v10
        default:
            var v11 uint16 = 578
            _ = v11
        }
        _ = v9
    } else {
        v5 -= v5
        fmt.Println(v5)
    }
    v12 := GenS1{F0: 42, F1: 1000}
    delete(v3, 0)
    var v13 uint32 = 178
    for fi14 := 0; fi14 < 12; fi14++ {
        if (fi14 >= fi14) {
            if (!((0.1 != 3.14159) || (v5 == 1))) {
                var v15 int64 = 9223372036854775807
                _ = v15
            } else {
                var v16 float32 = (float32((3.14159 / (0.1 + 1e10))) / -2.25)
                v17 := GenS2{F0: ((("a" + "hello") + (("" + "") + ("a" + ""))) + ("foo bar" + "")), F1: v5, F2: (!(v13 > 425))}
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS2{F0: "hello", F1: v5, F2: true}
                    v20 := []int{(fi18 ^ (fi18 | (-(fi14))))}
                    v5 = (-(v5))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v16
                _ = v17
            }
            switch "" {
            case "a":
                var v21 uint32 = (v13 | (v13 | (v13 * 527)))
                _ = v21
            case "hello":
                v22 := GenS3{F0: 0}
                _ = v22
            default:
                if ((-2.25 * (0.1 * 1e10)) == float64(0.0)) {
                    v23 := GenS4{F0: ((((("\n\t" + "") + ("\n\t" + ("\n\t" + "a"))) + "a") >= ("\n\t" + "\n\t")) || (255 != 513)), F1: (!false)}
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := GenS4{F0: (!true), F1: false}
                        var v26 uint32 = v13
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    if true {
                        v27 := GenS5{F0: true}
                        if (!true) {
                            v5 -= (v5 ^ v5)
                        }
                        if (!false) {
                            v5 -= (v5 << ((255) & 63))
                            v4 = 669
                        } else {
                            if true {
                                v13 += 0
                                v1["hello"] = "foo bar"
                            } else {
                                v28 := GenS3{F0: 0}
                                v4 -= (v4 >> ((4294967295) & 63))
                                _ = v28
                            }
                            if (2 >= 18446744073709551615) {
                                v29 := GenS6{F0: v4}
                                fmt.Println(fi14, v5)
                                v30 := GenS4{F0: (2147483647 <= 339), F1: (false && true)}
                                _ = v29
                                _ = v30
                            }
                        }
                        _ = v27
                    }
                    _ = v23
                } else {
                    var v31 int32 = -2147483648
                    _ = v31
                }
            }
            v32 := GenS1{F0: 595, F1: -676}
            _ = v32
        }
        v33 := GenS5{F0: (!(709 < -1))}
        _ = fi14
        _ = v33
    }
    var v34 int8 = 42
    var v35 uint8 = 42
    v36 := []int{689}
    v37 := v3[v13]
    v1[("hello" + ("a" + "\n\t"))] = (("hello" + "foo bar") + ("a" + "hello"))
    v38 := GenS7{F0: 1, F1: 1}
    if (true || (v13 != (v13 ^ v13))) {
        fmt.Println(v35, v13, v34)
    } else {
        if (v4 == v4) {
            var v39 uint = uint(v13)
            v34 = (v5 ^ (v5 ^ (v5 ^ 0)))
            var v40 uint64 = 1
            _ = v39
            _ = v40
        } else {
            var v41 float64 = 0.1
            var v42 bool = false
            _ = v41
            _ = v42
        }
    }
    switch "hello" {
    case "hello":
        _ = len(v36)
    default:
        v37 = true
    }
    var v43 int32 = -697
    for fi44 := 0; fi44 < 9; fi44++ {
        for fi45 := 0; fi45 < 13; fi45++ {
            v34 -= (v5 * (v5 | (v5 * v34)))
            _ = fi45
        }
        fmt.Println(v34, v13, v5)
        var v46 float64 = (((float64((-2.25 - 0.1)) * 1.5) * 1e10) - (1.5 * ((float64(float32(0.1)) / (((float64(float32(3.14159)) + (3.14159 / float64((float32(1e10) * 3.14159)))) / (3.14159 * 1e10)) * (0.1 - 1e10))) - 1e10)))
        _ = fi44
        _ = v46
    }
    if len(v36) > 0 {
        v36[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
}
