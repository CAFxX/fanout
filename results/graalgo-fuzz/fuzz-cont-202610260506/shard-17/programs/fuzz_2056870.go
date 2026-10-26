package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 int) int {
    var v6 int64 = 423
    var v7 uint64 = 0
    p0 = (p0 % ((p0) | 1))
    var v8 float64 = (3.14159 + 1e10)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (-(p0))
}
func genFn1(p0 uint8) uint32 {
    p0 = 1
    v9 := GenS1{F0: -457, F1: p0, F2: (241 >= 0)}
    var v10 int32 = -2147483648
    fmt.Println(p0)
    switch "hello" {
    case "\n\t":
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                if false {
                    v13 := []int{(fi12 / (((fi11 / ((fi11) | 1))) | 1)), 2147483647, fi12, genFn0(fi11)}
                    v13 = append(v13, (-(fi12)))
                    _ = v13
                }
                _ = fi12
            }
            v14 := GenS1{F0: 24, F1: p0, F2: true}
            _ = fi11
            _ = v14
        }
    case "foo bar":
        p0 = 255
    default:
        p0 -= 160
    }
    _ = p0
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    var v3 float32 = (1.5 / 1.0)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    switch "a" {
    case "\n\t":
        var v15 uint64 = 18446744073709551615
        _ = v15
    case "a":
        var v16 bool = (!((0 <= -128) && (!(genFn0(-805) == 9223372036854775807))))
        _ = v16
    }
    v17 := GenS2{F0: -128}
    fmt.Println(v1)
    var v18 uint8 = 11
    v19 := []int{}
    fmt.Println(v3)
    for fi20 := 0; fi20 < 19; fi20++ {
        if true {
            if (-2.25 >= 1e10) {
                v3 -= 3.14159
                for fi21 := 0; fi21 < 9; fi21++ {
                    v18 = (^(v18))
                    v22 := GenS2{F0: 45}
                    fmt.Println(v1, fi21, v18)
                    _ = fi21
                    _ = v22
                }
                v23 := GenS3{F0: -903}
                _ = v23
            } else {
                fmt.Println(v18, v3)
                switch "a" {
                case "foo bar":
                    if ((fi20 % ((fi20) | 1)) != (fi20 - genFn0((-(fi20))))) {
                        var v24 bool = (662 == 775)
                        v25 := GenS4{F0: 127, F1: "a"}
                        for fi26 := 0; fi26 < 20; fi26++ {
                            switch (^(fi26)) {
                            case 1:
                                if v24 {
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v1)
                                    case "a":
                                        if ((3.14159 > 0.1) && true) {
                                            switch (fi20 << ((273) & 63)) {
                                            case 1:
                                                v27 := GenS4{F0: v18, F1: "\n\t"}
                                                _ = v27
                                            case 922:
                                                switch fi20 {
                                                case 42:
                                                    fmt.Println(v18, v3)
                                                case 32767:
                                                    if len(v19) > 0 {
                                                        v19[0] = int(v18)
                                                    }
                                                case -1:
                                                    v28 := GenS1{F0: 0, F1: (v18 ^ (^(v18))), F2: true}
                                                    v18 += v18
                                                    _ = v28
                                                default:
                                                    var v29 int32 = -2147483648
                                                    _ = v29
                                                }
                                            }
                                            v30 := GenS2{F0: 52}
                                            var v31 uint16 = 121
                                            v31 = (v31 & 448)
                                            fmt.Println(v31)
                                            v32 := GenS3{F0: -1}
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        } else {
                                            var v33 int8 = 42
                                            _ = v33
                                        }
                                    default:
                                        v1 -= v1
                                    }
                                    if len(v19) > 0 {
                                        v19[0] = (^(fi20))
                                    }
                                    if v24 {
                                        var v34 uint8 = 83
                                        v35 := GenS5{F0: (fi26 % (((fi20 * 4294967295)) | 1))}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 int64 = 9223372036854775807
                                        fmt.Println(fi26)
                                        var v37 int = -940
                                        _ = v36
                                        _ = v37
                                    }
                                } else {
                                    v1 = 42
                                    var v38 bool = false
                                    v3 += (v3 + 0.1)
                                    _ = len(v4)
                                    v3 += (v3 / (v3 + float32((0.0 / float64(((v3 - v3) * v3))))))
                                    _ = v38
                                }
                            case -2147483648:
                                v24 = (208 <= 297)
                            case 0:
                                v39 := GenS4{F0: 48, F1: ((("" + "hello") + (("a" + "a") + "\n\t")) + "foo bar")}
                                _ = v39
                            }
                            _ = len(v19)
                            _ = fi26
                        }
                        _ = v24
                        _ = v25
                    }
                case "a":
                    v4[2147483647] = 0
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            }
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v40 := GenS6{F0: 32767, F1: ("\n\t" != "a"), F2: 32767}
            _ = v40
        } else {
            var v41 bool = (v1 > (^(v1)))
            if false {
                var v42 float64 = 0.1
                for fi43 := 0; fi43 < 16; fi43++ {
                    var v44 int64 = -864
                    _ = fi43
                    _ = v44
                }
                v45 := GenS6{F0: -1, F1: v41, F2: 1}
                _ = v42
                _ = v45
            }
            _ = v41
        }
        var v46 uint = 336
        var v47 int8 = -101
        _ = fi20
        _ = v46
        _ = v47
    }
    v48 := GenS6{F0: -618, F1: ((1.5 > 0.0) || true), F2: 255}
    var v49 uint64 = 136
    _ = len(v4)
    for fi50 := 0; fi50 < 7; fi50++ {
        v49 -= v49
        for fi51 := 0; fi51 < 19; fi51++ {
            if (true && true) {
                fmt.Println(fi51, v1, v3)
            }
            var v52 uint32 = 2147483648
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    v3 -= 3.14159
    var v53 int16 = -1
    v49 = v49
    v53 -= 1000
    v3 = 1e10
    v18 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v53
}
