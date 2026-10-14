package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{"hello", (("foo bar" + (("a" + ((((("foo bar" + ("\n\t" + "a")) + "\n\t") + "\n\t") + "\n\t") + ("foo bar" + (("hello" + "foo bar") + "foo bar")))) + ((("" + "a") + "a") + "foo bar"))) + ("hello" + "foo bar")), ("foo bar" + (("\n\t" + "a") + ""))}
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    v4 := []string{}
    v5 := make(map[int]int)
    var v6 int32 = 0
    v7 := GenS1{F0: 924, F1: 1}
    for fi8 := 0; fi8 < 16; fi8++ {
        v3 = v3
        for fi9 := 0; fi9 < 1; fi9++ {
            switch fi8 {
            case 32767:
                v10 := GenS1{F0: 585, F1: 0}
                _ = v10
            case 0:
                v11 := GenS1{F0: 1, F1: 0}
                _ = v11
            }
            _ = fi9
        }
        _ = fi8
    }
    var v12 int32 = (v6 - (^(v6)))
    switch "a" {
    case "hello":
        var v13 float32 = ((1e10 * (-2.25 - float32(float64(((((-2.25 * (3.14159 / -2.25)) + 3.14159) / float32(3.14159)) * 1.5))))) * (0.1 / (1.5 / 0.1)))
        _ = v13
    case "a":
        var v14 uint32 = 42
        _ = v14
    case "":
        v3 -= (v3 >> ((710) & 63))
    default:
        v6 = (-(v12))
    }
    fmt.Println(v3, v6, v3)
    v3 = (v3 + -689)
    switch ("a" + "\n\t") {
    case "a":
        var v15 string = "hello"
        _ = v15
    case "\n\t":
        v16 := GenS1{F0: 42, F1: 18446744073709551615}
        _ = v16
    }
    switch -411 {
    case -901:
        var v17 float64 = 0.1
        _ = v17
    case -9223372036854775808:
        delete(v2, 783)
    case 42:
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint32 = 87
            v20 := GenS1{F0: v19, F1: 1}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    switch 1 {
    case 32768:
        v21 := GenS2{F0: 132, F1: 243, F2: 42}
        _ = v21
    case 4294967295:
        if (!true) {
            var v22 int8 = 38
            var v23 uint32 = 76
            var v24 uint = 4294967295
            _ = v22
            _ = v23
            _ = v24
        } else {
            var v25 float64 = 1.5
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS3{F0: v6}
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
