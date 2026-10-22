package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8) int {
    p0 = (^(p0))
    v4 := GenS1{F0: 229, F1: 1}
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 -= (p0 | (p0 - (p0 << ((uint(p0)) & 63))))
        _ = fi5
    }
    var v6 int8 = 42
    _ = p0
    _ = v4
    _ = v6
    return -445
}
func genFn1(p0 uint16, p1 uint16) int {
    switch -595 {
    case -710:
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int32 = 466
            if (32767 > (p1 * p1)) {
                fmt.Println(v8, fi7, fi7)
                p0 = p0
                p1 += p0
            }
            _ = fi7
            _ = v8
        }
    case 2147483647:
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                if (uint64(p0) != 0) {
                    v11 := []int{genFn0(138), fi9}
                    p0 = (^(p1))
                    if false {
                        v12 := GenS2{F0: 40, F1: 173, F2: 65535}
                        fmt.Println(fi9, fi10)
                        _ = len(v11)
                        var v13 int8 = -1
                        var v14 bool = true
                        v14 = false
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                } else {
                    v15 := GenS2{F0: 1, F1: 42, F2: 65535}
                    var v16 uint64 = 0
                    _ = v15
                    _ = v16
                }
                var v17 string = ""
                _ = fi10
                _ = v17
            }
            _ = fi9
        }
    case -2147483648:
        for fi18 := 0; fi18 < 19; fi18++ {
            if (857 < 951) {
                p0 -= p1
                v19 := GenS3{F0: 65535, F1: uint32(p0), F2: fi18}
                _ = v19
            }
            var v20 float32 = 0.0
            if false {
                if ((2147483648 <= 269) && (float32(-2.25) != v20)) {
                    v21 := []int{genFn0(128), (fi18 << ((2) & 63))}
                    p0 = (p0 | p0)
                    _ = v21
                } else {
                    v22 := GenS4{F0: 127, F1: 541}
                    p0 += 1
                    v23 := GenS2{F0: 0, F1: 127, F2: 1}
                    _ = v22
                    _ = v23
                }
                p0 -= p1
            } else {
                for fi24 := 0; fi24 < 20; fi24++ {
                    if (!false) {
                        if false {
                            var v25 int = genFn0(0)
                            _ = v25
                        } else {
                            fmt.Println(p0, fi18, fi24)
                            v26 := GenS1{F0: 42, F1: -989}
                            _ = v26
                        }
                        var v27 uint = 2
                        p0 = p0
                        _ = v27
                    } else {
                        switch ("a" + "") {
                        case "":
                            v20 = (1.5 / float32(3.14159))
                        case "\n\t":
                            if ((int16(fi18) <= -1) && false) {
                                v20 -= v20
                            } else {
                                p0 -= p1
                            }
                        }
                    }
                    var v28 float64 = (3.14159 - 1e10)
                    var v29 int16 = 0
                    _ = fi24
                    _ = v28
                    _ = v29
                }
            }
            _ = fi18
            _ = v20
        }
    }
    if true {
        if (1e10 <= 0.0) {
            v30 := GenS1{F0: 1, F1: 9223372036854775807}
            for fi31 := 0; fi31 < 18; fi31++ {
                v32 := []int{fi31}
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 float32 = 3.14159
                    _ = fi33
                    _ = v34
                }
                var v35 string = "hello"
                _ = fi31
                _ = v32
                _ = v35
            }
            if ((-919 > -2147483648) && ((true && (!true)) && false)) {
                v36 := []int{}
                _ = v36
            }
            _ = v30
        }
        fmt.Println(p0, p1)
    }
    p1 = (p0 % (((p1 ^ 693)) | 1))
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint32 = 2
        _ = fi37
        _ = v38
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint32 = 613
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    var v39 uint = 42
    var v40 float64 = 3.14159
    var v41 int32 = 0
    switch -968 {
    case -19:
        v42 := v2[(v1 * (^(v1)))]
        _ = v42
    case 127:
        switch -9223372036854775808 {
        case 4294967295:
            v43 := GenS5{F0: 152, F1: 9223372036854775807}
            _ = v43
        case -2147483648:
            if false {
                switch 4294967295 {
                case 32767:
                    var v44 int8 = -121
                    _ = v44
                case 493:
                    var v45 int = -2147483648
                    _ = v45
                default:
                    v3 = float32(float64(((v3 - (3.14159 / v3)) / v3)))
                }
            } else {
                var v46 string = "hello"
                _ = v46
            }
        case 32767:
            v47 := v2[uint32(v39)]
            _ = v47
        }
    }
    v48 := v2[v1]
    var v49 int8 = -13
    v48 = v48
    if ((42 < -650) && ((v40 * (v40 * (0.1 + v40))) >= v40)) {
        var v50 int64 = 9223372036854775807
        _ = v50
    } else {
        v51 := []int{1, 9223372036854775807, 32768}
        _ = v51
    }
    var v52 uint = v39
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v48
    _ = v49
    _ = v52
}
