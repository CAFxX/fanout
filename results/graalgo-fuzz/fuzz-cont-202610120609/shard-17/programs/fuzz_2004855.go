package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int64; F2 int32 }
type GenS5 struct { F0 int; F1 int32; F2 uint }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint64, p2 uint) string {
    fmt.Println(p0)
    var v6 uint8 = 116
    for fi7 := 0; fi7 < 9; fi7++ {
        p1 += p1
        v8 := GenS1{F0: (v6 >> (((p2 ^ (p2 << ((p2) & 63)))) & 63)), F1: (^(v6))}
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int8 = int8(v6)
            v11 := []int{127}
            v12 := GenS1{F0: 0, F1: v6}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = 100
    var v3 int8 = 64
    v4 := make(map[int]int)
    var v5 uint = 42
    v13 := v4[-2147483648]
    fmt.Println(v3, v13, v1)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = -1
        var v16 int64 = v15
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = (^(v2))
    v13 += v1
    v1 += 734
    var v17 int = v1
    if (!(false && false)) {
        if true {
            fmt.Println(v13)
        } else {
            v4[(^(v1))] = -54
            var v18 int = v13
            _ = v18
        }
        if false {
            fmt.Println(v5, v3)
            _ = len(v4)
        } else {
            switch (v17 | (^(v13))) {
            case 1:
                fmt.Println(v2)
            case 32767:
                v19 := GenS1{F0: 42, F1: 128}
                _ = v19
            case 617:
                var v20 float64 = 3.14159
                _ = v20
            default:
                if false {
                    v3 = v2
                    var v21 int = 2147483647
                    v22 := GenS2{F0: -281, F1: 708, F2: v5}
                    _ = v21
                    _ = v22
                } else {
                    switch 127 {
                    case 32768:
                        v4[v17] = v13
                    case 9223372036854775807:
                        v23 := []int{(v17 - v13), (v17 * v13), (-(v17)), -312}
                        _ = v23
                    case 127:
                        v24 := GenS3{F0: 0}
                        _ = v24
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v3)
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            v25 := GenS2{F0: 110, F1: 614, F2: v5}
                            _ = v25
                        }
                    default:
                        if false {
                            fmt.Println(v1, v13, v1)
                        } else {
                            if (-638 >= 0) {
                                for fi26 := 0; fi26 < 7; fi26++ {
                                    _ = len(v4)
                                    if (2 != 42) {
                                        v17 -= (v13 / ((115) | 1))
                                    } else {
                                        var v27 string = ("foo bar" + "a")
                                        var v28 float64 = 3.14159
                                        var v29 uint16 = 255
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    v1 = v13
                                    _ = fi26
                                }
                                for fi30 := 0; fi30 < 20; fi30++ {
                                    v31 := GenS4{F0: 2147483647, F1: int64(v3), F2: -1}
                                    v5 += 42
                                    _ = fi30
                                    _ = v31
                                }
                            }
                            v13 += v13
                        }
                    }
                }
            }
        }
    } else {
        var v32 float64 = (-2.25 + ((1e10 * 0.0) / -2.25))
        v4[v13] = v17
        var v33 int8 = v2
        _ = v32
        _ = v33
    }
    var v34 int64 = -753
    switch (^(v17)) {
    case 127:
        if (0.1 <= (0.1 - 3.14159)) {
            v2 = v3
        }
    case 4294967295:
        var v35 bool = (0 < 2147483648)
        _ = v35
    }
    fmt.Println(v34, v3, v2)
    if false {
        var v36 int16 = 1
        _ = v36
    } else {
        var v37 int32 = -1
        var v38 bool = false
        v39 := GenS1{F0: 42, F1: 255}
        _ = v37
        _ = v38
        _ = v39
    }
    var v40 uint = v5
    if false {
        if (1e10 < 1e10) {
            v4[v13] = 127
            var v41 bool = true
            v42 := GenS1{F0: 127, F1: 166}
            _ = v41
            _ = v42
        } else {
            if true {
                if false {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        v1 += (v13 ^ (fi43 + (v17 - (v17 % ((-924) | 1)))))
                        if ((true && true) && false) {
                            if (!true) {
                                v44 := []int{(fi43 * -585), v1, fi43, (v1 / ((v17) | 1))}
                                v45 := GenS5{F0: 1, F1: -809, F2: v40}
                                _ = v44
                                _ = v45
                            } else {
                                switch ("" + ("\n\t" + ("hello" + ("a" + ("a" + "\n\t"))))) {
                                case "hello":
                                    for fi46 := 0; fi46 < 13; fi46++ {
                                        v34 = v34
                                        v13 = fi43
                                        var v47 uint = (v40 & (v40 & v40))
                                        _ = fi46
                                        _ = v47
                                    }
                                }
                            }
                        }
                        for fi48 := 0; fi48 < 3; fi48++ {
                            v49 := v4[v17]
                            _ = fi48
                            _ = v49
                        }
                        _ = fi43
                    }
                    for fi50 := 0; fi50 < 1; fi50++ {
                        var v51 float32 = float32(-2.25)
                        v52 := GenS6{F0: -720}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    v2 = (v2 & v2)
                } else {
                    var v53 int32 = -604
                    var v54 uint16 = 255
                    v55 := GenS3{F0: (v53 >> ((v40) & 63))}
                    _ = v53
                    _ = v54
                    _ = v55
                }
                var v56 int64 = -9223372036854775808
                if true {
                    v3 = int8(v1)
                }
                _ = v56
            }
            var v57 uint32 = 576
            _ = v57
        }
        v34 -= v34
        for fi58 := 0; fi58 < 11; fi58++ {
            var v59 float64 = (1e10 / 1.0)
            for fi60 := 0; fi60 < 16; fi60++ {
                v40 = 18446744073709551615
                if (!true) {
                    v61 := []int{(fi60 ^ v1), v17, (fi60 - 0), (fi58 + fi58)}
                    _ = v61
                }
                var v62 int8 = 1
                _ = fi60
                _ = v62
            }
            v2 = (v2 & (v3 ^ 1))
            _ = fi58
            _ = v59
        }
    }
    var v63 int8 = v3
    v3 = (-(v2))
    for fi64 := 0; fi64 < 7; fi64++ {
        fmt.Println(fi64, v40, fi64)
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v63
}
