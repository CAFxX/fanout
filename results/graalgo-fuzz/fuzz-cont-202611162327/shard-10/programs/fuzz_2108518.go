package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint8, p1 uint16) int16 {
    v6 := []int{}
    v7 := GenS1{F0: 986}
    var v8 float32 = (-2.25 / 0.1)
    var v9 int = int(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -710
}
func genFn1(p0 int) uint16 {
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint8 = 1
        fmt.Println(p0, v11)
        _ = fi10
        _ = v11
    }
    p0 -= p0
    v12 := []int{p0}
    _ = p0
    _ = v12
    return 255
}

func main() {
    var v1 int32 = 42
    v2 := []string{}
    v3 := []uint32{}
    var v4 int8 = -128
    var v5 uint16 = 494
    var v13 float64 = 0.0
    v13 += v13
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint8 = 60
        var v16 float32 = 1.5
        if true {
            if (v5 != 680) {
                var v17 uint = 18446744073709551615
                _ = v17
            } else {
                for fi18 := 0; fi18 < 5; fi18++ {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    v1 = (-(v1))
                    _ = fi18
                }
                fmt.Println(v16, v16, v13)
                for fi19 := 0; fi19 < 19; fi19++ {
                    v1 = -210
                    _ = fi19
                }
            }
            switch fi14 {
            case 12:
                var v20 int = 247
                _ = v20
            case 4294967295:
                v2 = append(v2, "")
            case 2147483647:
                var v21 uint8 = v15
                _ = v21
            default:
                fmt.Println(v16, v15, v15)
            }
            var v22 bool = true
            _ = v22
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 -= (v4 + (v4 ^ v4))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi23
    }
    if false {
        switch "" {
        case "hello":
            for fi24 := 0; fi24 < 17; fi24++ {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v26 := GenS1{F0: 4294967295}
                    if (1 < 4294967295) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v27 float64 = 1e10
                        v2 = append(v2, ("" + "\n\t"))
                        switch fi25 {
                        case 32768:
                            v13 += v13
                        case -344:
                            v5 -= (v5 % ((65535) | 1))
                        }
                        v28 := GenS2{F0: v4, F1: 4294967295}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 int = -2147483648
                        v29 = 42
                        if true {
                            var v30 int64 = -10
                            switch fi25 {
                            case 127:
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            case -534:
                                var v31 uint = 42
                                _ = v31
                            case -9223372036854775808:
                                v29 = int(v5)
                            }
                            fmt.Println(v1, v13, v1)
                            v32 := GenS3{F0: (!true)}
                            _ = v30
                            _ = v32
                        }
                        _ = v29
                    }
                    v1 = (v1 / (((v1 ^ 2147483647)) | 1))
                    _ = fi25
                    _ = v26
                }
                v13 += float64((0.1 - (-2.25 + 0.1)))
                _ = fi24
            }
        case "a":
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        case "\n\t":
            v33 := GenS1{F0: 884}
            _ = v33
        }
    } else {
        if (0 <= 0) {
            fmt.Println(v5)
            v13 = ((v13 - ((0.0 + (1.5 + v13)) + (v13 + 1.5))) - (((v13 / v13) * v13) / float64(1e10)))
            if ((-32768 > -616) || (!false)) {
                _ = len(v3)
                switch -962 {
                case 599:
                    var v34 int = -2147483648
                    _ = v34
                }
            } else {
                v4 += v4
            }
        } else {
            v1 = 1
            var v35 uint8 = 1
            _ = v35
        }
        fmt.Println(v1, v1, v4)
        if false {
            var v36 bool = (false || (357 < 2))
            switch 32767 {
            case 2147483647:
                var v37 int = 1
                _ = v37
            default:
                v2 = append(v2, "")
            }
            _ = v36
        } else {
            var v38 uint16 = v5
            _ = v38
        }
    }
    fmt.Println(v4)
    var v39 int16 = genFn0(255, v5)
    v2 = append(v2, "a")
    v40 := GenS4{F0: false, F1: (v1 - v1), F2: "a"}
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 int32 = v1
        var v43 int8 = v4
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 uint32 = 0
            _ = fi44
            _ = v45
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
