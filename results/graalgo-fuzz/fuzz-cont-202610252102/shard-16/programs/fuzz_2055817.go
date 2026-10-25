package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64) uint32 {
    p0 = (p0 >> ((493) & 63))
    p0 = 42
    if (0 <= 980) {
        var v7 int16 = 0
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                var v9 uint16 = 0
                fmt.Println(fi8, v9, p0)
                _ = v9
            }
            _ = fi8
        }
        v10 := []int{306, -558}
        _ = v10
        _ = v7
    } else {
        v11 := GenS1{F0: 0}
        _ = v11
    }
    _ = p0
    return 42
}
func genFn1(p0 uint, p1 uint, p2 string) int16 {
    if false {
        p1 -= (p1 * (p1 % ((255) | 1)))
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        p0 = 2
        _ = fi12
    }
    var v13 int8 = -128
    var v14 uint32 = 684
    var v15 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return (^(v15))
}

func main() {
    var v1 int8 = 1
    var v2 string = "a"
    var v3 uint16 = 65535
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    v6 := []int{4294967295}
    if false {
        var v16 uint = 255
        var v17 int32 = 1
        if (1e10 > 0.0) {
            var v18 float32 = -2.25
            if (("hello" >= (v2 + "a")) && true) {
                v19 := v4[((v2 + v2) + "\n\t")]
                v2 = v2
                var v20 int = 475
                _ = v19
                _ = v20
            }
            if true {
                var v21 int8 = v1
                _ = v21
            }
            _ = v18
        }
        _ = v16
        _ = v17
    }
    var v22 int32 = 188
    v4["\n\t"] = 0
    fmt.Println(v22)
    v23 := GenS2{F0: v22, F1: 42}
    var v24 string = (v2 + "hello")
    if false {
        switch -860 {
        case -1:
            v1 = v1
        default:
            switch 727 {
            case 42:
                delete(v4, "")
            case -9223372036854775808:
                v25 := GenS1{F0: 128}
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 bool = false
            fmt.Println(v22)
            v28 := GenS1{F0: 151}
            _ = fi26
            _ = v27
            _ = v28
        }
        v2 = v24
    } else {
        v29 := GenS2{F0: (v22 | v22), F1: 451}
        v30 := GenS3{F0: false, F1: 4294967295}
        if (2147483647 == 131) {
            v31 := v4["hello"]
            v2 = ("" + "a")
            fmt.Println(v1, v24)
            _ = v31
        } else {
            v32 := GenS1{F0: 255}
            _ = v32
        }
        _ = v29
        _ = v30
    }
    switch v24 {
    case "":
        v33 := GenS2{F0: (v22 - v22), F1: 65535}
        _ = v33
    case "a":
        v2 = (v2 + v2)
    case "hello":
        v3 = 255
    }
    v3 = v3
    var v34 bool = false
    var v35 uint8 = 221
    var v36 int = -478
    for fi37 := 0; fi37 < 20; fi37++ {
        delete(v4, ("foo bar" + "a"))
        _ = fi37
    }
    switch "hello" {
    case "a":
        var v38 float32 = -2.25
        _ = v38
    case "foo bar":
        var v39 uint64 = 79
        _ = v39
    default:
        for fi40 := 0; fi40 < 13; fi40++ {
            v35 = 120
            _ = fi40
        }
    }
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
