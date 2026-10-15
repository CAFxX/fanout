package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 int) uint16 {
    var v4 int16 = -996
    v5 := GenS1{F0: 0, F1: 29, F2: p0}
    v4 = (v4 | v4)
    _ = p0
    _ = v4
    _ = v5
    return 674
}
func genFn1() int64 {
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v6, v6)
        v8 := []int{(-(fi7)), (-(fi7))}
        _ = fi7
        _ = v8
    }
    _ = v6
    return -724
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    if ((!false) || (42 >= -564)) {
        var v9 int16 = 1000
        _ = v9
    } else {
        v10 := GenS1{F0: -128, F1: 127, F2: 9223372036854775807}
        if (1 < 524) {
            var v11 uint = 18446744073709551615
            v12 := v2[42]
            _ = v11
            _ = v12
        } else {
            var v13 string = ("" + "a")
            _ = v13
        }
        var v14 uint = 4294967295
        _ = v10
        _ = v14
    }
    switch 42 {
    case 184:
        var v15 int = 4294967295
        _ = v15
    case -699:
        var v16 int16 = 1000
        _ = v16
    default:
        var v17 uint8 = 1
        _ = v17
    }
    var v18 bool = true
    v18 = ((-619 > 163) && false)
    if (("a" + "foo bar") <= "a") {
        v19 := GenS2{F0: 59}
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v18, v18)
        _ = fi20
    }
    switch "foo bar" {
    case "\n\t":
        for fi21 := 0; fi21 < 18; fi21++ {
            fmt.Println(fi21, v18)
            if v18 {
                var v22 uint8 = 2
                _ = v22
            } else {
                switch "foo bar" {
                case "a":
                    var v23 int64 = -538
                    _ = v23
                }
            }
            v24 := []int{fi21, -1}
            _ = fi21
            _ = v24
        }
    default:
        if true {
            for fi25 := 0; fi25 < 1; fi25++ {
                switch ("a" + "foo bar") {
                case "a":
                    v26 := GenS1{F0: 100, F1: 1, F2: (fi25 - fi25)}
                    _ = v26
                case "":
                    v18 = true
                default:
                    v18 = (0.0 < 1e10)
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint32 = 65535
                    var v29 uint = 813
                    var v30 uint64 = 0
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                fmt.Println(fi25)
                _ = fi25
            }
        } else {
            v31 := []int{4294967295, 32767}
            _ = v31
        }
    }
    v18 = false
    _ = len(v2)
    v32 := GenS3{F0: 42, F1: ((v18 && v18) && ((4294967295 == 2) || (933 >= 18446744073709551615)))}
    v18 = v18
    var v33 float64 = float64(float32((float64((1.5 / 1.0)) - 0.0)))
    var v34 int8 = -1
    v34 = -1
    fmt.Println(v34, v33)
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
