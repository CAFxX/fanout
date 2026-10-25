package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint32) int64 {
    p0 = p0
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            p0 += p0
            p0 += (p0 & 65535)
            fmt.Println(fi4)
            v6 := GenS1{F0: (("hello" + ("hello" + ("foo bar" + "\n\t"))) + "hello"), F1: 210, F2: 103}
            p0 = p0
            p0 = 945
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := []int{443, 42, 42, 998}
    v7 = append(v7, 543)
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 string, p1 uint8, p2 int) int64 {
    v8 := GenS2{F0: 482, F1: (p0 + p0), F2: 0}
    v9 := GenS2{F0: 65535, F1: ("\n\t" + p0), F2: 82}
    v10 := []int{(^(p2)), p2, 759, (p2 % (((p2 / ((32767) | 1))) | 1))}
    switch p2 {
    case 32767:
        for fi11 := 0; fi11 < 18; fi11++ {
            p2 = (fi11 << ((255) & 63))
            _ = fi11
        }
    default:
        if false {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(p1, p0, p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return genFn0(970)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    v3 := []uint32{2147483648, 2147483647, 815}
    v1[32768] = 476
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(fi12, v2)
        if false {
            var v13 float64 = -2.25
            _ = v13
        } else {
            if false {
                if ((1e10 - v2) <= 0.1) {
                    var v14 int64 = genFn0(857)
                    v15 := v1[fi12]
                    var v16 int = (fi12 & (fi12 >> ((4294967295) & 63)))
                    _ = v14
                    _ = v15
                    _ = v16
                }
                switch fi12 {
                case -9223372036854775808:
                    var v17 int64 = 877
                    _ = v17
                case 42:
                    if true {
                        if false {
                            var v18 uint = 370
                            _ = v18
                        }
                        switch "\n\t" {
                        case "a":
                            var v19 string = "hello"
                            _ = v19
                        case "":
                            if (-1 != 42) {
                                v1[fi12] = 140
                            } else {
                                if (!(1 > -32768)) {
                                    v20 := GenS2{F0: 1, F1: ("\n\t" + "foo bar"), F2: -86}
                                    fmt.Println(fi12, v2)
                                    fmt.Println(v2)
                                    _ = v20
                                } else {
                                    fmt.Println(v2, fi12)
                                    v2 -= 0.1
                                    v2 += (float32(float64(0.0)) + float32((((0.1 + 1.5) / 1.5) + 0.1)))
                                }
                            }
                        case "hello":
                            var v21 int32 = 499
                            _ = v21
                        default:
                            var v22 uint8 = 1
                            _ = v22
                        }
                        v2 = (v2 - ((v2 + ((3.14159 * (0.0 - float32(3.14159))) / (0.0 - v2))) * (v2 * (1e10 + v2))))
                    }
                case -468:
                    if false {
                        v2 -= (0.1 / v2)
                        v2 -= -2.25
                        _ = len(v3)
                    } else {
                        v23 := GenS3{F0: 777, F1: 2}
                        for fi24 := 0; fi24 < 4; fi24++ {
                            fmt.Println(fi24, fi12)
                            _ = fi24
                        }
                        _ = v23
                    }
                default:
                    var v25 int16 = -32768
                    _ = v25
                }
                switch fi12 {
                case 773:
                    if false {
                        var v26 int32 = -1
                        v27 := GenS2{F0: 476, F1: "foo bar", F2: 1}
                        fmt.Println(fi12)
                        _ = v26
                        _ = v27
                    }
                }
            } else {
                var v28 int32 = 2147483647
                for fi29 := 0; fi29 < 20; fi29++ {
                    v30 := GenS1{F0: "", F1: 255, F2: 159}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v28, fi12)
                _ = v28
            }
            var v31 float32 = (1e10 / (v2 + (((1e10 - (v2 + -2.25)) - (1e10 / (v2 / v2))) / v2)))
            _ = v31
        }
        _ = fi12
    }
    var v32 int64 = 42
    var v33 string = ("\n\t" + "foo bar")
    var v34 bool = false
    var v35 float32 = v2
    v33 = (v33 + "a")
    _ = len(v3)
    v32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
