package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 uint16; F2 bool }
func genFn0(p0 string, p1 string) int16 {
    if false {
        p0 = p0
        var v4 int64 = -1
        _ = v4
    } else {
        var v5 int8 = 1
        _ = v5
    }
    v6 := GenS1{F0: 1}
    var v7 uint8 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint64) uint {
    var v8 int16 = int16(p1)
    v8 -= -1
    v9 := GenS2{F0: v8, F1: true}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "foo bar"
    var v3 int16 = -233
    v10 := GenS2{F0: (v3 - (v3 << ((18446744073709551615) & 63))), F1: (false && (32768 >= -697))}
    v2 = ((v2 + v2) + v2)
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint8 = 157
        var v14 int8 = 100
        v11 += v11
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 -= (-(v3))
    for fi15 := 0; fi15 < 19; fi15++ {
        if (!false) {
            if false {
                var v16 uint8 = 127
                delete(v1, (v2 + v2))
                _ = v16
            } else {
                if (617 >= -2147483648) {
                    if (2147483647 > 42) {
                        fmt.Println(v11, fi15)
                        if (!true) {
                            fmt.Println(v2, fi15, v11)
                            var v17 int16 = -588
                            _ = v17
                        } else {
                            fmt.Println(v3)
                            if (0.0 > 1.5) {
                                fmt.Println(v11)
                            }
                            v18 := []int{(-(fi15)), fi15}
                            _ = v18
                        }
                        switch fi15 {
                        case 32768:
                            v19 := v1[(v2 + (("" + (v2 + v2)) + v2))]
                            _ = v19
                        case 509:
                            var v20 int32 = -2147483648
                            _ = v20
                        case 0:
                            fmt.Println(v3)
                        default:
                            var v21 bool = false
                            _ = v21
                        }
                    } else {
                        var v22 float64 = (float64(1e10) / 3.14159)
                        v3 = -700
                        _ = v22
                    }
                    var v23 int8 = -82
                    _ = v23
                } else {
                    if true {
                        v24 := GenS3{F0: (v3 + (v3 >> ((genFn1(fi15, 18446744073709551615)) & 63))), F1: 42, F2: true}
                        v2 = (v2 + (v2 + v2))
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 4; fi25++ {
                            fmt.Println(v3, v3, fi25)
                            _ = fi25
                        }
                    }
                    var v26 float64 = 0.1
                    v27 := GenS3{F0: v3, F1: 42, F2: false}
                    _ = v26
                    _ = v27
                }
            }
            v3 = genFn0(v2, v2)
        } else {
            var v28 string = ("\n\t" + ((v2 + v2) + v2))
            fmt.Println(v11, v2)
            _ = v28
        }
        fmt.Println(v3)
        if (1e10 <= (0.0 + 0.1)) {
            var v29 string = v2
            if false {
                fmt.Println(v3)
            }
            v29 = (v2 + (("" + "") + v2))
            _ = v29
        } else {
            var v30 float32 = 1.5
            if false {
                if false {
                    v31 := v1[((v2 + v2) + ("a" + v2))]
                    var v32 uint8 = 2
                    _ = v31
                    _ = v32
                }
                v1[v2] = 42
                fmt.Println(v2)
            } else {
                v33 := GenS1{F0: 363}
                _ = v33
            }
            _ = v30
        }
        _ = fi15
    }
    var v34 bool = (0 < 1)
    var v35 float64 = -2.25
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := []int{fi36}
        switch "a" {
        case "\n\t":
            if len(v37) > 0 {
                v37[0] = 0
            }
        default:
            v11 = v11
        }
        _ = fi36
        _ = v37
    }
    var v38 uint = 65535
    v39 := []int{32768, 0}
    var v40 uint = uint(v11)
    var v41 uint8 = 42
    var v42 uint8 = (^(v41))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
