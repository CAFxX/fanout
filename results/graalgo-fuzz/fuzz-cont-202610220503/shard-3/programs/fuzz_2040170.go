package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 string = ""
    v2 := []string{}
    v3 := []string{"\n\t", "a", v1, ((((v1 + (v1 + "foo bar")) + v1) + v1) + "")}
    var v4 int64 = -9223372036854775808
    var v5 bool = (2 < 2)
    var v6 uint16 = 789
    if (-9223372036854775808 < v4) {
        var v7 int8 = -1
        v8 := GenS1{F0: "\n\t", F1: 2, F2: v7}
        fmt.Println(v6)
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: -36, F1: 1}
    for fi10 := 0; fi10 < 11; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v11 uint16 = v6
        fmt.Println(v4, v5)
        _ = fi10
        _ = v11
    }
    v4 -= 42
    if false {
        fmt.Println(v5, v4, v6)
        switch 42 {
        case 127:
            var v12 uint64 = 2
            _ = v12
        case 32767:
            var v13 int64 = v4
            _ = v13
        case -560:
            for fi14 := 0; fi14 < 13; fi14++ {
                _ = len(v3)
                var v15 string = ((v1 + "\n\t") + v1)
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(v5)
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            switch fi16 {
            case 42:
                var v17 float64 = 3.14159
                _ = v17
            case 128:
                v4 += v4
            case -1:
                v18 := GenS2{F0: -128, F1: 608}
                _ = v18
            default:
                v1 = v1
            }
            var v19 int16 = 42
            var v20 uint = 255
            _ = fi16
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: v1, F1: 2, F2: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint32 = 31
        _ = fi22
        _ = v23
    }
    var v24 uint = 255
    v6 -= (v6 + v6)
    if (!v5) {
        switch ("hello" + "foo bar") {
        case "hello":
            v25 := GenS2{F0: 39, F1: -109}
            _ = v25
        case "foo bar":
            v6 = (^(v6))
        case "":
            fmt.Println(v4, v1, v24)
        }
        v1 = "\n\t"
    }
    var v26 bool = (33 > 89)
    var v27 uint = (v24 & v24)
    if true {
        if len(v3) > 0 {
            v3[0] = v1
        }
    }
    v27 += (v27 * (v27 >> ((v24) & 63)))
    var v28 int16 = -39
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
