package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 bool, p1 int64, p2 int8) string {
    v4 := GenS1{F0: "", F1: 0}
    p1 = 9223372036854775807
    switch 0 {
    case -70:
        fmt.Println(p2)
    case -574:
        var v5 uint32 = 720
        _ = v5
    }
    v6 := GenS1{F0: (("foo bar" + ("" + "\n\t")) + "foo bar"), F1: 42}
    switch "hello" {
    case "a":
        v7 := []int{-9223372036854775808}
        _ = v7
    case "foo bar":
        if (!p0) {
            p1 = -1
        } else {
            p2 -= 1
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(p0, fi9)
                    var v10 uint8 = 255
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            switch ("foo bar" + "a") {
            case "hello":
                var v11 int32 = 42
                _ = v11
            case "\n\t":
                var v12 bool = (127 == -536)
                _ = v12
            case "":
                switch "a" {
                case "a":
                    for fi13 := 0; fi13 < 10; fi13++ {
                        var v14 int8 = (p2 / ((-1) | 1))
                        v15 := []int{fi13, (fi13 + fi13)}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case "\n\t":
                    v16 := []int{}
                    _ = v16
                case "hello":
                    var v17 int = 127
                    _ = v17
                }
            }
        }
    case "":
        v18 := GenS1{F0: "foo bar", F1: 673}
        _ = v18
    default:
        for fi19 := 0; fi19 < 13; fi19++ {
            if p0 {
                p2 -= -1
                fmt.Println(fi19)
            }
            var v20 int8 = (-(p2))
            _ = fi19
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return "a"
}
func genFn1() uint64 {
    var v21 float64 = (-2.25 * 1.5)
    var v22 uint8 = 127
    var v23 uint = 0
    fmt.Println(v23, v22)
    _ = v21
    _ = v22
    _ = v23
    return uint64(v23)
}

func main() {
    var v1 int16 = 0
    var v2 uint = 782
    v3 := []uint32{}
    var v24 uint64 = 0
    fmt.Println(v24, v2)
    var v25 uint8 = 18
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS1{F0: ("" + "a"), F1: v2}
        _ = fi26
        _ = v27
    }
    v24 = (v24 % (((^(v24))) | 1))
    switch "foo bar" {
    case "a":
        v28 := GenS1{F0: "a", F1: (v2 % ((v2) | 1))}
        _ = v28
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 float32 = 3.14159
            _ = fi29
            _ = v30
        }
    }
    v24 = (^(v24))
    var v31 float64 = (0.0 / float64(((1.5 * (1.5 + (0.0 / 1.0))) / (((-2.25 - 1e10) - 1.5) + float32((1.5 * (3.14159 + float64((((3.14159 / -2.25) - 0.0) / -2.25)))))))))
    var v32 float32 = 1.5
    switch "" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 2
        }
    default:
        var v33 int64 = 9223372036854775807
        _ = v33
    }
    var v34 uint16 = 714
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v34
}
