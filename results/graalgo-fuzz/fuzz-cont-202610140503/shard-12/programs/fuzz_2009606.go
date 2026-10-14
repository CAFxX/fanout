package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 int8, p1 string, p2 string) uint32 {
    if true {
        var v7 int64 = -9223372036854775808
        p0 = (-(p0))
        var v8 uint32 = 2147483647
        _ = v7
        _ = v8
    } else {
        if ("\n\t" <= "\n\t") {
            var v9 string = p1
            var v10 int = 4294967295
            v11 := GenS1{F0: "foo bar"}
            _ = v10
            _ = v11
            _ = v9
        } else {
            switch 42 {
            case -2147483648:
                fmt.Println(p1, p1, p2)
            case 1:
                var v12 uint16 = 143
                _ = v12
            }
            v13 := []int{-545, 9223372036854775807}
            for fi14 := 0; fi14 < 18; fi14++ {
                if (((-2.25 + 0.0) / (((-2.25 * -2.25) - float64(0.1)) - 0.0)) >= 0.1) {
                    fmt.Println(p1, p2, p1)
                    var v15 uint = 42
                    _ = v15
                }
                _ = fi14
            }
            _ = v13
        }
        if true {
            fmt.Println(p0, p0)
            p0 -= -128
            v16 := GenS1{F0: ("" + (("foo bar" + "a") + (p2 + ("" + p2))))}
            _ = v16
        } else {
            v17 := []int{573, 1, 4294967295, -830}
            v18 := GenS2{F0: 402}
            switch "hello" {
            case "hello":
                v17 = append(v17, 1)
            case "foo bar":
                fmt.Println(p2, p0)
            default:
                for fi19 := 0; fi19 < 3; fi19++ {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi19
                }
            }
            _ = v17
            _ = v18
        }
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if false {
            fmt.Println(p0, fi20, p0)
            v21 := GenS3{F0: 236}
            p2 = p1
            _ = v21
        } else {
            fmt.Println(p1, p2, fi20)
        }
        var v22 int = 127
        _ = fi20
        _ = v22
    }
    var v23 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return v23
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    var v4 int16 = 1000
    var v5 int16 = 936
    var v6 int64 = 9223372036854775807
    v24 := GenS3{F0: v3}
    var v25 float32 = 0.0
    v5 = v5
    for fi26 := 0; fi26 < 14; fi26++ {
        v25 = (float32((-2.25 - -2.25)) * (float32((1.5 / (((-2.25 - 0.0) / ((1.5 - (0.0 - 1.5)) + 1e10)) / -2.25))) + (v25 * v25)))
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS4{F0: -1, F1: fi26}
            if (!true) {
                if (!false) {
                    var v29 float64 = (0.0 * ((-2.25 + -2.25) / 1.5))
                    _ = v29
                } else {
                    v30 := GenS1{F0: ""}
                    var v31 int8 = 0
                    _ = v30
                    _ = v31
                }
                if false {
                    var v32 uint32 = genFn0(-1, (("foo bar" + ("" + "hello")) + "\n\t"), ("foo bar" + "\n\t"))
                    _ = v32
                } else {
                    fmt.Println(v4)
                    var v33 uint8 = 64
                    _ = v33
                }
            }
            v1[2] = 810
            _ = fi27
            _ = v28
        }
        var v34 float64 = 1e10
        _ = fi26
        _ = v34
    }
    v35 := GenS4{F0: 42, F1: -2147483648}
    switch 514 {
    case -1:
        delete(v1, 0)
    case 4294967295:
        fmt.Println(v4, v6)
    default:
        var v36 string = "foo bar"
        _ = v36
    }
    v37 := GenS4{F0: -43, F1: 0}
    for fi38 := 0; fi38 < 3; fi38++ {
        v3 = 2147483647
        switch "\n\t" {
        case "foo bar":
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 uint64 = 4294967295
                _ = fi39
                _ = v40
            }
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
