package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 uint) uint {
    var v5 uint8 = 157
    p0 += (^(p0))
    switch -97 {
    case 4294967295:
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{fi6, (-(fi6)), fi6, (fi6 - fi6)}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi6
            _ = v7
        }
    default:
        v8 := GenS1{F0: 575}
        _ = v8
    }
    if false {
        fmt.Println(p0, p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return p0
}
func genFn1() int16 {
    var v9 int64 = 837
    if true {
        fmt.Println(v9)
        switch "\n\t" {
        case "foo bar":
            var v10 uint32 = 4294967295
            _ = v10
        default:
            if false {
                for fi11 := 0; fi11 < 11; fi11++ {
                    v9 += (v9 / ((v9) | 1))
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
            }
        }
    } else {
        v13 := []int{889, 2147483647, -684}
        v9 = (-(v9))
        _ = v13
    }
    _ = v9
    return -1
}

func main() {
    var v1 int = -565
    v2 := make(map[string]bool)
    var v3 int64 = 0
    v4 := make(map[uint32]bool)
    v14 := []int{v1, (v1 ^ v1)}
    v15 := GenS1{F0: (v1 << ((65535) & 63))}
    if (42 > 329) {
        v2[("hello" + "foo bar")] = (genFn0(2) < genFn0(987))
    } else {
        if len(v14) > 0 {
            v14[0] = v1
        }
    }
    switch "hello" {
    case "a":
        switch v1 {
        case 127:
            var v16 uint64 = 801
            _ = v16
        case 128:
            var v17 uint32 = 0
            _ = v17
        case 32768:
            var v18 int = 0
            _ = v18
        }
    default:
        if true {
            if false {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS2{F0: 1, F1: 239, F2: 80}
                    var v21 uint8 = 171
                    if true {
                        for fi22 := 0; fi22 < 12; fi22++ {
                            if false {
                                v21 -= (v21 + 255)
                                var v23 int8 = 100
                                _ = v23
                            }
                            var v24 uint32 = 2
                            _ = fi22
                            _ = v24
                        }
                        v25 := GenS2{F0: 2, F1: 2, F2: 17}
                        _ = v25
                    } else {
                        var v26 int32 = 20
                        if false {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 string = ("\n\t" + "")
                                var v29 float64 = (3.14159 + -2.25)
                                var v30 float32 = float32(float64((0.0 + 1.5)))
                                v31 := GenS1{F0: int(v3)}
                                _ = fi27
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        _ = v26
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            } else {
                if (!true) {
                    var v32 uint = 4294967295
                    _ = v32
                } else {
                    v33 := GenS2{F0: 18446744073709551615, F1: 255, F2: 59}
                    _ = v33
                }
                var v34 float64 = (0.1 * 3.14159)
                var v35 uint32 = 2147483648
                _ = v34
                _ = v35
            }
            var v36 uint32 = 130
            var v37 int = (v1 / (((v1 & -45)) | 1))
            _ = v36
            _ = v37
        }
    }
    v38 := v4[65535]
    for fi39 := 0; fi39 < 16; fi39++ {
        switch v1 {
        case 32768:
            v1 -= (fi39 << ((475) & 63))
        case 32767:
            fmt.Println(v1, fi39, v3)
        case -116:
            v38 = (-2.25 != 1e10)
        }
        fmt.Println(fi39, v38)
        v1 += int(v3)
        _ = fi39
    }
    v40 := GenS1{F0: v1}
    var v41 uint8 = 128
    var v42 int = (v1 ^ -1)
    fmt.Println(v1, v41, v38)
    var v43 int8 = 100
    v1 = 32767
    fmt.Println(v43, v1, v38)
    v44 := GenS2{F0: 4294967295, F1: 255, F2: int8(v41)}
    var v45 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
