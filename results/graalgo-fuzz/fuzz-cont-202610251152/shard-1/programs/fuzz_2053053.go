package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
type GenS5 struct { F0 uint; F1 bool; F2 bool }
func genFn0() uint32 {
    var v6 uint64 = 742
    switch 9223372036854775807 {
    case 164:
        fmt.Println(v6)
    case 0:
        for fi7 := 0; fi7 < 19; fi7++ {
            if ((-2.25 - 0.0) > -2.25) {
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 float32 = 0.1
                    v6 = v6
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    if (false && false) {
                        fmt.Println(fi10, fi7)
                        v11 := GenS1{F0: -1}
                        _ = v11
                    } else {
                        v6 += v6
                        if ((fi7 << ((617) & 63)) != (fi10 % ((0) | 1))) {
                            v12 := []int{128}
                            var v13 int16 = 0
                            _ = len(v12)
                            v13 -= v13
                            _ = v12
                            _ = v13
                        }
                    }
                    var v14 bool = (125 != 3)
                    _ = fi10
                    _ = v14
                }
            }
            v15 := []int{(fi7 * (fi7 + fi7))}
            _ = fi7
            _ = v15
        }
    case -73:
        v16 := GenS2{F0: ("hello" + ""), F1: 2}
        _ = v16
    default:
        fmt.Println(v6)
    }
    v17 := GenS1{F0: 42}
    v18 := []int{-506, 9223372036854775807, 9223372036854775807, 32767}
    _ = v17
    _ = v18
    _ = v6
    return 1
}
func genFn1() uint16 {
    var v19 float64 = -2.25
    fmt.Println(v19)
    v19 += (-2.25 + (v19 - (-2.25 - -2.25)))
    v19 -= v19
    v20 := GenS3{F0: 42, F1: 151}
    _ = v19
    _ = v20
    return 255
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 float32 = v1
    v5 := []uint32{0, 499}
    var v21 int32 = 0
    var v22 int = -1
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v4, v1)
        if ((v21 ^ v21) <= (-(v21))) {
            v22 -= v22
            _ = len(v2)
            v1 = 0.1
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int32 = v21
            fmt.Println(v4, v25)
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v4 = (((((v4 * (v4 / v4)) - (v4 * (v4 * v1))) * 0.0) / v1) + v1)
    var v26 bool = true
    fmt.Println(v22, v1, v4)
    v27 := GenS4{F0: 4294967295, F1: (v22 ^ 593), F2: 255}
    switch "foo bar" {
    case "hello":
        v26 = v26
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "":
        v28 := GenS4{F0: 405, F1: (v22 % ((v22) | 1)), F2: genFn1()}
        _ = v28
    }
    if false {
        var v29 int16 = (v3 * (-(v3)))
        _ = v29
    }
    if (2147483648 <= genFn0()) {
        v30 := GenS5{F0: 846, F1: ((!true) && (732 != 356)), F2: (!(v26 && v26))}
        v1 -= (1.5 - -2.25)
        _ = v30
    } else {
        v31 := GenS5{F0: 0, F1: (((175 != genFn0()) || (65535 >= 1)) || ("\n\t" < "foo bar")), F2: v26}
        delete(v2, genFn0())
        _ = v31
    }
    if true {
        v32 := GenS1{F0: (v21 + (-(v21)))}
        if v26 {
            if true {
                v4 = v4
                if ((100 > 1) || (32767 < 1)) {
                    var v33 float32 = 1e10
                    var v34 float32 = 1.5
                    fmt.Println(v34, v22, v22)
                    _ = v33
                    _ = v34
                }
            }
        } else {
            var v35 int64 = -9223372036854775808
            var v36 uint8 = 128
            fmt.Println(v1)
            _ = v35
            _ = v36
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
