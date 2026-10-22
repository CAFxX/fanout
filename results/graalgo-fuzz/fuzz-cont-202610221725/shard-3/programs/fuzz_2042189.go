package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{0}
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 100, F1: 42}
    delete(v1, "foo bar")
    switch ("\n\t" + "\n\t") {
    case "a":
        var v5 int16 = -32768
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 18; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    }
    if (((3.14159 / 1e10) == float64(1.5)) && (false || (!false))) {
        var v7 uint8 = 174
        _ = v7
    }
    var v8 int16 = -1
    if true {
        v1["a"] = false
        v9 := GenS1{F0: 1, F1: 65535}
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v10 int16 = (v8 ^ v8)
            _ = v10
        }
        _ = v9
    }
    var v11 int = 71
    switch "\n\t" {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        v11 = (v11 | v11)
    case "foo bar":
        var v12 float64 = -2.25
        _ = v12
    default:
        fmt.Println(v11, v11)
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        if (!(255 >= 4294967295)) {
            switch fi13 {
            case 32767:
                var v14 int = fi13
                _ = v14
            case 568:
                v8 = 0
            }
        } else {
            v8 -= (-(v8))
            var v15 uint16 = 954
            _ = v15
        }
        fmt.Println(v8)
        _ = fi13
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(fi16, v8)
            v18 := v3[fi17]
            fmt.Println(fi16, v8)
            var v19 int64 = 653
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v20 := GenS2{F0: v8, F1: 128, F2: 42}
    fmt.Println(v8)
    switch -1 {
    case -9223372036854775808:
        _ = len(v3)
    case -410:
        v8 = v8
    }
    switch v11 {
    case 42:
        v8 -= v8
    case 4294967295:
        if (0 >= 155) {
            v11 = v11
        } else {
            v21 := GenS1{F0: 127, F1: 738}
            _ = v21
        }
    case 1:
        var v22 float64 = 3.14159
        _ = v22
    }
    v23 := v1[("hello" + "")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v8
}
