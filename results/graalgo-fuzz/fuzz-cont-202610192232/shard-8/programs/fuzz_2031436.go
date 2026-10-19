package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0() uint {
    var v7 int64 = -1
    fmt.Println(v7)
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 2, F1: -59, F2: v7}
        if (true || false) {
            switch "foo bar" {
            case "":
                switch fi8 {
                case 32767:
                    if ((((^(v7)) != v7) && true) && false) {
                        v10 := []int{}
                        if len(v10) > 0 {
                            v10[0] = (fi8 / ((43) | 1))
                        }
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 42, F1: 47, F2: (v7 % ((v7) | 1))}
                        fmt.Println(v7, fi8)
                        _ = v11
                    }
                case -328:
                    fmt.Println(fi8)
                default:
                    var v12 int = (fi8 / ((-687) | 1))
                    _ = v12
                }
            case "foo bar":
                v13 := []int{2147483647, (fi8 << ((42) & 63)), fi8}
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS2{F0: true, F1: "\n\t", F2: 2}
                    v16 := []int{42}
                    v7 += (v7 | (v7 / ((v7) | 1)))
                    _ = len(v16)
                    v7 = v7
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                var v17 string = ("foo bar" + "a")
                _ = v17
            }
        } else {
            var v18 bool = (!true)
            for fi19 := 0; fi19 < 2; fi19++ {
                v18 = (0.0 > 0.1)
                v20 := GenS3{F0: v7, F1: 2147483648}
                if v18 {
                    v21 := GenS4{F0: v7, F1: 1000, F2: 2}
                    v7 += (v7 % (((v7 - (v7 ^ (v7 << ((65535) & 63))))) | 1))
                    var v22 int64 = (v7 >> ((2) & 63))
                    _ = v21
                    _ = v22
                }
                if v18 {
                    v18 = false
                    var v23 int32 = -2147483648
                    _ = v23
                } else {
                    if false {
                        if true {
                            if false {
                                var v24 int64 = (v7 & v7)
                                v7 -= (v24 << ((927) & 63))
                                switch -2147483648 {
                                case -1:
                                    var v25 string = ("foo bar" + "hello")
                                    _ = v25
                                }
                                _ = v24
                            }
                        } else {
                            v18 = true
                            v7 += v7
                            switch "" {
                            case "a":
                                fmt.Println(v7, fi19)
                            }
                        }
                        v26 := []int{fi8}
                        _ = v26
                    }
                    if v18 {
                        v27 := GenS3{F0: v7, F1: 698}
                        _ = v27
                    } else {
                        switch "\n\t" {
                        case "":
                            if true {
                                v7 = v7
                                fmt.Println(fi19, fi8)
                            } else {
                                var v28 string = ("" + ("\n\t" + "a"))
                                _ = v28
                            }
                        default:
                            fmt.Println(fi8, v7, fi8)
                        }
                        v18 = (v18 && ((943 != 65535) || true))
                    }
                    switch fi19 {
                    case 32768:
                        if true {
                            v29 := GenS2{F0: (1 >= 269), F1: "a", F2: 164}
                            _ = v29
                        } else {
                            if true {
                                fmt.Println(v18, v18, v18)
                                v7 = v7
                            }
                        }
                    case 536:
                        v7 += (v7 - (v7 * 0))
                    }
                }
                _ = fi19
                _ = v20
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                v18 = (v18 || v18)
                _ = fi30
            }
            _ = v18
        }
        _ = fi8
        _ = v9
    }
    if (1 < 980) {
        v7 = (v7 | v7)
        v7 -= -1
        v7 += 1
    }
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 50
    var v3 string = (("a" + "") + "")
    var v4 int16 = 1000
    var v5 uint8 = (v2 << ((18446744073709551615) & 63))
    v6 := make(map[int]bool)
    v3 = ((("" + v3) + ((v3 + ("hello" + v3)) + v3)) + v3)
    v31 := []int{127, 0}
    v2 = v2
    delete(v1, -617)
    v5 = v2
    var v32 string = ("hello" + (v3 + v3))
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    var v33 uint = 2
    v31 = append(v31, -936)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
