package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 bool, p1 uint) uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{fi4, (fi4 ^ (fi4 % ((fi4) | 1))), fi4}
        var v6 uint32 = 2147483647
        v7 := GenS1{F0: 821, F1: (v6 - v6)}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    p1 = p1
    for fi8 := 0; fi8 < 7; fi8++ {
        if ((!true) && false) {
            if (!(!(p0 || p0))) {
                v9 := GenS1{F0: 18446744073709551615, F1: 2147483647}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 12; fi10++ {
                    p0 = (true && false)
                    _ = fi10
                }
                if false {
                    p1 -= p1
                    if (!false) {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            if ("a" == "a") {
                                p0 = ((fi8 & fi11) != (-(fi8)))
                                p0 = true
                                var v12 int = (fi11 << ((2) & 63))
                                _ = v12
                            }
                            _ = fi11
                        }
                    } else {
                        var v13 uint32 = 2147483647
                        if (127 >= 255) {
                            fmt.Println(p0, p0)
                            v14 := GenS1{F0: 925, F1: 392}
                            var v15 uint32 = v13
                            _ = v14
                            _ = v15
                        }
                        p1 -= p1
                        _ = v13
                    }
                    v16 := GenS1{F0: 2, F1: 2147483648}
                    _ = v16
                } else {
                    p0 = (3.14159 > 0.1)
                    p0 = p0
                    p1 += p1
                }
                p0 = (p0 && p0)
            }
            if p0 {
                var v17 float32 = float32((1.5 - 0.1))
                _ = v17
            } else {
                if p0 {
                    p1 = 1
                    var v18 uint16 = 1
                    _ = v18
                } else {
                    p0 = true
                    v19 := GenS1{F0: 2, F1: 249}
                    if p0 {
                        p0 = (1.5 != (-2.25 / ((float32((float64(0.1) / (1.5 + float64(float32(((-2.25 / -2.25) * 3.14159)))))) * 1.5) * (0.1 * -2.25))))
                    }
                    _ = v19
                }
            }
            var v20 float64 = 0.0
            _ = v20
        } else {
            p1 += (p1 ^ (^(p1)))
            v21 := GenS1{F0: 2, F1: 1}
            _ = v21
        }
        _ = fi8
    }
    if p0 {
        v22 := GenS1{F0: 18446744073709551615, F1: 317}
        _ = v22
    }
    _ = p0
    _ = p1
    return 223
}
func genFn1(p0 int16, p1 uint, p2 uint16) uint8 {
    var v23 float32 = 0.0
    v23 = float32(0.1)
    v23 += v23
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 151
}

func main() {
    var v1 float64 = 0.1
    var v2 float32 = 1e10
    var v3 bool = true
    var v24 float32 = v2
    if v3 {
        var v25 uint16 = 65535
        var v26 float32 = 3.14159
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := []int{(fi27 - fi27)}
            switch ("" + "") {
            case "\n\t":
                v29 := GenS2{F0: genFn0(((v3 || v3) && v3), 65535)}
                _ = v29
            case "hello":
                for fi30 := 0; fi30 < 19; fi30++ {
                    if len(v28) > 0 {
                        v28[0] = (fi27 & (fi30 & 1))
                    }
                    v2 = v24
                    _ = fi30
                }
            case "foo bar":
                var v31 int8 = -1
                _ = v31
            default:
                fmt.Println(v1)
            }
            _ = fi27
            _ = v28
        }
        v32 := GenS3{F0: -691, F1: 238}
        var v33 int32 = -1
        _ = v32
        _ = v33
    }
    v34 := GenS3{F0: -151, F1: 42}
    v35 := []int{428, 422}
    switch 127 {
    case -776:
        switch "hello" {
        case "a":
            v1 = (float64(float32(((v1 * float64(((3.14159 * v2) / (((v2 * v2) - 1e10) * v2)))) / v1))) + v1)
        case "foo bar":
            if len(v35) > 0 {
                v35[0] = -9223372036854775808
            }
        }
    case 0:
        var v36 uint16 = 1
        _ = v36
    case 32768:
        v2 = ((float32(((((1.5 - (v1 * v1)) / (v1 - 3.14159)) - float64(float32(((v1 / (v1 + (0.1 + (1e10 * (0.0 * v1))))) - v1)))) * -2.25)) * v24) - v24)
    default:
        v1 = v1
    }
    v37 := GenS3{F0: -32768, F1: 255}
    v38 := GenS4{F0: 337, F1: 2}
    var v39 int16 = 32767
    v39 -= (^(v39))
    v40 := GenS4{F0: 32767, F1: 8}
    v41 := GenS3{F0: (-(v39)), F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
