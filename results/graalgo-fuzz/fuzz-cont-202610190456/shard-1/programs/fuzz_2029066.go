package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 int8) uint {
    v5 := []int{-69, 127, 53, 209}
    for fi6 := 0; fi6 < 12; fi6++ {
        if (1 > 4294967295) {
            if len(v5) > 0 {
                v5[0] = fi6
            }
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint8 = 2
                p0 += 100
                _ = fi7
                _ = v8
            }
        } else {
            if false {
                v5 = append(v5, 2147483647)
                p0 = -1
                if true {
                    p0 += p0
                    switch 9223372036854775807 {
                    case -572:
                        switch "foo bar" {
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                var v9 int8 = 42
                                _ = v9
                            default:
                                p0 = p0
                            }
                        case "foo bar":
                            var v10 int32 = 121
                            _ = v10
                        }
                    default:
                        if true {
                            p0 = (p0 ^ -26)
                        } else {
                            var v11 int = 127
                            if ((!(v11 != 426)) && false) {
                                switch fi6 {
                                case 467:
                                    for fi12 := 0; fi12 < 2; fi12++ {
                                        var v13 float64 = 0.1
                                        _ = fi12
                                        _ = v13
                                    }
                                case -9223372036854775808:
                                    for fi14 := 0; fi14 < 4; fi14++ {
                                        v15 := GenS1{F0: (p0 / ((p0) | 1)), F1: 246}
                                        fmt.Println(p0)
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                            }
                            var v16 uint8 = 1
                            _ = v11
                            _ = v16
                        }
                    }
                    p0 = p0
                }
            } else {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 int32 = -343
                    switch 4294967295 {
                    case 4294967295:
                        v18 -= 2147483647
                    case 128:
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            v19 := GenS2{F0: (-(v18)), F1: true}
                            _ = v19
                        default:
                            p0 -= (p0 >> ((42) & 63))
                            v20 := GenS3{F0: (!(!(!true)))}
                            var v21 int32 = (-(v18))
                            _ = v20
                            _ = v21
                        }
                    default:
                        switch "hello" {
                        case "":
                            v5 = append(v5, fi17)
                        case "a":
                            v18 -= -1
                        case "hello":
                            p0 = (p0 % (((p0 * (p0 - (-(p0))))) | 1))
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                p0 += 42
            }
        }
        switch "a" {
        case "hello":
            if len(v5) > 0 {
                v5[0] = (fi6 ^ 588)
            }
        case "\n\t":
            fmt.Println(fi6)
        }
        _ = fi6
    }
    p0 = (p0 ^ 1)
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    var v3 float32 = 3.14159
    var v4 bool = (false || false)
    v4 = (581 != 1)
    switch 0 {
    case -2147483648:
        v1 = -923
    }
    v4 = (float64((v3 / v3)) == (0.1 * -2.25))
    v22 := GenS1{F0: 100, F1: 0}
    var v23 uint16 = 170
    var v24 uint16 = (v23 >> ((814) & 63))
    var v25 uint16 = (v23 - (v23 * v23))
    for fi26 := 0; fi26 < 3; fi26++ {
        if v4 {
            v27 := GenS3{F0: v4}
            var v28 int = 4294967295
            _ = v27
            _ = v28
        } else {
            v24 = v23
            switch "\n\t" {
            case "hello":
                v3 = (((3.14159 / (((v3 / float32((-2.25 + -2.25))) / (float32(3.14159) - 1.5)) - float32(-2.25))) + float32((0.1 * (1e10 + 0.1)))) - 3.14159)
            }
            v1 = v1
        }
        v23 += 1
        v29 := GenS1{F0: 127, F1: 449}
        _ = fi26
        _ = v29
    }
    _ = len(v2)
    v30 := []int{-475, -28}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
