package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int8, p1 int32) string {
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            var v7 uint32 = 2
            v8 := []int{fi6, (fi6 ^ (fi6 * fi6)), fi6, -9223372036854775808}
            if false {
                fmt.Println(v7, p1)
            } else {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int32 = p1
                    var v11 int16 = 241
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 float32 = 3.14159
                    _ = fi12
                    _ = v13
                }
                switch 42 {
                case 2147483647:
                    var v14 int64 = 9223372036854775807
                    _ = v14
                default:
                    if (!(true || false)) {
                        v7 = (v7 << ((201) & 63))
                    }
                }
            }
            _ = v7
            _ = v8
        } else {
            v15 := []int{(fi6 >> ((1) & 63)), (fi6 * 372), int(p0)}
            switch fi6 {
            case 4294967295:
                var v16 uint32 = 500
                _ = v16
            }
            _ = v15
        }
        var v17 int = (fi6 / (((fi6 - int(p0))) | 1))
        _ = fi6
        _ = v17
    }
    p0 -= 127
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint = 18446744073709551615
    var v5 uint32 = 2
    fmt.Println(v4)
    var v18 bool = (2147483647 < (v5 / ((v5) | 1)))
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: 831, F1: 42, F2: 42}
        _ = fi19
        _ = v20
    }
    delete(v2, 60)
    var v21 int32 = 0
    v4 = (v4 ^ 2)
    if v18 {
        v22 := GenS2{F0: 0, F1: v18, F2: 42}
        v4 += (v4 * (v4 >> ((v4) & 63)))
        _ = v22
    }
    if v18 {
        if true {
            for fi23 := 0; fi23 < 14; fi23++ {
                v18 = (-862 <= 0)
                var v24 float64 = 0.1
                _ = fi23
                _ = v24
            }
            if (800 >= 255) {
                v25 := GenS2{F0: -128, F1: ((((("a" + "hello") + ("\n\t" + "hello")) + "hello") + ("a" + "hello")) < "foo bar"), F2: 18446744073709551615}
                switch "" {
                case "\n\t":
                    v5 = (^(v5))
                case "foo bar":
                    _ = len(v2)
                case "":
                    v26 := GenS2{F0: -47, F1: true, F2: 4294967295}
                    _ = v26
                default:
                    switch 127 {
                    case 830:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            v28 := GenS1{F0: 65535, F1: 0, F2: 18446744073709551615}
                            var v29 bool = (!(v5 != 4294967295))
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    case -1:
                        if (v18 || (!(0 != -1))) {
                            var v30 int32 = (-(v21))
                            var v31 int64 = -9223372036854775808
                            fmt.Println(v5, v4, v4)
                            _ = v30
                            _ = v31
                        } else {
                            delete(v2, int(v5))
                        }
                    case 752:
                        var v32 float32 = 1e10
                        _ = v32
                    default:
                        v4 = 18446744073709551615
                    }
                }
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 int64 = -1
                    _ = fi33
                    _ = v34
                }
                _ = v25
            }
        } else {
            fmt.Println(v5, v18)
        }
    } else {
        v2[9223372036854775807] = -89
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 uint32 = 2147483648
        fmt.Println(v18, v21, fi35)
        _ = fi35
        _ = v36
    }
    v37 := GenS3{F0: 9223372036854775807, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
