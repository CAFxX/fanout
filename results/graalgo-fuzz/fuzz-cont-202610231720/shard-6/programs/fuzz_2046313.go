package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0() int8 {
    var v7 string = ("\n\t" + "hello")
    var v8 int8 = 100
    fmt.Println(v8, v8)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float32 = float32(-2.25)
        var v11 string = v7
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int8 = 127
            var v14 int32 = 1
            fmt.Println(v13, v13, v14)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 = ("\n\t" + (v7 + v7))
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8, p1 uint32) int16 {
    fmt.Println(p0, p1)
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 string = ("foo bar" + "a")
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v15
    return -794
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2, 2147483648, 1}
    var v3 int8 = 24
    v4 := make(map[string]uint32)
    var v5 uint16 = 255
    var v6 string = ("foo bar" + "foo bar")
    if true {
        v3 = (v3 | v3)
        fmt.Println(v6, v5)
        switch -562 {
        case -321:
            v6 = v6
        }
    }
    var v18 uint8 = 255
    fmt.Println(v5)
    _ = len(v1)
    if (!false) {
        v3 = (v3 % (((v3 - v3)) | 1))
        fmt.Println(v6, v18)
    } else {
        var v19 int = 4294967295
        v20 := GenS1{F0: 1000}
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS2{F0: v6, F1: 1}
            if ((((3.14159 + (((float32(0.1) * 1e10) / -2.25) - 3.14159)) / 0.1) != 0.1) || false) {
                switch v6 {
                case "hello":
                    v23 := GenS1{F0: 1000}
                    _ = v23
                case "foo bar":
                    v19 += 4294967295
                default:
                    var v24 int16 = 42
                    _ = v24
                }
            }
            v1[(v6 + (v6 + (v6 + (((("" + "") + v6) + "\n\t") + ("" + "")))))] = (genFn0() < v3)
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(v5)
            var v27 int16 = -1
            _ = fi26
            _ = v27
        }
        if (-523 < 457) {
            for fi28 := 0; fi28 < 18; fi28++ {
                v18 = v18
                var v29 uint64 = 0
                v5 = (v5 & v5)
                _ = fi28
                _ = v29
            }
        } else {
            v18 += (v18 & (v18 | (v18 ^ v18)))
            var v30 string = "foo bar"
            var v31 uint8 = (v18 >> ((746) & 63))
            _ = v30
            _ = v31
        }
        v32 := GenS1{F0: -58}
        _ = fi25
        _ = v32
    }
    var v33 int16 = -32768
    fmt.Println(v33, v18, v6)
    if true {
        v34 := GenS2{F0: "\n\t", F1: 690}
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
