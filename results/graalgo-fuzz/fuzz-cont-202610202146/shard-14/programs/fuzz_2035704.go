package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint, p1 uint32, p2 int16) uint32 {
    fmt.Println(p0, p2)
    if (0.0 != (0.1 - (float64((float32(1e10) * float32(3.14159))) + 1e10))) {
        v7 := GenS1{F0: 433, F1: (p1 / ((2) | 1))}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 bool, p1 int64) uint {
    switch -52 {
    case -2147483648:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{}
            v10 := GenS1{F0: fi8, F1: 272}
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 int64 = -436
        _ = v11
    }
    p0 = p0
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int16 = 42
    var v2 int8 = 107
    var v3 uint64 = 908
    var v4 int = 720
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    v6["a"] = "\n\t"
    v4 -= (v4 << ((255) & 63))
    v12 := GenS1{F0: v4, F1: genFn0(18446744073709551615, 2147483648, 916)}
    fmt.Println(v1)
    if (int8(v3) >= (v2 % ((0) | 1))) {
        v3 = (v3 | v3)
        var v13 int64 = 9223372036854775807
        v14 := GenS1{F0: 980, F1: 1}
        _ = v13
        _ = v14
    } else {
        if false {
            v2 = (v2 | -1)
            for fi15 := 0; fi15 < 6; fi15++ {
                v3 = (^(v3))
                v16 := []int{}
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int32 = 403
                    var v19 float64 = (0.1 - (0.0 + -2.25))
                    switch v5 {
                    case "\n\t":
                        v20 := GenS1{F0: fi17, F1: 2147483647}
                        _ = v20
                    case "hello":
                        var v21 float64 = -2.25
                        _ = v21
                    case "a":
                        v3 = v3
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
        } else {
            v22 := GenS2{F0: v4, F1: 42, F2: 0}
            var v23 int16 = (v1 + (v1 + 1))
            _ = v22
            _ = v23
        }
        if true {
            if true {
                v2 -= 127
                var v24 int32 = -250
                v24 = -961
                _ = v24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 bool = true
                    fmt.Println(v1, v2)
                    _ = fi26
                    _ = v27
                }
                var v28 int64 = 0
                _ = fi25
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := GenS2{F0: (v4 << ((557) & 63)), F1: 0, F2: 1}
            if true {
                v31 := []int{(fi29 / (((-(v4))) | 1))}
                var v32 int32 = 1
                _ = v31
                _ = v32
            }
            v33 := []int{v4, -2147483648, (fi29 - fi29)}
            _ = fi29
            _ = v30
            _ = v33
        }
    }
    var v34 uint16 = 895
    v4 -= (v4 % ((v4) | 1))
    if true {
        if true {
            fmt.Println(v2, v5, v3)
            fmt.Println(v2, v3)
        } else {
            if (v34 < (v34 & (v34 % ((v34) | 1)))) {
                fmt.Println(v3, v4)
                var v35 uint16 = (v34 / (((^(v34))) | 1))
                _ = v35
            } else {
                for fi36 := 0; fi36 < 1; fi36++ {
                    v3 += v3
                    switch v4 {
                    case 0:
                        v37 := GenS1{F0: v4, F1: 2}
                        _ = v37
                    default:
                        v1 = -32768
                    }
                    _ = fi36
                }
            }
        }
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 uint = 255
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 uint = 65535
        for fi42 := 0; fi42 < 9; fi42++ {
            if true {
                v43 := v6[v5]
                _ = v43
            } else {
                v5 = "foo bar"
                v3 -= 761
            }
            v44 := GenS1{F0: (v4 | fi40), F1: 96}
            v45 := GenS3{F0: (v1 >> ((genFn1((42 >= v34), 153)) & 63))}
            _ = fi42
            _ = v44
            _ = v45
        }
        v4 += v4
        _ = fi40
        _ = v41
    }
    if (!false) {
        fmt.Println(v5, v3, v34)
        v6[(v5 + "\n\t")] = v5
        fmt.Println(v2, v4, v1)
    }
    var v46 int = (v4 | (v4 - -2147483648))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
