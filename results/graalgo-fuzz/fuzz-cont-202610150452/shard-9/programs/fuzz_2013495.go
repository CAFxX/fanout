package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int32; F1 bool }
func genFn0() int16 {
    var v6 int32 = 1
    var v7 int32 = (v6 / ((v6) | 1))
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 uint32, p1 uint, p2 uint) int16 {
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(p1)
        _ = fi8
    }
    v9 := GenS1{F0: -2147483648}
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := []int{(fi10 - (fi10 << ((2) & 63))), fi10, 128}
        fmt.Println(p2)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 int32 = 42
    v3 := make(map[int]uint32)
    var v4 int32 = v2
    v5 := []string{"foo bar", "hello", "foo bar"}
    if (((((-2.25 / (0.0 - (1e10 + 0.1))) + (0.1 * 1e10)) + float32((float64(float32((3.14159 + (float64(0.0) + 0.0)))) + (-2.25 / (0.1 - (1.5 + 3.14159)))))) >= float32((((1e10 - 3.14159) / 3.14159) * 0.1))) || (!false)) {
        v4 = 879
        var v12 string = ""
        var v13 int16 = 400
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            v2 -= v4
            v15 := GenS2{F0: 127, F1: (("" + (("foo bar" + "foo bar") + ("hello" + (("\n\t" + "\n\t") + "\n\t")))) + ("hello" + "\n\t"))}
            _ = fi14
            _ = v15
        }
    }
    if false {
        fmt.Println(v2, v2)
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint32 = 1
            v2 -= v2
            var v18 uint = 826
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        switch int(v4) {
        case 42:
            var v19 uint16 = 32767
            _ = v19
        default:
            var v20 uint64 = 2
            _ = v20
        }
    }
    var v21 string = ("" + ("hello" + "a"))
    var v22 float64 = 3.14159
    if false {
        v21 = "hello"
        v23 := GenS1{F0: -421}
        _ = v23
    }
    var v24 int = 9223372036854775807
    v21 = v21
    var v25 int = v24
    var v26 int32 = v4
    fmt.Println(v2, v2)
    switch v25 {
    case 9223372036854775807:
        if false {
            switch (v25 | v25) {
            case -2147483648:
                v1 = append(v1, (-(v24)))
            case 2147483647:
                for fi27 := 0; fi27 < 9; fi27++ {
                    delete(v3, v24)
                    _ = fi27
                }
            default:
                v28 := GenS1{F0: v24}
                _ = v28
            }
        }
    case 2147483647:
        var v29 int = 980
        _ = v29
    }
    switch v21 {
    case "":
        v25 = 32768
    case "hello":
        if true {
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS1{F0: (v25 / (((fi30 * v24)) | 1))}
                var v32 int = (v25 | 127)
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    default:
        v33 := GenS3{F0: 0, F1: false}
        _ = v33
    }
    fmt.Println(v26)
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            if (9223372036854775807 > -9223372036854775808) {
                switch 1 {
                case 127:
                    var v36 uint = 1
                    _ = v36
                case 128:
                    if true {
                        var v37 int16 = 0
                        _ = v37
                    }
                }
            }
            _ = fi35
        }
        _ = fi34
    }
    v1 = append(v1, (v25 - v25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
