package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 int8) uint {
    v7 := []int{128, 32768, -38}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: true}
        var v10 uint64 = 409
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v7 = append(v7, fi11)
        if (!(((-(p0)) >= p0) || true)) {
            v12 := GenS1{F0: (207 != 32767)}
            _ = v12
        } else {
            p0 += 90
            if false {
                _ = len(v7)
                fmt.Println(fi11)
                p0 = 114
            } else {
                p0 += (^(p0))
            }
            v13 := GenS2{F0: 2, F1: 9223372036854775807, F2: 307}
            _ = v13
        }
        if false {
            if (0 == -2147483648) {
                if (871 < -763) {
                    v14 := GenS3{F0: 120, F1: 42}
                    p0 -= p0
                    var v15 int16 = -32768
                    _ = v14
                    _ = v15
                } else {
                    if len(v7) > 0 {
                        v7[0] = (fi11 & (fi11 ^ fi11))
                    }
                    p0 = p0
                }
            }
        }
        _ = fi11
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (float32(1e10) * (1.5 / (1e10 - ((1.5 - 0.0) * (((3.14159 * ((1e10 / -2.25) * (-2.25 * 3.14159))) + -2.25) - float32(float64(((1e10 + ((0.1 * float32(-2.25)) + 0.0)) + 3.14159))))))))
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    var v5 float32 = (v2 * (((0.1 + float32(1e10)) * 3.14159) / float32(0.1)))
    v6 := []string{"hello"}
    v16 := GenS4{F0: -32768, F1: -1, F2: 0}
    v17 := v1[2147483647]
    if (!(723 != 42)) {
        var v18 uint16 = 942
        v5 = 0.1
        _ = v18
    }
    fmt.Println(v5, v17, v17)
    if true {
        if (2 > genFn0(94)) {
            var v19 int8 = -11
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS1{F0: (42 <= -32768)}
            v17 = v17
            _ = fi20
            _ = v21
        }
    } else {
        v17 = v17
        fmt.Println(v5, v5)
        v5 += v2
    }
    v2 = (v5 * ((v5 - v2) - ((0.0 / v5) * v2)))
    if ((!true) || false) {
        var v22 uint32 = 116
        _ = v22
    } else {
        if false {
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 bool = (true && (-578 <= -2147483648))
                if v24 {
                    v2 = ((float32(float64(v5)) * 0.1) + 3.14159)
                } else {
                    v25 := GenS5{F0: 2}
                    v17 = v17
                    var v26 string = (v17 + (((("\n\t" + v17) + v17) + (v17 + ("foo bar" + ((v17 + v17) + v17)))) + (v17 + (v17 + "hello"))))
                    _ = v25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            var v27 string = v17
            _ = v27
        } else {
            var v28 bool = false
            if true {
                v29 := GenS6{F0: -1}
                v30 := GenS5{F0: 755}
                _ = v29
                _ = v30
            }
            v31 := GenS3{F0: 0, F1: 42}
            _ = v28
            _ = v31
        }
        if (2 != 2) {
            v17 = v17
            for fi32 := 0; fi32 < 4; fi32++ {
                v5 = float32((1e10 / 1e10))
                _ = fi32
            }
            v2 += (-2.25 - float32((1.5 + float64(v5))))
        }
        var v33 uint8 = 0
        _ = v33
    }
    fmt.Println(v2, v5)
    v34 := GenS1{F0: (82 > 0)}
    switch 32768 {
    case 1:
        var v35 uint8 = 127
        _ = v35
    default:
        v17 = v17
    }
    var v36 uint8 = 2
    var v37 uint = 135
    v37 = 2
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            if true {
                var v40 uint32 = 0
                v5 -= 0.1
                _ = v40
            } else {
                var v41 float64 = 0.1
                _ = v41
            }
            if false {
                if (true && false) {
                    v42 := GenS3{F0: 0, F1: 255}
                    fmt.Println(fi39, v36, v17)
                    if (!(2 != 901)) {
                        if (!true) {
                            v6 = append(v6, ((((v17 + ((v17 + ((v17 + v17) + "\n\t")) + ("a" + "hello"))) + v17) + ("hello" + ((v17 + v17) + (v17 + ((v17 + "hello") + v17))))) + (v17 + ("foo bar" + v17))))
                            v43 := GenS3{F0: 520, F1: uint16(v36)}
                            _ = v43
                        }
                    }
                    _ = v42
                }
                var v44 uint32 = 673
                _ = v44
            }
            v45 := GenS7{F0: 178, F1: 117}
            _ = fi39
            _ = v45
        }
        var v46 string = (v17 + (("hello" + "a") + (v17 + ("a" + "foo bar"))))
        _ = fi38
        _ = v46
    }
    var v47 int32 = -89
    if (1.5 <= (1e10 * 0.0)) {
        fmt.Println(v17)
        if ((-2.25 + (float32(1e10) - v2)) >= 3.14159) {
            switch -466 {
            case 42:
                for fi48 := 0; fi48 < 14; fi48++ {
                    fmt.Println(v37, v2)
                    _ = fi48
                }
            case 9223372036854775807:
                v49 := GenS8{F0: ("" + v17), F1: 212}
                _ = v49
            }
        }
    }
    var v50 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
