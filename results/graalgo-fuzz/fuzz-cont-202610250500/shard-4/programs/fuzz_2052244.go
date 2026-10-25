package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 bool }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 uint, p1 bool) uint {
    var v7 string = "\n\t"
    p1 = (p1 && false)
    var v8 int = 4294967295
    v9 := GenS1{F0: v8}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 uint64) uint8 {
    if true {
        v10 := GenS1{F0: -9223372036854775808}
        var v11 int8 = 100
        _ = v10
        _ = v11
    }
    if (127 > -104) {
        var v12 uint64 = 18446744073709551615
        var v13 float64 = (float64(1.5) * 3.14159)
        fmt.Println(p0, v13)
        _ = v12
        _ = v13
    }
    p0 = p0
    v14 := GenS1{F0: -1}
    _ = p0
    _ = v14
    return 127
}

func main() {
    var v1 int16 = 42
    var v2 int8 = 0
    var v3 uint32 = 225
    var v4 int16 = 0
    var v5 uint64 = 490
    v6 := []string{"a", "foo bar"}
    fmt.Println(v4)
    if (!((3.14159 >= 0.0) || (false || true))) {
        var v15 uint64 = (v5 | (v5 + v5))
        v3 = v3
        var v16 uint32 = 1
        _ = v15
        _ = v16
    } else {
        var v17 int32 = -1
        v1 = (v1 | (v1 << ((genFn0(18446744073709551615, false)) & 63)))
        _ = v17
    }
    v2 -= 1
    v1 = v4
    switch "a" {
    case "foo bar":
        if false {
            fmt.Println(v3, v4)
            v2 = v2
        } else {
            if (42 < -44) {
                var v18 float64 = (1e10 + 0.0)
                v18 -= float64(((-2.25 - (1e10 * (((0.1 + -2.25) + 1.5) + 3.14159))) / 1.0))
                _ = v18
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS1{F0: fi19}
                switch (-(fi19)) {
                case -9223372036854775808:
                    if (!((^(v1)) >= 42)) {
                        v21 := GenS1{F0: (fi19 * 2147483647)}
                        if (1.5 > (((0.0 - (float64(3.14159) / 1.0)) / 1e10) - float64(0.1))) {
                            v5 = v5
                        }
                        _ = v21
                    }
                default:
                    var v22 string = "\n\t"
                    _ = v22
                }
                fmt.Println(v2)
                _ = fi19
                _ = v20
            }
            switch int(v3) {
            case -224:
                switch 143 {
                case 797:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v3 += 2147483648
                        _ = fi23
                    }
                case 42:
                    switch ("" + "a") {
                    case "hello":
                        var v24 uint16 = 1
                        _ = v24
                    }
                }
            case -1:
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 uint16 = 255
                    if (!(1 < genFn1(42))) {
                        v6 = append(v6, "")
                        v27 := GenS2{F0: (v26 ^ (v26 - v26)), F1: v5}
                        if true {
                            v5 = (v5 / (((v5 % ((42) | 1))) | 1))
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if false {
                                    switch (("\n\t" + "\n\t") + "hello") {
                                    case "":
                                        v29 := GenS2{F0: 0, F1: (v5 / (((v5 * v5)) | 1))}
                                        _ = v29
                                    case "a":
                                        fmt.Println(v1, fi28, v1)
                                    case "\n\t":
                                        var v30 string = "a"
                                        v31 := GenS3{F0: (v5 ^ 2), F1: 299}
                                        v32 := GenS1{F0: (fi25 - (fi25 + fi25))}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    default:
                                        if false {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v26 = 679
                                            if false {
                                                var v33 uint32 = (v3 - v3)
                                                _ = v33
                                            } else {
                                                v34 := GenS1{F0: -962}
                                                if true {
                                                    v1 = v4
                                                    fmt.Println(v26, v26)
                                                    v35 := GenS4{F0: v4, F1: v1, F2: (!true)}
                                                    _ = v35
                                                } else {
                                                    var v36 int64 = 9223372036854775807
                                                    switch fi25 {
                                                    case 32768:
                                                        var v37 uint64 = (^(v5))
                                                        _ = v37
                                                    }
                                                    _ = v36
                                                }
                                                var v38 int = int(v4)
                                                _ = v34
                                                _ = v38
                                            }
                                        }
                                        switch (fi25 / ((42) | 1)) {
                                        case 32768:
                                            var v39 int8 = v2
                                            _ = v39
                                        case 648:
                                            v3 -= v3
                                        default:
                                            var v40 float64 = 3.14159
                                            _ = v40
                                        }
                                    }
                                    var v41 uint = 18446744073709551615
                                    _ = v41
                                }
                                v5 = 781
                                var v42 float32 = 0.0
                                _ = fi28
                                _ = v42
                            }
                            v3 -= v3
                        }
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
            case 127:
                switch "foo bar" {
                case "":
                    v5 += (v5 % ((v5) | 1))
                case "a":
                    var v43 int = 127
                    _ = v43
                }
            default:
                fmt.Println(v4)
            }
        }
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v44 := GenS3{F0: (v5 | v5), F1: 2147483647}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v45 := GenS5{F0: (v3 & 2147483647), F1: false}
    var v46 int8 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
