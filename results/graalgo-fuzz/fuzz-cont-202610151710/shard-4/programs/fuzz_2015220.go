package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int8; F1 int; F2 uint }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int64, p1 bool, p2 bool) uint8 {
    v5 := GenS1{F0: 128}
    var v6 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 int = 32768
    v4 := []string{}
    switch ("\n\t" + (("hello" + "a") + (("" + "") + ""))) {
    case "a":
        var v7 int16 = -996
        _ = v7
    case "\n\t":
        var v8 int32 = 1
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 float64 = 1.5
            v11 := v2[2]
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if (495 > 1000) {
            v2[863] = 2147483648
        }
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v3 -= 32768
        v3 = v3
        if ((true || ((-1 >= 32767) || true)) || (!false)) {
            v3 += fi12
            v13 := GenS1{F0: 42}
            if false {
                v3 = (v3 / ((v3) | 1))
                v14 := GenS1{F0: 62}
                switch fi12 {
                case 60:
                    v15 := GenS1{F0: 42}
                    _ = v15
                case -40:
                    v3 = 128
                default:
                    v3 += -1
                }
                _ = v14
            } else {
                if true {
                    var v16 uint32 = 0
                    var v17 bool = false
                    _ = v16
                    _ = v17
                }
            }
            _ = v13
        }
        _ = fi12
    }
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS1{F0: 255}
            _ = fi18
            _ = v19
        }
    } else {
        v20 := GenS2{F0: "", F1: (0 > -32768)}
        var v21 bool = (255 == 115)
        _ = v20
        _ = v21
    }
    v22 := v2[75]
    fmt.Println(v22)
    v3 += (v3 | -251)
    if false {
        fmt.Println(v22, v22)
        switch -9223372036854775808 {
        case -1:
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 float64 = 3.14159
                fmt.Println(v22, v22, v24)
                fmt.Println(fi23)
                _ = fi23
                _ = v24
            }
        }
        v25 := GenS3{F0: 42, F1: (v3 - (v3 ^ (v3 / ((2147483647) | 1)))), F2: 255}
        _ = v25
    } else {
        switch "\n\t" {
        case "foo bar":
            v3 += v3
        case "a":
            v26 := GenS4{F0: "hello", F1: 65535}
            _ = v26
        case "":
            for fi27 := 0; fi27 < 16; fi27++ {
                switch fi27 {
                case 9223372036854775807:
                    var v28 int = v3
                    _ = v28
                case 32767:
                    var v29 bool = (!true)
                    _ = v29
                case -1:
                    var v30 float32 = (((1.5 + 0.0) + 1.5) + 0.0)
                    _ = v30
                default:
                    var v31 float64 = 0.0
                    _ = v31
                }
                fmt.Println(fi27)
                v3 += (v3 + v3)
                _ = fi27
            }
        }
        v2[65535] = (v22 * (v22 * v22))
    }
    fmt.Println(v22, v22)
    var v32 uint32 = v22
    v22 += 42
    switch v3 {
    case -9223372036854775808:
        var v33 uint16 = 816
        _ = v33
    case 2147483647:
        var v34 uint16 = 241
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
