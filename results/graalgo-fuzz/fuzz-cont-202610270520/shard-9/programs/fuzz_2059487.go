package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    var v3 int32 = 42
    var v4 bool = (true && (35 != 20))
    v5 := make(map[string]string)
    v6 := GenS1{F0: 1000, F1: 2147483647, F2: "hello"}
    switch "" {
    case "foo bar":
        v2 = v2
    case "hello":
        switch "a" {
        case "\n\t":
            if (((("" + "hello") + "foo bar") + ("foo bar" + ("hello" + "\n\t"))) < "foo bar") {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v7 := GenS2{F0: v3, F1: -123, F2: true}
                _ = v7
            }
        case "hello":
            if (!v4) {
                var v8 int8 = -128
                v4 = false
                _ = v8
            }
        case "":
            if ("foo bar" == ("\n\t" + "\n\t")) {
                var v9 int16 = 42
                var v10 int = 9223372036854775807
                _ = v10
                _ = v9
            } else {
                v3 = v3
                var v11 uint8 = 2
                var v12 uint = 702
                _ = v11
                _ = v12
            }
        }
    }
    if ((v4 && true) || (83 > 111)) {
        var v13 float32 = float32(1e10)
        fmt.Println(v3, v2)
        v1 = append(v1, (("\n\t" + "\n\t") + ("foo bar" + "a")))
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 813, F2: ""}
    v15 := GenS1{F0: -1, F1: 68, F2: ("hello" + "\n\t")}
    v4 = v4
    switch ("\n\t" + "a") {
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS3{F0: -1, F1: true, F2: 0}
            _ = fi16
            _ = v17
        }
    case "foo bar":
        switch ("" + "hello") {
        case "hello":
            v18 := GenS1{F0: 1, F1: 1, F2: "foo bar"}
            _ = v18
        default:
            switch "hello" {
            case "":
                _ = len(v1)
            case "a":
                var v19 uint8 = 1
                _ = v19
            case "foo bar":
                fmt.Println(v2, v4, v2)
            }
        }
    default:
        var v20 uint = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int8 = 42
        _ = fi21
        _ = v22
    }
    fmt.Println(v2, v3)
    var v23 bool = ((-(v3)) != (v3 << ((2) & 63)))
    for fi24 := 0; fi24 < 12; fi24++ {
        v4 = false
        var v25 float64 = -2.25
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
