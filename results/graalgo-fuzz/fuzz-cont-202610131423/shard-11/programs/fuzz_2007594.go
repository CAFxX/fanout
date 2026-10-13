package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int32; F2 string }
type GenS7 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0() uint32 {
    var v5 bool = (!(false || true))
    v6 := GenS1{F0: 863, F1: 914, F2: ((!(!(65535 < 0))) || false)}
    var v7 int64 = 1
    if (255 <= 255) {
        switch 9223372036854775807 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 float32 = 1e10
                fmt.Println(v5, v5, v5)
                _ = fi8
                _ = v9
            }
        case 464:
            v10 := GenS2{F0: 128, F1: 32767, F2: (("a" + "a") + "hello")}
            _ = v10
        default:
            switch 568 {
            case -873:
                v7 = v7
            case -503:
                v5 = ((1 <= 146) || (!((v7 != v7) && v5)))
            default:
                v11 := []int{2147483647}
                _ = v11
            }
        }
        v12 := []int{127, 32767, -539, -182}
        _ = v12
    }
    var v13 int32 = 2147483647
    _ = v13
    _ = v5
    _ = v6
    _ = v7
    return 793
}
func genFn1(p0 uint16, p1 uint, p2 int8) int {
    v14 := []int{1, 0, 42}
    fmt.Println(p1)
    v15 := GenS3{F0: 4294967295, F1: p1, F2: 188}
    var v16 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 32767
}

func main() {
    v1 := []string{"\n\t", "a"}
    v2 := make(map[uint32]int)
    var v3 uint = 1
    v4 := []int{}
    delete(v2, 0)
    v2[genFn0()] = 2147483647
    if (-9223372036854775808 > 9223372036854775807) {
        for fi17 := 0; fi17 < 3; fi17++ {
            switch ("foo bar" + "a") {
            case "\n\t":
                var v18 uint = (v3 >> (((v3 ^ 18446744073709551615)) & 63))
                _ = v18
            case "hello":
                var v19 uint8 = 149
                _ = v19
            default:
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 float32 = 0.0
                    _ = fi20
                    _ = v21
                }
            }
            if true {
                var v22 uint64 = 570
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 float32 = 0.0
                    if (42 > 1) {
                        v3 = 823
                        v3 = (v3 & v3)
                    } else {
                        v4 = append(v4, fi17)
                        v22 -= (^(v22))
                        v3 = v3
                        v24 -= 3.14159
                    }
                    v25 := GenS1{F0: (v22 << ((2) & 63)), F1: -391, F2: (true && true)}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = v22
            } else {
                if (!(1000 >= 32767)) {
                    if true {
                        var v26 int32 = -326
                        var v27 int64 = -1
                        v28 := GenS4{F0: 0, F1: (v27 % ((v27) | 1)), F2: 337}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                } else {
                    if false {
                        v29 := GenS5{F0: "foo bar"}
                        for fi30 := 0; fi30 < 15; fi30++ {
                            delete(v2, 551)
                            var v31 int64 = 9223372036854775807
                            v32 := GenS3{F0: 272, F1: (v3 & 1), F2: 54}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    } else {
                        v3 += v3
                    }
                    switch "" {
                    case "foo bar":
                        fmt.Println(fi17, fi17)
                    }
                }
                for fi33 := 0; fi33 < 19; fi33++ {
                    v3 = 748
                    _ = fi33
                }
                delete(v2, 0)
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                v35 := GenS6{F0: 0, F1: 2147483647, F2: ""}
                v36 := v2[2]
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi17
        }
        v3 += (v3 & v3)
        v3 += v3
    } else {
        v3 -= (^(v3))
        var v37 int32 = 1
        _ = v37
    }
    v4 = append(v4, genFn1(645, (v3 ^ v3), 0))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v38 := v2[2147483648]
    v39 := v2[813]
    v40 := GenS6{F0: 40, F1: -1, F2: "hello"}
    v39 += 2147483647
    for fi41 := 0; fi41 < 9; fi41++ {
        v42 := GenS4{F0: -128, F1: 1, F2: 2147483647}
        var v43 uint64 = 502
        _ = fi41
        _ = v42
        _ = v43
    }
    v39 += v39
    fmt.Println(v38, v39)
    for fi44 := 0; fi44 < 15; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            v2[428] = (^(fi44))
            _ = fi45
        }
        var v46 float32 = float32((3.14159 * 0.1))
        if (true || ((true && false) || false)) {
            var v47 int8 = 42
            v46 = v46
            _ = v47
        } else {
            v48 := GenS7{F0: "hello", F1: 1, F2: 2}
            _ = v48
        }
        _ = fi44
        _ = v46
    }
    switch "hello" {
    case "a":
        v49 := GenS2{F0: 255, F1: 32767, F2: "a"}
        _ = v49
    case "foo bar":
        var v50 int16 = 42
        _ = v50
    }
    fmt.Println(v38, v38, v3)
    v2[2147483647] = v39
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
