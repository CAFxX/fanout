package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 int32, p2 string) uint8 {
    v7 := GenS1{F0: p1}
    v8 := GenS2{F0: 129, F1: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16) uint8 {
    v9 := []int{-482}
    p0 -= (p0 | (^(p0)))
    v9 = append(v9, -2147483648)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 uint16 = 442
    v4 := make(map[int]uint32)
    var v5 float32 = float32(-2.25)
    v6 := make(map[int]int)
    var v10 uint8 = genFn0(22, 2147483647, "hello")
    var v11 string = "hello"
    var v12 int16 = 825
    v13 := GenS3{F0: -1}
    switch 127 {
    case -2147483648:
        v14 := GenS1{F0: -1}
        _ = v14
    case 127:
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint64 = 18446744073709551615
            if false {
                v17 := v1[260]
                _ = v17
            } else {
                v16 += (v16 & (v16 | (v16 | 1)))
            }
            _ = fi15
            _ = v16
        }
    case 32767:
        if false {
            if false {
                var v18 int64 = 9223372036854775807
                var v19 float64 = float64((3.14159 + 1.5))
                if ((((1e10 / (float32(((float64(((1.5 + float32((v19 - v19))) / (v5 * v5))) * (v19 + (((float64(float32(float64(3.14159))) + v19) + ((v19 * (((0.0 * (1.5 / -2.25)) / 1.0) + v19)) + v19)) / (1.5 / (((0.0 + v19) + (0.1 * v19)) * (v19 - ((v19 + (v19 - v19)) / (v19 - v19)))))))) - v19)) / v5)) - (-2.25 - -2.25)) + 0.1) < (1.5 - (((0.0 + v5) * ((((-2.25 * (v5 / ((v5 * v5) / float32(v19)))) - v5) * 1e10) / v5)) / v5))) {
                    v20 := GenS3{F0: v18}
                    v10 = genFn0(139, -1, v11)
                    _ = v20
                }
                _ = v18
                _ = v19
            }
        }
    default:
        fmt.Println(v10, v10, v5)
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        switch 9223372036854775807 {
        case 4294967295:
            var v22 int8 = -128
            _ = v22
        case 9223372036854775807:
            switch (v11 + v11) {
            case "":
                for fi23 := 0; fi23 < 15; fi23++ {
                    v24 := []int{fi21, fi23, -870, (fi23 / (((fi23 ^ fi23)) | 1))}
                    v25 := GenS4{F0: v3}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case "\n\t":
                if true {
                    v5 -= (v5 + 0.0)
                    for fi26 := 0; fi26 < 9; fi26++ {
                        fmt.Println(v5)
                        if (true && (!(!(!(false || true))))) {
                            if false {
                                fmt.Println(v12, v12, v11)
                            }
                            v27 := []int{(-(fi26)), fi26, fi26}
                            _ = v27
                        }
                        _ = fi26
                    }
                }
            }
        case 707:
            v28 := GenS4{F0: v3}
            _ = v28
        default:
            v12 = -1
        }
        _ = fi21
    }
    switch (("a" + (v11 + ((v11 + "hello") + (((v11 + (v11 + v11)) + v11) + (v11 + (v11 + ("" + v11))))))) + v11) {
    case "foo bar":
        v29 := GenS5{F0: "hello", F1: v3}
        _ = v29
    case "a":
        if (1e10 < (1.5 / (((-2.25 / (float64(float32((3.14159 + 1e10))) - (1e10 * (0.1 * 3.14159)))) + 3.14159) / 0.1))) {
            fmt.Println(v11, v5, v12)
            if true {
                fmt.Println(v12, v3, v10)
            } else {
                _ = len(v1)
            }
        }
    case "hello":
        if (2147483648 >= 2) {
            v30 := GenS6{F0: 42}
            _ = v30
        } else {
            v12 = v12
            v31 := GenS6{F0: 2147483647}
            var v32 int16 = (v12 >> ((255) & 63))
            _ = v31
            _ = v32
        }
    default:
        var v33 int64 = 827
        _ = v33
    }
    var v34 bool = (0 == -2147483648)
    var v35 uint32 = 4294967295
    if true {
        for fi36 := 0; fi36 < 10; fi36++ {
            var v37 string = "hello"
            for fi38 := 0; fi38 < 9; fi38++ {
                fmt.Println(v37, v11, fi36)
                v12 -= (-(v12))
                v37 = v11
                _ = fi38
            }
            _ = fi36
            _ = v37
        }
        v5 -= (-2.25 * 0.0)
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        if (v34 && v34) {
            if true {
                if false {
                    var v40 int8 = 42
                    var v41 uint16 = 1
                    _ = v40
                    _ = v41
                } else {
                    var v42 int8 = 1
                    v3 = v3
                    v34 = ((0.1 * (0.1 - (v5 - v5))) < ((((v5 - ((v5 - -2.25) * ((0.0 - 0.0) + ((v5 * (3.14159 + ((((float32((1.5 / (((-2.25 * float64(v5)) * 1e10) * 0.0))) + -2.25) - v5) * v5) + float32(float64(v5))))) / 1e10)))) * (v5 - ((v5 - v5) - v5))) * (0.1 / v5)) / v5))
                    _ = v42
                }
                fmt.Println(v5)
                for fi43 := 0; fi43 < 2; fi43++ {
                    v12 = 125
                    var v44 uint32 = v35
                    var v45 uint16 = v3
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            } else {
                fmt.Println(v12, v5, v35)
                for fi46 := 0; fi46 < 13; fi46++ {
                    switch fi39 {
                    case 128:
                        v47 := GenS7{F0: v11, F1: 244}
                        _ = v47
                    case -396:
                        var v48 float32 = v5
                        _ = v48
                    case 2147483647:
                        switch fi46 {
                        case -9223372036854775808:
                            var v49 bool = (int64(fi46) == 1)
                            var v50 bool = v49
                            _ = v49
                            _ = v50
                        case -1:
                            fmt.Println(fi46, v10)
                        case 0:
                            _ = len(v1)
                        }
                    default:
                        v34 = v34
                    }
                    v35 += (v35 % ((v35) | 1))
                    _ = fi46
                }
                v3 = (v3 / ((uint16(v35)) | 1))
            }
        } else {
            fmt.Println(v10, v12)
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
