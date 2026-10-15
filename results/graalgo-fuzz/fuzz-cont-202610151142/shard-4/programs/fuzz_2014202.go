package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 65535
    var v4 uint64 = 2
    switch 127 {
    case 9223372036854775807:
        v5 := []int{42, -190}
        _ = v5
    case 335:
        v2 += (v2 | (^(v4)))
    default:
        var v6 int64 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{fi7}
        var v9 int8 = 1
        switch ("hello" + "a") {
        case "a":
            v9 = v9
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = ""
    if false {
        v11 := GenS1{F0: -488, F1: -1, F2: v2}
        _ = v11
    } else {
        var v12 int64 = 42
        v3 -= v3
        if ((v12 / (((v12 >> ((255) & 63))) | 1)) >= v12) {
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS2{F0: ((!((190 >= 0) && true)) && true)}
                switch fi13 {
                case 4294967295:
                    if (false || true) {
                        if (143 > 857) {
                            var v15 int16 = -1
                            v2 = 474
                            v16 := GenS1{F0: v12, F1: 4294967295, F2: (^(v4))}
                            _ = v15
                            _ = v16
                        } else {
                            var v17 int8 = 127
                            var v18 uint8 = 202
                            for fi19 := 0; fi19 < 10; fi19++ {
                                switch (fi19 | fi19) {
                                case -2147483648:
                                    switch v10 {
                                    case "\n\t":
                                        v18 = v18
                                    case "foo bar":
                                        var v20 int = -9223372036854775808
                                        _ = v20
                                    }
                                }
                                if true {
                                    if false {
                                        v4 -= 42
                                        switch v10 {
                                        case "a":
                                            v4 = (v4 >> ((111) & 63))
                                        }
                                    } else {
                                        fmt.Println(v2, v3, v2)
                                    }
                                    v4 = (v2 * v4)
                                    v21 := GenS3{F0: (fi19 - fi13), F1: (v2 * v4)}
                                    var v22 uint = 786
                                    _ = v21
                                    _ = v22
                                }
                                var v23 float32 = ((3.14159 * 1e10) - 1e10)
                                _ = fi19
                                _ = v23
                            }
                            _ = v17
                            _ = v18
                        }
                        fmt.Println(v2, v12)
                    }
                case 9223372036854775807:
                    v24 := GenS4{F0: 0, F1: -1}
                    _ = v24
                case 0:
                    for fi25 := 0; fi25 < 8; fi25++ {
                        fmt.Println(fi25)
                        fmt.Println(v12, fi13, fi25)
                        if true {
                            var v26 float32 = float32(1.5)
                            if (fi13 < fi25) {
                                v27 := GenS2{F0: false}
                                v28 := GenS5{F0: false}
                                fmt.Println(fi25, fi13, v12)
                                v26 -= (v26 + ((float32(float64(v26)) - v26) / (v26 - (float32((-2.25 - 0.1)) - v26))))
                                fmt.Println(v26, fi25, v3)
                                switch fi13 {
                                case 127:
                                    var v29 string = ("a" + v10)
                                    _ = v29
                                case -670:
                                    switch "\n\t" {
                                    case "a":
                                        switch (("a" + "\n\t") + v10) {
                                        case "":
                                            v26 = (float32(1e10) - ((float32(0.1) / 3.14159) * 0.1))
                                        case "foo bar":
                                            v12 += v12
                                        default:
                                            v10 = "a"
                                        }
                                        fmt.Println(fi13, v26, fi25)
                                        fmt.Println(v3, fi25, fi13)
                                    }
                                case -108:
                                    var v30 float32 = 0.0
                                    _ = v30
                                default:
                                    var v31 float32 = (float32((((float64(0.1) * 3.14159) * float64(float32(float64(3.14159)))) + (1e10 + (((float64(1e10) - (3.14159 - (0.0 - 3.14159))) / 1.0) * 1e10)))) + v26)
                                    _ = v31
                                }
                                var v32 int64 = (v12 >> ((18446744073709551615) & 63))
                                _ = v27
                                _ = v28
                                _ = v32
                            } else {
                                v33 := GenS2{F0: (!true)}
                                if false {
                                    var v34 int16 = 0
                                    var v35 uint16 = (^(v3))
                                    switch "\n\t" {
                                    case "":
                                        v34 = v34
                                    default:
                                        switch v10 {
                                        case "hello":
                                            if ((((-2.25 * 0.1) - (((1e10 + v26) + float32((1e10 - 0.0))) - v26)) + -2.25) < v26) {
                                                var v36 uint = 1
                                                _ = v36
                                            } else {
                                                v2 = 2
                                                var v37 string = v10
                                                _ = v37
                                            }
                                            v35 = (v3 + 0)
                                            var v38 int64 = int64(v35)
                                            _ = v38
                                        }
                                    }
                                    fmt.Println(v34, fi25, v26)
                                    var v39 int16 = (-(v34))
                                    _ = v34
                                    _ = v35
                                    _ = v39
                                } else {
                                    delete(v1, (((v10 + (("a" + (v10 + "hello")) + "a")) + v10) + "\n\t"))
                                    v40 := GenS3{F0: -916, F1: v2}
                                    v41 := GenS1{F0: (v12 * (v12 & 47)), F1: 32767, F2: 1}
                                    v10 = (((v10 + v10) + "\n\t") + "")
                                    _ = v40
                                    _ = v41
                                }
                                v42 := GenS1{F0: (v12 * 9223372036854775807), F1: (fi13 + fi13), F2: (v2 << ((448) & 63))}
                                _ = v33
                                _ = v42
                            }
                            _ = v26
                        }
                        if (!false) {
                            if (0.1 > 0.1) {
                                v3 = v3
                            }
                        }
                        _ = fi25
                    }
                }
                for fi43 := 0; fi43 < 3; fi43++ {
                    v4 += (v4 >> ((0) & 63))
                    if (((1.5 + 1.5) + float32((1e10 * (1e10 + float64(-2.25))))) != (1.5 - (float32(3.14159) + ((-2.25 * float32(-2.25)) / 1e10)))) {
                        fmt.Println(v2)
                    } else {
                        v3 = v3
                    }
                    var v44 float32 = 1e10
                    _ = fi43
                    _ = v44
                }
                _ = fi13
                _ = v14
            }
        } else {
            if (1 != 0) {
                for fi45 := 0; fi45 < 1; fi45++ {
                    v46 := []int{(fi45 * fi45), fi45}
                    _ = fi45
                    _ = v46
                }
                var v47 uint8 = 127
                var v48 float32 = (1.5 + (0.0 + (((((-2.25 - 0.0) - 3.14159) / 1.5) - (0.0 / 1e10)) * 0.0)))
                _ = v47
                _ = v48
            } else {
                var v49 float32 = (-2.25 + 1e10)
                v3 += (v3 * uint16(v2))
                var v50 uint8 = 2
                _ = v49
                _ = v50
            }
        }
        _ = v12
    }
    var v51 float32 = 1.5
    for fi52 := 0; fi52 < 8; fi52++ {
        v1[v10] = 42
        _ = fi52
    }
    var v53 int64 = 9223372036854775807
    v54 := []int{}
    for fi55 := 0; fi55 < 17; fi55++ {
        switch "a" {
        case "":
            var v56 uint16 = (v3 % (((v3 % (((^(v3))) | 1))) | 1))
            _ = v56
        case "\n\t":
            v4 += (v4 * 54)
        }
        _ = fi55
    }
    v2 += (v2 / (((^(v2))) | 1))
    v54 = append(v54, 127)
    v3 = (v3 * v3)
    if len(v54) > 0 {
        v54[0] = 2147483647
    }
    fmt.Println(v51)
    v3 = (v3 * v3)
    var v57 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v53
    _ = v54
    _ = v57
}
