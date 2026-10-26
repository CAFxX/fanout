package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 int64, p1 uint64, p2 int) uint64 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{p2, (fi4 | (p2 / (((p2 & (p2 << ((255) & 63)))) | 1))), int(p1), p2}
        _ = fi4
        _ = v5
    }
    v6 := []int{(p2 % ((p2) | 1))}
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 uint16 = 42
    var v2 int = 32767
    var v3 int = (v2 << ((643) & 63))
    switch "" {
    case "a":
        v8 := []int{}
        _ = v8
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := []int{(v3 * v2), (^(v2))}
                v12 := GenS1{F0: 246, F1: 112}
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := []int{(v3 + v3)}
            var v14 uint = 42
            _ = fi9
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 float32 = (((float32(1e10) + -2.25) + float32(float64(0.1))) / 3.14159)
        var v17 uint8 = 1
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := []int{(fi18 >> ((0) & 63))}
        _ = fi18
        _ = v19
    }
    if (0 == 100) {
        var v20 bool = false
        _ = v20
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(v3, fi21)
            _ = fi21
        }
        var v22 uint32 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint = 445
            switch "" {
            case "hello":
                v3 = 263
            default:
                fmt.Println(fi24, fi24, v2)
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v3 += v3
    v26 := GenS2{F0: -637, F1: 816}
    v27 := []int{v3, -9223372036854775808, (v3 - (v2 ^ (v3 % ((778) | 1))))}
    switch 127 {
    case -496:
        v1 += (v1 - v1)
    case 32768:
        v28 := GenS3{F0: 42, F1: 1, F2: 892}
        _ = v28
    case 893:
        switch ("" + ("" + "a")) {
        case "foo bar":
            var v29 int64 = -1
            _ = v29
        case "hello":
            var v30 uint8 = 0
            _ = v30
        case "":
            v1 += (v1 ^ v1)
        default:
            v27 = append(v27, v2)
        }
    default:
        if (!false) {
            v1 = (v1 | v1)
            var v31 int32 = -2147483648
            v3 = v2
            _ = v31
        } else {
            if false {
                if true {
                    var v32 uint8 = 200
                    if (0 >= 2) {
                        var v33 uint16 = v1
                        var v34 int32 = 641
                        if (!((v33 & v1) >= v33)) {
                            v3 += (v2 | int(v34))
                            fmt.Println(v3)
                        } else {
                            _ = len(v27)
                        }
                        _ = v33
                        _ = v34
                    } else {
                        if len(v27) > 0 {
                            v27[0] = -1
                        }
                        fmt.Println(v2)
                    }
                    _ = v32
                }
            } else {
                var v35 string = "foo bar"
                _ = v35
            }
            for fi36 := 0; fi36 < 17; fi36++ {
                switch v3 {
                case 9223372036854775807:
                    var v37 int16 = -727
                    _ = v37
                case 0:
                    v27 = append(v27, (fi36 + fi36))
                default:
                    v38 := GenS2{F0: 905, F1: -621}
                    _ = v38
                }
                v1 -= v1
                _ = fi36
            }
        }
    }
    if (-32768 > 1000) {
        switch 4294967295 {
        case 0:
            v1 = v1
        case -924:
            var v39 bool = true
            _ = v39
        case -9223372036854775808:
            switch 32768 {
            case -1:
                if (42 == 0) {
                    if false {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v1, v2, v3)
                        case "hello":
                            _ = len(v27)
                        case "\n\t":
                            v3 -= (v2 | (v3 + 0))
                        }
                    } else {
                        var v40 int64 = 1
                        fmt.Println(v2)
                        var v41 int = (v3 & 32767)
                        _ = v40
                        _ = v41
                    }
                    var v42 float32 = (float32(-2.25) / -2.25)
                    _ = v42
                } else {
                    if false {
                        fmt.Println(v1, v1)
                        switch v2 {
                        case 1:
                            v27 = append(v27, v3)
                        case 761:
                            var v43 uint8 = 0
                            _ = v43
                        default:
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        }
                        if false {
                            var v44 uint64 = genFn0(int64(v1), 358, v3)
                            _ = v44
                        }
                    }
                }
            case 9223372036854775807:
                if (!true) {
                    _ = len(v27)
                    switch "foo bar" {
                    case "":
                        var v45 uint8 = 127
                        _ = v45
                    default:
                        switch "hello" {
                        case "a":
                            var v46 uint64 = 2
                            _ = v46
                        case "\n\t":
                            v3 = 4294967295
                        case "foo bar":
                            v3 += v3
                        }
                    }
                }
            case -538:
                v3 -= v2
            default:
                for fi47 := 0; fi47 < 7; fi47++ {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        var v49 int64 = 87
                        var v50 uint = 115
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    var v51 int8 = 127
                    v52 := GenS4{F0: 2, F1: true}
                    _ = fi47
                    _ = v51
                    _ = v52
                }
            }
        }
    }
    v27 = append(v27, (v3 | (v3 | (v3 - (v2 % (((-(v2))) | 1))))))
    fmt.Println(v3, v2, v1)
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    if true {
        var v53 int64 = 1
        v53 = v53
        _ = v53
    } else {
        if true {
            v3 += v2
            var v54 string = (("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))) + "foo bar")
            _ = v54
        }
        switch (v2 & v3) {
        case 250:
            var v55 uint8 = 66
            _ = v55
        }
        var v56 bool = true
        _ = v56
    }
    var v57 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v57
}
