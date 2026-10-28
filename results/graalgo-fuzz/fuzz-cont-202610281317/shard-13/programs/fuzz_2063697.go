package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int16 }
func genFn0(p0 int) uint64 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 2147483647}
        if (false && (!true)) {
            v7 := []int{p0, 0}
            if false {
                fmt.Println(p0)
            }
            _ = v7
        }
        v8 := []int{4294967295, (^(fi5)), p0}
        _ = fi5
        _ = v6
        _ = v8
    }
    var v9 float64 = float64(1e10)
    v9 = (0.1 * 1e10)
    fmt.Println(v9)
    v10 := GenS1{F0: 337}
    _ = p0
    _ = v10
    _ = v9
    return 177
}
func genFn1(p0 uint8) uint32 {
    var v11 int = -341
    if (!false) {
        var v12 uint8 = p0
        v11 = 42
        switch (v11 & v11) {
        case 761:
            fmt.Println(v12, p0)
        case -246:
            var v13 int16 = 789
            _ = v13
        }
        _ = v12
    } else {
        v14 := GenS2{F0: 541, F1: 507}
        p0 -= (^(p0))
        _ = v14
    }
    v15 := GenS3{F0: 317}
    var v16 uint32 = 2147483648
    _ = p0
    _ = v11
    _ = v15
    _ = v16
    return (^(v16))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    v3 := []int{75, 417}
    v4 := []string{"foo bar", ("a" + "a"), "hello"}
    v2 -= float64(-2.25)
    if (2147483647 != 829) {
        var v17 int32 = 0
        fmt.Println(v2, v17)
        _ = v17
    }
    v2 = v2
    var v18 uint16 = 0
    fmt.Println(v2)
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS4{F0: 2147483648, F1: ("hello" + (("a" + "hello") + ("a" + "foo bar")))}
        var v21 float64 = (0.1 - -2.25)
        switch fi19 {
        case -2147483648:
            var v22 int8 = 1
            _ = v22
        case 0:
            var v23 int64 = 0
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        var v24 uint64 = 18446744073709551615
        switch -432 {
        case 9223372036854775807:
            v25 := GenS5{F0: 42}
            _ = v25
        default:
            v24 -= (v24 - v24)
        }
        var v26 uint64 = v24
        _ = v24
        _ = v26
    } else {
        delete(v1, 4294967295)
        v2 -= 1e10
    }
    var v27 uint64 = 4294967295
    v3 = append(v3, -596)
    switch ("\n\t" + ("" + "\n\t")) {
    case "\n\t":
        v4 = append(v4, "hello")
    case "":
        var v28 int32 = 0
        _ = v28
    }
    fmt.Println(v18)
    v18 = 65535
    fmt.Println(v18, v18, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
