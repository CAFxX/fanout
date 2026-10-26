package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8 }
func genFn0() string {
    if true {
        var v6 uint16 = 32767
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float32 = 0.1
        v9 := GenS1{F0: (fi7 >> ((805) & 63)), F1: (1 < 1)}
        switch "" {
        case "\n\t":
            var v10 int64 = 0
            _ = v10
        case "a":
            var v11 uint8 = 198
            _ = v11
        case "":
            v12 := GenS1{F0: (fi7 % ((2147483647) | 1)), F1: false}
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v13 uint64 = 149
    switch ("foo bar" + "\n\t") {
    case "":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 string = "hello"
            v16 := []int{(fi14 + fi14)}
            _ = fi14
            _ = v15
            _ = v16
        }
    case "hello":
        var v17 int16 = -32768
        _ = v17
    case "foo bar":
        v18 := []int{-9223372036854775808, -9223372036854775808, 0, 128}
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := []int{(fi19 % ((4294967295) | 1))}
        for fi21 := 0; fi21 < 18; fi21++ {
            v20 = append(v20, fi21)
            fmt.Println(fi19, fi19)
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    _ = v13
    return ""
}
func genFn1(p0 int8) int32 {
    fmt.Println(p0)
    var v22 uint = 453
    p0 = p0
    var v23 int16 = -962
    v24 := []int{-2147483648}
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 float64 = (0.1 / 1e10)
    var v4 uint8 = 0
    var v5 int64 = -9223372036854775808
    var v25 string = "foo bar"
    switch -208 {
    case 1:
        v26 := GenS1{F0: 2147483647, F1: true}
        _ = v26
    default:
        v27 := []int{42, -136, 42, 646}
        _ = v27
    }
    v3 += v3
    if (254 <= 510) {
        v28 := GenS1{F0: -424, F1: false}
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := v2[("a" + (v25 + (v25 + v25)))]
            _ = fi30
            _ = v31
        }
        v32 := []int{fi29, fi29}
        _ = fi29
        _ = v32
    }
    v5 -= (-(v5))
    if false {
        v4 -= (v4 / ((v4) | 1))
        switch 32768 {
        case 0:
            v33 := GenS1{F0: 9223372036854775807, F1: true}
            _ = v33
        }
        v3 += (v3 - v3)
    }
    switch 32768 {
    case 9223372036854775807:
        v25 = (v25 + "\n\t")
    case 127:
        v34 := GenS2{F0: (^(v4))}
        _ = v34
    case 1:
        v5 = (v5 << ((574) & 63))
    default:
        if false {
            for fi35 := 0; fi35 < 13; fi35++ {
                delete(v2, (((v25 + (v25 + v25)) + v25) + ("hello" + v25)))
                _ = fi35
            }
            fmt.Println(v3)
        } else {
            fmt.Println(v25, v4)
            if (!true) {
                fmt.Println(v5, v5, v3)
                v36 := GenS2{F0: v4}
                _ = v36
            }
        }
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 uint64 = 61
        var v39 float32 = ((0.1 - 0.1) / -2.25)
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
