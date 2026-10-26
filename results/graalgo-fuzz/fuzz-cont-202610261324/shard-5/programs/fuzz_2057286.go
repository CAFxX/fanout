package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int32; F1 int8; F2 int32 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int }
func genFn0(p0 uint16) int16 {
    v6 := []int{}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = true
        if v8 {
            var v9 int8 = -128
            v9 -= 38
            var v10 uint8 = 127
            _ = v10
            _ = v9
        }
        p0 = p0
        _ = fi7
        _ = v8
    }
    if true {
        p0 -= (p0 + 32767)
        var v11 float64 = ((((0.1 + (((((1e10 - (((3.14159 + 0.0) / 1.0) - (0.0 + 0.1))) + (0.0 * 3.14159)) / (-2.25 / 1.5)) / (float64(3.14159) / float64((0.1 * (0.1 - 1.5))))) - (1.5 / 0.1))) / (1.5 - (-2.25 / -2.25))) / ((3.14159 * (float64((1.5 - (1e10 + ((1e10 / 0.1) / (0.0 - 3.14159))))) + 0.1)) + ((((0.1 / 1.0) / (0.0 - float64(1.5))) / -2.25) + -2.25))) / float64((((3.14159 / -2.25) / float32(3.14159)) * -2.25)))
        var v12 uint16 = (p0 + (p0 << ((65535) & 63)))
        _ = v11
        _ = v12
    } else {
        p0 += (p0 + p0)
    }
    v13 := GenS1{F0: 112, F1: 65535}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v13
    _ = v6
    return -32768
}

func main() {
    var v1 int8 = -128
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    var v5 string = "foo bar"
    v16 := []int{32768}
    var v17 uint16 = 255
    if true {
        if (!false) {
            v4 += (v4 + (^(v4)))
            v17 = v17
        } else {
            fmt.Println(v2, v2)
            switch 200 {
            case 1:
                if true {
                    if (v1 >= (v1 >> ((v2) & 63))) {
                        var v18 uint8 = 189
                        _ = v18
                    }
                    fmt.Println(v5, v4)
                    var v19 int32 = -1
                    _ = v19
                } else {
                    var v20 string = "hello"
                    if len(v16) > 0 {
                        v16[0] = 15
                    }
                    switch "" {
                    case "":
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v2 += (v2 >> (((v2 % ((v2) | 1))) & 63))
                            v2 = v2
                            switch v20 {
                            case "a":
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            case "hello":
                                fmt.Println(fi21, v4)
                            case "foo bar":
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    var v23 uint64 = 233
                                    var v24 int32 = 0
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            default:
                                if (1 <= 42) {
                                    v25 := GenS2{F0: 2147483647, F1: 307, F2: 128}
                                    _ = v25
                                } else {
                                    switch (fi21 | 1) {
                                    case -207:
                                        v4 += -1
                                    case 9223372036854775807:
                                        var v26 uint16 = 42
                                        _ = v26
                                    case -9223372036854775808:
                                        if (4294967295 >= 582) {
                                            if true {
                                                for fi27 := 0; fi27 < 5; fi27++ {
                                                    v3[("\n\t" + ("a" + ("foo bar" + v5)))] = true
                                                    v28 := GenS3{F0: v2}
                                                    _ = fi27
                                                    _ = v28
                                                }
                                                v29 := GenS2{F0: -899, F1: 42, F2: 127}
                                                var v30 int16 = -437
                                                _ = v29
                                                _ = v30
                                            } else {
                                                var v31 uint32 = 2
                                                v32 := GenS4{F0: v5, F1: v20}
                                                _ = v31
                                                _ = v32
                                            }
                                            var v33 float64 = 0.1
                                            var v34 int16 = 1
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                }
                            }
                            _ = fi21
                        }
                    case "hello":
                        var v35 int16 = 1
                        _ = v35
                    case "\n\t":
                        switch -271 {
                        case 235:
                            if true {
                                if false {
                                    for fi36 := 0; fi36 < 20; fi36++ {
                                        var v37 int64 = (v4 + (v4 - 1))
                                        var v38 uint32 = 2147483648
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                fmt.Println(v5)
                                v39 := GenS3{F0: 550}
                                _ = v39
                            }
                        case -9223372036854775808:
                            v4 += 1
                        default:
                            var v40 int8 = (v1 & v1)
                            _ = v40
                        }
                    default:
                        if false {
                            var v41 uint = v2
                            var v42 int8 = int8(v4)
                            _ = v41
                            _ = v42
                        } else {
                            var v43 float64 = 3.14159
                            if (127 <= 9223372036854775807) {
                                v5 = v20
                                fmt.Println(v1)
                                for fi44 := 0; fi44 < 2; fi44++ {
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        v46 := v3[v20]
                                        switch -1 {
                                        case -990:
                                            v1 -= 100
                                        default:
                                            v43 = float64(0.0)
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                    _ = fi44
                                }
                            } else {
                                var v47 float32 = -2.25
                                v48 := v3[v5]
                                var v49 int16 = 1000
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            _ = v43
                        }
                    }
                    _ = v20
                }
            case 816:
                v50 := GenS5{F0: 42, F1: (^(v1)), F2: 69}
                _ = v50
            case 128:
                var v51 uint64 = 2
                _ = v51
            }
            v17 = (v17 & (v17 >> (((v2 | v2)) & 63)))
        }
    } else {
        fmt.Println(v2, v4, v17)
    }
    for fi52 := 0; fi52 < 4; fi52++ {
        v53 := GenS6{F0: v5, F1: v17}
        fmt.Println(fi52, v5)
        _ = fi52
        _ = v53
    }
    var v54 uint64 = 42
    v55 := GenS7{F0: int(v1)}
    v56 := v3["foo bar"]
    if false {
        var v57 float64 = 1.5
        _ = v57
    } else {
        v54 += 18446744073709551615
        var v58 int8 = -72
        _ = v58
    }
    fmt.Println(v17, v1, v54)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
}
