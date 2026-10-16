package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    fmt.Println(v1, v3)
    var v4 float64 = 1.5
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = v3
        _ = fi5
        _ = v6
    }
    var v7 float64 = (float64(1.5) - (((v4 * float64(0.0)) - ((float64(-2.25) / (0.1 / v4)) * v4)) / 0.1))
    if false {
        var v8 uint = 18446744073709551615
        _ = v8
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v4, v4, v1)
            _ = fi9
        }
    }
    if (-1 <= -217) {
        v10 := GenS1{F0: 312, F1: false}
        _ = v10
    }
    switch 9223372036854775807 {
    case 32767:
        if (807 <= 559) {
            var v11 bool = (("foo bar" + "hello") != ("hello" + "a"))
            v12 := GenS1{F0: 407, F1: (-601 < 126)}
            var v13 int32 = 876
            _ = v11
            _ = v12
            _ = v13
        } else {
            if false {
                var v14 int64 = -814
                v2[42] = 2147483648
                v15 := GenS2{F0: (v1 >> ((304) & 63)), F1: true, F2: 2}
                _ = v14
                _ = v15
            } else {
                fmt.Println(v3)
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 string = ""
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 uint8 = 0
                        _ = fi18
                        _ = v19
                    }
                    if true {
                        v20 := []int{42, (fi16 % ((4294967295) | 1)), fi16}
                        var v21 string = ("\n\t" + "\n\t")
                        v21 = v21
                        _ = v20
                        _ = v21
                    } else {
                        v22 := v2[0]
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                }
            }
            var v23 float64 = float64((1.5 / -2.25))
            v3 += v3
            _ = v23
        }
    default:
        var v24 int32 = 684
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if ((4294967295 == 474) || true) {
            var v26 string = ("a" + ("foo bar" + ""))
            _ = v26
        }
        v27 := GenS1{F0: 2, F1: false}
        for fi28 := 0; fi28 < 19; fi28++ {
            delete(v2, 12)
            var v29 bool = (v3 >= (^(v3)))
            v30 := []int{fi25, int(v3)}
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi25
        _ = v27
    }
    fmt.Println(v7, v7)
    if false {
        var v31 uint64 = 18446744073709551615
        for fi32 := 0; fi32 < 16; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 bool = (0 >= 18446744073709551615)
                var v35 uint = 4294967295
                v36 := GenS2{F0: v1, F1: v34, F2: v31}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            if true {
                v37 := GenS3{F0: 0, F1: 255}
                switch 619 {
                case -257:
                    fmt.Println(fi32, v4)
                case -663:
                    v38 := []int{32767, fi32}
                    _ = v38
                default:
                    v39 := GenS1{F0: 0, F1: false}
                    _ = v39
                }
                v31 = v31
                _ = v37
            } else {
                for fi40 := 0; fi40 < 3; fi40++ {
                    var v41 int16 = (v3 * (v3 | (v3 % ((v3) | 1))))
                    _ = fi40
                    _ = v41
                }
                var v42 int32 = 935
                _ = v42
            }
            v43 := GenS4{F0: 0, F1: 884, F2: 132}
            _ = fi32
            _ = v43
        }
        _ = v31
    }
    v44 := []int{9223372036854775807, -844, 0, 2147483647}
    switch "foo bar" {
    case "hello":
        var v45 int32 = -2147483648
        _ = v45
    case "a":
        fmt.Println(v7, v1)
    case "\n\t":
        v1 = v1
    }
    if len(v44) > 0 {
        v44[0] = -873
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v7
}
