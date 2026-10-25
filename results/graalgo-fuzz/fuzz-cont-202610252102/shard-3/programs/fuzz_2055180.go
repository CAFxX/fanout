package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 bool, p2 uint) uint16 {
    var v6 string = "hello"
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 string) uint16 {
    v8 := GenS1{F0: -628, F1: -1, F2: 18446744073709551615}
    if true {
        p0 = p0
        var v9 int32 = -1
        _ = v9
    } else {
        p0 = p0
    }
    var v10 uint64 = 692
    _ = p0
    _ = v10
    _ = v8
    return genFn0(859, true, 744)
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{}
    var v3 uint8 = 67
    v4 := []string{"hello", "hello", "\n\t"}
    v5 := make(map[string]bool)
    v1 += -2.25
    for fi11 := 0; fi11 < 14; fi11++ {
        v4 = append(v4, "\n\t")
        var v12 string = "hello"
        v1 += ((v1 / v1) * 1.5)
        _ = fi11
        _ = v12
    }
    switch 9223372036854775807 {
    case 32768:
        v1 -= v1
    case 9223372036854775807:
        v1 -= -2.25
    default:
        v1 = (((v1 * (float64((float32(0.1) - float32(v1))) + v1)) - 1.5) / v1)
    }
    var v13 int16 = 32767
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = 255
        v16 := GenS2{F0: ("\n\t" + "hello")}
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = v13
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint16 = 217
        switch "\n\t" {
        case "":
            v1 -= float64(0.0)
        case "a":
            if false {
                var v19 int16 = (v13 % ((int16(v3)) | 1))
                _ = v19
            }
        }
        _ = fi17
        _ = v18
    }
    v20 := GenS1{F0: v13, F1: v13, F2: 0}
    v21 := GenS3{F0: -9223372036854775808}
    fmt.Println(v1, v3)
    v13 = v13
    v13 = v13
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v22 := GenS2{F0: "\n\t"}
    v1 = float64(-2.25)
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint32 = 4294967295
        var v25 uint16 = 676
        _ = fi23
        _ = v24
        _ = v25
    }
    delete(v5, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
