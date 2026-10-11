package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16; F2 int }
type GenS4 struct { F0 int32; F1 int64; F2 uint }
type GenS5 struct { F0 int8; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint32) int16 {
    v6 := GenS1{F0: 255, F1: 255, F2: uint8(p1)}
    v7 := GenS2{F0: false}
    var v8 float32 = (-2.25 + 3.14159)
    if (!true) {
        var v9 uint64 = p0
        var v10 int8 = 100
        _ = v10
        _ = v9
    } else {
        var v11 uint = 69
        if true {
            fmt.Println(v8, v8, p0)
        } else {
            for fi12 := 0; fi12 < 7; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    p0 += p0
                    fmt.Println(p1)
                    _ = fi13
                }
                var v14 int8 = 43
                var v15 float64 = ((0.0 + float64(0.0)) - (float64(v8) + (((0.0 - (3.14159 * ((((1e10 / 3.14159) + (3.14159 / float64(0.1))) * (0.0 / 1.5)) + ((3.14159 / 1.0) + 1.5)))) * -2.25) - 1.5)))
                _ = fi12
                _ = v14
                _ = v15
            }
            if (!false) {
                if (1 == 1) {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 int32 = 42
                        v18 := GenS3{F0: fi16, F1: 32767, F2: fi16}
                        v11 += v11
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if false {
                        v11 += v11
                    }
                    v19 := []int{2147483647, -589, -236, 2147483647}
                    _ = v19
                } else {
                    switch "hello" {
                    case "":
                        var v20 float32 = v8
                        _ = v20
                    case "hello":
                        for fi21 := 0; fi21 < 14; fi21++ {
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 string = "a"
                                v11 = (v11 * v11)
                                var v24 string = ((v23 + "foo bar") + "foo bar")
                                if (-32768 >= -1) {
                                    v11 -= (v11 ^ (v11 ^ 411))
                                } else {
                                    v25 := GenS3{F0: (fi21 * -717), F1: uint16(v11), F2: fi21}
                                    fmt.Println(p1, v8, v24)
                                    var v26 string = "foo bar"
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(v8)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            var v27 uint32 = p1
                            _ = fi21
                            _ = v27
                        }
                    case "\n\t":
                        v8 = v8
                    }
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v29 := []int{fi28, (fi28 / (((fi28 >> (((v11 / ((v11) | 1))) & 63))) | 1)), fi28, (fi28 ^ -32)}
                        _ = fi28
                        _ = v29
                    }
                    v11 = v11
                }
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                switch (("a" + "a") + "a") {
                case "foo bar":
                    var v31 int64 = 0
                    _ = v31
                case "hello":
                    v32 := []int{fi30, -9223372036854775808, 2147483647}
                    _ = v32
                case "":
                    var v33 int64 = 42
                    _ = v33
                }
                v34 := GenS3{F0: int(v11), F1: 0, F2: 128}
                _ = fi30
                _ = v34
            }
        }
        var v35 uint8 = 255
        _ = v11
        _ = v35
    }
    var v36 int16 = 565
    _ = p0
    _ = p1
    _ = v36
    _ = v6
    _ = v7
    _ = v8
    return v36
}
func genFn1(p0 int) string {
    v37 := []int{1}
    p0 -= p0
    var v38 int64 = 0
    p0 = 0
    _ = p0
    _ = v37
    _ = v38
    return ("" + ((("" + "\n\t") + ("\n\t" + "foo bar")) + "foo bar"))
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{-1, 2147483647, -331, 4294967295}
    var v3 int8 = -1
    var v4 float32 = 3.14159
    v5 := make(map[uint32]bool)
    var v39 uint8 = 1
    switch "\n\t" {
    case "\n\t":
        v40 := GenS1{F0: 128, F1: 32767, F2: (v39 % (((v39 * v39)) | 1))}
        _ = v40
    case "":
        v39 += 252
    case "a":
        var v41 int8 = (v3 >> ((10) & 63))
        _ = v41
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 int16 = genFn0(42, 0)
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    v4 -= -2.25
    var v45 int64 = -931
    for fi46 := 0; fi46 < 9; fi46++ {
        if ((v4 != 1.5) && true) {
            for fi47 := 0; fi47 < 18; fi47++ {
                if false {
                    fmt.Println(fi47, v1)
                    var v48 float32 = 0.0
                    _ = v48
                }
                _ = fi47
            }
        }
        v45 = 1
        v49 := GenS3{F0: -1, F1: 65535, F2: (^(fi46))}
        _ = fi46
        _ = v49
    }
    fmt.Println(v45)
    fmt.Println(v39)
    v4 = 3.14159
    v50 := GenS4{F0: 51, F1: v45, F2: 4294967295}
    v4 = (1e10 * 0.0)
    for fi51 := 0; fi51 < 1; fi51++ {
        v52 := GenS5{F0: v3, F1: 1, F2: (-(v3))}
        switch ("" + "a") {
        case "a":
            v53 := v5[541]
            _ = v53
        case "foo bar":
            v54 := GenS6{F0: 2147483648}
            _ = v54
        case "":
            v39 -= v39
        default:
            if (fi51 <= fi51) {
                v39 -= v39
                for fi55 := 0; fi55 < 6; fi55++ {
                    delete(v5, 479)
                    v56 := GenS1{F0: 42, F1: 32767, F2: v39}
                    v4 = float32(0.0)
                    _ = fi55
                    _ = v56
                }
            } else {
                v57 := GenS1{F0: (v39 + (v39 + v39)), F1: 1, F2: 1}
                for fi58 := 0; fi58 < 2; fi58++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v59 uint8 = v39
                        var v60 uint = 613
                        if (false || (!(!true))) {
                            var v61 uint = 590
                            fmt.Println(v39, v61)
                            var v62 uint32 = 2147483647
                            _ = v61
                            _ = v62
                        }
                        _ = v59
                        _ = v60
                    }
                    _ = fi58
                }
                v4 -= ((0.0 + float32(1e10)) - 3.14159)
                _ = v57
            }
        }
        _ = fi51
        _ = v52
    }
    switch "a" {
    case "foo bar":
        var v63 float32 = v4
        _ = v63
    case "\n\t":
        v64 := GenS4{F0: -673, F1: v45, F2: 212}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v50
}
