package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0() uint64 {
    var v7 float32 = -2.25
    v7 = v7
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 32767}
        _ = fi8
        _ = v9
    }
    v10 := []int{-680}
    _ = v10
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{71, 2}
    var v2 uint64 = 0
    v3 := make(map[uint32]uint32)
    v4 := []string{("\n\t" + "foo bar"), "a", "", ""}
    v5 := make(map[uint32]bool)
    var v6 int = 537
    switch v6 {
    case -9223372036854775808:
        fmt.Println(v2, v6)
    case -1:
        var v11 bool = (!true)
        _ = v11
    }
    if len(v4) > 0 {
        v4[0] = ((("\n\t" + ("hello" + "\n\t")) + ("" + "a")) + (("" + "hello") + ("a" + "foo bar")))
    }
    switch "hello" {
    case "foo bar":
        var v12 int16 = 1000
        _ = v12
    case "":
        var v13 int64 = 437
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 string = ("hello" + ("hello" + ("\n\t" + "")))
        fmt.Println(fi14, v15)
        v16 := GenS2{F0: 1, F1: v2, F2: 0}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: 1000, F1: 42, F2: 65535}
    var v18 uint64 = (v2 * (v2 - (v2 * genFn0())))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v19 float64 = ((-2.25 - 1.5) - (((3.14159 - (1.5 + ((-2.25 - (1e10 - -2.25)) / 0.1))) + (float64(float32(0.1)) * -2.25)) + 1.5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
