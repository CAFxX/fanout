package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS6 struct { F0 int }
func genFn0(p0 uint32, p1 uint16, p2 uint16) int8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: p0, F1: 614, F2: p2}
        p2 = p2
        v7 := []int{int(p0)}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 7; fi8++ {
        p2 -= (p1 - p1)
        _ = fi8
    }
    v9 := []int{636, -761, 32768, -448}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}
func genFn1(p0 uint64, p1 int8, p2 bool) string {
    fmt.Println(p1, p2)
    var v10 int32 = 2147483647
    p1 += 100
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return ("\n\t" + ("hello" + "a"))
}

func main() {
    var v1 uint8 = 30
    var v2 int = 32768
    v3 := make(map[string]bool)
    var v4 uint = 65535
    _ = len(v3)
    delete(v3, "hello")
    var v11 uint64 = 416
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v3[""] = false
            if (false || false) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    if (false || (true && (!true))) {
                        var v14 uint16 = 995
                        if ((!true) && false) {
                            switch "hello" {
                            case "\n\t":
                                v15 := []int{(v2 % ((v2) | 1)), (fi13 + fi13), (^(fi13)), fi12}
                                _ = v15
                            case "foo bar":
                                v16 := []int{}
                                _ = v16
                            case "":
                                v14 = 0
                            default:
                                fmt.Println(fi13)
                            }
                        }
                        _ = v14
                    }
                default:
                    fmt.Println(v1, v11, v1)
                }
                var v17 uint64 = 2
                _ = v17
            } else {
                fmt.Println(fi13, v1)
                switch fi12 {
                case 73:
                    v18 := GenS1{F0: 336, F1: (v4 % (((v4 | (v4 | 2))) | 1)), F2: 32767}
                    _ = v18
                default:
                    switch "hello" {
                    case "\n\t":
                        var v19 uint = v4
                        _ = v19
                    default:
                        var v20 int16 = 1000
                        _ = v20
                    }
                }
            }
            if false {
                v21 := []int{(fi13 << (((v4 - (v4 >> ((v4) & 63)))) & 63)), (v2 * (-(v2))), v2}
                _ = v21
            } else {
                v22 := GenS2{F0: 2, F1: (v11 - 0)}
                _ = v22
            }
            _ = fi13
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS2{F0: 0, F1: (v11 * v11)}
            if (!false) {
                v2 = (fi12 & (v2 ^ (fi23 >> ((uint(v1)) & 63))))
                v25 := GenS3{F0: 0, F1: (v1 * (v1 << ((65535) & 63))), F2: 556}
                v4 = v4
                _ = v25
            }
            v26 := []int{(v2 % (((fi23 | 0)) | 1)), (^(fi23)), (fi12 / ((v2) | 1))}
            fmt.Println(fi12, v2)
            var v27 int16 = 1
            var v28 int64 = 42
            v29 := GenS4{F0: 599, F1: 526, F2: (v1 >> ((940) & 63))}
            _ = fi23
            _ = v24
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 int64 = 1
        _ = fi12
        _ = v30
    }
    var v31 int16 = 694
    if true {
        v3["foo bar"] = (v31 != (v31 >> (((v4 - 0)) & 63)))
        var v32 float32 = 1.5
        fmt.Println(v2, v1)
        _ = v32
    }
    if true {
        var v33 float64 = 1.5
        _ = v33
    } else {
        var v34 uint64 = 2
        _ = v34
    }
    var v35 int64 = -823
    v35 = (v35 - (v35 - int64(v2)))
    v36 := GenS5{F0: 437, F1: v31, F2: 255}
    for fi37 := 0; fi37 < 16; fi37++ {
        v1 = (v1 + 246)
        var v38 string = (("hello" + "a") + "")
        _ = fi37
        _ = v38
    }
    var v39 uint64 = 67
    v40 := GenS2{F0: 338, F1: v39}
    v1 += v1
    if false {
        v2 -= 128
        if (!false) {
            v41 := GenS5{F0: 32767, F1: (v31 << ((v4) & 63)), F2: 210}
            _ = v41
        } else {
            var v42 uint32 = 1
            _ = v42
        }
        switch 382 {
        case 9223372036854775807:
            var v43 uint8 = 112
            _ = v43
        case -9223372036854775808:
            if (!false) {
                var v44 uint32 = 65535
                fmt.Println(v39, v1, v4)
                _ = v44
            }
        case 840:
            v45 := v3[("hello" + ("a" + ("\n\t" + "foo bar")))]
            _ = v45
        }
    } else {
        v35 += (v35 / (((v35 - 12)) | 1))
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        if true {
            var v46 uint32 = 42
            _ = v46
        } else {
            switch "foo bar" {
            case "hello":
                v47 := GenS6{F0: (v2 ^ v2)}
                _ = v47
            default:
                v2 = v2
            }
            v39 = (v39 ^ v39)
            switch v2 {
            case 42:
                fmt.Println(v11)
            case 128:
                switch 1 {
                case -975:
                    v1 = (^(v1))
                case -599:
                    for fi48 := 0; fi48 < 9; fi48++ {
                        v39 = (^(v11))
                        for fi49 := 0; fi49 < 16; fi49++ {
                            var v50 float64 = ((((1.5 + 0.1) - -2.25) / 3.14159) + (1.5 / 1e10))
                            v51 := v3[("a" + (("\n\t" + (("a" + ("hello" + "a")) + "hello")) + "\n\t"))]
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        _ = fi48
                    }
                }
            case 515:
                var v52 uint32 = 800
                _ = v52
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
}
