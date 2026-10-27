package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0() uint32 {
    var v6 string = "foo bar"
    v7 := GenS1{F0: -1, F1: -6}
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint {
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 float32 = (((3.14159 / -2.25) + ((((((3.14159 / 1.5) + 1e10) - 1e10) - (float32((1.5 - (0.0 * (1.5 * float64(-2.25))))) - 1.5)) + 1e10) + (0.1 * 1e10))) + 1e10)
        var v10 string = "foo bar"
        var v11 bool = true
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint64 = 485
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: true}
        if false {
            v15 := []int{(fi13 | (fi13 ^ 987)), -9223372036854775808, -394, (fi13 * 32768)}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            if (42 >= 1) {
                if ((0.0 >= 1.5) && (!false)) {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 int8 = -128
                        fmt.Println(fi16, fi16, fi16)
                        if true {
                            var v18 int16 = 42
                            _ = v18
                        } else {
                            v12 += v12
                        }
                        v12 += v12
                        _ = fi16
                        _ = v17
                    }
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 int8 = 36
                        var v21 float64 = 1e10
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v12, v12)
                } else {
                    fmt.Println(v12, fi13)
                    if false {
                        var v22 float64 = float64(float32((float64((1e10 * (3.14159 / (float32(0.1) * 3.14159)))) * (0.1 - 0.0))))
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 16; fi23++ {
                    if true {
                        if ((true || (6 < -1)) && (2147483648 == 2147483647)) {
                            v12 += v12
                            fmt.Println(v12, fi13)
                        }
                        fmt.Println(fi23)
                    }
                    _ = fi23
                }
            } else {
                if len(v15) > 0 {
                    v15[0] = int(v12)
                }
                if ((v12 - 0) < (v12 % (((v12 | (^(v12)))) | 1))) {
                    v24 := GenS1{F0: 1000, F1: -128}
                    _ = v24
                }
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    v12 -= 311
    var v25 int = -2147483648
    _ = v12
    _ = v25
    return 18446744073709551615
}

func main() {
    var v1 float64 = -2.25
    var v2 int16 = -32768
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    for fi26 := 0; fi26 < 4; fi26++ {
        v5 = v5
        switch ((v5 + ("" + (v5 + v5))) + "foo bar") {
        case "foo bar":
            v27 := GenS3{F0: 42, F1: "\n\t"}
            _ = v27
        case "a":
            v28 := GenS1{F0: (v2 ^ (v2 / (((v2 ^ v2)) | 1))), F1: -39}
            _ = v28
        case "\n\t":
            if (v2 >= -1) {
                if true {
                    v5 = "hello"
                    fmt.Println(fi26, v5)
                } else {
                    v5 = v5
                    v29 := []int{}
                    fmt.Println(v5)
                    _ = v29
                }
            }
        }
        _ = fi26
    }
    var v30 float64 = float64((1e10 + 3.14159))
    v5 = ("foo bar" + ("foo bar" + ("foo bar" + v5)))
    v4[-516] = 2
    var v31 int32 = -1
    v3[-191] = (((("hello" + v5) + v5) + "\n\t") + "a")
    v32 := GenS2{F0: (100 > -128)}
    v30 -= float64((0.1 - 1e10))
    v33 := []int{55, 9223372036854775807, 0}
    if false {
        for fi34 := 0; fi34 < 3; fi34++ {
            v1 = (v30 + v1)
            v35 := GenS3{F0: 205, F1: v5}
            v30 -= 0.0
            _ = fi34
            _ = v35
        }
        var v36 string = v5
        var v37 int32 = 1
        _ = v36
        _ = v37
    }
    v2 = v2
    var v38 int32 = v31
    if (float32(v30) == (0.0 + (-2.25 / 1.0))) {
        var v39 uint = genFn1()
        _ = v39
    }
    v33 = append(v33, -9223372036854775808)
    switch (v5 + "") {
    case "":
        for fi40 := 0; fi40 < 15; fi40++ {
            v41 := GenS2{F0: (v2 == 222)}
            _ = fi40
            _ = v41
        }
    case "foo bar":
        v2 -= 1
    }
    switch "\n\t" {
    case "a":
        for fi42 := 0; fi42 < 16; fi42++ {
            if true {
                delete(v3, (fi42 * fi42))
                var v43 uint8 = 42
                var v44 uint32 = 42
                _ = v43
                _ = v44
            }
            v1 = v30
            if (v2 != (v2 ^ v2)) {
                fmt.Println(v5, v2)
                for fi45 := 0; fi45 < 11; fi45++ {
                    v46 := v4[527]
                    var v47 float64 = v30
                    v48 := GenS1{F0: (v2 - int16(v46)), F1: 114}
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
                v49 := GenS4{F0: (v31 * v38)}
                _ = v49
            }
            _ = fi42
        }
    case "foo bar":
        v5 = v5
    case "hello":
        var v50 int = 83
        _ = v50
    }
    var v51 int16 = (v2 ^ (v2 & -32768))
    if (true || true) {
        for fi52 := 0; fi52 < 1; fi52++ {
            for fi53 := 0; fi53 < 8; fi53++ {
                v54 := GenS5{F0: 0}
                _ = fi53
                _ = v54
            }
            _ = fi52
        }
        v1 = (0.1 / v30)
        v55 := GenS3{F0: 42, F1: "a"}
        _ = v55
    } else {
        fmt.Println(v1)
        switch (v5 + ((v5 + (v5 + (v5 + (((v5 + "\n\t") + "a") + "\n\t")))) + v5)) {
        case "":
            v2 -= (v2 >> ((898) & 63))
        case "a":
            var v56 int32 = v38
            _ = v56
        case "foo bar":
            var v57 uint64 = 497
            _ = v57
        default:
            for fi58 := 0; fi58 < 12; fi58++ {
                delete(v3, fi58)
                v59 := GenS6{F0: 0, F1: -128, F2: v31}
                _ = fi58
                _ = v59
            }
        }
        if true {
            var v60 uint32 = 42
            v30 += float64(-2.25)
            fmt.Println(v5, v51)
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v51
}
