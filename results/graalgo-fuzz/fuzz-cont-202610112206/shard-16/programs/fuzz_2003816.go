package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 int16 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        if true {
            var v7 int64 = -1
            _ = v7
        }
        _ = fi6
    }
    switch -913 {
    case -901:
        switch -473 {
        case 32768:
            v8 := GenS1{F0: "\n\t", F1: 0}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 float32 = (3.14159 * -2.25)
            v12 := GenS1{F0: "foo bar", F1: 255}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(fi9)
        switch ("hello" + "a") {
        case "a":
            if (false && false) {
                switch fi9 {
                case 128:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 uint = 506
                        v14 = v14
                        v15 := GenS1{F0: ("foo bar" + "\n\t"), F1: 83}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    if true {
                        v16 := []int{42, fi9, -668, -2147483648}
                        _ = v16
                    }
                }
                var v17 float64 = 3.14159
                _ = v17
            }
        case "":
            if (164 > 42) {
                v18 := []int{-1, 4294967295}
                _ = v18
            } else {
                if false {
                    var v19 int = fi9
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 int = (-(fi9))
                        var v22 string = ("\n\t" + "\n\t")
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v19 = fi9
                    _ = v19
                } else {
                    if ((3.14159 + float32(1e10)) == 1e10) {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            if false {
                                var v24 int16 = 42
                                fmt.Println(v24)
                                _ = v24
                            } else {
                                var v25 int32 = 57
                                var v26 int8 = -128
                                v26 = (v26 + (v26 + v26))
                                _ = v25
                                _ = v26
                            }
                            var v27 float64 = 0.0
                            _ = fi23
                            _ = v27
                        }
                        var v28 int16 = 32767
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 18; fi29++ {
                            fmt.Println(fi9, fi29)
                            v30 := []int{}
                            v31 := GenS2{F0: 929, F1: 32767}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        var v32 int64 = 42
                        if true {
                            if true {
                                v33 := []int{}
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    var v35 uint32 = 1
                                    _ = fi34
                                    _ = v35
                                }
                                v36 := GenS3{F0: (v32 - (v32 / ((-310) | 1))), F1: ("hello" + ("\n\t" + (("foo bar" + "") + ""))), F2: -193}
                                _ = v33
                                _ = v36
                            } else {
                                fmt.Println(v32, v32)
                                v32 += v32
                            }
                            if (v32 <= -265) {
                                if true {
                                    var v37 uint32 = 2147483647
                                    _ = v37
                                }
                                v32 = int64(fi9)
                                v38 := []int{(-(fi9)), (fi9 & -1)}
                                _ = v38
                            }
                            v32 += (v32 % ((-436) | 1))
                        } else {
                            if (!false) {
                                v39 := GenS3{F0: v32, F1: "foo bar", F2: 0}
                                _ = v39
                            }
                        }
                        _ = v32
                    }
                    v40 := []int{fi9}
                    _ = v40
                }
            }
        }
        _ = fi9
    }
    var v41 uint8 = 2
    _ = v41
    return 58
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{2}
    var v3 float64 = (float64(v1) / 0.1)
    var v4 uint64 = 906
    var v5 int = 699
    v2 = append(v2, 1)
    v1 += 0.1
    for fi42 := 0; fi42 < 15; fi42++ {
        for fi43 := 0; fi43 < 4; fi43++ {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "hello":
                v4 = v4
            case "":
                v2 = append(v2, 2)
            case "\n\t":
                v3 -= ((3.14159 + ((v3 - v3) + v3)) + 0.0)
            default:
                fmt.Println(v4, v1)
            }
            _ = fi43
        }
        _ = fi42
    }
    v5 = (v5 | v5)
    var v44 float32 = (v1 + ((v1 * ((3.14159 - (((v1 * ((float32(v3) + (((-2.25 * 3.14159) * 1e10) / 1e10)) - (((float32((v3 * ((v3 - (float64(v1) * (v3 * (((((1.5 / ((0.0 + float64(v1)) / (1.5 * v3))) + float64(0.0)) * float64(float32(v3))) - 0.0) - v3)))) * v3))) / v1) - (v1 * v1)) / v1))) + ((((v1 * v1) / 1e10) / v1) / v1)) / 3.14159)) + v1)) / (v1 / v1)))
    var v45 uint = 18446744073709551615
    v1 += v1
    v3 += (v3 + v3)
    fmt.Println(v3, v44)
    v4 -= (v4 % (((v4 | 912)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
