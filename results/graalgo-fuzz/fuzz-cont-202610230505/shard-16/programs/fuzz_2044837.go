package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 int32) bool {
    p0 = p0
    var v7 uint32 = 324
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 int32) uint8 {
    v8 := []int{-127, -378, 2147483647, 9223372036854775807}
    fmt.Println(p0)
    v8 = append(v8, -2147483648)
    _ = p0
    _ = v8
    return 39
}

func main() {
    v1 := []string{"hello", ("\n\t" + "")}
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    v4 := make(map[string]bool)
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    if (("a" + "hello") < ("" + ((("\n\t" + "foo bar") + "hello") + ""))) {
        switch "hello" {
        case "foo bar":
            switch -2147483648 {
            case 128:
                var v9 float32 = 1e10
                _ = v9
            case 772:
                v3 = (v3 - 663)
            default:
                var v10 int8 = 49
                _ = v10
            }
        case "":
            var v11 int64 = v3
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 5; fi12++ {
                v3 = v3
                var v13 float32 = 0.0
                _ = fi12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint16 = 32767
            _ = fi14
            _ = v15
        }
        v16 := v4["foo bar"]
        _ = v16
    }
    var v17 float32 = -2.25
    var v18 bool = (1 < 2)
    if (((v18 || v18) || false) && (!(!v18))) {
        if len(v1) > 0 {
            v1[0] = ""
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v17, fi19)
            _ = fi19
        }
    }
    switch ("a" + ((("hello" + "") + ("a" + "foo bar")) + "hello")) {
    case "a":
        for fi20 := 0; fi20 < 2; fi20++ {
            v3 = v3
            _ = len(v5)
            var v21 int64 = (-(v3))
            _ = fi20
            _ = v21
        }
    case "":
        var v22 uint8 = 255
        _ = v22
    }
    v1 = append(v1, "foo bar")
    switch "" {
    case "\n\t":
        var v23 bool = (1 <= 1)
        _ = v23
    case "foo bar":
        for fi24 := 0; fi24 < 18; fi24++ {
            if (101 <= 1) {
                switch "foo bar" {
                case "hello":
                    v25 := GenS1{F0: ("foo bar" + "hello")}
                    _ = v25
                case "":
                    v3 = (v3 - (v3 | v3))
                default:
                    v18 = v18
                }
                v26 := GenS2{F0: 0, F1: (v3 & (v3 | v3)), F2: (-2.25 != ((1e10 + 0.0) / ((0.1 * 1.5) + (3.14159 / 1e10))))}
                v17 -= v17
                _ = v26
            }
            v27 := GenS1{F0: ""}
            fmt.Println(v18)
            _ = fi24
            _ = v27
        }
    case "a":
        v3 = 0
    }
    var v28 string = ("foo bar" + "\n\t")
    v28 = ("\n\t" + "")
    fmt.Println(v18, v17)
    v29 := GenS3{F0: true, F1: -32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
