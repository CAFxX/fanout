package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 string; F1 int; F2 int }

func main() {
    var v1 uint64 = 586
    var v2 int8 = 42
    var v3 float64 = 1.5
    v3 -= v3
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float64 = ((v3 / v3) / v3)
        if ((("" + ("hello" + "hello")) <= ("foo bar" + "hello")) && false) {
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := GenS1{F0: 206, F1: 128, F2: -1}
                v5 += float64(1.5)
                if false {
                    v3 = float64(((1.5 - ((1e10 / ((-2.25 / (1.5 * 3.14159)) / 1.5)) - (1.5 - (1e10 - 3.14159)))) / -2.25))
                }
                _ = fi6
                _ = v7
            }
            var v8 uint = 0
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float64 = 1e10
                v11 := []int{32767, fi4, 32768}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint8 = 42
            _ = v12
        }
        _ = fi4
        _ = v5
    }
    var v13 int16 = 32767
    v3 -= (((v3 / v3) - (float64(float32(v3)) + 1.5)) * ((v3 + ((1.5 * v3) + float64(-2.25))) + v3))
    if true {
        v14 := []int{555, 2147483647, 748, 0}
        v1 -= 602
        v1 += v1
        _ = v14
    } else {
        fmt.Println(v3, v13, v2)
        switch ((("" + "\n\t") + ("hello" + "\n\t")) + (("" + ("foo bar" + "foo bar")) + "foo bar")) {
        case "hello":
            v15 := GenS1{F0: 1, F1: 156, F2: (v2 >> ((650) & 63))}
            _ = v15
        case "foo bar":
            var v16 int = 884
            _ = v16
        case "\n\t":
            fmt.Println(v2, v3)
        default:
            var v17 int8 = v2
            _ = v17
        }
        var v18 uint32 = 1
        _ = v18
    }
    v19 := []int{-2147483648, -9223372036854775808, 874, 906}
    var v20 int32 = -273
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    switch (("" + "") + (("\n\t" + "foo bar") + "foo bar")) {
    case "\n\t":
        v21 := GenS2{F0: ("a" + ("foo bar" + "")), F1: -2147483648, F2: 20}
        _ = v21
    case "":
        var v22 uint16 = 5
        _ = v22
    }
    var v23 uint8 = 127
    for fi24 := 0; fi24 < 3; fi24++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = len(v19)
        _ = fi24
    }
    var v25 float64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v3
}
