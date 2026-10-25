package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int, p1 bool) uint64 {
    var v6 int16 = -271
    var v7 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int64, p2 uint) uint64 {
    p1 = p1
    v8 := GenS1{F0: 104, F1: 42, F2: 509}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 bool = true
    v5 := []uint32{289, 822, 2, 2}
    if (127 != 892) {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint64 = 242
            _ = fi9
            _ = v10
        }
    }
    switch "hello" {
    case "\n\t":
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int = 2147483647
            _ = fi11
            _ = v12
        }
    case "a":
        var v13 float64 = (0.0 - 0.1)
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint = 4294967295
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v4 = v4
            _ = fi17
        }
        _ = fi16
    }
    v4 = ((((-2.25 / (-2.25 / 1.0)) * ((float64(float32((3.14159 - (1.5 - float64(float32(((3.14159 - (3.14159 / (3.14159 + 3.14159))) - (((3.14159 / (-2.25 + (0.0 / 0.1))) + 0.0) - 1.5)))))))) / float64(((1.5 - 0.1) - (1e10 * float32(0.1))))) * (1.5 - (3.14159 / (0.0 + 1e10))))) - 0.1) < 1e10)
    if v4 {
        switch ("\n\t" + ((("foo bar" + "a") + "a") + "a")) {
        case "":
            v18 := GenS2{F0: 42, F1: 1, F2: 0}
            _ = v18
        case "foo bar":
            v4 = false
        case "a":
            v4 = false
        }
    }
    switch "\n\t" {
    case "hello":
        for fi19 := 0; fi19 < 5; fi19++ {
            v4 = false
            v4 = (476 == genFn0(fi19, true))
            _ = fi19
        }
    case "\n\t":
        var v20 uint32 = 157
        _ = v20
    }
    fmt.Println(v4)
    fmt.Println(v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v21 int64 = 1
    switch 42 {
    case -567:
        v4 = ((-2.25 / 1.0) < 3.14159)
    case -653:
        for fi22 := 0; fi22 < 8; fi22++ {
            v21 = (v21 ^ v21)
            var v23 string = "\n\t"
            _ = fi22
            _ = v23
        }
    case 32768:
        var v24 int32 = 0
        _ = v24
    }
    v25 := GenS3{F0: 2, F1: (((float64(3.14159) * (0.1 - ((0.1 - 0.1) - (0.0 * float64(0.1))))) / -2.25) > -2.25), F2: v4}
    for fi26 := 0; fi26 < 14; fi26++ {
        _ = len(v5)
        _ = fi26
    }
    v21 -= v21
    if false {
        if (0 > 255) {
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 uint16 = 1
                var v29 uint16 = v28
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        switch 32768 {
        case 2147483647:
            v21 = v21
        case 4294967295:
            v1["\n\t"] = 74
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
