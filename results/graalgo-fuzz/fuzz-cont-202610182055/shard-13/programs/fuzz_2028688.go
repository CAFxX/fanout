package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 bool) string {
    if (p0 <= 0) {
        var v7 bool = (p1 || (144 < 645))
        v8 := GenS1{F0: 32767}
        _ = v7
        _ = v8
    } else {
        var v9 bool = (-178 <= 0)
        var v10 uint = (p0 & 4294967295)
        v9 = false
        _ = v10
        _ = v9
    }
    var v11 uint32 = 4294967295
    v12 := GenS2{F0: 42, F1: 1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return "a"
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{}
    var v3 bool = false
    var v4 int64 = 0
    v5 := make(map[uint32]string)
    var v6 bool = v3
    v6 = v3
    if false {
        var v13 uint16 = 2
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 bool = (1e10 != (-2.25 / 0.1))
            var v16 uint16 = (v13 * 1)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        switch ((("\n\t" + ((("\n\t" + ("" + "a")) + "a") + "\n\t")) + "") + (("a" + ("" + "hello")) + ("foo bar" + "foo bar"))) {
        case "":
            v6 = (!((!false) && (149 >= 255)))
        case "foo bar":
            fmt.Println(v3)
        case "hello":
            for fi17 := 0; fi17 < 20; fi17++ {
                v4 += 42
                _ = fi17
            }
        }
    }
    if v6 {
        var v18 uint64 = (v1 ^ (v1 % ((v1) | 1)))
        v3 = v6
        _ = v18
    } else {
        v4 = v4
    }
    var v19 int32 = 42
    var v20 float32 = (1.5 + 1e10)
    v21 := GenS1{F0: 9223372036854775807}
    v1 -= 4294967295
    v22 := GenS3{F0: 1, F1: -1}
    var v23 float64 = 1.5
    if (((v19 == (v19 >> ((255) & 63))) && (true || v6)) && (true || v3)) {
        v24 := GenS1{F0: 128}
        var v25 float64 = (1e10 - float64(v20))
        v26 := GenS4{F0: (v6 && v6)}
        _ = v24
        _ = v25
        _ = v26
    } else {
        var v27 uint = 823
        v28 := GenS4{F0: true}
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
