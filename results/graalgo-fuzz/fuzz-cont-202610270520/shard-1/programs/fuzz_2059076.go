package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint) uint8 {
    p0 = p0
    v4 := []int{772, 1}
    if (((!true) || true) && false) {
        v4 = append(v4, 441)
    } else {
        p0 += 255
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((!true) && false) {
            v4 = append(v4, -9223372036854775808)
            fmt.Println(p0)
        } else {
            p0 -= 75
            var v5 bool = true
            v5 = true
            _ = v5
        }
    }
    _ = p0
    _ = v4
    return 127
}
func genFn1(p0 uint8) bool {
    switch "foo bar" {
    case "hello":
        var v6 float64 = 0.0
        _ = v6
    }
    var v7 uint64 = 1
    v7 = v7
    _ = p0
    _ = v7
    return true
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = (-32768 > 273)
    var v3 string = "foo bar"
    v3 = v3
    var v8 string = v3
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = 76
        fmt.Println(fi9, v1, fi9)
        if (!v2) {
            fmt.Println(v3, v2, v1)
            switch 4294967295 {
            case 1:
                fmt.Println(v10, fi9)
            case 42:
                if (false || v2) {
                    v1 += (v1 / (((v1 << ((2) & 63))) | 1))
                } else {
                    v11 := []int{(fi9 & 2147483647)}
                    _ = v11
                }
            default:
                v12 := GenS1{F0: 65535}
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    var v13 float32 = 0.0
    var v14 uint16 = 0
    v14 -= 255
    var v15 uint64 = 2
    for fi16 := 0; fi16 < 10; fi16++ {
        if ((!(!(1 != 42))) || (!true)) {
            switch 32768 {
            case 128:
                v17 := []int{(fi16 + fi16), fi16, (fi16 & (^(fi16))), (fi16 + fi16)}
                _ = v17
            case 811:
                var v18 string = ("" + "a")
                _ = v18
            case 127:
                v1 = 2
            default:
                if (!("hello" == v3)) {
                    v19 := GenS2{F0: 1}
                    switch (fi16 & (fi16 * fi16)) {
                    case -79:
                        v3 = v8
                    case -621:
                        fmt.Println(fi16)
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 string = ("a" + "a")
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                if true {
                    var v23 int = -9223372036854775808
                    _ = v23
                } else {
                    if v2 {
                        v24 := []int{fi22, -1, 32768}
                        _ = v24
                    }
                    var v25 float64 = -2.25
                    v15 = (v15 << ((324) & 63))
                    _ = v25
                }
                _ = fi22
            }
        }
        _ = fi16
    }
    v26 := GenS2{F0: 42}
    var v27 int = 42
    var v28 float64 = (0.0 / 1.5)
    v29 := []int{v27}
    v30 := GenS1{F0: 453}
    if (47 > -1) {
        if (42 != 291) {
            var v31 string = v8
            _ = len(v29)
            v29 = append(v29, (-(v27)))
            _ = v31
        }
        var v32 uint64 = (^(v15))
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v8
}
