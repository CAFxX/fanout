package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool; F1 string; F2 int }
func genFn0() uint32 {
    if false {
        var v7 string = "hello"
        _ = v7
    } else {
        var v8 int = -728
        _ = v8
    }
    var v9 int16 = -118
    _ = v9
    return 65535
}
func genFn1(p0 int8, p1 int32, p2 int64) uint64 {
    v10 := []int{}
    v11 := GenS1{F0: (^(p1)), F1: 42, F2: 127}
    var v12 int16 = 0
    if (42 < 32767) {
        if (499 < 128) {
            p2 = -1
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: (p1 | p1), F1: 2, F2: (-(fi13))}
                _ = fi13
                _ = v14
            }
            switch "\n\t" {
            case "":
                var v15 float64 = float64(1e10)
                _ = v15
            case "foo bar":
                fmt.Println(v12)
            case "hello":
                if (!((p2 > (p2 / (((p2 % ((1) | 1))) | 1))) && true)) {
                    if true {
                        v16 := GenS2{F0: v12, F1: (-(p1)), F2: 2147483648}
                        p1 += (p1 - p1)
                        _ = v16
                    } else {
                        p1 -= (p1 | (p1 - p1))
                    }
                    if true {
                        if true {
                            fmt.Println(v12, p0)
                        } else {
                            switch -9223372036854775808 {
                            case -506:
                                var v17 int = 127
                                _ = v17
                            case 32767:
                                v18 := GenS3{F0: false, F1: "a", F2: -9223372036854775808}
                                _ = v18
                            default:
                                p2 -= -1
                            }
                        }
                        var v19 uint64 = 1
                        _ = v19
                    }
                    v12 = -1
                } else {
                    if true {
                        var v20 int64 = p2
                        _ = v20
                    }
                    if ((true || (42 == 210)) && true) {
                        fmt.Println(p1, p2)
                    } else {
                        p1 -= p1
                        v21 := GenS3{F0: ((p2 + p2) < (p2 ^ 0)), F1: "a", F2: 9223372036854775807}
                        for fi22 := 0; fi22 < 2; fi22++ {
                            var v23 uint8 = 0
                            v12 += v12
                            p1 = (p1 << ((65535) & 63))
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                }
            }
        }
        if (128 == 30) {
            p0 += (p0 >> ((90) & 63))
            v10 = append(v10, 571)
            var v24 string = "a"
            _ = v24
        }
    } else {
        var v25 uint32 = 0
        _ = v25
    }
    switch -466 {
    case -908:
        if (false && true) {
            if (638 != 65535) {
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS2{F0: (-(v12)), F1: (p1 % ((0) | 1)), F2: genFn0()}
                    var v28 uint8 = 54
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v12 = (v12 - v12)
                        _ = fi29
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v30 int32 = -112
                _ = v30
            } else {
                v31 := GenS1{F0: (p1 / (((p1 - -1)) | 1)), F1: 2147483647, F2: 32768}
                _ = v31
            }
        }
    case 127:
        if len(v10) > 0 {
            v10[0] = -1
        }
    case 787:
        p0 = p0
    default:
        p2 -= (p2 << ((255) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 int = -546
    var v6 uint16 = 0
    for fi32 := 0; fi32 < 13; fi32++ {
        if (3.14159 >= 0.1) {
            switch -540 {
            case 927:
                v5 = (fi32 << ((42) & 63))
            default:
                v33 := []int{v5, (fi32 | 68), v5, 1}
                _ = v33
            }
        }
        _ = fi32
    }
    var v34 uint8 = 0
    v35 := []int{(-(v5)), (v5 & v5), (v5 * (v5 / ((v5) | 1)))}
    var v36 uint64 = 0
    v37 := GenS1{F0: 0, F1: 65535, F2: (v5 >> ((4294967295) & 63))}
    v38 := GenS2{F0: -1, F1: -1, F2: 1}
    v39 := GenS3{F0: v1, F1: ("\n\t" + "hello"), F2: (v5 % ((v5) | 1))}
    var v40 int64 = 0
    v41 := v2["a"]
    switch "\n\t" {
    case "\n\t":
        var v42 int8 = 100
        _ = v42
    default:
        delete(v4, ("a" + ("\n\t" + "a")))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
