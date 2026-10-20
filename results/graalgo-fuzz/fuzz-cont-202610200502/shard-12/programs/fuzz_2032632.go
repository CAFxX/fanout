package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 uint8, p2 string) int32 {
    var v4 int = 42
    if false {
        var v5 uint = 2
        _ = v5
    } else {
        var v6 int = v4
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 0
}
func genFn1(p0 int, p1 string) int32 {
    if (!false) {
        v8 := []int{(-(p0)), (^(p0))}
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(p0, fi9, p1)
            fmt.Println(p1, p1, p1)
            if (3.14159 > (1.5 - (0.1 / 3.14159))) {
                var v10 uint16 = 2
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 bool = false
                    if ((fi9 / ((0) | 1)) >= fi9) {
                        var v13 uint64 = 2
                        _ = v13
                    }
                    p1 = "foo bar"
                    _ = fi11
                    _ = v12
                }
                v14 := GenS1{F0: 42}
                _ = v10
                _ = v14
            } else {
                switch (p0 >> ((1) & 63)) {
                case -2147483648:
                    p1 = ""
                case 42:
                    v15 := GenS1{F0: 32767}
                    _ = v15
                case 9223372036854775807:
                    fmt.Println(p1)
                default:
                    var v16 uint16 = 243
                    _ = v16
                }
                if true {
                    v17 := GenS1{F0: 94}
                    _ = len(v8)
                    _ = v17
                } else {
                    v18 := GenS1{F0: 1}
                    _ = v18
                }
            }
            _ = fi9
        }
        _ = v8
    }
    v19 := []int{(p0 & (p0 + 4294967295))}
    switch (p0 / (((p0 / ((p0) | 1))) | 1)) {
    case 9223372036854775807:
        fmt.Println(p1, p1)
    case 42:
        var v20 int64 = 9223372036854775807
        _ = v20
    case 128:
        p0 = 42
    }
    _ = p0
    _ = p1
    _ = v19
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = (0.0 / float32((3.14159 - -2.25)))
    var v3 bool = ((3.14159 - (1.5 + (((-2.25 - ((1.5 * 0.0) + 3.14159)) + ((0.0 - float64(3.14159)) + float64(float32(3.14159)))) - 1.5))) > -2.25)
    v21 := GenS1{F0: 1}
    v22 := []int{-222, 216}
    for fi23 := 0; fi23 < 3; fi23++ {
        switch fi23 {
        case -1:
            var v24 uint32 = 42
            _ = v24
        case 0:
            v22 = append(v22, fi23)
        default:
            fmt.Println(v1)
        }
        var v25 uint64 = 18446744073709551615
        fmt.Println(v2, v25, v2)
        _ = fi23
        _ = v25
    }
    var v26 uint16 = 801
    var v27 uint8 = 255
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            switch "\n\t" {
            case "":
                v22 = append(v22, fi29)
            case "foo bar":
                v30 := GenS1{F0: v26}
                _ = v30
            case "\n\t":
                _ = len(v22)
            }
            var v31 string = "hello"
            v26 = 255
            _ = fi29
            _ = v31
        }
        var v32 int = (^(fi28))
        _ = fi28
        _ = v32
    }
    v26 -= v26
    var v33 float64 = (1e10 + 3.14159)
    v22 = append(v22, 32768)
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 6; fi35++ {
            if ((v26 / ((v26) | 1)) >= 315) {
                if false {
                    if (!v3) {
                        var v36 float32 = (v2 / 3.14159)
                        _ = len(v22)
                        v36 += (v36 / (1e10 + (((-2.25 * (float32(((v33 / ((((v33 - (3.14159 - (v33 / (v33 + 3.14159)))) + v33) - v33) - (1.5 / (v33 * v33)))) / (-2.25 - 0.0))) * (v36 * ((v36 / v2) - (1.5 - v2))))) - ((float32(0.0) + (v36 / ((3.14159 - (((0.0 + (0.0 - (float32(float64(v2)) / (3.14159 - 1.5)))) / ((0.1 * (float32(v33) - v2)) * -2.25)) / v36)) * 3.14159))) * v2)) / (1.5 + 1e10))))
                        _ = v36
                    } else {
                        var v37 string = "a"
                        v22 = append(v22, (fi35 * fi34))
                        _ = v37
                    }
                    var v38 int16 = 1000
                    _ = v38
                }
                switch fi34 {
                case 4294967295:
                    fmt.Println(v27, fi34, v2)
                case 2147483647:
                    if (42 >= 2) {
                        var v39 uint8 = (v27 ^ 1)
                        v40 := GenS2{F0: -773}
                        _ = v39
                        _ = v40
                    } else {
                        v3 = (!((4294967295 == 255) || v3))
                        var v41 string = "hello"
                        v42 := GenS3{F0: (v27 + (v27 & 170)), F1: v26}
                        _ = v41
                        _ = v42
                    }
                }
                var v43 int = (-(fi35))
                _ = v43
            } else {
                var v44 float32 = 1.5
                _ = v44
            }
            if (!v3) {
                var v45 uint16 = (^(v26))
                var v46 int32 = 2147483647
                _ = v45
                _ = v46
            } else {
                v2 = v2
                if (0.1 == -2.25) {
                    var v47 float32 = 3.14159
                    v48 := GenS3{F0: v27, F1: (v26 << ((220) & 63))}
                    if v3 {
                        switch ("hello" + "\n\t") {
                        case "hello":
                            var v49 uint = 355
                            _ = v49
                        case "\n\t":
                            var v50 bool = (76 < 100)
                            _ = v50
                        case "":
                            var v51 string = ("\n\t" + "hello")
                            _ = v51
                        default:
                            if len(v22) > 0 {
                                v22[0] = (fi34 * 0)
                            }
                            v2 += (v2 - v47)
                        }
                        fmt.Println(fi35, fi34)
                    }
                    _ = v47
                    _ = v48
                } else {
                    if len(v22) > 0 {
                        v22[0] = (-(fi34))
                    }
                    v27 = v27
                    v33 -= v33
                    v26 += uint16(v27)
                }
                switch "\n\t" {
                case "a":
                    fmt.Println(v26)
                default:
                    v52 := GenS1{F0: (v26 << ((255) & 63))}
                    _ = v52
                }
            }
            v1 -= v1
            fmt.Println(v27, v3)
            _ = fi35
        }
        switch (-(fi34)) {
        case -2147483648:
            for fi53 := 0; fi53 < 17; fi53++ {
                fmt.Println(v26, v1)
                _ = fi53
            }
        case 2147483647:
            if v3 {
                fmt.Println(v3, fi34)
                v54 := GenS1{F0: (v26 << ((806) & 63))}
                _ = v54
            }
        case 32768:
            if len(v22) > 0 {
                v22[0] = fi34
            }
        }
        for fi55 := 0; fi55 < 3; fi55++ {
            v27 = (v27 - v27)
            v56 := GenS4{F0: (fi34 + (fi55 >> ((uint(v27)) & 63))), F1: fi34}
            var v57 int16 = 42
            _ = fi55
            _ = v56
            _ = v57
        }
        _ = fi34
    }
    fmt.Println(v26, v27)
    fmt.Println(v33)
    var v58 int = -9223372036854775808
    v1 += (v1 / ((4294967295) | 1))
    switch "a" {
    case "foo bar":
        _ = len(v22)
    }
    fmt.Println(v58, v27)
    if v3 {
        var v59 uint32 = 322
        _ = v59
    }
    if len(v22) > 0 {
        v22[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v58
}
