package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64) uint32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 42}
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    var v9 string = ""
    _ = p0
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 uint, p2 uint8) uint64 {
    if (100 < 100) {
        v10 := GenS2{F0: 0, F1: 42}
        _ = v10
    } else {
        switch 32768 {
        case 1:
            p0 += p0
        case -1:
            var v11 uint8 = (^(p2))
            _ = v11
        default:
            var v12 float32 = -2.25
            _ = v12
        }
        v13 := []int{957}
        fmt.Println(p2, p1, p0)
        _ = v13
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("hello" + "a"), ("a" + "\n\t"), "foo bar"}
    var v3 float64 = 0.0
    var v4 int16 = 1
    var v5 uint64 = 739
    if (1 < 255) {
        var v15 uint = 2
        _ = v15
    } else {
        v16 := GenS2{F0: genFn0(genFn1(872, 42, 0)), F1: 1}
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "\n\t" {
        case "\n\t":
            switch "hello" {
            case "":
                v5 = (v5 % (((v5 & (v5 << ((4294967295) & 63)))) | 1))
            case "a":
                for fi18 := 0; fi18 < 7; fi18++ {
                    v3 = (v3 - ((float64(1e10) - float64((-2.25 - -2.25))) / 0.1))
                    _ = fi18
                }
            case "\n\t":
                v19 := GenS2{F0: 0, F1: v5}
                _ = v19
            default:
                switch "" {
                case "\n\t":
                    if true {
                        v20 := GenS2{F0: 0, F1: (v5 ^ (v5 << ((42) & 63)))}
                        v21 := GenS1{F0: -1}
                        var v22 float64 = v3
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                case "":
                    var v23 uint32 = 1
                    _ = v23
                case "foo bar":
                    fmt.Println(fi17)
                default:
                    var v24 bool = ((!true) && ((^(v5)) != 42))
                    _ = v24
                }
            }
        case "a":
            v25 := v1[(fi17 + fi17)]
            _ = v25
        }
        _ = fi17
    }
    fmt.Println(v5, v4, v3)
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS3{F0: 0}
        v5 += genFn1(514, 717, 255)
        var v28 int = fi26
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint32 = 2
    var v30 uint8 = 0
    if len(v2) > 0 {
        v2[0] = ("" + (("a" + ((("hello" + "\n\t") + "a") + "foo bar")) + ("foo bar" + "\n\t")))
    }
    if (9223372036854775807 != 9223372036854775807) {
        v2 = append(v2, "")
    }
    fmt.Println(v29, v5, v29)
    var v31 uint32 = (v29 + (^(v29)))
    v29 = v31
    var v32 uint = 255
    fmt.Println(v32)
    var v33 uint64 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
