package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 int; F1 bool; F2 int64 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
func genFn0(p0 bool, p1 int16) uint16 {
    p1 -= 1
    v7 := GenS1{F0: p1, F1: -560, F2: 100}
    _ = p0
    _ = p1
    _ = v7
    return 85
}
func genFn1(p0 uint16, p1 bool) uint64 {
    v8 := GenS1{F0: 42, F1: 9223372036854775807, F2: 0}
    if p1 {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 395, F1: 700, F2: 42}
            _ = fi9
            _ = v10
        }
        if (((1e10 / (1.5 - (-2.25 * 0.1))) < 0.0) && (4294967295 > 2147483647)) {
            var v11 int8 = 42
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    if p1 {
                        v14 := []int{(-(fi12)), (fi12 - (^(fi12)))}
                        _ = v14
                    } else {
                        fmt.Println(fi13, fi12)
                        if (42 == 104) {
                            v15 := GenS1{F0: 32767, F1: (fi12 * (^(fi13))), F2: (v11 - v11)}
                            _ = v15
                        } else {
                            if p1 {
                                p0 = 249
                            } else {
                                v16 := GenS2{F0: fi13, F1: false, F2: -589}
                                v17 := GenS1{F0: 42, F1: 963, F2: v11}
                                switch "" {
                                case "a":
                                    fmt.Println(v11)
                                case "\n\t":
                                    var v18 float32 = -2.25
                                    _ = v18
                                case "":
                                    var v19 uint32 = 1
                                    _ = v19
                                }
                                var v20 int64 = -696
                                p0 = (^(p0))
                                var v21 uint16 = p0
                                var v22 uint16 = 255
                                _ = v16
                                _ = v17
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            if (!(-32768 > -1)) {
                                v11 += v11
                                p0 += p0
                                var v23 int64 = 42
                                _ = v23
                            } else {
                                fmt.Println(fi12, p1, v11)
                            }
                        }
                    }
                    _ = fi13
                }
                v24 := []int{fi12, fi12, (fi12 / ((fi12) | 1))}
                _ = fi12
                _ = v24
            }
            _ = v11
        } else {
            var v25 int = 959
            p1 = (!(206 == 128))
            var v26 float64 = float64((-2.25 / 1.0))
            _ = v25
            _ = v26
        }
    } else {
        if p1 {
            p0 -= (p0 | 42)
            var v27 int8 = 1
            v28 := []int{208, 497}
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: 2147483647, F1: false, F2: 0}
        v30 := GenS3{F0: 255, F1: "foo bar", F2: 2}
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v8
    return 2
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 543
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    v5 := []uint32{65535, 154, 42}
    v6 := make(map[int]string)
    v1 = ("a" + (v1 + (v1 + ("" + ""))))
    switch 127 {
    case 701:
        _ = len(v3)
    case -9223372036854775808:
        var v31 uint8 = 241
        _ = v31
    default:
        switch 863 {
        case 383:
            v32 := GenS4{F0: 18446744073709551615, F1: (v1 + "\n\t")}
            _ = v32
        case 4294967295:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v33 := GenS5{F0: true}
            _ = v33
        }
    }
    switch "hello" {
    case "":
        var v34 uint16 = v2
        _ = v34
    case "foo bar":
        var v35 uint64 = 109
        _ = v35
    case "a":
        fmt.Println(v1, v1)
    default:
        v1 = v1
    }
    switch 127 {
    case 507:
        var v36 float32 = (-2.25 - 1e10)
        _ = v36
    default:
        _ = len(v5)
    }
    v37 := GenS1{F0: 647, F1: -1, F2: 49}
    switch 32767 {
    case -322:
        v2 = (v2 * v2)
    case 1:
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 int32 = -159
            var v40 uint16 = v2
            _ = fi38
            _ = v39
            _ = v40
        }
    case -9223372036854775808:
        v2 = (v2 >> ((4294967295) & 63))
    default:
        v2 = 32767
    }
    v41 := GenS2{F0: -559, F1: false, F2: -1}
    fmt.Println(v2, v1)
    v42 := GenS6{F0: int32(v2)}
    var v43 float32 = 3.14159
    v43 += 1e10
    if false {
        var v44 int8 = 0
        var v45 uint = 18446744073709551615
        var v46 uint64 = 4294967295
        _ = v44
        _ = v45
        _ = v46
    } else {
        v43 += v43
        if len(v5) > 0 {
            v5[0] = 299
        }
        var v47 int16 = -1
        _ = v47
    }
    v48 := GenS6{F0: -2147483648}
    v1 = (v1 + ("" + (v1 + (v1 + (v1 + ((((v1 + v1) + (v1 + v1)) + "a") + (v1 + (v1 + "\n\t"))))))))
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v5
    _ = v6
}
