package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 int64, p1 int64, p2 bool) uint32 {
    fmt.Println(p0)
    var v4 uint = 65535
    var v5 string = ("hello" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := []uint32{42}
    v2 := []uint32{}
    v3 := []string{"hello", "hello", "hello", "\n\t"}
    for fi6 := 0; fi6 < 7; fi6++ {
        switch ("\n\t" + ("\n\t" + ("\n\t" + ""))) {
        case "foo bar":
            v3 = append(v3, "a")
        default:
            var v7 uint8 = 255
            _ = v7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 1, F1: 42}
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            v10 := GenS1{F0: -364, F1: genFn0(9223372036854775807, 0, (-843 >= 1))}
            fmt.Println(fi6, fi6)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(fi12)
            for fi13 := 0; fi13 < 19; fi13++ {
                switch "a" {
                case "foo bar":
                    v14 := GenS1{F0: 42, F1: 4294967295}
                    _ = v14
                case "":
                    _ = len(v3)
                case "hello":
                    var v15 int = (-(fi12))
                    _ = v15
                default:
                    v16 := GenS2{F0: 0, F1: 255}
                    _ = v16
                }
                var v17 int32 = -2147483648
                var v18 uint64 = 324
                _ = fi13
                _ = v17
                _ = v18
            }
            var v19 int32 = -2147483648
            _ = fi12
            _ = v19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int8 = 100
            _ = len(v2)
            _ = fi20
            _ = v21
        }
    }
    var v22 float32 = (-2.25 - 3.14159)
    v22 -= v22
    _ = len(v2)
    v23 := GenS2{F0: -9223372036854775808, F1: 417}
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            switch fi24 {
            case -9223372036854775808:
                switch 4294967295 {
                case 420:
                    v22 += (v22 - 1.5)
                case 142:
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (1 < 121) {
                            if true {
                                v22 -= 0.0
                                fmt.Println(v22, fi25, fi25)
                                var v26 float32 = -2.25
                                _ = v26
                            }
                        }
                    } else {
                        v22 -= v22
                        v2 = append(v2, genFn0(-9223372036854775808, 0, false))
                    }
                default:
                    var v27 uint8 = 162
                    _ = v27
                }
                fmt.Println(fi25)
                var v28 bool = (false && true)
                _ = v28
            case -1:
                fmt.Println(fi24)
            case 597:
                var v29 int64 = 0
                switch 4294967295 {
                case -9223372036854775808:
                    v30 := GenS2{F0: v29, F1: 32767}
                    _ = v30
                case 2147483647:
                    v22 -= 1.5
                default:
                    v29 = v29
                }
                var v31 uint16 = 1
                _ = v29
                _ = v31
            }
            v22 -= float32((0.0 - 3.14159))
            v32 := GenS3{F0: 47, F1: 532}
            _ = fi25
            _ = v32
        }
        fmt.Println(v22, fi24)
        _ = fi24
    }
    _ = len(v2)
    var v33 int64 = -1
    var v34 int32 = -1
    var v35 int64 = (v33 % (((v33 >> ((255) & 63))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
