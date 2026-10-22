package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int, p1 int8) uint {
    v5 := []int{(p0 / ((-9223372036854775808) | 1)), p0}
    p1 = (p1 - p1)
    v5 = append(v5, (p0 << ((42) & 63)))
    var v6 float64 = float64((-2.25 / 1e10))
    p1 -= (p1 - (p1 + -15))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint8, p1 int8, p2 int16) int {
    for fi7 := 0; fi7 < 1; fi7++ {
        p2 = 1000
        if true {
            switch "\n\t" {
            case "\n\t":
                for fi8 := 0; fi8 < 4; fi8++ {
                    fmt.Println(p0, fi7)
                    var v9 uint8 = (p0 >> ((2) & 63))
                    fmt.Println(p1, fi7)
                    var v10 uint = 19
                    switch 9223372036854775807 {
                    case 1:
                        v11 := GenS1{F0: (-(p2))}
                        _ = v11
                    case 32767:
                        var v12 uint32 = 65535
                        _ = v12
                    default:
                        var v13 bool = ((0.0 - (1e10 * 1.5)) <= 0.1)
                        _ = v13
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "":
                fmt.Println(p1)
            }
        } else {
            var v14 int32 = -1
            fmt.Println(v14, p0)
            if false {
                if (int32(p2) != (v14 << ((2) & 63))) {
                    var v15 float32 = (0.0 / (0.1 * -2.25))
                    _ = v15
                } else {
                    if true {
                        if (!false) {
                            var v16 uint16 = 1
                            _ = v16
                        }
                        switch fi7 {
                        case 32768:
                            p0 = 178
                        }
                    } else {
                        var v17 float64 = 0.1
                        v14 += (v14 + v14)
                        _ = v17
                    }
                    var v18 int8 = (p1 * (p1 << ((4294967295) & 63)))
                    p0 = 220
                    _ = v18
                }
                v14 = v14
                if (((p1 / ((p1) | 1)) >= p1) || (false || false)) {
                    v19 := []int{}
                    if true {
                        var v20 int8 = -128
                        _ = v20
                    } else {
                        var v21 uint16 = 1
                        if (!(((0.0 + 1.5) - 0.0) == (3.14159 + 0.0))) {
                            var v22 int = fi7
                            _ = v22
                        } else {
                            if false {
                                v14 += v14
                            }
                            fmt.Println(p0, v14, v21)
                        }
                        v23 := GenS1{F0: (p2 * 275)}
                        _ = v21
                        _ = v23
                    }
                    _ = v19
                } else {
                    if ((1e10 * (((-2.25 - -2.25) + float64((0.0 * 1.5))) * 0.1)) <= 0.1) {
                        var v24 int = fi7
                        v25 := GenS2{F0: true, F1: 1}
                        v26 := []int{(v24 | 246), (fi7 ^ int(v14)), 1, fi7}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(p1)
                    p0 += p0
                }
            }
            _ = v14
        }
        v27 := GenS2{F0: false, F1: 216}
        _ = fi7
        _ = v27
    }
    v28 := []int{-639, 9223372036854775807, 2147483647, 32768}
    v29 := GenS3{F0: (p2 - p2), F1: (p1 * p1), F2: ((("foo bar" + "a") + "") + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return 32768
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 128
    v4 := make(map[uint32]bool)
    v30 := GenS3{F0: -1, F1: 0, F2: "a"}
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 float32 = (((float32(((1e10 * (-2.25 + (1e10 / -2.25))) - 3.14159)) / 1e10) + 1e10) + (3.14159 + (0.0 * 1.5)))
        var v33 uint16 = 751
        v34 := v2[4294967295]
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 uint32 = 42
    var v36 bool = false
    fmt.Println(v35)
    var v37 uint = 1
    if true {
        v3 -= v3
    }
    v37 = (v37 >> (((v37 ^ v37)) & 63))
    v38 := GenS4{F0: 855}
    v39 := v2[v35]
    v40 := []int{128, 733}
    var v41 int16 = 1
    fmt.Println(v41)
    v42 := GenS2{F0: ((uint16(v37) <= 65535) || (true || (!v36))), F1: v37}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
