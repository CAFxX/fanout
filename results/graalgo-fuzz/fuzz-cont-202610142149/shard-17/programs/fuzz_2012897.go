package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 string }

func main() {
    v1 := []string{("hello" + ""), "a", "a"}
    v2 := make(map[uint32]bool)
    var v3 bool = (!true)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 float64 = float64(3.14159)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        delete(v2, 745)
        if v3 {
            var v7 uint8 = 127
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v3)
    if (1 <= 42) {
        var v8 uint64 = 937
        for fi9 := 0; fi9 < 15; fi9++ {
            _ = len(v2)
            for fi10 := 0; fi10 < 19; fi10++ {
                v8 = v8
                _ = fi10
            }
            var v11 int = (fi9 - 128)
            _ = fi9
            _ = v11
        }
        _ = v8
    }
    v3 = (32768 == -9223372036854775808)
    v2[109] = v3
    fmt.Println(v3)
    if v3 {
        switch "hello" {
        case "":
            if (65535 == 135) {
                v3 = (v3 || (-368 < 1000))
            }
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            var v12 int16 = 152
            _ = v12
        default:
            v3 = (0.0 <= 0.1)
        }
        fmt.Println(v3)
        v3 = (-128 < 61)
    } else {
        v13 := GenS1{F0: 0, F1: "hello"}
        _ = v13
    }
    v14 := GenS2{F0: 2, F1: 127, F2: 2}
    switch 119 {
    case -845:
        switch -602 {
        case -630:
            var v15 bool = (760 <= 585)
            _ = v15
        case 9223372036854775807:
            v16 := GenS2{F0: 6, F1: 255, F2: 2}
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint = 362
            _ = fi17
            _ = v18
        }
    }
    var v19 int16 = 1
    if v3 {
        v20 := GenS3{F0: 2, F1: 100}
        for fi21 := 0; fi21 < 4; fi21++ {
            v19 = (v19 - v19)
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 uint64 = 18446744073709551615
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    }
    if true {
        delete(v2, 4294967295)
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v3, fi24, v19)
            _ = fi24
        }
        var v25 int8 = -39
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint = 65535
        var v28 int16 = 0
        if ((true && false) || false) {
            v29 := GenS1{F0: 42, F1: "hello"}
            var v30 uint16 = 65535
            v31 := GenS3{F0: v27, F1: -114}
            _ = v29
            _ = v30
            _ = v31
        } else {
            v32 := GenS4{F0: 4294967295, F1: "hello"}
            _ = v32
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
