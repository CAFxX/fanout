package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 int16, p1 bool) uint8 {
    if (2 <= 42) {
        var v5 uint = 255
        fmt.Println(p1, p1, v5)
        _ = v5
    }
    v6 := GenS1{F0: -1, F1: 0}
    if (2 == 2) {
        switch -908 {
        case 4294967295:
            v7 := GenS1{F0: -9223372036854775808, F1: 42}
            _ = v7
        default:
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int16 = 50
                    switch "a" {
                    case "a":
                        v11 := GenS2{F0: p1, F1: "a"}
                        _ = v11
                    case "foo bar":
                        var v12 int8 = -124
                        _ = v12
                    default:
                        fmt.Println(fi8, v10, fi9)
                    }
                    _ = fi9
                    _ = v10
                }
                p0 -= (^(p0))
                _ = fi8
            }
        }
    }
    if (9223372036854775807 == 9223372036854775807) {
        switch 127 {
        case 1:
            p1 = p1
        case 32767:
            if (true || (-2147483648 == 1)) {
                switch 0 {
                case -639:
                    v13 := []int{128, 1}
                    _ = v13
                case 1:
                    p0 -= 1
                default:
                    p1 = p1
                }
            }
        }
    } else {
        var v14 uint8 = 255
        var v15 int64 = 458
        var v16 int16 = (p0 >> ((230) & 63))
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := []int{691, -791, -353, 716}
    _ = p0
    _ = p1
    _ = v17
    _ = v6
    return 127
}

func main() {
    v1 := []string{("a" + "a")}
    v2 := make(map[string]int)
    var v3 bool = ((true && (true || (!(-108 < 96)))) && ((0 <= -58) || (!false)))
    v4 := []string{""}
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int16 = 0
        _ = len(v4)
        var v20 int = (fi18 & (fi18 % (((fi18 * 801)) | 1)))
        _ = fi18
        _ = v19
        _ = v20
    }
    if (42 > 42) {
        fmt.Println(v3)
        _ = len(v1)
        v1 = append(v1, "")
    }
    if v3 {
        v21 := GenS1{F0: -1, F1: 77}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int64 = 1
            _ = fi22
            _ = v23
        }
        if ((1e10 + 0.0) == 1e10) {
            var v24 uint8 = 66
            if (v3 && v3) {
                v24 = v24
                fmt.Println(v3, v3)
                fmt.Println(v3, v24)
            } else {
                var v25 uint32 = 1
                _ = v25
            }
            var v26 uint32 = 0
            _ = v24
            _ = v26
        } else {
            var v27 string = "foo bar"
            var v28 int32 = 2147483647
            _ = v27
            _ = v28
        }
    }
    if v3 {
        v29 := GenS3{F0: 1}
        var v30 uint64 = 4294967295
        v2["foo bar"] = 128
        _ = v29
        _ = v30
    } else {
        var v31 uint64 = 4294967295
        var v32 uint = 1
        var v33 uint64 = v31
        _ = v31
        _ = v32
        _ = v33
    }
    v3 = (!(v3 && (!(false || (((807 > 128) && (127 > 172)) || v3)))))
    var v34 int32 = 42
    v34 += (-(v34))
    v35 := GenS1{F0: 788, F1: 326}
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 int32 = -1
        v38 := GenS2{F0: (true || (100 < -128)), F1: (("hello" + "\n\t") + "")}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v34, v34)
    v3 = ((true || (1 >= -1)) && false)
    v34 = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v39 := GenS4{F0: ((v3 && (!(50 < 66))) && true)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
