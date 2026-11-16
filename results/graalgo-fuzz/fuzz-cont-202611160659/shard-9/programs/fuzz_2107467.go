package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 int32 }

func main() {
    v1 := []uint32{}
    var v2 float32 = (((float32(((1e10 * 0.0) + 0.1)) - float32(float64(-2.25))) / (float32((0.1 + 1e10)) * (1.5 * -2.25))) * (1e10 / ((0.1 * -2.25) - (0.1 * (float32(1e10) / (float32(-2.25) / 0.1))))))
    v3 := []uint32{4294967295}
    v4 := []string{"a", "a", (("foo bar" + "foo bar") + "\n\t")}
    switch -83 {
    case 32768:
        var v5 uint = 22
        _ = v5
    case 42:
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: (fi6 % (((fi7 ^ (fi6 & -997))) | 1)), F1: -1}
                v2 = 0.0
                if false {
                    var v9 int = 4294967295
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    v10 := GenS1{F0: (v9 / ((fi6) | 1)), F1: -316}
                    _ = v10
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            var v11 bool = ((950 == 1000) || ((3.14159 * (-2.25 + -2.25)) > 1e10))
            _ = fi6
            _ = v11
        }
    case 4294967295:
        var v12 uint = 4294967295
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v2, fi13)
        if false {
            var v14 float32 = v2
            switch fi13 {
            case 458:
                v2 = float32(-2.25)
            case 32768:
                var v15 int8 = 0
                _ = v15
            case 1:
                fmt.Println(v14, fi13, v14)
            default:
                fmt.Println(fi13, fi13)
            }
            if (!true) {
                var v16 bool = ((42 >= -2147483648) || false)
                _ = v16
            }
            _ = v14
        } else {
            for fi17 := 0; fi17 < 11; fi17++ {
                v2 = 1.5
                _ = fi17
            }
        }
        fmt.Println(fi13, v2)
        _ = fi13
    }
    if (1 == 0) {
        switch ("" + "") {
        case "a":
            fmt.Println(v2)
        case "":
            var v18 uint32 = 100
            _ = v18
        case "foo bar":
            v2 -= (v2 - (v2 + v2))
        default:
            v2 = (v2 * 1e10)
        }
    }
    v2 = 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = (1e10 * v2)
    if false {
        fmt.Println(v2)
        v19 := GenS2{F0: ("a" + ("" + "a")), F1: 100}
        _ = v19
    }
    v2 = float32(0.1)
    var v20 uint = 255
    v21 := GenS3{F0: 18446744073709551615, F1: 0}
    v22 := GenS4{F0: (766 > 4294967295), F1: -119}
    var v23 string = "\n\t"
    v24 := GenS5{F0: 16}
    if (-1 == -1) {
        switch 402 {
        case 1:
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v23 = ("foo bar" + "hello")
                var v25 string = (v23 + ((((("foo bar" + "foo bar") + "foo bar") + ("a" + v23)) + (v23 + ((v23 + "hello") + (v23 + "")))) + v23))
                _ = v25
            }
        }
    }
    switch "a" {
    case "":
        v26 := GenS2{F0: "\n\t", F1: -32}
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 18; fi27++ {
            if (v20 != (^(v20))) {
                fmt.Println(v23, fi27)
            } else {
                if false {
                    var v28 uint16 = 602
                    fmt.Println(fi27, v28, v20)
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v30 := GenS6{F0: 8, F1: -1}
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            }
            _ = fi27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
