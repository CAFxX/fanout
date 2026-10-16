package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 uint32, p2 bool) uint32 {
    v5 := []int{128, -2147483648, 127, 32767}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 32767, F1: 257}
        var v8 int = fi6
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = append(v5, 32767)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 / ((2147483647) | 1))
}
func genFn1(p0 uint32, p1 int8) uint32 {
    var v9 int64 = -1
    p1 -= (p1 + p1)
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v9
    return (p0 | 4294967295)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint = 255
    var v4 bool = (!(-394 <= 42))
    v3 -= v3
    v4 = v4
    if (1 >= 1) {
        fmt.Println(v4, v3)
    }
    v4 = true
    switch 42 {
    case -2147483648:
        var v10 uint = v3
        _ = v10
    case 42:
        var v11 int = 32767
        _ = v11
    }
    _ = len(v2)
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "a" {
        case "foo bar":
            fmt.Println(v4, v4)
        }
        _ = fi12
    }
    if (975 >= -79) {
        v13 := GenS1{F0: 32767, F1: -1}
        var v14 uint8 = 201
        v14 = v14
        _ = v13
        _ = v14
    }
    var v15 int = 621
    v15 = v15
    for fi16 := 0; fi16 < 6; fi16++ {
        v15 = 0
        _ = fi16
    }
    fmt.Println(v15, v15, v4)
    if ((((((0.1 * -2.25) + 1.5) - 1.5) - 3.14159) + 0.1) < 1.5) {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int = fi17
            for fi19 := 0; fi19 < 3; fi19++ {
                if v4 {
                    v4 = true
                }
                v4 = v4
                v20 := GenS1{F0: fi17, F1: 32767}
                if ((3.14159 - 0.0) < 1.5) {
                    if (((v3 * 42) <= v3) && false) {
                        if (((0.1 / 1.5) - (0.1 - 1.5)) <= float32(1.5)) {
                            v21 := v1["hello"]
                            var v22 int32 = -345
                            if (((!false) || v4) || v4) {
                                delete(v1, "hello")
                                v23 := []int{}
                                _ = v23
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                    v24 := GenS1{F0: v15, F1: -32768}
                    _ = v24
                } else {
                    v3 = v3
                    v3 = (v3 & v3)
                }
                _ = fi19
                _ = v20
            }
            v4 = false
            _ = fi17
            _ = v18
        }
        var v25 uint = v3
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if v4 {
            fmt.Println(fi26)
            v15 = v15
            switch fi26 {
            case 4294967295:
                v4 = (-1 >= 597)
            case 9223372036854775807:
                switch fi26 {
                case 92:
                    v15 -= (fi26 + 2147483647)
                case 54:
                    switch "hello" {
                    case "a":
                        v4 = (v4 && false)
                    case "hello":
                        v15 -= (v15 * v15)
                    case "foo bar":
                        if (v4 || (0.0 > 3.14159)) {
                            v2[2147483648] = 2147483647
                            var v27 uint = v3
                            var v28 string = ""
                            _ = v27
                            _ = v28
                        } else {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                switch ("a" + "foo bar") {
                                case "a":
                                    var v30 int64 = -1
                                    _ = v30
                                }
                                switch ("" + ("hello" + "foo bar")) {
                                case "a":
                                    v4 = v4
                                default:
                                    fmt.Println(v15)
                                }
                                _ = fi29
                            }
                            for fi31 := 0; fi31 < 11; fi31++ {
                                v32 := v2[454]
                                switch (fi31 / ((0) | 1)) {
                                case 1:
                                    v2[v32] = 2147483647
                                case 763:
                                    if v4 {
                                        var v33 bool = v4
                                        var v34 uint16 = 8
                                        _ = v33
                                        _ = v34
                                    }
                                    fmt.Println(fi26, fi31)
                                default:
                                    var v35 uint64 = 18446744073709551615
                                    _ = v35
                                }
                                _ = fi31
                                _ = v32
                            }
                        }
                    }
                default:
                    v15 = v15
                }
            case 127:
                if (42 > 0) {
                    var v36 uint = v3
                    _ = v36
                }
            }
        } else {
            var v37 uint32 = 4294967295
            v38 := GenS2{F0: 8, F1: -547, F2: -4}
            _ = v37
            _ = v38
        }
        v3 -= v3
        if v4 {
            var v39 int32 = 2147483647
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int8 = -46
                if v4 {
                    switch ("hello" + "foo bar") {
                    case "":
                        var v42 string = "a"
                        _ = v42
                    case "hello":
                        var v43 bool = (v39 > v39)
                        _ = v43
                    }
                } else {
                    v44 := GenS1{F0: (v15 % ((-564) | 1)), F1: -32768}
                    v15 = v15
                    if v4 {
                        fmt.Println(v3)
                        v39 = v39
                        var v45 uint8 = 42
                        _ = v45
                    } else {
                        if ((v41 == (v41 / ((-17) | 1))) || (638 >= 418)) {
                            v46 := GenS3{F0: 1000}
                            _ = v46
                        } else {
                            if (-1 > 1) {
                                var v47 int8 = v41
                                _ = v47
                            }
                            v41 -= (v41 - v41)
                            switch fi40 {
                            case 127:
                                v15 = 0
                            case 42:
                                switch fi40 {
                                case 388:
                                    v15 -= -1
                                    fmt.Println(fi26, v39)
                                    if (!false) {
                                        var v48 uint64 = 2
                                        _ = v48
                                    } else {
                                        v49 := GenS2{F0: 127, F1: 1, F2: (v41 + v41)}
                                        _ = v49
                                    }
                                    v3 += v3
                                case -9223372036854775808:
                                    switch "a" {
                                    case "hello":
                                        _ = len(v2)
                                    case "\n\t":
                                        v3 -= (v3 - 18446744073709551615)
                                    case "foo bar":
                                        if ((true || (2 > 2147483648)) && v4) {
                                            var v50 int64 = 1
                                            _ = v50
                                        }
                                    default:
                                        fmt.Println(v39)
                                    }
                                    v51 := []int{(fi40 & (v15 | 0)), v15, fi26}
                                    if (2 < 48) {
                                        _ = len(v51)
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    _ = v51
                                case 32767:
                                    var v52 string = "foo bar"
                                    _ = v52
                                }
                            default:
                                var v53 uint8 = 42
                                v4 = v4
                                _ = v53
                            }
                            if (!v4) {
                                v4 = true
                            }
                            if (v4 && (!v4)) {
                                v54 := v1[(("\n\t" + "hello") + "foo bar")]
                                var v55 string = "\n\t"
                                var v56 int8 = -86
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                        }
                    }
                    _ = v44
                }
                fmt.Println(fi40, fi26)
                _ = fi40
                _ = v41
            }
            v57 := GenS4{F0: v39, F1: -97, F2: 0}
            _ = v39
            _ = v57
        }
        _ = fi26
    }
    switch (v15 % ((-887) | 1)) {
    case 1:
        if ((-1 > -1) && true) {
            v4 = ((890 != 32767) && (474 == -2147483648))
            var v58 bool = ((305 >= 18446744073709551615) || v4)
            _ = v58
        }
    }
    var v59 bool = (3.14159 == float32((((0.0 - 0.0) * 1.5) * 0.0)))
    fmt.Println(v59)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v59
}
