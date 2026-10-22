package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8, p1 string) int16 {
    var v4 string = p1
    if true {
        if (1 == 267) {
            fmt.Println(p1, p0)
        } else {
            v5 := GenS1{F0: p0, F1: 803}
            v6 := []int{179}
            p0 -= (p0 - 255)
            _ = v5
            _ = v6
        }
    }
    p0 += (p0 | p0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 393
}
func genFn1(p0 bool, p1 string, p2 int16) int16 {
    fmt.Println(p0)
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{(fi9 ^ -9223372036854775808), 1}
            _ = fi9
            _ = v10
        }
        v11 := []int{fi8, (^(fi8)), (fi8 / ((9223372036854775807) | 1))}
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 * p2)
}

func main() {
    var v1 int8 = -1
    var v2 bool = true
    var v3 int32 = -636
    v2 = (("foo bar" + "\n\t") != "foo bar")
    switch "foo bar" {
    case "a":
        v12 := []int{42, 32768}
        _ = v12
    default:
        fmt.Println(v1, v1)
    }
    v2 = (1.5 == 1.5)
    var v13 string = "a"
    if false {
        v3 += -2147483648
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint = 255
            v16 := []int{fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := []int{}
        _ = v17
    } else {
        v2 = (32767 != 1000)
    }
    v3 += (v3 | (v3 | (v3 - (v3 * v3))))
    var v18 uint16 = 0
    v13 = "\n\t"
    v1 = v1
    if true {
        switch "" {
        case "foo bar":
            v18 += v18
        default:
            var v19 uint32 = 1
            _ = v19
        }
        var v20 float64 = 0.1
        _ = v20
    } else {
        switch int(v1) {
        case 1:
            if (128 < 1) {
                v18 = 65535
                var v21 int64 = 456
                _ = v21
            } else {
                for fi22 := 0; fi22 < 20; fi22++ {
                    v13 = v13
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v13 = ("\n\t" + (v13 + "a"))
                        var v24 float32 = (1e10 - 1e10)
                        v25 := GenS1{F0: 255, F1: fi23}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
            }
        default:
            var v26 int16 = genFn1(v2, "", 1)
            _ = v26
        }
        v1 = -1
        v27 := GenS1{F0: 128, F1: 0}
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS1{F0: 128, F1: -823}
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 float32 = 0.1
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    var v32 bool = false
    v33 := GenS2{F0: v32, F1: 0}
    v34 := GenS2{F0: ((v18 / (((v18 << ((42) & 63))) | 1)) == 42), F1: 23}
    v35 := GenS2{F0: v2, F1: 226}
    if (!(169 != 1)) {
        var v36 float64 = 3.14159
        switch 127 {
        case 32767:
            v37 := []int{9223372036854775807, int(v1), 706, 0}
            _ = v37
        case 1:
            v38 := []int{334, 9223372036854775807, 0}
            _ = v38
        case -796:
            if (0 < 2147483648) {
                switch (v13 + "foo bar") {
                case "":
                    var v39 float32 = (-2.25 * 0.0)
                    _ = v39
                case "hello":
                    switch v13 {
                    case "":
                        fmt.Println(v3, v2)
                    case "hello":
                        v40 := []int{-1}
                        _ = v40
                    default:
                        v41 := []int{}
                        _ = v41
                    }
                default:
                    var v42 float64 = -2.25
                    _ = v42
                }
                v18 = (v18 + v18)
            } else {
                v13 = v13
                if (!(0 <= 0)) {
                    v13 = ((v13 + (("a" + v13) + (v13 + v13))) + ("foo bar" + ("" + v13)))
                }
            }
        default:
            fmt.Println(v1)
        }
        v43 := GenS2{F0: (!(!(3.14159 > v36))), F1: 0}
        _ = v36
        _ = v43
    } else {
        fmt.Println(v3, v2)
        v44 := []int{}
        v13 = v13
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
