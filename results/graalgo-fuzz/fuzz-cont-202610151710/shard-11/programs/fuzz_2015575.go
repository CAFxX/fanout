package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0() int32 {
    v4 := GenS1{F0: 149, F1: 100, F2: 4294967295}
    if (-2.25 >= 0.1) {
        var v5 int32 = 533
        _ = v5
    } else {
        var v6 string = ("" + "")
        v7 := []int{765, -286}
        _ = v6
        _ = v7
    }
    _ = v4
    return 42
}
func genFn1() uint32 {
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8)
        var v9 int64 = 9223372036854775807
        var v10 float64 = 1e10
        _ = fi8
        _ = v10
        _ = v9
    }
    if ((0.0 / (-2.25 + -2.25)) == float32(-2.25)) {
        switch -809 {
        case -1:
            v11 := []int{}
            _ = v11
        case 443:
            var v12 uint32 = 726
            _ = v12
        default:
            v13 := []int{-2147483648, -371, 32768, -235}
            _ = v13
        }
        v14 := []int{4294967295, 127, -1}
        switch "hello" {
        case "foo bar":
            v15 := GenS2{F0: 874}
            _ = v15
        case "\n\t":
            switch 4294967295 {
            case -854:
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 uint16 = 0
                    var v18 int = -2147483648
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                var v19 int = 42
                _ = v19
            }
        default:
            if true {
                switch "\n\t" {
                case "foo bar":
                    _ = len(v14)
                case "hello":
                    if (!false) {
                        var v20 int64 = -1
                        _ = v20
                    } else {
                        var v21 uint32 = 1
                        fmt.Println(v21)
                        _ = v21
                    }
                }
                var v22 int64 = 42
                _ = v22
            }
        }
        _ = v14
    } else {
        var v23 uint64 = 0
        switch (("" + (((("foo bar" + "foo bar") + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))) + "") {
        case "\n\t":
            v23 += v23
        default:
            var v24 uint32 = 832
            _ = v24
        }
        _ = v23
    }
    return 2
}

func main() {
    var v1 float64 = (((3.14159 * ((-2.25 + 3.14159) + 1e10)) + 3.14159) + float64((1e10 * 0.1)))
    var v2 int32 = -827
    v3 := make(map[uint32]string)
    var v25 int32 = (v2 >> ((0) & 63))
    var v26 float64 = 1.5
    var v27 bool = false
    for fi28 := 0; fi28 < 11; fi28++ {
        delete(v3, genFn1())
        delete(v3, genFn1())
        _ = fi28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(v25)
        _ = fi29
    }
    var v30 int = 1
    switch "a" {
    case "hello":
        if (v27 && v27) {
            v31 := v3[genFn1()]
            var v32 uint32 = 2147483648
            var v33 float64 = 0.0
            _ = v31
            _ = v32
            _ = v33
        }
    default:
        var v34 int64 = -25
        _ = v34
    }
    fmt.Println(v1, v30)
    var v35 uint = 42
    v25 += v2
    v1 = v1
    for fi36 := 0; fi36 < 18; fi36++ {
        fmt.Println(v2)
        var v37 int = v30
        switch v37 {
        case 725:
            v30 -= 127
        default:
            fmt.Println(v25)
        }
        _ = fi36
        _ = v37
    }
    v38 := []int{128, (-(v30)), -144}
    v39 := GenS3{F0: -27}
    v25 = v25
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v35
    _ = v38
    _ = v39
}
