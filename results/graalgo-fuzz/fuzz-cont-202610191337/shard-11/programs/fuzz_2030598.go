package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint8, p2 string) bool {
    v7 := []int{32767, 0}
    var v8 uint64 = 302
    var v9 int = 42
    var v10 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1() int32 {
    if ((((((1.5 - (-2.25 / (1.5 - -2.25))) - (0.1 - (0.1 * (1.5 - float64(3.14159))))) / (0.1 / float64(1e10))) / 1.0) / ((float64(((1e10 + ((1e10 + 0.0) + 0.0)) * 3.14159)) / ((0.0 / 1.5) + 3.14159)) + -2.25)) > 3.14159) {
        var v11 uint16 = 2
        for fi12 := 0; fi12 < 4; fi12++ {
            v11 += v11
            var v13 float32 = 0.0
            fmt.Println(v11)
            _ = fi12
            _ = v13
        }
        var v14 string = "a"
        _ = v11
        _ = v14
    }
    var v15 int16 = -30
    if ((!((v15 ^ -1) != (v15 / ((v15) | 1)))) || (!false)) {
        var v16 int = 1
        v15 = (v15 ^ v15)
        v16 = (v16 & (^(v16)))
        _ = v16
    } else {
        fmt.Println(v15)
    }
    _ = v15
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 int16 = -999
    var v5 float32 = (-2.25 / ((1.5 / 1.5) / 1.5))
    var v6 int32 = -1
    for fi17 := 0; fi17 < 7; fi17++ {
        if false {
            var v18 float32 = ((v5 + 0.0) + ((v5 * float32((3.14159 / ((-2.25 / (0.1 / float64(1.5))) - ((3.14159 / 1.0) - (0.0 + -2.25)))))) - 3.14159))
            if true {
                if false {
                    var v19 int8 = 42
                    v19 += v19
                    switch (fi17 * (fi17 >> ((255) & 63))) {
                    case 9223372036854775807:
                        v2 = v2
                    case 42:
                        switch (fi17 >> ((65535) & 63)) {
                        case -2147483648:
                            v4 -= (v4 / ((v4) | 1))
                        case 128:
                            v20 := v1[878]
                            _ = v20
                        case 131:
                            var v21 string = ("foo bar" + "\n\t")
                            _ = v21
                        }
                    case -1:
                        for fi22 := 0; fi22 < 4; fi22++ {
                            switch fi17 {
                            case 135:
                                switch "" {
                                case "foo bar":
                                    v23 := []int{(fi22 & 9223372036854775807), (fi17 % ((fi17) | 1)), fi22, fi17}
                                    _ = v23
                                case "a":
                                    v24 := []int{}
                                    _ = v24
                                default:
                                    switch (fi17 - (fi17 | fi22)) {
                                    case 127:
                                        fmt.Println(fi22, v2)
                                    case 830:
                                        if (2 < 231) {
                                            fmt.Println(v6, v4)
                                            v2 = -9223372036854775808
                                            v25 := GenS1{F0: false, F1: v4, F2: 0}
                                            _ = v25
                                        }
                                    case -1:
                                        v19 -= int8(v6)
                                    default:
                                        v26 := []int{-149, fi17, -9223372036854775808, fi17}
                                        _ = v26
                                    }
                                }
                            case -461:
                                switch "hello" {
                                case "a":
                                    v18 = (v5 / v5)
                                    var v27 uint16 = 777
                                    var v28 int = fi17
                                    _ = v27
                                    _ = v28
                                default:
                                    if false {
                                        v29 := GenS2{F0: 42, F1: 65535}
                                        var v30 int64 = (v2 / ((v2) | 1))
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 int16 = (v4 + (v4 & v4))
                                        var v32 uint32 = 42
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            default:
                                if true {
                                    fmt.Println(v5, v18)
                                    v33 := []int{}
                                    if (!(!false)) {
                                        v5 = -2.25
                                    }
                                    _ = v33
                                } else {
                                    var v34 uint32 = 490
                                    _ = v34
                                }
                            }
                            _ = fi22
                        }
                    }
                    _ = v19
                }
            }
            _ = v18
        }
        var v35 int64 = 1
        switch "hello" {
        case "":
            var v36 int = 32768
            _ = v36
        case "\n\t":
            var v37 int = fi17
            _ = v37
        default:
            var v38 int = (fi17 ^ -807)
            _ = v38
        }
        _ = fi17
        _ = v35
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        v40 := GenS3{F0: 113}
        v5 = v5
        _ = fi39
        _ = v40
    }
    v5 = 0.1
    v1[2] = (!((false || (835 != 2147483647)) && (((0.1 / 1.5) - 0.0) != -2.25)))
    delete(v1, 1)
    v4 += v4
    v4 -= v4
    v5 += v5
    var v41 uint64 = 2
    fmt.Println(v4, v5, v5)
    if false {
        if (42 > 255) {
            var v42 float32 = v5
            if true {
                fmt.Println(v5, v41, v2)
                v1[542] = true
                v5 += v42
            }
            var v43 int = int(v4)
            _ = v42
            _ = v43
        } else {
            v44 := GenS2{F0: (v41 >> ((0) & 63)), F1: 887}
            for fi45 := 0; fi45 < 5; fi45++ {
                v46 := GenS1{F0: (v6 < v6), F1: (v4 % ((v4) | 1)), F2: v41}
                if (v4 >= v4) {
                    var v47 int8 = 53
                    _ = v47
                }
                _ = fi45
                _ = v46
            }
            _ = v44
        }
        fmt.Println(v41, v6)
    } else {
        var v48 int16 = v4
        _ = v48
    }
    v49 := GenS3{F0: 93}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v49
    _ = v5
    _ = v6
}
