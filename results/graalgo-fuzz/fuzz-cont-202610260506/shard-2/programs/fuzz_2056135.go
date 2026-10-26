package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
func genFn0(p0 bool, p1 string) uint32 {
    v5 := []int{32768}
    for fi6 := 0; fi6 < 14; fi6++ {
        switch "\n\t" {
        case "foo bar":
            var v7 uint64 = 606
            _ = v7
        default:
            for fi8 := 0; fi8 < 6; fi8++ {
                p0 = true
                if ((((-2.25 + (-2.25 + (0.0 - 0.1))) - (-2.25 + 1.5)) * 1.5) == 3.14159) {
                    v9 := GenS1{F0: 42, F1: 32767}
                    p1 = ""
                    _ = v9
                } else {
                    switch p1 {
                    case "a":
                        v10 := GenS1{F0: 2, F1: (fi8 * (^(fi6)))}
                        _ = v10
                    case "":
                        if p0 {
                            var v11 int8 = 127
                            p1 = (p1 + p1)
                            _ = len(v5)
                            _ = v11
                        }
                    default:
                        fmt.Println(fi8, fi6, p1)
                        if (!(42 > -1)) {
                            v12 := GenS1{F0: 255, F1: fi8}
                            _ = v12
                        }
                        v5 = append(v5, fi8)
                        if p0 {
                            v5 = append(v5, (fi6 >> ((18446744073709551615) & 63)))
                            switch p1 {
                            case "hello":
                                var v13 bool = p0
                                _ = v13
                            default:
                                fmt.Println(p1)
                            }
                        } else {
                            if p0 {
                                switch fi8 {
                                case 4294967295:
                                    v14 := GenS2{F0: -8, F1: 42, F2: p0}
                                    _ = v14
                                case 656:
                                    _ = len(v5)
                                case 79:
                                    fmt.Println(p1)
                                }
                                fmt.Println(p1)
                            } else {
                                p0 = (809 > 963)
                                if (32767 != 32767) {
                                    var v15 int16 = 32767
                                    if ((0.0 - (-2.25 * (3.14159 + 1.5))) != 1e10) {
                                        switch "a" {
                                        case "":
                                            var v16 int16 = (v15 | (v15 % (((v15 % ((-816) | 1))) | 1)))
                                            _ = v16
                                        case "a":
                                            v15 = v15
                                        default:
                                            var v17 uint32 = 65535
                                            _ = v17
                                        }
                                    }
                                    var v18 string = p1
                                    _ = v15
                                    _ = v18
                                }
                            }
                        }
                    }
                    v19 := GenS1{F0: 321, F1: (fi6 / ((fi6) | 1))}
                    _ = v19
                }
                p0 = true
                _ = fi8
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 bool = (928 >= 0)
    var v2 bool = v1
    var v3 uint = 350
    var v4 float32 = (float32(0.0) / 0.1)
    var v20 uint8 = 0
    var v21 uint32 = genFn0((!v2), ("foo bar" + "hello"))
    v21 = 0
    if v2 {
        if false {
            for fi22 := 0; fi22 < 16; fi22++ {
                if ((127 > 0) || ("foo bar" >= (("\n\t" + "hello") + (("\n\t" + "") + "a")))) {
                    v23 := []int{(-(fi22)), 32768}
                    v24 := GenS3{F0: fi22, F1: v3, F2: v20}
                    switch (fi22 % (((-(fi22))) | 1)) {
                    case 127:
                        v21 = 2147483648
                    case 32767:
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v4 += 1e10
                            v4 = (0.0 + (0.0 - (v4 - v4)))
                            _ = fi25
                        }
                    case 4294967295:
                        fmt.Println(v20, v1)
                    default:
                        switch fi22 {
                        case 127:
                            switch "foo bar" {
                            case "foo bar":
                                v21 = (v21 / (((v21 | v21)) | 1))
                            case "a":
                                v26 := GenS3{F0: fi22, F1: (v3 >> (((v3 * (v3 % ((255) | 1)))) & 63)), F2: v20}
                                _ = v26
                            case "\n\t":
                                v20 += v20
                            }
                        case 2147483647:
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        default:
                            v20 = v20
                        }
                    }
                    _ = v23
                    _ = v24
                }
                _ = fi22
            }
        } else {
            v2 = (true && (!(v2 || (3.14159 < 0.0))))
            var v28 uint16 = 203
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := GenS3{F0: -1, F1: v3, F2: (^(v20))}
        v31 := GenS3{F0: (fi29 ^ -1), F1: (v3 << (((v3 >> (((v3 << (((v3 - (v3 >> (((v3 ^ (^(v3)))) & 63)))) & 63))) & 63))) & 63)), F2: 1}
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS2{F0: 100, F1: 32767, F2: (v3 == v3)}
    fmt.Println(v4, v4)
    if (v3 > 2) {
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 uint = (v3 / ((v3) | 1))
            _ = fi33
            _ = v34
        }
        fmt.Println(v4, v4, v3)
        for fi35 := 0; fi35 < 11; fi35++ {
            switch (fi35 * (fi35 - (fi35 % ((fi35) | 1)))) {
            case -129:
                v36 := GenS2{F0: 0, F1: -1, F2: (!(v1 && (!false)))}
                _ = v36
            case -237:
                v37 := GenS4{F0: 18446744073709551615, F1: ("foo bar" + "\n\t"), F2: 0}
                _ = v37
            default:
                v21 -= 65535
            }
            v38 := GenS1{F0: (v3 | v3), F1: fi35}
            _ = fi35
            _ = v38
        }
    }
    v39 := []int{0, 364, -9223372036854775808, -1}
    v4 = ((v4 - ((float32(1.5) - (((1e10 - -2.25) * float32((0.1 - 0.1))) * v4)) / v4)) * 3.14159)
    var v40 int8 = -1
    var v41 string = "hello"
    v42 := GenS1{F0: 614, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
