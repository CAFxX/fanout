package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint }
func genFn0() uint {
    var v5 string = ("hello" + ("hello" + "foo bar"))
    v5 = (((v5 + v5) + "foo bar") + v5)
    v6 := GenS1{F0: 400}
    v7 := GenS2{F0: 42, F1: 0, F2: 2}
    _ = v5
    _ = v6
    _ = v7
    return 385
}
func genFn1(p0 string) uint64 {
    p0 = ("hello" + "")
    switch p0 {
    case "foo bar":
        fmt.Println(p0)
    default:
        switch 32768 {
        case -696:
            if (!true) {
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 uint8 = 127
                    if true {
                        for fi10 := 0; fi10 < 15; fi10++ {
                            switch p0 {
                            case "\n\t":
                                if ((242 >= 231) || true) {
                                    if (!(!false)) {
                                        var v11 int16 = 42
                                        _ = v11
                                    }
                                }
                            default:
                                if false {
                                    v12 := GenS3{F0: int64(v9), F1: 65535, F2: 402}
                                    v13 := []int{(^(fi8)), (-(fi8))}
                                    _ = v12
                                    _ = v13
                                }
                            }
                            switch "hello" {
                            case "hello":
                                p0 = "\n\t"
                            case "\n\t":
                                var v14 int64 = -406
                                _ = v14
                            default:
                                v15 := []int{fi8, (fi10 / ((fi10) | 1)), -2147483648, (fi10 % ((704) | 1))}
                                _ = v15
                            }
                            var v16 uint = 480
                            _ = fi10
                            _ = v16
                        }
                    } else {
                        v17 := []int{973, fi8, (fi8 / (((-(fi8))) | 1)), fi8}
                        if true {
                            for fi18 := 0; fi18 < 7; fi18++ {
                                p0 = p0
                                _ = fi18
                            }
                        }
                        _ = v17
                    }
                    _ = fi8
                    _ = v9
                }
                v19 := GenS2{F0: 2147483647, F1: 1, F2: 1}
                _ = v19
            }
        case 2147483647:
            p0 = p0
        }
    }
    v20 := GenS1{F0: 989}
    var v21 uint = 255
    fmt.Println(v21, p0)
    _ = p0
    _ = v20
    _ = v21
    return 425
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = false
    var v3 float64 = (v1 / (v1 - v1))
    v4 := make(map[string]bool)
    var v22 uint16 = 2
    fmt.Println(v2)
    v23 := GenS3{F0: 42, F1: (^(v22)), F2: (v22 / ((968) | 1))}
    v4[""] = (uint64(v22) != 2)
    if v2 {
        var v24 int8 = 100
        var v25 uint64 = 666
        _ = v24
        _ = v25
    } else {
        v2 = (1.5 != 3.14159)
        var v26 uint64 = 2
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v2 = (v2 && false)
        v28 := GenS4{F0: 42}
        if false {
            if (!(((0.0 - 0.0) * 3.14159) == 0.1)) {
                v29 := GenS5{F0: v22}
                v30 := GenS2{F0: 2147483648, F1: 891, F2: 255}
                _ = v29
                _ = v30
            } else {
                if v2 {
                    var v31 string = ""
                    v1 -= v1
                    _ = v31
                } else {
                    v32 := GenS2{F0: 42, F1: 263, F2: 285}
                    _ = v32
                }
            }
            var v33 string = (("a" + "hello") + "a")
            var v34 int32 = 1
            _ = v33
            _ = v34
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v2)
    v35 := GenS6{F0: 259, F1: 414}
    var v36 int = 1
    fmt.Println(v3)
    fmt.Println(v1, v1, v3)
    if (false && (2147483648 != 1)) {
        var v37 uint32 = 0
        _ = v37
    }
    v38 := GenS4{F0: -1}
    switch v36 {
    case 9223372036854775807:
        for fi39 := 0; fi39 < 14; fi39++ {
            _ = len(v4)
            v40 := GenS4{F0: 0}
            v22 -= (v22 % (((v22 * v22)) | 1))
            _ = fi39
            _ = v40
        }
    case 128:
        v4[("" + "a")] = true
    case 471:
        v1 = -2.25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
}
