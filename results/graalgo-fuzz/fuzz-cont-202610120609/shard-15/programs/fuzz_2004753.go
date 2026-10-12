package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
func genFn0(p0 uint64) uint {
    var v7 bool = true
    v8 := GenS1{F0: true, F1: 158}
    _ = p0
    _ = v7
    _ = v8
    return 304
}
func genFn1() uint8 {
    var v9 int = -2147483648
    var v10 bool = true
    v11 := GenS2{F0: "", F1: v9}
    _ = v10
    _ = v11
    _ = v9
    return 68
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 0
    var v3 int16 = v2
    v4 := make(map[int]uint32)
    var v5 string = ("" + "\n\t")
    var v6 int32 = -1
    if false {
        v4[-1] = 2
        v12 := GenS1{F0: false, F1: v6}
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(v3, fi13)
            switch fi13 {
            case 1:
                v1 += int8(fi14)
            }
            v2 += v3
            _ = fi14
        }
        v1 -= (v1 & 42)
        _ = fi13
    }
    var v15 uint16 = 0
    v16 := GenS2{F0: v5, F1: 516}
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS1{F0: false, F1: (v6 & v6)}
        _ = fi17
        _ = v18
    }
    if false {
        var v19 uint32 = 707
        v3 += 42
        _ = v19
    }
    if true {
        fmt.Println(v15, v15)
    }
    v20 := GenS2{F0: "", F1: 1}
    v21 := GenS2{F0: v5, F1: 2147483647}
    v22 := []int{0, 1, 32768, -2147483648}
    switch v5 {
    case "\n\t":
        v3 = (v2 << ((1) & 63))
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            if false {
                v22 = append(v22, fi23)
            }
            v2 = v2
            v25 := GenS1{F0: false, F1: (v6 << ((255) & 63))}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if len(v22) > 0 {
        v22[0] = -2147483648
    }
    v6 -= -2147483648
    switch v5 {
    case "\n\t":
        v1 = 42
    case "foo bar":
        for fi26 := 0; fi26 < 19; fi26++ {
            if (true || false) {
                var v27 uint64 = 1
                var v28 uint8 = 1
                v6 += (v6 / (((v6 % (((v6 - v6)) | 1))) | 1))
                _ = v27
                _ = v28
            }
            if false {
                v15 += (v15 - v15)
                var v29 uint64 = 18446744073709551615
                _ = v29
            } else {
                v30 := GenS1{F0: false, F1: (v6 << ((255) & 63))}
                if false {
                    v3 -= 32767
                    v3 -= (v2 | (v3 - (-(v2))))
                } else {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v6 = v6
                        if true {
                            v32 := v4[(fi26 | int(v3))]
                            switch (fi26 * fi26) {
                            case -2147483648:
                                v2 -= v2
                            case 499:
                                if true {
                                    var v33 uint32 = (v32 % (((v32 % (((v32 | v32)) | 1))) | 1))
                                    var v34 uint64 = 1
                                    _ = v33
                                    _ = v34
                                }
                            case -64:
                                switch v5 {
                                case "foo bar":
                                    v35 := GenS1{F0: (1 >= 42), F1: v6}
                                    _ = v35
                                case "\n\t":
                                    var v36 uint64 = 2
                                    _ = v36
                                case "hello":
                                    v2 = (v2 + (^(v3)))
                                default:
                                    v4[377] = v32
                                }
                            default:
                                _ = len(v22)
                            }
                            v37 := GenS2{F0: "a", F1: (fi26 ^ (fi26 & (fi31 & (fi26 | (fi31 & (fi31 ^ 32768))))))}
                            _ = v32
                            _ = v37
                        } else {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            if len(v22) > 0 {
                                v22[0] = fi26
                            }
                            switch (v5 + v5) {
                            case "\n\t":
                                v3 = 0
                            case "":
                                var v38 int = int(v6)
                                _ = v38
                            case "foo bar":
                                var v39 float64 = float64(1e10)
                                _ = v39
                            }
                        }
                        var v40 int = fi26
                        _ = fi31
                        _ = v40
                    }
                    var v41 int16 = v2
                    for fi42 := 0; fi42 < 18; fi42++ {
                        v43 := GenS2{F0: ("hello" + ((v5 + "\n\t") + "\n\t")), F1: (fi26 & fi26)}
                        v44 := GenS2{F0: v5, F1: (fi26 & int(v15))}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    _ = v41
                }
                _ = v30
            }
            switch 32768 {
            case -1:
                var v45 int = fi26
                _ = v45
            case -9223372036854775808:
                v6 -= (v6 / (((v6 & v6)) | 1))
            }
            _ = fi26
        }
    default:
        var v46 int16 = int16(v1)
        _ = v46
    }
    v47 := GenS3{F0: v3}
    v48 := GenS4{F0: 1, F1: 128, F2: v6}
    for fi49 := 0; fi49 < 14; fi49++ {
        v50 := GenS4{F0: 42, F1: 159, F2: (v6 / (((-(v6))) | 1))}
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
