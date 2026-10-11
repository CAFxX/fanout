package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 uint8; F1 uint; F2 string }
func genFn0() string {
    v7 := GenS1{F0: 2}
    var v8 int = -1
    v9 := GenS2{F0: v8, F1: 497}
    _ = v7
    _ = v8
    _ = v9
    return (("" + ("\n\t" + "a")) + "a")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (-2.25 / -2.25)
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    var v6 int64 = 9223372036854775807
    v10 := []int{1, 1}
    var v11 float32 = 0.0
    var v12 uint16 = 255
    if false {
        switch 127 {
        case -9223372036854775808:
            if true {
                var v13 int16 = 1000
                v14 := GenS3{F0: (true || (v11 == ((((1.5 * (v11 - float32(v2))) - 3.14159) / (v11 * -2.25)) / v11))), F1: 687, F2: v5}
                switch 908 {
                case 42:
                    if true {
                        var v15 uint32 = 0
                        _ = v15
                    } else {
                        switch 32768 {
                        case 32767:
                            var v16 int = 32768
                            _ = v16
                        case -133:
                            _ = len(v1)
                        default:
                            if (0 == 255) {
                                switch v5 {
                                case "a":
                                    v17 := GenS2{F0: 128, F1: 0}
                                    _ = v17
                                default:
                                    var v18 float64 = 0.0
                                    _ = v18
                                }
                                if ((v12 + v12) != 32767) {
                                    for fi19 := 0; fi19 < 11; fi19++ {
                                        delete(v4, (^(fi19)))
                                        var v20 string = ""
                                        _ = fi19
                                        _ = v20
                                    }
                                } else {
                                    v21 := GenS1{F0: 2147483647}
                                    _ = v21
                                }
                                var v22 uint32 = 4294967295
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 16; fi23++ {
                                    switch "foo bar" {
                                    case "a":
                                        _ = len(v4)
                                    case "\n\t":
                                        v1[65535] = (false || (!true))
                                    default:
                                        var v24 int64 = (v6 - v6)
                                        _ = v24
                                    }
                                    v25 := GenS2{F0: fi23, F1: 2147483648}
                                    v3[""] = v5
                                    _ = fi23
                                    _ = v25
                                }
                            }
                        }
                    }
                case -9223372036854775808:
                    switch 0 {
                    case 0:
                        delete(v1, 4294967295)
                    case -319:
                        var v26 int64 = v6
                        _ = v26
                    }
                default:
                    v2 -= (v2 / (1.5 / 1.5))
                }
                _ = v13
                _ = v14
            }
        default:
            switch v5 {
            case "":
                v2 = 1e10
            default:
                for fi27 := 0; fi27 < 7; fi27++ {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 uint64 = 32
                        switch fi28 {
                        case -2147483648:
                            v2 = v2
                        case 784:
                            v5 = v5
                        case 32767:
                            v30 := GenS4{F0: 255, F1: 367, F2: v5}
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
            }
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v10 = append(v10, fi31)
        var v32 int16 = 42
        v3[((("a" + v5) + "\n\t") + v5)] = v5
        _ = fi31
        _ = v32
    }
    var v33 uint64 = 42
    fmt.Println(v33, v33)
    _ = len(v10)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
