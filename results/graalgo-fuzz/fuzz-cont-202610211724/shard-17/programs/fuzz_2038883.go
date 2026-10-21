package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint, p2 int8) uint8 {
    v7 := []int{255}
    var v8 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 163
}
func genFn1(p0 uint16, p1 int8, p2 int) int8 {
    if true {
        switch 0 {
        case -9223372036854775808:
            var v9 bool = false
            _ = v9
        case -929:
            if (3.14159 != -2.25) {
                switch "hello" {
                case "\n\t":
                    var v10 uint = 255
                    _ = v10
                }
            } else {
                if (true || (float64(((((float32(((1e10 * 1e10) + 0.1)) - 0.0) / -2.25) / float32((float64(0.0) - ((1e10 + float64(float32(1.5))) + float64((0.1 / (3.14159 - (-2.25 + (1e10 * ((1e10 + 0.1) + 0.1)))))))))) + 0.0)) < (float64(((0.1 / (-2.25 - float32(1.5))) - 3.14159)) + 1e10))) {
                    var v11 uint64 = 2
                    _ = v11
                } else {
                    fmt.Println(p1)
                    var v12 int32 = -627
                    _ = v12
                }
            }
        case 127:
            v13 := GenS1{F0: 32767, F1: (p0 % ((23) | 1)), F2: 42}
            _ = v13
        default:
            p2 += p2
        }
        v14 := GenS1{F0: (p0 + (p0 & (p0 - (p0 & 2)))), F1: p0, F2: int16(p2)}
        _ = v14
    } else {
        p2 -= p2
    }
    var v15 int32 = -1
    v16 := GenS1{F0: p0, F1: (^(p0)), F2: -1}
    if ((-943 != 42) || true) {
        v17 := []int{42, (p2 - (p2 + p2)), (^(p2))}
        p2 -= p2
        v17 = append(v17, (p2 + (^(p2))))
        _ = v17
    }
    var v18 float32 = (0.1 - -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v18
    return (p1 >> ((255) & 63))
}

func main() {
    v1 := []int{}
    var v2 int32 = 42
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v5 int32 = (v2 % ((v2) | 1))
    v6 := make(map[uint32]string)
    var v19 uint64 = 18446744073709551615
    delete(v4, 2147483648)
    var v20 int8 = 1
    if (42 != 770) {
        var v21 float64 = float64(0.0)
        v2 = v5
        switch ("hello" + "a") {
        case "foo bar":
            v22 := GenS2{F0: ((v21 * (0.1 / 1.0)) == 1e10)}
            _ = v22
        case "":
            for fi23 := 0; fi23 < 17; fi23++ {
                fmt.Println(v2, fi23)
                if (2 <= 42) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 uint16 = uint16(v19)
                        switch "hello" {
                        case "":
                            v26 := v6[2147483648]
                            _ = v26
                        case "hello":
                            var v27 float64 = float64(-2.25)
                            _ = v27
                        default:
                            if false {
                                v28 := GenS3{F0: v5, F1: 2, F2: 1000}
                                _ = v28
                            }
                        }
                        if false {
                            if (!true) {
                                var v29 uint64 = (v19 | (v19 * 0))
                                switch "a" {
                                case "\n\t":
                                    if (0.0 == (3.14159 * (((0.0 - 1.5) + -2.25) + float32(-2.25)))) {
                                        v29 = 217
                                        if (uint8(v29) <= 255) {
                                            if (false || false) {
                                                v21 = 1e10
                                                v25 += (v25 << ((18446744073709551615) & 63))
                                            }
                                            v30 := GenS2{F0: (2147483648 <= 2147483648)}
                                            _ = v30
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    } else {
                                        v20 += genFn1(v25, genFn1((v25 & (v25 * (v25 << ((780) & 63)))), v20, (fi23 * fi24)), fi24)
                                    }
                                case "hello":
                                    var v31 uint8 = uint8(v25)
                                    v21 = 1e10
                                    v32 := GenS1{F0: v25, F1: (v25 * (v25 | uint16(v31))), F2: -1}
                                    _ = v31
                                    _ = v32
                                case "foo bar":
                                    v5 -= -2147483648
                                default:
                                    var v33 int64 = 0
                                    _ = v33
                                }
                                v2 += v2
                                _ = v29
                            }
                            if (v2 == v2) {
                                fmt.Println(v19, fi24, fi24)
                                v5 = 1
                                delete(v3, "")
                                v2 = (^(v5))
                                if len(v1) > 0 {
                                    v1[0] = -9223372036854775808
                                }
                            }
                            switch (fi23 & (fi23 / ((fi24) | 1))) {
                            case -457:
                                fmt.Println(v25, v21)
                            case 128:
                                v34 := v6[2147483647]
                                _ = v34
                            default:
                                var v35 uint64 = (v19 - (^(v19)))
                                _ = v35
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                    var v36 uint8 = 46
                    v19 += v19
                    _ = v36
                }
                v1 = append(v1, fi23)
                _ = fi23
            }
        case "a":
            switch "" {
            case "a":
                for fi37 := 0; fi37 < 8; fi37++ {
                    if (42 != 1) {
                        switch fi37 {
                        case -9223372036854775808:
                            switch (fi37 + 128) {
                            case 42:
                                if (int16(v5) <= 42) {
                                    switch fi37 {
                                    case -1:
                                        v38 := v3["hello"]
                                        _ = v38
                                    case 497:
                                        v2 = (v2 * v5)
                                    case -723:
                                        v20 = (v20 * v20)
                                    }
                                    fmt.Println(fi37, v21, v19)
                                }
                            case 901:
                                var v39 float32 = -2.25
                                _ = v39
                            case -119:
                                fmt.Println(fi37, v21)
                            }
                        case -280:
                            var v40 uint8 = genFn0(-504, 309, 100)
                            _ = v40
                        case 1:
                            for fi41 := 0; fi41 < 1; fi41++ {
                                v42 := v6[2]
                                _ = fi41
                                _ = v42
                            }
                        default:
                            var v43 string = "\n\t"
                            _ = v43
                        }
                        v5 += -830
                    } else {
                        fmt.Println(fi37, fi37, fi37)
                    }
                    _ = fi37
                }
            }
        default:
            fmt.Println(v5, v21, v2)
        }
        _ = v21
    }
    var v44 uint64 = (v19 ^ v19)
    v45 := GenS1{F0: 32767, F1: 712, F2: 0}
    if (4294967295 <= (v19 | 18446744073709551615)) {
        var v46 uint32 = 2147483647
        v47 := GenS4{F0: ""}
        _ = v46
        _ = v47
    } else {
        fmt.Println(v44)
    }
    if ((v20 * (v20 + v20)) == genFn1(65535, v20, -9223372036854775808)) {
        for fi48 := 0; fi48 < 6; fi48++ {
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            _ = fi48
        }
    } else {
        v3[("hello" + "hello")] = "\n\t"
    }
    fmt.Println(v20, v44, v5)
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v20 = -128
    var v49 int16 = 1
    if true {
        var v50 float32 = float32((0.1 * 3.14159))
        v51 := GenS1{F0: 0, F1: 32767, F2: int16(v20)}
        v2 -= v5
        _ = v50
        _ = v51
    } else {
        v52 := GenS5{F0: 146}
        var v53 int = -604
        v54 := GenS5{F0: 42}
        _ = v52
        _ = v53
        _ = v54
    }
    v20 -= v20
    var v55 int32 = (v2 >> ((255) & 63))
    if false {
        var v56 int8 = (v20 >> ((1) & 63))
        for fi57 := 0; fi57 < 7; fi57++ {
            delete(v6, 65535)
            v58 := GenS4{F0: (("\n\t" + (("a" + "a") + ("hello" + "hello"))) + ("\n\t" + "\n\t"))}
            _ = fi57
            _ = v58
        }
        var v59 int64 = int64(v44)
        _ = v56
        _ = v59
    } else {
        v60 := GenS3{F0: 974, F1: 145, F2: (v49 << ((18446744073709551615) & 63))}
        v61 := v4[2147483648]
        _ = v60
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v49
    _ = v5
    _ = v55
    _ = v6
}
