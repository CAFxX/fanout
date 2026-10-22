package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 int64, p1 uint64, p2 int) int8 {
    fmt.Println(p0, p1, p0)
    p0 += -1
    if (!true) {
        p0 = -214
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 = -9223372036854775808
            var v6 string = "hello"
            _ = fi5
            _ = v6
        }
        p2 = (p2 | (p2 + p2))
    } else {
        p1 = 283
        var v7 float64 = ((float64(-2.25) * 0.0) - (1.5 * 0.1))
        fmt.Println(p2, p2)
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        p1 -= (^(p1))
        p2 = (p2 & fi8)
        for fi9 := 0; fi9 < 20; fi9++ {
            switch "hello" {
            case "a":
                var v10 int16 = 1
                _ = v10
            case "foo bar":
                var v11 int64 = p0
                _ = v11
            }
            fmt.Println(fi8, p2, p0)
            p0 = (p0 - (p0 & p0))
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = -715
    var v2 float64 = 1.5
    var v3 int32 = 42
    v4 := make(map[string]string)
    switch "" {
    case "a":
        if true {
            v1 = v1
        }
    case "foo bar":
        if (65535 < 2) {
            v2 = 1.5
        } else {
            var v12 int64 = (v1 ^ (v1 | (v1 << ((65535) & 63))))
            v12 -= v12
            v12 = (-(v12))
            _ = v12
        }
    case "":
        var v13 uint32 = 42
        _ = v13
    default:
        switch 32767 {
        case 32767:
            switch (("\n\t" + "foo bar") + "hello") {
            case "foo bar":
                if (!(!("hello" <= "a"))) {
                    v3 -= 42
                    fmt.Println(v3)
                }
            case "a":
                var v14 bool = false
                _ = v14
            case "":
                switch -104 {
                case 127:
                    v3 = v3
                case 128:
                    if (65535 > 65535) {
                        v3 += 0
                    }
                case 451:
                    v15 := GenS1{F0: 4294967295}
                    _ = v15
                default:
                    if ((1.5 - (3.14159 - ((float64((1e10 + (1e10 - (1e10 / ((float32(float64(3.14159)) / float32(v2)) - 3.14159))))) - float64(1.5)) / v2))) <= ((0.1 + (1.5 + 1e10)) / (1.5 + (1.5 / v2)))) {
                        v2 = v2
                        v16 := GenS2{F0: v3, F1: true, F2: -2147483648}
                        var v17 bool = (v1 > v1)
                        _ = v16
                        _ = v17
                    }
                }
            default:
                v1 = (v1 >> ((2) & 63))
            }
        case -2147483648:
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 int16 = 42
                _ = fi18
                _ = v19
            }
        case 4294967295:
            v3 = (v3 - 2147483647)
        }
    }
    var v20 uint32 = 923
    v3 = v3
    v21 := GenS2{F0: v3, F1: ((v1 | -584) >= v1), F2: 0}
    for fi22 := 0; fi22 < 18; fi22++ {
        if false {
            v23 := GenS2{F0: 42, F1: (!false), F2: fi22}
            var v24 string = ""
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: 42}
        _ = fi22
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(fi26, fi26)
        var v27 float64 = (3.14159 + v2)
        var v28 float32 = 1e10
        _ = fi26
        _ = v27
        _ = v28
    }
    if true {
        var v29 int32 = 1
        if false {
            var v30 int64 = (^(v1))
            _ = v30
        }
        _ = v29
    } else {
        fmt.Println(v3, v3, v1)
    }
    v1 = v1
    v3 = v3
    var v31 int32 = v3
    for fi32 := 0; fi32 < 5; fi32++ {
        v33 := GenS3{F0: 0, F1: true}
        if (-785 <= v1) {
            delete(v4, "hello")
            v34 := GenS2{F0: v3, F1: false, F2: fi32}
            _ = v34
        } else {
            v35 := GenS1{F0: v20}
            if false {
                var v36 uint64 = 18446744073709551615
                v37 := GenS1{F0: v20}
                var v38 string = "\n\t"
                _ = v36
                _ = v37
                _ = v38
            }
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    var v39 float64 = 0.1
    v40 := GenS4{F0: uint(v20)}
    v41 := GenS5{F0: 2}
    v42 := GenS6{F0: -128, F1: 42}
    v43 := []int{0}
    v2 = v2
    var v44 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
