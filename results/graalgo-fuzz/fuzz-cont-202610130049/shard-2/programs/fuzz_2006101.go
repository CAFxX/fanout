package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 string }
func genFn0() int64 {
    var v7 string = "a"
    switch 128 {
    case -9223372036854775808:
        switch 4294967295 {
        case -896:
            var v8 float32 = 0.0
            _ = v8
        case -838:
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v7 = "a"
                    _ = fi10
                }
                _ = fi9
            }
        }
    default:
        v11 := []int{}
        _ = v11
    }
    var v12 int32 = 2147483647
    switch -9223372036854775808 {
    case 127:
        v7 = ("\n\t" + v7)
    case -1:
        if false {
            switch v7 {
            case "\n\t":
                for fi13 := 0; fi13 < 14; fi13++ {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v15 := GenS1{F0: (v7 + v7), F1: 128}
                        fmt.Println(fi13, v7)
                        fmt.Println(v7)
                        _ = fi14
                        _ = v15
                    }
                    v12 = v12
                    _ = fi13
                }
            case "foo bar":
                var v16 int = 32768
                _ = v16
            case "":
                var v17 int64 = 9223372036854775807
                _ = v17
            default:
                switch -9223372036854775808 {
                case 2147483647:
                    v7 = (v7 + (v7 + (("a" + "") + "hello")))
                case -2147483648:
                    v18 := GenS1{F0: "", F1: 128}
                    _ = v18
                case 272:
                    v7 = v7
                default:
                    v7 = ""
                }
            }
        } else {
            switch ((("\n\t" + "hello") + "\n\t") + v7) {
            case "foo bar":
                var v19 int32 = v12
                _ = v19
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS1{F0: (v7 + ((v7 + "hello") + v7)), F1: 255}
                for fi22 := 0; fi22 < 14; fi22++ {
                    if true {
                        switch fi20 {
                        case -982:
                            fmt.Println(v7, fi22, fi20)
                        default:
                            switch v7 {
                            case "a":
                                v7 = "foo bar"
                                switch v7 {
                                case "hello":
                                    v23 := GenS1{F0: v7, F1: 2}
                                    _ = v23
                                }
                                v24 := GenS1{F0: v7, F1: 39}
                                var v25 float32 = (1e10 + -2.25)
                                _ = v24
                                _ = v25
                            case "hello":
                                v26 := []int{}
                                _ = v26
                            default:
                                if (((float32(3.14159) - (float32(((3.14159 / 3.14159) + (0.0 / float64(3.14159)))) - 0.1)) + (0.1 - (1.5 - ((0.0 / -2.25) / 1.0)))) != 0.0) {
                                    var v27 int = -1
                                    _ = v27
                                }
                            }
                        }
                        v12 = v12
                    }
                    v12 = v12
                    if false {
                        v12 -= v12
                        v28 := GenS1{F0: (v7 + v7), F1: 2}
                        if (true && true) {
                            v7 = v7
                            v7 = "\n\t"
                            switch v7 {
                            case "\n\t":
                                v29 := GenS2{F0: true}
                                _ = v29
                            case "foo bar":
                                v30 := GenS3{F0: 2, F1: 1, F2: (((((v7 + v7) + ((v7 + v7) + "a")) + "a") + "a") + v7)}
                                _ = v30
                            default:
                                if false {
                                    v31 := []int{}
                                    if true {
                                        if len(v31) > 0 {
                                            v31[0] = (fi22 << ((18446744073709551615) & 63))
                                        }
                                        var v32 float32 = float32(3.14159)
                                        _ = v32
                                    }
                                    _ = v31
                                }
                            }
                        }
                        _ = v28
                    } else {
                        v33 := GenS3{F0: 71, F1: 55, F2: v7}
                        _ = v33
                    }
                    _ = fi22
                }
                var v34 bool = true
                _ = fi20
                _ = v21
                _ = v34
            }
            v35 := GenS2{F0: (float32(((float64(float32(0.1)) * (0.1 - -2.25)) / ((3.14159 + 1.5) / 1.5))) >= 1e10)}
            _ = v35
        }
    case 128:
        for fi36 := 0; fi36 < 20; fi36++ {
            fmt.Println(v12, fi36)
            if (true || (!false)) {
                v12 = (v12 >> ((42) & 63))
            }
            v37 := GenS4{F0: 55}
            _ = fi36
            _ = v37
        }
    }
    v38 := GenS5{F0: -1, F1: v7}
    _ = v12
    _ = v38
    _ = v7
    return -354
}
func genFn1(p0 bool, p1 bool) int8 {
    v39 := []int{4294967295, 1}
    switch 379 {
    case 42:
        if p0 {
            if len(v39) > 0 {
                v39[0] = 127
            }
        } else {
            v39 = append(v39, -1)
            p0 = true
            v40 := GenS2{F0: ("\n\t" == ("" + "\n\t"))}
            _ = v40
        }
    case -89:
        fmt.Println(p1, p0)
    }
    var v41 uint8 = 0
    _ = len(v39)
    _ = p0
    _ = p1
    _ = v39
    _ = v41
    return 127
}

func main() {
    var v1 float64 = -2.25
    var v2 string = ""
    var v3 uint32 = 4294967295
    var v4 bool = false
    v5 := make(map[int]int)
    var v6 uint8 = 0
    v42 := []int{}
    v1 = (0.1 * 1.5)
    var v43 int64 = -1
    for fi44 := 0; fi44 < 3; fi44++ {
        v3 += v3
        switch "foo bar" {
        case "hello":
            fmt.Println(v43)
        }
        v3 = 4294967295
        _ = fi44
    }
    var v45 uint = 494
    v46 := GenS2{F0: (v4 && v4)}
    v2 = v2
    v43 -= 1
    v2 = ((("\n\t" + (v2 + "")) + (("a" + v2) + v2)) + v2)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
