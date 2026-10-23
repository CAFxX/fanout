package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{("hello" + ("foo bar" + ((("hello" + (("\n\t" + "a") + "a")) + ("" + "")) + (("hello" + "\n\t") + "")))), "hello", ("\n\t" + ("\n\t" + ("" + ("\n\t" + ("\n\t" + ""))))), (("foo bar" + "foo bar") + "")}
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: 0}
    if len(v2) > 0 {
        v2[0] = ""
    }
    v5 := GenS1{F0: 255}
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 float64 = (-2.25 + float64((0.1 / (0.1 * 3.14159))))
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS2{F0: 212}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint = 4294967295
                switch (fi10 / ((fi10) | 1)) {
                case -2147483648:
                    var v12 int = 102
                    _ = v12
                }
                _ = len(v3)
                _ = fi10
                _ = v11
            }
            _ = fi6
            _ = v7
        }
    case "foo bar":
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 bool = ((false && (true || true)) && (!false))
            v14 = v14
            _ = fi13
            _ = v14
        }
    case "":
        v15 := GenS3{F0: 1, F1: 2, F2: 2147483648}
        _ = v15
    default:
        var v16 int8 = 1
        _ = v16
    }
    v17 := GenS3{F0: 255, F1: 255, F2: 1}
    switch -1 {
    case 32767:
        if (0 > -1) {
            for fi18 := 0; fi18 < 13; fi18++ {
                v1["foo bar"] = "foo bar"
                _ = fi18
            }
        }
    case 32768:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) > 0 {
        v2[0] = ((((("" + "") + "foo bar") + "a") + "hello") + ("foo bar" + ((("hello" + "") + ("foo bar" + "hello")) + ("a" + "hello"))))
    }
    var v19 float64 = (float64(0.1) + (1e10 * -2.25))
    v19 -= 3.14159
    v20 := GenS4{F0: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
