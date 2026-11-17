package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64, p1 uint16, p2 int8) uint64 {
    p0 = (^(p0))
    for fi6 := 0; fi6 < 13; fi6++ {
        if ((-(fi6)) <= fi6) {
            v7 := []int{fi6}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                p1 -= (p1 % (((p1 & p1)) | 1))
                p2 -= p2
                _ = fi8
            }
            if true {
                p1 += 42
                v9 := GenS1{F0: p0, F1: -32768, F2: false}
                _ = v9
            } else {
                if false {
                    p1 = (p1 - (p1 - p1))
                } else {
                    if true {
                        var v10 int16 = 0
                        v11 := []int{fi6, (fi6 / ((fi6) | 1))}
                        _ = v10
                        _ = v11
                    } else {
                        switch int(p2) {
                        case 1:
                            for fi12 := 0; fi12 < 5; fi12++ {
                                if false {
                                    var v13 string = "foo bar"
                                    fmt.Println(fi6, p1)
                                    var v14 int16 = int16(p0)
                                    _ = v13
                                    _ = v14
                                }
                                v15 := GenS2{F0: ((4294967295 > 4294967295) || (42 >= 0))}
                                var v16 bool = true
                                _ = fi12
                                _ = v15
                                _ = v16
                            }
                        case 128:
                            for fi17 := 0; fi17 < 2; fi17++ {
                                var v18 uint16 = uint16(p0)
                                var v19 float64 = (3.14159 * 1e10)
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        case 42:
                            p2 += p2
                        default:
                            v20 := GenS2{F0: false}
                            _ = v20
                        }
                        v21 := GenS1{F0: (p0 - (p0 >> ((255) & 63))), F1: 42, F2: (((false || (!true)) && (!(42 == 255))) || true)}
                        v22 := GenS3{F0: (!true)}
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 9; fi23++ {
                        p2 = 104
                        _ = fi23
                    }
                }
                switch "" {
                case "a":
                    p2 += p2
                case "foo bar":
                    var v24 float64 = 0.0
                    _ = v24
                }
            }
            v25 := GenS3{F0: (p1 >= (p1 | (p1 >> ((421) & 63))))}
            _ = v25
        }
        _ = fi6
    }
    p0 -= 4294967295
    var v26 uint32 = 4294967295
    var v27 int32 = -10
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return p0
}
func genFn1(p0 int8, p1 bool) int16 {
    switch -1 {
    case 1:
        fmt.Println(p1, p0)
    case -1:
        if true {
            p0 += 126
            var v28 int8 = (-(p0))
            _ = v28
        }
    default:
        switch ("" + ("\n\t" + ("\n\t" + (((((("a" + "hello") + ("\n\t" + "foo bar")) + "hello") + "a") + "a") + "\n\t")))) {
        case "a":
            p1 = ((3.14159 == 0.0) && ((255 == 42) || p1))
        }
    }
    var v29 uint32 = 2
    var v30 uint32 = 42
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return -32768
}

func main() {
    var v1 uint = 792
    var v2 uint8 = 255
    var v3 int8 = 127
    var v4 float64 = 1.5
    var v5 float64 = v4
    v31 := GenS2{F0: (42 > 261)}
    if (2147483647 == 0) {
        if false {
            fmt.Println(v3)
            fmt.Println(v2, v4, v1)
            v1 -= (v1 - v1)
        } else {
            fmt.Println(v3, v3)
        }
        v32 := []int{}
        _ = v32
    }
    var v33 uint16 = 1
    if ((v3 / ((111) | 1)) >= 42) {
        v34 := []int{9223372036854775807, 2147483647, 32767}
        _ = v34
    } else {
        var v35 float64 = (0.0 * v4)
        fmt.Println(v3, v33, v33)
        _ = v35
    }
    v36 := GenS2{F0: ((v3 * (v3 % ((v3) | 1))) > (^(v3)))}
    v3 += (-(v3))
    switch ("hello" + "hello") {
    case "a":
        v37 := []int{-785, 128, 32768}
        _ = v37
    default:
        fmt.Println(v2, v2)
    }
    if (true && (true || false)) {
        for fi38 := 0; fi38 < 16; fi38++ {
            v39 := []int{(fi38 * (-(fi38))), -365, (fi38 ^ 32767)}
            _ = fi38
            _ = v39
        }
        var v40 uint32 = 412
        var v41 float64 = (((0.0 * 3.14159) / 1.0) + -2.25)
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: false}
    v3 -= -108
    v33 += (v33 << ((uint(v33)) & 63))
    var v43 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
