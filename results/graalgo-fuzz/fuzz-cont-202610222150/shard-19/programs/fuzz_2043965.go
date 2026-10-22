package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int8; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 bool, p1 int16) int {
    switch 32767 {
    case 9223372036854775807:
        var v5 float32 = -2.25
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return -236
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 669
    var v3 int = 91
    v4 := []int{}
    var v6 float64 = ((float64((float32(1.5) + float32(((((1e10 * (-2.25 * 1.5)) / 1e10) / -2.25) - 1.5)))) + ((0.0 - 1e10) * 1.5)) + (0.1 - 3.14159))
    if (!(493 >= 481)) {
        switch (v3 & (v3 % ((genFn0((9223372036854775807 > -220), -781)) | 1))) {
        case -742:
            fmt.Println(v6, v3, v3)
        case -1:
            v7 := GenS1{F0: 171, F1: -2147483648}
            _ = v7
        case 995:
            v2 = v2
        default:
            v8 := GenS2{F0: 1000, F1: -1, F2: -2147483648}
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS2{F0: 1, F1: (v3 - fi10), F2: (^(fi10))}
                _ = fi10
                _ = v11
            }
            switch "foo bar" {
            case "hello":
                if true {
                    var v12 uint32 = v2
                    _ = v12
                } else {
                    v13 := GenS3{F0: v3}
                    _ = v13
                }
            case "a":
                var v14 uint16 = 274
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 2; fi15++ {
                    v2 = (v2 - (v2 * (v2 >> ((1) & 63))))
                    v2 += v2
                    v16 := GenS4{F0: (fi9 * genFn0(false, -383)), F1: -128, F2: -1}
                    _ = fi15
                    _ = v16
                }
            }
            v3 -= genFn0(false, 0)
            _ = fi9
        }
        v1[42] = (3.14159 > 1.5)
    }
    var v17 bool = true
    fmt.Println(v17, v2, v3)
    v17 = (42 < 0)
    v17 = true
    if (((42 > 1) && true) || (1 < (v3 << ((0) & 63)))) {
        if v17 {
            if v17 {
                var v18 bool = v17
                v19 := GenS5{F0: 2, F1: 32767, F2: 0}
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 int8 = -23
                    fmt.Println(v18, v21, v6)
                    _ = fi20
                    _ = v21
                }
                _ = v18
                _ = v19
            } else {
                for fi22 := 0; fi22 < 9; fi22++ {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS6{F0: (^(v2)), F1: -2147483648, F2: 1}
                        v25 := GenS3{F0: v3}
                        if v17 {
                            v17 = (true && ((480 != 665) || v17))
                            var v26 uint16 = 255
                            var v27 uint64 = 18446744073709551615
                            if len(v4) > 0 {
                                v4[0] = 128
                            }
                            if (!v17) {
                                var v28 string = ("\n\t" + "\n\t")
                                if false {
                                    fmt.Println(v2, v2, v17)
                                }
                                _ = v28
                            }
                            _ = v26
                            _ = v27
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    if ((v2 << ((2) & 63)) <= (v2 ^ 2147483647)) {
                        v4 = append(v4, v3)
                        fmt.Println(fi22)
                        _ = len(v4)
                    }
                    _ = fi22
                }
                var v29 uint = 693
                var v30 uint32 = (^(v2))
                _ = v29
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v32 := v1[(v2 + v2)]
        if v32 {
            var v33 uint8 = 127
            v34 := GenS3{F0: (v3 % ((v3) | 1))}
            _ = v33
            _ = v34
        }
        if true {
            var v35 int = (fi31 | (-(fi31)))
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 int = genFn0(((false && (100 <= 1)) && v32), 0)
                if false {
                    v38 := v1[v2]
                    _ = v38
                }
                _ = fi36
                _ = v37
            }
            var v39 uint = 2
            _ = v35
            _ = v39
        } else {
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 uint = 0
                if len(v4) > 0 {
                    v4[0] = genFn0(false, 0)
                }
                _ = fi40
                _ = v41
            }
        }
        _ = fi31
        _ = v32
    }
    switch "foo bar" {
    case "a":
        if v17 {
            for fi42 := 0; fi42 < 17; fi42++ {
                v2 = 2
                _ = fi42
            }
            delete(v1, v2)
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1[v2] = (!(1 < 65535))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
