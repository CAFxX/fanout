package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 int; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint16; F1 int8; F2 bool }
type GenS8 struct { F0 int32; F1 int16 }
func genFn0(p0 int64) int32 {
    v6 := GenS1{F0: 0, F1: -240}
    v7 := []int{-9223372036854775808, 2147483647}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 151, F1: 1000}
        var v10 string = "\n\t"
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool) bool {
    fmt.Println(p0)
    p0 = (!true)
    var v11 float32 = ((0.1 * ((1e10 + 0.1) + (0.0 - 1e10))) / 3.14159)
    _ = p0
    _ = v11
    return p0
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 3.14159
    var v3 int = -860
    var v4 int64 = 42
    var v5 float64 = ((3.14159 - (v2 / ((1e10 + float64(float32(float64(((float32(1.5) + 3.14159) / (1.5 - 0.0)))))) / ((v2 - v2) * float64((0.1 / 0.1)))))) + v2)
    v5 -= v2
    var v12 string = v1
    var v13 float64 = 1.5
    if true {
        var v14 uint64 = 892
        for fi15 := 0; fi15 < 17; fi15++ {
            v12 = v12
            _ = fi15
        }
        _ = v14
    } else {
        v16 := []int{}
        var v17 string = v1
        _ = v16
        _ = v17
    }
    fmt.Println(v12, v2)
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := []int{(-(v3)), (v3 / ((40) | 1))}
        if true {
            if (!true) {
                var v20 int64 = 9223372036854775807
                _ = v20
            }
        } else {
            if (!(0 == -1)) {
                if false {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS2{F0: 4294967295, F1: 0}
                        _ = fi21
                        _ = v22
                    }
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v12 = ("foo bar" + (v12 + v1))
                    case "":
                        switch 340 {
                        case 128:
                            var v23 int32 = 1
                            _ = v23
                        default:
                            var v24 int = v3
                            _ = v24
                        }
                    case "\n\t":
                        switch (v3 % ((v3) | 1)) {
                        case -2147483648:
                            var v25 uint32 = 4294967295
                            _ = v25
                        default:
                            if len(v19) > 0 {
                                v19[0] = (v3 - fi18)
                            }
                        }
                    }
                    v19 = append(v19, (fi18 - -298))
                }
                _ = len(v19)
                v4 -= 781
            } else {
                fmt.Println(v5)
                _ = len(v19)
            }
            var v26 int8 = 42
            for fi27 := 0; fi27 < 6; fi27++ {
                if ((1.5 + (float32(-2.25) / -2.25)) <= ((1.5 + float32(0.0)) - 0.0)) {
                    switch -926 {
                    case 128:
                        v28 := GenS3{F0: 65535, F1: (fi18 / (((v3 << ((0) & 63))) | 1)), F2: fi27}
                        _ = v28
                    case 32767:
                        fmt.Println(v26, v2, v3)
                    case 1:
                        var v29 float64 = ((v2 * v5) * v13)
                        v30 := GenS4{F0: v4}
                        _ = v29
                        _ = v30
                    default:
                        v1 = (("foo bar" + ("a" + v12)) + (((v12 + "a") + v1) + (("" + ("a" + ("hello" + (v12 + (("foo bar" + v1) + ""))))) + "a")))
                    }
                    _ = len(v19)
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                } else {
                    v31 := GenS1{F0: 1, F1: 100}
                    _ = len(v19)
                    _ = v31
                }
                _ = fi27
            }
            _ = v26
        }
        _ = fi18
        _ = v19
    }
    switch v12 {
    case "":
        var v32 int = v3
        _ = v32
    case "a":
        var v33 int = v3
        _ = v33
    case "foo bar":
        if true {
            v34 := GenS3{F0: 519, F1: (v3 - 42), F2: v3}
            v35 := []int{4294967295, 305, 127, 9223372036854775807}
            var v36 float64 = ((((3.14159 - 1.5) / (v5 * (v13 * (v2 / v13)))) / v5) * v2)
            _ = v34
            _ = v35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        if ((-1 < -291) && true) {
            v38 := GenS5{F0: 1}
            var v39 uint32 = 2147483647
            _ = v38
            _ = v39
        } else {
            switch (v3 ^ v3) {
            case 50:
                var v40 float32 = 0.1
                _ = v40
            default:
                v41 := GenS6{F0: -29, F1: 0}
                _ = v41
            }
        }
        v42 := []int{}
        _ = fi37
        _ = v42
    }
    v4 += v4
    v12 = (v12 + v1)
    if false {
        for fi43 := 0; fi43 < 2; fi43++ {
            v44 := []int{fi43, (v3 + fi43), v3}
            v45 := GenS7{F0: 427, F1: 127, F2: (!false)}
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    var v46 int8 = -66
    for fi47 := 0; fi47 < 20; fi47++ {
        v48 := []int{-9223372036854775808, fi47}
        var v49 int64 = v4
        _ = fi47
        _ = v48
        _ = v49
    }
    v50 := []int{(v3 | (v3 << ((1) & 63))), 422}
    if false {
        v51 := GenS8{F0: 42, F1: 675}
        var v52 float32 = 1e10
        _ = v51
        _ = v52
    } else {
        if (v1 > ("\n\t" + v1)) {
            var v53 bool = (!false)
            if len(v50) > 0 {
                v50[0] = (v3 | (v3 ^ v3))
            }
            fmt.Println(v4)
            _ = v53
        }
        v1 = (v1 + (v12 + "a"))
    }
    fmt.Println(v4)
    v4 = 485
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
}
