package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
func genFn0(p0 uint64, p1 uint64) int32 {
    p1 -= p0
    var v7 bool = (534 <= -9223372036854775808)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 0
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 bool) int32 {
    if true {
        v10 := GenS1{F0: -184}
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: "\n\t", F1: -1}
            v13 := GenS1{F0: -403}
            if (((p0 || p0) || p0) || p0) {
                p0 = p0
                v14 := GenS1{F0: (fi11 << ((234) & 63))}
                for fi15 := 0; fi15 < 20; fi15++ {
                    p0 = p0
                    _ = fi15
                }
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        v16 := []int{127}
        switch 32768 {
        case 57:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        case -9223372036854775808:
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 float32 = 1e10
                if len(v16) > 0 {
                    v16[0] = -1
                }
                _ = fi17
                _ = v18
            }
        }
        if len(v16) > 0 {
            v16[0] = 446
        }
        _ = v16
    }
    v19 := []int{}
    p0 = p0
    v20 := GenS3{F0: false}
    _ = p0
    _ = v19
    _ = v20
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 87
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    var v6 int = -947
    if (!((true || ((42 == 9223372036854775807) || ((false && false) && (!(!(-315 <= 9223372036854775807)))))) && (-814 >= -763))) {
        v21 := []int{v6}
        v22 := GenS4{F0: 65535, F1: 1}
        var v23 uint = 2
        _ = v21
        _ = v22
        _ = v23
    } else {
        if true {
            if false {
                _ = len(v5)
            } else {
                fmt.Println(v2, v2)
                switch (("a" + "a") + "") {
                case "hello":
                    if ((0.0 + (float64(1e10) * float64(-2.25))) == 0.0) {
                        v24 := GenS5{F0: 42}
                        v2 = 0
                        var v25 bool = (!true)
                        _ = v24
                        _ = v25
                    }
                case "a":
                    v26 := GenS2{F0: "a", F1: (v2 - v2)}
                    _ = v26
                case "foo bar":
                    v27 := GenS6{F0: (v6 - (-(v6)))}
                    _ = v27
                default:
                    v2 += v2
                }
            }
        } else {
            var v28 int = v6
            v29 := []int{(v28 ^ int(v2)), -405, (-(v6))}
            _ = v28
            _ = v29
        }
    }
    v2 += (v2 + (v2 % (((v2 ^ (v2 % ((v2) | 1)))) | 1)))
    var v30 float64 = 0.0
    if ((2 > 65535) && (!((false || (0 > 2147483648)) || (((1.5 / 1.0) - -2.25) != 3.14159)))) {
        fmt.Println(v6, v2, v6)
        if false {
            if false {
                v6 += -2147483648
                switch int(v2) {
                case 52:
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v32 := GenS2{F0: "\n\t", F1: (^(v2))}
                        _ = fi31
                        _ = v32
                    }
                }
                v33 := GenS6{F0: -9223372036854775808}
                _ = v33
            }
        }
        fmt.Println(v30)
    }
    if ((((v2 - v2) <= (^(v2))) || false) && (!(!true))) {
        if (v6 > 9223372036854775807) {
            var v34 int = 1
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 uint32 = 0
                v37 := GenS3{F0: false}
                var v38 int16 = -32768
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = v34
        }
        v2 -= (^(v2))
        v39 := []int{v6}
        _ = v39
    }
    v40 := []int{(v6 / (((v6 & (v6 & (v6 & v6)))) | 1))}
    _ = len(v40)
    var v41 int32 = 267
    v40 = append(v40, 517)
    var v42 int8 = (v2 * v2)
    v41 = v41
    if (-2.25 <= (-2.25 / float32(1e10))) {
        if (((0.0 + (-2.25 * v30)) < 3.14159) && true) {
            v43 := GenS4{F0: 156, F1: 572}
            var v44 float32 = -2.25
            v44 -= float32((((v30 + (0.0 * (v30 / ((v30 - v30) - 3.14159)))) * v30) + float64(((v44 * 0.1) / (-2.25 / v44)))))
            _ = v43
            _ = v44
        } else {
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 int64 = 1
                _ = fi45
                _ = v46
            }
            var v47 uint8 = 2
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
