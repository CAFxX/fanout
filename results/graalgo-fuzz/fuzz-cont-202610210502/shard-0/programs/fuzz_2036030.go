package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 string, p1 uint16, p2 uint32) int64 {
    v6 := GenS1{F0: 32767, F1: 42}
    fmt.Println(p0, p2)
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint8 = 1
            v9 += 127
            _ = fi8
            _ = v9
        }
        p0 = (p0 + (p0 + p0))
        v10 := []int{(fi7 - fi7), 32767, (fi7 + 487), fi7}
        _ = fi7
        _ = v10
    }
    p0 = p0
    switch -2147483648 {
    case 237:
        p1 = p1
    case -155:
        v11 := GenS2{F0: 713}
        _ = v11
    case 32767:
        v12 := GenS3{F0: -128}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    var v4 int64 = -68
    v5 := []string{(("\n\t" + "hello") + "a"), "foo bar", "\n\t"}
    v3[0] = 4294967295
    var v13 float64 = (1e10 - 3.14159)
    if (118 == 127) {
        switch "\n\t" {
        case "\n\t":
            switch "" {
            case "foo bar":
                for fi14 := 0; fi14 < 10; fi14++ {
                    _ = len(v1)
                    _ = fi14
                }
            default:
                var v15 uint8 = 42
                _ = v15
            }
        default:
            switch -2147483648 {
            case 9223372036854775807:
                var v16 uint64 = 18446744073709551615
                _ = v16
            case 847:
                v4 = (v4 & v4)
            }
        }
        var v17 uint = 18446744073709551615
        _ = v17
    }
    fmt.Println(v2, v4, v2)
    var v18 int16 = 788
    if ((((v4 / ((v4) | 1)) < (v4 - v4)) || false) || ((v4 & genFn0("a", 0, 442)) != (-(v4)))) {
        var v19 uint16 = 2
        for fi20 := 0; fi20 < 16; fi20++ {
            if false {
                v4 -= 799
                v18 = 42
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v21 string = "a"
                _ = v21
            }
            v4 -= (v4 & v4)
            if false {
                if len(v5) > 0 {
                    v5[0] = (("\n\t" + (("foo bar" + "") + "foo bar")) + "a")
                }
            } else {
                v19 -= (v19 & v19)
            }
            _ = fi20
        }
        _ = v19
    } else {
        switch "hello" {
        case "foo bar":
            switch -403 {
            case -2147483648:
                switch "\n\t" {
                case "\n\t":
                    v22 := GenS1{F0: v18, F1: 42}
                    _ = v22
                case "a":
                    if false {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v24 := v3[2]
                            v4 += v4
                            v18 -= v18
                            _ = fi23
                            _ = v24
                        }
                        switch "\n\t" {
                        case "":
                            var v25 uint64 = 366
                            _ = v25
                        default:
                            fmt.Println(v2)
                        }
                    } else {
                        fmt.Println(v13)
                    }
                default:
                    v26 := GenS4{F0: 255, F1: (373 != 969)}
                    _ = v26
                }
            }
        case "":
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 bool = (32767 > 544)
                switch fi27 {
                case 32767:
                    for fi29 := 0; fi29 < 5; fi29++ {
                        var v30 int32 = 2147483647
                        var v31 int16 = v18
                        var v32 int32 = int32(v18)
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                case 0:
                    v28 = (!((!v28) && v28))
                }
                v2 = 127
                _ = fi27
                _ = v28
            }
        default:
            v13 = v13
        }
        var v33 uint32 = 2147483647
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int64 = (v4 + (v4 % (((v4 * 42)) | 1)))
            var v36 string = ("hello" + "hello")
            if ((0 == 1) && true) {
                v36 = v36
                var v37 int8 = -43
                v13 = (v13 + v13)
                _ = v37
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    }
    var v38 int32 = -2147483648
    v13 -= 0.1
    v18 = v18
    v13 = (((((v13 + (0.0 / v13)) - float64(float32((v13 / (1.5 * 1e10))))) / v13) * float64((-2.25 * (((0.0 / -2.25) / 0.1) - -2.25)))) / (float64(0.0) / v13))
    v18 = (v18 * (v18 + (v18 | v18)))
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v39 string = "a"
        _ = v39
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "":
        delete(v3, 1)
    default:
        _ = len(v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
