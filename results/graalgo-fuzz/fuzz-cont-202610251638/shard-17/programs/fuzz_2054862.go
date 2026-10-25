package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int; F1 uint }
func genFn0(p0 int8) string {
    v6 := []int{1, -1}
    p0 = (-(p0))
    _ = p0
    _ = v6
    return (("foo bar" + "\n\t") + "\n\t")
}

func main() {
    var v1 uint64 = 383
    var v2 int = -75
    var v3 uint64 = v1
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 363
    fmt.Println(v5, v1)
    if (1 < 1) {
        var v7 int = 42
        _ = v7
    } else {
        if (-562 >= -2147483648) {
            v1 -= (^(v3))
        } else {
            switch "" {
            case "":
                v3 = v1
            case "foo bar":
                delete(v4, v5)
            default:
                v1 = v1
            }
            var v8 uint16 = 65535
            if false {
                v9 := []int{v2, (v2 >> ((16) & 63)), (-(v2))}
                _ = v9
            }
            _ = v8
        }
    }
    switch "hello" {
    case "":
        switch -943 {
        case -732:
            for fi10 := 0; fi10 < 2; fi10++ {
                if (!false) {
                    if false {
                        switch "a" {
                        case "":
                            v2 = (v2 - (v2 - (v2 + 999)))
                        case "hello":
                            if ((true && false) && false) {
                                v1 = v3
                                var v11 int = fi10
                                v1 += 4294967295
                                _ = v11
                            }
                        case "foo bar":
                            v2 += (fi10 & 35)
                        default:
                            v12 := GenS1{F0: (("" + "") + "\n\t"), F1: -1}
                            _ = v12
                        }
                        v5 -= (v5 << ((0) & 63))
                        var v13 uint64 = 2
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 65535}
                        _ = v14
                    }
                    v15 := GenS3{F0: "hello", F1: 2, F2: 838}
                    if false {
                        v2 += fi10
                        var v16 string = (("" + "a") + "hello")
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v16 = v16
                            v4[v5] = 4294967295
                            if (true && (false && false)) {
                                fmt.Println(fi17, v16)
                                var v18 string = v16
                                _ = v18
                            } else {
                                fmt.Println(fi17, fi17)
                            }
                            v19 := GenS1{F0: v16, F1: -1}
                            _ = fi17
                            _ = v19
                        }
                        _ = v16
                    }
                    _ = v15
                }
                var v20 int32 = -1
                if (false && ((1.5 * 1.5) == 3.14159)) {
                    var v21 float32 = 1e10
                    _ = v21
                } else {
                    fmt.Println(v20)
                }
                _ = fi10
                _ = v20
            }
        }
    case "a":
        v2 = (v2 / (((v2 - v2)) | 1))
    default:
        var v22 int32 = 0
        _ = v22
    }
    fmt.Println(v3, v3, v5)
    v2 = v2
    var v23 int64 = 42
    if true {
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(v1, fi24, v23)
            v3 = (v3 - (^(v1)))
            _ = fi24
        }
    } else {
        v25 := GenS2{F0: 42}
        v26 := []int{(^(v2)), (v2 - v2), -240, v2}
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint64 = 1
                _ = fi28
                _ = v29
            }
            v30 := GenS2{F0: 2}
            _ = fi27
            _ = v30
        }
        _ = v25
        _ = v26
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int8 = 100
        _ = fi31
        _ = v32
    }
    if ((-2.25 >= (((-2.25 * (0.1 - -2.25)) + (((1.5 - float32(0.1)) + (float32(3.14159) - ((3.14159 + float32(0.1)) + 3.14159))) * (0.0 - (3.14159 / 0.1)))) / 1.0)) || (42 == 1)) {
        v4[(v5 ^ v5)] = v5
        if false {
            v1 -= (^(v3))
        }
        for fi33 := 0; fi33 < 2; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                v35 := []int{}
                _ = fi34
                _ = v35
            }
            v36 := GenS4{F0: 336, F1: 4294967295}
            v37 := []int{(^(fi33)), (fi33 >> ((0) & 63)), (v2 % ((-605) | 1)), (v2 >> ((264) & 63))}
            _ = fi33
            _ = v36
            _ = v37
        }
    } else {
        var v38 float64 = (-2.25 * 1.5)
        v3 = (v3 | v1)
        _ = v38
    }
    switch v2 {
    case -486:
        v39 := GenS5{F0: v2, F1: 4294967295}
        _ = v39
    default:
        for fi40 := 0; fi40 < 11; fi40++ {
            v23 += v23
            _ = fi40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
