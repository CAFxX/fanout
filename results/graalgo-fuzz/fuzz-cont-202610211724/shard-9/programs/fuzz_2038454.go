package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint64; F1 bool; F2 int }
type GenS5 struct { F0 uint; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 string) int8 {
    v6 := GenS1{F0: 65535, F1: 598, F2: "\n\t"}
    v7 := GenS1{F0: 255, F1: 1, F2: p0}
    if (!(!false)) {
        var v8 uint8 = 210
        fmt.Println(v8, p0)
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{}
            p0 = p0
            _ = len(v10)
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v11 := []int{42, -965, 151, 294}
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    return 86
}
func genFn1() int16 {
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 float64 = (-2.25 * (1e10 / 3.14159))
            v14 = ((v14 * (v14 / 1e10)) + 3.14159)
            v14 = v14
            v15 := GenS2{F0: 805}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int8 = -40
        _ = fi12
        _ = v16
    }
    switch -1 {
    case -676:
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS1{F0: 42, F1: -94, F2: "foo bar"}
            _ = fi17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint64 = 225
            _ = fi19
            _ = v20
        }
    }
    v21 := []int{9223372036854775807}
    _ = v21
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ((("\n\t" + ("\n\t" + "")) + ("" + "\n\t")) + "a")
    v3 := make(map[uint32]int)
    var v4 float64 = 3.14159
    v5 := make(map[string]string)
    var v22 int = 2147483647
    for fi23 := 0; fi23 < 11; fi23++ {
        if true {
            v24 := []int{(v22 << ((255) & 63)), v22, (v22 % (((fi23 >> ((0) & 63))) | 1))}
            switch "\n\t" {
            case "\n\t":
                var v25 float32 = 3.14159
                _ = v25
            case "":
                var v26 int64 = -1
                _ = v26
            default:
                v4 = (1.5 * float64(((1e10 + 0.0) - 1.5)))
            }
            var v27 int8 = 1
            _ = v24
            _ = v27
        }
        _ = fi23
    }
    fmt.Println(v4, v22, v2)
    if true {
        v28 := []int{v22, 1, v22, v22}
        var v29 int8 = 0
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 uint8 = 128
            if true {
                fmt.Println(fi30, v22)
            }
            if ((!(int16(v31) == 866)) || (!true)) {
                v32 := GenS3{F0: 127, F1: v2}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 16; fi33++ {
                    if false {
                        var v34 uint32 = 2
                        fmt.Println(v2)
                        switch 128 {
                        case -2147483648:
                            var v35 uint16 = 388
                            _ = v35
                        default:
                            v36 := []int{v22}
                            _ = v36
                        }
                        _ = v34
                    }
                    switch v22 {
                    case -2147483648:
                        if true {
                            v37 := GenS4{F0: 4294967295, F1: (false || (967 >= 2)), F2: fi33}
                            var v38 bool = false
                            _ = v37
                            _ = v38
                        }
                    }
                    _ = fi33
                }
            }
            _ = fi30
            _ = v31
        }
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 int8 = 1
            _ = fi39
            _ = v40
        }
        v41 := GenS1{F0: 148, F1: 0, F2: v2}
        _ = v41
    }
    var v42 uint16 = 42
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 uint32 = 65535
        v45 := GenS2{F0: 42}
        v2 = ""
        _ = fi43
        _ = v44
        _ = v45
    }
    if true {
        v3[9] = -1
        v46 := GenS4{F0: 4294967295, F1: true, F2: -1}
        v42 = (^(v42))
        _ = v46
    } else {
        var v47 string = "foo bar"
        var v48 uint32 = 205
        _ = v47
        _ = v48
    }
    v49 := v3[315]
    v50 := GenS1{F0: 255, F1: 42, F2: v2}
    v51 := v1[(v22 % ((v22) | 1))]
    if (v4 < (v4 * 1.5)) {
        v52 := GenS5{F0: 42, F1: genFn1(), F2: (^(v42))}
        _ = v52
    } else {
        v49 = v51
        v49 -= (v49 ^ v22)
        fmt.Println(v4, v2, v49)
    }
    v53 := GenS6{F0: 120, F1: 2, F2: 127}
    if true {
        fmt.Println(v49)
        v51 = v51
    } else {
        v22 += (v49 * v49)
        v54 := GenS2{F0: 869}
        fmt.Println(v42, v2, v4)
        _ = v54
    }
    v55 := []int{-436, v49, (v22 & (v49 % ((v51) | 1)))}
    fmt.Println(v2, v51, v4)
    v56 := GenS7{F0: 577, F1: 2, F2: -252}
    switch -2147483648 {
    case 42:
        var v57 string = v2
        _ = v57
    case 850:
        v58 := v5[v2]
        _ = v58
    case 921:
        if false {
            var v59 float32 = -2.25
            v60 := GenS7{F0: 266, F1: 42, F2: -2147483648}
            v49 = v51
            _ = v59
            _ = v60
        }
    default:
        v22 += (^(v51))
    }
    var v61 float64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v55
    _ = v56
    _ = v61
}
