package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0(p0 bool, p1 int64, p2 int64) bool {
    fmt.Println(p2, p1, p0)
    v6 := []int{128, 127}
    switch "hello" {
    case "foo bar":
        v7 := GenS1{F0: 4294967295, F1: 513}
        _ = v7
    case "a":
        switch -166 {
        case 576:
            var v8 int16 = 270
            _ = v8
        case 128:
            p1 = (p2 % ((-9223372036854775808) | 1))
        }
    }
    v9 := GenS2{F0: 2}
    p1 = (p2 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return (264 > 973)
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    v4 := []string{"\n\t", ("foo bar" + "hello"), v1, v1}
    v5 := make(map[string]string)
    v1 = v1
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            v1 = ("a" + "")
            var v12 float64 = 3.14159
            if (!(false && (!true))) {
                v13 := GenS3{F0: (fi10 % ((fi10) | 1)), F1: 32767, F2: 940}
                v2[810] = -2147483648
                v4 = append(v4, "\n\t")
                _ = v13
            } else {
                if true {
                    v14 := v5["\n\t"]
                    v14 = v14
                    fmt.Println(fi10, v12, v14)
                    _ = v14
                } else {
                    var v15 int16 = 0
                    var v16 uint8 = 110
                    v4 = append(v4, "a")
                    _ = v15
                    _ = v16
                }
            }
            v3 = append(v3, 785)
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
        var v17 float64 = 1.5
        _ = fi10
        _ = v17
    }
    _ = len(v5)
    var v18 uint32 = 2
    var v19 uint64 = 42
    v20 := v5[v1]
    var v21 uint8 = 42
    v3 = append(v3, (^(v18)))
    switch "" {
    case "\n\t":
        switch -288 {
        case -756:
            v20 = "\n\t"
        case 4294967295:
            if (v20 == "\n\t") {
                if false {
                    var v22 int = 459
                    fmt.Println(v18, v1, v1)
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 int8 = 0
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                v18 = (v18 - (v18 * uint32(v21)))
            }
        case 724:
            for fi25 := 0; fi25 < 2; fi25++ {
                v20 = (v20 + "\n\t")
                v26 := GenS4{F0: (v19 ^ (^(v19))), F1: 1, F2: 1}
                _ = fi25
                _ = v26
            }
        }
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        var v27 uint8 = (v21 - v21)
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
