package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint16 }
type GenS5 struct { F0 uint; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16 }
func genFn0(p0 bool, p1 int8, p2 string) bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 2, F1: -160}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            p0 = (1000 != -1)
            v11 := []int{(-(fi7)), 2147483647, (fi7 >> ((145) & 63)), (fi7 / (((fi7 * fi10)) | 1))}
            switch -300 {
            case 1:
                v12 := GenS2{F0: 2147483648}
                _ = v12
            case 4294967295:
                switch "foo bar" {
                case "hello":
                    if len(v11) > 0 {
                        v11[0] = fi10
                    }
                case "a":
                    _ = len(v11)
                case "":
                    if ("hello" >= p2) {
                        if (42 == 326) {
                            p1 += 92
                            if (1 > 0) {
                                if (!p0) {
                                    v13 := GenS3{F0: 2147483648, F1: p1, F2: 1}
                                    _ = v13
                                }
                                var v14 int32 = 2147483647
                                v15 := GenS3{F0: 1, F1: (p1 % (((p1 ^ (p1 ^ p1))) | 1)), F2: 186}
                                var v16 int8 = 1
                                switch -550 {
                                case 1:
                                    p0 = p0
                                default:
                                    switch 0 {
                                    case 4294967295:
                                        _ = len(v11)
                                    case -504:
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    case 2147483647:
                                        if (32 != 42) {
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        } else {
                                            fmt.Println(p0)
                                        }
                                    default:
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                    if true {
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    }
                                    var v18 int16 = -969
                                    _ = v18
                                }
                                var v19 uint32 = 959
                                _ = v14
                                _ = v15
                                _ = v16
                                _ = v19
                            } else {
                                v20 := GenS4{F0: 0, F1: ((p2 + p2) + ("hello" + p2)), F2: 1}
                                var v21 string = p2
                                p0 = p0
                                _ = v20
                                _ = v21
                            }
                            p1 -= p1
                        } else {
                            fmt.Println(fi10, p2)
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            fmt.Println(fi7, p1)
                        }
                    }
                default:
                    v22 := GenS4{F0: 266, F1: ("" + p2), F2: 831}
                    _ = v22
                }
            case -790:
                switch fi10 {
                case 42:
                    var v23 int16 = 1000
                    _ = v23
                case -1:
                    var v24 float64 = 1.5
                    v11 = append(v11, 32767)
                    _ = v24
                }
            }
            _ = fi10
            _ = v11
        }
        var v25 bool = (0.0 == 3.14159)
        _ = fi7
        _ = v25
    }
    var v26 float32 = (-2.25 / 1.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return (p0 || p0)
}

func main() {
    var v1 uint64 = 690
    var v2 uint64 = 621
    var v3 uint16 = 42
    var v4 int8 = 42
    var v5 int16 = 763
    v6 := make(map[int]string)
    for fi27 := 0; fi27 < 17; fi27++ {
        if ((42 != 65535) && (true || ((-(fi27)) < fi27))) {
            delete(v6, fi27)
            var v28 float64 = (float64(1.5) + (((3.14159 - 1e10) / 1.0) * 0.1))
            var v29 bool = (false && true)
            _ = v28
            _ = v29
        }
        switch ("foo bar" + "a") {
        case "foo bar":
            v1 = v1
        default:
            v30 := []int{(fi27 * fi27)}
            _ = v30
        }
        _ = fi27
    }
    var v31 int8 = (v4 | (v4 | (v4 - v4)))
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 string = ("\n\t" + "foo bar")
        _ = fi32
        _ = v33
    }
    v31 = (v4 & -93)
    v4 += v4
    v34 := GenS5{F0: 1, F1: (-(v5)), F2: v1}
    if (1 == 65535) {
        if (79 >= (v3 - 32767)) {
            var v35 int32 = -2147483648
            v36 := GenS6{F0: 255}
            _ = v35
            _ = v36
        } else {
            v37 := GenS4{F0: 18446744073709551615, F1: "a", F2: 65535}
            v38 := GenS7{F0: v5}
            _ = v37
            _ = v38
        }
    }
    var v39 uint16 = v3
    v40 := GenS7{F0: v5}
    var v41 uint64 = (v1 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
